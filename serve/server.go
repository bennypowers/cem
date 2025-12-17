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
	"maps"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/netutil"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/cors"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/inject"
	"bennypowers.dev/cem/serve/middleware/requestlogger"
	"bennypowers.dev/cem/serve/middleware/routes"
	"bennypowers.dev/cem/serve/middleware/shadowroot"
	"bennypowers.dev/cem/serve/middleware/transform"
	W "bennypowers.dev/cem/workspace"
	"golang.org/x/net/html"
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
	transformPool   *transform.Pool  // Worker pool for transforms with backpressure
	running         bool
	mu              sync.RWMutex
	generateSession *G.GenerateSession
	fs              platform.FileSystem // Filesystem abstraction for testability
	shutdown        chan struct{}       // Signal channel for graceful shutdown
	// Workspace mode fields
	isWorkspace       bool                          // True if serving a monorepo workspace
	workspaceRoot     string                        // Root directory of workspace
	workspacePackages []middleware.WorkspacePackage // Discovered packages with manifests
	// Cached routing table for demo routes (both workspace and single-package mode)
	demoRoutes           map[string]*routes.DemoRouteEntry
	importMap            *importmappkg.ImportMap     // Cached import map (workspace or single-package)
	sourceControlRootURL string                      // Source control root URL for demo routing
	templates            *routes.TemplateRegistry    // Template registry for HTML rendering
	pathMappings         map[string]string           // Path mappings for src/dist separation
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
	// No defaulting here - cmd/serve.go handles defaults via viper.IsSet()
	// Tests must explicitly set transform config values

	s := &Server{
		port:                 config.Port,
		config:               config,
		logger:               logger.NewDefaultLogger(),
		shutdown:             make(chan struct{}),
		sourceControlRootURL: config.SourceControlRootURL,
	}

	// Use provided filesystem or default to os package
	if config.FS != nil {
		s.fs = config.FS
	} else {
		s.fs = platform.NewOSFileSystem()
	}

	// Create WebSocket manager if reload is enabled
	if config.Reload {
		s.wsManager = newWebSocketManager()
		s.wsManager.SetLogger(s.logger)
	}

	// Initialize transform cache (500MB default)
	s.transformCache = transform.NewCache(500 * 1024 * 1024)

	// Initialize transform pool with adaptive sizing based on available CPUs
	// NumCPU/2 leaves headroom for HTTP serving, manifest generation, file watching
	// Cap at 8 to prevent thread explosion (esbuild creates ~3 OS threads per worker)
	maxWorkers := min(max(runtime.NumCPU()/2, 2), 8)
	queueDepth := maxWorkers * 12 // Proportional buffering for burst traffic
	s.transformPool = transform.NewPool(maxWorkers, queueDepth)

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

// buildConfigOverride converts serve config override to import map config override
// Creates defensive copies to prevent shared mutable state.
func (s *Server) buildConfigOverride() *importmappkg.ImportMap {
	if len(s.config.ImportMap.Override.Imports) == 0 && len(s.config.ImportMap.Override.Scopes) == 0 {
		return nil
	}

	result := &importmappkg.ImportMap{
		Imports: make(map[string]string, len(s.config.ImportMap.Override.Imports)),
		Scopes:  make(map[string]map[string]string, len(s.config.ImportMap.Override.Scopes)),
	}

	// Copy imports
	maps.Copy(result.Imports, s.config.ImportMap.Override.Imports)

	// Deep copy scopes
	for scopeKey, scopeMap := range s.config.ImportMap.Override.Scopes {
		result.Scopes[scopeKey] = make(map[string]string, len(scopeMap))
		maps.Copy(result.Scopes[scopeKey], scopeMap)
	}

	return result
}

// DemoRoutes returns the pre-computed demo routing table (both workspace and single-package mode)
func (s *Server) DemoRoutes() any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.demoRoutes
}

// SourceControlRootURL returns the source control root URL for demo routing
func (s *Server) SourceControlRootURL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sourceControlRootURL
}

// FileSystem returns the filesystem abstraction
func (s *Server) FileSystem() platform.FileSystem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.fs
}

