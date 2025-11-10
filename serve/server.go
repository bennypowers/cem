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
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	C "bennypowers.dev/cem/cmd/config"
	G "bennypowers.dev/cem/generate"
	V "bennypowers.dev/cem/internal/version"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// PackageContext holds information about a single package in the workspace
type PackageContext struct {
	Name     string       // Package name from package.json
	Path     string       // Absolute path to package directory
	Manifest []byte       // Generated custom elements manifest
	Config   *C.CemConfig // Package-specific config
}

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
	tsconfigRaw     string          // Cached tsconfig.json for transforms
	transformCache  *TransformCache // LRU cache for transformed files
	running         bool
	mu              sync.RWMutex
	generateSession *G.GenerateSession
	// Workspace mode fields
	isWorkspace       bool                       // True if serving a monorepo workspace
	workspaceRoot     string                     // Root directory of workspace
	workspacePackages []PackageContext           // Discovered packages with manifests
	workspaceRoutes   map[string]*DemoRouteEntry // Combined routing table
	importMap         *ImportMap                 // Cached import map (workspace or single-package)
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

	// Initialize transform cache (500MB default)
	s.transformCache = NewTransformCache(500 * 1024 * 1024)

	// Set up handler with middleware pipeline
	s.setupMiddleware()

	return s, nil
}

// Port returns the server's port
func (s *Server) Port() int {
	return s.port
}

// IsWorkspace returns true if server is running in workspace mode
func (s *Server) IsWorkspace() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isWorkspace
}

