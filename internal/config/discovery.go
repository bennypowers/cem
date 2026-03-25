package config

import (
	"os"
	"path/filepath"
	"strings"
)

// ConfigPaths lists the config file paths to search, in order of precedence.
var ConfigPaths = []string{
	".config/cem.yaml",
	".config/cem.yml",
	".config/cem.json",
	".config/cem.jsonc",
	".cem.yaml",
	".cem.yml",
	".cem.json",
	".cem.jsonc",
}

// FindConfigFile searches for a CEM config file in the given root directory.
// Returns the absolute path of the first match, or "" if none found.
func FindConfigFile(root string) string {
	for _, relPath := range ConfigPaths {
		absPath := filepath.Join(root, relPath)
		if _, err := os.Stat(absPath); err == nil {
			return absPath
		}
	}
	return ""
}

// FormatFromPath returns the config format based on file extension.
// Returns "yaml", "json", "jsonc", or "" for unknown extensions.
func FormatFromPath(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".yaml", ".yml":
		return "yaml"
	case ".json":
		return "json"
	case ".jsonc":
		return "jsonc"
	default:
		return ""
	}
}
