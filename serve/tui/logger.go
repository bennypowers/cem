package tui

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"

	"bennypowers.dev/cem/internal/logging"
	itui "bennypowers.dev/cem/internal/tui"
	"bennypowers.dev/cem/serve/logger"
)

// Logger implements logger.Logger by sending tea.Msg to a bubbletea Program.
type Logger struct {
	program   *tea.Program
	mu        sync.Mutex
	logs      []logger.LogEntry
	maxLogs   int
	wsManager logger.Broadcaster
	pending   []tea.Msg
}

// NewLogger creates a logger that buffers messages until SetProgram is called.
func NewLogger() *Logger {
	return &Logger{
		logs:    make([]logger.LogEntry, 0, 100),
		maxLogs: 100,
	}
}

// SetProgram atomically sets the program and returns buffered messages.
// After this call, log() sends directly via p.Send().
func (l *Logger) SetProgram(p *tea.Program) []tea.Msg {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.program = p
	pending := l.pending
	l.pending = nil
	return pending
}

// SetWebSocketManager sets the WebSocket manager for broadcasting logs.
func (l *Logger) SetWebSocketManager(wsManager logger.Broadcaster) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.wsManager = wsManager
}

// Logs returns a copy of the log ring buffer.
func (l *Logger) Logs() []logger.LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()
	out := make([]logger.LogEntry, len(l.logs))
	copy(out, l.logs)
	return out
}

func (l *Logger) log(level, levelType, msg string, args ...any) {
	formatted := fmt.Sprintf(msg, args...)
	now := time.Now()

	entry := logger.LogEntry{
		Type:    levelType,
		Date:    now.Format(time.RFC3339),
		Message: formatted,
	}

	l.mu.Lock()
	l.logs = append(l.logs, entry)
	if len(l.logs) > l.maxLogs {
		l.logs = l.logs[len(l.logs)-l.maxLogs:]
	}
	p := l.program
	ws := l.wsManager
	shouldSend := logging.ShouldDisplay(levelType)
	if shouldSend && p == nil {
		l.pending = append(l.pending, LogMsg{Entry: entry, Level: level})
	}
	l.mu.Unlock()

	if shouldSend && p != nil {
		p.Send(LogMsg{Entry: entry, Level: level})
	}

	if ws != nil {
		wsEntry := entry
		wsEntry.Message = ansi.Strip(wsEntry.Message)
		wsMsg := logger.LogMessage{Type: "logs", Logs: []logger.LogEntry{wsEntry}}
		if data, err := json.Marshal(wsMsg); err == nil {
			_ = ws.Broadcast(data)
		}
	}
}


func (l *Logger) Info(msg string, args ...any)    { l.log("INFO", "info", msg, args...) }
func (l *Logger) Warning(msg string, args ...any) { l.log("WARN", "warning", msg, args...) }
func (l *Logger) Error(msg string, args ...any)   { l.log("ERROR", "error", msg, args...) }
func (l *Logger) Debug(msg string, args ...any)   { l.log("DEBUG", "debug", msg, args...) }
func (l *Logger) Success(msg string, args ...any) { l.log(" OK ", "success", msg, args...) }
func (l *Logger) Trace(msg string, args ...any)   { l.log("TRACE", "trace", msg, args...) }

// Clear clears the log buffer and tells the TUI to clear the viewport.
func (l *Logger) Clear() {
	l.mu.Lock()
	l.logs = l.logs[:0]
	p := l.program
	l.mu.Unlock()
	if p != nil {
		p.Send(ClearMsg{})
	}
}

// SetStatus sends a status update to the TUI.
func (l *Logger) SetStatus(status string) {
	l.mu.Lock()
	p := l.program
	if p == nil {
		l.pending = append(l.pending, StatusMsg(status))
	}
	l.mu.Unlock()
	if p != nil {
		p.Send(StatusMsg(status))
	}
}

// LogDurations sends structured duration data to both TUI and WebSocket.
func (l *Logger) LogDurations(title string, entries []itui.DurationData) {
	if !logging.ShouldDisplay("debug") {
		return
	}

	durations := make([]logger.DurationEntry, len(entries))
	for i, e := range entries {
		durations[i] = logger.DurationEntry{
			Name:     e.Name,
			Duration: e.Duration,
			Percent:  e.Percent,
		}
	}

	entry := logger.LogEntry{
		Type:    "debug",
		Date:    time.Now().Format(time.RFC3339),
		Message: title,
		Data: &logger.LogData{
			Kind:      "durations",
			Durations: durations,
		},
	}

	l.mu.Lock()
	p := l.program
	ws := l.wsManager
	logMsg := LogMsg{Entry: entry, Level: "DEBUG"}
	if p == nil {
		l.pending = append(l.pending, logMsg)
	}
	l.mu.Unlock()

	if p != nil {
		p.Send(logMsg)
	}

	if ws != nil {
		wsMsg := logger.LogMessage{Type: "logs", Logs: []logger.LogEntry{entry}}
		if data, err := json.Marshal(wsMsg); err == nil {
			_ = ws.Broadcast(data)
		}
	}
}

// Start is a no-op; bubbletea manages lifecycle.
func (l *Logger) Start() {}

// Stop is a no-op; bubbletea manages lifecycle.
func (l *Logger) Stop() {}
