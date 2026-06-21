package tui

import "bennypowers.dev/cem/internal/logging"

// ShouldDisplay reports whether a log entry of the given level type
// should be shown at the current verbosity.
func ShouldDisplay(levelType string) bool {
	switch levelType {
	case "trace":
		return logging.AtLevel(logging.LogLevelTrace)
	case "debug":
		return logging.AtLevel(logging.LogLevelDebug)
	case "info":
		return logging.AtLevel(logging.LogLevelInfo)
	case "success":
		return logging.CurrentVerbosity() >= logging.VerbosityNormal
	case "warning", "error":
		return true
	default:
		return false
	}
}
