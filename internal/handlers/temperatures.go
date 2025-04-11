package handlers

import (
	"fmt"
	"net/http"
	"time"

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

func HandleTemperatures(w http.ResponseWriter, r *http.Request) {
	swxco := NewSWXCOFiles()
	swxcoFiles := swxco.getImagePaths()

	data := struct {
		Menu        *menu.Menu
		CurrentPath string
		SWXCOFiles  map[string]string
	}{
		Menu:        menu.New(),
		CurrentPath: r.URL.Path,
		SWXCOFiles:  swxcoFiles,
	}

	// Check if menu creation failed silently
	if data.Menu == nil {
		http.Error(w, "menu.New() returned nil", http.StatusInternalServerError)
		return
	}

	// Get the specific template set for this page
	ts, err := templates.Get("temperatures")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get template set 'temperatures': %v", err), http.StatusInternalServerError)
		return
	}

	// Add a check to ensure ts is not nil, although Get should handle this
	if ts == nil {
		http.Error(w, "Got nil template set from templates.Get", http.StatusInternalServerError)
		return
	}
	// Execute the main template definition within this set
	if err := ts.ExecuteTemplate(w, "temperatures", data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template 'temperatures': %v", err), http.StatusInternalServerError)
		return
	}
}
