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
	"bennypowers.dev/cem/internal/platform"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/modulegraph"
	"bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// DocumentManager interface for document operations
type DocumentManager interface {
	OpenDocument(uri, content string, version int32) Document
	UpdateDocument(uri, content string, version int32) Document
	UpdateDocumentWithChanges(uri, content string, version int32, changes []protocol.TextDocumentContentChangeEvent) Document
	CloseDocument(uri string)
	Document(uri string) Document
	AllDocuments() []Document
	Close()
}

// Workspace interface for workspace operations
type Workspace interface {
	Root() string
	Cleanup() error
}

// Registry interface for manifest operations
type Registry interface {
	AddManifest(manifest *M.Package)
	AllTagNames() []string
	Element(tagName string) (*M.CustomElement, bool)
	Attributes(tagName string) (map[string]*M.Attribute, bool)
	Slots(tagName string) ([]M.Slot, bool)
	ElementDefinition(tagName string) (ElementDefinition, bool)
	FindCustomElementDeclaration(tagName string) *M.CustomElementDeclaration
	ManifestCount() int
	ElementCount() int
}

// ServerContext provides all dependencies needed for LSP methods
// This unified context eliminates the need for method-specific context interfaces
type ServerContext interface {
	// Server lifecycle
	InitializeManifests() error
	UpdateWorkspaceFromLSP(rootURI *string, workspaceFolders []protocol.WorkspaceFolder) error
	SetAdditionalPackages(packages []string)
	Close() error

	// Document operations
	DocumentManager() (DocumentManager, error)
	Document(uri string) Document
	AllDocuments() []Document

	// Workspace operations
	Workspace() Workspace
	WorkspaceRoot() string

	// Filesystem operations
	FileSystem() platform.FileSystem

	// Logging
	DebugLog(format string, args ...any)

	// Registry operations (embedded Registry interface)
	Registry

	// Element operations for advanced features
	ElementSource(tagName string) (string, bool)
	ElementDescription(tagName string) (string, bool)

	// Query operations for tree-sitter
	QueryManager() (*queries.QueryManager, error)

	// Module graph operations for re-export analysis
	ModuleGraph() *modulegraph.ModuleGraph
}
