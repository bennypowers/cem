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

package logger

import "log"

// Logger is the logging interface used throughout the serve package
type Logger interface {
	Info(msg string, args ...any)
	Warning(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
	Success(msg string, args ...any)
	Trace(msg string, args ...any)
}

// Broadcaster defines the interface for broadcasting log messages to WebSocket clients
type Broadcaster interface {
	Broadcast([]byte) error
}

// LogMessage represents a log message broadcast to clients
type LogMessage struct {
	Type string     `json:"type"`
	Logs []LogEntry `json:"logs"`
}

// LogEntry represents a single structured log entry
type LogEntry struct {
	Type    string   `json:"type"`
	Date    string   `json:"date"`
	Message string   `json:"message"`
	Data    *LogData `json:"data,omitempty"`
}

// LogData carries structured payload for specialized log rendering.
type LogData struct {
	Kind      string          `json:"kind"`
	Durations []DurationEntry `json:"durations,omitempty"`
}

// DurationEntry represents a named duration for bar chart rendering.
type DurationEntry struct {
	Name     string  `json:"name"`
	Duration string  `json:"duration"`
	Percent  float64 `json:"percent"`
}

// defaultLogger is a simple logger implementation using standard log package
type defaultLogger struct{}

func (l *defaultLogger) Info(msg string, args ...any)    { log.Printf("[INFO] "+msg, args...) }
func (l *defaultLogger) Warning(msg string, args ...any) { log.Printf("[WARN] "+msg, args...) }
func (l *defaultLogger) Error(msg string, args ...any)   { log.Printf("[ERROR] "+msg, args...) }
func (l *defaultLogger) Debug(msg string, args ...any)   { log.Printf("[DEBUG] "+msg, args...) }
func (l *defaultLogger) Success(msg string, args ...any) { log.Printf("[SUCCESS] "+msg, args...) }
func (l *defaultLogger) Trace(msg string, args ...any)   { log.Printf("[TRACE] "+msg, args...) }

// NewDefaultLogger creates a simple logger using the standard log package
func NewDefaultLogger() Logger {
	return &defaultLogger{}
}
