package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"wichitaradar/internal/cache"
	"wichitaradar/internal/middleware"
	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// Severity ranks for picking the worst active condition.
const (
	sevNone     = 0
	sevAdvisory = 1
	sevWatch    = 2
	sevWarning  = 3
	sevFlash    = 4
)

// floodGauge describes one NWS/NWPS river gauge we display.
type floodGauge struct {
	ID   string  // NWPS LID, e.g. "ICTK1"
	Name string  // Human-readable label
	Lat  float64 // Decimal degrees
	Lon  float64
}

// gaugesForWichita lists NWPS gauges in and around Sedgwick County (Arkansas
// River, Little Arkansas, Cowskin, Ninnescah, Walnut basins). NWS does not
// operate a flood-forecast gauge directly in Wichita on the Arkansas River
// (Cheney Reservoir regulates upstream flow); upstream/downstream gauges plus
// the Cowskin and Little Arkansas give the best picture of city-relevant flood
// conditions. Coordinates come from the NWPS gauge metadata.
var gaugesForWichita = []floodGauge{
	{ID: "COWK1", Name: "Cowskin Creek at 119th St W, Wichita", Lat: 37.70169, Lon: -97.48055},
	{ID: "DRBK1", Name: "Arkansas River at Derby", Lat: 37.54431, Lon: -97.27557},
	{ID: "MULK1", Name: "Arkansas River near Mulvane", Lat: 37.4755, Lon: -97.2613},
	{ID: "HAVK1", Name: "Arkansas River near Haven", Lat: 37.94614, Lon: -97.77512},
	{ID: "SEDK1", Name: "Little Arkansas River near Sedgwick", Lat: 37.88305, Lon: -97.42416},
	{ID: "HTDK1", Name: "Little Arkansas River near Halstead", Lat: 38.02853, Lon: -97.54054},
	{ID: "ALMK1", Name: "Little Arkansas River at Alta Mills", Lat: 38.1122, Lon: -97.59194},
	{ID: "BLPK1", Name: "Ninnescah River near Belle Plaine", Lat: 37.3916, Lon: -97.3391},
	{ID: "PECK1", Name: "Ninnescah River near Peck", Lat: 37.4569, Lon: -97.4236},
	{ID: "AGAK1", Name: "Walnut River at Augusta", Lat: 37.67056, Lon: -96.98},
}

// Alert is a normalized NWS flood alert ready for display.
type Alert struct {
	Event        string
	Headline     string
	Severity     string
	Description  string
	AreaDesc     string
	Expires      string
	Sent         string
	SeverityRank int
}

// Gauge is a normalized gauge reading ready for display.
type Gauge struct {
	ID         string
	Name       string
	Lat        float64
	Lon        float64
	Stage      string // Current observed stage, formatted (e.g., "4.21 ft")
	Status     string // None / Action / Minor / Moderate / Major / Unknown
	StatusRank int    // For sorting and worst-status computation
	FloodStage string // e.g., "13.0 ft" (or "" if not provided)
	Observed   string // Observation time
	Trend      string // "Rising" / "Falling" / "Steady" / ""
}

// Closure is a normalized active road closure entry (flood or storm related)
// sourced from Sedgwick County's public road-closure FeatureService.
type Closure struct {
	Street      string
	Reason      string // "Flooding" / "Storm Damage"
	Description string
	City        string
	Begin       string
	End         string
	AltRoute    string
	StartedAt   string
	Lat         float64
	Lon         float64
}

// FloodingData is the template payload.
type FloodingData struct {
	Menu            *menu.Menu
	CurrentPath     string
	RefreshInterval int

	BannerLevel       string // "none" / "advisory" / "watch" / "warning" / "flash"
	BannerText        string
	Alerts            []Alert
	Gauges            []Gauge
	Closures          []Closure
	AlertsAvailable   bool
	GaugesAvailable   bool
	ClosuresAvailable bool

	AlertsGeoJSON   template.JS
	GaugesGeoJSON   template.JS
	ClosuresGeoJSON template.JS

	GeneratedAt string
}

// --- NWS alerts -----------------------------------------------------------

