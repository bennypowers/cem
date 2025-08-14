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

	"bennypowers.dev/cem/lsp/helpers"
	serverMethods "bennypowers.dev/cem/lsp/methods/server"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
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

func (s *ServerAdapter) GetDocumentManager() serverMethods.DocumentManager {
	return s.server.documents
}

// GetTextDocumentManager returns the document manager for text document methods
func (s *ServerAdapter) GetTextDocumentManager() textDocument.DocumentManager {
	return s.server.documents
}

func (s *ServerAdapter) GetWorkspace() serverMethods.Workspace {
	return s.server.workspace
}

func (s *ServerAdapter) DebugLog(format string, args ...interface{}) {
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

func (s *ServerAdapter) GetDocument(uri string) types.Document {
	return s.server.documents.GetDocument(uri)
}

func (s *ServerAdapter) GetElement(tagName string) (*M.CustomElement, bool) {
	return s.server.registry.GetElement(tagName)
}

func (s *ServerAdapter) GetAttributes(tagName string) (map[string]*M.Attribute, bool) {
	return s.server.registry.GetAttributes(tagName)
}

func (s *ServerAdapter) GetSlots(tagName string) ([]M.Slot, bool) {
	return s.server.registry.GetSlots(tagName)
}

// Definition Context implementations

func (s *ServerAdapter) GetElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return s.server.registry.GetElementDefinition(tagName)
}

func (s *ServerAdapter) GetElementSource(tagName string) (string, bool) {
	// Get the element definition which contains source information
	definition, exists := s.server.registry.GetElementDefinition(tagName)
	if !exists {
		return "", false
	}

	// Prefer package name over module path for better import suggestions
	packageName := definition.GetPackageName()
	if packageName != "" {
		// Return the package name for npm package imports
		return packageName, true
	}

	// Fallback to module path for local/workspace elements
	modulePath := definition.GetModulePath()
	if modulePath != "" {
		return modulePath, true
	}

	return "", false
}

func (s *ServerAdapter) GetWorkspaceRoot() string {
	if s.server.workspace != nil {
		return s.server.workspace.Root()
	}
	return ""
}

func (s *ServerAdapter) GetRawDocumentManager() interface{} {
	return s.server.documents
}

// Completion Context implementations

func (s *ServerAdapter) GetAllTagNames() []string {
	return s.server.registry.GetAllTagNames()
}
