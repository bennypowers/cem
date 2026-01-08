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
	"time"

	"bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware/transform"
	"bennypowers.dev/cem/serve/middleware/types"
)

// TransformConfig holds transform-specific configuration
type TransformConfig struct {
	TypeScript TypeScriptConfig
	CSS        CSSConfig
}

// TypeScriptConfig holds TypeScript transform configuration
type TypeScriptConfig struct {
	Enabled bool
	Target  transform.Target
}

// CSSConfig holds CSS transform configuration
type CSSConfig struct {
	Enabled bool
	Include []string // Glob patterns to include
	Exclude []string // Glob patterns to exclude
}

// DemosConfig holds demo rendering configuration
type DemosConfig struct {
	Rendering string // Default rendering mode: "light", "shadow", or "iframe"
}

// Config represents the dev server configuration
type Config struct {
	Port                 int
	Reload               bool
	Target               transform.Target      // Transform target (default: ES2022) - deprecated, use Transforms.TypeScript.Target
	Transforms           TransformConfig       // Transform configuration
	ImportMap            types.ImportMapConfig // Import map override configuration
	Demos                DemosConfig           // Demo rendering configuration
	ConfigFile           string                // Path to config file (for error reporting)
	WatchIgnore          []string              // Glob patterns to ignore in file watcher (e.g., ["_site/**", "dist/**"])
	SourceControlRootURL string                // Source control root URL for demo routing (e.g., "https://github.com/user/repo/tree/main/")
	FS                   platform.FileSystem   // Optional filesystem for testing (defaults to os package)
	URLRewrites          []config.URLRewrite   // URL rewrites for request path mapping (e.g., "/dist/:path*" -> "/src/{{.path}}")
	WebSocketManager     WebSocketManager      // Optional WebSocket manager for testing (created automatically if nil and Reload=true)
}

// ReloadMessage represents a WebSocket reload event
type ReloadMessage struct {
	Type   string   `json:"type"`
	Reason string   `json:"reason"`
	Files  []string `json:"files"`
}

// LogMessage represents a WebSocket log update
type LogMessage struct {
	Type string   `json:"type"`
	Logs []string `json:"logs"`
}

// ErrorMessage represents a WebSocket error notification (e.g., transform errors)
type ErrorMessage struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Message string `json:"message"`
	File    string `json:"file,omitempty"`
}

// Logger is a type alias for the logger.Logger interface
type Logger = logger.Logger

// WebSocketManager manages WebSocket connections for live reload
type WebSocketManager interface {
	ConnectionCount() int
	Broadcast(message []byte) error
	BroadcastToPages(message []byte, pageURLs []string) error
	BroadcastShutdown() error
	CloseAll() error
	HandleConnection(w http.ResponseWriter, r *http.Request)
	SetLogger(logger Logger)
}

// FileWatcher watches files for changes
type FileWatcher interface {
	Watch(path string) error
	WatchPaths(paths []string) error // Watch specific file paths (for config files)
	Events() <-chan FileEvent
	Close() error
	SetIgnorePatterns(watchDir string, patterns []string)
}

// FileEvent represents a file system event
type FileEvent struct {
	Path           string   // Primary file path (for single file events)
	Paths          []string // All changed file paths (for batched events)
	EventType      string   // Event type for primary file (create/delete/modify)
	HasCreates     bool     // True if any files were created
	HasDeletes     bool     // True if any files were deleted
	HasPackageJSON bool     // True if package.json was modified
	Timestamp      time.Time
}
