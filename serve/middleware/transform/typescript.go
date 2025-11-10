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
	"os"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/serve/middleware"
)

// TypeScriptConfig holds configuration for TypeScript transformation
type TypeScriptConfig struct {
	WatchDirFunc     func() string // Function to get current watch directory
	TsconfigRawFunc  func() string // Function to get current tsconfig.json content
	Cache            *Cache
	Logger           Logger
	ErrorBroadcaster ErrorBroadcaster
	Target           string
}

// NewTypeScript creates a middleware that transforms TypeScript files to JavaScript
func NewTypeScript(config TypeScriptConfig) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
				tsPathNorm := strings.TrimPrefix(tsPath, "/")
				tsPathNorm = filepath.FromSlash(tsPathNorm)
				fullTsPath := filepath.Join(watchDir, tsPathNorm)

				// Get file stat for cache key
				fileInfo, err := os.Stat(fullTsPath)
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
					source, err := os.ReadFile(fullTsPath)
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
						target = "ES2022"
					}

					// Transform TypeScript to JavaScript
					result, err := TransformTypeScript(source, TransformOptions{
						Loader:      LoaderTS,
						Target:      Target(target),
						Sourcemap:   SourceMapInline,
						Sourcefile:  tsPathNorm, // Use normalized path for source maps
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
