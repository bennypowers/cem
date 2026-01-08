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

package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	V "bennypowers.dev/cem/internal/version"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
)

// notFoundDetector wraps http.ResponseWriter to detect and intercept 404 status codes
type notFoundDetector struct {
	http.ResponseWriter
	is404       bool
	wroteHeader bool
	buffer      []byte
}

func (n *notFoundDetector) WriteHeader(statusCode int) {
	if !n.wroteHeader {
		n.wroteHeader = true
		if statusCode == http.StatusNotFound {
			n.is404 = true
			// Don't write the header - we'll render our custom 404 page
			return
		}
		n.ResponseWriter.WriteHeader(statusCode)
	}
}

func (n *notFoundDetector) Write(b []byte) (int, error) {
	if !n.wroteHeader {
		n.WriteHeader(http.StatusOK)
	}
	if n.is404 {
		// Capture the 404 response but don't write it
		n.buffer = append(n.buffer, b...)
		return len(b), nil
	}
	return n.ResponseWriter.Write(b)
}

// Config holds configuration for the internal routes middleware
type Config struct {
	// Context provides access to dev server state
	Context middleware.DevServerContext

	// LogsFunc returns the current logs (if logger supports it)
	LogsFunc func() []logger.LogEntry

	// WebSocketHandler handles WebSocket upgrade requests for live reload
	WebSocketHandler http.HandlerFunc

	// Templates holds the parsed template registry with context
	Templates *TemplateRegistry
}

// New creates a middleware that handles internal CEM routes and demo routing
func New(config Config) middleware.Middleware {
	// Create template registry if not provided
	if config.Templates == nil {
		config.Templates = NewTemplateRegistry(config.Context)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Handle internal CEM routes first
			switch {
			case r.URL.Path == "/__cem/reload":
				// WebSocket live reload endpoint
				if config.WebSocketHandler != nil {
					config.WebSocketHandler(w, r)
				} else {
					http.NotFound(w, r)
				}
				return
			case r.URL.Path == "/__cem/logs":
				serveLogs(w, r, config)
				return
			case r.URL.Path == "/__cem/api/markdown":
				serveMarkdownAPI(w, r, config)
				return
			case r.URL.Path == "/__cem/debug":
				serveDebugInfo(w, r, config)
				return
			case strings.HasPrefix(r.URL.Path, "/__cem/"):
				serveInternalModules(w, r, config)
				return
			case r.URL.Path == "/custom-elements.json":
				serveManifest(w, r, config)
				return
			}

			// Try to serve as demo route
			if serveDemoRoute(w, r, config) {
				return
			}

			// Try to serve index listing (for root path)
			if serveIndexListing(w, r, config) {
				return
			}

			// Try to serve static assets from demo package directories (works in both workspace and single-package mode)
			if servePackageStaticAsset(w, r, config) {
				return
			}

			// Not a special route - try static files with 404 detection
			detector := &notFoundDetector{ResponseWriter: w}
			next.ServeHTTP(detector, r)

			// If static handler returned 404, show our custom 404 page
			if detector.is404 {
				serve404Page(w, r, config)
			}
		})
	}
}

