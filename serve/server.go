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
	"fmt"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/netutil"

	"bennypowers.dev/cem/cmd/config"
	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/routes"
	"bennypowers.dev/cem/serve/middleware/transform"
	"bennypowers.dev/cem/serve/middleware/types"
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
	demoRoutes              map[string]*types.DemoRouteEntry
	importMap               *importmappkg.ImportMap       // Cached import map (workspace or single-package)
	importMapGraph          *importmappkg.DependencyGraph // Dependency graph for incremental updates
	sourceControlRootURL    string                        // Source control root URL for demo routing
	templates               *routes.TemplateRegistry      // Template registry for HTML rendering
	urlRewrites             []config.URLRewrite           // URL rewrites for request path mapping
	pathResolver            *transform.PathResolver       // Cached path resolver (initialized once)
	pathResolverSourceFiles []string                      // Files that pathResolver depends on (for hot-reload)
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

	// Validate URL rewrites early to fail fast
	if len(config.URLRewrites) > 0 {
		if err := transform.ValidateURLRewrites(config.URLRewrites); err != nil {
			return nil, fmt.Errorf("invalid URL rewrites: %w", err)
		}
	}

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
		// Use provided WebSocketManager (for testing) or create default
		if config.WebSocketManager != nil {
			s.wsManager = config.WebSocketManager
		} else {
			s.wsManager = newWebSocketManager()
		}
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

// FileSystem returns the filesystem abstraction
func (s *Server) FileSystem() platform.FileSystem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.fs
}

// DemoRenderingMode returns the configured default rendering mode for demos
func (s *Server) DemoRenderingMode() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rendering := s.config.Demos.Rendering
	if rendering == "" {
		return "light" // default
	}
	return rendering
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

// DebounceDuration returns the debounce duration for file watching
func (s *Server) DebounceDuration() time.Duration {
	return 50 * time.Millisecond
}

// Logger returns the server logger
func (s *Server) Logger() middleware.Logger {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.logger
}

// SetLogger sets the server's logger
func (s *Server) SetLogger(log Logger) {
	s.mu.Lock()
	s.logger = log
	if s.wsManager != nil {
		s.wsManager.SetLogger(log)
		// Set WebSocket manager on logger for broadcasting logs
		if wsSetter, ok := log.(interface{ SetWebSocketManager(logger.Broadcaster) }); ok {
			wsSetter.SetWebSocketManager(s.wsManager)
		}
	}

	// Unlock before calling setupMiddleware to avoid deadlock
	// setupMiddleware calls accessor methods that acquire s.mu.RLock()
	s.mu.Unlock()

	// Rebuild middleware pipeline with new logger (no lock needed)
	s.setupMiddleware()
}

// WebSocketManager returns the WebSocket manager (nil if reload disabled)
func (s *Server) WebSocketManager() WebSocketManager {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.wsManager
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

		// Explicitly watch config files (cem.yaml, tsconfig.json) that affect path resolution
		// These bypass ignore patterns to ensure config hot-reload always works
		if len(s.pathResolverSourceFiles) > 0 {
			if err := fw.WatchPaths(s.pathResolverSourceFiles); err != nil {
				s.logger.Warning("Failed to watch some config files: %v", err)
				// Continue anyway - main watch is set up
			}
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

	// CRITICAL: Signal shutdown to all goroutines FIRST
	// This allows handleFileChanges and other goroutines to see the shutdown signal
	// and start exiting before we close their channels
	s.logger.Debug("Signaling shutdown to all goroutines")
	close(s.shutdown)

	// Give goroutines a brief moment to see shutdown signal and start exiting
	time.Sleep(50 * time.Millisecond)

	s.logger.Debug("Closing file watcher")
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
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.shutdown
}
