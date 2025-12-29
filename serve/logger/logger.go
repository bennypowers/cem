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

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm"
	"golang.org/x/term"
)

// Logger is the logging interface used throughout the serve package
type Logger interface {
	Info(msg string, args ...any)
	Warning(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
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
	Type    string `json:"type"`    // "info", "warn", "error", "debug"
	Date    string `json:"date"`    // ISO8601 timestamp
	Message string `json:"message"` // Log message content
}

// defaultLogger is a simple logger implementation using standard log package
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

// NewDefaultLogger creates a simple logger using the standard log package
func NewDefaultLogger() Logger {
	return &defaultLogger{}
}

// ptermLogger implements Logger interface using pterm live rendering
type ptermLogger struct {
	verbose       bool
	logs          []LogEntry // Structured logs for web interface
	terminalLogs  []string   // Colored logs for terminal display
	pendingLogs   []pendingLog // Logs buffered before area starts
	maxLogs       int
	maxTermLogs   int
	mu            sync.Mutex
	renderMu      sync.Mutex // Serializes area.Update() calls to prevent interleaved output
	interactive   bool
	area          *pterm.AreaPrinter
	status        string
	wsManager     Broadcaster // WebSocket manager for broadcasting logs
}

// pendingLog represents a log entry waiting to be displayed
type pendingLog struct {
	level     string
	levelType string
	message   string
	timestamp string
}

// NewPtermLogger creates a new pterm-based logger with live rendering
func NewPtermLogger(verbose bool) Logger {
	interactive := term.IsTerminal(int(os.Stdout.Fd()))

	logger := &ptermLogger{
		verbose:      verbose,
		logs:         make([]LogEntry, 0),
		terminalLogs: make([]string, 0),
		pendingLogs:  make([]pendingLog, 0),
		maxLogs:      100,
		maxTermLogs:  50, // Keep last 50 logs visible in terminal
		interactive:  interactive,
		status:       "Starting...",
	}

	// Don't start live area immediately - let caller do it after initial setup
	// This prevents initial logs from appearing above the live area

	return logger
}

// Start starts the live rendering area (call after initial setup is complete)
func (l *ptermLogger) Start() {
	l.mu.Lock()
	if !l.interactive {
		l.mu.Unlock()
		return
	}

	if l.area != nil {
		l.mu.Unlock()
		l.render()
		return
	}

	// Capture pending logs before unlocking
	pending := l.pendingLogs
	l.pendingLogs = nil // Clear the pending buffer

	l.mu.Unlock()

	area, _ := pterm.DefaultArea.Start()

	l.mu.Lock()
	if l.area != nil {
		l.mu.Unlock()
		if area != nil {
			_ = area.Stop()
		}
		return
	}

	l.area = area

	// Format and add all pending logs now that area is ready
	for _, p := range pending {
		l.formatAndBufferLog(p.level, p.levelType, p.message, p.timestamp)
	}

	l.mu.Unlock()

	if area != nil {
		l.render()
	}
}

// SetStatus updates the status line and re-renders
func (l *ptermLogger) SetStatus(status string) {
	l.mu.Lock()
	l.status = status
	l.mu.Unlock()
	if l.interactive {
		l.render()
	}
}

// SetWebSocketManager sets the WebSocket manager for broadcasting logs
func (l *ptermLogger) SetWebSocketManager(wsManager Broadcaster) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.wsManager = wsManager
}

// render updates the live terminal display
func (l *ptermLogger) render() {
	l.mu.Lock()

	if !l.interactive || l.area == nil {
		l.mu.Unlock()
		return
	}

	var sb strings.Builder

	// Render logs with colors
	for _, log := range l.terminalLogs {
		sb.WriteString(log + "\n")
	}

	// Add separator with gray line
	sb.WriteString("\n" + pterm.FgGray.Sprint(strings.Repeat("─", 80)) + "\n")

	// Status line at bottom (colors are applied at call site in cmd/serve.go)
	// Example: ● Running on http://localhost:8080 | Live reload: true | Press Ctrl+C to stop
	sb.WriteString(pterm.FgLightGreen.Sprint("● ") + l.status)

	// Hand off pointer before calling into pterm API
	area := l.area
	output := sb.String()
	l.mu.Unlock()

	// Serialize area.Update() calls to prevent concurrent renders from interleaving output
	l.renderMu.Lock()
	area.Update(output)
	l.renderMu.Unlock()
}

