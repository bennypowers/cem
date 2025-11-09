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
)

// Config represents the dev server configuration
type Config struct {
	Port   int
	Reload bool
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

// Logger interface for dev server logging
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
}

// WebSocketManager manages WebSocket connections for live reload
type WebSocketManager interface {
	ConnectionCount() int
	Broadcast(message []byte) error
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
	Path      string   // Primary file path (for single file events)
	Paths     []string // All changed file paths (for batched events)
	EventType string
	Timestamp time.Time
}
