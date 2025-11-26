// Package types provides shared interfaces for middleware packages
package types

// Logger provides structured logging for middleware
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warning(msg string, args ...any)
	Error(msg string, args ...any)
}

// ErrorBroadcaster sends error messages to connected WebSocket clients
type ErrorBroadcaster interface {
	BroadcastError(title, message, filename string)
}