// Start starts the HTTP server
func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return fmt.Errorf("server already running")
	}

	// Reinitialize shutdown channel if server was previously stopped
	select {
	case <-s.shutdown:
		s.shutdown = make(chan struct{})
	default:
		// Channel is still open
	}

	// Bind the socket first to catch port binding errors before returning success
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to bind port %d: %w", s.port, err)
	}

	// Limit total concurrent connections to prevent resource exhaustion
	// This caps HTTP handler goroutines even with HTTP/2 multiplexing
	listener = netutil.LimitListener(listener, 100)
	s.listener = listener

	// Configure HTTP/2 to limit concurrent streams per connection
	// This prevents goroutine explosion when browsers multiplex many requests
	http2Server := &http2.Server{
		MaxConcurrentStreams: 250, // Standard HTTP/2 limit
	}

	s.server = &http.Server{
		Handler:      s.handler,
		ReadTimeout:  30 * time.Second, // Protects against slow-read attacks
		WriteTimeout: 60 * time.Second, // Sufficient for transforms + manifest generation
		IdleTimeout:  90 * time.Second, // Balances keep-alive reuse vs resource cleanup
	}

	// Configure HTTP/2 with concurrent stream limits
	_ = http2.ConfigureServer(s.server, http2Server)

	// Start file watcher if watch directory is set
	if s.watchDir != "" {
		fw, err := newFileWatcher(s.DebounceDuration(), s.logger)
		if err != nil {
			_ = listener.Close()
			return fmt.Errorf("failed to create file watcher: %w", err)
		}

		// Set custom ignore patterns from config
		fw.SetIgnorePatterns(s.watchDir, s.config.WatchIgnore)

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
	s.logger.Debug("Server bound to port %d", s.port)
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
		// Give clients time to receive the shutdown message and close gracefully
		time.Sleep(250 * time.Millisecond)
		// Send close frames and close all WebSocket connections
		if err := s.wsManager.CloseAll(); err != nil {
			s.logger.Error("Failed to close WebSocket connections: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if s.server != nil {
		// Shutdown waits for connections to finish, but with active browser sessions
		// this can timeout. Just log the error and continue with cleanup.
		if err := s.server.Shutdown(ctx); err != nil {
			s.logger.Warning("Server shutdown timeout (active connections): %v", err)
			// Force close by closing the listener
			if s.listener != nil {
				_ = s.listener.Close()
			}
		}
	}

	// Note: server.Shutdown() already closes the listener, so we don't need to close it again
	// The listener is only closed explicitly in error paths during Start()

	if s.watcher != nil {
		if err := s.watcher.Close(); err != nil {
			s.logger.Error("Failed to close file watcher: %v", err)
		}
	}

	// Close transform pool to stop accepting new tasks
	if s.transformPool != nil {
		s.transformPool.Close()
	}

	if s.generateSession != nil {
		s.generateSession.Close()
	}

	s.running = false

	// Signal shutdown to all goroutines
	close(s.shutdown)

	s.logger.Info("Server stopped")
	return nil
}

// IsRunning returns whether the server is currently running
func (s *Server) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}

// Done returns a channel that's closed when the server shuts down.
// This allows goroutines to gracefully cancel work during shutdown.
func (s *Server) Done() <-chan struct{} {
	return s.shutdown
}

// WebSocketManager returns the WebSocket manager (nil if reload disabled)
func (s *Server) WebSocketManager() WebSocketManager {
	return s.wsManager
}

