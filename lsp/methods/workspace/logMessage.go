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
	"bennypowers.dev/cem/lsp/helpers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// LogMessageContext provides the dependencies needed for logging functionality
type LogMessageContext interface {
	GetLogger() *helpers.Logger
}

// LogInfo logs an info message using the context logger
func LogInfo(ctx LogMessageContext, context *glsp.Context, format string, args ...interface{}) {
	if logger := ctx.GetLogger(); logger != nil {
		logger.Info(format, args...)
	}
}

// LogWarning logs a warning message using the context logger
func LogWarning(ctx LogMessageContext, context *glsp.Context, format string, args ...interface{}) {
	if logger := ctx.GetLogger(); logger != nil {
		logger.Warning(format, args...)
	}
}

// LogError logs an error message using the context logger
func LogError(ctx LogMessageContext, context *glsp.Context, format string, args ...interface{}) {
	if logger := ctx.GetLogger(); logger != nil {
		logger.Error(format, args...)
	}
}

// LogDebug logs a debug message using the context logger
func LogDebug(ctx LogMessageContext, context *glsp.Context, format string, args ...interface{}) {
	if logger := ctx.GetLogger(); logger != nil {
		logger.Log(format, args...)
	}
}

// ShowMessage sends a message to be displayed to the user
func ShowMessage(ctx LogMessageContext, context *glsp.Context, messageType protocol.MessageType, message string) {
	if context != nil {
		go func() {
			context.Notify(protocol.ServerWindowShowMessage, &protocol.ShowMessageParams{
				Type:    messageType,
				Message: message,
			})
		}()
	}
}
