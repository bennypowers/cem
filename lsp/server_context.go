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
	"path"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	serverMethods "bennypowers.dev/cem/lsp/methods/server"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// ServerAdapter adapts the Server to the various method contexts
type ServerAdapter struct {
	server *Server
}

// NewServerAdapter creates a new adapter for the server
func NewServerAdapter(server *Server) *ServerAdapter {
	return &ServerAdapter{server: server}
}

// Server Context implementations

func (s *ServerAdapter) DocumentManager() serverMethods.DocumentManager {
	return s.server.documents
}

// TextDocumentManager returns the document manager for text document methods
func (s *ServerAdapter) TextDocumentManager() textDocument.DocumentManager {
	return s.server.documents
}

func (s *ServerAdapter) Workspace() serverMethods.Workspace {
	return s.server.workspace
}

func (s *ServerAdapter) DebugLog(format string, args ...any) {
	helpers.SafeDebugLog(format, args...)
}

func (s *ServerAdapter) InitializeManifests() error {
	// Initialize the manifest registry
	if err := s.server.registry.LoadFromWorkspace(s.server.workspace); err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	// Start file watching for manifest changes
	if err := s.server.registry.StartFileWatching(s.server.handleManifestReload); err != nil {
		helpers.SafeDebugLog("Warning: Could not start file watching: %v", err)
		// Don't fail startup if file watching fails
	}

	// Start generate watcher for local project source file changes
	if err := s.server.registry.StartGenerateWatcher(); err != nil {
		helpers.SafeDebugLog("Warning: Could not start generate watcher: %v", err)
		// Don't fail startup if generate watcher fails
	}

	return nil
}

// Hover Context implementations

func (s *ServerAdapter) Document(uri string) types.Document {
	return s.server.documents.Document(uri)
}

func (s *ServerAdapter) Element(tagName string) (*M.CustomElement, bool) {
	return s.server.registry.Element(tagName)
}

func (s *ServerAdapter) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return s.server.registry.Attributes(tagName)
}

func (s *ServerAdapter) Slots(tagName string) ([]M.Slot, bool) {
	return s.server.registry.Slots(tagName)
}

// Definition Context implementations

func (s *ServerAdapter) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return s.server.registry.ElementDefinition(tagName)
}

func (s *ServerAdapter) ElementSource(tagName string) (string, bool) {
	// Get the element definition which contains source information
	definition, exists := s.server.registry.ElementDefinition(tagName)
	if !exists {
		helpers.SafeDebugLog("[ELEMENT_SOURCE] No definition found for tag '%s'", tagName)
		return "", false
	}

	// Prefer package name over module path for better import suggestions
	packageName := definition.PackageName()
	modulePath := definition.ModulePath()

	helpers.SafeDebugLog("[ELEMENT_SOURCE] Tag '%s': packageName='%s' (len=%d), modulePath='%s' (len=%d)",
		tagName, packageName, len(packageName), modulePath, len(modulePath))

	if packageName != "" && modulePath != "" {
		// For npm packages, combine package name with module path
		// e.g. "@rhds/elements" + "rh-card/rh-card.js" = "@rhds/elements/rh-card/rh-card.js"
		result := path.Join(packageName, modulePath)
		helpers.SafeDebugLog("[ELEMENT_SOURCE] Returning combined path: '%s'", result)
		return result, true
	} else if packageName != "" {
		// Just package name (fallback)
		helpers.SafeDebugLog("[ELEMENT_SOURCE] Returning package name only: '%s'", packageName)
		return packageName, true
	} else if modulePath != "" {
		// Fallback to module path for local/workspace elements
		helpers.SafeDebugLog("[ELEMENT_SOURCE] Returning module path only: '%s'", modulePath)
		return modulePath, true
	}

	helpers.SafeDebugLog("[ELEMENT_SOURCE] No source available for tag '%s'", tagName)
	return "", false
}

func (s *ServerAdapter) WorkspaceRoot() string {
	if s.server.workspace != nil {
		root := s.server.workspace.Root()
		helpers.SafeDebugLog("[WORKSPACE_ROOT] Returning workspace root: '%s' (len=%d)", root, len(root))
		return root
	}
	helpers.SafeDebugLog("[WORKSPACE_ROOT] No workspace context available")
	return ""
}

func (s *ServerAdapter) RawDocumentManager() interface{} {
	return s.server.documents
}

// UpdateWorkspaceFromLSP updates the workspace context based on LSP initialize parameters
func (s *ServerAdapter) UpdateWorkspaceFromLSP(rootURI *string, workspaceFolders []protocol.WorkspaceFolder) error {
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Starting workspace context update")
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Current workspace root before update: '%s'", s.server.workspace.Root())

	var newRoot string

	// Prefer workspace folders over rootURI
	if len(workspaceFolders) > 0 {
		// Use the first workspace folder URI
		folder := workspaceFolders[0]
		newRoot = folder.URI
		helpers.SafeDebugLog("[UPDATE_WORKSPACE] Using workspace folder URI: %s", newRoot)
	} else if rootURI != nil {
		newRoot = *rootURI
		helpers.SafeDebugLog("[UPDATE_WORKSPACE] Using root URI: %s", newRoot)
	} else {
		helpers.SafeDebugLog("[UPDATE_WORKSPACE] No workspace information provided by LSP client")
		return nil
	}

	// Convert file:// URI to file system path
	if strings.HasPrefix(newRoot, "file://") {
		newRoot = strings.TrimPrefix(newRoot, "file://")
		helpers.SafeDebugLog("[SERVER_ADAPTER] Converted URI to path: %s", newRoot)
	}

	// Create new workspace context with the correct root
	newWorkspace := W.NewFileSystemWorkspaceContext(newRoot)
	if err := newWorkspace.Init(); err != nil {
		helpers.SafeDebugLog("[SERVER_ADAPTER] Failed to initialize new workspace context: %v", err)
		return fmt.Errorf("failed to initialize workspace context: %w", err)
	}

	// Update server's workspace
	oldRoot := s.server.workspace.Root()
	s.server.workspace = newWorkspace
	newActualRoot := s.server.workspace.Root()
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Successfully updated workspace context")
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Old root: '%s' -> New root: '%s'", oldRoot, newActualRoot)

	return nil
}

// Completion Context implementations

func (s *ServerAdapter) AllTagNames() []string {
	return s.server.registry.AllTagNames()
}

func (s *ServerAdapter) ElementDescription(tagName string) (string, bool) {
	// For now, use a simple approach to get basic element info
	// This can be enhanced later to pull from the full declaration data
	element, exists := s.server.registry.Element(tagName)
	if !exists {
		return "", false
	}

	// Check if we have basic element information that could serve as description
	// For now, return empty since CustomElement doesn't directly have description
	// but this establishes the interface for when more detailed data is available
	_ = element
	return "", false
}
