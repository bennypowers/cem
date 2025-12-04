/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package transform

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/types"
	DS "github.com/bmatcuk/doublestar/v4"
)

// CSSConfig holds configuration for CSS transformation
type CSSConfig struct {
	WatchDirFunc     func() string      // Function to get current watch directory
	Logger           types.Logger
	ErrorBroadcaster types.ErrorBroadcaster // Sends errors to browser error overlay
	ConfigFile       string             // Path to config file (for error reporting)
	Enabled          bool               // Enable/disable CSS transformation
	Include          []string           // Glob patterns to include (empty means all .css files)
	Exclude          []string           // Glob patterns to exclude
	FS               platform.FileSystem // Filesystem abstraction for testability
}

// shouldTransformCSS checks if a CSS file should be transformed based on include/exclude patterns
func shouldTransformCSS(cssPath string, include []string, exclude []string, logger types.Logger, errorBroadcaster types.ErrorBroadcaster, configFile string) bool {
	// Opt-in behavior: if no include patterns specified, don't transform
	if len(include) == 0 {
		return false
	}

	// If include patterns are specified, file must match at least one
	if len(include) > 0 {
		matched := false
		for _, pattern := range include {
			match, err := DS.Match(pattern, cssPath)
			if err != nil {
				// Invalid glob pattern - log and broadcast error
				errMsg := fmt.Sprintf("Invalid CSS include pattern '%s': %v", pattern, err)
				if logger != nil {
					logger.Error(errMsg)
				}
				if errorBroadcaster != nil {
					errorBroadcaster.BroadcastError(
						"CSS Transform Config Error",
						errMsg,
						"", // No specific file for config errors
					)
				}
				// Treat invalid pattern as non-matching
				continue
			}
			if match {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	// If exclude patterns are specified, file must not match any
	if len(exclude) > 0 {
		for _, pattern := range exclude {
			match, err := DS.Match(pattern, cssPath)
			if err != nil {
				// Invalid glob pattern - log and broadcast error
				errMsg := fmt.Sprintf("Invalid CSS exclude pattern '%s': %v", pattern, err)
				if logger != nil {
					logger.Error(errMsg)
				}
				if errorBroadcaster != nil {
					errorBroadcaster.BroadcastError(
						"CSS Transform Config Error",
						errMsg,
						configFile,
					)
				}
				// Treat invalid pattern as non-matching (don't exclude)
				continue
			}
			if match {
				return false
			}
		}
	}

	// File matched include patterns and didn't match exclude patterns - transform it
	return true
}

// NewCSS creates a middleware that transforms CSS files to JavaScript modules
func NewCSS(config CSSConfig) middleware.Middleware {
	// Default to real filesystem if not provided
	fs := config.FS
	if fs == nil {
		fs = platform.NewOSFileSystem()
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip if transformation is disabled
			if !config.Enabled {
				next.ServeHTTP(w, r)
				return
			}

			// Get watch directory dynamically
			watchDir := config.WatchDirFunc()
			if watchDir == "" {
				// No watch directory set yet, pass to next handler
				next.ServeHTTP(w, r)
				return
			}

			requestPath := r.URL.Path

			// Only handle .css files
			if filepath.Ext(requestPath) != ".css" {
				next.ServeHTTP(w, r)
				return
			}

			// Strip leading slash and normalize
			watchDirClean := filepath.Clean(watchDir)
			cssPath := strings.TrimPrefix(requestPath, "/")
			cssPath = filepath.Clean(filepath.FromSlash(cssPath))
			fullCssPath := filepath.Join(watchDirClean, cssPath)

			// Reject attempts to escape the watch directory
			if rel, err := filepath.Rel(watchDirClean, fullCssPath); err != nil || strings.HasPrefix(rel, "..") {
				http.NotFound(w, r)
				return
			}

			// Check if this CSS file has import attributes query parameter
			// If so, we should always transform it regardless of include/exclude patterns
			hasImportAttrs := false
			queryParams := r.URL.Query()
			for key := range queryParams {
				if strings.HasPrefix(key, "__cem-import-attrs[") {
					hasImportAttrs = true
					break
				}
			}

			// If no import attributes, check include/exclude patterns
			if !hasImportAttrs && !shouldTransformCSS(cssPath, config.Include, config.Exclude, config.Logger, config.ErrorBroadcaster, config.ConfigFile) {
				// Don't transform this CSS file, pass to next handler
				next.ServeHTTP(w, r)
				return
			}

			// Read CSS file
			source, err := fs.ReadFile(fullCssPath)
			if err != nil {
				// File doesn't exist - let next handler handle 404
				next.ServeHTTP(w, r)
				return
			}

			// Transform CSS to JavaScript module
			transformed := TransformCSS(source, requestPath)

			// Serve as JavaScript module
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
			if _, err := w.Write([]byte(transformed)); err != nil {
				config.Logger.Error("Failed to write CSS transform response: %v", err)
			}
		})
	}
}
