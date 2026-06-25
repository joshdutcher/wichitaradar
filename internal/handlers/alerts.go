package handlers

import (
	"encoding/json"
	"fmt"
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

// AlertsData is the template payload for the /alerts page.
type AlertsData struct {
	Menu            *menu.Menu
	CurrentPath     string
	RefreshInterval int

	Available   bool
	Alerts      []Alert
	GeneratedAt string
}

// alertCategoryRank classifies an alert event for sort + banner color. Higher
// is more urgent. This complements flooding.go's flood-only severityRank so the
// /alerts page can rank ALL hazard types consistently.
func alertCategoryRank(event string) int {
	e := strings.ToLower(event)
	switch {
	case strings.Contains(e, "tornado warning"),
		strings.Contains(e, "flash flood warning"),
		strings.Contains(e, "extreme"):
		return 4
	case strings.Contains(e, "warning"):
		return 3
	case strings.Contains(e, "watch"):
		return 2
	case strings.Contains(e, "advisory"),
		strings.Contains(e, "statement"),
		strings.Contains(e, "alert"),
		strings.Contains(e, "outlook"):
		return 1
	}
	return 0
}

func parseAllNWSAlerts(r io.Reader) ([]Alert, error) {
	var resp nwsAlertsResponse
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("America/Chicago")
	alerts := make([]Alert, 0, len(resp.Features))

	for _, raw := range resp.Features {
		var f nwsAlertFeature
		if err := json.Unmarshal(raw, &f); err != nil {
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
			SeverityRank: alertCategoryRank(f.Properties.Event),
		})
	}

	sort.SliceStable(alerts, func(i, j int) bool {
		return alerts[i].SeverityRank > alerts[j].SeverityRank
	})
	return alerts, nil
}

// HandleAlerts returns the /alerts handler. It shares the alertsCache used by
// /flooding so both pages hit api.weather.gov at most once per cache TTL.
func HandleAlerts(alertsCache cache.CacheProvider) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		data := AlertsData{
			Menu:            menu.New(),
			CurrentPath:     r.URL.Path,
			RefreshInterval: 300,
		}

		const alertURL = "https://api.weather.gov/alerts/active?zone=KSC173"
		if ar, err := alertsCache.GetContent(alertURL, "", "alerts-ksc173.json"); err == nil {
			alerts, perr := parseAllNWSAlerts(ar)
			_ = ar.Close()
			if perr == nil {
				data.Alerts = alerts
				data.Available = true
			}
		}

		loc, _ := time.LoadLocation("America/Chicago")
		data.GeneratedAt = time.Now().In(loc).Format("Mon Jan 2 3:04:05 PM MST")

		ts, err := templates.Get("alerts")
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to get template set 'alerts': %w", err))
		}
		if err := ts.ExecuteTemplate(w, "alerts", data); err != nil {
			return middleware.InternalError(fmt.Errorf("failed to render template 'alerts': %w", err))
		}
		return nil
	}
}
