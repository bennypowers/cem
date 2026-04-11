/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package serve

import (
	"errors"
	"fmt"
	"html"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/routes"
)

// BuildConfig holds configuration for static site generation.
type BuildConfig struct {
	// OutputDir is the directory to write the static site to.
	OutputDir string
	// BasePath is the URL base path for deployment (e.g., "/docs/components/").
	BasePath string
	// ImportMode controls dependency resolution: vendor (default), esm, jspm, unpkg.
	ImportMode string
}

// siteRoot returns the output directory including the base path.
// All files are written under this directory so URLs resolve correctly.
func (c BuildConfig) siteRoot() string {
	if c.BasePath == "" {
		return c.OutputDir
	}
	return filepath.Join(c.OutputDir, c.BasePath)
}

// Build generates a static site from the dev server's demo pages.
// It renders each page in-process through the middleware chain
// and writes the results to disk.
func (s *Server) Build(config BuildConfig) error {
	if config.OutputDir == "" {
		config.OutputDir = "dist"
	}

	// Normalize base path: ensure leading slash, no trailing slash
	if config.BasePath != "" {
		config.BasePath = "/" + strings.Trim(config.BasePath, "/")
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

	// Render pages in-process through the middleware chain.
	// No HTTP server is started -- requests are handled directly via
	// handler.ServeHTTP, avoiding TCP connections and the EOF errors
	// that httptest.Server causes on resource-constrained CI runners.
	handler := s.Handler()

	// Write the index listing page
	if err := s.buildPage(handler, "/", config); err != nil {
		return fmt.Errorf("build index: %w", err)
	}

	// Write each demo page
	var pageErrors error
	for route := range demoRoutes {
		if err := s.buildPage(handler, route, config); err != nil {
			s.logger.Error("Failed to build %s: %v", route, err)
			pageErrors = errors.Join(pageErrors, fmt.Errorf("%s: %w", route, err))
			continue
		}
	}
	if pageErrors != nil {
		return fmt.Errorf("failed to build pages: %w", pageErrors)
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
	if err := s.buildUserSources(handler, config, demoRoutes); err != nil {
		return fmt.Errorf("build user sources: %w", err)
	}

	// Handle dependencies based on import mode
	switch config.ImportMode {
	case "esm", "jspm", "unpkg":
		// CDN mode: rewrite import map in all HTML files to use CDN URLs
		if err := s.rewriteImportMapToCDN(config); err != nil {
			return fmt.Errorf("rewrite import map to CDN: %w", err)
		}
	default:
		// Vendor mode (default): copy node_modules to output
		if err := s.buildDependencies(handler, config); err != nil {
			return fmt.Errorf("build dependencies: %w", err)
		}
	}

	// Write manifest as static JSON
	if manifest, err := s.Manifest(); err == nil && len(manifest) > 0 {
		outPath := filepath.Join(config.siteRoot(), "custom-elements.json")
		if err := os.WriteFile(outPath, manifest, 0o644); err != nil {
			return fmt.Errorf("write manifest: %w", err)
		}
	}

	// Pre-render health API as static JSON
	if err := s.buildHealthJSON(handler, config); err != nil {
		s.logger.Warning("Failed to build health JSON: %v", err)
	}

	// Concatenate lightdom CSS into a single file
	if err := s.buildLightdomCSS(config); err != nil {
		return fmt.Errorf("build lightdom CSS: %w", err)
	}

	// Generate sitemap
	if err := s.buildSitemap(demoRoutes, config); err != nil {
		s.logger.Warning("Failed to generate sitemap: %v", err)
	}

	s.logger.Info("Built %d pages", len(demoRoutes)+1)
	return nil
}

// buildPage renders a page through the handler and writes it to disk.
// The template's {{if .StaticBuild}} block handles chrome-bundle.js injection.
func (s *Server) buildPage(handler http.Handler, route string, config BuildConfig) error {
	body, err := renderURL(handler, route)
	if err != nil {
		return err
	}

	// Rewrite absolute URLs with base path prefix
	if config.BasePath != "" {
		body = rewriteBasePath(body, config.BasePath)
	}

	// Determine output path
	outPath := filepath.Join(config.siteRoot(), route)
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
	cemDir := filepath.Join(config.siteRoot(), "__cem")

	for _, name := range []string{"chrome-bundle.js", "chrome-bundle.js.map"} {
		data, err := routes.TemplatesFS.ReadFile("templates/" + name)
		if err != nil {
			s.logger.Warning("Chrome bundle missing: %s: %v", name, err)
			continue
		}
		// Rewrite absolute paths in JS bundle for base path deployment
		if config.BasePath != "" && strings.HasSuffix(name, ".js") {
			data = rewriteAssetPaths(data, config.BasePath)
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
	cemDir := filepath.Join(config.siteRoot(), "__cem")

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
			return fmt.Errorf("reading internal module %s: %w", path, err)
		}

		// Reverse the route handler's path mapping:
		// js/foo.js -> foo.js (served at /__cem/foo.js)
		// css/foo.css -> foo.css (served at /__cem/foo.css)
		// images/foo.svg -> foo.svg (served at /__cem/foo.svg)
		// elements/... stays as-is
		outRel := rel
		for _, prefix := range []string{"js/", "css/", "images/"} {
			if after, found := strings.CutPrefix(rel, prefix); found {
				outRel = after
				break
			}
		}

		// Rewrite absolute paths in JS/CSS assets for base path deployment
		if config.BasePath != "" {
			ext := filepath.Ext(outRel)
			if ext == ".js" || ext == ".css" {
				data = rewriteAssetPaths(data, config.BasePath)
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

	if err := fs.WalkDir(routes.InternalModules, "templates/elements", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if !strings.Contains(filepath.Base(path), "lightdom") || !strings.HasSuffix(path, ".css") {
			return nil
		}
		data, readErr := routes.InternalModules.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		combined.Write(data)
		combined.WriteByte('\n')
		return nil
	}); err != nil {
		return err
	}

	// Always write the file so the <link> in demo-chrome.html resolves
	outPath := filepath.Join(config.siteRoot(), "__cem", "lightdom.css")
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(combined.String()), 0o644)
}

// buildHealthJSON renders the health API endpoint and writes it as static JSON.
func (s *Server) buildHealthJSON(handler http.Handler, config BuildConfig) error {
	body, err := renderURL(handler, "/__cem/api/health")
	if err != nil {
		return err
	}

	outPath := filepath.Join(config.siteRoot(), "__cem", "api", "health")
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
	outPath := filepath.Join(config.siteRoot(), "__cem", "websocket-client.js")
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(stub), 0o644)
}

// buildUserSources walks the user's source directory and renders each file
// through the handler, which applies TS→JS and CSS transforms.
func (s *Server) buildUserSources(handler http.Handler, config BuildConfig, demoRoutes map[string]*middleware.DemoRouteEntry) error {
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

	// Resolve output directory to an absolute path for skipping during walk
	absOutputDir, err := filepath.Abs(config.OutputDir)
	if err != nil {
		return err
	}

	// Build set of demo source file paths to skip (already rendered with chrome by buildPage)
	demoSourceFiles := make(map[string]bool, len(demoRoutes))
	for _, entry := range demoRoutes {
		demoSourceFiles[entry.FilePath] = true
	}

	count := 0
	err = filepath.WalkDir(watchDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip node_modules (handled by buildDependencies)
		rel, _ := filepath.Rel(watchDir, path)
		if strings.HasPrefix(rel, "node_modules") {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		// Skip the output directory to avoid infinite recursion
		absPath, _ := filepath.Abs(path)
		if d.IsDir() && absPath == absOutputDir {
			return fs.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		// Skip demo source files (already rendered with chrome by buildPage)
		if demoSourceFiles[rel] {
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

			body, fetchErr := renderURL(handler, urlPath)
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
			outPath := filepath.Join(config.siteRoot(), relDir, rel)
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
				jsBody, jsErr := renderURL(handler, cssModuleURL)
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
// Renders through the handler so transforms are applied.
func (s *Server) buildDependencies(handler http.Handler, config BuildConfig) error {
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
		body, err := renderURL(handler, urlPath)
		if err != nil {
			s.logger.Debug("Skip dependency %s: %v", urlPath, err)
			continue
		}

		outPath := filepath.Join(config.siteRoot(), urlPath)
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
			mapBody, mapErr := renderURL(handler, mapURL)
			if mapErr == nil {
				mapPath := outPath + ".map"
				_ = os.WriteFile(mapPath, mapBody, 0o644)
			}
		}
	}

	s.logger.Info("Vendored %d dependencies", count)
	return nil
}

// buildSitemap generates a sitemap.xml listing all built pages.
func (s *Server) buildSitemap(demoRoutes map[string]*middleware.DemoRouteEntry, config BuildConfig) error {
	// Collect and sort routes for deterministic output
	urls := make([]string, 0, len(demoRoutes)+1)
	urls = append(urls, "/")
	for route := range demoRoutes {
		urls = append(urls, route)
	}
	sort.Strings(urls)

	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	sb.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	for _, u := range urls {
		loc := config.BasePath + u
		sb.WriteString("  <url><loc>" + html.EscapeString(loc) + "</loc></url>\n")
	}
	sb.WriteString("</urlset>\n")

	outPath := filepath.Join(config.siteRoot(), "sitemap.xml")
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(sb.String()), 0o644)
}

// rewriteImportMapToCDN rewrites import map entries in all HTML files
// to use CDN URLs instead of local /node_modules/ paths.
func (s *Server) rewriteImportMapToCDN(config BuildConfig) error {
	root := config.siteRoot()
	count := 0

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".html") {
			return err
		}

		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}

		original := string(data)
		rewritten := rewriteNodeModulesToCDN(original, config.ImportMode)
		if rewritten == original {
			return nil
		}

		count++
		return os.WriteFile(path, []byte(rewritten), 0o644)
	})

	if err != nil {
		return err
	}

	s.logger.Info("Rewrote import maps to %s CDN in %d files", config.ImportMode, count)
	return nil
}