type nwsAlertsResponse struct {
	Features []json.RawMessage `json:"features"`
}

type nwsAlertFeature struct {
	Geometry   json.RawMessage `json:"geometry"`
	Properties struct {
		Event       string    `json:"event"`
		Headline    string    `json:"headline"`
		Severity    string    `json:"severity"`
		Description string    `json:"description"`
		AreaDesc    string    `json:"areaDesc"`
		Sent        time.Time `json:"sent"`
		Expires     time.Time `json:"expires"`
	} `json:"properties"`
}

func severityRank(event string) int {
	e := strings.ToLower(event)
	switch {
	case strings.Contains(e, "flash flood warning"):
		return sevFlash
	case strings.Contains(e, "flood warning"):
		return sevWarning
	case strings.Contains(e, "flood watch"), strings.Contains(e, "flash flood watch"):
		return sevWatch
	case strings.Contains(e, "flood advisory"), strings.Contains(e, "hydrologic"):
		return sevAdvisory
	}
	return sevNone
}

func parseNWSAlerts(r io.Reader) ([]Alert, json.RawMessage, error) {
	var resp nwsAlertsResponse
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, err
	}

	loc, _ := time.LoadLocation("America/Chicago")
	alerts := make([]Alert, 0, len(resp.Features))
	geomFeatures := make([]json.RawMessage, 0, len(resp.Features))

	for _, raw := range resp.Features {
		var f nwsAlertFeature
		if err := json.Unmarshal(raw, &f); err != nil {
			continue
		}
		rank := severityRank(f.Properties.Event)
		if rank == sevNone {
			continue
		}
		expires := ""
		if !f.Properties.Expires.IsZero() {
			expires = f.Properties.Expires.In(loc).Format("Mon 3:04 PM MST")
		}
		sent := ""
		if !f.Properties.Sent.IsZero() {
			sent = f.Properties.Sent.In(loc).Format("Mon 3:04 PM MST")
		}
		alerts = append(alerts, Alert{
			Event:        f.Properties.Event,
			Headline:     f.Properties.Headline,
			Severity:     f.Properties.Severity,
			Description:  f.Properties.Description,
			AreaDesc:     f.Properties.AreaDesc,
			Expires:      expires,
			Sent:         sent,
			SeverityRank: rank,
		})
		geomFeatures = append(geomFeatures, raw)
	}

	sort.SliceStable(alerts, func(i, j int) bool {
		return alerts[i].SeverityRank > alerts[j].SeverityRank
	})

	geo := struct {
		Type     string            `json:"type"`
		Features []json.RawMessage `json:"features"`
	}{
		Type:     "FeatureCollection",
		Features: geomFeatures,
	}
	geoJSON, _ := json.Marshal(geo)
	return alerts, geoJSON, nil
}

// --- NWPS gauges ----------------------------------------------------------

// nwpsGaugeResponse models the subset of NWPS /gauges/{id} we use. The API
// returns flood-stage thresholds (nested under flood.categories) and the most
// recent observation.
type nwpsGaugeResponse struct {
	LID    string `json:"lid"`
	Name   string `json:"name"`
	Status struct {
		Observed struct {
			Primary       float64 `json:"primary"`
			PrimaryUnit   string  `json:"primaryUnit"`
			ValidTime     string  `json:"validTime"`
			FloodCategory string  `json:"floodCategory"`
		} `json:"observed"`
	} `json:"status"`
	Flood struct {
		StageUnits string `json:"stageUnits"`
		Categories struct {
			Action struct {
				Stage float64 `json:"stage"`
			} `json:"action"`
			Minor struct {
				Stage float64 `json:"stage"`
			} `json:"minor"`
			Moderate struct {
				Stage float64 `json:"stage"`
			} `json:"moderate"`
			Major struct {
				Stage float64 `json:"stage"`
			} `json:"major"`
		} `json:"categories"`
	} `json:"flood"`
}

func statusRank(s string) int {
	switch strings.ToLower(s) {
	case "major":
		return 4
	case "moderate":
		return 3
	case "minor":
		return 2
	case "action":
		return 1
	case "":
		return 0
	case "no flooding", "none":
		return 0
	}
	return 0
}

