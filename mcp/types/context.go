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
	"bennypowers.dev/cem/lsp/types"
	"bennypowers.dev/cem/manifest"
)

// MCPContext interface for accessing custom elements manifest data
// This allows tools to access context without circular dependency
type MCPContext interface {
	ElementInfo(tagName string) (ElementInfo, error)
	AllElements() map[string]ElementInfo
	LoadManifests() error
	GetManifestSchemaVersions() []string
	DocumentManager() types.DocumentManager

	// Lazy-computed cached methods for performance
	CommonPrefixes() []string
	AllCSSProperties() []string
}

// ElementInfo represents element information using manifest types directly
type ElementInfo interface {
	// Core manifest data
	Declaration() *manifest.CustomElementDeclaration

	// Convenience accessors that delegate to manifest types
	TagName() string
	Name() string
	Summary() string
	Description() string
	Module() string
	Package() string

	// Member accessors returning manifest types
	Attributes() []manifest.Attribute
	Slots() []manifest.Slot
	Events() []manifest.Event
	CssProperties() []manifest.CssCustomProperty
	CssParts() []manifest.CssPart
	CssStates() []manifest.CssCustomState


	// MCP-specific extensions
	Guidelines() []string
	Examples() []Example
}

// Legacy type aliases - use manifest types directly instead
type Attribute = manifest.Attribute
type Slot = manifest.Slot
type Event = manifest.Event
type CssProperty = manifest.CssCustomProperty
type CssPart = manifest.CssPart
type CssState = manifest.CssCustomState

type Example interface {
	Title() string
	Description() string
	Code() string
	Language() string
}

// ExampleInfo represents usage examples
type ExampleInfo struct {
	TitleValue       string `json:"title"`
	DescriptionValue string `json:"description,omitempty"`
	CodeValue        string `json:"code"`
	LanguageValue    string `json:"language,omitempty"`
}

func (e ExampleInfo) Title() string       { return e.TitleValue }
func (e ExampleInfo) Description() string { return e.DescriptionValue }
func (e ExampleInfo) Code() string        { return e.CodeValue }
func (e ExampleInfo) Language() string    { return e.LanguageValue }
