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

const (
	// wuResultTTL is how long we serve a previously-computed probe result
	// before re-probing upstream. Keeps the 24-HEAD fan-out off the hot path.
	wuResultTTL = 10 * time.Minute
	// wuLastGoodMaxAge bounds how stale a last-known-good URL can be before
	// we stop using it as a network-error fallback.
	wuLastGoodMaxAge = 3 * time.Hour
)

type wuCacheEntry struct {
	url string
	at  time.Time
}

var (
	wuMu           sync.Mutex
	wuResult       map[string]string
	wuResultAt     time.Time
	wuLastGoodURLs = map[string]wuCacheEntry{}
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

type probeResult int

const (
	probeMissing probeResult = iota // clean HTTP miss (e.g. 404)
	probeFound                      // HTTP 200
	probeNetErr                     // transport error (DNS, timeout, TCP reset)
)

// getImagePaths returns the freshest available image URL per region. When a
// region has no HTTP hit but at least one probe failed with a transport error
// (DNS, timeout, etc.), it falls back to the last known-good URL, or failing
// that, the current-hour URL — so a server-side resolver hiccup doesn't blank
// a slot whose image the user's browser could have loaded fine.
func (s *SWXCOFiles) getImagePaths() map[string]string {
	type hit struct {
		imagetype string
		hoursAgo  int
		url       string
		result    probeResult
	}
	hits := make(chan hit, len(s.images)*s.maxHoursAgo)
	var wg sync.WaitGroup

	for _, imagetype := range s.images {
		for hoursAgo := 0; hoursAgo < s.maxHoursAgo; hoursAgo++ {
			wg.Add(1)
			go func(imagetype string, hoursAgo int) {
				defer wg.Done()
				timeToCheck := s.runtime.Add(-time.Duration(hoursAgo) * time.Hour)
				imageUrl := s.urlFor(imagetype, timeToCheck)
				hits <- hit{imagetype, hoursAgo, imageUrl, checkRemoteFile(imageUrl)}
			}(imagetype, hoursAgo)
		}
	}
	wg.Wait()
	close(hits)

	best := make(map[string]int)
	netErrByType := make(map[string]bool)
	for h := range hits {
		switch h.result {
		case probeFound:
			if existing, ok := best[h.imagetype]; !ok || h.hoursAgo < existing {
				best[h.imagetype] = h.hoursAgo
				s.imagePaths[h.imagetype] = h.url
			}
		case probeNetErr:
			netErrByType[h.imagetype] = true
		}
	}

	now := time.Now().UTC()
	wuMu.Lock()
	defer wuMu.Unlock()
	for _, imagetype := range s.images {
		if url, ok := s.imagePaths[imagetype]; ok && url != "" {
			wuLastGoodURLs[imagetype] = wuCacheEntry{url: url, at: now}
			continue
		}
		if netErrByType[imagetype] {
			if lg, ok := wuLastGoodURLs[imagetype]; ok && now.Sub(lg.at) < wuLastGoodMaxAge {
				s.imagePaths[imagetype] = lg.url
				continue
			}
			s.imagePaths[imagetype] = s.urlFor(imagetype, s.runtime)
			continue
		}
		s.imagePaths[imagetype] = ""
	}
	return s.imagePaths
}

func (s *SWXCOFiles) urlFor(imagetype string, t time.Time) string {
	return fmt.Sprintf("%s/%s/%s/%s00z.jpg",
		s.prefix,
		imagetype,
		t.Format("20060102"),
		t.Format("15"))
}

func checkRemoteFile(url string) probeResult {
	resp, err := imageCheckClient.Head(url)
	if err != nil {
		// HEAD doesn't return err for HTTP statuses — any error here is a
		// transport-level failure (DNS, timeout, TCP reset, TLS handshake).
		return probeNetErr
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode == http.StatusOK {
		return probeFound
	}
	return probeMissing
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

// HandleWUTemperatureImages probes the Weather Underground static-map host in
// parallel and returns the most recent available image URL per region as JSON.
// Results are memoized for wuResultTTL to keep the fan-out off the hot path.
func HandleWUTemperatureImages(w http.ResponseWriter, r *http.Request) error {
	paths := cachedWUPaths()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	if err := json.NewEncoder(w).Encode(paths); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to encode WU image paths: %w", err))
	}
	return nil
}

func cachedWUPaths() map[string]string {
	wuMu.Lock()
	if wuResult != nil && time.Since(wuResultAt) < wuResultTTL {
		paths := make(map[string]string, len(wuResult))
		for k, v := range wuResult {
			paths[k] = v
		}
		wuMu.Unlock()
		return paths
	}
	wuMu.Unlock()

	paths := NewSWXCOFiles().getImagePaths()

	wuMu.Lock()
	wuResult = make(map[string]string, len(paths))
	for k, v := range paths {
		wuResult[k] = v
	}
	wuResultAt = time.Now()
	wuMu.Unlock()
	return paths
}
