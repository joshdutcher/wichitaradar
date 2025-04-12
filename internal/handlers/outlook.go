package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// WeatherStory represents a weather story image
type WeatherStory struct {
	URL   string
	Alt   string
	Order int
}

type Graphicast struct {
	StartTime   string `xml:"StartTime"`
	EndTime     string `xml:"EndTime"`
	Radar       string `xml:"radar"`
	SmallImage  string `xml:"SmallImage"`
	Description string `xml:"description"`
	Order       int    `xml:"order"`
}

type Graphicasts struct {
	Graphicasts []Graphicast `xml:"graphicast"`
}

type WeatherFeed struct {
	Graphicasts Graphicasts `xml:"graphicasts"`
}

// HandleOutlook handles the outlook page
func HandleOutlook(w http.ResponseWriter, r *http.Request) {
	// Fetch weather stories
	stories, err := getWeatherStories()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get weather stories: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert stories to the correct type
	processedStories := make([]struct {
		Image       string
		Description string
	}, len(stories))

	for i, story := range stories {
		processedStories[i] = struct {
			Image       string
			Description string
		}{
			Image:       story.URL,
			Description: story.Alt,
		}
	}

	// Create template data
	data := struct {
		Menu        *menu.Menu
		CurrentPath string
		Stories     []struct {
			Image       string
			Description string
		}
	}{
		Menu:        menu.New(),
		CurrentPath: r.URL.Path,
		Stories:     processedStories,
	}

	// Check if menu creation failed silently
	if data.Menu == nil {
		http.Error(w, "menu.New() returned nil", http.StatusInternalServerError)
		return
	}

	// Get the specific template set for this page
	ts, err := templates.Get("outlook")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get template set 'outlook': %v", err), http.StatusInternalServerError)
		return
	}

	// Execute the main template definition within this set
	if err := ts.ExecuteTemplate(w, "outlook", data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template 'outlook': %v", err), http.StatusInternalServerError)
		return
	}
}

func getWeatherStories() ([]WeatherStory, error) {
	// Ensure the directory exists
	xmlDir := "scraped/xml"
	if err := os.MkdirAll(xmlDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create xml directory: %v", err)
	}

	xmlPath := filepath.Join(xmlDir, "wxstory.xml")

	// Check if file exists and is less than 5 minutes old
	shouldDownload := true
	if info, err := os.Stat(xmlPath); err == nil {
		if time.Since(info.ModTime()) < 5*time.Minute {
			shouldDownload = false
		}
	}

	// Download the file if needed
	if shouldDownload {
		resp, err := http.Get("https://www.weather.gov/ict/wxstory/wxstory.xml")
		if err != nil {
			return nil, fmt.Errorf("failed to download XML: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		// Create the file
		file, err := os.Create(xmlPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()

		// Copy the response body to the file
		if _, err := io.Copy(file, resp.Body); err != nil {
			return nil, fmt.Errorf("failed to write file: %v", err)
		}
	}

	// Read and parse the local file
	file, err := os.Open(xmlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Parse XML
	var feed WeatherFeed
	decoder := xml.NewDecoder(file)
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity

	if err := decoder.Decode(&feed); err != nil {
		// Don't return error, just use default story
		return []WeatherStory{{
			URL:   "/img/nostories.png",
			Alt:   "No Weather Stories!",
			Order: 0,
		}}, nil
	}

	// Process stories
	var stories []WeatherStory
	timeNow := time.Now().Unix()

	for _, graphicast := range feed.Graphicasts.Graphicasts {
		// Parse Unix timestamps
		startTime, err := strconv.ParseInt(graphicast.StartTime, 10, 64)
		if err != nil {
			continue
		}
		endTime, err := strconv.ParseInt(graphicast.EndTime, 10, 64)
		if err != nil {
			continue
		}

		// Check if story is current and not a radar image
		if timeNow < endTime && timeNow >= startTime && graphicast.Radar != "true" {
			// Clean up image URL
			imageUrl := strings.TrimLeft(graphicast.SmallImage, "/")
			if !strings.HasPrefix(imageUrl, "http://") && !strings.HasPrefix(imageUrl, "https://") {
				imageUrl = "http://weather.gov/" + imageUrl
			}

			// Add random query param to prevent caching
			imageUrl += "?" + fmt.Sprintf("%d", rand.Intn(900000)+100000)

			// Clean up description
			description := strings.Join(strings.Fields(strings.TrimSpace(graphicast.Description)), " ")

			stories = append(stories, WeatherStory{
				URL:   imageUrl,
				Alt:   description,
				Order: graphicast.Order,
			})
		}
	}

	// If no stories, add default
	if len(stories) == 0 {
		stories = []WeatherStory{{
			URL:   "/img/nostories.png",
			Alt:   "No Weather Stories!",
			Order: 0,
		}}
	}

	// Sort stories by order
	sort.Slice(stories, func(i, j int) bool {
		return stories[i].Order < stories[j].Order
	})

	return stories, nil
}