// serveInternalModules serves embedded static assets from embed.FS
func serveInternalModules(w http.ResponseWriter, r *http.Request, config Config) {
	// Strip /__cem/ prefix to get the file path within the embedded FS
	// Request: /__cem/foo.js -> templates/js/foo.js in embed.FS
	// Request: /__cem/logo.svg -> templates/images/logo.svg in embed.FS
	reqPath := strings.TrimPrefix(r.URL.Path, "/__cem/")

	// Determine subdirectory based on file extension and path
	var path string
	switch {
	case strings.HasPrefix(reqPath, "elements/"):
		// Element files (e.g., elements/cem-drawer/cem-drawer.js)
		path = "templates/" + reqPath
	case strings.HasSuffix(reqPath, ".js"):
		path = "templates/js/" + reqPath
	case strings.HasSuffix(reqPath, ".css"):
		path = "templates/css/" + reqPath
	case strings.HasSuffix(reqPath, ".svg"), strings.HasSuffix(reqPath, ".png"), strings.HasSuffix(reqPath, ".jpg"), strings.HasSuffix(reqPath, ".jpeg"):
		path = "templates/images/" + reqPath
	default:
		// Try js as default for backward compatibility
		path = "templates/js/" + reqPath
	}

	data, err := InternalModules.ReadFile(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Set content type based on file extension
	contentType := "application/octet-stream"
	switch {
	case strings.HasSuffix(reqPath, ".js"):
		contentType = "application/javascript; charset=utf-8"
	case strings.HasSuffix(reqPath, ".css"):
		contentType = "text/css; charset=utf-8"
	case strings.HasSuffix(reqPath, ".html"):
		contentType = "text/html; charset=utf-8"
	case strings.HasSuffix(reqPath, ".svg"):
		contentType = "image/svg+xml"
	case strings.HasSuffix(reqPath, ".png"):
		contentType = "image/png"
	case strings.HasSuffix(reqPath, ".jpg"), strings.HasSuffix(reqPath, ".jpeg"):
		contentType = "image/jpeg"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "no-cache")
	if _, err := w.Write(data); err != nil {
		config.Context.Logger().Error("Failed to write static asset response: %v", err)
	}
}

// serveManifest serves the custom elements manifest
func serveManifest(
	w http.ResponseWriter,
	_ *http.Request,
	config Config,
) {
	var manifestBytes []byte

	if config.Context.IsWorkspace() {
		// Workspace mode: create workspace manifest with packages array
		middlewarePackages := config.Context.WorkspacePackages()

		// Parse each package manifest and add package name
		type PackageWithName struct {
			Name string `json:"name"`
			M.Package
		}

		packages := make([]PackageWithName, 0, len(middlewarePackages))

		for _, pkg := range middlewarePackages {
			if len(pkg.Manifest) == 0 {
				continue
			}

			var parsed M.Package
			if err := json.Unmarshal(pkg.Manifest, &parsed); err != nil {
				config.Context.Logger().Warning("Failed to parse manifest for package %s: %v", pkg.Name, err)
				continue
			}

			packages = append(packages, PackageWithName{
				Name:    pkg.Name,
				Package: parsed,
			})
		}

		if len(packages) == 0 {
			http.Error(w, "No manifests available in workspace", http.StatusNotFound)
			return
		}

		// Create workspace manifest structure
		workspaceManifest := map[string]any{
			"packages": packages,
		}

		// Serialize workspace manifest
		var err error
		manifestBytes, err = json.Marshal(workspaceManifest)
		if err != nil {
			config.Context.Logger().Error("Failed to serialize workspace manifest: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		// Single-package mode: use regular manifest
		var err error
		manifestBytes, err = config.Context.Manifest()
		if err != nil {
			config.Context.Logger().Error("Failed to get manifest: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if len(manifestBytes) == 0 {
			http.Error(w, "Manifest not available", http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(manifestBytes); err != nil {
		config.Context.Logger().Error("Failed to write manifest response: %v", err)
	}
}

// serveLogs serves the plain text logs for the debug console
func serveLogs(
	w http.ResponseWriter,
	_ *http.Request,
	config Config,
) {
	var logs []logger.LogEntry
	if config.LogsFunc != nil {
		logs = config.LogsFunc()
	}

	if logs == nil {
		logs = []logger.LogEntry{}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		config.Context.Logger().Error("Failed to encode logs response: %v", err)
	}
}

// serveDebugInfo serves debug information for the debug overlay
func serveDebugInfo(
	w http.ResponseWriter,
	_ *http.Request,
	config Config,
) {
	// Get watch directory from context
	watchDir := config.Context.WatchDir()

	// Get import map (pre-computed during server initialization)
	var importMapObj any = config.Context.ImportMap()

	// Get manifest for demo info
	manifestBytes, err := config.Context.Manifest()
	if err != nil {
		config.Context.Logger().Warning("Failed to get manifest for debug info: %v", err)
		manifestBytes = nil
	}

	// Parse manifest to get demo info
	var demos []map[string]any
	if len(manifestBytes) > 0 {
		var pkg M.Package
		if err := json.Unmarshal(manifestBytes, &pkg); err == nil {
			for _, renderableDemo := range pkg.RenderableDemos() {
				demoURL := renderableDemo.Demo.URL
				// Extract local route from canonical URL (strip origin)
				localRoute := demoURL
				if parsed, err := url.Parse(demoURL); err == nil && parsed.Path != "" {
					// Use just the path component for the local route
					localRoute = parsed.Path
				}

				demoInfo := map[string]any{
					"tagName":      renderableDemo.CustomElementDeclaration.TagName,
					"description":  renderableDemo.Demo.Description,
					"canonicalURL": demoURL,
					"localRoute":   localRoute,
				}

				// Add file path if available
				if renderableDemo.Demo.Source != nil && renderableDemo.Demo.Source.Href != "" {
					demoInfo["filePath"] = renderableDemo.Demo.Source.Href
				}

				demos = append(demos, demoInfo)
			}
		}
	}

	debugInfo := map[string]any{
		"watchDir":     watchDir,
		"manifestSize": fmt.Sprintf("%d bytes", len(manifestBytes)),
		"version":      V.GetVersion(),
		"os":           fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		"demos":        demos,
		"demoCount":    len(demos),
		"importMap":    importMapObj,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(debugInfo); err != nil {
		config.Context.Logger().Error("Failed to encode debug info response: %v", err)
	}
}

// serveIndexListing serves the index listing page for root path. Returns true if handled, false otherwise.
func serveIndexListing(w http.ResponseWriter, r *http.Request, config Config) bool {
	// Only handle root path
	if r.URL.Path != "/" {
		return false
	}

	config.Context.Logger().Debug("serveIndexListing: handling root path /")

	// Get import map as JSON
	var importMapJSON string
	if im := config.Context.ImportMap(); im != nil {
		if importMap, ok := im.(*importmappkg.ImportMap); ok {
			importMapJSON = importMap.ToJSON()
		}
	}

	// Extract state from request cookie for SSR
	state := GetStateFromRequest(r, config.Context.Logger())

	var html string
	var err error

	if config.Context.IsWorkspace() {
		// Workspace mode: always render workspace listing
		middlewarePackages := config.Context.WorkspacePackages()
		config.Context.Logger().Debug("serveIndexListing: IsWorkspace=true, got %d middleware packages", len(middlewarePackages))
		packages := make([]PackageContext, len(middlewarePackages))
		for i, pkg := range middlewarePackages {
			config.Context.Logger().Debug("serveIndexListing: Package %d: Name=%s, Manifest len=%d", i, pkg.Name, len(pkg.Manifest))
			packages[i] = PackageContext{
				Name:     pkg.Name,
				Path:     pkg.Path,
				Manifest: pkg.Manifest,
			}
		}
		html, err = RenderWorkspaceListing(config.Templates, config.Context, packages, importMapJSON, state)
	} else {
		config.Context.Logger().Debug("serveIndexListing: NOT in workspace mode, single-package mode")
		// Single-package mode: check if index.html exists
		watchDir := config.Context.WatchDir()
		if watchDir != "" {
			indexPath := filepath.Join(watchDir, "index.html")
			var err error

			// Use injected filesystem if available, otherwise fall back to os.Stat
			if filesystem := config.Context.FileSystem(); filesystem != nil {
				_, err = filesystem.Stat(indexPath)
			} else {
				_, err = os.Stat(indexPath)
			}

			if err == nil {
				// index.html exists, let static handler serve it
				return false
			}
		}

		// No index.html, render element listing
		manifestBytes, err2 := config.Context.Manifest()
		if err2 != nil {
			config.Context.Logger().Warning("Failed to get manifest for index listing: %v", err2)
			return false
		}
		config.Context.Logger().Debug("serveIndexListing: got manifest bytes, len=%d", len(manifestBytes))
		// Get package name from package.json
		pkgName := ""
		if pkg, err := config.Context.PackageJSON(); err == nil && pkg != nil {
			pkgName = pkg.Name
		}
		config.Context.Logger().Debug("serveIndexListing: calling RenderElementListing with package name=%s", pkgName)
		html, err = RenderElementListing(config.Templates, config.Context, manifestBytes, importMapJSON, pkgName, state)
		config.Context.Logger().Debug("serveIndexListing: RenderElementListing returned, err=%v, html len=%d", err, len(html))
	}

	if err != nil {
		config.Context.Logger().Error("Failed to render index listing: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return true
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(html)); err != nil {
		config.Context.Logger().Error("Failed to write index listing response: %v", err)
	}
	return true
}

// serveDemoRoute attempts to serve a demo route. Returns true if handled, false otherwise.
func serveDemoRoute(w http.ResponseWriter, r *http.Request, config Config) bool {
	// Parse query parameters
	queryParams := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}

	// Get pre-computed routing table from context (used in both workspace and single-package mode)
	routes := config.Context.DemoRoutes()
	if routes == nil {
		// No routes available - not a demo route
		return false
	}

	// Normalize requested path (ensure trailing slash for directory-style)
	normalizedPath := r.URL.Path
	if normalizedPath != "/" && normalizedPath[len(normalizedPath)-1] != '/' && filepath.Ext(normalizedPath) == "" {
		normalizedPath += "/"
	}

	// Look up route in table
	entry, found := routes[normalizedPath]
	if !found {
		// Not a demo route
		return false
	}

	// Debug: log entry details
	config.Context.Logger().Debug("Demo route entry: localRoute=%s, packagePath='%s', filePath='%s', isWorkspace=%v",
		entry.LocalRoute, entry.PackagePath, entry.FilePath, config.Context.IsWorkspace())

	// Build navigation data for the drawer
	var navigationHTML template.HTML
	var packageName string

	// Get package name from package.json for masthead
	if pkg, err := config.Context.PackageJSON(); err == nil && pkg != nil && pkg.Name != "" {
		packageName = pkg.Name
	}

	if config.Context.IsWorkspace() {
		// Workspace mode: build navigation from all packages
		middlewarePackages := config.Context.WorkspacePackages()
		packages := make([]PackageContext, len(middlewarePackages))
		for i, pkg := range middlewarePackages {
			packages[i] = PackageContext{
				Name:     pkg.Name,
				Path:     pkg.Path,
				Manifest: pkg.Manifest,
			}
		}
		var navErr error
		navigationHTML, navErr = BuildWorkspaceNavigation(config.Templates, packages, r.URL.Path)
		if navErr != nil {
			config.Context.Logger().Error("Failed to build workspace navigation for %s: %v", r.URL.Path, navErr)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return true
		}
		// Keep packageName from package.json, don't override with "Workspace"
	} else {
		// Single-package mode: build navigation from manifest
		manifestBytes, err := config.Context.Manifest()
		if err == nil && len(manifestBytes) > 0 {
			var navErr error
			navigationHTML, packageName, navErr = BuildSinglePackageNavigation(config.Templates, manifestBytes, packageName, r.URL.Path)
			if navErr != nil {
				config.Context.Logger().Error("Failed to build navigation: %v", navErr)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return true
			}
		}
	}

	// Render the demo
	html, err := renderDemoFromRoute(entry, queryParams, config, navigationHTML, packageName, r)
	if err != nil {
		config.Context.Logger().Error("Failed to render demo: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return true
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(html)); err != nil {
		config.Context.Logger().Error("Failed to write demo response: %v", err)
	}
	return true
}

// renderDemoFromRoute renders a demo page with chrome using routing table entry
func renderDemoFromRoute(entry *DemoRouteEntry, queryParams map[string]string, config Config, navigationHTML template.HTML, packageName string, r *http.Request) (string, error) {
	// Determine base directory for file path
	var baseDir string
	if config.Context.IsWorkspace() && entry.PackagePath != "" {
		// Workspace mode: use package directory
		baseDir = entry.PackagePath
		config.Context.Logger().Debug("Workspace mode demo: baseDir=%s, filePath=%s, packagePath=%s", baseDir, entry.FilePath, entry.PackagePath)
	} else {
		// Single-package mode: use watch directory
		baseDir = config.Context.WatchDir()
		config.Context.Logger().Debug("Single-package mode demo: baseDir=%s, filePath=%s, isWorkspace=%v, packagePath=%s",
			baseDir, entry.FilePath, config.Context.IsWorkspace(), entry.PackagePath)
	}

	// Read demo HTML file using injected filesystem
	demoPath := filepath.Join(baseDir, entry.FilePath)
	demoHTML, err := config.Context.FileSystem().ReadFile(demoPath)
	if err != nil {
		return "", fmt.Errorf("reading demo file %s: %w", entry.FilePath, err)
	}

	// Get import map as JSON (pre-computed during server initialization)
	var importMapJSON string
	if im := config.Context.ImportMap(); im != nil {
		if importMap, ok := im.(*importmappkg.ImportMap); ok {
			importMapJSON = importMap.ToJSON()
		}
	}

	// Prepare chrome data
	demoTitle := entry.Demo.Description
	if demoTitle == "" {
		demoTitle = filepath.Base(entry.FilePath)
	}

	// Get source URL (file location) and canonical URL (demo URL)
	var sourceURL string
	if entry.Demo.Source != nil {
		sourceURL = entry.Demo.Source.Href
	}

	// Generate knobs (default to all categories if not specified)
	var knobsHTML template.HTML
	enabledKnobs := queryParams["knobs"]
	if enabledKnobs == "" {
		// Default to all knob categories
		enabledKnobs = "attributes properties css-properties"
	}

	// Fetch manifest and generate knobs for ALL custom elements in demo
	var manifestBytes []byte
	var parsedManifest *M.Package      // Declare parsedManifest here
	var packages []PackageWithManifest // For workspace mode package-level tree

	if config.Context.IsWorkspace() {
		// Workspace mode: aggregate manifests from all packages
		middlewarePackages := config.Context.WorkspacePackages()
		var aggregatedManifest *M.Package
		for _, pkg := range middlewarePackages {
			if len(pkg.Manifest) == 0 {
				continue
			}

			var parsed M.Package
			if err := json.Unmarshal(pkg.Manifest, &parsed); err != nil {
				continue
			}

			// Track packages with their modules for tree organization
			if len(parsed.Modules) > 0 {
				packages = append(packages, PackageWithManifest{
					Name:    pkg.Name,
					Modules: parsed.Modules,
				})
			}

			if aggregatedManifest == nil {
				aggregatedManifest = &parsed
			} else {
				// Merge modules from this package into the aggregated manifest
				aggregatedManifest.Modules = append(aggregatedManifest.Modules, parsed.Modules...)
			}
		}

		if aggregatedManifest != nil {
			parsedManifest = aggregatedManifest
			manifestBytes, err = M.SerializeToBytes(aggregatedManifest)
			if err != nil {
				config.Context.Logger().Warning("Failed to serialize aggregated manifest: %v", err)
			}
		}
	} else {
		// Single-package mode: use context manifest
		manifestBytes, err = config.Context.Manifest()
		if err == nil && len(manifestBytes) > 0 {
			var pkg M.Package
			if err := json.Unmarshal(manifestBytes, &pkg); err == nil {
				parsedManifest = &pkg
			}
		}
	}

	// Generate knobs if we have a manifest
	if parsedManifest != nil {
		// Phase 5b: Discover all custom elements in demo (not just primary element)
		// This supports compositional components like accordion > accordion-header + accordion-panel
		allKnobGroups, err := GenerateKnobsForAllElements(parsedManifest, demoHTML, enabledKnobs)
		if err != nil {
			config.Context.Logger().Warning("Failed to generate knobs for demo elements: %v", err)
			_ = config.Context.BroadcastError(
				"Knobs Generation Error",
				fmt.Sprintf("Failed to generate knobs: %v", err),
				entry.TagName,
			)
		} else if len(allKnobGroups) > 0 {
			// Render all knob groups
			knobsHTML, err = RenderKnobsHTML(config.Templates, allKnobGroups)
			if err != nil {
				config.Context.Logger().Warning("Failed to render knobs HTML: %v", err)
				_ = config.Context.BroadcastError(
					"Knobs Render Error",
					fmt.Sprintf("Failed to render knobs HTML: %v", err),
					entry.TagName,
				)
			}
		}
	}

	// Extract state from request cookie for SSR
	state := GetStateFromRequest(r, config.Context.Logger())

	// Determine rendering mode: config default, overridden by query parameter
	renderingMode := config.Context.DemoRenderingMode()

	// Check for conflicting query parameters
	hasRendering := queryParams["rendering"] != ""
	hasShadow := queryParams["shadow"] != ""

	if hasRendering && hasShadow {
		// Both parameters present - broadcast error and prefer ?rendering
		config.Context.Logger().Warning("Both ?rendering and ?shadow query parameters present - using ?rendering value")
		_ = config.Context.BroadcastError(
			"Query Parameter Conflict",
			"Both ?rendering and ?shadow query parameters are present. Using ?rendering value. Remove ?shadow to avoid this warning.",
			entry.TagName,
		)
	}

	if renderingParam, ok := queryParams["rendering"]; ok {
		// Validate query parameter
		switch renderingParam {
		case "light", "shadow", "chromeless":
			renderingMode = renderingParam
		case "iframe":
			// Iframe not yet implemented - broadcast error and fallback to shadow
			config.Context.Logger().Warning("iframe rendering mode requested but not implemented, falling back to shadow")
			_ = config.Context.BroadcastError(
				"Rendering Mode Error",
				"iframe rendering mode is not yet implemented. Falling back to shadow mode.",
				entry.TagName,
			)
			renderingMode = "shadow"
		default:
			config.Context.Logger().Warning("invalid rendering mode '%s', using config default '%s'", renderingParam, renderingMode)
			_ = config.Context.BroadcastError(
				"Invalid Rendering Mode",
				fmt.Sprintf("Invalid rendering mode '%s'. Valid values are 'light', 'shadow', 'chromeless', or 'iframe'. Using config default '%s'.", renderingParam, renderingMode),
				entry.TagName,
			)
		}
	} else if queryParams["shadow"] == "true" {
		// Backward compatibility: ?shadow=true overrides to shadow mode
		// Only apply if ?rendering is not present (checked above)
		renderingMode = "shadow"
	}

	chromeData := ChromeData{
		TagName:        entry.TagName,
		DemoTitle:      demoTitle,
		DemoHTML:       template.HTML(demoHTML),
		Description:    template.HTML(entry.Demo.Description),
		ImportMap:      template.HTML(importMapJSON),
		EnabledKnobs:   enabledKnobs,
		KnobsHTML:      knobsHTML,
		RenderingMode:  renderingMode,
		SourceURL:      sourceURL,      // Link to source file
		CanonicalURL:   entry.Demo.URL, // Link to canonical demo
		PackageName:    packageName,
		NavigationHTML: navigationHTML,
		ManifestJSON:   template.JS(manifestBytes),
		Manifest:       parsedManifest,
		Packages:       packages, // Workspace packages with modules (for package-level tree)
		State:          state,    // Persisted UI state for SSR
	}

	// Render demo in the appropriate mode
	return renderDemo(config.Templates, config.Context, chromeData)
}

// levenshteinDistance computes the Levenshtein distance between two strings
func levenshteinDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	// Create a matrix to store distances
	matrix := make([][]int, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s2)+1)
	}

	// Initialize first row and column
	for i := 0; i <= len(s1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(s2); j++ {
		matrix[0][j] = j
	}

	// Fill in the matrix
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[len(s1)][len(s2)]
}

// URLSuggestion represents a suggested URL with its distance
type URLSuggestion struct {
	URL      string
	TagName  string
	DemoName string
	Distance int
}

// serve404Page renders a 404 page with URL suggestions
func serve404Page(w http.ResponseWriter, r *http.Request, config Config) {
	requestedPath := r.URL.Path

	// For non-HTML resources (CSS, JS, images, etc.), return simple 404 with correct MIME type
	// This prevents MIME type mismatches with X-Content-Type-Options: nosniff
	switch {
	case strings.HasSuffix(requestedPath, ".css"):
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("/* 404 Not Found */\n")); err != nil {
			config.Context.Logger().Error("Failed to write CSS 404 response: %v", err)
		}
		return
	case strings.HasSuffix(requestedPath, ".js"):
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("// 404 Not Found\n")); err != nil {
			config.Context.Logger().Error("Failed to write JS 404 response: %v", err)
		}
		return
	case strings.HasSuffix(requestedPath, ".json"):
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("{}\n")); err != nil {
			config.Context.Logger().Error("Failed to write JSON 404 response: %v", err)
		}
		return
	case strings.HasSuffix(requestedPath, ".svg"):
		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(http.StatusNotFound)
		return
	case strings.HasSuffix(requestedPath, ".png"), strings.HasSuffix(requestedPath, ".jpg"), strings.HasSuffix(requestedPath, ".jpeg"), strings.HasSuffix(requestedPath, ".gif"), strings.HasSuffix(requestedPath, ".webp"):
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get all available demo routes
	routes := config.Context.DemoRoutes()
	if routes == nil {
		// No routes available - render simple 404
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Compute Levenshtein distance for all routes
	var suggestions []URLSuggestion
	for route, entry := range routes {
		distance := levenshteinDistance(requestedPath, route)
		demoName := entry.Demo.Description
		if demoName == "" {
			demoName = filepath.Base(entry.FilePath)
		}
		suggestions = append(suggestions, URLSuggestion{
			URL:      route,
			TagName:  entry.TagName,
			DemoName: demoName,
			Distance: distance,
		})
	}

	// Sort by distance (closest first)
	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].Distance < suggestions[j].Distance
	})

	// Take top 5 suggestions
	maxSuggestions := 5
	if len(suggestions) > maxSuggestions {
		suggestions = suggestions[:maxSuggestions]
	}

	// Render 404 content with template
	var notFoundHTML strings.Builder
	notFoundData := map[string]any{
		"RequestedPath": requestedPath,
		"Suggestions":   suggestions,
	}
	if err := config.Templates.NotFoundTemplate.Execute(&notFoundHTML, notFoundData); err != nil {
		config.Context.Logger().Error("Failed to execute 404 template: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Get navigation data
	var navigationHTML template.HTML
	var packageName string

	if config.Context.IsWorkspace() {
		middlewarePackages := config.Context.WorkspacePackages()
		packages := make([]PackageContext, len(middlewarePackages))
		for i, pkg := range middlewarePackages {
			packages[i] = PackageContext{
				Name:     pkg.Name,
				Path:     pkg.Path,
				Manifest: pkg.Manifest,
			}
		}
		var navErr error
		navigationHTML, navErr = BuildWorkspaceNavigation(config.Templates, packages, r.URL.Path)
		if navErr != nil {
			config.Context.Logger().Error("Failed to build workspace navigation for %s: %v", r.URL.Path, navErr)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		packageName = "Workspace"
	} else {
		manifestBytes, err := config.Context.Manifest()
		if err == nil && len(manifestBytes) > 0 {
			pkgName := ""
			if pkg, err := config.Context.PackageJSON(); err == nil && pkg != nil {
				pkgName = pkg.Name
			}
			var navErr error
			navigationHTML, packageName, navErr = BuildSinglePackageNavigation(config.Templates, manifestBytes, pkgName, r.URL.Path)
			if navErr != nil {
				config.Context.Logger().Error("Failed to build navigation: %v", navErr)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
	}

	// Get import map
	var importMapJSON string
	if im := config.Context.ImportMap(); im != nil {
		if importMap, ok := im.(*importmappkg.ImportMap); ok {
			importMapJSON = importMap.ToJSON()
		}
	}

	// Extract state from request cookie for SSR
	state := GetStateFromRequest(r, config.Context.Logger())

	// Render 404 page with chrome
	chromeData := ChromeData{
		TagName:        "cem-serve",
		DemoTitle:      "404 Not Found",
		DemoHTML:       template.HTML(notFoundHTML.String()),
		ImportMap:      template.HTML(importMapJSON),
		PackageName:    packageName,
		NavigationHTML: navigationHTML,
		State:          state, // Persisted UI state for SSR
	}

	html, err := renderDemo(config.Templates, config.Context, chromeData)
	if err != nil {
		config.Context.Logger().Error("Failed to render 404 page: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte(html)); err != nil {
		config.Context.Logger().Error("Failed to write 404 response: %v", err)
	}
}

// servePackageStaticAsset serves static assets (CSS, JS, images) from demo package directories.
// Resolves URL paths like /elements/announcement/foo.css to the correct package directory (elements/rh-announcement/).
// Works in both workspace and single-package mode by using path mappings and routing table.
func servePackageStaticAsset(w http.ResponseWriter, r *http.Request, config Config) bool {
	requestPath := r.URL.Path
	ext := filepath.Ext(requestPath)
	var fullPath string

	// Strategy 1: Try PathResolver with URL rewrites
	resolver := config.Context.PathResolver()
	if resolver != nil {
		// For static assets, source ext == request ext
		resolvedPath := resolver.ResolveSourcePath(requestPath, ext, ext)
		if resolvedPath != "" {
			// Convert relative path to absolute path
			fullPath = filepath.Join(config.Context.WatchDir(), strings.TrimPrefix(resolvedPath, "/"))
			config.Context.Logger().Debug("servePackageStaticAsset: resolved via URL rewrites: %s -> %s", requestPath, fullPath)
		}
	}

	// Strategy 2: Fall back to routing table resolution
	if fullPath == "" {
		fullPath = resolveViaRoutingTable(requestPath, config)
	}

	if fullPath == "" {
		return false
	}

	// Try to read the file
	content, err := config.Context.FileSystem().ReadFile(fullPath)
	if err != nil {
		// File doesn't exist
		config.Context.Logger().Debug("servePackageStaticAsset: file not found: %v", err)
		return false
	}

	config.Context.Logger().Debug("servePackageStaticAsset: successfully serving %s (%d bytes)", fullPath, len(content))

	// Set correct MIME type based on file extension
	ext = filepath.Ext(fullPath)
	switch ext {
	case ".js", ".mjs", ".cjs":
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	case ".css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case ".html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case ".json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	case ".webp":
		w.Header().Set("Content-Type", "image/webp")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	w.Header().Set("Cache-Control", "no-cache")
	if _, err := w.Write(content); err != nil {
		config.Context.Logger().Error("Failed to write static asset response: %v", err)
	}

	return true
}

// resolveViaRoutingTable resolves static assets using the demo routing table.
// Returns the full file path or empty string if not found.
func resolveViaRoutingTable(requestPath string, config Config) string {
	// Get routing table
	routes := config.Context.DemoRoutes()
	if routes == nil {
		return ""
	}

	// Find the longest matching demo route prefix
	// e.g., /elements/announcement/rh-announcement-lightdom.css should match /elements/announcement/demo/
	var matchedEntry *DemoRouteEntry
	var matchedPrefix string

	for route, entry := range routes {
		// Extract the package/element prefix from the demo route
		// e.g., /elements/announcement/demo/ -> /elements/announcement/
		if !strings.Contains(route, "/demo/") {
			continue
		}

		prefix := route[:strings.Index(route, "/demo/")+1] // Include trailing slash: /elements/announcement/

		// Check if request path starts with this prefix
		if strings.HasPrefix(requestPath, prefix) && len(prefix) > len(matchedPrefix) {
			matchedEntry = entry
			matchedPrefix = prefix
		}
	}

	if matchedEntry == nil {
		config.Context.Logger().Debug("resolveViaRoutingTable: no matching route for %s", requestPath)
		return ""
	}

	// Extract the relative path after the prefix
	// e.g., /elements/announcement/rh-announcement-lightdom.css -> rh-announcement-lightdom.css
	relativePath := strings.TrimPrefix(requestPath, matchedPrefix)
	if relativePath == "" {
		return ""
	}

	// Build full path - different logic for workspace vs single-package mode
	var fullPath string
	if config.Context.IsWorkspace() && matchedEntry.PackagePath != "" {
		// Workspace mode: use PackagePath directly
		fullPath = filepath.Join(matchedEntry.PackagePath, relativePath)
		// Security: Reject path traversal attempts
		if rel, err := filepath.Rel(matchedEntry.PackagePath, fullPath); err != nil || strings.HasPrefix(rel, "..") {
			config.Context.Logger().Debug("resolveViaRoutingTable: path traversal rejected for %s", requestPath)
			return ""
		}
		config.Context.Logger().Debug("resolveViaRoutingTable (workspace): trying %s (package=%s, relative=%s)", fullPath, matchedEntry.PackagePath, relativePath)
	} else {
		// Single-package mode: extract package directory from FilePath
		// e.g., FilePath="elements/rh-announcement/demo/index.html" -> packageDir="elements/rh-announcement"
		packageDir := matchedEntry.FilePath
		if idx := strings.Index(packageDir, "/demo/"); idx >= 0 {
			packageDir = packageDir[:idx]
		} else {
			// Fallback: use directory of FilePath
			packageDir = filepath.Dir(matchedEntry.FilePath)
		}
		fullPath = filepath.Join(config.Context.WatchDir(), packageDir, relativePath)
		// Security: Reject path traversal attempts
		expectedBase := filepath.Join(config.Context.WatchDir(), packageDir)
		if rel, err := filepath.Rel(expectedBase, fullPath); err != nil || strings.HasPrefix(rel, "..") {
			config.Context.Logger().Debug("resolveViaRoutingTable: path traversal rejected for %s", requestPath)
			return ""
		}
		config.Context.Logger().Debug("resolveViaRoutingTable (single-package): trying %s (watchDir=%s, packageDir=%s, relative=%s)",
			fullPath, config.Context.WatchDir(), packageDir, relativePath)
	}

	return fullPath
}
