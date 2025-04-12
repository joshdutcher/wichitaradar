package handlers

import (
	"bytes"
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
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get working directory: %v\n", err)
		return nil, fmt.Errorf("failed to get working directory: %v", err)
	}

	// Ensure the directory exists and is writable
	xmlDir := filepath.Join(workDir, "scraped/xml")
	if err := os.MkdirAll(xmlDir, 0755); err != nil {
		fmt.Printf("Failed to create xml directory: %v\n", err)
		return nil, fmt.Errorf("failed to create xml directory: %v", err)
	}

	xmlPath := filepath.Join(xmlDir, "wxstory.xml")

	// Check if we need to download a new file
	shouldDownload := true
	if info, err := os.Stat(xmlPath); err == nil {
		// File exists, check its age
		age := time.Since(info.ModTime())
		fmt.Printf("Cached XML file is %v old\n", age)
		if age < 5*time.Minute {
			shouldDownload = false
			fmt.Printf("Using cached XML file (downloaded %v ago)\n", age)
		} else {
			fmt.Println("Cached XML file is too old, downloading fresh")
		}
	} else {
		fmt.Println("No cached XML file found, downloading fresh")
	}

	var body []byte
	if shouldDownload {
		// Download fresh XML file
		fmt.Println("Downloading fresh XML file")
		// DO NOT CHANGE THIS URL - it is the correct source for Wichita weather stories XML
		resp, err := http.Get("https://www.weather.gov/source/ict/wxstory/wxstory.xml")
		if err != nil {
			fmt.Printf("Failed to download XML: %v\n", err)
			return nil, fmt.Errorf("failed to download XML: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		// Read the response body
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}

		// Save to cache
		if err := os.WriteFile(xmlPath, body, 0644); err != nil {
			fmt.Printf("Failed to write cache file: %v\n", err)
			// Continue anyway, we have the data in memory
		} else {
			fmt.Println("Successfully cached XML file")
		}
	} else {
		// Read from cache
		var err error
		body, err = os.ReadFile(xmlPath)
		if err != nil {
			fmt.Printf("Failed to read cache file: %v\n", err)
			return nil, fmt.Errorf("failed to read cache file: %v", err)
		}
		fmt.Println("Successfully read from cache")
	}

	// Print first 500 characters of XML for debugging
	fmt.Printf("XML content (first 500 chars): %s\n", string(body[:min(500, len(body))]))

	// Parse XML
	var feed WeatherFeed
	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity

	if err := decoder.Decode(&feed); err != nil {
		fmt.Printf("Failed to parse XML: %v\n", err)
		// Don't return error, just use default story
		return []WeatherStory{{
			URL:   "/static/img/nostories.png",
			Alt:   "No Weather Stories!",
			Order: 0,
		}}, nil
	}

	fmt.Printf("Found %d graphicasts in XML\n", len(feed.Graphicasts.Graphicasts))

	// Process stories
	var stories []WeatherStory
	timeNow := time.Now().Unix()

	for _, graphicast := range feed.Graphicasts.Graphicasts {
		// Parse Unix timestamps
		startTime, err := strconv.ParseInt(graphicast.StartTime, 10, 64)
		if err != nil {
			fmt.Printf("Failed to parse start time: %v\n", err)
			continue
		}
		endTime, err := strconv.ParseInt(graphicast.EndTime, 10, 64)
		if err != nil {
			fmt.Printf("Failed to parse end time: %v\n", err)
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
			fmt.Printf("Added story: %s\n", description)
		}
	}

	// If no stories, add default
	if len(stories) == 0 {
		fmt.Println("No valid stories found, using default")
		stories = []WeatherStory{{
			URL:   "/static/img/nostories.png",
			Alt:   "No Weather Stories!",
			Order: 0,
		}}
	}

	// Sort stories by order
	sort.Slice(stories, func(i, j int) bool {
		return stories[i].Order < stories[j].Order
	})

	fmt.Printf("Returning %d stories\n", len(stories))
	return stories, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