func categorize(stage float64, c nwpsGaugeResponse) string {
	switch {
	case c.Flood.Categories.Major.Stage > 0 && stage >= c.Flood.Categories.Major.Stage:
		return "Major"
	case c.Flood.Categories.Moderate.Stage > 0 && stage >= c.Flood.Categories.Moderate.Stage:
		return "Moderate"
	case c.Flood.Categories.Minor.Stage > 0 && stage >= c.Flood.Categories.Minor.Stage:
		return "Minor"
	case c.Flood.Categories.Action.Stage > 0 && stage >= c.Flood.Categories.Action.Stage:
		return "Action"
	}
	return "None"
}

func parseNWPSGauge(r io.Reader, g floodGauge) Gauge {
	var resp nwpsGaugeResponse
	body, err := io.ReadAll(r)
	if err != nil {
		return Gauge{ID: g.ID, Name: g.Name, Lat: g.Lat, Lon: g.Lon, Status: "Unknown"}
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return Gauge{ID: g.ID, Name: g.Name, Lat: g.Lat, Lon: g.Lon, Status: "Unknown"}
	}

	status := resp.Status.Observed.FloodCategory
	if status == "" || strings.EqualFold(status, "no_flooding") {
		status = categorize(resp.Status.Observed.Primary, resp)
	}
	// Normalize casing: NWPS sometimes returns "minor" / "no_flooding".
	status = titleCase(strings.ReplaceAll(strings.ToLower(status), "_", " "))
	if status == "" || status == "No Flooding" {
		status = "None"
	}

	// NWPS uses -999 as a sentinel for "no current observation".
	stage := ""
	primary := resp.Status.Observed.Primary
	if primary > -900 && (primary != 0 || resp.Status.Observed.PrimaryUnit != "") {
		unit := resp.Status.Observed.PrimaryUnit
		if unit == "" {
			unit = "ft"
		}
		stage = fmt.Sprintf("%.2f %s", primary, unit)
	} else {
		status = "Unknown"
	}

	floodStage := ""
	unit := resp.Flood.StageUnits
	if unit == "" {
		unit = "ft"
	}
	if resp.Flood.Categories.Minor.Stage > 0 {
		floodStage = fmt.Sprintf("%.1f %s", resp.Flood.Categories.Minor.Stage, unit)
	}

	observed := ""
	if t, err := time.Parse(time.RFC3339, resp.Status.Observed.ValidTime); err == nil {
		if loc, err := time.LoadLocation("America/Chicago"); err == nil {
			observed = t.In(loc).Format("Mon 3:04 PM MST")
		} else {
			observed = t.Format("Mon 15:04 MST")
		}
	}

	return Gauge{
		ID:         g.ID,
		Name:       g.Name,
		Lat:        g.Lat,
		Lon:        g.Lon,
		Stage:      stage,
		Status:     status,
		StatusRank: statusRank(status),
		FloodStage: floodStage,
		Observed:   observed,
	}
}

func titleCase(s string) string {
	b := []byte(s)
	upper := true
	for i, c := range b {
		if c == ' ' {
			upper = true
			continue
		}
		if upper && c >= 'a' && c <= 'z' {
			b[i] = c - 32
		}
		upper = false
	}
	return string(b)
}

func gaugesGeoJSON(gauges []Gauge) []byte {
	type feature struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties map[string]any `json:"properties"`
	}
	type collection struct {
		Type     string    `json:"type"`
		Features []feature `json:"features"`
	}
	c := collection{Type: "FeatureCollection"}
	for _, g := range gauges {
		f := feature{Type: "Feature"}
		f.Geometry.Type = "Point"
		f.Geometry.Coordinates = []float64{g.Lon, g.Lat}
		f.Properties = map[string]any{
			"id":         g.ID,
			"name":       g.Name,
			"stage":      g.Stage,
			"status":     g.Status,
			"floodStage": g.FloodStage,
			"observed":   g.Observed,
		}
		c.Features = append(c.Features, f)
	}
	out, _ := json.Marshal(c)
	return out
}

