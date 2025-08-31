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
	"bennypowers.dev/cem/internal/logging"
	"github.com/tliron/glsp"
)

// Debug logging state - shared across all textDocument methods
// (Currently no shared state needed)

// SetGlobalLoggerContext sets the LSP context for the centralized logger
func SetGlobalLoggerContext(context *glsp.Context) {
	logging.SetLSPContext(context)
}

// SetDebugLoggingEnabled controls whether debug logging is enabled
// This is controlled by the LSP $/setTrace notification
func SetDebugLoggingEnabled(enabled bool) {
	logging.SetDebugEnabled(enabled)
}

// IsDebugLoggingEnabled returns whether debug logging is currently enabled
func IsDebugLoggingEnabled() bool {
	return logging.IsDebugEnabled()
}

// SafeDebugLog safely calls the centralized logger for debug messages
func SafeDebugLog(format string, args ...any) {
	logging.Debug(format, args...)
}