// rewriteNodeModulesToCDN replaces /node_modules/ paths with CDN URLs.
// Handles both import map JSON and HTML attribute references.
func rewriteNodeModulesToCDN(html, mode string) string {
	// Find and replace /node_modules/PKG/path patterns
	// Package names: bare (lit) or scoped (@lit/reactive-element)
	var result strings.Builder
	result.Grow(len(html))

	prefix := "/node_modules/"
	for {
		idx := strings.Index(html, prefix)
		if idx == -1 {
			result.WriteString(html)
			break
		}

		// Write everything before /node_modules/ and skip the prefix
		result.WriteString(html[:idx])
		html = html[idx+len(prefix):]

		// Extract package name and subpath
		pkg, subpath := parseNodeModulePath(html)
		if pkg == "" {
			result.WriteString(prefix)
			continue
		}

		// Advance past the consumed path
		consumed := pkg
		if subpath != "" {
			consumed += "/" + subpath
		}
		html = html[len(consumed):]

		// Generate CDN URL
		cdnURL := cdnURL(mode, pkg, subpath)
		result.WriteString(cdnURL)
	}

	return result.String()
}

// parseNodeModulePath extracts package name and subpath from a path
// starting after /node_modules/. Returns (pkg, subpath).
func parseNodeModulePath(s string) (string, string) {
	// Find the end of the path (quote, whitespace, or HTML delimiter)
	end := strings.IndexFunc(s, func(r rune) bool {
		return r == '"' || r == '\'' || r == ' ' || r == '>' || r == '\n' || r == '\r' || r == '\t'
	})
	if end == -1 {
		return "", ""
	}
	fullPath := s[:end]

	var pkg, subpath string
	if strings.HasPrefix(fullPath, "@") {
		// Scoped package: @scope/name/subpath
		parts := strings.SplitN(fullPath, "/", 3)
		if len(parts) < 2 {
			return "", ""
		}
		pkg = parts[0] + "/" + parts[1]
		if len(parts) == 3 {
			subpath = parts[2]
		}
	} else {
		// Bare package: name/subpath
		parts := strings.SplitN(fullPath, "/", 2)
		pkg = parts[0]
		if len(parts) == 2 {
			subpath = parts[1]
		}
	}

	return pkg, subpath
}

