package handlers

import (
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"wichitaradar/internal/cache"
	"wichitaradar/internal/testutils"
)

func TestParseWeatherFeed(t *testing.T) {
	// Use a fixed time for testing
	now := time.Unix(1743757570, 0)

	tests := []struct {
		name    string
		xmlData string
		want    []WeatherStory
	}{
		{
			name: "valid XML with one active story",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<feed><graphicasts><graphicast><StartTime>1743757560</StartTime><EndTime>1743798600</EndTime><description>Rain Showers Continue</description><SmallImage>/images/ict/wxstory/Tab1FileL.png</SmallImage><order>1</order><radar>false</radar></graphicast></graphicasts></feed>`,
			want: []WeatherStory{
				{
					URL:   "https://weather.gov/images/ict/wxstory/Tab1FileL.png",
					Alt:   "Rain Showers Continue",
					Order: 1,
				},
			},
		},
		{
			name: "expired story",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<feed><graphicasts><graphicast><StartTime>1643757560</StartTime><EndTime>1643798600</EndTime><description>Old Story</description><SmallImage>/images/ict/wxstory/Tab1FileL.png</SmallImage><order>1</order><radar>false</radar></graphicast></graphicasts></feed>`,
			want: []WeatherStory{{
				URL:   "/static/img/nostories.png",
				Alt:   "No Weather Stories!",
				Order: 0,
			}},
		},
		{
			name: "invalid XML",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<feed><graphicasts><graphicast><invalidTag></graphicast></graphicasts></feed>`,
			want: []WeatherStory{{
				URL:   "/static/img/nostories.png",
				Alt:   "No Weather Stories!",
				Order: 0,
			}},
		},
		{
			name: "radar image should be skipped",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<feed><graphicasts><graphicast><StartTime>1743757560</StartTime><EndTime>1743798600</EndTime><description>Radar Image</description><SmallImage>/images/ict/wxstory/radar.gif</SmallImage><order>1</order><radar>true</radar></graphicast></graphicasts></feed>`,
			want: []WeatherStory{{
				URL:   "/static/img/nostories.png",
				Alt:   "No Weather Stories!",
				Order: 0,
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use strings.NewReader instead of writing to a file
			got, err := getWeatherStoriesFromReader(strings.NewReader(tt.xmlData), now)
			if err != nil {
				t.Fatalf("getWeatherStoriesFromReader() error = %v", err)
			}

			// Compare results
			if len(got) != len(tt.want) {
				t.Errorf("getWeatherStoriesFromReader() got %d stories, want %d", len(got), len(tt.want))
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

func TestHandleOutlook(t *testing.T) {
	// --- Test Setup Start ---
	// Init templates (required by the handler)
	testutils.InitTemplates(t)

	// Create a temporary directory for the test cache
	tempDir, err := os.MkdirTemp("", "outlook_test_cache")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a mock wxstory.xml file with timestamps relative to the current time
	currentTime := time.Now()
	startTime := fmt.Sprintf("%d", currentTime.Add(-1*time.Hour).Unix())
	endTime := fmt.Sprintf("%d", currentTime.Add(1*time.Hour).Unix())
	testXML := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<feed><graphicasts><graphicast>
	<StartTime>%s</StartTime><EndTime>%s</EndTime>
	<description>Mock Outlook Story</description>
	<SmallImage>/images/ict/wxstory/mock.png</SmallImage>
	<order>1</order><radar>false</radar>
</graphicast></graphicasts></feed>`, startTime, endTime)
	mockXMLPath := filepath.Join(tempDir, "wxstory.xml")
	if err := os.WriteFile(mockXMLPath, []byte(testXML), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a test cache instance pointing to the temp dir
	testCache := cache.New(tempDir, 5*time.Minute) // Use a short cache time for testing

	// Create the handler instance using the factory and test cache
	handler := HandleOutlook(testCache)
	// --- Test Setup End ---

	// Create request and recorder
	req := httptest.NewRequest("GET", "/outlook", nil)
	rr := httptest.NewRecorder()

	// Execute the handler
	handler(rr, req)

	// Check status code
	if status := rr.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v", status, 200)
	}

	// Check if the body contains expected content (e.g., the mock story description)
	expectedContent := "Mock Outlook Story"
	if !strings.Contains(rr.Body.String(), expectedContent) {
		t.Errorf("handler body does not contain expected string %q", expectedContent)
	}
}