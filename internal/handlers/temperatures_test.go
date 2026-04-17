package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"wichitaradar/internal/middleware"
	"wichitaradar/internal/testutils"
	"wichitaradar/pkg/templates"
)

func resetWUCache() {
	wuMu.Lock()
	defer wuMu.Unlock()
	wuResult = nil
	wuResultAt = time.Time{}
	wuLastGoodURLs = map[string]wuCacheEntry{}
}

func newTestSWXCO(prefix string) *SWXCOFiles {
	return &SWXCOFiles{
		runtime:     time.Now().UTC(),
		prefix:      prefix,
		images:      []string{"usa", "ddc"},
		imagePaths:  make(map[string]string),
		maxHoursAgo: 3,
	}
}

func TestGetImagePaths_FoundReturnsURL(t *testing.T) {
	resetWUCache()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	s := newTestSWXCO(srv.URL)
	paths := s.getImagePaths()
	for _, k := range s.images {
		if !strings.HasPrefix(paths[k], srv.URL) {
			t.Errorf("%s: expected URL under %s, got %q", k, srv.URL, paths[k])
		}
	}
}

func TestGetImagePaths_CleanMissReturnsEmpty(t *testing.T) {
	resetWUCache()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	s := newTestSWXCO(srv.URL)
	paths := s.getImagePaths()
	for _, k := range s.images {
		if paths[k] != "" {
			t.Errorf("%s: expected empty string on clean miss, got %q", k, paths[k])
		}
	}
}

func TestGetImagePaths_NetErrorFallsBackToCurrentHour(t *testing.T) {
	resetWUCache()
	// Port 1 reliably refuses connections → probeNetErr for every probe.
	s := newTestSWXCO("http://127.0.0.1:1")
	paths := s.getImagePaths()
	for _, k := range s.images {
		want := s.urlFor(k, s.runtime)
		if paths[k] != want {
			t.Errorf("%s: expected current-hour fallback %q, got %q", k, want, paths[k])
		}
	}
}

func TestGetImagePaths_NetErrorPrefersLastGood(t *testing.T) {
	resetWUCache()
	wuMu.Lock()
	wuLastGoodURLs["usa"] = wuCacheEntry{url: "https://example.invalid/usa-last-good.jpg", at: time.Now().UTC()}
	wuMu.Unlock()

	s := newTestSWXCO("http://127.0.0.1:1")
	paths := s.getImagePaths()
	if paths["usa"] != "https://example.invalid/usa-last-good.jpg" {
		t.Errorf("usa: expected last-good URL, got %q", paths["usa"])
	}
	// ddc has no last-good → current-hour fallback
	if paths["ddc"] != s.urlFor("ddc", s.runtime) {
		t.Errorf("ddc: expected current-hour fallback, got %q", paths["ddc"])
	}
}

func TestGetImagePaths_StaleLastGoodIgnored(t *testing.T) {
	resetWUCache()
	wuMu.Lock()
	wuLastGoodURLs["usa"] = wuCacheEntry{
		url: "https://example.invalid/usa-stale.jpg",
		at:  time.Now().UTC().Add(-wuLastGoodMaxAge - time.Minute),
	}
	wuMu.Unlock()

	s := newTestSWXCO("http://127.0.0.1:1")
	paths := s.getImagePaths()
	if paths["usa"] == "https://example.invalid/usa-stale.jpg" {
		t.Errorf("usa: expected stale last-good to be ignored")
	}
	if paths["usa"] != s.urlFor("usa", s.runtime) {
		t.Errorf("usa: expected current-hour fallback, got %q", paths["usa"])
	}
}

func TestCachedWUPaths_MemoizesResult(t *testing.T) {
	resetWUCache()
	var hits int32
	var hitsMu sync.Mutex
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hitsMu.Lock()
		hits++
		hitsMu.Unlock()
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	// Seed the cache by running a probe directly and stashing the result.
	s := newTestSWXCO(srv.URL)
	paths := s.getImagePaths()
	wuMu.Lock()
	wuResult = paths
	wuResultAt = time.Now()
	wuMu.Unlock()

	hitsMu.Lock()
	before := hits
	hitsMu.Unlock()

	for i := 0; i < 3; i++ {
		_ = cachedWUPaths()
	}

	hitsMu.Lock()
	after := hits
	hitsMu.Unlock()
	if after != before {
		t.Errorf("expected 0 additional upstream hits while cache fresh, got %d", after-before)
	}
}

func TestHandleTemperatures(t *testing.T) {
	tests := []struct {
		name           string
		setupTemplates func(*testing.T)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful render",
			setupTemplates: func(t *testing.T) {
				templates.Reset()
				testutils.InitTemplates(t)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Current Temps",
		},
		{
			name: "template error",
			setupTemplates: func(t *testing.T) {
				templates.Reset()
				// Don't initialize templates to force an error
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupTemplates != nil {
				tt.setupTemplates(t)
			}

			// Create request
			req := httptest.NewRequest("GET", "/temperatures", nil)
			w := httptest.NewRecorder()

			// Execute handler
			middleware.ErrorHandler(HandleTemperatures).ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Check response body
			body := w.Body.String()
			if !bytes.Contains([]byte(body), []byte(tt.expectedBody)) {
				t.Errorf("expected body to contain %q, got %q", tt.expectedBody, body)
			}
		})
	}
}