// SetWatchDir sets the directory to watch for file changes
func (s *Server) SetWatchDir(dir string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Convert to absolute path for consistent behavior
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("converting to absolute path: %w", err)
	}
	s.watchDir = absDir

	// Load tsconfig - try tsconfig.settings.json first (common in monorepos),
	// then fall back to tsconfig.json
	tsconfigPaths := []string{
		filepath.Join(absDir, "tsconfig.settings.json"),
		filepath.Join(absDir, "tsconfig.json"),
	}

	loaded := false
	for _, tsconfigPath := range tsconfigPaths {
		if data, err := s.fs.ReadFile(tsconfigPath); err == nil {
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

	// Parse tsconfig for path mappings (supports src/dist separation)
	var pathMappings map[string]string
	for _, tsconfigPath := range tsconfigPaths {
		if mappings, err := transform.ParseTsConfig(tsconfigPath, s.fs); err == nil {
			pathMappings = mappings
			if len(mappings) > 0 {
				s.logger.Debug("Extracted path mappings from %s: %v", tsconfigPath, mappings)
			}
			break
		}
	}

	// Merge with explicit config overrides (config takes precedence)
	if s.config.PathMappings != nil {
		if pathMappings == nil {
			pathMappings = make(map[string]string)
		}
		for k, v := range s.config.PathMappings {
			pathMappings[k] = v
			s.logger.Debug("Config override path mapping: %s -> %s", k, v)
		}
	}

	s.pathMappings = pathMappings

	// Generate import map for single-package mode
	// (Workspace mode generates in InitializeWorkspaceMode instead)
	if !s.isWorkspace && s.config.ImportMap.Generate {
		importMap, err := importmappkg.Generate(dir, &importmappkg.Config{
			InputMapPath:   s.config.ImportMap.OverrideFile,
			ConfigOverride: s.buildConfigOverride(),
			Logger:         s.logger,
			FS:             s.fs,
		})
		if err != nil {
			s.logger.Warning("Failed to generate import map: %v", err)
			s.importMap = nil
		} else {
			s.importMap = importMap
			s.logger.Debug("Generated import map for single-package mode")
		}
	} else if !s.config.ImportMap.Generate {
		s.logger.Debug("Import map generation disabled")
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
	routingTable, err := routes.BuildDemoRoutingTable(manifest, s.sourceControlRootURL)
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

// PackageJSON returns parsed package.json from the watch directory
func (s *Server) PackageJSON() (*middleware.PackageJSON, error) {
	watchDir := s.WatchDir()
	if watchDir == "" {
		return nil, nil
	}

	packageJSONPath := filepath.Join(watchDir, "package.json")

	var data []byte
	var err error

	// Use injected filesystem if available, otherwise fall back to os.ReadFile
	if filesystem := s.FileSystem(); filesystem != nil {
		data, err = filesystem.ReadFile(packageJSONPath)
	} else {
		data, err = os.ReadFile(packageJSONPath)
	}

	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var pkg struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	return &middleware.PackageJSON{
		Name:    pkg.Name,
		Version: pkg.Version,
	}, nil
}

// DebounceDuration returns the debounce duration for file watching
func (s *Server) DebounceDuration() time.Duration {
	return 50 * time.Millisecond
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

// TryLoadExistingManifest attempts to load a manifest from disk
// Returns the manifest size in bytes and any error
// Returns 0, nil if no manifest file exists (not an error condition)
func (s *Server) TryLoadExistingManifest() (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.watchDir == "" {
		return 0, fmt.Errorf("no watch directory set")
	}

	// Determine manifest path
	workspace := W.NewFileSystemWorkspaceContext(s.watchDir)
	if err := workspace.Init(); err != nil {
		return 0, fmt.Errorf("initializing workspace: %w", err)
	}

	manifestPath := workspace.CustomElementsManifestPath()
	if manifestPath == "" {
		return 0, nil // No manifest path configured
	}

	// Check if manifest file exists
	if _, err := s.fs.Stat(manifestPath); err != nil {
		if os.IsNotExist(err) {
			return 0, nil // File doesn't exist, not an error
		}
		return 0, fmt.Errorf("checking manifest file: %w", err)
	}

	// Read existing manifest
	manifestBytes, err := s.fs.ReadFile(manifestPath)
	if err != nil {
		return 0, fmt.Errorf("reading manifest file: %w", err)
	}

	// Validate it's valid JSON
	var pkg any
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return 0, fmt.Errorf("invalid manifest JSON: %w", err)
	}

	// Store the manifest
	s.manifest = manifestBytes

	// Build routing table from manifest
	routingTable, err := routes.BuildDemoRoutingTable(manifestBytes, s.sourceControlRootURL)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table from cached manifest: %v", err)
		s.demoRoutes = nil
	} else {
		s.demoRoutes = routingTable
		s.logger.Debug("Built routing table with %d demo routes from cached manifest", len(routingTable))
	}

	return len(manifestBytes), nil
}

// RegenerateManifest triggers manifest regeneration
// RegenerateManifest performs a full manifest regeneration
// Returns the manifest size in bytes and any error
func (s *Server) RegenerateManifest() (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.watchDir == "" {
		return 0, fmt.Errorf("no watch directory set")
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
		return 0, fmt.Errorf("initializing workspace: %w", err)
	}

	session, err := G.NewGenerateSession(workspace)
	if err != nil {
		return 0, fmt.Errorf("creating generate session: %w", err)
	}

	// Configure adaptive worker count to prevent goroutine explosion under concurrent load
	// Use same formula as transform pool for consistency
	maxWorkers := min(max(runtime.NumCPU()/2, 2), 8)
	session.SetMaxWorkers(maxWorkers)

	s.generateSession = session

	// Generate manifest
	ctx := context.Background()
	pkg, err := s.generateSession.GenerateFullManifest(ctx)
	if err != nil {
		return 0, fmt.Errorf("generating manifest: %w", err)
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
		return 0, fmt.Errorf("marshaling manifest: %w", err)
	}

	s.manifest = manifestBytes

	// Build routing table from manifest
	routingTable, err := routes.BuildDemoRoutingTable(manifestBytes, s.sourceControlRootURL)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table: %v", err)
		s.demoRoutes = nil
	} else {
		s.demoRoutes = routingTable
		s.logger.Debug("Built routing table with %d demo routes", len(routingTable))
	}
	return len(manifestBytes), nil
}

// RegenerateManifestIncremental incrementally updates the manifest for changed files
// Returns the manifest size in bytes and any error
func (s *Server) RegenerateManifestIncremental(changedFiles []string) (int, error) {
	s.mu.Lock()

	if s.watchDir == "" {
		s.mu.Unlock()
		return 0, fmt.Errorf("no watch directory set")
	}

	// If no generate session exists yet, do a full regeneration
	if s.generateSession == nil {
		s.mu.Unlock() // Unlock before calling RegenerateManifest which will lock
		return s.RegenerateManifest()
	}

	// Use incremental processing with existing session
	ctx := context.Background()
	pkg, err := s.generateSession.ProcessChangedFiles(ctx, changedFiles)
	if err != nil {
		// If incremental processing fails, fall back to full regeneration
		s.logger.Warning("Incremental manifest generation failed, falling back to full regeneration: %v", err)
		s.mu.Unlock() // Unlock before calling RegenerateManifest which will lock
		return s.RegenerateManifest()
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
		s.mu.Unlock()
		return 0, fmt.Errorf("marshaling manifest: %w", err)
	}

	s.manifest = manifestBytes

	// Build routing table from manifest
	routingTable, err := routes.BuildDemoRoutingTable(manifestBytes, s.sourceControlRootURL)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table: %v", err)
		s.demoRoutes = nil
	} else {
		s.demoRoutes = routingTable
		s.logger.Debug("Built routing table with %d demo routes", len(routingTable))
	}

	s.mu.Unlock()
	return len(manifestBytes), nil
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

	// Discover packages with manifests
	packages, err := discoverWorkspacePackages(s.watchDir)
	if err != nil {
		return fmt.Errorf("discovering workspace packages: %w", err)
	}

	// If no packages have customElements field, fall back to single-package mode
	// This handles workspaces where packages haven't been configured yet
	if len(packages) == 0 {
		s.logger.Debug("Workspace detected but no packages have customElements field - using single-package mode")
		s.isWorkspace = false
		return nil
	}

	s.logger.Debug("Detected workspace mode - found %d packages with manifests", len(packages))
	s.isWorkspace = true
	s.workspaceRoot = s.watchDir
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
	s.logger.Debug("Built routing table with %d demo routes", len(workspaceRoutingTable))

	// Generate workspace import map using middleware package
	if s.config.ImportMap.Generate {
		importMap, err := importmappkg.Generate(s.workspaceRoot, &importmappkg.Config{
			InputMapPath:      s.config.ImportMap.OverrideFile,
			ConfigOverride:    s.buildConfigOverride(),
			WorkspacePackages: packages,
			Logger:            s.logger,
			FS:                s.fs,
		})
		if err != nil {
			s.logger.Warning("Failed to generate workspace import map: %v", err)
			s.importMap = nil
		} else {
			s.importMap = importMap
			s.logger.Info("Generated workspace import map")
		}
	} else {
		s.logger.Debug("Import map generation disabled")
	}

	return nil
}

