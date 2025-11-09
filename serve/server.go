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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	G "bennypowers.dev/cem/generate"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// Server represents the development server
type Server struct {
	port            int
	config          Config
	server          *http.Server
	listener        net.Listener
	handler         http.Handler
	logger          Logger
	wsManager       WebSocketManager
	watcher         FileWatcher
	watchDir        string
	manifest        []byte
	sourceFiles     map[string]bool // Set of source files to watch
	running         bool
	mu              sync.RWMutex
	generateSession *G.GenerateSession
}

// NewServer creates a new server with the given port
func NewServer(port int) (*Server, error) {
	config := Config{
		Port:   port,
		Reload: true,
	}
	return NewServerWithConfig(config)
}

// NewServerWithConfig creates a new server with the given config
func NewServerWithConfig(config Config) (*Server, error) {
	s := &Server{
		port:   config.Port,
		config: config,
		logger: &defaultLogger{},
	}

	// Create WebSocket manager if reload is enabled
	if config.Reload {
		s.wsManager = newWebSocketManager()
		s.wsManager.SetLogger(s.logger)
	}

	// Set up handler with middleware pipeline
	s.setupMiddleware()

	return s, nil
}

// Port returns the server's port
func (s *Server) Port() int {
	return s.port
}

// Start starts the HTTP server
func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return fmt.Errorf("server already running")
	}

	// Bind the socket first to catch port binding errors before returning success
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to bind port %d: %w", s.port, err)
	}
	s.listener = listener

	s.server = &http.Server{
		Handler: s.handler,
	}

	// Start file watcher if watch directory is set
	if s.watchDir != "" {
		fw, err := newFileWatcher(s.DebounceDuration(), s.logger)
		if err != nil {
			_ = listener.Close()
			return fmt.Errorf("failed to create file watcher: %w", err)
		}

		if err := fw.Watch(s.watchDir); err != nil {
			if closeErr := fw.Close(); closeErr != nil {
				s.logger.Error("Failed to close file watcher during cleanup: %v", closeErr)
			}
			_ = listener.Close()
			return fmt.Errorf("failed to watch directory: %w", err)
		}

		s.watcher = fw

		// Start file change handler
		go s.handleFileChanges()
	}

	// Start server in goroutine with pre-bound listener
	go func() {
		if err := s.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			s.logger.Error("Server error: %v", err)
		}
	}()

	s.running = true
	s.logger.Info("Server started on port %d", s.port)
	return nil
}

// Close stops the server gracefully
func (s *Server) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if s.server != nil {
		if err := s.server.Shutdown(ctx); err != nil {
			return err
		}
	}

	// Note: server.Shutdown() already closes the listener, so we don't need to close it again
	// The listener is only closed explicitly in error paths during Start()

	if s.watcher != nil {
		if err := s.watcher.Close(); err != nil {
			s.logger.Error("Failed to close file watcher: %v", err)
		}
	}

	if s.generateSession != nil {
		s.generateSession.Close()
	}

	s.running = false
	s.logger.Info("Server stopped")
	return nil
}

// IsRunning returns whether the server is currently running
func (s *Server) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}

// WebSocketManager returns the WebSocket manager (nil if reload disabled)
func (s *Server) WebSocketManager() WebSocketManager {
	return s.wsManager
}

// SetWatchDir sets the directory to watch for file changes
func (s *Server) SetWatchDir(dir string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.watchDir = dir
	return nil
}

// WatchDir returns the current watch directory
func (s *Server) WatchDir() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.watchDir
}

// SetManifest sets the current manifest
func (s *Server) SetManifest(manifest []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Defensive copy to prevent caller from mutating our internal state
	s.manifest = make([]byte, len(manifest))
	copy(s.manifest, manifest)
	return nil
}

// Manifest returns the current manifest
func (s *Server) Manifest() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Defensive copy to prevent caller from mutating our internal state
	if s.manifest == nil {
		return nil, nil
	}
	manifestCopy := make([]byte, len(s.manifest))
	copy(manifestCopy, s.manifest)
	return manifestCopy, nil
}

// DebounceDuration returns the debounce duration (150ms per spec)
func (s *Server) DebounceDuration() time.Duration {
	return 150 * time.Millisecond
}

// Handler returns the HTTP handler
func (s *Server) Handler() http.Handler {
	return s.handler
}

// Logger returns the server logger
func (s *Server) Logger() Logger {
	return s.logger
}

