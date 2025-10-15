package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestParseAndAllow(t *testing.T) {
	tests := []struct {
		name       string
		rawURL     string
		wantHost   string
		wantOK     bool
		wantScheme string
	}{
		{
			name:       "valid wichitaradar.com URL",
			rawURL:     "https://wichitaradar.com/image.png",
			wantHost:   "wichitaradar.com",
			wantOK:     true,
			wantScheme: "https",
		},
		{
			name:       "valid www.wichitaradar.com URL",
			rawURL:     "https://www.wichitaradar.com/image.png",
			wantHost:   "www.wichitaradar.com",
			wantOK:     true,
			wantScheme: "https",
		},
		{
			name:       "valid static.wichitaradar.com URL",
			rawURL:     "https://static.wichitaradar.com/image.png",
			wantHost:   "static.wichitaradar.com",
			wantOK:     true,
			wantScheme: "https",
		},
		{
			name:    "third-party domain rejected",
			rawURL:  "https://nickletto.com/metric",
			wantOK:  false,
		},
		{
			name:    "tracking pixel domain rejected",
			rawURL:  "https://boxclone.com/tracking.gif",
			wantOK:  false,
		},
		{
			name:    "invalid URL rejected",
			rawURL:  "not-a-valid-url",
			wantOK:  false,
		},
		{
			name:    "relative URL rejected",
			rawURL:  "/relative/path.png",
			wantOK:  false,
		},
		{
			name:    "URL without scheme rejected",
			rawURL:  "wichitaradar.com/image.png",
			wantOK:  false,
		},
		{
			name:    "empty URL rejected",
			rawURL:  "",
			wantOK:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, gotOK := parseAndAllow(tt.rawURL)

			if gotOK != tt.wantOK {
				t.Errorf("parseAndAllow() ok = %v, want %v", gotOK, tt.wantOK)
				return
			}

			if tt.wantOK {
				if gotURL == nil {
					t.Fatal("parseAndAllow() returned nil URL when ok=true")
				}
				if strings.ToLower(gotURL.Host) != tt.wantHost {
					t.Errorf("parseAndAllow() host = %v, want %v", gotURL.Host, tt.wantHost)
				}
				if gotURL.Scheme != tt.wantScheme {
					t.Errorf("parseAndAllow() scheme = %v, want %v", gotURL.Scheme, tt.wantScheme)
				}
			}
		})
	}
}