// resolveSourceFile checks if a .js file has a corresponding .ts source
func (s *Server) resolveSourceFile(path string) string {
	// If it's a .js file, check if .ts exists
	if filepath.Ext(path) == ".js" {
		tsPath := path[:len(path)-3] + ".ts"
		watchDir := s.WatchDir()
		if watchDir != "" {
			fullTsPath := filepath.Join(watchDir, tsPath)
			if fs := s.FileSystem(); fs != nil {
				if _, err := fs.Stat(fullTsPath); err == nil {
					return tsPath
				}
			} else if _, err := os.Stat(fullTsPath); err == nil {
				return tsPath
			}
		}
	}
	return path
}

// extractModuleImports parses an HTML file and extracts ES module import specifiers
// Returns both import specifiers from inline scripts and src URLs from script tags
func (s *Server) extractModuleImports(htmlPath string) ([]string, error) {
	s.mu.RLock()
	fs := s.fs
	s.mu.RUnlock()

	content, err := fs.ReadFile(htmlPath)
	if err != nil {
		return nil, err
	}

	imports := make(map[string]bool) // Use map to deduplicate
	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return nil, err
	}

	// Import regex patterns
	importFromRe := regexp.MustCompile(`import\s+(?:[^'"]*?)\s*from\s*['"]([^'"]+)['"]`)
	importRe := regexp.MustCompile(`import\s*['"]([^'"]+)['"]`)

	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			// Check if it's a module script
			isModule := false
			var srcAttr string
			for _, attr := range n.Attr {
				if attr.Key == "type" && attr.Val == "module" {
					isModule = true
				}
				if attr.Key == "src" {
					srcAttr = attr.Val
				}
			}

			if isModule {
				// External module script
				if srcAttr != "" {
					imports[srcAttr] = true
				}

				// Inline module script - extract import statements
				if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
					scriptContent := n.FirstChild.Data

					// Extract "import ... from 'specifier'" statements
					matches := importFromRe.FindAllStringSubmatch(scriptContent, -1)
					for _, match := range matches {
						if len(match) > 1 {
							imports[match[1]] = true
						}
					}

					// Extract "import 'specifier'" statements
					matches = importRe.FindAllStringSubmatch(scriptContent, -1)
					for _, match := range matches {
						if len(match) > 1 && !strings.Contains(match[0], " from ") {
							imports[match[1]] = true
						}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}

	visit(doc)

	// Convert map to slice
	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}

	return result, nil
}

