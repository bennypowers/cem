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
package testhelpers

import (
	"slices"

	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
)

// MockRegistry provides a unified mock implementation of Registry for all tests
type MockRegistry struct {
	TagNames       []string
	Elements       map[string]*M.CustomElement
	AttributesMap  map[string]map[string]*M.Attribute
	SlotsMap       map[string][]M.Slot
	ElementDefsMap map[string]types.ElementDefinition
}

// Verify MockRegistry implements Registry
var _ types.Registry = (*MockRegistry)(nil)

// NewMockRegistry creates a new mock registry
func NewMockRegistry() *MockRegistry {
	return &MockRegistry{
		TagNames:       []string{},
		Elements:       make(map[string]*M.CustomElement),
		AttributesMap:  make(map[string]map[string]*M.Attribute),
		SlotsMap:       make(map[string][]M.Slot),
		ElementDefsMap: make(map[string]types.ElementDefinition),
	}
}

// AllTagNames returns all available tag names
func (r *MockRegistry) AllTagNames() []string {
	return r.TagNames
}

// Element returns the custom element for the given tag name
func (r *MockRegistry) Element(tagName string) (*M.CustomElement, bool) {
	element, exists := r.Elements[tagName]
	return element, exists
}

// Attributes returns the attributes for the given tag name
func (r *MockRegistry) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	attrs, exists := r.AttributesMap[tagName]
	return attrs, exists
}

// Slots returns the slots for the given tag name
func (r *MockRegistry) Slots(tagName string) ([]M.Slot, bool) {
	slots, exists := r.SlotsMap[tagName]
	return slots, exists
}

// ElementDefinition returns the element definition for the given tag name
func (r *MockRegistry) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	def, exists := r.ElementDefsMap[tagName]
	return def, exists
}

// Helper methods for setting up test data

// AddElement adds a custom element to the mock registry
func (r *MockRegistry) AddElement(tagName string, element *M.CustomElement) {
	r.Elements[tagName] = element
	// Also add to tag names if not already present
	if !slices.Contains(r.TagNames, tagName) {
		r.TagNames = append(r.TagNames, tagName)
	}
}

// AddAttributes adds attributes for a tag name
func (r *MockRegistry) AddAttributes(tagName string, attrs map[string]*M.Attribute) {
	r.AttributesMap[tagName] = attrs
}

// AddSlots adds slots for a tag name
func (r *MockRegistry) AddSlots(tagName string, slots []M.Slot) {
	r.SlotsMap[tagName] = slots
}

// AddElementDefinition adds an element definition
func (r *MockRegistry) AddElementDefinition(tagName string, def types.ElementDefinition) {
	r.ElementDefsMap[tagName] = def
}

// MockElementDefinition implements types.ElementDefinition for testing
type MockElementDefinition struct {
	CustomElement  *M.CustomElement
	ModulePathStr  string
	PackageNameStr string
	SourceHrefStr  string
}

func (d *MockElementDefinition) ModulePath() string {
	return d.ModulePathStr
}

func (d *MockElementDefinition) PackageName() string {
	return d.PackageNameStr
}

func (d *MockElementDefinition) SourceHref() string {
	return d.SourceHrefStr
}

func (d *MockElementDefinition) Element() *M.CustomElement {
	return d.CustomElement
}

// Verify MockElementDefinition implements ElementDefinition
var _ types.ElementDefinition = (*MockElementDefinition)(nil)

// AddManifest adds a complete manifest to the mock registry
// This simulates the behavior of the real Registry.AddManifest method
func (r *MockRegistry) AddManifest(manifest *M.Package) {
	for _, module := range manifest.Modules {
		for _, declaration := range module.Declarations {
			if customElementDecl, ok := declaration.(*M.CustomElementDeclaration); ok {
				element := &customElementDecl.CustomElement
				tagName := element.TagName

				// Add the element
				r.AddElement(tagName, element)

				// Add attributes
				if len(element.Attributes) > 0 {
					attrs := make(map[string]*M.Attribute)
					for i := range element.Attributes {
						attr := &element.Attributes[i]
						attrs[attr.Name] = attr
					}
					r.AddAttributes(tagName, attrs)
				}

				// Add slots
				if len(element.Slots) > 0 {
					r.AddSlots(tagName, element.Slots)
				}

				// Add element definition
				elementDef := &MockElementDefinition{
					CustomElement:  element,
					ModulePathStr:  module.Path,
					PackageNameStr: "",
				}
				r.AddElementDefinition(tagName, elementDef)
			}
		}
	}
}
