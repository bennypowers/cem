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
)

// DefaultTarget is the default TypeScript transformation target.
const DefaultTarget = "ES2022"

// TypeScriptConfig holds configuration for TypeScript transformation
type TypeScriptConfig struct {
	WatchDirFunc     func() string // Function to get current watch directory
	TsconfigRawFunc  func() string // Function to get current tsconfig.json content
	Cache            *Cache
	Logger           types.Logger
	ErrorBroadcaster types.ErrorBroadcaster
	Target           string
	Enabled          bool                // Enable/disable TypeScript transformation
	FS               platform.FileSystem // Filesystem abstraction for testability
}

// NewTypeScript creates a middleware that transforms TypeScript files to JavaScript
func NewTypeScript(config TypeScriptConfig) middleware.Middleware {
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
			ext := filepath.Ext(requestPath)

			// Handle both .js requests (with .ts source) and direct .ts requests
			var tsPath string
			switch ext {
			case ".js":
				// Check if .ts file exists for .js request
				tsPath = requestPath[:len(requestPath)-3] + ".ts"
			case ".ts":
				// Direct .ts request
				tsPath = requestPath
			}

			if tsPath != "" {
				// Strip leading slash and normalize path separators before joining
				watchDirClean := filepath.Clean(watchDir)
				tsPathNorm := strings.TrimPrefix(tsPath, "/")
				tsPathNorm = filepath.Clean(filepath.FromSlash(tsPathNorm))
				fullTsPath := filepath.Join(watchDirClean, tsPathNorm)

				// Reject attempts to escape the watch directory
				if rel, err := filepath.Rel(watchDirClean, fullTsPath); err != nil || strings.HasPrefix(rel, "..") {
					http.NotFound(w, r)
					return
				}

				// Get file stat for cache key
				fileInfo, err := fs.Stat(fullTsPath)
				if err == nil {
					// .ts file exists - check cache first
					cacheKey := CacheKey{
						Path:    fullTsPath,
						ModTime: fileInfo.ModTime(),
						Size:    fileInfo.Size(),
					}

					// Try to get from cache
					if cached, found := config.Cache.Get(cacheKey); found {
						config.Logger.Debug("Cache hit for %s", tsPathNorm)
						// Serve cached transformed JavaScript
						w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
						if _, err := w.Write(cached.Code); err != nil {
							config.Logger.Error("Failed to write cached transform response: %v", err)
						}
						return
					}

					// Cache miss - read file and transform
					config.Logger.Debug("Cache miss for %s", tsPathNorm)
					source, err := fs.ReadFile(fullTsPath)
					if err != nil {
						config.Logger.Error("Failed to read TypeScript file %s: %v", tsPathNorm, err)
						http.Error(w, "Failed to read file", http.StatusInternalServerError)
						return
					}

					// Get tsconfig for transform (dynamically from function)
					var tsconfigRaw string
					if config.TsconfigRawFunc != nil {
						tsconfigRaw = config.TsconfigRawFunc()
					}

					// Get configured target (defaults to ES2022 if not set)
					target := config.Target
					if target == "" {
						target = DefaultTarget
					}

					// Transform TypeScript to JavaScript
					result, err := TransformTypeScript(source, TransformOptions{
						Loader:      LoaderTS,
						Target:      Target(target),
						Sourcemap:   SourceMapInline,
						Sourcefile:  fullTsPath, // Use absolute path for dependency resolution
						TsconfigRaw: tsconfigRaw,
					})
					if err != nil {
						config.Logger.Error("Failed to transform TypeScript file %s: %v", tsPathNorm, err)

						// Broadcast error to browser overlay
						if config.ErrorBroadcaster != nil {
							config.ErrorBroadcaster.BroadcastError(
								"TypeScript Transform Error",
								err.Error(),
								tsPathNorm,
							)
						}

						http.Error(w, fmt.Sprintf("Transform error: %v", err), http.StatusInternalServerError)
						return
					}

					// Store in cache
					if len(result.Dependencies) > 0 {
						config.Logger.Debug("Caching %s with %d dependencies: %v", tsPathNorm, len(result.Dependencies), result.Dependencies)
					}
					config.Cache.Set(cacheKey, result.Code, result.Dependencies)

					// Serve transformed JavaScript
					w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
					if _, err := w.Write(result.Code); err != nil {
						config.Logger.Error("Failed to write transform response: %v", err)
					}
					return
				}
			}

			// Not a TypeScript file or file doesn't exist - pass to next handler
			next.ServeHTTP(w, r)
		})
	}
}