// cdnURL constructs a CDN URL for a package.
func cdnURL(mode, pkg, subpath string) string {
	switch mode {
	case "esm":
		if subpath != "" {
			return "https://esm.sh/" + pkg + "/" + subpath
		}
		return "https://esm.sh/" + pkg
	case "jspm":
		if subpath != "" {
			return "https://ga.jspm.io/npm:" + pkg + "/" + subpath
		}
		return "https://ga.jspm.io/npm:" + pkg
	case "unpkg":
		if subpath != "" {
			return "https://unpkg.com/" + pkg + "/" + subpath
		}
		return "https://unpkg.com/" + pkg
	default:
		// Shouldn't happen, return original
		return "/node_modules/" + pkg
	}
}

// renderURL renders a URL path through the handler in-process, returning the body bytes.
// No HTTP server or TCP connection is involved.
func renderURL(handler http.Handler, urlPath string) ([]byte, error) {
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request for %s: %w", urlPath, err)
	}
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", w.Code)
	}

	return w.Body.Bytes(), nil
}

// rewriteAssetPaths rewrites absolute paths in JS and CSS assets for base path deployment.
// Handles patterns like:
//   - JS: import("/__cem/..."), fetch("/custom-elements.json"), new URL("/__cem/...", ...)
//   - CSS: @import url('/__cem/...')
func rewriteAssetPaths(data []byte, basePath string) []byte {
	s := string(data)

	// Rewrite JS patterns: import("/...", fetch("/..., new URL("/...
	for _, prefix := range []string{
		`import("`, `import('`,
		`fetch("`, `fetch('`,
		`URL("`, `URL('`,
		`url('`, `url("`,
	} {
		s = rewriteAttrPaths(s, prefix, basePath)
	}

	return []byte(s)
}

