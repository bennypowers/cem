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
package workspace

import (
	"bennypowers.dev/cem/internal/logging"
	"go.lsp.dev/protocol"
)

// LogInfo logs an info message using the centralized logger
func LogInfo(format string, args ...any) {
	logging.Info(format, args...)
}

// LogWarning logs a warning message using the centralized logger
func LogWarning(format string, args ...any) {
	logging.Warning(format, args...)
}

// LogError logs an error message using the centralized logger
func LogError(format string, args ...any) {
	logging.Error(format, args...)
}

// LogDebug logs a debug message using the centralized logger
func LogDebug(format string, args ...any) {
	logging.Debug(format, args...)
}

// ShowMessage sends a message to be displayed to the user
func ShowMessage(messageType protocol.MessageType, message string) {
	// TODO: wire up notification via the new protocol connection
	logging.Info("[SHOW_MESSAGE] type=%d: %s", messageType, message)
}
