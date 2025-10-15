package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
)

// ==============================
// Configuration knobs (tweak me)
// ==============================

// Only these hosts are considered "ours" and eligible for logging/escalation.
// Add your CDN / subdomains here.
var allowedImageHosts = map[string]bool{
	"wichitaradar.com":     true,
	"www.wichitaradar.com": true,
	// External providers embedded in the app
	"img.shields.io":               true,
	"www.spc.noaa.gov":             true,
	"s.w-x.co":                     true,
	"sirocco.accuweather.com":      true,
	"media.psg.nexstardigital.net": true,
	"graphical.weather.gov":        true,
	"www.weather.gov":              true,
	"weather.gov":                  true,
	"radar.weather.gov":            true,
	"x-hv1.pivotalweather.com":     true,
	"cdn.star.nesdis.noaa.gov":     true,
}

// Ignore trivial 1×1 pixels (often trackers) to cut noise at the source.
const ignoreOneByOnePixels = true

// We normalize URLs by host + path and DROP query strings to avoid a million uniques.
const dropQueryStringsInKey = true

// How long a failure must persist (with no subsequent success) before we’ll consider escalating.
var failureThreshold = 4 * time.Hour // was ~1h; raise to 4–12h to be safer

// Minimum number of failures for a given normalized key before we escalate.
const reportMinCount = 10

// We won’t escalate the same key more than once within this cool-down window.
var escalateCooldown = 6 * time.Hour

// Background sweep cadence for persistent failures.
var sweepInterval = 2 * time.Minute

// ==============================
// Data structures
// ==============================

