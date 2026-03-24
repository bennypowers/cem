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

		// Skip SSR artifacts and element JS (chrome bundle replaces them)
		base := filepath.Base(rel)
		if base == "ssr-bundle.js" || strings.HasSuffix(rel, ".qbc") {
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
