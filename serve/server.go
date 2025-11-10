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
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/cors"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/inject"
	"bennypowers.dev/cem/serve/middleware/requestlogger"
	"bennypowers.dev/cem/serve/middleware/routes"
	"bennypowers.dev/cem/serve/middleware/transform"
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
	sourceFiles     map[string]bool  // Set of source files to watch
	tsconfigRaw     string           // Cached tsconfig.json for transforms
	transformCache  *transform.Cache // LRU cache for transformed files
	running         bool
	mu              sync.RWMutex
	generateSession *G.GenerateSession
	// Workspace mode fields
	isWorkspace       bool                                // True if serving a monorepo workspace
	workspaceRoot     string                              // Root directory of workspace
	workspacePackages []middleware.WorkspacePackage       // Discovered packages with manifests
	// Cached routing table for demo routes (both workspace and single-package mode)
	demoRoutes map[string]*routes.DemoRouteEntry
	importMap  *importmappkg.ImportMap // Cached import map (workspace or single-package)
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
		logger: logger.NewDefaultLogger(),
	}

	// Create WebSocket manager if reload is enabled
	if config.Reload {
		s.wsManager = newWebSocketManager()
		s.wsManager.SetLogger(s.logger)
	}

	// Initialize transform cache (500MB default)
	s.transformCache = transform.NewCache(500 * 1024 * 1024)

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

// WorkspacePackages returns discovered workspace packages (workspace mode only)
func (s *Server) WorkspacePackages() []middleware.WorkspacePackage {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.workspacePackages
}

// ImportMap returns the cached import map (may be nil)
func (s *Server) ImportMap() middleware.ImportMap {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.importMap
}

// DemoRoutes returns the pre-computed demo routing table (both workspace and single-package mode)
func (s *Server) DemoRoutes() any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.demoRoutes
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

	// Generate import map for single-package mode
	// (Workspace mode generates in InitializeWorkspaceMode instead)
	if !s.isWorkspace {
		importMap, err := importmappkg.Generate(dir, nil)
		if err != nil {
			s.logger.Warning("Failed to generate import map: %v", err)
			s.importMap = nil
		} else {
			s.importMap = importMap
			s.logger.Debug("Generated import map for single-package mode")
		}
	}

	return nil
}

// WatchDir returns the current watch directory
func (s *Server) WatchDir() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.watchDir
}

// TsconfigRaw returns the current tsconfig.json content
func (s *Server) TsconfigRaw() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.tsconfigRaw
}

// SetManifest sets the current manifest and builds the routing table
func (s *Server) SetManifest(manifest []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Defensive copy to prevent caller from mutating our internal state
	s.manifest = make([]byte, len(manifest))
	copy(s.manifest, manifest)

	// Build routing table from manifest
	routingTable, err := routes.BuildDemoRoutingTable(manifest)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table: %v", err)
		s.demoRoutes = nil
	} else {
		s.demoRoutes = routingTable
		s.logger.Debug("Built routing table with %d demo routes", len(routingTable))
	}

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
func (s *Server) Logger() middleware.Logger {
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
		if wsSetter, ok := logger.(interface{ SetWebSocketManager(any) }); ok {
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

	// Build routing table from manifest
	routingTable, err := routes.BuildDemoRoutingTable(manifestBytes)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table: %v", err)
		s.demoRoutes = nil
	} else {
		s.demoRoutes = routingTable
		s.logger.Debug("Built routing table with %d demo routes", len(routingTable))
	}

	s.logger.Info("Manifest regenerated (%d bytes)", len(manifestBytes))
	return nil
}

