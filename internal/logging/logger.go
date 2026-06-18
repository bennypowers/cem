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

// Package logging provides centralized diagnostic output for cem.
//
// All diagnostic messages (debug, info, warning, error, success) should go
// through this package rather than calling fmt directly. The logger adapts
// output to the active mode: CLI mode uses lipgloss for styled terminal
// output; LSP mode routes messages over the LSP protocol.
//
// Quiet (-q) suppresses info and debug. Verbose (-v) enables debug. These
// flags are respected automatically by all log functions.
//
// This package does NOT own terminal UI primitives (spinners, live areas,
// colored display formatting). Those stay at their callsites.
package logging

import (
	"fmt"
	"os"
	"sync"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/bennypowers/glsp"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

var (
	infoPrefix    = lipgloss.NewStyle().Foreground(lipgloss.Blue).Render(" INFO ")
	successPrefix = lipgloss.NewStyle().Foreground(lipgloss.Green).Render(" SUCCESS ")
	warningPrefix = lipgloss.NewStyle().Foreground(lipgloss.Yellow).Render(" WARNING ")
	errorPrefix   = lipgloss.NewStyle().Foreground(lipgloss.Red).Render(" ERROR ")
	debugPrefix   = lipgloss.NewStyle().Foreground(lipgloss.Cyan).Render(" DEBUG ")
)

// Verbosity controls how much diagnostic output is produced.
type Verbosity int

const (
	VerbosityQuiet   Verbosity = -1 // -q: warnings and errors only
	VerbosityNormal  Verbosity = 0  // default: + success
	VerbosityVerbose Verbosity = 1  // -v: + info
	VerbosityDebug   Verbosity = 2  // -vv: + debug (timings, config, cache)
	VerbosityTrace   Verbosity = 3  // -vvv: + trace (per-file, tree-sitter)
)

func (v Verbosity) String() string {
	switch v {
	case VerbosityQuiet:
		return "quiet"
	case VerbosityNormal:
		return "normal"
	case VerbosityVerbose:
		return "verbose"
	case VerbosityDebug:
		return "debug"
	case VerbosityTrace:
		return "trace"
	default:
		return "unknown"
	}
}

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	LogLevelTrace LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
	case LogLevelTrace:
		return "TRACE"
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger provides centralized logging that adapts to CLI vs LSP contexts
type Logger struct {
	mu         sync.RWMutex
	mode       LoggerMode
	lspContext *glsp.Context
	verbosity  Verbosity
}

// LoggerMode determines how logs are output
type LoggerMode int

const (
	// ModeCLI uses lipgloss-styled output to stderr
	ModeCLI LoggerMode = iota
	// ModeLSP uses LSP protocol messages (window/showMessage, window/logMessage)
	ModeLSP
	// ModeServe uses lipgloss-styled output like ModeCLI. Verbosity gating is
	// controlled by the shared Verbosity level, same as all other modes.
	ModeServe
)

// Global logger instance
var globalLogger = &Logger{
	mode:      ModeCLI,
	verbosity: VerbosityNormal,
}

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	return globalLogger
}

// SetMode configures the logger for CLI or LSP operation
func (l *Logger) SetMode(mode LoggerMode) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.mode = mode
}

// SetLSPContext sets the LSP context for LSP mode logging
func (l *Logger) SetLSPContext(context *glsp.Context) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lspContext = context
	l.mode = ModeLSP
}

// SetVerbosity sets the verbosity level.
func (l *Logger) SetVerbosity(v Verbosity) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.verbosity = v
}

// Verbosity returns the current verbosity level.
func (l *Logger) Verbosity() Verbosity {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.verbosity
}

// AtLevel reports whether a message at the given level would be logged.
func (l *Logger) AtLevel(level LogLevel) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.shouldLog(level)
}

func (l *Logger) shouldLog(level LogLevel) bool {
	switch level {
	case LogLevelTrace:
		return l.verbosity >= VerbosityTrace
	case LogLevelDebug:
		return l.verbosity >= VerbosityDebug
	case LogLevelInfo:
		return l.verbosity >= VerbosityVerbose
	case LogLevelWarning, LogLevelError:
		return true
	}
	return false
}

// SetDebugEnabled is a convenience shim: true sets VerbosityDebug, false
// drops back to VerbosityNormal (unless already quieter).
func (l *Logger) SetDebugEnabled(enabled bool) {
	l.mu.Lock()
	v := l.verbosity
	l.mu.Unlock()
	if enabled && v < VerbosityDebug {
		l.SetVerbosity(VerbosityDebug)
	} else if !enabled && v >= VerbosityDebug {
		l.SetVerbosity(VerbosityNormal)
	}
}

// IsDebugEnabled returns whether debug logging is enabled (verbosity >= Debug).
func (l *Logger) IsDebugEnabled() bool {
	return l.AtLevel(LogLevelDebug)
}