// resolveImportToPath converts an import specifier to a file path
// Handles bare specifiers (via import map), relative paths, and absolute paths
// contextDir is the directory of the HTML file making the import (for resolving relative imports)
func (s *Server) resolveImportToPath(importSpec string, contextDir string) []string {
	paths := make([]string, 0, 2)

	// If it's a relative path, resolve it relative to the context directory
	if strings.HasPrefix(importSpec, "./") || strings.HasPrefix(importSpec, "../") {
		var resolved string
		if contextDir != "" {
			// Resolve relative to the HTML file's directory
			resolved = filepath.Join(contextDir, importSpec)
		} else {
			// No context available, use import spec as-is
			resolved = importSpec
		}
		normalized := filepath.Clean(resolved)
		paths = append(paths, normalized)
		s.logger.Debug("Resolved relative import %s (context: %s) -> %s", importSpec, contextDir, normalized)
		return paths
	}

	// If it's an absolute path, use it directly
	if strings.HasPrefix(importSpec, "/") {
		normalized := filepath.Clean(importSpec)
		paths = append(paths, normalized)
		s.logger.Debug("Resolved absolute import %s -> %s", importSpec, normalized)
		return paths
	}

	// Try to resolve via import map
	s.mu.RLock()
	importMap := s.importMap
	s.mu.RUnlock()

	if importMap != nil {
		// First try exact match
		if resolved, ok := importMap.Imports[importSpec]; ok {
			s.logger.Debug("Import map entry: %s -> %s", importSpec, resolved)
			if parsedURL, err := url.Parse(resolved); err == nil && parsedURL.Path != "" {
				paths = append(paths, parsedURL.Path)
				s.logger.Debug("Resolved bare import %s -> %s (via import map exact match)", importSpec, parsedURL.Path)
			} else {
				s.logger.Debug("Failed to parse URL or extract path from: %s", resolved)
			}
		} else {
			// Try prefix matching for entries ending with "/"
			// Find the longest matching prefix
			var longestPrefix string
			var longestPrefixValue string
			for key, value := range importMap.Imports {
				if strings.HasSuffix(key, "/") && strings.HasPrefix(importSpec, key) {
					if len(key) > len(longestPrefix) {
						longestPrefix = key
						longestPrefixValue = value
					}
				}
			}

			if longestPrefix != "" {
				// Replace the prefix
				suffix := strings.TrimPrefix(importSpec, longestPrefix)
				resolved := longestPrefixValue + suffix
				s.logger.Debug("Import map prefix match: %s (prefix: %s -> %s)", importSpec, longestPrefix, longestPrefixValue)
				if parsedURL, err := url.Parse(resolved); err == nil && parsedURL.Path != "" {
					paths = append(paths, parsedURL.Path)
					s.logger.Debug("Resolved bare import %s -> %s (via import map prefix match)", importSpec, parsedURL.Path)
				} else {
					s.logger.Debug("Failed to parse URL or extract path from: %s", resolved)
				}
			} else {
				s.logger.Debug("Import %s not found in import map (have %d entries)", importSpec, len(importMap.Imports))
			}
		}

		// TODO(enhancement): Add support for scoped imports for path-specific resolution.
		// The import map spec supports a 'scopes' section for context-dependent module resolution.
		// This would allow different import resolutions based on the importing file's path.
		// Currently, we only resolve from the global 'imports' section, which handles most cases.
		// Ref: https://github.com/WICG/import-maps#scoping-examples
	} else {
		s.logger.Debug("No import map available to resolve %s", importSpec)
	}

	return paths
}

