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
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm"
	"golang.org/x/term"
)

// ptermLogger implements Logger interface using pterm live rendering
type ptermLogger struct {
	verbose       bool
	logs          []string // Plain text logs for web interface
	terminalLogs  []string // Colored logs for terminal display
	maxLogs       int
	maxTermLogs   int
	mu            sync.Mutex
	interactive   bool
	area          *pterm.AreaPrinter
	status        string
	wsManager     WebSocketManager
}

// NewPtermLogger creates a new pterm-based logger with live rendering
func NewPtermLogger(verbose bool) Logger {
	interactive := term.IsTerminal(int(os.Stdout.Fd()))

	logger := &ptermLogger{
		verbose:      verbose,
		logs:         make([]string, 0),
		terminalLogs: make([]string, 0),
		maxLogs:      100,
		maxTermLogs:  50, // Keep last 50 logs visible in terminal
		interactive:  interactive,
		status:       "Starting...",
	}

	if interactive {
		logger.area, _ = pterm.DefaultArea.Start()
		logger.render()
	}

	return logger
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
func (l *ptermLogger) SetWebSocketManager(wsManager WebSocketManager) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.wsManager = wsManager
}

// broadcastLogs sends current logs to all WebSocket clients
func (l *ptermLogger) broadcastLogs() {
	if l.wsManager == nil {
		return
	}

	msg := LogMessage{
		Type: "logs",
		Logs: l.logs,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return
	}

	l.wsManager.Broadcast(msgBytes)
}

// render updates the live terminal display
func (l *ptermLogger) render() {
	if !l.interactive || l.area == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var sb strings.Builder

	// Render logs with colors
	for _, log := range l.terminalLogs {
		sb.WriteString(log + "\n")
	}

	// Add separator
	sb.WriteString("\n" + strings.Repeat("─", 80) + "\n")

	// Status line at bottom (no trailing newline to avoid empty cursor line)
	sb.WriteString(pterm.FgLightGreen.Sprint("● ") + l.status)

	l.area.Update(sb.String())
}

// Stop stops the live rendering
func (l *ptermLogger) Stop() {
	if l.area != nil {
		l.area.Stop()
	}
}

func (l *ptermLogger) log(level, color, msg string, args ...interface{}) {
	formatted := fmt.Sprintf(msg, args...)
	timestamp := time.Now().Format("15:04:05")

	// Store plain text log with timestamp for web interface
	l.mu.Lock()
	plainLog := fmt.Sprintf("[%s] %s %s", timestamp, level, formatted)
	l.logs = append(l.logs, plainLog)
	if len(l.logs) > l.maxLogs {
		l.logs = l.logs[len(l.logs)-l.maxLogs:]
	}

	// Get copy of logs for broadcasting (before unlock)
	logsCopy := make([]string, len(l.logs))
	copy(logsCopy, l.logs)

	if l.interactive {
		// Create colored log line for terminal
		var prefix, coloredMsg string
		timestampStr := pterm.FgGray.Sprint(timestamp)

		switch color {
		case "info":
			prefix = pterm.FgCyan.Sprint("INFO ")
			coloredMsg = formatted
		case "error":
			prefix = pterm.FgRed.Sprint("ERROR")
			coloredMsg = pterm.FgRed.Sprint(formatted)
		case "debug":
			prefix = pterm.FgGray.Sprint("DEBUG")
			coloredMsg = pterm.FgGray.Sprint(formatted)
		}

		// Calculate padding for right-aligned timestamp
		// Use visual width (uncolored text length) not string length (which includes ANSI codes)
		width := 80
		if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 {
			width = w
		}

		// Visual length is just the actual text without ANSI codes
		visualLen := len(level) + 1 + len(formatted)
		padding := width - visualLen - 9 // 9 to prevent overflow (8 for timestamp + 1 buffer)
		if padding < 1 {
			padding = 1
		}

		terminalLog := fmt.Sprintf("%s %s%s%s", prefix, coloredMsg, strings.Repeat(" ", padding), timestampStr)
		l.terminalLogs = append(l.terminalLogs, terminalLog)

		if len(l.terminalLogs) > l.maxTermLogs {
			l.terminalLogs = l.terminalLogs[len(l.terminalLogs)-l.maxTermLogs:]
		}

		l.mu.Unlock()
		l.render()
	} else {
		l.mu.Unlock()
		// Non-interactive: standard pterm output
		if color == "info" {
			pterm.Info.Println(formatted)
		} else if color == "error" {
			pterm.Error.Println(formatted)
		} else if color == "debug" {
			pterm.Debug.Println(formatted)
		}
	}

	// Broadcast logs to WebSocket clients (after unlock to avoid blocking)
	if l.wsManager != nil {
		msg := LogMessage{
			Type: "logs",
			Logs: logsCopy,
		}
		if msgBytes, err := json.Marshal(msg); err == nil {
			l.wsManager.Broadcast(msgBytes)
		}
	}
}

func (l *ptermLogger) GetLogs() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	// Return a copy to avoid race conditions
	logsCopy := make([]string, len(l.logs))
	copy(logsCopy, l.logs)
	return logsCopy
}

func (l *ptermLogger) Info(msg string, args ...interface{}) {
	l.log("INFO", "info", msg, args...)
}

func (l *ptermLogger) Error(msg string, args ...interface{}) {
	l.log("ERROR", "error", msg, args...)
}

func (l *ptermLogger) Debug(msg string, args ...interface{}) {
	if l.verbose {
		l.log("DEBUG", "debug", msg, args...)
	}
}

// basicLogger is a simple implementation for testing
type basicLogger struct{}

func (l *basicLogger) Info(msg string, args ...interface{}) {
	fmt.Printf("[INFO] "+msg+"\n", args...)
}

func (l *basicLogger) Error(msg string, args ...interface{}) {
	fmt.Printf("[ERROR] "+msg+"\n", args...)
}

func (l *basicLogger) Debug(msg string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+msg+"\n", args...)
}
