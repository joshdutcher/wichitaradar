package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
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
	XMLName     xml.Name    `xml:"xml"`
	Graphicasts Graphicasts `xml:"graphicasts"`
}

var outlookTimestampRe = regexp.MustCompile(`["']otlk_(\d{4})["']`)

// extractOutlookTimestamp extracts the issuance timestamp from SPC outlook HTML.
func extractOutlookTimestamp(html string) string {
	m := outlookTimestampRe.FindStringSubmatch(html)
	if len(m) < 2 {
		return ""
	}
	return m[1]
}

// resolveSPCOutlookURL fetches an SPC outlook HTML page via cache, extracts the
// current issuance timestamp, and returns the corresponding PNG URL. On any
// error it returns the fallback URL.
func resolveSPCOutlookURL(spcCache cache.CacheProvider, pageURL, dayPrefix, fallbackURL string) string {
	reader, err := spcCache.GetContent(pageURL, "https://www.spc.noaa.gov/products/outlook/", dayPrefix+".html")
	if err != nil {
		return fallbackURL
	}
	defer func() {
		_ = reader.Close()
	}()

	body, err := io.ReadAll(reader)
	if err != nil {
		return fallbackURL
	}

	ts := extractOutlookTimestamp(string(body))
	if ts == "" {
		return fallbackURL
	}

	return fmt.Sprintf("https://www.spc.noaa.gov/products/outlook/%s_%s.png", dayPrefix, ts)
}

// HandleOutlook handles the outlook page
func HandleOutlook(xmlCache, spcCache cache.CacheProvider) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Fetch weather stories using the cache
		url := "https://www.weather.gov/source/ict/wxstory/wxstory.xml"
		xmlReader, err := xmlCache.GetContent(url, "https://www.weather.gov/ict/", "wxstory.xml")
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to get XML: %w", err))
		}
		defer func() {
			_ = xmlReader.Close()
		}()

		// Parse stories from XML
		stories, err := parseWeatherStories(xmlReader)
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to parse XML: %w", err))
		}

		// Resolve SPC convective outlook image URLs
		spcBase := "https://www.spc.noaa.gov/products/outlook/"
		outlookURLs := make([]string, 3)
		for i := 1; i <= 3; i++ {
			dayPrefix := fmt.Sprintf("day%dotlk", i)
			pageURL := spcBase + dayPrefix + ".html"
			fallbackURL := spcBase + dayPrefix + ".gif"
			outlookURLs[i-1] = resolveSPCOutlookURL(spcCache, pageURL, dayPrefix, fallbackURL)
		}

		return renderOutlook(w, r, stories, outlookURLs)
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

func renderOutlook(w http.ResponseWriter, r *http.Request, stories []WeatherStory, outlookURLs []string) error {
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
		Menu        *menu.Menu
		CurrentPath string
		Stories     []struct {
			Image       string
			Description string
		}
		OutlookURLs     []string
		RefreshInterval int
	}{
		Menu:            menu.New(),
		CurrentPath:     r.URL.Path,
		Stories:         processedStories,
		OutlookURLs:     outlookURLs,
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