// logCacheStats periodically logs cache statistics
func (s *Server) logCacheStats(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if s.transformCache != nil {
				stats := s.transformCache.Stats()
				s.logger.Debug("Transform cache stats: %d entries, %.1f%% hit rate, %d MB / %d MB",
					stats.Entries,
					stats.HitRate,
					stats.SizeBytes/(1024*1024),
					stats.MaxSize/(1024*1024),
				)
			}
		case <-s.shutdown:
			return
		}
	}
}

// handleFileChanges listens for file change events and triggers reload
func (s *Server) handleFileChanges() {
	if s.watcher == nil {
		return
	}

	for event := range s.watcher.Events() {
		// Process all files in the batched event
		filesToProcess := event.Paths
		if len(filesToProcess) == 0 {
			filesToProcess = []string{event.Path}
		}

		// Filter to only relevant source files
		relevantFiles := make([]string, 0, len(filesToProcess))
		for _, filePath := range filesToProcess {
			ext := filepath.Ext(filePath)
			isRelevant := ext == ".ts" || ext == ".js" || ext == ".css" || ext == ".html"

			if !isRelevant {
				var relPath string
				if s.watchDir != "" {
					if rel, err := filepath.Rel(s.watchDir, filePath); err == nil {
						relPath = rel
					} else {
						relPath = filePath
					}
				} else {
					relPath = filePath
				}
				s.logger.Debug("Ignoring non-source file: %s", relPath)
				continue
			}

			relevantFiles = append(relevantFiles, filePath)
		}

		// Skip if no relevant files
		if len(relevantFiles) == 0 {
			continue
		}

		// Process the first relevant file (for manifest regeneration, etc.)
		// In the future, we could optimize to only regenerate once for multiple files
		changedPath := relevantFiles[0]
		relPath := changedPath
		if s.watchDir != "" {
			if rel, err := filepath.Rel(s.watchDir, changedPath); err == nil {
				relPath = rel
			}
		}

		// Resolve .js to .ts if source exists
		displayPath := s.resolveSourceFile(relPath)
		s.logger.Info("File changed: %s", displayPath)

		ext := filepath.Ext(changedPath)

		// Collect all affected files from both transform cache and module graph
		// This gives us complete dependency tracking:
		// - Transform cache: tracks dependencies from esbuild transforms
		// - Module graph: tracks ALL module imports from manifest generation
		affectedFiles := make(map[string]bool)

		// Add transform cache invalidations
		if s.transformCache != nil {
			invalidatedFiles := s.transformCache.Invalidate(changedPath)
			if len(invalidatedFiles) > 0 {
				s.logger.Debug("Transform cache invalidated %d files: %v", len(invalidatedFiles), invalidatedFiles)
				for _, file := range invalidatedFiles {
					affectedFiles[file] = true
				}
			}
		}

		// Add module graph affected files (includes non-transformed .js files)
		moduleGraphFiles := s.getModuleGraphAffectedFiles(changedPath)
		if len(moduleGraphFiles) > 0 {
			for _, file := range moduleGraphFiles {
				affectedFiles[file] = true
			}
		}

		// Convert map to slice for downstream use
		var invalidatedFiles []string
		for file := range affectedFiles {
			invalidatedFiles = append(invalidatedFiles, file)
		}

		s.logger.Debug("Total affected files: %d (transform cache + module graph)", len(invalidatedFiles))

		// Regenerate manifest if a source file changed
		if ext == ".ts" || ext == ".js" {
			manifestStart := time.Now()
			s.logger.Debug("Regenerating manifest incrementally for %s file change...", ext)
			manifestSize, err := s.RegenerateManifestIncremental([]string{changedPath})
			manifestDuration := time.Since(manifestStart)
			if err != nil {
				s.logger.Error("Failed to regenerate manifest incrementally: %v", err)
				// Continue anyway - we still want to reload the page
			} else {
				s.logger.Info("Manifest regenerated incrementally (%d bytes) in %v", manifestSize, manifestDuration)
			}

			// Note: Import map regeneration is skipped for .ts/.js changes
			// Import maps are built from package.json exports, not source files
			// They only need to be regenerated when package.json changes
		}

		// Regenerate import map only when necessary and enabled:
		// - package.json was modified (exports may have changed)
		// - files were created (new modules may need to be mapped)
		// - files were deleted (old modules may need to be unmapped)
		if s.config.ImportMap.Generate && (event.HasPackageJSON || event.HasCreates || event.HasDeletes) {
			importMapStart := time.Now()
			s.mu.Lock()
			if s.isWorkspace {
				// Workspace mode: regenerate workspace import map
				importMap, err := importmappkg.Generate(s.workspaceRoot, &importmappkg.Config{
					InputMapPath:      s.config.ImportMap.OverrideFile,
					ConfigOverride:    s.buildConfigOverride(),
					WorkspacePackages: s.workspacePackages,
					Logger:            s.logger,
					FS:                s.fs,
				})
				if err != nil {
					s.logger.Warning("Failed to regenerate workspace import map: %v", err)
				} else {
					s.importMap = importMap
					importMapDuration := time.Since(importMapStart)
					s.logger.Info("Regenerated workspace import map in %v", importMapDuration)
				}
			} else {
				// Single-package mode: regenerate import map
				importMap, err := importmappkg.Generate(s.watchDir, &importmappkg.Config{
					InputMapPath:   s.config.ImportMap.OverrideFile,
					ConfigOverride: s.buildConfigOverride(),
					Logger:         s.logger,
					FS:             s.fs,
				})
				if err != nil {
					s.logger.Warning("Failed to regenerate import map: %v", err)
				} else {
					s.importMap = importMap
					importMapDuration := time.Since(importMapStart)
					s.logger.Info("Regenerated import map in %v", importMapDuration)
				}
			}
			s.mu.Unlock()
		}

		// Smart reload: only reload pages that import the changed file or its dependents
		affectedPageURLs := s.getAffectedPageURLs(changedPath, invalidatedFiles)

		if len(affectedPageURLs) == 0 {
			// If smart reload found no affected pages, check if we're in a "no routes" state
			// (e.g. no manifest yet). In this case, fallback to broadcasting to all clients.
			s.mu.RLock()
			noRoutes := len(s.demoRoutes) == 0
			s.mu.RUnlock()

			if noRoutes {
				s.logger.Debug("No demo routes found, falling back to broadcast all for %s", relPath)
				// Create a "broadcast all" message (using empty affected list effectively broadcasts to all if we change logic,
				// but here we just pass the file and rely on client side or simply assume all pages need reload)
				// Actually BroadcastToPages with nil/empty list skips.
				// We should use Broadcast() instead.
				files := []string{relPath}
				if err := s.BroadcastReload(files, "file-change-fallback"); err != nil {
					s.logger.Error("Failed to broadcast reload: %v", err)
				}
				continue
			}

			s.logger.Debug("No pages affected by changes to %s", relPath)
			continue
		}

		// Broadcast reload only to affected pages
		files := []string{relPath}
		msgBytes, err := s.CreateReloadMessage(files, "file-change")
		if err != nil {
			s.logger.Error("Failed to create reload message: %v", err)
			continue
		}

		if s.wsManager != nil {
			err = s.wsManager.BroadcastToPages(msgBytes, affectedPageURLs)
			if err != nil {
				s.logger.Error("Failed to broadcast reload: %v", err)
			}
		}
	}
}

