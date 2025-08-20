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
package textDocument

import (
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Document interface for textDocument operations - uses shared types
type Document interface {
	FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch
	FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string)
}

// Type aliases for backward compatibility
type CustomElementMatch = types.CustomElementMatch
type AttributeMatch = types.AttributeMatch

// ElementDefinition represents a custom element with its source information
type ElementDefinition interface {
	GetModulePath() string
	GetSourceHref() string
}
