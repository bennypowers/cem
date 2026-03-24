/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package serve

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/routes"
)

// BuildConfig holds configuration for static site generation.
type BuildConfig struct {
	// OutputDir is the directory to write the static site to.
	OutputDir string
	// BasePath is the URL base path for deployment (e.g., "/docs/components/").
	BasePath string
}

// Build generates a static site from the dev server's demo pages.
// It spins up the full middleware chain, renders each page via internal
// HTTP requests, and writes the results to disk.
func (s *Server) Build(config BuildConfig) error {
	if config.OutputDir == "" {
		config.OutputDir = "dist"
	}

	s.logger.Info("Building static site to %s", config.OutputDir)

	// Enable static build mode so templates render chrome-bundle.js
	// instead of individual module scripts + hydration support
	s.staticBuild = true
	s.setupMiddleware()

	// Get demo routes from the pre-computed routing table
	demoRoutes := s.DemoRoutes()
	if len(demoRoutes) == 0 {
		return fmt.Errorf("no demo routes found")
	}

	// Create a test server backed by the real handler
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()

	// Write the index listing page
	if err := s.buildPage(ts, "/", config); err != nil {
		return fmt.Errorf("build index: %w", err)
	}

	// Write each demo page
	for route := range demoRoutes {
		if err := s.buildPage(ts, route, config); err != nil {
			s.logger.Warning("Failed to build %s: %v", route, err)
			continue
		}
	}

	// Copy chrome bundle
	if err := s.copyChrome(config); err != nil {
		return fmt.Errorf("copy chrome: %w", err)
	}

	// Copy static assets (CSS, images, utility JS)
	if err := s.copyStaticAssets(config); err != nil {
		return fmt.Errorf("copy assets: %w", err)
	}

	// Stub out websocket client (no live reload in static builds)
	if err := s.writeWebSocketStub(config); err != nil {
		return fmt.Errorf("write websocket stub: %w", err)
	}

	// Transform and copy user source files (TS→JS, CSS modules, etc.)
	if err := s.buildUserSources(ts, config); err != nil {
		return fmt.Errorf("build user sources: %w", err)
	}

	// Vendor node_modules dependencies referenced by import map
	if err := s.buildDependencies(ts, config); err != nil {
		return fmt.Errorf("build dependencies: %w", err)
	}

	// Write manifest as static JSON
	if manifest, err := s.Manifest(); err == nil && len(manifest) > 0 {
		outPath := filepath.Join(config.OutputDir, "custom-elements.json")
		if err := os.WriteFile(outPath, manifest, 0o644); err != nil {
			return fmt.Errorf("write manifest: %w", err)
		}
	}

	// Pre-render health API as static JSON
	if err := s.buildHealthJSON(ts, config); err != nil {
		s.logger.Warning("Failed to build health JSON: %v", err)
	}

	// Concatenate lightdom CSS into a single file
	if err := s.buildLightdomCSS(config); err != nil {
		return fmt.Errorf("build lightdom CSS: %w", err)
	}

	s.logger.Info("Built %d pages", len(demoRoutes)+1)
	return nil
}

// buildPage fetches a page from the internal server and writes it to disk.
// The template's {{if .StaticBuild}} block handles chrome-bundle.js injection.
func (s *Server) buildPage(ts *httptest.Server, route string, config BuildConfig) error {
	resp, err := ts.Client().Get(ts.URL + route)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d for %s", resp.StatusCode, route)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Determine output path
	outPath := filepath.Join(config.OutputDir, route)
	if strings.HasSuffix(route, "/") || filepath.Ext(route) == "" {
		outPath = filepath.Join(outPath, "index.html")
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}

	return os.WriteFile(outPath, body, 0o644)
}

// copyChrome copies the chrome bundle and sourcemap to the output directory.
func (s *Server) copyChrome(config BuildConfig) error {
	cemDir := filepath.Join(config.OutputDir, "__cem")

	for _, name := range []string{"chrome-bundle.js", "chrome-bundle.js.map"} {
		data, err := routes.TemplatesFS.ReadFile("templates/" + name)
		if err != nil {
			continue
		}
		outPath := filepath.Join(cemDir, name)
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(outPath, data, 0o644); err != nil {
			return err
		}
	}

	return nil
}

