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

package serve

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	M "bennypowers.dev/cem/manifest"
)

// DemoRouteEntry maps a local route to demo metadata
type DemoRouteEntry struct {
	LocalRoute string    // e.g., "/elements/accordion/demo/"
	TagName    string    // e.g., "rh-accordion"
	Demo       *M.Demo   // Demo metadata from manifest
	FilePath   string    // Relative file path from watch dir
}

// buildDemoRoutingTable creates a routing table from manifest and workspace config
func (s *Server) buildDemoRoutingTable(manifestBytes []byte) (map[string]*DemoRouteEntry, error) {
	if len(manifestBytes) == 0 {
		return nil, fmt.Errorf("no manifest available")
	}

	var pkg M.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return nil, fmt.Errorf("parsing manifest: %w", err)
	}

	// Get the source control root URL from config if available
	var sourceControlRootURL string
	if s.generateSession != nil {
		if workspace := s.generateSession.WorkspaceContext(); workspace != nil {
			if config, err := workspace.Config(); err == nil && config.SourceControlRootUrl != "" {
				sourceControlRootURL = config.SourceControlRootUrl
			}
		}
	}

	routes := make(map[string]*DemoRouteEntry)

	for _, renderableDemo := range pkg.RenderableDemos() {
		demoURL := renderableDemo.Demo.URL

		// Extract local route from canonical URL (strip origin)
		localRoute := demoURL
		if parsed, err := url.Parse(demoURL); err == nil && parsed.Path != "" {
			localRoute = parsed.Path
		}

		// Normalize relative paths (./foo -> /foo)
		if strings.HasPrefix(localRoute, "./") {
			localRoute = localRoute[1:] // Remove leading dot
		}
		if !strings.HasPrefix(localRoute, "/") {
			localRoute = "/" + localRoute
		}

		// Normalize: ensure trailing slash for directory-style URLs
		if localRoute != "/" && localRoute[len(localRoute)-1] != '/' && filepath.Ext(localRoute) == "" {
			localRoute += "/"
		}

		// Determine file path from Source.Href
		filePath := demoURL
		if renderableDemo.Demo.Source != nil && renderableDemo.Demo.Source.Href != "" {
			sourceHref := renderableDemo.Demo.Source.Href

			// If we have a sourceControlRootURL and the source href starts with it,
			// everything after the root URL is the file path
			if sourceControlRootURL != "" && strings.HasPrefix(sourceHref, sourceControlRootURL) {
				filePath = strings.TrimPrefix(sourceHref, sourceControlRootURL)
			} else if parsed, err := url.Parse(sourceHref); err == nil && parsed.Path != "" {
				// Fallback: extract the filename and combine with module directory
				moduleDir := filepath.Dir(renderableDemo.Module.Path)
				filename := filepath.Base(parsed.Path)
				filePath = filepath.Join(moduleDir, "demo", filename)
			}
		} else if strings.HasPrefix(demoURL, "./") || strings.HasPrefix(demoURL, "/") {
			// Relative path - use as-is (strip leading ./ if present)
			filePath = strings.TrimPrefix(demoURL, "./")
			filePath = strings.TrimPrefix(filePath, "/")
		}

		// Normalize filePath to ensure it's always relative (security: prevent directory traversal)
		// Strip leading "./" and "/" repeatedly to ensure relative path
		for strings.HasPrefix(filePath, "./") || strings.HasPrefix(filePath, "/") {
			filePath = strings.TrimPrefix(filePath, "./")
			filePath = strings.TrimPrefix(filePath, "/")
		}

		routes[localRoute] = &DemoRouteEntry{
			LocalRoute: localRoute,
			TagName:    renderableDemo.CustomElementDeclaration.TagName,
			Demo:       renderableDemo.Demo,
			FilePath:   filePath,
		}
	}

	return routes, nil
}

// renderDemoFromRoute renders a demo page with chrome using routing table entry
func (s *Server) renderDemoFromRoute(entry *DemoRouteEntry, queryParams map[string]string) (string, error) {
	s.mu.RLock()
	watchDir := s.watchDir
	s.mu.RUnlock()

	// Read demo HTML file
	demoPath := filepath.Join(watchDir, entry.FilePath)
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		return "", fmt.Errorf("reading demo file %s: %w", entry.FilePath, err)
	}

	// Generate import map
	var importMapJSON string
	if watchDir != "" {
		importMap, err := GenerateImportMap(watchDir, nil)
		if err != nil {
			s.logger.Warning("Failed to generate import map for demo: %v", err)
		} else if importMap != nil && len(importMap.Imports) > 0 {
			importMapBytes, err := json.MarshalIndent(importMap, "  ", "  ")
			if err != nil {
				s.logger.Warning("Failed to marshal import map: %v", err)
			} else {
				importMapJSON = string(importMapBytes)
			}
		}
	}

	// Prepare chrome data
	demoTitle := entry.Demo.Description
	if demoTitle == "" {
		demoTitle = filepath.Base(entry.FilePath)
	}

	chromeData := ChromeData{
		TagName:      entry.TagName,
		DemoTitle:    demoTitle,
		DemoHTML:     template.HTML(demoHTML),
		Description:  template.HTML(entry.Demo.Description),
		ImportMap:    template.HTML(importMapJSON),
		EnabledKnobs: queryParams["knobs"],
		ShadowMode:   queryParams["shadow"] == "true",
		SourceURL:    entry.Demo.URL,
	}

	// Render with chrome
	return renderDemoChrome(chromeData)
}

// serveDemoRoutes handles demo routing using the manifest-based routing table
func (s *Server) serveDemoRoutes(path string, queryParams map[string]string) (string, error) {
	s.mu.RLock()
	manifestBytes := s.manifest
	s.mu.RUnlock()

	// Build routing table from manifest
	routes, err := s.buildDemoRoutingTable(manifestBytes)
	if err != nil {
		return "", fmt.Errorf("building demo routing table: %w", err)
	}

	// Normalize requested path (ensure trailing slash for directory-style)
	normalizedPath := path
	if normalizedPath != "/" && normalizedPath[len(normalizedPath)-1] != '/' && filepath.Ext(normalizedPath) == "" {
		normalizedPath += "/"
	}

	// Look up route in table
	entry, found := routes[normalizedPath]
	if !found {
		return "", fmt.Errorf("demo not found for route: %s", path)
	}

	return s.renderDemoFromRoute(entry, queryParams)
}