type imageErrorPayload struct {
	Src       string `json:"src"`
	Referrer  string `json:"referrer,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	Error     string `json:"error,omitempty"`
	// Optional client timestamp; we treat it as informational only.
	Timestamp string `json:"timestamp,omitempty"`
}

type imageSuccessPayload struct {
	Src string `json:"src"`
}

type failureRecord struct {
	Host         string
	Path         string
	OriginalSrc  string // last seen full src (for debugging)
	FirstFailure time.Time
	LastFailure  time.Time
	LastReported time.Time
	LastSuccess  time.Time
	FailureCount int
}

var (
	mu           sync.Mutex
	failedImages = make(map[string]*failureRecord) // key = normalized(host+path) OR full src
	once         sync.Once
)

// ==================================
// Public bootstrap: start the sweeper
// ==================================

func InitImageFailureMonitor() {
	once.Do(func() {
		go persistentFailureSweeper()
	})
}

// ==================================
// HTTP Handlers
// ==================================

// POST /image-error
// Body: { src, referrer?, userAgent?, width?, height?, error?, timestamp? }
func HandleImageError(w http.ResponseWriter, r *http.Request) error {
	var p imageErrorPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return nil
	}
	if p.Src == "" {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	u, ok := parseAndAllow(p.Src)
	if !ok {
		// Not our host (or bad URL) → ignore
		w.WriteHeader(http.StatusOK)
		return nil
	}

	// Optionally drop obvious tracker pixels early
	if ignoreOneByOnePixels && p.Width == 1 && p.Height == 1 {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	key := normalizeKey(u)

	now := time.Now()
	mu.Lock()
	rec := failedImages[key]
	if rec == nil {
		rec = &failureRecord{
			Host:         strings.ToLower(u.Host),
			Path:         u.Path,
			OriginalSrc:  p.Src,
			FirstFailure: now,
			LastFailure:  now,
			FailureCount: 1,
		}
		failedImages[key] = rec
	} else {
		rec.LastFailure = now
		rec.FailureCount++
		rec.OriginalSrc = p.Src // keep last-seen for context
	}
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	return nil
}

// POST /image-success
// Body: { src }
func HandleImageSuccess(w http.ResponseWriter, r *http.Request) error {
	var p imageSuccessPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return nil
	}
	if p.Src == "" {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	u, ok := parseAndAllow(p.Src)
	if !ok {
		// Success for a non-allowed host is irrelevant.
		w.WriteHeader(http.StatusOK)
		return nil
	}

	key := normalizeKey(u)
	now := time.Now()

	mu.Lock()
	if rec, exists := failedImages[key]; exists {
		rec.LastSuccess = now
		// If we’ve had a success after the failure window started, clear the record.
		if now.After(rec.FirstFailure) {
			delete(failedImages, key)
		}
	}
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	return nil
}

// ==============================
// Background Sweeper & Escalation
// ==============================

func persistentFailureSweeper() {
	t := time.NewTicker(sweepInterval)
	defer t.Stop()

	for range t.C {
		sweepPersistentFailures()
	}
}

func sweepPersistentFailures() {
	now := time.Now()

	var toEscalate []*failureRecord
	var toLog []*failureRecord

	mu.Lock()
	for key, rec := range failedImages {
		_ = key // key unused except for map ops

		// Basic guards: only our hosts; minimum count; cool-down respected
		if !allowedImageHosts[rec.Host] {
			continue
		}
		if rec.FailureCount < reportMinCount {
			continue
		}
		if !rec.LastReported.IsZero() && now.Sub(rec.LastReported) < escalateCooldown {
			continue
		}
		// Must have persisted for the threshold window
		if now.Sub(rec.FirstFailure) < failureThreshold {
			continue
		}
		// If a success ever occurred after the first failure, treat issue as resolved.
		if !rec.LastSuccess.IsZero() && rec.LastSuccess.After(rec.FirstFailure) {
			continue
		}

		// If we passed all checks, queue for escalation & logging
		toEscalate = append(toEscalate, rec)
		toLog = append(toLog, rec)
		// Mark "reported" stamp so we don't spam
		rec.LastReported = now
	}
	mu.Unlock()

	// Side-effect outside lock: file log + Sentry
	for _, rec := range toLog {
		logPersistent(rec, now)
	}
	for _, rec := range toEscalate {
		sendSentry(rec)
	}
}

func logPersistent(rec *failureRecord, now time.Time) {
	// Append to daily file, best-effort only.
	today := now.Format("2006-01-02")
	_ = os.MkdirAll("logs", 0o750)
	f := filepath.Join("logs", fmt.Sprintf("persistent-image-errors-%s.log", today))
	msg := fmt.Sprintf("[%s] Persistent failure detected\nHost: %s\nPath: %s\nLastSrc: %s\nFirstFailure: %s\nLastFailure: %s\nFailures: %d\n\n",
		now.Format(time.RFC3339),
		rec.Host,
		rec.Path,
		rec.OriginalSrc,
		rec.FirstFailure.Format(time.RFC3339),
		rec.LastFailure.Format(time.RFC3339),
		rec.FailureCount,
	)
	_ = appendFile(f, msg)
}

func sendSentry(rec *failureRecord) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetTag("image.host", rec.Host)
		scope.SetTag("image.path", rec.Path)
		scope.SetTag("image.src.last", rec.OriginalSrc)
		scope.SetLevel(sentry.LevelWarning)
		scope.SetFingerprint([]string{"image-load-persistent", rec.Host, rec.Path})
		sentry.CaptureMessage("Persistent image loading failure detected")
	})
}

// ==============================
// Helpers
// ==============================

func parseAndAllow(raw string) (*url.URL, bool) {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, false
	}
	host := strings.ToLower(u.Host)
	if !allowedImageHosts[host] {
		return nil, false
	}
	return u, true
}

func normalizeKey(u *url.URL) string {
	host := strings.ToLower(u.Host)
	path := u.Path
	if path == "" {
		path = "/"
	}
	if dropQueryStringsInKey {
		return host + "|" + path
	}
	return host + "|" + path + "?" + u.RawQuery
}

func appendFile(path, s string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		// Fall back to stderr — don't crash the sweeper for logging problems.
		log.Printf("appendFile: %v", err)
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = f.WriteString(s)
	return err
}
