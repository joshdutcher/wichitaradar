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

	"wichitaradar/internal/cache"
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

// NewOutlookHandler creates an HTTP handler func for the outlook page,
// using the provided cache for weather stories.
func NewOutlookHandler(xmlCache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// --- Fetch weather stories using the cache ---
		var stories []WeatherStory
		var xmlReader io.Reader

		filename := "wxstory.xml"
		cacheFile := filepath.Join(xmlCache.GetCacheDir(), filename)
		url := "https://www.weather.gov/source/ict/wxstory/wxstory.xml"

		if xmlCache.Expired(filename) {
			fmt.Println("Outlook Handler: Downloading fresh XML file")
			if err := xmlCache.DownloadFile(url, filename, "https://www.weather.gov/ict/"); err != nil {
				fmt.Printf("Outlook Handler: Failed to download XML: %v\n", err)
				// Fallback to default story on download error
				stories = getDefaultWeatherStory()
			}
		}

		// If stories haven't been set by fallback, read from cache
		if stories == nil {
			file, err := os.Open(cacheFile)
			if err != nil {
				fmt.Printf("Outlook Handler: Failed to open cached XML: %v\n", err)
				stories = getDefaultWeatherStory()
			} else {
				xmlReader = file
				defer file.Close()
			}
		}

		// If we have an XML reader (from cache), parse it
		if xmlReader != nil {
			parsedStories, err := getWeatherStoriesFromReader(xmlReader, time.Now())
			if err != nil {
				// This shouldn't happen based on getWeatherStoriesFromReader logic, but handle defensively
				fmt.Printf("Outlook Handler: Error parsing XML from reader: %v\n", err)
				stories = getDefaultWeatherStory()
			} else {
				stories = parsedStories
			}
		}
		// --- End Fetch weather stories ---

		// Convert stories to the correct type for the template
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
}

// getWeatherStoriesFromReader parses weather stories from an XML reader
func getWeatherStoriesFromReader(xmlReader io.Reader, now time.Time) ([]WeatherStory, error) {
	var feed WeatherFeed
	decoder := xml.NewDecoder(xmlReader)
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity

	if err := decoder.Decode(&feed); err != nil {
		fmt.Printf("Failed to parse XML: %v\n", err)
		// Don't return error, just use default story
		return getDefaultWeatherStory(), nil
	}

	fmt.Printf("Found %d graphicasts in XML\n", len(feed.Graphicasts.Graphicasts))

	// Process stories
	var stories []WeatherStory
	timeNow := now.Unix()

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
				imageUrl = "https://weather.gov/" + imageUrl
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
		stories = getDefaultWeatherStory()
	}

	// Sort stories by order
	sort.Slice(stories, func(i, j int) bool {
		return stories[i].Order < stories[j].Order
	})

	fmt.Printf("Returning %d stories\n", len(stories))
	return stories, nil
}

// getDefaultWeatherStory returns the default story when fetching fails
func getDefaultWeatherStory() []WeatherStory {
	return []WeatherStory{{
		URL:   "/static/img/nostories.png",
		Alt:   "No Weather Stories!",
		Order: 0,
	}}
}