// SetLogger sets the server's logger
func (s *Server) SetLogger(logger Logger) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.logger = logger
	if s.wsManager != nil {
		s.wsManager.SetLogger(logger)
		// Set WebSocket manager on logger for broadcasting logs
		if wsSetter, ok := logger.(interface{ SetWebSocketManager(WebSocketManager) }); ok {
			wsSetter.SetWebSocketManager(s.wsManager)
		}
	}
	// Rebuild middleware pipeline with new logger
	s.setupMiddleware()
}

// CreateReloadMessage creates a reload message JSON
func (s *Server) CreateReloadMessage(files []string, reason string) ([]byte, error) {
	msg := ReloadMessage{
		Type:   "reload",
		Reason: reason,
		Files:  files,
	}
	return json.Marshal(msg)
}

// BroadcastReload broadcasts a reload event to all WebSocket clients
func (s *Server) BroadcastReload(files []string, reason string) error {
	if s.wsManager == nil {
		return nil // Reload disabled
	}

	msgBytes, err := s.CreateReloadMessage(files, reason)
	if err != nil {
		return err
	}

	s.logger.Debug("Broadcasting reload: reason=%s, files=%v, clients=%d",
		reason, files, s.wsManager.ConnectionCount())

	return s.wsManager.Broadcast(msgBytes)
}

// RegenerateManifest triggers manifest regeneration
func (s *Server) RegenerateManifest() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.watchDir == "" {
		return fmt.Errorf("no watch directory set")
	}

	// Close old session if it exists
	if s.generateSession != nil {
		s.generateSession.Close()
		s.generateSession = nil
	}

	// Create fresh workspace context and session for live reload
	// This ensures we always read the latest file contents
	workspace := W.NewFileSystemWorkspaceContext(s.watchDir)
	if err := workspace.Init(); err != nil {
		return fmt.Errorf("initializing workspace: %w", err)
	}

	session, err := G.NewGenerateSession(workspace)
	if err != nil {
		return fmt.Errorf("creating generate session: %w", err)
	}
	s.generateSession = session

	// Generate manifest
	ctx := context.Background()
	pkg, err := s.generateSession.GenerateFullManifest(ctx)
	if err != nil {
		return fmt.Errorf("generating manifest: %w", err)
	}

	// Extract source files from manifest for targeted file watching
	sourceFiles := make(map[string]bool)
	for _, module := range pkg.Modules {
		if module.Path != "" {
			sourceFiles[module.Path] = true
		}
	}
	s.sourceFiles = sourceFiles
	s.logger.Debug("Tracking %d source files from manifest", len(sourceFiles))

	// Marshal to JSON
	manifestBytes, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling manifest: %w", err)
	}

	s.manifest = manifestBytes
	s.logger.Info("Manifest regenerated (%d bytes)", len(manifestBytes))
	return nil
}

// Use registers a middleware function
func (s *Server) Use(middleware func(http.Handler) http.Handler) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Apply middleware to existing handler
	if s.handler != nil {
		s.handler = middleware(s.handler)
	}

	return nil
}

// resolveSourceFile checks if a .js file has a corresponding .ts source
func (s *Server) resolveSourceFile(path string) string {
	// If it's a .js file, check if .ts exists
	if filepath.Ext(path) == ".js" {
		tsPath := path[:len(path)-3] + ".ts"
		fullTsPath := filepath.Join(s.watchDir, tsPath)
		if _, err := os.Stat(fullTsPath); err == nil {
			return tsPath
		}
	}
	return path
}

// handleFileChanges listens for file change events and triggers reload
func (s *Server) handleFileChanges() {
	if s.watcher == nil {
		return
	}

	for event := range s.watcher.Events() {
		relPath := event.Path
		if s.watchDir != "" {
			if rel, err := filepath.Rel(s.watchDir, event.Path); err == nil {
				relPath = rel
			}
		}

		// Check if this file is in our source files set
		s.mu.RLock()
		isSourceFile := s.sourceFiles != nil && s.sourceFiles[relPath]
		s.mu.RUnlock()

		// Skip files that aren't in the manifest
		if !isSourceFile {
			s.logger.Debug("Ignoring non-manifest file: %s", relPath)
			continue
		}

		// Resolve .js to .ts if source exists
		displayPath := s.resolveSourceFile(relPath)
		s.logger.Info("File changed: %s", displayPath)

		// Regenerate manifest if a source file changed
		ext := filepath.Ext(event.Path)
		if ext == ".ts" || ext == ".js" {
			s.logger.Debug("Regenerating manifest for %s file change...", ext)
			err := s.RegenerateManifest()
			if err != nil {
				s.logger.Error("Failed to regenerate manifest: %v", err)
				// Continue anyway - we still want to reload the page
			} else {
				s.logger.Debug("Manifest regenerated successfully")
			}
		}

		// Broadcast reload to all WebSocket clients
		files := []string{relPath}
		err := s.BroadcastReload(files, "file-change")
		if err != nil {
			s.logger.Error("Failed to broadcast reload: %v", err)
		}
	}
}