// ImportMap returns the cached import map, generating it if needed for single-package mode
func (s *Server) ImportMap() (*ImportMap, error) {
	s.mu.RLock()
	cached := s.importMap
	isWorkspace := s.isWorkspace
	watchDir := s.watchDir
	s.mu.RUnlock()

	// If we have a cached import map (workspace mode), return it
	if cached != nil {
		return cached, nil
	}

	// Single-package mode: generate on demand
	if !isWorkspace && watchDir != "" {
		return GenerateImportMap(watchDir, nil)
	}

	return nil, nil
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

	// Start cache stats logger (every 5 minutes)
	if s.transformCache != nil {
		go s.logCacheStats(5 * time.Minute)
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

	// Send shutdown notification to connected clients before shutting down
	if s.wsManager != nil {
		if err := s.wsManager.BroadcastShutdown(); err != nil {
			s.logger.Error("Failed to broadcast shutdown notification: %v", err)
		}
		// Give clients a moment to receive the shutdown message
		time.Sleep(100 * time.Millisecond)
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

	// Load tsconfig - try tsconfig.settings.json first (common in monorepos),
	// then fall back to tsconfig.json
	tsconfigPaths := []string{
		filepath.Join(dir, "tsconfig.settings.json"),
		filepath.Join(dir, "tsconfig.json"),
	}

	loaded := false
	for _, tsconfigPath := range tsconfigPaths {
		if data, err := os.ReadFile(tsconfigPath); err == nil {
			s.tsconfigRaw = string(data)
			s.logger.Debug("Loaded TypeScript config from %s", tsconfigPath)
			loaded = true
			break
		}
	}

	if !loaded {
		s.tsconfigRaw = ""
		s.logger.Debug("No tsconfig found, using default transform settings")
	}

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

// BroadcastError broadcasts an error notification to all WebSocket clients
func (s *Server) BroadcastError(title, message, file string) error {
	if s.wsManager == nil {
		return nil // WebSocket disabled
	}

	msg := ErrorMessage{
		Type:    "error",
		Title:   title,
		Message: message,
		File:    file,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	s.logger.Debug("Broadcasting error: title=%s, file=%s, clients=%d",
		title, file, s.wsManager.ConnectionCount())

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

// discoverWorkspacePackages finds all packages in workspace with manifests
func discoverWorkspacePackages(rootDir string) ([]PackageContext, error) {
	packages, err := W.LoadWorkspaceManifests(rootDir)
	if err != nil {
		return nil, err
	}

	// Convert workspace.PackageWithManifest to serve.PackageContext
	contexts := make([]PackageContext, len(packages))
	for i, pkg := range packages {
		contexts[i] = PackageContext{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: pkg.Manifest,
			Config:   nil, // TODO: Load config from package
		}
	}

	return contexts, nil
}

// InitializeWorkspaceMode detects and initializes workspace mode if applicable
func (s *Server) InitializeWorkspaceMode() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.watchDir == "" {
		return fmt.Errorf("no watch directory set")
	}

	// Check if this is a workspace
	if !W.IsWorkspaceMode(s.watchDir) {
		s.isWorkspace = false
		return nil
	}

	s.logger.Info("Detected workspace mode - discovering packages...")
	s.isWorkspace = true
	s.workspaceRoot = s.watchDir

	// Discover packages with manifests
	packages, err := discoverWorkspacePackages(s.watchDir)
	if err != nil {
		return fmt.Errorf("discovering workspace packages: %w", err)
	}

	if len(packages) == 0 {
		return fmt.Errorf("no packages with customElements field found in workspace")
	}

	s.logger.Info("Found %d packages with manifests", len(packages))
	s.workspacePackages = packages

	// Build combined routing table
	routes, err := BuildWorkspaceRoutingTable(packages)
	if err != nil {
		return fmt.Errorf("building workspace routing table: %w", err)
	}

	s.workspaceRoutes = routes
	s.logger.Info("Built routing table with %d demo routes", len(routes))

	// Generate workspace import map using unified GenerateImportMap
	importMap, err := GenerateImportMap(s.workspaceRoot, &ImportMapConfig{
		WorkspacePackages: packages,
		Logger:            s.logger,
	})
	if err != nil {
		return fmt.Errorf("generating workspace import map: %w", err)
	}

	s.importMap = importMap
	s.logger.Info("Generated workspace import map")

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

// logCacheStats periodically logs cache statistics
func (s *Server) logCacheStats(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		// Check if server is still running
		s.mu.RLock()
		running := s.running
		s.mu.RUnlock()

		if !running {
			return
		}

		select {
		case <-ticker.C:
			if s.transformCache != nil {
				stats := s.transformCache.Stats()
				s.logger.Info("Transform cache stats: %d entries, %.1f%% hit rate, %d MB / %d MB",
					stats.Entries,
					stats.HitRate,
					stats.SizeBytes/(1024*1024),
					stats.MaxSize/(1024*1024),
				)
			}
		case <-time.After(1 * time.Second):
			// Check running status periodically
			continue
		}
	}
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

		// Invalidate transform cache for this file and its dependents
		if s.transformCache != nil {
			invalidated := s.transformCache.Invalidate(event.Path)
			if len(invalidated) > 0 {
				s.logger.Debug("Invalidated %d cached transforms", len(invalidated))
			}
		}

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

	// Debug info endpoint
	mux.HandleFunc("/__cem-debug", s.serveDebugInfo)

	// Internal JavaScript modules
	mux.HandleFunc("/__cem/", s.serveInternalModules)

	// Main handler: try demo routing first, then fall back to static files
	mux.HandleFunc("/", s.serveMainHandler)

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

// serveInternalModules serves embedded JavaScript modules from embed.FS
func (s *Server) serveInternalModules(w http.ResponseWriter, r *http.Request) {
	// Strip /__cem/ prefix to get the file path within the embedded FS
	// Request: /__cem/foo.js -> templates/js/foo.js in embed.FS
	path := strings.TrimPrefix(r.URL.Path, "/__cem/")
	path = "templates/js/" + path

	data, err := internalModules.ReadFile(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	if _, err := w.Write(data); err != nil {
		s.logger.Error("Failed to write JavaScript module response: %v", err)
	}
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

// getServerVersion returns version info for debug display
func getServerVersion() string {
	return V.GetVersion()
}

// serveDebugInfo serves debug information for the debug overlay
func (s *Server) serveDebugInfo(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	watchDir := s.watchDir
	manifestBytes := s.manifest
	s.mu.RUnlock()

	// Get import map
	var importMapObj any
	importMap, err := s.ImportMap()
	if err != nil {
		s.logger.Warning("Failed to get import map for debug info: %v", err)
	} else if importMap != nil && len(importMap.Imports) > 0 {
		importMapObj = importMap
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
		"version":      getServerVersion(),
		"os":           fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		"demos":        demos,
		"demoCount":    len(demos),
		"importMap":    importMapObj,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(debugInfo); err != nil {
		s.logger.Error("Failed to encode debug info response: %v", err)
	}
}

// serveMainHandler handles all non-special routes: tries demo routing first, then falls back to static files
func (s *Server) serveMainHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}

	// Try to serve as demo route
	html, err := s.serveDemoRoutes(r.URL.Path, queryParams)
	if err != nil {
		// Check if it's a routing table error (e.g., duplicate routes) vs just not found
		if strings.Contains(err.Error(), "duplicate demo route") {
			s.logger.Error("Demo routing error: %v", err)
			http.Error(w, "Internal Server Error: duplicate demo routes", http.StatusInternalServerError)
			return
		}
		// Not a demo route - fall through to static files
		s.serveStaticFiles(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(html)); err != nil {
		s.logger.Error("Failed to write demo response: %v", err)
	}
}

// serveStaticFiles serves static files from the watch directory
func (s *Server) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	watchDir := s.watchDir
	isWorkspace := s.isWorkspace
	s.mu.RUnlock()

	if watchDir == "" {
		http.NotFound(w, r)
		return
	}

	// Check if requesting root and no index.html exists
	if r.URL.Path == "/" {
		if isWorkspace {
			// Workspace mode: always serve index listing
			s.serveIndex(w, r)
			return
		}
		if _, err := http.Dir(watchDir).Open("index.html"); err != nil {
			// No index.html, serve index listing
			s.serveIndex(w, r)
			return
		}
	}

	// Transform TypeScript files to JavaScript
	requestPath := r.URL.Path
	ext := filepath.Ext(requestPath)

	// Handle both .js requests (with .ts source) and direct .ts requests
	var tsPath string
	switch ext {
	case ".js":
		// Check if .ts file exists for .js request
		tsPath = requestPath[:len(requestPath)-3] + ".ts"
	case ".ts":
		// Direct .ts request
		tsPath = requestPath
	}

	if tsPath != "" {
		// Strip leading slash and normalize path separators before joining
		tsPathNorm := strings.TrimPrefix(tsPath, "/")
		tsPathNorm = filepath.FromSlash(tsPathNorm)
		fullTsPath := filepath.Join(watchDir, tsPathNorm)

		// Get file stat for cache key
		fileInfo, err := os.Stat(fullTsPath)
		if err == nil {
			// .ts file exists - check cache first
			cacheKey := CacheKey{
				Path:    fullTsPath,
				ModTime: fileInfo.ModTime(),
				Size:    fileInfo.Size(),
			}

			// Try to get from cache
			if cached, found := s.transformCache.Get(cacheKey); found {
				s.logger.Debug("Cache hit for %s", tsPathNorm)
				// Serve cached transformed JavaScript
				w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
				if _, err := w.Write(cached.Code); err != nil {
					s.logger.Error("Failed to write cached transform response: %v", err)
				}
				return
			}

			// Cache miss - read file and transform
			s.logger.Debug("Cache miss for %s", tsPathNorm)
			source, err := os.ReadFile(fullTsPath)
			if err != nil {
				s.logger.Error("Failed to read TypeScript file %s: %v", tsPathNorm, err)
				http.Error(w, "Failed to read file", http.StatusInternalServerError)
				return
			}

			// Get tsconfig for transform
			s.mu.RLock()
			tsconfigRaw := s.tsconfigRaw
			s.mu.RUnlock()

			// Get configured target (defaults to ES2022 if not set)
			target := s.config.Target
			if target == "" {
				target = ES2022
			}

			// Transform TypeScript to JavaScript
			result, err := TransformTypeScript(source, TransformOptions{
				Loader:      LoaderTS,
				Target:      target,
				Sourcemap:   SourceMapInline,
				Sourcefile:  tsPathNorm, // Use normalized path for source maps
				TsconfigRaw: tsconfigRaw,
			})
			if err != nil {
				s.logger.Error("Failed to transform TypeScript file %s: %v", tsPathNorm, err)

				// Broadcast error to browser overlay
				s.BroadcastError(
					"TypeScript Transform Error",
					err.Error(),
					tsPathNorm,
				)

				http.Error(w, fmt.Sprintf("Transform error: %v", err), http.StatusInternalServerError)
				return
			}

			// Store in cache
			s.transformCache.Set(cacheKey, result.Code, result.Dependencies)

			// Serve transformed JavaScript
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
			if _, err := w.Write(result.Code); err != nil {
				s.logger.Error("Failed to write transform response: %v", err)
			}
			return
		}
	}

	// Transform CSS files to JavaScript modules (constructable stylesheets)
	if filepath.Ext(requestPath) == ".css" {
		// Strip leading slash and normalize
		cssPath := strings.TrimPrefix(requestPath, "/")
		cssPath = filepath.FromSlash(cssPath)
		fullCssPath := filepath.Join(watchDir, cssPath)

		// Read CSS file
		source, err := os.ReadFile(fullCssPath)
		if err != nil {
			// File doesn't exist - let normal file server handle 404
			fileServer := http.FileServer(http.Dir(watchDir))
			fileServer.ServeHTTP(w, r)
			return
		}

		// Transform CSS to JavaScript module
		transformed := TransformCSS(source, requestPath)

		// Serve as JavaScript module
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		if _, err := w.Write([]byte(transformed)); err != nil {
			s.logger.Error("Failed to write CSS transform response: %v", err)
		}
		return
	}

	// Set correct MIME type for JavaScript files (for plain .js files without .ts source)
	if filepath.Ext(requestPath) == ".js" {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	}

	// Serve files from watch directory
	fileServer := http.FileServer(http.Dir(watchDir))
	fileServer.ServeHTTP(w, r)
}

// getImportMapJSON returns the import map as indented JSON string
func (s *Server) getImportMapJSON() string {
	importMap, err := s.ImportMap()
	if err != nil {
		s.logger.Warning("Failed to get import map: %v", err)
		return ""
	}

	if importMap == nil || len(importMap.Imports) == 0 {
		return ""
	}

	importMapBytes, err := json.MarshalIndent(importMap, "  ", "  ")
	if err != nil {
		s.logger.Warning("Failed to marshal import map: %v", err)
		return ""
	}

	return string(importMapBytes)
}

// serveIndex serves the index page (workspace or single-package listing)
func (s *Server) serveIndex(w http.ResponseWriter, _ *http.Request) {
	importMapJSON := s.getImportMapJSON()

	var html string
	var err error

	s.mu.RLock()
	isWorkspace := s.isWorkspace
	s.mu.RUnlock()

	if isWorkspace {
		// Workspace mode: render workspace listing
		s.mu.RLock()
		packages := s.workspacePackages
		routes := s.workspaceRoutes
		s.mu.RUnlock()

		html, err = renderWorkspaceListing(packages, routes, importMapJSON)
	} else {
		// Single-package mode: render element listing
		s.mu.RLock()
		manifestBytes := s.manifest
		s.mu.RUnlock()

		html, err = renderElementListing(manifestBytes, importMapJSON)
	}

	if err != nil {
		s.logger.Error("Failed to render listing: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(html)); err != nil {
		s.logger.Error("Failed to write listing response: %v", err)
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

func (l *defaultLogger) Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

func (l *defaultLogger) Warning(msg string, args ...any) {
	log.Printf("[WARN] "+msg, args...)
}

func (l *defaultLogger) Error(msg string, args ...any) {
	log.Printf("[ERROR] "+msg, args...)
}

func (l *defaultLogger) Debug(msg string, args ...any) {
	log.Printf("[DEBUG] "+msg, args...)
}
