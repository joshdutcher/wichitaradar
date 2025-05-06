package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"wichitaradar/internal/middleware"

	"github.com/getsentry/sentry-go"
)

type ImageError struct {
	Src       string `json:"src"`
	Alt       string `json:"alt"`
	Page      string `json:"page"`
	Timestamp string `json:"timestamp"`
}

// Global state for tracking image failures
var (
	failedImages     = make(map[string]*ImageFailure)
	failedImagesLock sync.RWMutex
)

type ImageFailure struct {
	FirstFailure time.Time
	LastFailure  time.Time
	LastReported time.Time
	FailureCount int
}

const (
	FailureThreshold = 4 * time.Hour
	CheckInterval    = 5 * time.Minute
)

func init() {
	// Start background checker
	go checkPersistentFailures()
}

func checkPersistentFailures() {
	ticker := time.NewTicker(CheckInterval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		failedImagesLock.Lock()
		for src, failure := range failedImages {
			// Only report if:
			// 1. The image has been failing for more than FailureThreshold
			// 2. We haven't reported it recently
			if now.Sub(failure.FirstFailure) >= FailureThreshold &&
				now.Sub(failure.LastReported) >= FailureThreshold {

				// Log to daily file
				today := now.Format("2006-01-02")
				logFile := filepath.Join("logs", fmt.Sprintf("persistent-image-errors-%s.log", today))

				errorMsg := fmt.Sprintf("[%s] Persistent failure detected:\n"+
					"Image: %s\n"+
					"First failure: %s\n"+
					"Last failure: %s\n"+
					"Duration: %s\n"+
					"Failure count: %d\n\n",
					now.Format(time.RFC3339),
					src,
					failure.FirstFailure.Format(time.RFC3339),
					failure.LastFailure.Format(time.RFC3339),
					now.Sub(failure.FirstFailure).Round(time.Minute),
					failure.FailureCount)

				if err := os.WriteFile(logFile, []byte(errorMsg), 0644); err != nil {
					fmt.Printf("Failed to write to log file: %v\n", err)
				}

				// Report to Sentry
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("image", src)
					scope.SetTag("duration", now.Sub(failure.FirstFailure).Round(time.Minute).String())
					scope.SetTag("failure_count", fmt.Sprintf("%d", failure.FailureCount))
					scope.SetExtra("first_failure", failure.FirstFailure.Format(time.RFC3339))
					scope.SetExtra("last_failure", failure.LastFailure.Format(time.RFC3339))
					sentry.CaptureMessage("Persistent image loading failure detected")
				})

				// Update last reported time
				failure.LastReported = now
			}
		}
		failedImagesLock.Unlock()
	}
}

func HandleImageError(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return middleware.BadRequestError(fmt.Errorf("method not allowed"), "Method not allowed")
	}

	var errorData ImageError
	if err := json.NewDecoder(r.Body).Decode(&errorData); err != nil {
		return middleware.BadRequestError(err, "Invalid request body")
	}

	// Update failure tracking
	failedImagesLock.Lock()
	defer failedImagesLock.Unlock()

	now := time.Now()
	if failure, exists := failedImages[errorData.Src]; exists {
		failure.LastFailure = now
		failure.FailureCount++
	} else {
		failedImages[errorData.Src] = &ImageFailure{
			FirstFailure: now,
			LastFailure:  now,
			FailureCount: 1,
		}
	}

	// Also log to daily file for all errors
	today := now.Format("2006-01-02")
	logFile := filepath.Join("logs", fmt.Sprintf("image-errors-%s.log", today))

	errorMsg := fmt.Sprintf("[%s] Page: %s, Image: %s, Alt: %s\n",
		errorData.Timestamp,
		errorData.Page,
		errorData.Src,
		errorData.Alt)

	if err := os.WriteFile(logFile, []byte(errorMsg), 0644); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to write to log file: %w", err))
	}

	w.WriteHeader(http.StatusOK)
	return nil
}