func TestNormalizeKey(t *testing.T) {
	tests := []struct {
		name    string
		rawURL  string
		want    string
	}{
		{
			name:    "URL with query string normalized",
			rawURL:  "https://wichitaradar.com/image.png?cachebuster=123",
			want:    "wichitaradar.com|/image.png",
		},
		{
			name:    "URL without query string",
			rawURL:  "https://www.wichitaradar.com/path/image.jpg",
			want:    "www.wichitaradar.com|/path/image.jpg",
		},
		{
			name:    "URL with multiple query params",
			rawURL:  "https://static.wichitaradar.com/img.png?v=1&t=2",
			want:    "static.wichitaradar.com|/img.png",
		},
		{
			name:    "root path URL",
			rawURL:  "https://wichitaradar.com",
			want:    "wichitaradar.com|/",
		},
		{
			name:    "URL with fragment",
			rawURL:  "https://wichitaradar.com/page#section",
			want:    "wichitaradar.com|/page",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.Parse(tt.rawURL)
			if err != nil {
				t.Fatalf("url.Parse() error = %v", err)
			}

			got := normalizeKey(u)
			if got != tt.want {
				t.Errorf("normalizeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleImageError(t *testing.T) {
	// Reset failedImages map before each test
	mu.Lock()
	failedImages = make(map[string]*failureRecord)
	mu.Unlock()

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		checkState     func(*testing.T)
	}{
		{
			name: "valid error for allowed host",
			requestBody: `{
				"src": "https://wichitaradar.com/image.png",
				"referrer": "https://wichitaradar.com/",
				"userAgent": "test-agent",
				"width": 100,
				"height": 100,
				"error": "Failed to load"
			}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				mu.Lock()
				defer mu.Unlock()
				key := "wichitaradar.com|/image.png"
				rec, exists := failedImages[key]
				if !exists {
					t.Error("Expected failure record to be created")
					return
				}
				if rec.FailureCount != 1 {
					t.Errorf("FailureCount = %d, want 1", rec.FailureCount)
				}
				if rec.Host != "wichitaradar.com" {
					t.Errorf("Host = %s, want wichitaradar.com", rec.Host)
				}
			},
		},
		{
			name: "error with 1x1 pixel filtered",
			requestBody: `{
				"src": "https://wichitaradar.com/tracking.png",
				"width": 1,
				"height": 1
			}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				mu.Lock()
				defer mu.Unlock()
				if len(failedImages) != 0 {
					t.Error("Expected 1x1 pixel to be filtered out")
				}
			},
		},
		{
			name: "error for disallowed host ignored",
			requestBody: `{
				"src": "https://nickletto.com/metric.gif",
				"width": 100,
				"height": 100
			}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				mu.Lock()
				defer mu.Unlock()
				if len(failedImages) != 0 {
					t.Error("Expected third-party domain to be ignored")
				}
			},
		},
		{
			name:           "empty src ignored",
			requestBody:    `{"src": ""}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				mu.Lock()
				defer mu.Unlock()
				if len(failedImages) != 0 {
					t.Error("Expected empty src to be ignored")
				}
			},
		},
		{
			name:           "invalid JSON returns bad request",
			requestBody:    `{invalid json}`,
			expectedStatus: http.StatusBadRequest,
			checkState:     nil,
		},
		{
			name: "multiple errors increment count",
			requestBody: `{
				"src": "https://www.wichitaradar.com/test.jpg",
				"width": 200,
				"height": 150
			}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				// Send second error for same image
				payload := imageErrorPayload{
					Src:    "https://www.wichitaradar.com/test.jpg",
					Width:  200,
					Height: 150,
				}
				body, _ := json.Marshal(payload)
				req := httptest.NewRequest("POST", "/api/image-error", bytes.NewReader(body))
				w := httptest.NewRecorder()
				_ = HandleImageError(w, req)

				mu.Lock()
				defer mu.Unlock()
				key := "www.wichitaradar.com|/test.jpg"
				rec, exists := failedImages[key]
				if !exists {
					t.Error("Expected failure record to exist")
					return
				}
				if rec.FailureCount != 2 {
					t.Errorf("FailureCount = %d, want 2", rec.FailureCount)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset state before each test
			mu.Lock()
			failedImages = make(map[string]*failureRecord)
			mu.Unlock()

			req := httptest.NewRequest("POST", "/api/image-error", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			err := HandleImageError(w, req)
			if err != nil {
				t.Fatalf("HandleImageError() error = %v", err)
			}

			if w.Code != tt.expectedStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.expectedStatus)
			}

			if tt.checkState != nil {
				tt.checkState(t)
			}
		})
	}
}

func TestHandleImageSuccess(t *testing.T) {
	tests := []struct {
		name           string
		setupState     func()
		requestBody    string
		expectedStatus int
		checkState     func(*testing.T)
	}{
		{
			name: "success clears failure record",
			setupState: func() {
				mu.Lock()
				defer mu.Unlock()
				failedImages["wichitaradar.com|/image.png"] = &failureRecord{
					Host:         "wichitaradar.com",
					Path:         "/image.png",
					FirstFailure: time.Now().Add(-1 * time.Hour),
					FailureCount: 5,
				}
			},
			requestBody:    `{"src": "https://wichitaradar.com/image.png"}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				mu.Lock()
				defer mu.Unlock()
				key := "wichitaradar.com|/image.png"
				if _, exists := failedImages[key]; exists {
					t.Error("Expected failure record to be cleared after success")
				}
			},
		},
		{
			name:           "success for non-existent failure is no-op",
			setupState:     func() {},
			requestBody:    `{"src": "https://wichitaradar.com/never-failed.png"}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				mu.Lock()
				defer mu.Unlock()
				if len(failedImages) != 0 {
					t.Error("Expected no state changes for non-existent failure")
				}
			},
		},
		{
			name:           "success for disallowed host ignored",
			setupState:     func() {},
			requestBody:    `{"src": "https://external.com/image.png"}`,
			expectedStatus: http.StatusOK,
			checkState: func(t *testing.T) {
				mu.Lock()
				defer mu.Unlock()
				if len(failedImages) != 0 {
					t.Error("Expected third-party success to be ignored")
				}
			},
		},
		{
			name:           "empty src ignored",
			setupState:     func() {},
			requestBody:    `{"src": ""}`,
			expectedStatus: http.StatusOK,
			checkState:     nil,
		},
		{
			name:           "invalid JSON returns bad request",
			setupState:     func() {},
			requestBody:    `{bad json}`,
			expectedStatus: http.StatusBadRequest,
			checkState:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset state
			mu.Lock()
			failedImages = make(map[string]*failureRecord)
			mu.Unlock()

			// Setup test state
			if tt.setupState != nil {
				tt.setupState()
			}

			req := httptest.NewRequest("POST", "/api/image-success", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			err := HandleImageSuccess(w, req)
			if err != nil {
				t.Fatalf("HandleImageSuccess() error = %v", err)
			}

			if w.Code != tt.expectedStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.expectedStatus)
			}

			if tt.checkState != nil {
				tt.checkState(t)
			}
		})
	}
}

func TestSweepPersistentFailures(t *testing.T) {
	// Save original config
	origThreshold := failureThreshold
	origCooldown := escalateCooldown

	// Use shorter durations for testing
	failureThreshold = 100 * time.Millisecond
	escalateCooldown = 50 * time.Millisecond

	// Restore original config after test
	defer func() {
		failureThreshold = origThreshold
		escalateCooldown = origCooldown
	}()

	tests := []struct {
		name       string
		setupState func()
		wantLog    bool
		wantSentry bool
	}{
		{
			name: "recent failure not escalated",
			setupState: func() {
				mu.Lock()
				defer mu.Unlock()
				failedImages["wichitaradar.com|/recent.png"] = &failureRecord{
					Host:         "wichitaradar.com",
					Path:         "/recent.png",
					FirstFailure: time.Now(),
					LastFailure:  time.Now(),
					FailureCount: reportMinCount + 1,
				}
			},
			wantLog:    false,
			wantSentry: false,
		},
		{
			name: "persistent failure escalated",
			setupState: func() {
				mu.Lock()
				defer mu.Unlock()
				failedImages["wichitaradar.com|/persistent.png"] = &failureRecord{
					Host:         "wichitaradar.com",
					Path:         "/persistent.png",
					FirstFailure: time.Now().Add(-200 * time.Millisecond),
					LastFailure:  time.Now(),
					FailureCount: reportMinCount + 5,
				}
			},
			wantLog:    true,
			wantSentry: true,
		},
		{
			name: "failure below minimum count not escalated",
			setupState: func() {
				mu.Lock()
				defer mu.Unlock()
				failedImages["wichitaradar.com|/low-count.png"] = &failureRecord{
					Host:         "wichitaradar.com",
					Path:         "/low-count.png",
					FirstFailure: time.Now().Add(-200 * time.Millisecond),
					LastFailure:  time.Now(),
					FailureCount: reportMinCount - 1,
				}
			},
			wantLog:    false,
			wantSentry: false,
		},
		{
			name: "resolved failure not escalated",
			setupState: func() {
				mu.Lock()
				defer mu.Unlock()
				failedImages["wichitaradar.com|/resolved.png"] = &failureRecord{
					Host:         "wichitaradar.com",
					Path:         "/resolved.png",
					FirstFailure: time.Now().Add(-200 * time.Millisecond),
					LastFailure:  time.Now().Add(-100 * time.Millisecond),
					LastSuccess:  time.Now(), // Success after failure
					FailureCount: reportMinCount + 1,
				}
			},
			wantLog:    false,
			wantSentry: false,
		},
		{
			name: "cooldown period prevents re-escalation",
			setupState: func() {
				mu.Lock()
				defer mu.Unlock()
				failedImages["wichitaradar.com|/cooldown.png"] = &failureRecord{
					Host:         "wichitaradar.com",
					Path:         "/cooldown.png",
					FirstFailure: time.Now().Add(-200 * time.Millisecond),
					LastFailure:  time.Now(),
					LastReported: time.Now().Add(-40 * time.Millisecond), // Within cooldown (50ms)
					FailureCount: reportMinCount + 1,
				}
			},
			wantLog:    false,
			wantSentry: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset state
			mu.Lock()
			failedImages = make(map[string]*failureRecord)
			mu.Unlock()

			// Setup test state
			if tt.setupState != nil {
				tt.setupState()
			}

			// Capture LastReported timestamps before sweep
			mu.Lock()
			beforeReported := make(map[string]time.Time)
			for key, rec := range failedImages {
				beforeReported[key] = rec.LastReported
			}
			mu.Unlock()

			// Run sweep
			sweepPersistentFailures()

			// Check if escalation occurred by comparing LastReported before/after
			mu.Lock()
			defer mu.Unlock()
			escalated := false
			for key, rec := range failedImages {
				// If LastReported changed, an escalation occurred
				if beforeTime := beforeReported[key]; !rec.LastReported.Equal(beforeTime) {
					escalated = true
					break
				}
			}

			if escalated != tt.wantSentry {
				t.Errorf("escalation = %v, want %v", escalated, tt.wantSentry)
			}
		})
	}
}

func TestInitImageFailureMonitor(t *testing.T) {
	// Test that InitImageFailureMonitor can be called multiple times safely
	InitImageFailureMonitor()
	InitImageFailureMonitor()
	InitImageFailureMonitor()

	// If we get here without panicking, the sync.Once protection works
	t.Log("InitImageFailureMonitor successfully handled multiple calls")
}