// SetQuietEnabled is a convenience shim: true sets VerbosityQuiet, false
// restores VerbosityNormal (unless already more verbose).
func (l *Logger) SetQuietEnabled(enabled bool) {
	l.mu.Lock()
	v := l.verbosity
	l.mu.Unlock()
	if enabled {
		l.SetVerbosity(VerbosityQuiet)
	} else if v == VerbosityQuiet {
		l.SetVerbosity(VerbosityNormal)
	}
}

// IsQuietEnabled returns whether quiet mode is active (verbosity == Quiet).
func (l *Logger) IsQuietEnabled() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.verbosity == VerbosityQuiet
}

// Trace logs a trace message (only shown at -vvv).
func (l *Logger) Trace(format string, args ...any) {
	l.log(LogLevelTrace, format, args...)
}

// Debug logs a debug message (only shown at -vv or higher).
func (l *Logger) Debug(format string, args ...any) {
	l.log(LogLevelDebug, format, args...)
}

// Info logs an informational message
func (l *Logger) Info(format string, args ...any) {
	l.log(LogLevelInfo, format, args...)
}

// Warning logs a warning message
func (l *Logger) Warning(format string, args ...any) {
	l.log(LogLevelWarning, format, args...)
}

// Error logs an error message (goes to log output, not popup)
func (l *Logger) Error(format string, args ...any) {
	l.log(LogLevelError, format, args...)
}

// Critical logs a critical error that shows a popup notification in LSP mode
func (l *Logger) Critical(format string, args ...any) {
	l.mu.RLock()
	mode := l.mode
	lspContext := l.lspContext
	l.mu.RUnlock()

	message := fmt.Sprintf(format, args...)

	switch mode {
	case ModeCLI:
		lipgloss.Fprintf(os.Stderr, "%s %s\n", errorPrefix, message) //nolint:errcheck
	case ModeLSP:
		if lspContext != nil {
			// Always use window/showMessage for critical errors (popup)
			go func() {
				lspContext.Notify(protocol.ServerWindowShowMessage, &protocol.ShowMessageParams{
					Type:    protocol.MessageTypeError,
					Message: message,
				})
			}()
		} else {
			// Fallback to stderr
			fmt.Fprintf(os.Stderr, "[CRITICAL] %s\n", message)
		}
	}
}

// Notify sends an Info-level message as a popup notification (window/showMessage)
// This is for user-facing notifications that should be prominently displayed
func (l *Logger) Notify(format string, args ...any) {
	l.mu.RLock()
	mode := l.mode
	lspContext := l.lspContext
	l.mu.RUnlock()

	message := fmt.Sprintf(format, args...)

	switch mode {
	case ModeCLI:
		lipgloss.Fprintf(os.Stderr, "%s %s\n", infoPrefix, message) //nolint:errcheck
	case ModeLSP:
		if lspContext != nil {
			go func() {
				lspContext.Notify(protocol.ServerWindowShowMessage, &protocol.ShowMessageParams{
					Type:    protocol.MessageTypeInfo,
					Message: message,
				})
			}()
		} else {
			// Fallback to stderr
			fmt.Fprintf(os.Stderr, "[NOTIFY] %s\n", message)
		}
	}
}

// NotifyWithActions sends an Info-level message with action buttons that can open URLs
// This uses window/showMessageRequest for interactive notifications
func (l *Logger) NotifyWithActions(message string, actions []MessageAction) {
	l.mu.RLock()
	mode := l.mode
	lspContext := l.lspContext
	l.mu.RUnlock()

	switch mode {
	case ModeCLI:
		lipgloss.Fprintf(os.Stderr, "%s %s\n", infoPrefix, message) //nolint:errcheck
		for _, action := range actions {
			if action.URL != "" {
				lipgloss.Fprintf(os.Stderr, "%s   %s: %s\n", infoPrefix, action.Title, action.URL) //nolint:errcheck
			}
		}
	case ModeLSP:
		if lspContext != nil {
			// Convert to LSP protocol format
			actionItems := make([]protocol.MessageActionItem, len(actions))
			for i, action := range actions {
				actionItems[i] = protocol.MessageActionItem{
					Title: action.Title,
				}
			}

			// Send the request and handle the response
			go func() {
				var selectedAction *protocol.MessageActionItem
				lspContext.Call(string(protocol.ServerWindowShowMessageRequest), &protocol.ShowMessageRequestParams{
					Type:    protocol.MessageTypeInfo,
					Message: message,
					Actions: actionItems,
				}, &selectedAction)

				// Handle action selection
				if selectedAction != nil {
					// Find the corresponding action and open its URL
					for _, action := range actions {
						if action.Title == selectedAction.Title && action.URL != "" {
							// Use window/showDocument to open the URL
							external := true
							var showDocResult *protocol.ShowDocumentResult
							lspContext.Call(string(protocol.ServerWindowShowDocument), &protocol.ShowDocumentParams{
								URI:      protocol.URI(action.URL),
								External: &external,
							}, &showDocResult)
							break
						}
					}
				}
			}()
		} else {
			// Fallback to stderr
			fmt.Fprintf(os.Stderr, "[NOTIFY] %s\n", message)
			for _, action := range actions {
				if action.URL != "" {
					fmt.Fprintf(os.Stderr, "  %s: %s\n", action.Title, action.URL)
				}
			}
		}
	}
}

