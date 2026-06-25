package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"wichitaradar/internal/testutils"
)

// stubCache returns a different canned response per filename prefix so we can
// distinguish the alert call from individual gauge calls.
type stubCache struct {
	responses map[string]string // keyed by filename prefix
	fallback  string
}

func (s *stubCache) GetContent(url, referer string, filename ...string) (io.ReadCloser, error) {
	name := ""
	if len(filename) > 0 {
		name = filename[0]
	}
	for prefix, body := range s.responses {
		if strings.HasPrefix(name, prefix) {
			return io.NopCloser(strings.NewReader(body)), nil
		}
	}
	if s.fallback != "" {
		return io.NopCloser(strings.NewReader(s.fallback)), nil
	}
	return nil, fmt.Errorf("no stub for %s", name)
}

func TestSeverityRank(t *testing.T) {
	cases := []struct {
		event string
		want  int
	}{
		{"Flash Flood Warning", sevFlash},
		{"Flood Warning", sevWarning},
		{"Flood Watch", sevWatch},
		{"Flash Flood Watch", sevWatch},
		{"Flood Advisory", sevAdvisory},
		{"Tornado Warning", sevNone},
	}
	for _, c := range cases {
		if got := severityRank(c.event); got != c.want {
			t.Errorf("severityRank(%q) = %d, want %d", c.event, got, c.want)
		}
	}
}

func TestParseNWSAlerts(t *testing.T) {
	body := `{
	  "features": [
	    {
	      "geometry": {"type":"Polygon","coordinates":[[[-97.5,37.6],[-97.2,37.6],[-97.2,37.8],[-97.5,37.8],[-97.5,37.6]]]},
	      "properties": {
	        "event":"Flash Flood Warning",
	        "headline":"Flash flooding likely",
	        "severity":"Severe",
	        "description":"Heavy rain producing flash flooding.",
	        "areaDesc":"Sedgwick, KS",
	        "sent":"2026-06-25T18:00:00+00:00",
	        "expires":"2026-06-25T21:00:00+00:00"
	      }
	    },
	    {
	      "geometry": null,
	      "properties": {
	        "event":"Wind Advisory",
	        "headline":"Windy",
	        "severity":"Minor"
	      }
	    }
	  ]
	}`
	alerts, geo, err := parseNWSAlerts(strings.NewReader(body))
	if err != nil {
		t.Fatalf("parseNWSAlerts error: %v", err)
	}
	if len(alerts) != 1 {
		t.Fatalf("got %d alerts, want 1", len(alerts))
	}
	if alerts[0].Event != "Flash Flood Warning" {
		t.Errorf("event = %q", alerts[0].Event)
	}
	if alerts[0].SeverityRank != sevFlash {
		t.Errorf("rank = %d", alerts[0].SeverityRank)
	}
	if !strings.Contains(string(geo), "FeatureCollection") {
		t.Errorf("geojson missing FeatureCollection: %s", geo)
	}
}

func TestParseNWPSGauge(t *testing.T) {
	body := `{
	  "lid":"ICTK1",
	  "name":"Arkansas at Wichita",
	  "status":{
	    "observed":{"primary":14.2,"primaryUnit":"ft","validTime":"2026-06-25T18:00:00Z","floodCategory":"minor"}
	  },
	  "flood":{"stageUnits":"ft","categories":{"action":{"stage":11.0},"minor":{"stage":13.0},"moderate":{"stage":16.0},"major":{"stage":20.0}}}
	}`
	g := parseNWPSGauge(strings.NewReader(body), floodGauge{ID: "ICTK1", Name: "Arkansas at Wichita", Lat: 37.69, Lon: -97.34})
	if g.Status != "Minor" {
		t.Errorf("status = %q, want Minor", g.Status)
	}
	if !strings.HasPrefix(g.Stage, "14.20") {
		t.Errorf("stage = %q", g.Stage)
	}
	if g.FloodStage != "13.0 ft" {
		t.Errorf("flood stage = %q", g.FloodStage)
	}
	if g.StatusRank != 2 {
		t.Errorf("rank = %d", g.StatusRank)
	}
}

func TestParseNWPSGaugeCategorizesFromStage(t *testing.T) {
	body := `{
	  "lid":"X","name":"X",
	  "status":{"observed":{"primary":17.0,"primaryUnit":"ft","validTime":"","floodCategory":""}},
	  "flood":{"stageUnits":"ft","categories":{"action":{"stage":11.0},"minor":{"stage":13.0},"moderate":{"stage":16.0},"major":{"stage":20.0}}}
	}`
	g := parseNWPSGauge(strings.NewReader(body), floodGauge{ID: "X", Name: "X"})
	if g.Status != "Moderate" {
		t.Errorf("status = %q, want Moderate", g.Status)
	}
}

