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

import "bennypowers.dev/cem/lsp/types"

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

// ElementInfo represents element information as needed by tools
type ElementInfo interface {
	TagName() string
	Name() string
	Description() string
	Module() string
	Package() string
	Attributes() []Attribute
	Slots() []Slot
	Events() []Event
	CssProperties() []CssProperty
	CssParts() []CssPart
	CssStates() []CssState
	Guidelines() []string
	Examples() []Example
	ItemsByKind(kind string) []Item
}

// Item interfaces for different element features
type Item interface {
	Name() string
	Description() string
	Guidelines() []string
	Examples() []string
	Kind() string
}

type Attribute interface {
	Item
	Type() string
	Default() string
	Required() bool
	Values() []string
}

type Slot interface {
	Item
}

type Event interface {
	Item
	Type() string
}

type CssProperty interface {
	Item
	Syntax() string
	Inherits() bool
	Initial() string
}

type CssPart interface {
	Item
}

type CssState interface {
	Item
}

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