// copyStaticAssets copies CSS, images, JS, and lightdom CSS from the embedded FS.
func (s *Server) copyStaticAssets(config BuildConfig) error {
	cemDir := filepath.Join(config.OutputDir, "__cem")

	// Copy all assets from the embedded FS that the HTML references.
	// Copy assets using the same path mapping as the route handler.
	// /__cem/foo.js -> templates/js/foo.js
	// /__cem/foo.css -> templates/css/foo.css
	// /__cem/foo.svg -> templates/images/foo.svg
	// /__cem/elements/... -> templates/elements/...
	return fs.WalkDir(routes.InternalModules, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		rel := strings.TrimPrefix(path, "templates/")

		// Skip SSR artifacts, element JS (chrome bundle replaces them),
		// and websocket client (replaced by stub)
		base := filepath.Base(rel)
		if base == "ssr-bundle.js" || strings.HasSuffix(rel, ".qbc") {
			return nil
		}
		if base == "websocket-client.js" || base == "websocket-client.js.map" || base == "websocket-client.test.js" {
			return nil
		}
		if strings.HasPrefix(rel, "elements/") && strings.HasSuffix(rel, ".js") {
			return nil
		}
		if strings.HasPrefix(rel, "elements/") && strings.HasSuffix(rel, ".js.map") {
			return nil
		}

		data, err := routes.InternalModules.ReadFile(path)
		if err != nil {
			return nil
		}

		// Reverse the route handler's path mapping:
		// js/foo.js -> foo.js (served at /__cem/foo.js)
		// css/foo.css -> foo.css (served at /__cem/foo.css)
		// images/foo.svg -> foo.svg (served at /__cem/foo.svg)
		// elements/... stays as-is
		outRel := rel
		for _, prefix := range []string{"js/", "css/", "images/"} {
			if strings.HasPrefix(rel, prefix) {
				outRel = strings.TrimPrefix(rel, prefix)
				break
			}
		}

		outPath := filepath.Join(cemDir, outRel)
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}
		return os.WriteFile(outPath, data, 0o644)
	})
}

// buildLightdomCSS concatenates all lightdom CSS files into one sheet.
func (s *Server) buildLightdomCSS(config BuildConfig) error {
	var combined strings.Builder

	_ = fs.WalkDir(routes.InternalModules, "templates/elements", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if !strings.Contains(filepath.Base(path), "lightdom") || !strings.HasSuffix(path, ".css") {
			return nil
		}
		data, err := routes.InternalModules.ReadFile(path)
		if err != nil {
			return nil
		}
		combined.Write(data)
		combined.WriteByte('\n')
		return nil
	})

	if combined.Len() == 0 {
		return nil
	}

	outPath := filepath.Join(config.OutputDir, "__cem", "lightdom.css")
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(combined.String()), 0o644)
}

// buildHealthJSON fetches health data from the internal server and writes it as static JSON.
func (s *Server) buildHealthJSON(ts *httptest.Server, config BuildConfig) error {
	resp, err := ts.Client().Get(ts.URL + "/__cem/api/health")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health API returned %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	outPath := filepath.Join(config.OutputDir, "__cem", "api", "health")
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outPath, body, 0o644)
}

// writeWebSocketStub writes a no-op websocket client module for static builds.
// The chrome component dynamically imports this, so it must exist and export
// the expected CEMReloadClient class.
func (s *Server) writeWebSocketStub(config BuildConfig) error {
	stub := `// Static build stub - no live reload
export class CEMReloadClient {
  init() {}
  retry() {}
  destroy() {}
}
`
	outPath := filepath.Join(config.OutputDir, "__cem", "websocket-client.js")
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(stub), 0o644)
}