func TestParseClosures(t *testing.T) {
	body := `{"features":[
	  {"attributes":{"street":"Lincoln St","subtype":"ROAD_CLOSED_FLOOD","description":"High water","direction":"E/W","starttime":1751823000000,"altroute":"Use Maple","City":"Wichita","BeginBoundary":"Hydraulic","EndBoundary":"I-135"},"geometry":{"x":-97.3,"y":37.65}},
	  {"attributes":{"street":"K-42","subtype":"ROAD_CLOSED_STORM_DAMAGE","description":"Downed tree","City":"Haysville"},"geometry":{"x":-97.36,"y":37.55}}
	]}`
	got, err := parseClosures(strings.NewReader(body))
	if err != nil {
		t.Fatalf("parseClosures error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d closures, want 2", len(got))
	}
	if got[0].Reason != "Flooding" {
		t.Errorf("first reason = %q, want Flooding", got[0].Reason)
	}
	if got[1].Reason != "Storm Damage" {
		t.Errorf("second reason = %q, want Storm Damage", got[1].Reason)
	}
	if got[0].City != "Wichita" || got[0].Begin != "Hydraulic" {
		t.Errorf("attributes not parsed: %+v", got[0])
	}
}

func TestComputeBanner(t *testing.T) {
	cases := []struct {
		name      string
		alerts    []Alert
		gauges    []Gauge
		wantLevel string
	}{
		{"none", nil, nil, "none"},
		{"flash beats gauge", []Alert{{Event: "Flash Flood Warning", SeverityRank: sevFlash}}, []Gauge{{StatusRank: 4}}, "flash"},
		{"gauge minor", nil, []Gauge{{Name: "Arkansas", StatusRank: 2}}, "watch"},
		{"gauge action", nil, []Gauge{{Name: "Arkansas", StatusRank: 1}}, "advisory"},
		{"warning", []Alert{{Event: "Flood Warning", SeverityRank: sevWarning}}, nil, "warning"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lvl, _ := computeBanner(c.alerts, c.gauges)
			if lvl != c.wantLevel {
				t.Errorf("level = %q, want %q", lvl, c.wantLevel)
			}
		})
	}
}

func TestHandleFloodingRenders(t *testing.T) {
	testutils.InitTemplates(t)

	alertBody := `{"features":[]}`
	gaugeBody := `{
	  "lid":"ICTK1","name":"Arkansas at Wichita",
	  "status":{"observed":{"primary":4.21,"primaryUnit":"ft","validTime":"2026-06-25T18:00:00Z","floodCategory":"no_flooding"}},
	  "flood":{"stageUnits":"ft","categories":{"action":{"stage":11.0},"minor":{"stage":13.0},"moderate":{"stage":16.0},"major":{"stage":20.0}}}
	}`

	closureBody := `{"features":[{"attributes":{"street":"K-15 / Hydraulic","subtype":"ROAD_CLOSED_FLOOD","description":"Standing water across roadway","direction":"N/S","starttime":1751823000000,"altroute":"Use Broadway","City":"Wichita","BeginBoundary":"31st St S","EndBoundary":"47th St S"},"geometry":{"x":-97.33,"y":37.62}}]}`

	alertCache := &stubCache{responses: map[string]string{"alerts-": alertBody}}
	gaugeCache := &stubCache{responses: map[string]string{"gauge-": gaugeBody}}
	closureCache := &stubCache{responses: map[string]string{"closures-": closureBody}}

	handler := HandleFlooding(alertCache, gaugeCache, closureCache)
	req := httptest.NewRequest("GET", "/flooding", nil)
	w := httptest.NewRecorder()

	if err := handler(w, req); err != nil {
		t.Fatalf("handler error: %v", err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("status = %d", w.Code)
	}
	body := w.Body.String()
	for _, want := range []string{
		"No flood alerts or warnings are currently in effect",
		"Arkansas River at Derby",
		"4.21 ft",
		"id=\"flood-map\"",
		"flood-disclaimer",
		"K-15 / Hydraulic",
		"Flooding",
		"Use Broadway",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("body missing %q", want)
		}
	}
}

func TestHandleFloodingDegradesGracefully(t *testing.T) {
	testutils.InitTemplates(t)

	errCache := &testutils.MockErrorCacheProvider{}
	handler := HandleFlooding(errCache, errCache, errCache)
	req := httptest.NewRequest("GET", "/flooding", nil)
	w := httptest.NewRecorder()

	if err := handler(w, req); err != nil {
		t.Fatalf("handler returned error on upstream failure: %v", err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "temporarily unavailable") {
		t.Error("expected degraded message in body")
	}
}
