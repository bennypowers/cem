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
package types

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Document interface for LSP document operations
type Document interface {
	FindElementAtPosition(position protocol.Position, dm any) *CustomElementMatch
	FindAttributeAtPosition(position protocol.Position, dm any) (*AttributeMatch, string)
	Content() string
	Version() int32
	URI() string
	FindCustomElements(dm any) ([]CustomElementMatch, error)
	AnalyzeCompletionContextTS(position protocol.Position, dm any) *CompletionAnalysis
}

// CustomElementMatch represents a found custom element
type CustomElementMatch struct {
	TagName    string
	Range      protocol.Range
	Attributes map[string]AttributeMatch
}

// AttributeMatch represents a found attribute
type AttributeMatch struct {
	Name  string
	Value string
	Range protocol.Range
}

// ElementDefinition represents a custom element with its source information
type ElementDefinition interface {
	ModulePath() string
	SourceHref() string
}
