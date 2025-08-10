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
package helpers

import (
	"fmt"
	"sync"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Logger provides LSP-aware logging capabilities
type Logger struct {
	enabled bool
	context *glsp.Context
	mu      sync.RWMutex
}

// NewLogger creates a new LSP logger
func NewLogger() *Logger {
	return &Logger{
		enabled: false,
		context: nil,
	}
}

// Enable turns on logging with the given LSP context
func (l *Logger) Enable(context *glsp.Context) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.enabled = true
	l.context = context
}

// Disable turns off logging
func (l *Logger) Disable() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.enabled = false
	l.context = nil
}

// IsEnabled returns whether logging is currently enabled
func (l *Logger) IsEnabled() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.enabled
}

// SetEnabled sets the logging state
func (l *Logger) SetEnabled(enabled bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.enabled = enabled
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(protocol.MessageTypeInfo, format, args...)
}

// Warning logs a warning message
func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(protocol.MessageTypeWarning, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(protocol.MessageTypeError, format, args...)
}

// Log logs a debug message (same as Info for LSP)
func (l *Logger) Log(format string, args ...interface{}) {
	l.log(protocol.MessageTypeLog, format, args...)
}

// log sends a log message through the LSP protocol
func (l *Logger) log(messageType protocol.MessageType, format string, args ...interface{}) {
	l.mu.RLock()
	enabled := l.enabled
	context := l.context
	l.mu.RUnlock()

	if !enabled || context == nil {
		return
	}

	message := fmt.Sprintf(format, args...)

	// Send log message to LSP client
	go func() {
		context.Notify(protocol.ServerWindowLogMessage, &protocol.LogMessageParams{
			Type:    messageType,
			Message: message,
		})
	}()
}

// Global logger instance
var globalLogger = NewLogger()

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	return globalLogger
}

// SetGlobalLoggerContext sets the context for the global logger
func SetGlobalLoggerContext(context *glsp.Context) {
	globalLogger.Enable(context)
}

// Debug logging state and function - shared across all textDocument methods
var (
	debugLog        func(format string, args ...any)
	debugLogEnabled bool = false // Default to disabled, enabled via $/setTrace
	debugLogMutex   sync.RWMutex
)

// SetDebugLogger sets the debug logging function for all textDocument methods
func SetDebugLogger(fn func(format string, args ...any)) {
	debugLogMutex.Lock()
	defer debugLogMutex.Unlock()
	debugLog = fn
}

// SetDebugLoggingEnabled controls whether debug logging is enabled
// This is controlled by the LSP $/setTrace notification
func SetDebugLoggingEnabled(enabled bool) {
	debugLogMutex.Lock()
	defer debugLogMutex.Unlock()
	debugLogEnabled = enabled
}

// IsDebugLoggingEnabled returns whether debug logging is currently enabled
func IsDebugLoggingEnabled() bool {
	debugLogMutex.RLock()
	defer debugLogMutex.RUnlock()
	return debugLogEnabled
}

// SafeDebugLog safely calls debugLog if it's not nil and debug logging is enabled
func SafeDebugLog(format string, args ...any) {
	debugLogMutex.RLock()
	enabled := debugLogEnabled
	logFn := debugLog
	debugLogMutex.RUnlock()
	
	if enabled && logFn != nil {
		logFn(format, args...)
	}
}