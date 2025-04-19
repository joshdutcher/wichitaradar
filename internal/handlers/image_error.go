package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"wichitaradar/internal/middleware"
)

type ImageError struct {
	Src       string `json:"src"`
	Alt       string `json:"alt"`
	Page      string `json:"page"`
	Timestamp string `json:"timestamp"`
}

func HandleImageError(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return middleware.BadRequestError(fmt.Errorf("method not allowed"), "Method not allowed")
	}

	var errorData ImageError
	if err := json.NewDecoder(r.Body).Decode(&errorData); err != nil {
		return middleware.BadRequestError(err, "Invalid request body")
	}

	// Create logs directory if it doesn't exist
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to create logs directory: %w", err))
	}

	// Log to a daily file
	today := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, fmt.Sprintf("image-errors-%s.log", today))

	// Format the error message
	errorMsg := fmt.Sprintf("[%s] Page: %s, Image: %s, Alt: %s\n",
		errorData.Timestamp,
		errorData.Page,
		errorData.Src,
		errorData.Alt)

	// Append to log file
	if err := os.WriteFile(logFile, []byte(errorMsg), 0644); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to write to log file: %w", err))
	}

	w.WriteHeader(http.StatusOK)
	return nil
}