// formatAndBufferLog formats a log entry and adds it to terminalLogs buffer.
// Must be called with lock held.
func (l *ptermLogger) formatAndBufferLog(level, levelType, message, timestamp string) {
	var prefix, coloredMsg string
	timestampStr := pterm.FgGray.Sprint(timestamp)

	switch levelType {
	case "info":
		prefix = pterm.FgCyan.Sprint("INFO ")
		coloredMsg = message
	case "warning":
		prefix = pterm.FgYellow.Sprint("WARN ")
		coloredMsg = pterm.FgYellow.Sprint(message)
	case "error":
		prefix = pterm.FgRed.Sprint("ERROR")
		coloredMsg = pterm.FgRed.Sprint(message)
	case "debug":
		prefix = pterm.FgGray.Sprint("DEBUG")
		coloredMsg = pterm.FgGray.Sprint(message)
	}

	// Calculate padding for right-aligned timestamp
	width := 80
	if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 {
		width = w
	}

	// Visual length is just the actual text without ANSI codes
	visualLen := len(level) + 1 + len(message)
	timestampBuffer := 10
	padding := max(width-visualLen-timestampBuffer, 1)

	terminalLog := fmt.Sprintf(" %s %s%s%s", prefix, coloredMsg, strings.Repeat(" ", padding), timestampStr)
	l.terminalLogs = append(l.terminalLogs, terminalLog)

	if len(l.terminalLogs) > l.maxTermLogs {
		l.terminalLogs = l.terminalLogs[len(l.terminalLogs)-l.maxTermLogs:]
	}
}

// Stop stops the live rendering
func (l *ptermLogger) Stop() {
	l.mu.Lock()
	area := l.area
	l.area = nil
	l.mu.Unlock()

	if area != nil {
		if err := area.Stop(); err != nil {
			// Log to stderr since our logging system may be shutting down
			fmt.Fprintf(os.Stderr, "Warning: failed to stop area printer: %v\n", err)
		}
	}
}

func (l *ptermLogger) log(level, levelType, msg string, args ...any) {
	formatted := fmt.Sprintf(msg, args...)
	now := time.Now()
	timestamp := now.Format("15:04:05")

	// Store structured log entry for web interface
	l.mu.Lock()
	logEntry := LogEntry{
		Type:    levelType, // "info", "warning", "error", "debug"
		Date:    now.Format(time.RFC3339),
		Message: formatted,
	}
	l.logs = append(l.logs, logEntry)
	if len(l.logs) > l.maxLogs {
		l.logs = l.logs[len(l.logs)-l.maxLogs:]
	}

	// Capture wsManager reference while holding lock to avoid race with SetWebSocketManager
	ws := l.wsManager

	shouldPrint := levelType != "debug" || l.verbose

	if shouldPrint {
		if l.interactive {
			// Interactive mode: either buffer semantic data or format+render
			if l.area != nil {
				// Area is ready: format and render immediately
				l.formatAndBufferLog(level, levelType, formatted, timestamp)
				l.mu.Unlock()
				l.render()
			} else {
				// Area not ready yet: buffer semantic data for later formatting
				l.pendingLogs = append(l.pendingLogs, pendingLog{
					level:     level,
					levelType: levelType,
					message:   formatted,
					timestamp: timestamp,
				})
				l.mu.Unlock()
			}
		} else {
			l.mu.Unlock()
			// Non-interactive: standard pterm output
			switch levelType {
			case "info":
				pterm.Info.Println(formatted)
			case "warning":
				pterm.Warning.Println(formatted)
			case "error":
				pterm.Error.Println(formatted)
			case "debug":
				pterm.Debug.Println(formatted)
			}
		}
	} else {
		l.mu.Unlock()
	}

	// Broadcast just this individual log entry to WebSocket clients
	// Full log history is sent on page load via /__cem/logs endpoint

	if ws != nil {
		// Send only the new log entry (not the entire history)
		msg := LogMessage{
			Type: "logs",
			Logs: []LogEntry{logEntry},
		}
		if msgBytes, err := json.Marshal(msg); err == nil {
			// Broadcast error intentionally ignored - failures occur when clients
			// disconnect and we can't log them here without causing infinite recursion
			_ = ws.Broadcast(msgBytes)
		}
	}
}

func (l *ptermLogger) Logs() []LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()
	// Return a copy to avoid race conditions
	logsCopy := make([]LogEntry, len(l.logs))
	copy(logsCopy, l.logs)
	return logsCopy
}

func (l *ptermLogger) Info(msg string, args ...any) {
	l.log("INFO", "info", msg, args...)
}

func (l *ptermLogger) Warning(msg string, args ...any) {
	l.log("WARN", "warning", msg, args...)
}

func (l *ptermLogger) Error(msg string, args ...any) {
	l.log("ERROR", "error", msg, args...)
}

func (l *ptermLogger) Debug(msg string, args ...any) {
	l.log("DEBUG", "debug", msg, args...)
}

// Clear clears all logs from the buffer
func (l *ptermLogger) Clear() {
	l.mu.Lock()
	l.logs = make([]LogEntry, 0)
	l.terminalLogs = make([]string, 0)
	l.mu.Unlock()
	if l.interactive {
		l.render()
	}
}