// setupMiddleware configures the middleware pipeline
func (s *Server) setupMiddleware() {
	// Create a mux for routing
	mux := http.NewServeMux()

	// WebSocket endpoint for live reload
	if s.wsManager != nil {
		mux.HandleFunc("/__cem-reload", s.wsManager.HandleConnection)
	}

	// Manifest endpoint
	mux.HandleFunc("/custom-elements.json", s.serveManifest)

	// Logs endpoint for debug console
	mux.HandleFunc("/__cem-logs", s.serveLogs)

	// Static file server for all other routes
	mux.HandleFunc("/", s.serveStaticFiles)

	// Apply middleware in reverse order (last to first)
	var handler http.Handler = mux

	// 8. WebSocket client injection (into HTML)
	handler = injectWebSocketClient(handler, s.config.Reload)

	// 7. Static fallback (already in mux)

	// 6. Demo rendering (Phase 3)
	// TODO: Add demo rendering middleware

	// 5. CSS transform (Phase 4)
	// TODO: Add CSS transform middleware

	// 4. TypeScript transform (Phase 4)
	// TODO: Add TypeScript transform middleware

	// 3. Import map injection (Phase 2)
	// TODO: Add import map injection middleware

	// 2. CORS headers
	handler = corsMiddleware(handler)

	// 1. Logging
	handler = loggingMiddleware(s.logger)(handler)

	s.handler = handler
}

