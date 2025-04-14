package handlers

import (
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"wichitaradar/internal/cache"
)

// Cache for XML files (15 minutes)
var xmlCache *cache.Cache

func init() {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	xmlCache = cache.New(cache.GetXMLCacheDir(workDir), 15*time.Minute)
}

// HandleXML handles requests for XML files
func HandleXML(w http.ResponseWriter, r *http.Request) {
	// Extract the XML path from the URL
	xmlPath := r.URL.Query().Get("path")
	if xmlPath == "" {
		http.Error(w, "Missing XML path", http.StatusBadRequest)
		return
	}

	// Construct the full URL with the correct base path
	url := "https://www.weather.gov/source/" + xmlPath

	// Get the cached XML file
	cacheFile := filepath.Join(xmlCache.GetCacheDir(), "wxstory.xml")
	if xmlCache.Expired("wxstory.xml") {
		if err := xmlCache.DownloadFile(url, "wxstory.xml", "https://www.weather.gov/ict/"); err != nil {
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
