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

	helpers.SafeDebugLog("[ELEMENT_SOURCE] Tag '%s': packageName='%s', modulePath='%s'", tagName, packageName, modulePath)

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
		return s.server.workspace.Root()
	}
	return ""
}

func (s *ServerAdapter) RawDocumentManager() interface{} {
	return s.server.documents
}

// Completion Context implementations

func (s *ServerAdapter) AllTagNames() []string {
	return s.server.registry.AllTagNames()
}
