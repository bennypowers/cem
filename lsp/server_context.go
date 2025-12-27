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

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/modulegraph"
	"bennypowers.dev/cem/queries"
	W "bennypowers.dev/cem/workspace"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Verify Server implements all context interfaces
var _ types.ServerContext = (*Server)(nil)

// Server Context implementations - Server directly implements all context interfaces

func (s *Server) DocumentManager() (types.DocumentManager, error) {
	if s.documents == nil {
		return nil, fmt.Errorf("document manager not initialized")
	}
	return s.documents, nil
}

func (s *Server) Workspace() types.Workspace {
	return s.workspace
}

func (s *Server) DebugLog(format string, args ...any) {
	helpers.SafeDebugLog(format, args...)
}

func (s *Server) InitializeManifests() error {
	// Initialize the manifest registry
	if err := s.registry.LoadFromWorkspace(s.workspace); err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	// Start file watching for manifest changes
	if err := s.registry.StartFileWatching(s.handleManifestReload); err != nil {
		helpers.SafeDebugLog("Warning: Could not start file watching: %v", err)
		// Don't fail startup if file watching fails
	}

	// Start generate watcher for local project source file changes
	// Note: Uses a grace period to ignore initial event flood from fsnotify
	if err := s.registry.StartGenerateWatcher(); err != nil {
		helpers.SafeDebugLog("Warning: Could not start generate watcher: %v", err)
		// Don't fail startup if generate watcher fails
	}

	return nil
}

// Document operations

func (s *Server) Document(uri string) types.Document {
	return s.documents.Document(uri)
}

func (s *Server) AllDocuments() []types.Document {
	return s.documents.AllDocuments()
}

// Registry operations

func (s *Server) AllTagNames() []string {
	return s.registry.AllTagNames()
}

func (s *Server) Element(tagName string) (*M.CustomElement, bool) {
	return s.registry.Element(tagName)
}

func (s *Server) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return s.registry.Attributes(tagName)
}

func (s *Server) Slots(tagName string) ([]M.Slot, bool) {
	return s.registry.Slots(tagName)
}

func (s *Server) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return s.registry.ElementDefinition(tagName)
}

// ManifestCount returns the number of loaded manifests
func (s *Server) ManifestCount() int {
	s.registry.mu.RLock()
	defer s.registry.mu.RUnlock()
	return len(s.registry.Manifests)
}

// ElementCount returns the number of loaded elements
func (s *Server) ElementCount() int {
	s.registry.mu.RLock()
	defer s.registry.mu.RUnlock()
	return len(s.registry.Elements)
}

func (s *Server) AddManifest(manifest *M.Package) {
	s.registry.AddManifest(manifest)
}

func (s *Server) ElementSource(tagName string) (string, bool) {
	// Get the element definition which contains source information
	definition, exists := s.registry.ElementDefinition(tagName)
	if !exists {
		return "", false
	}

	packageName := definition.PackageName()
	modulePath := definition.ModulePath()

	// Always use bare specifiers - they work for both npm packages and local modules
	// when configured with import maps or package.json exports

	if packageName != "" && modulePath != "" {
		// Ensure we use relative module path, not absolute
		// If modulePath is absolute (starts with /), extract the relative part
		relativeModulePath := modulePath
		if strings.HasPrefix(modulePath, "/") {
			// Extract relative path from workspace root
			workspaceRoot := s.WorkspaceRoot()
			if workspaceRoot != "" && strings.HasPrefix(modulePath, workspaceRoot) {
				// Remove workspace root prefix to get relative path
				relativeModulePath = strings.TrimPrefix(modulePath, workspaceRoot)
				relativeModulePath = strings.TrimPrefix(relativeModulePath, "/")
			}
		}

		// Combine package name with relative module path for bare specifier
		// e.g. "cem-lsp-demo" + "components/button-element.js" = "cem-lsp-demo/components/button-element.js"
		// e.g. "@rhds/elements" + "rh-card/rh-card.js" = "@rhds/elements/rh-card/rh-card.js"
		result := path.Join(packageName, relativeModulePath)
		return result, true
	} else if packageName != "" {
		// Just package name (fallback)
		return packageName, true
	} else if modulePath != "" {
		// Just module path for elements without package info
		return modulePath, true
	}

	return "", false
}

func (s *Server) WorkspaceRoot() string {
	if s.workspace != nil {
		return s.workspace.Root()
	}
	return ""
}

func (s *Server) FileSystem() platform.FileSystem {
	return platform.NewOSFileSystem()
}

func (s *Server) QueryManager() (*queries.QueryManager, error) {
	if s.documents == nil {
		return nil, fmt.Errorf("document manager not initialized")
	}
	return s.documents.QueryManager(), nil
}

func (s *Server) ModuleGraph() *modulegraph.ModuleGraph {
	if s.registry == nil {
		return nil
	}
	return s.registry.GetModuleGraph()
}

// UpdateWorkspaceFromLSP updates the workspace context based on LSP initialize parameters
func (s *Server) UpdateWorkspaceFromLSP(rootURI *string, workspaceFolders []protocol.WorkspaceFolder) error {
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Starting workspace context update")
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Current workspace root before update: '%s'", s.workspace.Root())

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

	// Search upward for the actual workspace root
	// This handles cases where the LSP client opens a file in a subdirectory
	// like /path/to/repo/elements/ instead of /path/to/repo/
	workspaceRoot, err := W.FindWorkspaceRoot(newRoot)
	if err != nil {
		helpers.SafeDebugLog("[SERVER_ADAPTER] Warning: Could not find workspace root, using provided path: %v", err)
		// Keep newRoot as-is when search fails
	} else if workspaceRoot != newRoot {
		helpers.SafeDebugLog("[SERVER_ADAPTER] Found workspace root: %s (original: %s)", workspaceRoot, newRoot)
		newRoot = workspaceRoot
	}

	// Create new workspace context with the correct root
	newWorkspace := W.NewFileSystemWorkspaceContext(newRoot)
	if err := newWorkspace.Init(); err != nil {
		helpers.SafeDebugLog("[SERVER_ADAPTER] Failed to initialize new workspace context: %v", err)
		return fmt.Errorf("failed to initialize workspace context: %w", err)
	}

	// Update server's workspace
	oldRoot := s.workspace.Root()
	s.workspace = newWorkspace
	newActualRoot := s.workspace.Root()
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Successfully updated workspace context")
	helpers.SafeDebugLog("[UPDATE_WORKSPACE] Old root: '%s' -> New root: '%s'", oldRoot, newActualRoot)

	return nil
}

func (s *Server) ElementDescription(tagName string) (string, bool) {
	// For now, use a simple approach to get basic element info
	// This can be enhanced later to pull from the full declaration data
	element, exists := s.registry.Element(tagName)
	if !exists {
		return "", false
	}

	// Check if we have basic element information that could serve as description
	// For now, return empty since CustomElement doesn't directly have description
	// but this establishes the interface for when more detailed data is available
	_ = element
	return "", false
}
