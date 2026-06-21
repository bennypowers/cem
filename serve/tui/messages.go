package tui

import "bennypowers.dev/cem/serve/logger"

// LogMsg carries a log entry from the server goroutine to the TUI.
type LogMsg struct {
	Entry logger.LogEntry
	Level string // display label: "INFO", "WARN", "ERROR", " OK ", "DEBUG", "TRACE"
}

// StatusMsg updates the status line footer.
type StatusMsg string

// ClearMsg clears the log viewport.
type ClearMsg struct{}

// ServerDoneMsg signals that the server goroutine has shut down.
type ServerDoneMsg struct{ Err error }

// RebuildResultMsg carries the result of an async manifest rebuild.
type RebuildResultMsg struct {
	Size int
	Err  error
}

// OpenBrowserResultMsg carries the result of an async browser open.
type OpenBrowserResultMsg struct{ Err error }

// VerbosityCycledMsg is sent after the verbosity level is cycled.
type VerbosityCycledMsg struct{ Label string }

// ServerReadyMsg signals that server initialization is complete.
type ServerReadyMsg struct {
	Port      int
	Reload    bool
	WatchDone <-chan struct{}
}

// ServerInitErrorMsg signals that server initialization failed.
type ServerInitErrorMsg struct{ Err error }

// ShutdownMsg is sent after the shutdown callback completes.
type ShutdownMsg struct{}
