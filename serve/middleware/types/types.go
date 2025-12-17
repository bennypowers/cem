// Package types provides shared interfaces and types for middleware packages
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

// ImportMapOverride represents import map overrides with structure matching importmap.ImportMap
// This is used for configuration to avoid circular dependencies while maintaining the same structure
type ImportMapOverride struct {
	Imports map[string]string            `mapstructure:"imports" yaml:"imports" json:"imports"`
	Scopes  map[string]map[string]string `mapstructure:"scopes" yaml:"scopes,omitempty" json:"scopes,omitempty"`
}

// ImportMapConfig holds import map configuration
// This is shared between cmd/config and serve packages to avoid duplication
type ImportMapConfig struct {
	Generate     bool              `mapstructure:"generate" yaml:"generate"`
	OverrideFile string            `mapstructure:"overrideFile" yaml:"overrideFile"`
	Override     ImportMapOverride `mapstructure:"override" yaml:"override"`
}
