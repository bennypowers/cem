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
	"io/fs"
	"net/http"
	"os"
	"time"

	"bennypowers.dev/cem/serve/logger"
)

// FileSystem abstracts filesystem operations for testability
type FileSystem interface {
	ReadFile(name string) ([]byte, error)
	Stat(name string) (fs.FileInfo, error)
}

// osFileSystem implements FileSystem using the standard os package
type osFileSystem struct{}

func (osFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (osFileSystem) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

// Config represents the dev server configuration
type Config struct {
	Port   int
	Reload bool
	Target Target // Transform target (default: ES2022)
	FS     FileSystem // Optional filesystem for testing (defaults to os package)
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
	HandleConnection(w http.ResponseWriter, r *http.Request)
	SetLogger(logger Logger)
}

// FileWatcher watches files for changes
type FileWatcher interface {
	Watch(path string) error
	Events() <-chan FileEvent
	Close() error
}

// FileEvent represents a file system event
type FileEvent struct {
	Path             string   // Primary file path (for single file events)
	Paths            []string // All changed file paths (for batched events)
	EventType        string   // Event type for primary file (create/delete/modify)
	HasCreates       bool     // True if any files were created
	HasDeletes       bool     // True if any files were deleted
	HasPackageJSON   bool     // True if package.json was modified
	Timestamp        time.Time
}
