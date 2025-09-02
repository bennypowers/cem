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
package logging

import (
	"fmt"
	"os"
	"sync"

	"github.com/pterm/pterm"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// init configures pterm styles to use foreground colors only (no backgrounds)
// This creates cleaner, more readable output similar to pterm logger examples
func init() {
	// Modify existing printers to use foreground colors only, no backgrounds
	// Preserve original functionality while changing styling
	
	pterm.Info = *pterm.Info.WithPrefix(pterm.Prefix{
		Text:  "INFO",
		Style: pterm.NewStyle(pterm.FgBlue),
	}).WithMessageStyle(&pterm.ThemeDefault.DefaultText)
	
	pterm.Success = *pterm.Success.WithPrefix(pterm.Prefix{
		Text:  "SUCCESS", 
		Style: pterm.NewStyle(pterm.FgGreen),
	}).WithMessageStyle(&pterm.ThemeDefault.DefaultText)
	
	pterm.Warning = *pterm.Warning.WithPrefix(pterm.Prefix{
		Text:  "WARNING",
		Style: pterm.NewStyle(pterm.FgYellow),
	}).WithMessageStyle(&pterm.ThemeDefault.DefaultText)
	
	pterm.Error = *pterm.Error.WithPrefix(pterm.Prefix{
		Text:  "ERROR",
		Style: pterm.NewStyle(pterm.FgRed),
	}).WithMessageStyle(&pterm.ThemeDefault.DefaultText)
	
	pterm.Debug = *pterm.Debug.WithPrefix(pterm.Prefix{
		Text:  "DEBUG",
		Style: pterm.NewStyle(pterm.FgCyan),
	}).WithMessageStyle(&pterm.ThemeDefault.DefaultText)
}

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
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
	mu           sync.RWMutex
	mode         LoggerMode
	lspContext   *glsp.Context
	debugEnabled bool
	quietEnabled bool
}

// LoggerMode determines how logs are output
type LoggerMode int

const (
	// ModeCLI uses pterm for colorized CLI output
	ModeCLI LoggerMode = iota
	// ModeLSP uses LSP protocol messages (window/showMessage, window/logMessage)
	ModeLSP
)

// Global logger instance
var globalLogger = &Logger{
	mode:         ModeCLI, // Default to CLI mode
	debugEnabled: false,
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

// SetDebugEnabled controls whether debug messages are shown
func (l *Logger) SetDebugEnabled(enabled bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.debugEnabled = enabled
}

// IsDebugEnabled returns whether debug logging is enabled
func (l *Logger) IsDebugEnabled() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.debugEnabled
}

// SetQuietEnabled controls whether quiet mode is active (suppresses INFO and DEBUG)
func (l *Logger) SetQuietEnabled(enabled bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.quietEnabled = enabled
}

// IsQuietEnabled returns whether quiet mode is active
func (l *Logger) IsQuietEnabled() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.quietEnabled
}

// Debug logs a debug message (only shown if debug is enabled)
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
		// Use pterm Error for CLI (same as Error)
		pterm.Error.Println(message)
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
		// Use pterm Info for CLI
		pterm.Info.Println(message)
	case ModeLSP:
		if lspContext != nil {
			// Use window/showMessage for Info-level popup notifications
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
		// For CLI, just show the message and URLs
		pterm.Info.Println(message)
		for _, action := range actions {
			if action.URL != "" {
				pterm.Info.Printf("  %s: %s\n", action.Title, action.URL)
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

// Success logs a success message (treated as Info in LSP mode)
func (l *Logger) Success(format string, args ...any) {
	l.mu.RLock()
	mode := l.mode
	quietEnabled := l.quietEnabled
	l.mu.RUnlock()

	// Skip success messages if quiet mode is enabled (success is above warning)
	if quietEnabled {
		return
	}

	if mode == ModeCLI {
		// Use pterm Success for CLI
		pterm.Success.Printf(format+"\n", args...)
	} else {
		// Treat as Info for LSP
		l.log(LogLevelInfo, format, args...)
	}
}

// log is the internal logging implementation
func (l *Logger) log(level LogLevel, format string, args ...any) {
	l.mu.RLock()
	mode := l.mode
	lspContext := l.lspContext
	debugEnabled := l.debugEnabled
	quietEnabled := l.quietEnabled
	l.mu.RUnlock()

	// Skip debug messages if debug is not enabled
	if level == LogLevelDebug && !debugEnabled {
		return
	}
	
	// Skip INFO and DEBUG messages if quiet mode is enabled
	if quietEnabled && (level == LogLevelInfo || level == LogLevelDebug) {
		return
	}

	message := fmt.Sprintf(format, args...)

	switch mode {
	case ModeCLI:
		l.logCLI(level, message)
	case ModeLSP:
		l.logLSP(level, message, lspContext)
	}
}

// logCLI handles CLI-mode logging using pterm
func (l *Logger) logCLI(level LogLevel, message string) {
	switch level {
	case LogLevelDebug:
		pterm.Debug.Println(message)
	case LogLevelInfo:
		pterm.Info.Println(message)
	case LogLevelWarning:
		pterm.Warning.Println(message)
	case LogLevelError:
		pterm.Error.Println(message)
	}
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
	case LogLevelDebug:
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