// rewriteBasePath prefixes absolute URLs in HTML with the base path.
// Uses a simple but correct approach: rewrite all absolute URL patterns
// that appear in HTML attributes and import map JSON.
func rewriteBasePath(body []byte, basePath string) []byte {
	s := string(body)

	// Rewrite HTML attributes with absolute paths: href="/...", src="/..."
	for _, attr := range []string{`href="`, `src="`} {
		s = rewriteAttrPaths(s, attr, basePath)
	}

	// Rewrite import map JSON values: ": "/..."
	s = rewriteJSONPaths(s, basePath)

	// Rewrite dynamic import() paths in inline scripts
	for _, quote := range []string{`'`, `"`} {
		s = rewriteAttrPaths(s, "import("+quote, basePath)
	}

	return []byte(s)
}

// rewriteAttrPaths rewrites absolute paths in HTML attributes.
// e.g., href="/foo" → href="/base/foo"
func rewriteAttrPaths(s, attrPrefix, basePath string) string {
	var result strings.Builder
	result.Grow(len(s))

	for {
		idx := strings.Index(s, attrPrefix)
		if idx == -1 {
			result.WriteString(s)
			break
		}

		// Write everything up to and including the attr prefix
		result.WriteString(s[:idx+len(attrPrefix)])
		s = s[idx+len(attrPrefix):]

		// Check if the value starts with "/" (absolute path) but not "//" (protocol-relative)
		if len(s) > 0 && s[0] == '/' && (len(s) < 2 || s[1] != '/') {
			// Don't double-prefix if already has the base path
			if !strings.HasPrefix(s, basePath+"/") && !strings.HasPrefix(s, basePath+`"`) {
				result.WriteString(basePath)
			}
		}
	}

	return result.String()
}

