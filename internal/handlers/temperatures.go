package handlers

import (
	"fmt"
	"net/http"
	"time"

	"wichitaradar/internal/middleware"
	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

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
	for _, imagetype := range s.images {
		imageFound := false
		for hoursAgo := 0; hoursAgo < s.maxHoursAgo; hoursAgo++ {
			timeToCheck := s.runtime.Add(-time.Duration(hoursAgo) * time.Hour)
			imageUrl := fmt.Sprintf("%s/%s/%s/%s00z.jpg",
				s.prefix,
				imagetype,
				timeToCheck.Format("20060102"),
				timeToCheck.Format("15"))

			if checkRemoteFile(imageUrl) {
				s.imagePaths[imagetype] = imageUrl
				imageFound = true
				break
			}
		}
		if !imageFound {
			s.imagePaths[imagetype] = ""
		}
	}
	return s.imagePaths
}

func checkRemoteFile(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	return resp.StatusCode == http.StatusOK
}

func HandleTemperatures(w http.ResponseWriter, r *http.Request) error {
	swxco := NewSWXCOFiles()
	swxcoFiles := swxco.getImagePaths()

	// Create menu - if it fails, return an error that will bubble up
	m := menu.New()
	if m == nil {
		return middleware.InternalError(fmt.Errorf("menu.New() returned nil"))
	}

	// Get template set - if it fails, return an error that will bubble up
	ts, err := templates.Get("temperatures")
	if err != nil {
		return middleware.InternalError(fmt.Errorf("failed to get template set 'temperatures': %w", err))
	}

	data := struct {
		Menu            *menu.Menu
		CurrentPath     string
		SWXCOFiles      map[string]string
		RefreshInterval int
	}{
		Menu:            m,
		CurrentPath:     r.URL.Path,
		SWXCOFiles:      swxcoFiles,
		RefreshInterval: 0, // No auto-refresh for temperatures page
	}

	// Execute template - if it fails, return an error that will bubble up
	if err := ts.ExecuteTemplate(w, "temperatures", data); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to render template 'temperatures': %w", err))
	}

	return nil
}
