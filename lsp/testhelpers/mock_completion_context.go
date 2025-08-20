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
	"fmt"
	"slices"

	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
)

// MockCompletionContext provides a reusable mock implementation of CompletionContext for tests
type MockCompletionContext struct {
	Documents     map[string]types.Document
	TagNames      []string
	Elements      map[string]*M.CustomElement
	AttributesMap map[string]map[string]*M.Attribute
	SlotsMap      map[string][]M.Slot
}

// NewMockCompletionContext creates a new mock completion context
func NewMockCompletionContext() *MockCompletionContext {
	return &MockCompletionContext{
		Documents:     make(map[string]types.Document),
		TagNames:      []string{},
		Elements:      make(map[string]*M.CustomElement),
		AttributesMap: make(map[string]map[string]*M.Attribute),
		SlotsMap:      make(map[string][]M.Slot),
	}
}

// Document returns the document for the given URI
func (m *MockCompletionContext) Document(uri string) types.Document {
	return m.Documents[uri]
}

// AllTagNames returns all available tag names
func (m *MockCompletionContext) AllTagNames() []string {
	return m.TagNames
}

// Element returns the custom element for the given tag name
func (m *MockCompletionContext) Element(tagName string) (*M.CustomElement, bool) {
	element, exists := m.Elements[tagName]
	return element, exists
}

// Attributes returns the attributes for the given tag name
func (m *MockCompletionContext) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	attrs, exists := m.AttributesMap[tagName]
	return attrs, exists
}

// Slots returns the slots for the given tag name
func (m *MockCompletionContext) Slots(tagName string) ([]M.Slot, bool) {
	slots, exists := m.SlotsMap[tagName]
	return slots, exists
}

// DocumentManager returns an error since test contexts don't provide DocumentManager
func (m *MockCompletionContext) DocumentManager() (any, error) {
	return nil, fmt.Errorf("DocumentManager not available in test context")
}

// AddDocument adds a document to the mock context
func (m *MockCompletionContext) AddDocument(uri string, doc types.Document) {
	m.Documents[uri] = doc
}

// AddElement adds a custom element to the mock context
func (m *MockCompletionContext) AddElement(tagName string, element *M.CustomElement) {
	m.Elements[tagName] = element
	// Also add to tag names if not already present
	if slices.Contains(m.TagNames, tagName) {
		return
	}
	m.TagNames = append(m.TagNames, tagName)
}

// AddAttributes adds attributes for a tag name
func (m *MockCompletionContext) AddAttributes(tagName string, attrs map[string]*M.Attribute) {
	m.AttributesMap[tagName] = attrs
}

// AddSlots adds slots for a tag name
func (m *MockCompletionContext) AddSlots(tagName string, slots []M.Slot) {
	m.SlotsMap[tagName] = slots
}

