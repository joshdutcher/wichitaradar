package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"wichitaradar/internal/testutils"
)

func TestGetDefaultWeatherStory(t *testing.T) {
	got := getDefaultWeatherStory()
	want := []WeatherStory{{
		URL:   "/static/img/nostories.png",
		Alt:   "No Weather Stories!",
		Order: 0,
	}}

	if len(got) != len(want) {
		t.Errorf("got %d stories, want %d", len(got), len(want))
		return
	}

	if got[0].URL != want[0].URL {
		t.Errorf("URL = %v, want %v", got[0].URL, want[0].URL)
	}
	if got[0].Alt != want[0].Alt {
		t.Errorf("Alt = %v, want %v", got[0].Alt, want[0].Alt)
	}
	if got[0].Order != want[0].Order {
		t.Errorf("Order = %v, want %v", got[0].Order, want[0].Order)
	}
}

func TestGetWeatherStoriesFromReader(t *testing.T) {
	// Calculate timestamps relative to now
	now := time.Now().Unix()
	startTime := now - 3600 // 1 hour ago
	endTime := now + 3600   // 1 hour from now

	tests := []struct {
		name    string
		xmlData string
		want    []WeatherStory
	}{
		{
			name: "valid XML with one active story",
			xmlData: fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<xml><graphicasts><graphicast><StartTime>%d</StartTime><EndTime>%d</EndTime><description>Rain Showers Continue</description><SmallImage>/images/ict/wxstory/Tab1FileL.png</SmallImage><order>1</order><radar>0</radar></graphicast></graphicasts></xml>`, startTime, endTime),
			want: []WeatherStory{
				{
					URL:   "/images/ict/wxstory/Tab1FileL.png",
					Alt:   "Rain Showers Continue",
					Order: 1,
				},
			},
		},
		{
			name: "expired story",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<xml><graphicasts><graphicast><StartTime>1643757560</StartTime><EndTime>1643798600</EndTime><description>Old Story</description><SmallImage>/images/ict/wxstory/Tab1FileL.png</SmallImage><order>1</order><radar>false</radar></graphicast></graphicasts></xml>`,
			want: []WeatherStory{{
				URL:   "/static/img/nostories.png",
				Alt:   "No Weather Stories!",
				Order: 0,
			}},
		},
		{
			name: "invalid XML",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<xml><graphicasts><graphicast><invalidTag></graphicast></graphicasts></xml>`,
			want: []WeatherStory{{
				URL:   "/static/img/nostories.png",
				Alt:   "No Weather Stories!",
				Order: 0,
			}},
		},
		{
			name: "radar image should be skipped",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<xml><graphicasts><graphicast><StartTime>1743757560</StartTime><EndTime>1743798600</EndTime><description>Radar Image</description><SmallImage>/images/ict/wxstory/radar.gif</SmallImage><order>1</order><radar>true</radar></graphicast></graphicasts></xml>`,
			want: []WeatherStory{{
				URL:   "/static/img/nostories.png",
				Alt:   "No Weather Stories!",
				Order: 0,
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseWeatherStories(strings.NewReader(tt.xmlData))
			if err != nil {
				t.Fatalf("parseWeatherStories() error = %v", err)
			}

			if len(got) != len(tt.want) {
				t.Errorf("parseWeatherStories() got %d stories, want %d", len(got), len(tt.want))
				return
			}

			for i := range got {
				// Remove random query param from URL before comparing
				gotURL := strings.Split(got[i].URL, "?")[0]
				if gotURL != tt.want[i].URL {
					t.Errorf("Story %d URL = %v, want %v", i, gotURL, tt.want[i].URL)
				}
				if got[i].Alt != tt.want[i].Alt {
					t.Errorf("Story %d Alt = %v, want %v", i, got[i].Alt, tt.want[i].Alt)
				}
				if got[i].Order != tt.want[i].Order {
					t.Errorf("Story %d Order = %v, want %v", i, got[i].Order, tt.want[i].Order)
				}
			}
		})
	}
}

// mockCacheProvider implements cache.CacheProvider for testing
type mockCacheProvider struct {
	content string
}

func (m *mockCacheProvider) GetContent(url string, referer string, filename ...string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(m.content)), nil
}