// buildUserSources walks the user's source directory and fetches each file
// through the test server, which applies TS→JS and CSS transforms.
func (s *Server) buildUserSources(ts *httptest.Server, config BuildConfig) error {
	watchDir := s.WatchDir()
	if watchDir == "" {
		return nil
	}

	// Compute the URL prefix: the relative path from CWD to watchDir
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	relDir, err := filepath.Rel(cwd, watchDir)
	if err != nil {
		return err
	}

	count := 0
	err = filepath.WalkDir(watchDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		// Skip node_modules (handled by buildDependencies)
		rel, _ := filepath.Rel(watchDir, path)
		if strings.HasPrefix(rel, "node_modules") {
			return nil
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".ts", ".js", ".css", ".html", ".json", ".svg", ".png", ".jpg", ".gif", ".woff", ".woff2":
			// Dev server serves files relative to watchDir
			urlPath := "/" + rel
			if ext == ".ts" {
				urlPath = strings.TrimSuffix(urlPath, ".ts") + ".js"
			}

			body, fetchErr := fetchURL(ts, urlPath)
			if fetchErr != nil {
				// Non-fatal: file may not be directly servable
				return nil
			}

			// For JS/TS files, rewrite CSS import attribute query params to .js extension
			// The dev server uses ?__cem-import-attrs[type]=css but static servers
			// don't handle query params, so we write CSS-as-JS wrappers as .css.js files
			if ext == ".ts" || ext == ".js" {
				body = []byte(strings.ReplaceAll(string(body),
					"?__cem-import-attrs[type]=css", ".js"))
			}

			// Write to the CWD-relative path so import map URLs resolve
			outPath := filepath.Join(config.OutputDir, relDir, rel)
			if ext == ".ts" {
				outPath = strings.TrimSuffix(outPath, ".ts") + ".js"
			}
			if mkErr := os.MkdirAll(filepath.Dir(outPath), 0o755); mkErr != nil {
				return mkErr
			}
			if writeErr := os.WriteFile(outPath, body, 0o644); writeErr != nil {
				return writeErr
			}
			count++

			// For CSS files, also create a .css.js wrapper (CSS module for JS import)
			if ext == ".css" {
				cssModuleURL := urlPath + "?__cem-import-attrs[type]=css"
				jsBody, jsErr := fetchURL(ts, cssModuleURL)
				if jsErr == nil {
					jsPath := outPath + ".js"
					if writeErr := os.WriteFile(jsPath, jsBody, 0o644); writeErr != nil {
						return writeErr
					}
					count++
				}
			}

			return nil
		default:
			return nil
		}
	})

	if err != nil {
		return err
	}

	s.logger.Info("Transformed %d user source files", count)
	return nil
}

// buildDependencies copies node_modules files referenced by the import map.
// Fetches through the test server so transforms are applied.
func (s *Server) buildDependencies(ts *httptest.Server, config BuildConfig) error {
	im := s.ImportMap()
	if im == nil {
		return nil
	}
	importMap, ok := im.(*importmap.ImportMap)
	if !ok || importMap == nil {
		return nil
	}

	// Collect unique target paths from import map
	seen := make(map[string]bool)
	var paths []string
	for _, target := range importMap.Imports {
		if !strings.HasPrefix(target, "/node_modules/") {
			continue
		}
		if seen[target] {
			continue
		}
		seen[target] = true
		paths = append(paths, target)
	}
	// Also check scopes
	for _, scope := range importMap.Scopes {
		for _, target := range scope {
			if !strings.HasPrefix(target, "/node_modules/") {
				continue
			}
			if seen[target] {
				continue
			}
			seen[target] = true
			paths = append(paths, target)
		}
	}

	count := 0
	for _, urlPath := range paths {
		body, err := fetchURL(ts, urlPath)
		if err != nil {
			s.logger.Debug("Skip dependency %s: %v", urlPath, err)
			continue
		}

		outPath := filepath.Join(config.OutputDir, urlPath)
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(outPath, body, 0o644); err != nil {
			return err
		}
		count++

		// Also copy source map if it exists
		if strings.HasSuffix(urlPath, ".js") {
			mapURL := urlPath + ".map"
			mapBody, mapErr := fetchURL(ts, mapURL)
			if mapErr == nil {
				mapPath := outPath + ".map"
				_ = os.WriteFile(mapPath, mapBody, 0o644)
			}
		}
	}

	s.logger.Info("Vendored %d dependencies", count)
	return nil
}

// fetchURL fetches a URL path from the test server, returning the body bytes.
func fetchURL(ts *httptest.Server, urlPath string) ([]byte, error) {
	resp, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
