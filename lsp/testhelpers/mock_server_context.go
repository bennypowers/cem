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
	"sync"

	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// MockServerContext provides a unified mock implementation of ServerContext for all tests
type MockServerContext struct {
	mu               sync.RWMutex
	Documents        map[string]types.Document
	TagNames         []string
	Elements         map[string]*M.CustomElement
	AttributesMap    map[string]map[string]*M.Attribute
	SlotsMap         map[string][]M.Slot
	ElementDefsMap   map[string]types.ElementDefinition
	DescriptionsMap  map[string]string
	WorkspaceRootStr string
	DocumentMgr      types.DocumentManager
	QueryMgr         *queries.QueryManager
	ModuleGraphInst  *types.ModuleGraph
	types.Registry
}

// MockElementDefinition implements types.ElementDefinition for testing
type MockElementDefinition struct {
	ModulePathStr  string
	PackageNameStr string
	SourceHrefStr  string
	ElementPtr     *M.CustomElement
}

func (m *MockElementDefinition) ModulePath() string {
	return m.ModulePathStr
}

func (m *MockElementDefinition) PackageName() string {
	return m.PackageNameStr
}

func (m *MockElementDefinition) SourceHref() string {
	return m.SourceHrefStr
}

func (m *MockElementDefinition) Element() *M.CustomElement {
	return m.ElementPtr
}

// MockManifestResolver implements types.ManifestResolver for testing
type MockManifestResolver struct {
	ctx *MockServerContext
}

func (r *MockManifestResolver) FindManifestModulesForImportPath(importPath string) []string {
	if r.ctx == nil {
		return nil
	}

	r.ctx.mu.RLock()
	defer r.ctx.mu.RUnlock()

	var matchingModules []string
	// Search through element definitions to find matches
	for _, elementDef := range r.ctx.ElementDefsMap {
		modulePath := elementDef.ModulePath()
		if modulePath == "" {
			continue
		}

		// Simple path matching for test
		if importPath == "@rhds/elements/rh-tabs/rh-tabs.js" && modulePath == "rh-tabs/rh-tabs.js" {
			matchingModules = append(matchingModules, modulePath)
		}
	}
	return matchingModules
}

func (r *MockManifestResolver) GetManifestModulePath(filePath string) string {
	if r.ctx == nil {
		return ""
	}

	// For test fixtures, map TypeScript files to their manifest module paths
	if filePath == "rh-tabs/rh-tabs.ts" {
		return "rh-tabs/rh-tabs.js"
	}
	if filePath == "rh-tabs/rh-tab.ts" {
		return "rh-tabs/rh-tab.js"
	}
	if filePath == "rh-tabs/rh-tab-panel.ts" {
		return "rh-tabs/rh-tab-panel.js"
	}
	return ""
}

func (r *MockManifestResolver) GetElementsFromManifestModule(manifestModulePath string) []string {
	if r.ctx == nil {
		return nil
	}

	r.ctx.mu.RLock()
	defer r.ctx.mu.RUnlock()

	var elements []string
	// Find elements that belong to this manifest module
	for tagName, elementDef := range r.ctx.ElementDefsMap {
		if elementDef.ModulePath() == manifestModulePath {
			elements = append(elements, tagName)
		}
	}
	return elements
}

// Verify MockServerContext implements ServerContext
var _ types.ServerContext = (*MockServerContext)(nil)
var _ types.ElementDefinition = (*MockElementDefinition)(nil)
var _ types.ManifestResolver = (*MockManifestResolver)(nil)

// NewMockServerContext creates a new mock server context
func NewMockServerContext() *MockServerContext {
	return &MockServerContext{
		Documents:        make(map[string]types.Document),
		TagNames:         []string{},
		Elements:         make(map[string]*M.CustomElement),
		AttributesMap:    make(map[string]map[string]*M.Attribute),
		SlotsMap:         make(map[string][]M.Slot),
		ElementDefsMap:   make(map[string]types.ElementDefinition),
		DescriptionsMap:  make(map[string]string),
		WorkspaceRootStr: "/test/workspace",
	}
}

// Server lifecycle methods
func (m *MockServerContext) InitializeManifests() error {
	return nil
}

func (m *MockServerContext) UpdateWorkspaceFromLSP(rootURI *string, workspaceFolders []protocol.WorkspaceFolder) error {
	if rootURI != nil {
		m.WorkspaceRootStr = *rootURI
	}
	return nil
}

// Document operations
func (m *MockServerContext) DocumentManager() (types.DocumentManager, error) {
	if m.DocumentMgr != nil {
		return m.DocumentMgr, nil
	}
	return nil, fmt.Errorf("DocumentManager not available in test context")
}

func (m *MockServerContext) Document(uri string) types.Document {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Documents[uri]
}

func (m *MockServerContext) AllDocuments() []types.Document {
	m.mu.RLock()
	defer m.mu.RUnlock()
	docs := make([]types.Document, 0, len(m.Documents))
	for _, doc := range m.Documents {
		docs = append(docs, doc)
	}
	return docs
}

// Workspace operations
func (m *MockServerContext) Workspace() types.Workspace {
	return &MockWorkspace{root: m.WorkspaceRootStr}
}

func (m *MockServerContext) WorkspaceRoot() string {
	return m.WorkspaceRootStr
}

// Logging
func (m *MockServerContext) DebugLog(format string, args ...any) {
	// Silent for tests
}

