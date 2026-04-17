package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"wichitaradar/internal/middleware"
	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HTTP client with timeout for external image validation
var imageCheckClient = &http.Client{
	Timeout: 5 * time.Second,
}

type SWXCOFiles struct {
	runtime     time.Time
	prefix      string
	images      []string
	imagePaths  map[string]string
	maxHoursAgo int
}

func NewSWXCOFiles() *SWXCOFiles {
	return &SWXCOFiles{
		runtime:     time.Now().UTC(),
		prefix:      "https://s.w-x.co/staticmaps/wu/fee4c/temp_cur",
		images:      []string{"usa", "ddc"},
		imagePaths:  make(map[string]string),
		maxHoursAgo: 12,
	}
}

func (s *SWXCOFiles) getImagePaths() map[string]string {
	type hit struct {
		imagetype string
		hoursAgo  int
		url       string
	}
	hits := make(chan hit, len(s.images)*s.maxHoursAgo)
	var wg sync.WaitGroup

	for _, imagetype := range s.images {
		for hoursAgo := 0; hoursAgo < s.maxHoursAgo; hoursAgo++ {
			wg.Add(1)
			go func(imagetype string, hoursAgo int) {
				defer wg.Done()
				timeToCheck := s.runtime.Add(-time.Duration(hoursAgo) * time.Hour)
				imageUrl := fmt.Sprintf("%s/%s/%s/%s00z.jpg",
					s.prefix,
					imagetype,
					timeToCheck.Format("20060102"),
					timeToCheck.Format("15"))
				if checkRemoteFile(imageUrl) {
					hits <- hit{imagetype, hoursAgo, imageUrl}
				}
			}(imagetype, hoursAgo)
		}
	}
	wg.Wait()
	close(hits)

	best := make(map[string]int)
	for h := range hits {
		if existing, ok := best[h.imagetype]; !ok || h.hoursAgo < existing {
			best[h.imagetype] = h.hoursAgo
			s.imagePaths[h.imagetype] = h.url
		}
	}
	for _, imagetype := range s.images {
		if _, ok := s.imagePaths[imagetype]; !ok {
			s.imagePaths[imagetype] = ""
		}
	}
	return s.imagePaths
}

func checkRemoteFile(url string) bool {
	resp, err := imageCheckClient.Head(url)
	if err != nil {
		return false
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return resp.StatusCode == http.StatusOK
}

func HandleTemperatures(w http.ResponseWriter, r *http.Request) error {
	m := menu.New()
	if m == nil {
		return middleware.InternalError(fmt.Errorf("menu.New() returned nil"))
	}

	ts, err := templates.Get("temperatures")
	if err != nil {
		return middleware.InternalError(fmt.Errorf("failed to get template set 'temperatures': %w", err))
	}

	data := struct {
		Menu            *menu.Menu
		CurrentPath     string
		RefreshInterval int
	}{
		Menu:            m,
		CurrentPath:     r.URL.Path,
		RefreshInterval: 0,
	}

	if err := ts.ExecuteTemplate(w, "temperatures", data); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to render template 'temperatures': %w", err))
	}

	return nil
}

// HandleWUTemperatureImages probes the Weather Underground static-map host
// in parallel and returns the most recent available image URL per region as JSON.
// Missing/stale regions come back with an empty string so the client can render
// a failure state.
func HandleWUTemperatureImages(w http.ResponseWriter, r *http.Request) error {
	paths := NewSWXCOFiles().getImagePaths()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	if err := json.NewEncoder(w).Encode(paths); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to encode WU image paths: %w", err))
	}
	return nil
}