// discoverWorkspacePackages finds all packages in workspace with manifests
func discoverWorkspacePackages(rootDir string) ([]middleware.WorkspacePackage, error) {
	packages, err := W.LoadWorkspaceManifests(rootDir)
	if err != nil {
		return nil, err
	}

	// Convert workspace.PackageWithManifest to middleware.WorkspacePackage
	contexts := make([]middleware.WorkspacePackage, len(packages))
	for i, pkg := range packages {
		contexts[i] = middleware.WorkspacePackage{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: pkg.Manifest,
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

	// Build routing table from workspace packages
	pkgContexts := make([]routes.PackageContext, len(packages))
	for i, pkg := range packages {
		pkgContexts[i] = routes.PackageContext{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: pkg.Manifest,
		}
	}
	workspaceRoutingTable, err := routes.BuildWorkspaceRoutingTable(pkgContexts)
	if err != nil {
		return fmt.Errorf("building workspace routing table: %w", err)
	}
	s.demoRoutes = workspaceRoutingTable
	s.logger.Info("Built routing table with %d demo routes", len(workspaceRoutingTable))

	// Convert middleware.WorkspacePackage to importmappkg.WorkspacePackage
	workspacePkgs := make([]importmappkg.WorkspacePackage, len(packages))
	for i, pkg := range packages {
		workspacePkgs[i] = importmappkg.WorkspacePackage{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: pkg.Manifest,
		}
	}

	// Generate workspace import map using middleware package
	importMap, err := importmappkg.Generate(s.workspaceRoot, &importmappkg.Config{
		WorkspacePackages: workspacePkgs,
		Logger:            s.logger,
	})
	if err != nil {
		return fmt.Errorf("generating workspace import map: %w", err)
	}

	s.importMap = importMap
	s.logger.Info("Generated workspace import map")

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

		// Check if this is a relevant source file
		ext := filepath.Ext(event.Path)
		isRelevant := ext == ".ts" || ext == ".js" || ext == ".css"

		// Skip files that aren't relevant source files
		if !isRelevant {
			s.logger.Debug("Ignoring non-source file: %s", relPath)
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
		if ext == ".ts" || ext == ".js" {
			s.logger.Debug("Regenerating manifest for %s file change...", ext)
			err := s.RegenerateManifest()
			if err != nil {
				s.logger.Error("Failed to regenerate manifest: %v", err)
				// Continue anyway - we still want to reload the page
			} else {
				s.logger.Debug("Manifest regenerated successfully")
			}

			// Regenerate import map (both workspace and single-package mode)
			s.mu.Lock()
			if s.isWorkspace {
				// Workspace mode: regenerate workspace import map
				workspacePkgs := make([]importmappkg.WorkspacePackage, len(s.workspacePackages))
				for i, pkg := range s.workspacePackages {
					workspacePkgs[i] = importmappkg.WorkspacePackage{
						Name:     pkg.Name,
						Path:     pkg.Path,
						Manifest: pkg.Manifest,
					}
				}
				importMap, err := importmappkg.Generate(s.workspaceRoot, &importmappkg.Config{
					WorkspacePackages: workspacePkgs,
					Logger:            s.logger,
				})
				if err != nil {
					s.logger.Warning("Failed to regenerate workspace import map: %v", err)
				} else {
					s.importMap = importMap
					s.logger.Debug("Regenerated workspace import map")
				}
			} else {
				// Single-package mode: regenerate import map
				importMap, err := importmappkg.Generate(s.watchDir, nil)
				if err != nil {
					s.logger.Warning("Failed to regenerate import map: %v", err)
				} else {
					s.importMap = importMap
					s.logger.Debug("Regenerated import map")
				}
			}
			s.mu.Unlock()
		}

		// Broadcast reload to all WebSocket clients
		files := []string{relPath}
		err := s.BroadcastReload(files, "file-change")
		if err != nil {
			s.logger.Error("Failed to broadcast reload: %v", err)
		}
	}
}

// errorBroadcaster adapts Server.BroadcastError to the middleware interface
type errorBroadcaster struct {
	*Server
}

func (e errorBroadcaster) BroadcastError(title, message, filename string) {
	_ = e.Server.BroadcastError(title, message, filename) // Ignore error
}

// getLogs returns logs if the logger supports it
func (s *Server) getLogs() []string {
	if logGetter, ok := s.logger.(interface{ Logs() []string }); ok {
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
	// Middlewares are applied in reverse order (last to first in the chain)
	// Terminal handler: static files
	s.handler = middleware.Chain(
		http.HandlerFunc(s.serveStaticFiles), // Static file server (terminal handler)
		inject.New(s.config.Reload, "/__cem/websocket-client.js"), // WebSocket injection
		importmappkg.New(importmappkg.MiddlewareConfig{ // Import map injection
			Context: s,
		}),
		transform.NewCSS(transform.CSSConfig{ // CSS transform
			WatchDirFunc: s.WatchDir,
			Logger:       s.logger,
		}),
		transform.NewTypeScript(transform.TypeScriptConfig{ // TypeScript transform
			WatchDirFunc:     s.WatchDir,
			TsconfigRawFunc:  s.TsconfigRaw,
			Cache:            s.transformCache,
			Logger:           s.logger,
			ErrorBroadcaster: errorBroadcaster{s},
			Target:           string(s.config.Target),
		}),
		routes.New(routes.Config{ // Internal CEM routes (includes WebSocket, demos, listings)
			Context:          s,
			LogsFunc:         s.getLogs,
			WebSocketHandler: wsHandler,
		}),
		cors.New(),                 // CORS headers
		requestlogger.New(s.logger), // HTTP request logging
	)
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

	// Set correct MIME type for JavaScript files
	if filepath.Ext(r.URL.Path) == ".js" {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	}

	// Serve files from watch directory
	fileServer := http.FileServer(http.Dir(watchDir))
	fileServer.ServeHTTP(w, r)
}