// --- Sedgwick County road closures ---------------------------------------

// closuresQueryURL is the Sedgwick County Public Works road-closure
// FeatureServer query, filtered to currently-in-progress flood and
// storm-damage closures (Status=2 means in progress per the layer's coded
// domain; 1=planned, 3=completed). Coverage includes the city of Wichita.
const closuresQueryURL = "https://services7.arcgis.com/McLat6HlPl45bNBv/arcgis/rest/services/Public_PW_RoadClosures_WGS84/FeatureServer/0/query" +
	"?where=Status%3D2+AND+subtype+IN+%28%27ROAD_CLOSED_FLOOD%27%2C%27ROAD_CLOSED_STORM_DAMAGE%27%29" +
	"&outFields=street%2Csubtype%2Cdescription%2Cdirection%2Cstarttime%2Caltroute%2CCity%2CBeginBoundary%2CEndBoundary" +
	"&returnGeometry=true&outSR=4326&f=json"

// subtypeLabel maps the layer's coded subtype values to display labels.
var subtypeLabel = map[string]string{
	"ROAD_CLOSED_FLOOD":        "Flooding",
	"ROAD_CLOSED_STORM_DAMAGE": "Storm Damage",
}

type arcgisClosureResponse struct {
	Features []struct {
		Attributes struct {
			Street        string `json:"street"`
			Subtype       string `json:"subtype"`
			Description   string `json:"description"`
			Direction     string `json:"direction"`
			StartTime     int64  `json:"starttime"` // epoch millis
			AltRoute      string `json:"altroute"`
			City          string `json:"City"`
			BeginBoundary string `json:"BeginBoundary"`
			EndBoundary   string `json:"EndBoundary"`
		} `json:"attributes"`
		Geometry struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"geometry"`
	} `json:"features"`
}

func parseClosures(r io.Reader) ([]Closure, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var resp arcgisClosureResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	loc, _ := time.LoadLocation("America/Chicago")
	out := make([]Closure, 0, len(resp.Features))
	for _, f := range resp.Features {
		a := f.Attributes
		reason := subtypeLabel[a.Subtype]
		if reason == "" {
			reason = a.Subtype
		}
		started := ""
		if a.StartTime > 0 {
			started = time.UnixMilli(a.StartTime).In(loc).Format("Mon Jan 2, 3:04 PM MST")
		}
		out = append(out, Closure{
			Street:      a.Street,
			Reason:      reason,
			Description: a.Description,
			City:        a.City,
			Begin:       a.BeginBoundary,
			End:         a.EndBoundary,
			AltRoute:    a.AltRoute,
			StartedAt:   started,
			Lat:         f.Geometry.Y,
			Lon:         f.Geometry.X,
		})
	}
	return out, nil
}

func closuresGeoJSON(closures []Closure) []byte {
	type feature struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties map[string]any `json:"properties"`
	}
	type collection struct {
		Type     string    `json:"type"`
		Features []feature `json:"features"`
	}
	c := collection{Type: "FeatureCollection"}
	for _, cl := range closures {
		if cl.Lat == 0 && cl.Lon == 0 {
			continue
		}
		f := feature{Type: "Feature"}
		f.Geometry.Type = "Point"
		f.Geometry.Coordinates = []float64{cl.Lon, cl.Lat}
		f.Properties = map[string]any{
			"street":      cl.Street,
			"reason":      cl.Reason,
			"description": cl.Description,
			"city":        cl.City,
			"begin":       cl.Begin,
			"end":         cl.End,
			"altRoute":    cl.AltRoute,
			"startedAt":   cl.StartedAt,
		}
		c.Features = append(c.Features, f)
	}
	out, _ := json.Marshal(c)
	return out
}

// --- Handler --------------------------------------------------------------

// HandleFlooding returns the /flooding handler.
func HandleFlooding(alertCache, gaugeCache, closureCache cache.CacheProvider) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		data := FloodingData{
			Menu:            menu.New(),
			CurrentPath:     r.URL.Path,
			RefreshInterval: 300,
		}

		// --- Alerts (Sedgwick County, FIPS zone KSC173) ---
		alertURL := "https://api.weather.gov/alerts/active?zone=KSC173"
		if ar, err := alertCache.GetContent(alertURL, "", "alerts-ksc173.json"); err == nil {
			alerts, geo, perr := parseNWSAlerts(ar)
			_ = ar.Close()
			if perr == nil {
				data.Alerts = alerts
				data.AlertsGeoJSON = template.JS(geo)
				data.AlertsAvailable = true
			}
		}

		// --- Gauges ---
		gaugeReadings := make([]Gauge, 0, len(gaugesForWichita))
		anyGauge := false
		for _, g := range gaugesForWichita {
			url := fmt.Sprintf("https://api.water.noaa.gov/nwps/v1/gauges/%s", strings.ToLower(g.ID))
			gr, err := gaugeCache.GetContent(url, "", "gauge-"+strings.ToLower(g.ID)+".json")
			if err != nil {
				gaugeReadings = append(gaugeReadings, Gauge{ID: g.ID, Name: g.Name, Lat: g.Lat, Lon: g.Lon, Status: "Unknown"})
				continue
			}
			gaugeReadings = append(gaugeReadings, parseNWPSGauge(gr, g))
			_ = gr.Close()
			anyGauge = true
		}
		data.Gauges = gaugeReadings
		data.GaugesAvailable = anyGauge
		data.GaugesGeoJSON = template.JS(gaugesGeoJSON(gaugeReadings))

		// --- Closures (Sedgwick County, flood/storm subtypes, in-progress only) ---
		if cr, err := closureCache.GetContent(closuresQueryURL, "", "closures-sgw.json"); err == nil {
			closures, perr := parseClosures(cr)
			_ = cr.Close()
			if perr == nil {
				data.Closures = closures
				data.ClosuresAvailable = true
				data.ClosuresGeoJSON = template.JS(closuresGeoJSON(closures))
			}
		}

		// --- Banner ---
		data.BannerLevel, data.BannerText = computeBanner(data.Alerts, data.Gauges)

		loc, _ := time.LoadLocation("America/Chicago")
		data.GeneratedAt = time.Now().In(loc).Format("Mon Jan 2 3:04:05 PM MST")

		ts, err := templates.Get("flooding")
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to get template set 'flooding': %w", err))
		}
		if err := ts.ExecuteTemplate(w, "flooding", data); err != nil {
			return middleware.InternalError(fmt.Errorf("failed to render template 'flooding': %w", err))
		}
		return nil
	}
}

func computeBanner(alerts []Alert, gauges []Gauge) (level, text string) {
	worstAlert := 0
	var worstAlertEvent string
	for _, a := range alerts {
		if a.SeverityRank > worstAlert {
			worstAlert = a.SeverityRank
			worstAlertEvent = a.Event
		}
	}
	worstGauge := 0
	var worstGaugeName string
	for _, g := range gauges {
		if g.StatusRank > worstGauge {
			worstGauge = g.StatusRank
			worstGaugeName = g.Name
		}
	}

	switch {
	case worstAlert >= sevFlash:
		return "flash", fmt.Sprintf("A %s is in effect for Sedgwick County.", worstAlertEvent)
	case worstAlert >= sevWarning:
		return "warning", fmt.Sprintf("A %s is in effect for Sedgwick County.", worstAlertEvent)
	case worstAlert >= sevWatch:
		return "watch", fmt.Sprintf("A %s is in effect for Sedgwick County.", worstAlertEvent)
	case worstAlert >= sevAdvisory:
		return "advisory", fmt.Sprintf("A %s is in effect for Sedgwick County.", worstAlertEvent)
	case worstGauge >= 3:
		return "warning", fmt.Sprintf("%s is at moderate or major flood stage.", worstGaugeName)
	case worstGauge == 2:
		return "watch", fmt.Sprintf("%s is at minor flood stage.", worstGaugeName)
	case worstGauge == 1:
		return "advisory", fmt.Sprintf("%s is near action stage.", worstGaugeName)
	}
	return "none", "No flood alerts or warnings are currently in effect for Sedgwick County."
}
