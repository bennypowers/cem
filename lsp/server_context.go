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
	"sync"

	"bennypowers.dev/cem/generate"
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

	// Load additional packages from initializationOptions
	if len(s.additionalPackages) > 0 {
		helpers.SafeDebugLog("Loading %d additional packages from initializationOptions", len(s.additionalPackages))
		if err := s.registry.LoadAdditionalPackages(s.additionalPackages); err != nil {
			helpers.SafeDebugLog("Warning: Error loading additional packages: %v", err)
			// Don't fail startup if additional package loading fails
		}
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

// SetAdditionalPackages sets the additional packages to load during manifest initialization
func (s *Server) SetAdditionalPackages(packages []string) {
	s.additionalPackages = packages
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
	mainTags := s.registry.AllTagNames()
	ephemeralTags := s.ephemeralRegistry.AllTagNames()
	if len(ephemeralTags) == 0 {
		return mainTags
	}

	// Merge, deduplicating ephemeral tags already in main
	seen := make(map[string]bool, len(mainTags))
	for _, tag := range mainTags {
		seen[tag] = true
	}
	for _, tag := range ephemeralTags {
		if !seen[tag] {
			mainTags = append(mainTags, tag)
		}
	}
	return mainTags
}

func (s *Server) Element(tagName string) (*M.CustomElement, bool) {
	if el, ok := s.registry.Element(tagName); ok {
		return el, true
	}
	return s.ephemeralRegistry.Element(tagName)
}

func (s *Server) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	if attrs, ok := s.registry.Attributes(tagName); ok {
		return attrs, true
	}
	return s.ephemeralRegistry.Attributes(tagName)
}

func (s *Server) Slots(tagName string) ([]M.Slot, bool) {
	if slots, ok := s.registry.Slots(tagName); ok {
		return slots, true
	}
	return s.ephemeralRegistry.Slots(tagName)
}

func (s *Server) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	if def, ok := s.registry.ElementDefinition(tagName); ok {
		return def, true
	}
	return s.ephemeralRegistry.ElementDefinition(tagName)
}

func (s *Server) FindCustomElementDeclaration(tagName string) *M.CustomElementDeclaration {
	if decl := s.registry.FindCustomElementDeclaration(tagName); decl != nil {
		return decl
	}
	return s.ephemeralRegistry.FindCustomElementDeclaration(tagName)
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
		// Fall back to ephemeral registry
		return s.ephemeralRegistry.ElementSource(tagName)
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
	// Use FindCustomElementDeclaration to get the full declaration with summary/description
	decl := s.registry.FindCustomElementDeclaration(tagName)
	if decl == nil {
		// Fall back to ephemeral registry
		return s.ephemeralRegistry.ElementDescription(tagName)
	}

	// Prefer description over summary
	if decl.Description != "" {
		return decl.Description, true
	}
	if decl.Summary != "" {
		return decl.Summary, true
	}

	return "", false
}

// ephemeralQueryManager lazily creates a QueryManager with GenerateQueries()
// for use by the ephemeral synthesis pipeline
var (
	ephemeralQM     *queries.QueryManager
	ephemeralQMOnce sync.Once
	ephemeralQMErr  error
)

func getEphemeralQueryManager() (*queries.QueryManager, error) {
	ephemeralQMOnce.Do(func() {
		ephemeralQM, ephemeralQMErr = queries.NewQueryManager(queries.GenerateQueries())
	})
	return ephemeralQM, ephemeralQMErr
}

// SynthesizeEphemeralElements runs synthesis of element declarations from a
// TypeScript/JavaScript document that defines custom elements locally.
// The synthesized data is stored in the ephemeral registry and acts as a
// fallback for all LSP features.
//
// Synthesis runs synchronously so that ephemeral data is available before
// diagnostics fire. The per-file cost is ~1-5ms, well within LSP budgets.
//
// If the document is no longer open (e.g., after DidClose), ephemeral data
// for the URI is removed immediately.
func (s *Server) SynthesizeEphemeralElements(uri string) {
	doc := s.Document(uri)
	if doc == nil {
		// Document closed — remove any stale ephemeral data
		s.ephemeralRegistry.Remove(uri)
		return
	}

	lang := doc.Language()
	if lang != "typescript" && lang != "javascript" {
		return
	}

	content, err := doc.Content()
	if err != nil || content == "" {
		return
	}

	contentBytes := []byte(content)

	// Fast check: does this file define any custom elements?
	lspQM, err := s.QueryManager()
	if err != nil {
		return
	}

	definedTags := queries.FindDefinedElementTags(contentBytes, lspQM)
	if len(definedTags) == 0 {
		// No local definitions — remove any stale ephemeral data
		s.ephemeralRegistry.Remove(uri)
		return
	}

	// Check if all defined tags are already in the main registry
	allInMain := true
	for _, tag := range definedTags {
		if _, ok := s.registry.Element(tag); !ok {
			allInMain = false
			break
		}
	}
	if allInMain {
		// All tags are already known via manifests — no need for ephemeral data
		s.ephemeralRegistry.Remove(uri)
		return
	}

	// Run the full generate pipeline on this file's source
	genQM, err := getEphemeralQueryManager()
	if err != nil {
		helpers.SafeDebugLog("[EPHEMERAL] Failed to get generate query manager: %v", err)
		return
	}

	parser := queries.RetrieveTypeScriptParser()
	defer queries.PutTypeScriptParser(parser)

	mp, err := generate.NewEphemeralModuleProcessor(uri, contentBytes, parser, genQM)
	if err != nil {
		helpers.SafeDebugLog("[EPHEMERAL] Failed to create module processor for %s: %v", uri, err)
		return
	}
	defer mp.Close()

	module, _, _, _, errs := mp.Collect()
	if errs != nil {
		helpers.SafeDebugLog("[EPHEMERAL] Errors collecting from %s: %v", uri, errs)
		// Continue — partial results are still useful
	}

	if module == nil {
		return
	}

	// Check if the module has any CustomElementDeclarations
	hasCustomElements := false
	for _, decl := range module.Declarations {
		if _, ok := decl.(*M.CustomElementDeclaration); ok {
			hasCustomElements = true
			break
		}
	}
	if !hasCustomElements {
		s.ephemeralRegistry.Remove(uri)
		return
	}

	// Wrap in a Package and update the ephemeral registry
	pkg := &M.Package{
		SchemaVersion: "2.1.0",
		Modules:       []M.Module{*module},
	}
	s.ephemeralRegistry.Update(uri, pkg)

	helpers.SafeDebugLog("[EPHEMERAL] Synthesized elements for %s: %d declarations", uri, len(module.Declarations))
}