// serveManifest serves the custom elements manifest
func (s *Server) serveManifest(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	manifest := s.manifest
	s.mu.RUnlock()

	if len(manifest) == 0 {
		http.Error(w, "Manifest not available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(manifest); err != nil {
		s.logger.Error("Failed to write manifest response: %v", err)
	}
}

// serveLogs serves the plain text logs for the debug console
func (s *Server) serveLogs(w http.ResponseWriter, r *http.Request) {
	// Get logs from logger if it supports Logs()
	if logGetter, ok := s.logger.(interface{ Logs() []string }); ok {
		logs := logGetter.Logs()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(logs); err != nil {
			s.logger.Error("Failed to encode logs response: %v", err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte("[]")); err != nil {
			s.logger.Error("Failed to write empty logs response: %v", err)
		}
	}
}

// serveStaticFiles serves static files from the watch directory
func (s *Server) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	watchDir := s.watchDir
	s.mu.RUnlock()

	if watchDir == "" {
		http.NotFound(w, r)
		return
	}

	// Check if requesting root and no index.html exists
	if r.URL.Path == "/" {
		if _, err := http.Dir(watchDir).Open("index.html"); err != nil {
			// No index.html, serve default page
			s.serveDefaultIndex(w, r)
			return
		}
	}

	// Serve files from watch directory
	fileServer := http.FileServer(http.Dir(watchDir))
	fileServer.ServeHTTP(w, r)
}

// serveDefaultIndex serves a default index page with debug info
func (s *Server) serveDefaultIndex(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	watchDir := s.watchDir
	manifestSize := len(s.manifest)
	manifestBytes := s.manifest
	wsConnections := 0
	if s.wsManager != nil {
		wsConnections = s.wsManager.ConnectionCount()
	}
	s.mu.RUnlock()

	// Parse manifest to extract custom elements using manifest package
	elementsList := ""
	if len(manifestBytes) > 0 {
		var pkg M.Package
		if err := json.Unmarshal(manifestBytes, &pkg); err == nil {
			tagNames := pkg.GetAllTagNames()
			if len(tagNames) > 0 {
				for _, tagName := range tagNames {
					elementsList += fmt.Sprintf("<li><code>&lt;%s&gt;</code></li>", tagName)
				}
			} else {
				elementsList = "<li><em>No custom elements found</em></li>"
			}
		} else {
			elementsList = "<li><em>Error parsing manifest</em></li>"
		}
	} else {
		elementsList = "<li><em>Manifest not generated yet</em></li>"
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CEM Dev Server</title>
  <style>
    body {
      font-family: system-ui, -apple-system, sans-serif;
      max-width: 800px;
      margin: 40px auto;
      padding: 0 20px;
      line-height: 1.6;
      background: #0f172a;
      color: #e2e8f0;
    }
    h1 { color: #60a5fa; }
    .info {
      background: #1e293b;
      border-left: 4px solid #60a5fa;
      padding: 1rem;
      margin: 1rem 0;
      border-radius: 4px;
    }
    .debug {
      background: #1e293b;
      color: #10b981;
      padding: 1rem;
      border-radius: 4px;
      font-family: 'Courier New', monospace;
      font-size: 14px;
      overflow-x: auto;
      white-space: pre-wrap;
    }
    .status {
      display: inline-block;
      padding: 4px 8px;
      border-radius: 4px;
      font-size: 12px;
      font-weight: bold;
      background: #10b981;
      color: #0f172a;
    }
    code {
      background: #334155;
      padding: 2px 6px;
      border-radius: 3px;
      color: #fbbf24;
    }
    ul {
      list-style: none;
      padding: 0;
    }
    ul li {
      padding: 0.5rem;
      background: #1e293b;
      margin: 0.25rem 0;
      border-radius: 4px;
    }
  </style>
</head>
<body>
  <h1>🚀 CEM Development Server</h1>

  <div class="info">
    <p><strong>Status:</strong> <span class="status">Running</span></p>
    <p><strong>Watch Directory:</strong> <code>%s</code></p>
    <p><strong>Manifest:</strong> %d bytes</p>
    <p><strong>WebSocket Connections:</strong> %d</p>
  </div>

  <details>
    <summary><h2 style="display: inline; cursor: pointer;">Custom Elements</h2></summary>
    <ul>%s</ul>
  </details>

  <h2>Debug Console</h2>
  <div class="debug" id="messages">Waiting for WebSocket messages...</div>

  <script>
    const messagesDiv = document.getElementById('messages');

    // Fetch initial logs
    async function loadInitialLogs() {
      try {
        const response = await fetch('/__cem-logs');
        const logs = await response.json();
        messagesDiv.textContent = logs.length > 0 ? logs.join('\n') : 'No logs yet...';
        messagesDiv.scrollTop = messagesDiv.scrollHeight;
      } catch (e) {
        console.error('Failed to fetch logs:', e);
      }
    }

    // Load initial logs
    loadInitialLogs();

    // Listen for log updates via WebSocket
    window.addEventListener('load', () => {
      setTimeout(() => {
        if (window.__cemReloadSocket) {
          console.log('[cem-serve] WebSocket client connected');

          // Store original onmessage handler
          const originalOnMessage = window.__cemReloadSocket.onmessage;

          // Handle both reload and log messages
          window.__cemReloadSocket.onmessage = function(event) {
            const data = JSON.parse(event.data);

            if (data.type === 'logs') {
              // Update logs display
              messagesDiv.textContent = data.logs.length > 0 ? data.logs.join('\n') : 'No logs yet...';
              messagesDiv.scrollTop = messagesDiv.scrollHeight;
            } else if (data.type === 'reload') {
              // Call original handler for reload
              if (originalOnMessage) {
                originalOnMessage.call(this, event);
              }
            }
          };
        }
      }, 100);
    });
  </script>
</body>
</html>`, watchDir, manifestSize, wsConnections, elementsList)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(html)); err != nil {
		s.logger.Error("Failed to write default index response: %v", err)
	}
}

// corsMiddleware adds CORS headers to responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

// loggingMiddleware logs all requests except internal endpoints
func loggingMiddleware(logger Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip logging for internal polling endpoints
			if r.URL.Path != "/__cem-logs" && r.URL.Path != "/__cem-reload" {
				logger.Info("%s %s", r.Method, r.URL.Path)
			}
			next.ServeHTTP(w, r)
		})
	}
}

// defaultLogger is a simple logger implementation
type defaultLogger struct{}

func (l *defaultLogger) Info(msg string, args ...interface{}) {
	log.Printf("[INFO] "+msg, args...)
}

func (l *defaultLogger) Warning(msg string, args ...interface{}) {
	log.Printf("[WARN] "+msg, args...)
}

func (l *defaultLogger) Error(msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg, args...)
}

func (l *defaultLogger) Debug(msg string, args ...interface{}) {
	log.Printf("[DEBUG] "+msg, args...)
}