func TestExtractOutlookTimestamp(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "valid timestamp with double quotes",
			html: `<script>var defined = "otlk_1630";</script>`,
			want: "1630",
		},
		{
			name: "valid timestamp with single quotes",
			html: `<script>var defined = 'otlk_0100';</script>`,
			want: "0100",
		},
		{
			name: "no match",
			html: `<html><body>no outlook data here</body></html>`,
			want: "",
		},
		{
			name: "empty string",
			html: "",
			want: "",
		},
		{
			name: "multiple matches returns first",
			html: `"otlk_1300" and "otlk_1630"`,
			want: "1300",
		},
		{
			name: "realistic SPC HTML snippet",
			html: `<script language="JavaScript">
var defined = "otlk_2000";
var defined2 = "otlk_2000_torn";
</script>`,
			want: "2000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractOutlookTimestamp(tt.html)
			if got != tt.want {
				t.Errorf("extractOutlookTimestamp() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestResolveSPCOutlookURL(t *testing.T) {
	fallbackURL := "https://www.spc.noaa.gov/products/outlook/day1otlk.gif"

	tests := []struct {
		name      string
		html      string
		cacheErr  bool
		dayPrefix string
		want      string
	}{
		{
			name:      "success extracts PNG URL",
			html:      `<script>var defined = "otlk_1630";</script>`,
			dayPrefix: "day1otlk",
			want:      "https://www.spc.noaa.gov/products/outlook/day1otlk_1630.png",
		},
		{
			name:      "cache error returns fallback",
			cacheErr:  true,
			dayPrefix: "day1otlk",
			want:      fallbackURL,
		},
		{
			name:      "no timestamp match returns fallback",
			html:      `<html><body>nothing here</body></html>`,
			dayPrefix: "day1otlk",
			want:      fallbackURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cache interface {
				GetContent(url string, referer string, filename ...string) (io.ReadCloser, error)
			}
			if tt.cacheErr {
				cache = &testutils.MockErrorCacheProvider{}
			} else {
				cache = &testutils.MockCacheProvider{Content: tt.html}
			}

			pageURL := "https://www.spc.noaa.gov/products/outlook/" + tt.dayPrefix + ".html"
			got := resolveSPCOutlookURL(cache, pageURL, tt.dayPrefix, fallbackURL)
			if got != tt.want {
				t.Errorf("resolveSPCOutlookURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestHandleOutlook(t *testing.T) {
	// Initialize templates
	testutils.InitTemplates(t)

	// Calculate timestamps relative to now
	now := time.Now().Unix()
	startTime := now - 3600 // 1 hour ago
	endTime := now + 3600   // 1 hour from now

	// Create mock cache with test XML
	mockXML := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<xml>
	<graphicasts>
		<graphicast>
			<StartTime>%d</StartTime>
			<EndTime>%d</EndTime>
			<radar>0</radar>
			<SmallImage>test.jpg</SmallImage>
			<description>Mock Story</description>
			<order>1</order>
		</graphicast>
	</graphicasts>
</xml>`, startTime, endTime)

	// Create mock caches
	mockXMLCache := &testutils.MockCacheProvider{Content: mockXML}
	mockSPCCache := &testutils.MockCacheProvider{Content: `<script>var defined = "otlk_1630";</script>`}

	// Create test request
	req := httptest.NewRequest("GET", "/outlook", nil)
	w := httptest.NewRecorder()

	// Call handler
	handler := HandleOutlook(mockXMLCache, mockSPCCache)
	err := handler(w, req)
	if err != nil {
		t.Fatalf("HandleOutlook failed: %v", err)
	}

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}

	// Check response body
	body := w.Body.String()
	if !strings.Contains(body, "Mock Story") {
		t.Error("response body does not contain 'Mock Story'")
	}

	// Verify PNG URLs are rendered (not old GIF URLs)
	if !strings.Contains(body, "day1otlk_1630.png") {
		t.Error("response body does not contain day1otlk_1630.png")
	}
	if !strings.Contains(body, "day2otlk_1630.png") {
		t.Error("response body does not contain day2otlk_1630.png")
	}
	if !strings.Contains(body, "day3otlk_1630.png") {
		t.Error("response body does not contain day3otlk_1630.png")
	}
}

func TestHandleOutlookWithSPCFallback(t *testing.T) {
	// Initialize templates
	testutils.InitTemplates(t)

	now := time.Now().Unix()
	startTime := now - 3600
	endTime := now + 3600

	mockXML := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<xml>
	<graphicasts>
		<graphicast>
			<StartTime>%d</StartTime>
			<EndTime>%d</EndTime>
			<radar>0</radar>
			<SmallImage>test.jpg</SmallImage>
			<description>Mock Story</description>
			<order>1</order>
		</graphicast>
	</graphicasts>
</xml>`, startTime, endTime)

	mockXMLCache := &testutils.MockCacheProvider{Content: mockXML}
	mockSPCCache := &testutils.MockErrorCacheProvider{}

	req := httptest.NewRequest("GET", "/outlook", nil)
	w := httptest.NewRecorder()

	handler := HandleOutlook(mockXMLCache, mockSPCCache)
	err := handler(w, req)
	if err != nil {
		t.Fatalf("HandleOutlook with SPC fallback failed: %v", err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}

	// Verify fallback GIF URLs are rendered
	body := w.Body.String()
	if !strings.Contains(body, "day1otlk.gif") {
		t.Error("response body does not contain fallback day1otlk.gif")
	}
	if !strings.Contains(body, "day2otlk.gif") {
		t.Error("response body does not contain fallback day2otlk.gif")
	}
	if !strings.Contains(body, "day3otlk.gif") {
		t.Error("response body does not contain fallback day3otlk.gif")
	}
}