// getModuleGraphAffectedFiles queries the generate session's dependency tracker
// to find all modules transitively affected by a file change.
// Returns filesystem paths of affected modules.
func (s *Server) getModuleGraphAffectedFiles(changedPath string) []string {
	s.mu.RLock()
	session := s.generateSession
	watchDir := s.watchDir
	s.mu.RUnlock()

	if session == nil || watchDir == "" {
		return nil
	}

	// Get the workspace context from the session
	ctx := session.WorkspaceContext()
	if ctx == nil {
		return nil
	}

	// Convert filesystem path to module path
	modulePath, err := ctx.FSPathToModule(changedPath)
	if err != nil {
		s.logger.Debug("Failed to convert FS path to module path for %s: %v", changedPath, err)
		return nil
	}

	// Query the dependency tracker
	depTracker := session.DependencyTracker()
	if depTracker == nil {
		return nil
	}

	affectedModulePaths := depTracker.GetModulesAffectedByFiles([]string{modulePath})
	if len(affectedModulePaths) == 0 {
		return nil
	}

	// Convert module paths back to filesystem paths
	affectedFSPaths := make([]string, 0, len(affectedModulePaths))
	for _, modPath := range affectedModulePaths {
		fsPath := ctx.ModulePathToFS(modPath)
		affectedFSPaths = append(affectedFSPaths, fsPath)
	}

	s.logger.Debug("Module graph analysis: %d modules affected by %s", len(affectedFSPaths), modulePath)
	return affectedFSPaths
}

