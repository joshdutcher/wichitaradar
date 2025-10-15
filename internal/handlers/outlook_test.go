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

	// Create mock cache
	mockCache := &testutils.MockCacheProvider{Content: mockXML}

	// Create test request
	req := httptest.NewRequest("GET", "/outlook", nil)
	w := httptest.NewRecorder()

	// Call handler
	handler := HandleOutlook(mockCache)
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
}
