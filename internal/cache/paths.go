package cache

import (
	"path/filepath"
)

// GetCacheDirs returns the paths to all cache directories relative to the working directory
func GetCacheDirs(workDir string) []string {
	return []string{
		filepath.Join(workDir, "scraped/xml"),
		filepath.Join(workDir, "scraped/images"),
		filepath.Join(workDir, "scraped/temp"),
	}
}

// GetXMLCacheDir returns the path to the XML cache directory
func GetXMLCacheDir(workDir string) string {
	return filepath.Join(workDir, "scraped/xml")
}

// GetImagesCacheDir returns the path to the images cache directory
func GetImagesCacheDir(workDir string) string {
	return filepath.Join(workDir, "scraped/images")
}

// GetTempCacheDir returns the path to the temporary cache directory
func GetTempCacheDir(workDir string) string {
	return filepath.Join(workDir, "scraped/temp")
}