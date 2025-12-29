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
	"net/http"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/cors"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/inject"
	"bennypowers.dev/cem/serve/middleware/requestlogger"
	"bennypowers.dev/cem/serve/middleware/routes"
	"bennypowers.dev/cem/serve/middleware/shadowroot"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// Handler returns the HTTP handler
func (s *Server) Handler() http.Handler {
	return s.handler
}

// errorBroadcaster adapts Server.BroadcastError to the middleware interface
type errorBroadcaster struct {
	*Server
}

func (e errorBroadcaster) BroadcastError(title, message, filename string) {
	_ = e.Server.BroadcastError(title, message, filename) // Ignore error
}

// getLogs returns logs if the logger supports it
func (s *Server) getLogs() []logger.LogEntry {
	if logGetter, ok := s.logger.(interface{ Logs() []logger.LogEntry }); ok {
		return logGetter.Logs()
	}
	return nil
}

// setupMiddleware configures the middleware pipeline
func (s *Server) setupMiddleware() {
	// Get WebSocket handler if reload is enabled
	var wsHandler http.HandlerFunc
	if s.wsManager != nil {
		wsHandler = s.wsManager.HandleConnection
	}

	// Apply middleware using Chain helper
	// Initialize template registry for HTML rendering
	s.templates = routes.NewTemplateRegistry(s)

	// Middlewares are applied in reverse order (last to first in the chain)
	// Terminal handler: static files
	s.handler = middleware.Chain(
		http.HandlerFunc(s.serveStaticFiles), // Static file server (terminal handler)
		shadowroot.New(
			s.logger,
			errorBroadcaster{s},
			routes.TemplatesFS,
			func(elementName string, data any) (string, error) {
				html, err := routes.RenderElementShadowRoot(s.templates, elementName, data)
				return string(html), err
			},
		), // Shadow root injection (last - processes final HTML)
		inject.New(s.config.Reload, "/__cem/websocket-client.js"), // WebSocket injection
		importmappkg.New(importmappkg.MiddlewareConfig{ // Import map injection
			Context: s,
		}),
		transform.NewCSS(transform.CSSConfig{ // CSS transform
			WatchDirFunc:     s.WatchDir,
			Logger:           s.logger,
			ErrorBroadcaster: errorBroadcaster{s},
			ConfigFile:       s.config.ConfigFile,
			Enabled:          s.config.Transforms.CSS.Enabled,
			Include:          s.config.Transforms.CSS.Include,
			Exclude:          s.config.Transforms.CSS.Exclude,
			FS:               s.fs,
			PathResolver:     s.pathResolver,
		}),
		transform.NewTypeScript(transform.TypeScriptConfig{ // TypeScript transform
			WatchDirFunc:     s.WatchDir,
			TsconfigRawFunc:  s.TsconfigRaw,
			Cache:            s.transformCache,
			Pool:             s.transformPool,
			Logger:           s.logger,
			ErrorBroadcaster: errorBroadcaster{s},
			Target:           string(s.config.Transforms.TypeScript.Target),
			Enabled:          s.config.Transforms.TypeScript.Enabled,
			FS:               s.fs,
			PathResolver:     s.pathResolver,
		}),
		routes.New(routes.Config{ // Internal CEM routes (includes WebSocket, demos, listings)
			Context:          s,
			LogsFunc:         s.getLogs,
			WebSocketHandler: wsHandler,
			Templates:        s.templates,
		}),
		cors.New(),                  // CORS headers
		requestlogger.New(s.logger), // HTTP request logging
	)
}

// serveStaticFiles serves static files from the watch directory using injected filesystem
func (s *Server) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	watchDir := s.watchDir
	fs := s.fs
	s.mu.RUnlock()

	if watchDir == "" {
		http.NotFound(w, r)
		return
	}

	// Clean the URL path
	requestPath := filepath.Clean(r.URL.Path)
	if requestPath == "." {
		requestPath = "/"
	}

	// Build full file path
	fullPath := filepath.Join(watchDir, strings.TrimPrefix(requestPath, "/"))

	// Reject path traversal attempts
	if rel, err := filepath.Rel(watchDir, fullPath); err != nil || strings.HasPrefix(rel, "..") {
		http.NotFound(w, r)
		return
	}

	// Try to read the file using injected filesystem
	content, err := fs.ReadFile(fullPath)
	if err != nil {
		// Check if it's a directory - if so, try index.html
		if stat, statErr := fs.Stat(fullPath); statErr == nil && stat.IsDir() {
			indexPath := filepath.Join(fullPath, "index.html")
			if indexContent, indexErr := fs.ReadFile(indexPath); indexErr == nil {
				content = indexContent
				fullPath = indexPath
				err = nil
			}
		}

		if err != nil {
			http.NotFound(w, r)
			return
		}
	}

	// Set correct MIME type
	ext := filepath.Ext(fullPath)
	switch ext {
	case ".js", ".mjs", ".cjs":
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	case ".css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case ".html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case ".map":
	case ".json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case ".woff":
		w.Header().Set("Content-Type", "font/woff")
	case ".woff2":
		w.Header().Set("Content-Type", "font/woff2")
	}

	// Write the content
	if _, err := w.Write(content); err != nil {
		s.logger.Error("Failed to write static file response: %v", err)
	}
}