// Registry operations (embedded Registry interface)
func (m *MockServerContext) AllTagNames() []string {
	if m.Registry != nil {
		return m.Registry.AllTagNames()
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.TagNames
}

func (m *MockServerContext) Element(tagName string) (*M.CustomElement, bool) {
	if m.Registry != nil {
		return m.Registry.Element(tagName)
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	element, exists := m.Elements[tagName]
	return element, exists
}

func (m *MockServerContext) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	if m.Registry != nil {
		return m.Registry.Attributes(tagName)
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	attrs, exists := m.AttributesMap[tagName]
	return attrs, exists
}

func (m *MockServerContext) Slots(tagName string) ([]M.Slot, bool) {
	if m.Registry != nil {
		return m.Registry.Slots(tagName)
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	slots, exists := m.SlotsMap[tagName]
	return slots, exists
}

func (m *MockServerContext) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	if m.Registry != nil {
		return m.Registry.ElementDefinition(tagName)
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	def, exists := m.ElementDefsMap[tagName]
	return def, exists
}

func (m *MockServerContext) AddManifest(manifest *M.Package) {
	if m.Registry != nil {
		m.Registry.AddManifest(manifest)
		return
	}

	// If no Registry is set, populate the internal maps directly
	for _, module := range manifest.Modules {
		for _, decl := range module.Declarations {
			if customElement, ok := decl.(*M.CustomElementDeclaration); ok {
				element := &customElement.CustomElement
				tagName := element.TagName

				// Add to tag names if not already present
				if !slices.Contains(m.TagNames, tagName) {
					m.TagNames = append(m.TagNames, tagName)
				}

				// Add element
				m.Elements[tagName] = element

				// Add attributes
				if len(element.Attributes) > 0 {
					attrMap := make(map[string]*M.Attribute)
					for i := range element.Attributes {
						attr := &element.Attributes[i]
						attrMap[attr.Name] = attr
					}
					m.AttributesMap[tagName] = attrMap
				}

				// Add slots
				if len(element.Slots) > 0 {
					m.SlotsMap[tagName] = element.Slots
				}

				// Add element definition
				elementDef := &MockElementDefinition{
					ModulePathStr:  module.Path,
					PackageNameStr: "", // Could be extracted from package.json if needed
					SourceHrefStr:  "", // Could be computed if needed
					ElementPtr:     element,
				}
				m.ElementDefsMap[tagName] = elementDef
			}
		}
	}
}

// Element operations for advanced features
func (m *MockServerContext) ElementSource(tagName string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if def, exists := m.ElementDefsMap[tagName]; exists {
		return def.ModulePath(), true
	}
	return "", false
}

func (m *MockServerContext) ElementDescription(tagName string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	desc, exists := m.DescriptionsMap[tagName]
	return desc, exists
}

// Query operations for tree-sitter
func (m *MockServerContext) QueryManager() (*queries.QueryManager, error) {
	if m.QueryMgr != nil {
		return m.QueryMgr, nil
	}
	return nil, fmt.Errorf("QueryManager not available in test context")
}

func (m *MockServerContext) ModuleGraph() *types.ModuleGraph {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Return the persistent module graph instance
	if m.ModuleGraphInst == nil {
		// Initialize with a mock manifest resolver and QueryManager
		m.ModuleGraphInst = types.NewModuleGraph(m.QueryMgr)
		// Set workspace root if available
		if m.WorkspaceRootStr != "" {
			m.ModuleGraphInst.SetWorkspaceRoot(m.WorkspaceRootStr)
		}
		// Create a mock manifest resolver that uses this context
		mockResolver := &MockManifestResolver{ctx: m}
		m.ModuleGraphInst.SetManifestResolver(mockResolver)
	}
	return m.ModuleGraphInst
}

// Helper methods for setting up test data

// AddDocument adds a document to the mock context
func (m *MockServerContext) AddDocument(uri string, doc types.Document) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Documents[uri] = doc
}

// AddElement adds a custom element to the mock context
func (m *MockServerContext) AddElement(tagName string, element *M.CustomElement) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Elements[tagName] = element
	// Also add to tag names if not already present
	if !slices.Contains(m.TagNames, tagName) {
		m.TagNames = append(m.TagNames, tagName)
	}
}

// AddAttributes adds attributes for a tag name
func (m *MockServerContext) AddAttributes(tagName string, attrs map[string]*M.Attribute) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.AttributesMap[tagName] = attrs
}

// AddSlots adds slots for a tag name
func (m *MockServerContext) AddSlots(tagName string, slots []M.Slot) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.SlotsMap[tagName] = slots
}

// AddElementDefinition adds an element definition
func (m *MockServerContext) AddElementDefinition(tagName string, def types.ElementDefinition) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ElementDefsMap[tagName] = def
}

// SetWorkspaceRoot sets the workspace root for tests
func (m *MockServerContext) SetWorkspaceRoot(root string) {
	m.WorkspaceRootStr = root
}

// SetDocumentManager sets the document manager for tests
func (m *MockServerContext) SetDocumentManager(dm types.DocumentManager) {
	m.DocumentMgr = dm
}

// SetDocumentManager sets the document manager for tests
func (m *MockServerContext) SetRegistry(registry types.Registry) {
	m.Registry = registry
}

// SetQueryManager sets the query manager for tests
func (m *MockServerContext) SetQueryManager(qm *queries.QueryManager) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.QueryMgr = qm
	// Reset the module graph instance so it gets recreated with the new QueryManager
	m.ModuleGraphInst = nil
}

// AddElementDescription adds a description for an element
func (m *MockServerContext) AddElementDescription(tagName string, description string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.DescriptionsMap[tagName] = description
}

// MockWorkspace provides a simple workspace implementation for tests
type MockWorkspace struct {
	root string
}

func (w *MockWorkspace) Root() string {
	return w.root
}

func (w *MockWorkspace) Cleanup() error {
	return nil
}
