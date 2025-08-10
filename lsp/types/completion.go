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

// CompletionContextType represents what the user is trying to complete
type CompletionContextType int

const (
	CompletionUnknown             CompletionContextType = iota
	CompletionTagName                                   // After < or inside tag name
	CompletionAttributeName                             // Inside a tag, completing attribute
	CompletionAttributeValue                            // Inside attribute value quotes
	CompletionLitEventBinding                           // @event-name (lit event binding)
	CompletionLitPropertyBinding                        // .property (lit property binding)
	CompletionLitBooleanAttribute                       // ?attribute (lit boolean attribute)
)

// CompletionAnalysis holds the analysis of cursor position for completion
type CompletionAnalysis struct {
	Type          CompletionContextType
	TagName       string // For attribute completion, which tag we're in
	AttributeName string // For value completion, which attribute
	TriggerChar   string // What character triggered completion
	LineContent   string // Content of the current line
	IsLitTemplate bool   // True if we're in a tagged template literal (not innerHTML)
	LitSyntax     string // The Lit syntax prefix: "@", ".", or "?"
}