// rewriteJSONPaths rewrites absolute paths in import map JSON,
// both values (": "/...") and scope keys ("/node_modules/...": {).
func rewriteJSONPaths(s, basePath string) string {
	// Rewrite JSON values: ": "/..."
	for _, pattern := range []string{`": "/`, `": '/`} {
		quote := pattern[len(pattern)-2 : len(pattern)-1] // extract the quote char
		var result strings.Builder
		result.Grow(len(s))

		for {
			idx := strings.Index(s, pattern)
			if idx == -1 {
				result.WriteString(s)
				break
			}

			// Write up to and including ": <quote>
			prefix := `": ` + quote
			result.WriteString(s[:idx+len(prefix)])
			s = s[idx+len(prefix):]

			// s now starts with "/" - prefix with base path if not already present
			if !strings.HasPrefix(s, basePath+"/") {
				result.WriteString(basePath)
			}
		}

		s = result.String()
	}

	// Rewrite scope keys: "/node_modules/...": { and similar
	// Pattern: start-of-key quote, slash, ...path..., quote, colon
	// We look for `"/` at the start of a JSON key in scopes
	s = rewriteJSONScopeKeys(s, basePath)

	return s
}

// rewriteJSONScopeKeys rewrites absolute paths used as JSON object keys.
// e.g., "/node_modules/lit/": { → "/base/node_modules/lit/": {
// Only matches keys that look like scope entries: "/path/": {
func rewriteJSONScopeKeys(s, basePath string) string {
	var result strings.Builder
	result.Grow(len(s))

	// Scope keys appear as: "/path/": {
	// We match `"/` but exclude `"//` (protocol-relative URLs)
	// and `": "/` (JSON values, handled by rewriteJSONPaths)
	pattern := `"/`
	for {
		idx := strings.Index(s, pattern)
		if idx == -1 {
			result.WriteString(s)
			break
		}

		// Distinguish JSON keys from values: values are preceded by `: `
		prefix := s[:idx]
		isValue := strings.HasSuffix(prefix, `: `)

		// Write everything up to and including the opening quote
		result.WriteString(s[:idx+1]) // include the "
		s = s[idx+1:]                 // s starts with /

		if isValue {
			continue
		}

		// Skip protocol-relative URLs: //...
		if len(s) > 1 && s[1] == '/' {
			continue
		}

		// Verify this is actually a scope key: "/path/": {
		// Look ahead for closing quote, colon, optional whitespace, opening brace
		if !looksLikeScopeKey(s) {
			continue
		}

		// Prefix with base path if not already present
		if !strings.HasPrefix(s, basePath+"/") && !strings.HasPrefix(s, basePath+`"`) {
			result.WriteString(basePath)
		}
	}

	return result.String()
}

// looksLikeScopeKey checks whether the text starting at s (after the opening
// quote and slash) matches the pattern for a JSON scope key: /path/": {
// i.e., there is a closing `"` followed by `:` and `{` with optional whitespace.
func looksLikeScopeKey(s string) bool {
	closeQuote := strings.IndexByte(s, '"')
	if closeQuote == -1 {
		return false
	}
	rest := strings.TrimLeft(s[closeQuote+1:], " \t\n\r")
	if len(rest) == 0 || rest[0] != ':' {
		return false
	}
	rest = strings.TrimLeft(rest[1:], " \t\n\r")
	return len(rest) > 0 && rest[0] == '{'
}
