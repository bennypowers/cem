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
	"strings"

	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	V "bennypowers.dev/cem/internal/version"
	M "bennypowers.dev/cem/manifest"
)

// Config holds configuration for the internal routes middleware
type Config struct {
	// Context provides access to dev server state
	Context middleware.DevServerContext

	// LogsFunc returns the current logs (if logger supports it)
	LogsFunc func() []logger.LogEntry

	// WebSocketHandler handles WebSocket upgrade requests for live reload
	WebSocketHandler http.HandlerFunc
}

// New creates a middleware that handles internal CEM routes and demo routing
func New(config Config) middleware.Middleware {
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
			case strings.HasPrefix(r.URL.Path, "/__cem/"):
				serveInternalModules(w, r, config)
				return
			case r.URL.Path == "/custom-elements.json":
				serveManifest(w, r, config)
				return
			case r.URL.Path == "/__cem-logs":
				serveLogs(w, r, config)
				return
			case r.URL.Path == "/__cem-debug":
				serveDebugInfo(w, r, config)
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

			// Not a special route - pass to next handler (static files)
			next.ServeHTTP(w, r)
		})
	}
}

// serveInternalModules serves embedded JavaScript modules from embed.FS
func serveInternalModules(w http.ResponseWriter, r *http.Request, config Config) {
	// Strip /__cem/ prefix to get the file path within the embedded FS
	// Request: /__cem/foo.js -> templates/js/foo.js in embed.FS
	path := strings.TrimPrefix(r.URL.Path, "/__cem/")
	path = "templates/js/" + path

	data, err := InternalModules.ReadFile(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	if _, err := w.Write(data); err != nil {
		config.Context.Logger().Error("Failed to write JavaScript module response: %v", err)
	}
}

// serveManifest serves the custom elements manifest
func serveManifest(w http.ResponseWriter, r *http.Request, config Config) {
	manifest, err := config.Context.Manifest()
	if err != nil {
		config.Context.Logger().Error("Failed to get manifest: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if len(manifest) == 0 {
		http.Error(w, "Manifest not available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(manifest); err != nil {
		config.Context.Logger().Error("Failed to write manifest response: %v", err)
	}
}

// serveLogs serves the plain text logs for the debug console
func serveLogs(w http.ResponseWriter, r *http.Request, config Config) {
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
func serveDebugInfo(w http.ResponseWriter, r *http.Request, config Config) {
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
	var demos []map[string]interface{}
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

				demos = append(demos, map[string]interface{}{
					"tagName":      renderableDemo.CustomElementDeclaration.TagName,
					"description":  renderableDemo.Demo.Description,
					"canonicalURL": demoURL,
					"localRoute":   localRoute,
				})
			}
		}
	}

	debugInfo := map[string]interface{}{
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

	// Get import map as JSON
	var importMapJSON string
	if im := config.Context.ImportMap(); im != nil {
		if importMap, ok := im.(*importmappkg.ImportMap); ok {
			importMapJSON = importMap.ToJSON()
		}
	}

	var html string
	var err error

	if config.Context.IsWorkspace() {
		// Workspace mode: always render workspace listing
		middlewarePackages := config.Context.WorkspacePackages()
		packages := make([]PackageContext, len(middlewarePackages))
		for i, pkg := range middlewarePackages {
			packages[i] = PackageContext{
				Name:     pkg.Name,
				Path:     pkg.Path,
				Manifest: pkg.Manifest,
			}
		}
		html, err = RenderWorkspaceListing(packages, importMapJSON)
	} else {
		// Single-package mode: check if index.html exists
		watchDir := config.Context.WatchDir()
		if watchDir != "" {
			if _, err := os.Stat(filepath.Join(watchDir, "index.html")); err == nil {
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
		html, err = RenderElementListing(manifestBytes, importMapJSON)
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
	routesAny := config.Context.DemoRoutes()
	if routesAny == nil {
		// No routes available - not a demo route
		return false
	}

	routes, ok := routesAny.(map[string]*DemoRouteEntry)
	if !ok {
		config.Context.Logger().Error("Demo routes has unexpected type")
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

	// Render the demo
	html, err := renderDemoFromRoute(entry, queryParams, config)
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
func renderDemoFromRoute(entry *DemoRouteEntry, queryParams map[string]string, config Config) (string, error) {
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

	// Read demo HTML file
	demoPath := filepath.Join(baseDir, entry.FilePath)
	demoHTML, err := os.ReadFile(demoPath)
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

	chromeData := ChromeData{
		TagName:      entry.TagName,
		DemoTitle:    demoTitle,
		DemoHTML:     template.HTML(demoHTML),
		Description:  template.HTML(entry.Demo.Description),
		ImportMap:    template.HTML(importMapJSON),
		EnabledKnobs: queryParams["knobs"],
		ShadowMode:   queryParams["shadow"] == "true",
		SourceURL:    sourceURL,      // Link to source file
		CanonicalURL: entry.Demo.URL, // Link to canonical demo
	}

	// Render with chrome
	return renderDemoChrome(chromeData)
}