// MessageAction represents an action that can be taken from a notification
type MessageAction struct {
	Title string
	URL   string
}

// Success logs a success message. In LSP mode, emits as Info-level
// window/logMessage bypassing the normal verbosity gate so LSP clients
// see success messages at default verbosity without polluting stdio.
func (l *Logger) Success(format string, args ...any) {
	l.mu.RLock()
	mode := l.mode
	verbosity := l.verbosity
	lspContext := l.lspContext
	l.mu.RUnlock()

	if verbosity < VerbosityNormal {
		return
	}

	switch mode {
	case ModeCLI, ModeServe:
		lipgloss.Fprintf(os.Stderr, "%s %s\n", successPrefix, fmt.Sprintf(format, args...)) //nolint:errcheck
	case ModeLSP:
		message := fmt.Sprintf(format, args...)
		l.logLSP(LogLevelInfo, message, lspContext)
	}
}

// log is the internal logging implementation
func (l *Logger) log(level LogLevel, format string, args ...any) {
	l.mu.RLock()
	mode := l.mode
	lspContext := l.lspContext
	ok := l.shouldLog(level)
	l.mu.RUnlock()

	if !ok {
		return
	}

	message := fmt.Sprintf(format, args...)

	switch mode {
	case ModeCLI, ModeServe:
		l.logCLI(level, message)
	case ModeLSP:
		l.logLSP(level, message, lspContext)
	}
}

func (l *Logger) logCLI(level LogLevel, message string) {
	var prefix string
	switch level {
	case LogLevelTrace, LogLevelDebug:
		prefix = debugPrefix
	case LogLevelInfo:
		prefix = infoPrefix
	case LogLevelWarning:
		prefix = warningPrefix
	case LogLevelError:
		prefix = errorPrefix
	}
	lipgloss.Fprintf(os.Stderr, "%s %s\n", prefix, message) //nolint:errcheck
}

// logLSP handles LSP-mode logging using LSP protocol messages
func (l *Logger) logLSP(level LogLevel, message string, context *glsp.Context) {
	if context == nil {
		// Fallback to stderr if no LSP context available
		fmt.Fprintf(os.Stderr, "[%s] %s\n", level.String(), message)
		return
	}

	// Map log levels to LSP message types
	var messageType protocol.MessageType
	switch level {
	case LogLevelTrace, LogLevelDebug:
		messageType = protocol.MessageTypeLog
	case LogLevelInfo:
		messageType = protocol.MessageTypeInfo
	case LogLevelWarning:
		messageType = protocol.MessageTypeWarning
	case LogLevelError:
		messageType = protocol.MessageTypeError
	}

	// Use window/logMessage for all standard log levels (non-intrusive)
	// Only Critical() method uses window/showMessage for popup notifications
	go func() {
		context.Notify(protocol.ServerWindowLogMessage, &protocol.LogMessageParams{
			Type:    messageType,
			Message: message,
		})
	}()
}

// Convenience functions for global logger

func Trace(format string, args ...any) {
	globalLogger.Trace(format, args...)
}

func Debug(format string, args ...any) {
	globalLogger.Debug(format, args...)
}

func Info(format string, args ...any) {
	globalLogger.Info(format, args...)
}

func Warning(format string, args ...any) {
	globalLogger.Warning(format, args...)
}

func Error(format string, args ...any) {
	globalLogger.Error(format, args...)
}

func Critical(format string, args ...any) {
	globalLogger.Critical(format, args...)
}

func Notify(format string, args ...any) {
	globalLogger.Notify(format, args...)
}

func NotifyWithActions(message string, actions []MessageAction) {
	globalLogger.NotifyWithActions(message, actions)
}

func Success(format string, args ...any) {
	globalLogger.Success(format, args...)
}

func SetMode(mode LoggerMode) {
	globalLogger.SetMode(mode)
}

func SetLSPContext(context *glsp.Context) {
	globalLogger.SetLSPContext(context)
}

func SetVerbosity(v Verbosity) {
	globalLogger.SetVerbosity(v)
}

func CurrentVerbosity() Verbosity {
	return globalLogger.Verbosity()
}

func AtLevel(level LogLevel) bool {
	return globalLogger.AtLevel(level)
}

func SetDebugEnabled(enabled bool) {
	globalLogger.SetDebugEnabled(enabled)
}

func IsDebugEnabled() bool {
	return globalLogger.IsDebugEnabled()
}

func SetQuietEnabled(enabled bool) {
	globalLogger.SetQuietEnabled(enabled)
}

func IsQuietEnabled() bool {
	return globalLogger.IsQuietEnabled()
}
