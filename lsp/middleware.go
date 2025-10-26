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
package lsp

import (
	"fmt"
	"runtime/debug"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
)

// method wraps an LSP handler that returns (result, error) with middleware
// Returns the underlying function type so it's compatible with protocol.Handler field types
func method[P, R any](
	s *Server,
	method string,
	handler func(types.ServerContext, *glsp.Context, P) (R, error),
) func(*glsp.Context, P) (R, error) {
	return func(ctx *glsp.Context, params P) (result R, err error) {
		// Panic recovery - prevents LSP server crashes
		defer func() {
			if r := recover(); r != nil {
				helpers.SafeDebugLog("[LSP] PANIC in %s: %v\nStack trace:\n%s",
					method, r, string(debug.Stack()))
				err = fmt.Errorf("internal error in %s", method)
				var zero R
				result = zero
			}
		}()

		// Request logging
		helpers.SafeDebugLog("[LSP] %s started", method)

		// Execute handler
		result, err = handler(s, ctx, params)

		// Error context wrapping
		if err != nil {
			helpers.SafeDebugLog("[LSP] %s error: %v", method, err)
			return result, fmt.Errorf("%s: %w", method, err)
		}

		// Success logging
		helpers.SafeDebugLog("[LSP] %s completed successfully", method)
		return result, nil
	}
}

// notify wraps an LSP notification handler that returns only error
func notify[P any](
	s *Server,
	method string,
	handler func(types.ServerContext, *glsp.Context, P) error,
) func(*glsp.Context, P) error {
	return func(ctx *glsp.Context, params P) (err error) {
		defer func() {
			if r := recover(); r != nil {
				helpers.SafeDebugLog("[LSP] PANIC in %s: %v\nStack trace:\n%s",
					method, r, string(debug.Stack()))
				err = fmt.Errorf("internal error in %s", method)
			}
		}()

		helpers.SafeDebugLog("[LSP] %s started", method)
		err = handler(s, ctx, params)

		if err != nil {
			helpers.SafeDebugLog("[LSP] %s error: %v", method, err)
			return fmt.Errorf("%s: %w", method, err)
		}

		helpers.SafeDebugLog("[LSP] %s completed successfully", method)
		return nil
	}
}

// noParam wraps an LSP handler that takes no params (like Shutdown)
func noParam(
	s *Server,
	method string,
	handler func(types.ServerContext, *glsp.Context) error,
) func(*glsp.Context) error {
	return func(ctx *glsp.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				helpers.SafeDebugLog("[LSP] PANIC in %s: %v\nStack trace:\n%s",
					method, r, string(debug.Stack()))
				err = fmt.Errorf("internal error in %s", method)
			}
		}()

		helpers.SafeDebugLog("[LSP] %s started", method)
		err = handler(s, ctx)

		if err != nil {
			helpers.SafeDebugLog("[LSP] %s error: %v", method, err)
			return fmt.Errorf("%s: %w", method, err)
		}

		helpers.SafeDebugLog("[LSP] %s completed successfully", method)
		return nil
	}
}
