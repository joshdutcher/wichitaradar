package handlers

import (
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"wichitaradar/internal/cache"
)

// HandleXML creates an HTTP handler func for XML requests, using the provided cache.
func HandleXML(xmlCache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the XML path from the URL
		xmlPath := r.URL.Query().Get("path")
		if xmlPath == "" {
			http.Error(w, "Missing XML path", http.StatusBadRequest)
			return
		}

		// Construct the full URL with the correct base path
		// Note: This handler is currently hardcoded for wxstory.xml
		// Consider making the filename dynamic if needed later.
		url := "https://www.weather.gov/source/" + xmlPath
		filename := "wxstory.xml" // Hardcoded for now

		// Get the cached XML file path using the provided cache instance
		cacheFile := filepath.Join(xmlCache.GetCacheDir(), filename)

		// Check expiry and download if needed using the provided cache instance
		if xmlCache.Expired(filename) {
			if err := xmlCache.DownloadFile(url, filename, "https://www.weather.gov/ict/"); err != nil {
				http.Error(w, "Failed to fetch XML", http.StatusInternalServerError)
				return
			}
		}

		// Read the cached file
		xmlData, err := os.ReadFile(cacheFile)
		if err != nil {
			http.Error(w, "Failed to read cached XML", http.StatusInternalServerError)
			return
		}

		// Log the XML file size
		log.Printf("XML file size: %d bytes", len(xmlData))

		// Convert XML data to string for easier manipulation
		xmlStr := string(xmlData)

		// Pre-process XML to handle common issues
		xmlStr = strings.ReplaceAll(xmlStr, "&", "&amp;")    // Fix unescaped ampersands
		xmlStr = strings.ReplaceAll(xmlStr, "<![CDATA[", "") // Remove CDATA markers
		xmlStr = strings.ReplaceAll(xmlStr, "]]>", "")

		// Create a more lenient XML decoder
		decoder := xml.NewDecoder(strings.NewReader(xmlStr))
		decoder.Strict = false
		decoder.AutoClose = xml.HTMLAutoClose
		decoder.Entity = xml.HTMLEntity

		// Parse the XML into a generic structure
		var result interface{}
		if err := decoder.Decode(&result); err != nil {
			log.Printf("Warning: XML parsing error: %v", err)
			// Return the original XML even if parsing fails
			w.Header().Set("Content-Type", "application/xml")
			w.Write(xmlData)
			return
		}

		// Return the original XML
		w.Header().Set("Content-Type", "application/xml")
		w.Write(xmlData)
	}
}