// getAffectedPageURLs returns page URLs that import the changed file or its dependents
func (s *Server) getAffectedPageURLs(changedPath string, invalidatedFiles []string) []string {
	s.logger.Debug("getAffectedPageURLs called with changedPath=%s, invalidatedFiles=%d", changedPath, len(invalidatedFiles))
	s.mu.RLock()
	watchDir := s.watchDir
	s.mu.RUnlock()

	// Get all files affected by this change (from cache invalidation)
	// Convert absolute paths to relative paths for comparison
	affectedFiles := make(map[string]bool)

	// Convert changed path to relative
	if rel, err := filepath.Rel(watchDir, changedPath); err == nil {
		affectedFiles[rel] = true
		affectedFiles["/"+rel] = true // Also store with leading slash
		s.logger.Debug("Changed file (relative): %s", rel)
		s.logger.Debug("Affected files map keys: %v", affectedFiles)
	} else {
		affectedFiles[changedPath] = true
		s.logger.Debug("Could not make relative, using absolute: %s", changedPath)
	}

	// Add all transitively invalidated files (also convert to relative)
	for _, path := range invalidatedFiles {
		if rel, err := filepath.Rel(watchDir, path); err == nil {
			affectedFiles[rel] = true
			affectedFiles["/"+rel] = true
		} else {
			affectedFiles[path] = true
		}
	}

	s.logger.Debug("Found %d affected file paths to check", len(affectedFiles))

	// Get demo routes to find HTML file paths
	s.mu.RLock()
	demoRoutes := s.demoRoutes
	s.mu.RUnlock()

	if demoRoutes == nil {
		s.logger.Debug("No demo routes available for smart reload")
		return nil
	}

	s.logger.Debug("Checking %d demo routes for affected imports", len(demoRoutes))
	affectedPages := make([]string, 0)

	// For each demo route, check if it imports any affected files
	for routePath, routeEntry := range demoRoutes {
		// First, check if the changed file IS this demo HTML file
		if affectedFiles[routeEntry.FilePath] ||
			affectedFiles["/"+routeEntry.FilePath] ||
			affectedFiles[strings.TrimPrefix(routeEntry.FilePath, "/")] {
			affectedPages = append(affectedPages, routePath)
			s.logger.Debug("Page %s is the changed file itself", routePath)
			continue // Already added, skip import checking for this route
		}

		htmlPath := filepath.Join(watchDir, routeEntry.FilePath)

		// Extract imports from HTML
		imports, err := s.extractModuleImports(htmlPath)
		if err != nil {
			s.logger.Debug("Failed to parse %s: %v", routeEntry.FilePath, err)
			continue
		}

		if len(imports) == 0 {
			continue // Skip routes with no imports
		}

		// Get the directory of the HTML file for resolving relative imports
		demoDir := filepath.Dir(routeEntry.FilePath)

		// Resolve imports to file paths and check if any match affected files
		for _, importSpec := range imports {
			resolvedPaths := s.resolveImportToPath(importSpec, demoDir)
			if len(resolvedPaths) == 0 {
				continue
			}

			for _, resolvedPath := range resolvedPaths {
				// Normalize the resolved path
				normalizedResolved := filepath.Clean(resolvedPath)

				// Also try with .ts extension instead of .js
				normalizedResolvedTS := normalizedResolved
				if strings.HasSuffix(normalizedResolved, ".js") {
					normalizedResolvedTS = normalizedResolved[:len(normalizedResolved)-3] + ".ts"
				}

				// Check if this resolved path matches any affected file
				if affectedFiles[normalizedResolved] ||
					affectedFiles[normalizedResolvedTS] ||
					affectedFiles[strings.TrimPrefix(normalizedResolved, "/")] ||
					affectedFiles[strings.TrimPrefix(normalizedResolvedTS, "/")] {
					affectedPages = append(affectedPages, routePath)
					s.logger.Debug("Page %s imports affected file %s (via %s)", routePath, normalizedResolved, importSpec)
					goto nextRoute // Found a match, move to next route
				}
			}
		}
	nextRoute:
	}

	if len(affectedPages) == 0 {
		s.logger.Debug("No pages import any of the affected files")
	}

	return affectedPages
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
		shadowroot.New(s.logger, errorBroadcaster{s}, routes.TemplatesFS, func(elementName string, data any) (string, error) {
			html, err := routes.RenderElementShadowRoot(s.templates, elementName, data)
			return string(html), err
		}), // Shadow root injection (last - processes final HTML)
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
			PathMappings:     s.pathMappings,
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
			PathMappings:     s.pathMappings,
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
	case ".json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	}

	// Write the content
	if _, err := w.Write(content); err != nil {
		s.logger.Error("Failed to write static file response: %v", err)
	}
}
