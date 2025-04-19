package cache

import "path/filepath"

// GetCacheDirs returns all cache directories
func GetCacheDirs(projectRoot string) []string {
	return []string{
		GetXMLCacheDir(projectRoot),
		GetImagesCacheDir(projectRoot),
		GetTempCacheDir(projectRoot),
	}
}

// GetImagesCacheDir returns the path to the images cache directory
func GetImagesCacheDir(projectRoot string) string {
	return filepath.Join(projectRoot, "scraped", "images")
}

// GetAnimatedCacheDir returns the path to the animated images cache directory
func GetAnimatedCacheDir(projectRoot string) string {
	return filepath.Join(projectRoot, "scraped", "animated")
}

// GetXMLCacheDir returns the path to the XML cache directory
func GetXMLCacheDir(projectRoot string) string {
	return filepath.Join(projectRoot, "scraped", "xml")
}

// GetTempCacheDir returns the path to the temporary cache directory
func GetTempCacheDir(projectRoot string) string {
	return filepath.Join(projectRoot, "scraped", "temp")
}