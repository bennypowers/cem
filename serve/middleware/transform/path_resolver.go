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
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/middleware/types"
)

// PathResolver resolves TypeScript source files from JavaScript request paths.
// It handles both in-place compilation (co-located .ts/.js files) and
// src/dist separation patterns (e.g., /dist/foo.js -> /src/foo.ts).
type PathResolver struct {
	watchDir     string
	pathMappings map[string]string // e.g., {"/dist/": "/src/"}
	fs           platform.FileSystem
	logger       types.Logger
}

// NewPathResolver creates a new PathResolver with the given configuration.
func NewPathResolver(watchDir string, pathMappings map[string]string, fs platform.FileSystem, logger types.Logger) *PathResolver {
	return &PathResolver{
		watchDir:     watchDir,
		pathMappings: pathMappings,
		fs:           fs,
		logger:       logger,
	}
}

// ResolveSourcePath resolves a source file path using path mappings.
// This is the generic path resolution logic shared by TypeScript and CSS resolvers.
//
// Parameters:
//   - requestPath: The requested path (e.g., "/dist/components/button.js")
//   - sourceExt: The source file extension to look for (e.g., ".ts" or ".css")
//   - requestExt: The request extension to replace (e.g., ".js" or ".css")
//
// Returns the resolved source path, or empty string if not found.
//
// Resolution strategy:
// 1. Try co-located file (backward compatibility for in-place compilation)
// 2. Try each explicit path mapping (from tsconfig.json or user config)
func (pr *PathResolver) ResolveSourcePath(requestPath string, sourceExt string, requestExt string) string {
	// Skip internal /__cem/ paths
	if strings.HasPrefix(requestPath, "/__cem/") {
		return ""
	}

	// Check request has expected extension
	if !strings.HasSuffix(requestPath, requestExt) {
		return ""
	}

	// Replace request extension with source extension
	sourceFileName := requestPath[:len(requestPath)-len(requestExt)] + sourceExt

	// Strategy 1: Co-located file (backward compatibility)
	// This handles in-place compilation where source and output are in same directory
	if pr.fileExists(sourceFileName) {
		if pr.logger != nil {
			pr.logger.Debug("PathResolver: co-located source found: %s", sourceFileName)
		}
		return sourceFileName
	}

	// Strategy 2: Try each explicit path mapping
	if pr.pathMappings != nil {
		for urlPrefix, sourcePrefix := range pr.pathMappings {
			if after, ok := strings.CutPrefix(requestPath, urlPrefix); ok {
				// Replace URL prefix with source prefix
				relativePath := after
				sourcePath := filepath.ToSlash(filepath.Join(sourcePrefix, relativePath))

				// Replace request extension with source extension
				if strings.HasSuffix(sourcePath, requestExt) {
					sourcePath = sourcePath[:len(sourcePath)-len(requestExt)] + sourceExt
				}

				// Ensure leading slash for consistency
				if !strings.HasPrefix(sourcePath, "/") {
					sourcePath = "/" + sourcePath
				}

				if pr.fileExists(sourcePath) {
					if pr.logger != nil {
						pr.logger.Debug("PathResolver: mapped source found: %s -> %s (mapping: %s -> %s)",
							requestPath, sourcePath, urlPrefix, sourcePrefix)
					}
					return sourcePath
				}
			}
		}
	}

	// Not found via any strategy
	if pr.logger != nil {
		pr.logger.Debug("PathResolver: no source found for %s", requestPath)
	}
	return ""
}

// ResolveTsSource finds the .ts source file for a requested .js file.
// Returns the request path to the .ts file (relative to watchDir with leading slash),
// or empty string if not found.
func (pr *PathResolver) ResolveTsSource(requestedJsPath string) string {
	return pr.ResolveSourcePath(requestedJsPath, ".ts", ".js")
}

// fileExists checks if a file exists at the given request path (relative to watchDir).
// The path should have a leading slash and use forward slashes.
func (pr *PathResolver) fileExists(requestPath string) bool {
	// Strip leading slash and normalize path separators
	pathNorm := strings.TrimPrefix(requestPath, "/")
	pathNorm = filepath.Clean(filepath.FromSlash(pathNorm))

	// Join with watchDir
	fullPath := filepath.Join(pr.watchDir, pathNorm)

	// Security: Reject attempts to escape watchDir
	rel, err := filepath.Rel(pr.watchDir, fullPath)
	if err != nil || strings.HasPrefix(rel, "..") {
		return false
	}

	// Check if file exists
	if pr.fs == nil {
		return false
	}

	stat, err := pr.fs.Stat(fullPath)
	if err != nil {
		return false
	}

	// Must be a regular file, not a directory
	return !stat.IsDir()
}
