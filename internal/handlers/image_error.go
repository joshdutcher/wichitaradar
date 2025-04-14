package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type ImageError struct {
	Src       string `json:"src"`
	Alt       string `json:"alt"`
	Page      string `json:"page"`
	Timestamp string `json:"timestamp"`
}

func HandleImageError(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var errorData ImageError
	if err := json.NewDecoder(r.Body).Decode(&errorData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create logs directory if it doesn't exist
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		http.Error(w, "Failed to create logs directory", http.StatusInternalServerError)
		return
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
		http.Error(w, "Failed to write to log file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}