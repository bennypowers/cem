package config

import (
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/platform"
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
func FindConfigFile(root string, fsys platform.FileSystem) string {
	for _, relPath := range ConfigPaths {
		absPath := filepath.Join(root, relPath)
		if info, err := fsys.Stat(absPath); err == nil && !info.IsDir() {
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
