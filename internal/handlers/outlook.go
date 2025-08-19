package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"wichitaradar/internal/cache"
	"wichitaradar/internal/middleware"
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
	XMLName     xml.Name `xml:"xml"`
	Graphicasts Graphicasts `xml:"graphicasts"`
}

// HandleOutlook handles the outlook page
func HandleOutlook(cache cache.CacheProvider) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Fetch weather stories using the cache
		url := "https://www.weather.gov/source/ict/wxstory/wxstory.xml"
		xmlReader, err := cache.GetContent(url, "https://www.weather.gov/ict/", "wxstory.xml")
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to get XML: %w", err))
		}
		defer xmlReader.Close()

		// Parse stories from XML
		stories, err := parseWeatherStories(xmlReader)
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to parse XML: %w", err))
		}

		return renderOutlook(w, r, stories)
	}
}

func parseWeatherStories(r io.Reader) ([]WeatherStory, error) {
	var feed WeatherFeed

	// Read and parse XML
	xmlBytes, err := io.ReadAll(r)
	if err != nil {
		return getDefaultWeatherStory(), nil
	}

	if err := xml.NewDecoder(bytes.NewReader(xmlBytes)).Decode(&feed); err != nil {
		return getDefaultWeatherStory(), nil
	}

	timeNow := time.Now().Unix()
	var stories []WeatherStory
	for _, g := range feed.Graphicasts.Graphicasts {
		startTime, _ := strconv.ParseInt(g.StartTime, 10, 64)
		endTime, _ := strconv.ParseInt(g.EndTime, 10, 64)
		if startTime <= timeNow && timeNow <= endTime && g.Radar == "0" {
			stories = append(stories, WeatherStory{
				URL:   g.SmallImage,
				Alt:   g.Description,
				Order: g.Order,
			})
		}
	}

	if len(stories) == 0 {
		return getDefaultWeatherStory(), nil
	}

	// Sort stories by order
	sort.Slice(stories, func(i, j int) bool {
		return stories[i].Order < stories[j].Order
	})

	return stories, nil
}

func renderOutlook(w http.ResponseWriter, r *http.Request, stories []WeatherStory) error {
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

	data := struct {
		Menu            *menu.Menu
		CurrentPath     string
		Stories         []struct {
			Image       string
			Description string
		}
		RefreshInterval int
	}{
		Menu:            menu.New(),
		CurrentPath:     r.URL.Path,
		Stories:         processedStories,
		RefreshInterval: 1800,
	}

	if data.Menu == nil {
		return middleware.InternalError(fmt.Errorf("menu.New() returned nil"))
	}

	ts, err := templates.Get("outlook")
	if err != nil {
		return middleware.InternalError(fmt.Errorf("failed to get template set 'outlook': %w", err))
	}

	if err := ts.ExecuteTemplate(w, "outlook", data); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to render template 'outlook': %w", err))
	}

	return nil
}

// getDefaultWeatherStory returns the default story when fetching fails
func getDefaultWeatherStory() []WeatherStory {
	return []WeatherStory{{
		URL:   "/static/img/nostories.png",
		Alt:   "No Weather Stories!",
		Order: 0,
	}}
}
