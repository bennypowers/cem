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
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/middleware/types"
	"github.com/dunglas/go-urlpattern"
	"github.com/gosimple/slug"
)

// According to the WHATWG URLPattern specification, the URLPattern constructor requires
// a valid absolute base URL to resolve relative patterns, even if only pattern matching
// (and not actual URL resolution) is needed. We use the RFC-defined example domain
// "https://example.com" as a standard placeholder, since it is guaranteed to be a valid,
// non-resolvable URL. The actual value does not matter as long as it is a valid absolute URL.
const urlPatternBaseURL = "https://example.com"

// patternMapping represents a compiled URL pattern and Go template mapping
type patternMapping struct {
	pattern  *urlpattern.URLPattern // Compiled URL pattern
	template *template.Template      // Compiled Go template
	fromStr  string                  // Original pattern string (for logging)
	toStr    string                  // Original template string (for logging)
}

// PathResolver resolves TypeScript source files from JavaScript request paths.
// It handles both in-place compilation (co-located .ts/.js files) and
// pattern-based path mappings (e.g., /elements/:slug/:rest* -> /elements/rh-{{.slug}}/{{.rest}}).
type PathResolver struct {
	watchDir        string
	patternMappings []patternMapping // All path mappings as URLPattern + template
	fs              platform.FileSystem
	logger          types.Logger
}

// NewPathResolver creates a new PathResolver with the given configuration.
// All URL rewrites are compiled as URLPattern + template pairs during initialization.
func NewPathResolver(watchDir string, urlRewrites []config.URLRewrite, fs platform.FileSystem, logger types.Logger) *PathResolver {
	pr := &PathResolver{
		watchDir:        watchDir,
		patternMappings: make([]patternMapping, 0, len(urlRewrites)),
		fs:              fs,
		logger:          logger,
	}

	// Compile all URL rewrites as URLPattern + template pairs
	for _, rewrite := range urlRewrites {
		pm, err := compilePatternMapping(rewrite.URLPattern, rewrite.URLTemplate)
		if err != nil {
			// Log error but continue - validation should catch this earlier
			if logger != nil {
				logger.Debug("PathResolver: skipping invalid URL rewrite %s -> %s: %v", rewrite.URLPattern, rewrite.URLTemplate, err)
			}
			continue
		}
		pr.patternMappings = append(pr.patternMappings, *pm)
	}

	return pr
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
// For extensionless files (requestExt == ""), the path is used as-is without modification.
// This allows pattern matching on paths like /README or /docs/guide.
//
// Resolution strategy:
// 1. Try co-located file (in-place compilation)
// 2. Try pattern-based path mappings (URLPattern + templates)
func (pr *PathResolver) ResolveSourcePath(requestPath string, sourceExt string, requestExt string) string {
	// Skip internal /__cem/ paths
	if strings.HasPrefix(requestPath, "/__cem/") {
		return ""
	}

	// Check request has expected extension (empty string is valid - matches extensionless files)
	if !strings.HasSuffix(requestPath, requestExt) {
		return ""
	}

	// Replace request extension with source extension
	// For extensionless files (requestExt == ""), this is a no-op that returns requestPath unchanged
	sourceFileName := requestPath[:len(requestPath)-len(requestExt)] + sourceExt

	// Strategy 1: Co-located file
	// This handles in-place compilation where source and output are in same directory
	if pr.fileExists(sourceFileName) {
		if pr.logger != nil {
			pr.logger.Debug("PathResolver: co-located source found: %s", sourceFileName)
		}
		return sourceFileName
	}

	// Strategy 2: Try pattern-based path mappings
	for _, pm := range pr.patternMappings {
		resolvedPath, ok := pr.resolvePatternMapping(requestPath, &pm)
		if !ok {
			continue
		}

		// Replace request extension with source extension if needed
		if sourceExt != requestExt && strings.HasSuffix(resolvedPath, requestExt) {
			resolvedPath = resolvedPath[:len(resolvedPath)-len(requestExt)] + sourceExt
		}

		// Check if the resolved file exists
		if pr.fileExists(resolvedPath) {
			if pr.logger != nil {
				pr.logger.Debug("PathResolver: pattern-mapped source found: %s -> %s (pattern: %s -> %s)",
					requestPath, resolvedPath, pm.fromStr, pm.toStr)
			}
			return resolvedPath
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

// compilePatternMapping creates a patternMapping from pattern/template strings.
// Both the pattern and template are compiled and validated.
func compilePatternMapping(from, to string) (*patternMapping, error) {
	// Compile URLPattern
	pattern, err := urlpattern.New(from, urlPatternBaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("invalid URL pattern %q: %w", from, err)
	}

	// Create template with allowed functions
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"slug":  slug.Make, // Reuse slug function from demo discovery
	}

	// Compile Go template
	tmpl, err := template.New("pathMapping").Funcs(funcMap).Parse(to)
	if err != nil {
		return nil, fmt.Errorf("invalid template %q: %w", to, err)
	}

	return &patternMapping{
		pattern:  pattern,
		template: tmpl,
		fromStr:  from,
		toStr:    to,
	}, nil
}

// resolvePatternMapping matches a request path against a pattern and executes the template.
// Returns the resolved path and true if successful, empty string and false otherwise.
func (pr *PathResolver) resolvePatternMapping(requestPath string, pm *patternMapping) (string, bool) {
	// Match pattern using URLPattern
	testURL := urlPatternBaseURL + requestPath
	result := pm.pattern.Exec(testURL, "")
	if result == nil {
		return "", false // Pattern didn't match
	}

	// Execute template with captured params
	data := make(map[string]any)
	for key, value := range result.Pathname.Groups {
		data[key] = value
	}

	var buf bytes.Buffer
	if err := pm.template.Execute(&buf, data); err != nil {
		if pr.logger != nil {
			pr.logger.Warning("PathResolver: template execution failed for %s (pattern: %s): %v",
				requestPath, pm.fromStr, err)
		}
		return "", false
	}

	resolvedPath := buf.String()

	// Ensure leading slash for consistency
	if !strings.HasPrefix(resolvedPath, "/") {
		resolvedPath = "/" + resolvedPath
	}

	return resolvedPath, true
}

// ValidateURLRewrites validates all URL rewrites to ensure they can be compiled.
// This should be called during server initialization to fail fast on invalid patterns/templates.
func ValidateURLRewrites(urlRewrites []config.URLRewrite) error {
	for _, rewrite := range urlRewrites {
		if _, err := compilePatternMapping(rewrite.URLPattern, rewrite.URLTemplate); err != nil {
			return err
		}
	}
	return nil
}
