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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp/helpers"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// ElementDefinition stores a custom element with its source information
type ElementDefinition struct {
	CustomElement *M.CustomElement
	modulePath    string             // Path from the manifest module
	Source        *M.SourceReference // Source reference if available
	packageName   string             // Package name from package.json (if loaded from a package)
}

// ModulePath returns the module path for this element definition
func (ed *ElementDefinition) ModulePath() string {
	return ed.modulePath
}

// PackageName returns the package name for this element definition
func (ed *ElementDefinition) PackageName() string {
	return ed.packageName
}

// SourceHref returns the source href for this element definition
func (ed *ElementDefinition) SourceHref() string {
	if ed.Source != nil {
		return ed.Source.Href
	}
	return ""
}

// Element returns the custom element for this definition
func (ed *ElementDefinition) Element() *M.CustomElement {
	return ed.CustomElement
}

// Registry manages all loaded custom elements manifests and provides
// fast lookup capabilities for LSP features
type Registry struct {
	// Main data protection
	mu sync.RWMutex
	// Elements maps tag names to their definitions
	Elements map[string]*M.CustomElement
	// ElementDefinitions maps tag names to their definitions with source information
	ElementDefinitions map[string]*ElementDefinition
	// Attributes maps element tag names to their available attributes
	attributes map[string]map[string]*M.Attribute
	// Manifests stores all loaded manifest packages
	Manifests []*M.Package
	// ManifestPaths tracks the file paths of loaded manifests for file watching
	ManifestPaths []string
	// ManifestPackageNames tracks the package name for each manifest path
	ManifestPackageNames map[string]string
	// File watching
	fileWatcher platform.FileWatcher
	watcherMu   sync.RWMutex
	onReload    func() // Callback when manifests are reloaded
	// Generate watching for local project
	generateWatcher *InProcessGenerateWatcher
	generateMu      sync.RWMutex
	localWorkspace  W.WorkspaceContext // Track the local workspace for generate watching
}

// NewRegistry creates a new empty registry with the given file watcher.
// For production use, pass platform.NewFSNotifyFileWatcher().
// For testing, pass platform.NewMockFileWatcher().
func NewRegistry(fileWatcher platform.FileWatcher) *Registry {
	return &Registry{
		Elements:             make(map[string]*M.CustomElement),
		ElementDefinitions:   make(map[string]*ElementDefinition),
		attributes:           make(map[string]map[string]*M.Attribute),
		Manifests:            make([]*M.Package, 0),
		ManifestPaths:        make([]string, 0),
		ManifestPackageNames: make(map[string]string),
		fileWatcher:          fileWatcher,
	}
}

// NewRegistryWithDefaults creates a registry with production defaults.
// This is a convenience constructor for production use.
func NewRegistryWithDefaults() (*Registry, error) {
	fileWatcher, err := platform.NewFSNotifyFileWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create file watcher: %w", err)
	}
	return NewRegistry(fileWatcher), nil
}

// LoadFromWorkspace loads all available custom elements manifests from
// the workspace context
func (r *Registry) LoadFromWorkspace(workspace W.WorkspaceContext) error {
	helpers.SafeDebugLog("Loading manifests from workspace...")

	// Clear existing data
	r.clear()

	// 1. Load manifest from workspace itself (if available)
	if err := r.loadWorkspaceManifest(workspace); err != nil {
		helpers.SafeDebugLog("Warning: Could not load workspace manifest: %v", err)
	}

	// 2. Load manifests from node_modules packages
	if err := r.loadNodeModulesManifests(workspace); err != nil {
		helpers.SafeDebugLog("Warning: Could not load node_modules manifests: %v", err)
	}

	// 3. Load manifests specified in config
	if err := r.loadConfigManifests(workspace); err != nil {
		helpers.SafeDebugLog("Warning: Could not load config manifests: %v", err)
	}

	helpers.SafeDebugLog("Loaded %d custom elements from %d manifests", len(r.Elements), len(r.Manifests))
	return nil
}

// clear resets the registry
func (r *Registry) clear() {
	r.Elements = make(map[string]*M.CustomElement)
	r.ElementDefinitions = make(map[string]*ElementDefinition)
	r.attributes = make(map[string]map[string]*M.Attribute)
	r.Manifests = r.Manifests[:0]
	r.ManifestPaths = r.ManifestPaths[:0]
	r.ManifestPackageNames = make(map[string]string)
}

// clearDataOnly resets the registry data but preserves manifest paths for watching
func (r *Registry) clearDataOnly() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Elements = make(map[string]*M.CustomElement)
	r.ElementDefinitions = make(map[string]*ElementDefinition)
	r.attributes = make(map[string]map[string]*M.Attribute)
	r.Manifests = r.Manifests[:0]
	// Note: ManifestPaths are preserved for file watching
}

// loadWorkspaceManifest loads the manifest from the current workspace
func (r *Registry) loadWorkspaceManifest(workspace W.WorkspaceContext) error {
	helpers.SafeDebugLog("Attempting to load workspace manifest...")
	helpers.SafeDebugLog("Workspace root: %s", workspace.Root())
	helpers.SafeDebugLog("Workspace manifest path: %s", workspace.CustomElementsManifestPath())

	pkg, err := workspace.Manifest()
	if err != nil {
		helpers.SafeDebugLog("Error loading workspace manifest: %v", err)
		return err
	}

	if pkg != nil {
		helpers.SafeDebugLog("Successfully loaded workspace manifest with %d modules", len(pkg.Modules))
		for i, module := range pkg.Modules {
			helpers.SafeDebugLog("  Module [%d]: %s with %d declarations", i, module.Path, len(module.Declarations))
			for j, decl := range module.Declarations {
				if customElementDecl, ok := decl.(*M.CustomElementDeclaration); ok {
					helpers.SafeDebugLog("    Declaration [%d]: %s (tag: %s)", j, customElementDecl.Name, customElementDecl.TagName)
				}
			}
		}

		// Try to get package name from workspace package.json
		var packageName string
		if packageJSON, err := workspace.PackageJSON(); err == nil && packageJSON != nil {
			packageName = packageJSON.Name
			helpers.SafeDebugLog("Package name from workspace package.json: '%s'", packageName)
		} else {
			helpers.SafeDebugLog("Could not read workspace package.json: %v", err)
		}

		r.addManifest(pkg, packageName)
		// Track the manifest file path for watching
		manifestPath := workspace.CustomElementsManifestPath()
		if manifestPath != "" {
			helpers.SafeDebugLog("Tracking workspace manifest path: %s", manifestPath)
			r.addManifestPathWithPackageName(manifestPath, packageName)
		} else {
			helpers.SafeDebugLog("Warning: No manifest path returned from workspace")
		}

		// Store the local workspace for potential generate watching
		r.localWorkspace = workspace

		helpers.SafeDebugLog("Loaded workspace manifest with %d modules", len(pkg.Modules))
	} else {
		helpers.SafeDebugLog("Workspace manifest is nil")
	}

	return nil
}

// loadNodeModulesManifests loads manifests from node_modules packages
func (r *Registry) loadNodeModulesManifests(workspace W.WorkspaceContext) error {
	// Look for node_modules directory
	nodeModulesPath := filepath.Join(workspace.Root(), "node_modules")
	entries, err := os.ReadDir(nodeModulesPath)
	if err != nil {
		return fmt.Errorf("could not read node_modules: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pkgPath := filepath.Join(nodeModulesPath, entry.Name())

		// Handle scoped packages (@scope/package)
		if strings.HasPrefix(entry.Name(), "@") {
			scopedEntries, err := os.ReadDir(pkgPath)
			if err != nil {
				continue
			}
			for _, scopedEntry := range scopedEntries {
				if scopedEntry.IsDir() {
					scopedPkgPath := filepath.Join(pkgPath, scopedEntry.Name())
					r.loadPackageManifest(scopedPkgPath)
				}
			}
		} else {
			r.loadPackageManifest(pkgPath)
		}
	}

	return nil
}

// loadPackageManifest loads a manifest from a specific package directory
func (r *Registry) loadPackageManifest(packagePath string) {
	// Read package.json to find customElements field
	packageJSONPath := filepath.Join(packagePath, "package.json")
	packageJSON, err := r.readPackageJSON(packageJSONPath)
	if err != nil {
		return
	}

	if packageJSON.CustomElements == "" {
		return
	}

	// Resolve the manifest path
	manifestPath := filepath.Join(packagePath, packageJSON.CustomElements)
	if !filepath.IsAbs(manifestPath) {
		manifestPath = filepath.Join(packagePath, packageJSON.CustomElements)
	}

	// Load the manifest
	if pkg, err := r.loadManifestFileWithPackageName(manifestPath, packageJSON.Name); err == nil {
		r.addManifest(pkg, packageJSON.Name)
		helpers.SafeDebugLog("Loaded manifest from %s (%s)", packageJSON.Name, manifestPath)
	}
}

// loadConfigManifests loads manifests specified in the config
func (r *Registry) loadConfigManifests(workspace W.WorkspaceContext) error {
	// TODO: Add config support for specifying additional manifest paths
	// This would allow users to specify manifests from non-npm sources
	return nil
}

// readPackageJSON reads and parses a package.json file
func (r *Registry) readPackageJSON(path string) (*M.PackageJSON, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg M.PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	// If this package.json references a custom elements manifest, track it for watching
	if pkg.CustomElements != "" {
		r.addManifestPath(path)
	}

	return &pkg, nil
}

// loadManifestFileWithPackageName loads a custom elements manifest from a file with package name
func (r *Registry) loadManifestFileWithPackageName(path string, packageName string) (*M.Package, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg M.Package
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	// Track this file path for watching with package name
	r.addManifestPathWithPackageName(path, packageName)

	return &pkg, nil
}

// loadManifestFile loads a custom elements manifest from a file
func (r *Registry) loadManifestFile(path string) (*M.Package, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg M.Package
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	// Track this file path for watching
	r.addManifestPath(path)

	return &pkg, nil
}

// addManifest adds a manifest package to the registry with package name context
func (r *Registry) addManifest(manifest *M.Package, packageName string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Manifests = append(r.Manifests, manifest)

	// Index all custom elements from all modules
	for _, module := range manifest.Modules {
		for _, decl := range module.Declarations {
			// Check if this is a custom element declaration
			if customElementDecl, ok := decl.(*M.CustomElementDeclaration); ok {
				element := &customElementDecl.CustomElement
				// Index by tag name
				if element.TagName != "" {
					r.Elements[element.TagName] = element

					// Store the element definition with source information
					elementDef := &ElementDefinition{
						CustomElement: element,
						modulePath:    module.Path,
						Source:        customElementDecl.Source,
						packageName:   packageName,
					}
					// Check if element already exists
					if existing, exists := r.ElementDefinitions[element.TagName]; exists {
						helpers.SafeDebugLog("[REGISTRY] WARNING: Overriding existing element '%s' (old packageName='%s', new packageName='%s')", element.TagName, existing.PackageName(), packageName)
					}
					helpers.SafeDebugLog("[REGISTRY] Registering element '%s' with packageName='%s', modulePath='%s'", element.TagName, packageName, module.Path)
					r.ElementDefinitions[element.TagName] = elementDef

					// Index attributes for this element
					if element.Attributes != nil {
						attrMap := make(map[string]*M.Attribute)
						for i := range element.Attributes {
							attr := &element.Attributes[i]
							attrMap[attr.Name] = attr
						}
						r.attributes[element.TagName] = attrMap
					}
				}
			}
		}
	}
}

// Element returns the custom element definition for a tag name
func (r *Registry) Element(tagName string) (*M.CustomElement, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	element, exists := r.Elements[tagName]
	helpers.SafeDebugLog("[REGISTRY] GetElement('%s'): exists=%t", tagName, exists)
	if exists {
		helpers.SafeDebugLog("[REGISTRY] Element '%s' has %d attributes, %d events", tagName, len(element.Attributes), len(element.Events))
	}
	return element, exists
}

// Attributes returns the available attributes for a custom element tag
func (r *Registry) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	attrs, exists := r.attributes[tagName]
	helpers.SafeDebugLog("[REGISTRY] Attributes('%s'): exists=%t", tagName, exists)
	if exists {
		helpers.SafeDebugLog("[REGISTRY] Element '%s' has %d attributes", tagName, len(attrs))
		for attrName := range attrs {
			helpers.SafeDebugLog("[REGISTRY]   - attribute: %s", attrName)
		}
	}
	return attrs, exists
}

// Slots returns the available slots for a custom element tag
func (r *Registry) Slots(tagName string) ([]M.Slot, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	element, exists := r.Elements[tagName]
	if !exists {
		helpers.SafeDebugLog("[REGISTRY] GetSlots('%s'): element not found", tagName)
		return nil, false
	}

	helpers.SafeDebugLog("[REGISTRY] GetSlots('%s'): exists=%t, slots=%d", tagName, exists, len(element.Slots))
	for _, slot := range element.Slots {
		helpers.SafeDebugLog("[REGISTRY]   - slot: '%s'", slot.Name)
	}

	return element.Slots, true
}

// ElementDefinition returns the element definition with source information for a tag name
func (r *Registry) ElementDefinition(tagName string) (*ElementDefinition, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	definition, exists := r.ElementDefinitions[tagName]
	helpers.SafeDebugLog("[REGISTRY] GetElementDefinition('%s'): exists=%t", tagName, exists)
	if exists {
		helpers.SafeDebugLog("[REGISTRY] Element '%s' module path: %s", tagName, definition.modulePath)
		if definition.Source != nil {
			helpers.SafeDebugLog("[REGISTRY] Element '%s' source href: %s", tagName, definition.Source.Href)
		}
	}
	return definition, exists
}

// AllTagNames returns all registered custom element tag names
func (r *Registry) AllTagNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tags := make([]string, 0, len(r.Elements))
	for tag := range r.Elements {
		tags = append(tags, tag)
	}
	return tags
}

// AddManifest adds a custom elements manifest to the registry
// This allows programmatic addition of manifests from various sources
func (r *Registry) AddManifest(pkg *M.Package) {
	// Fallback to empty string when package name is not available
	r.addManifest(pkg, "")
}

// StartFileWatching initializes file watching for manifest changes
func (r *Registry) StartFileWatching(onReload func()) error {
	r.watcherMu.Lock()
	defer r.watcherMu.Unlock()

	if r.fileWatcher == nil {
		return fmt.Errorf("no file watcher configured")
	}

	r.onReload = onReload

	// Add all known manifest paths to the watcher
	for _, path := range r.ManifestPaths {
		if err := r.fileWatcher.Add(path); err != nil {
			helpers.SafeDebugLog("Warning: Could not watch manifest file %s: %v", path, err)
		} else {
			helpers.SafeDebugLog("Watching manifest file: %s", path)
		}
	}

	// Start watching in a goroutine
	go r.watchFiles()

	return nil
}

// StopFileWatching stops file watching
func (r *Registry) StopFileWatching() error {
	r.watcherMu.Lock()
	defer r.watcherMu.Unlock()

	if r.fileWatcher == nil {
		return nil
	}

	err := r.fileWatcher.Close()
	r.onReload = nil
	return err
}

// watchFiles handles file system events in a goroutine
func (r *Registry) watchFiles() {
	// Get channel references under lock to avoid race condition
	r.watcherMu.RLock()
	if r.fileWatcher == nil {
		r.watcherMu.RUnlock()
		return
	}
	events := r.fileWatcher.Events()
	errors := r.fileWatcher.Errors()
	r.watcherMu.RUnlock()

	for {
		select {
		case event, ok := <-events:
			if !ok {
				return
			}
			r.handleFileChange(event)
		case err, ok := <-errors:
			if !ok {
				return
			}
			helpers.SafeDebugLog("File watcher error: %v", err)
		}
	}
}

// handleFileChange processes a file system change event
func (r *Registry) handleFileChange(event platform.FileWatchEvent) {
	// Only handle write and create events (ignores temporary files, etc.)
	if event.Op&platform.Write == 0 && event.Op&platform.Create == 0 {
		return
	}

	// Check if this is a manifest file we care about
	isManifestFile := slices.Contains(r.ManifestPaths, event.Name)

	if !isManifestFile {
		return
	}

	helpers.SafeDebugLog("Manifest file changed: %s", event.Name)

	// Trigger reload callback if available
	if r.onReload != nil {
		r.onReload()
	}
}

// AddManifestPath adds a manifest file path to the tracking list for file watching
// This is useful for testing scenarios where you need to watch specific manifest files
func (r *Registry) AddManifestPath(path string) {
	r.addManifestPath(path)
}

// ReloadManifestsDirectly bypasses workspace caching by directly reading manifest files from disk
// This is used when manifest files change and we need to force a fresh read
func (r *Registry) ReloadManifestsDirectly() error {
	helpers.SafeDebugLog("Starting direct manifest reload")

	// Save manifest paths before clearing
	manifestPaths := make([]string, len(r.ManifestPaths))
	copy(manifestPaths, r.ManifestPaths)

	// Clear existing data but preserve paths
	r.clearDataOnly()

	// Reload manifests by directly reading files from the tracked manifest paths
	for _, manifestPath := range manifestPaths {
		// Read and parse the manifest file directly
		pkg, err := r.loadManifestFile(manifestPath)
		if err != nil {
			helpers.SafeDebugLog("Warning: Could not reload manifest %s: %v", manifestPath, err)
			continue
		}

		// Get the associated package name for this manifest path
		packageName := r.ManifestPackageNames[manifestPath]
		helpers.SafeDebugLog("Reloading manifest %s with package name: %s", manifestPath, packageName)

		// Add to registry with the preserved package name
		r.addManifest(pkg, packageName)
		helpers.SafeDebugLog("Reloaded manifest: %s with %d modules", manifestPath, len(pkg.Modules))
	}

	r.mu.RLock()
	elementCount := len(r.Elements)
	r.mu.RUnlock()
	helpers.SafeDebugLog("Direct reload complete: %d elements available", elementCount)

	return nil
}

// StartGenerateWatcher starts the generate watcher for the local project if one exists
// This auto-regenerates the manifest when source files change
func (r *Registry) StartGenerateWatcher() error {
	r.generateMu.Lock()
	defer r.generateMu.Unlock()

	// Only start if we have a local workspace and no watcher is running
	if r.localWorkspace == nil || r.generateWatcher != nil {
		return nil
	}

	// Get the workspace root directory
	workspaceRoot := r.localWorkspace.Root()
	manifestPath := r.localWorkspace.CustomElementsManifestPath()

	if manifestPath == "" {
		return nil // No manifest path configured
	}

	// Get globs from workspace configuration
	cfg, err := r.localWorkspace.Config()
	var globs []string
	if err == nil && len(cfg.Generate.Files) > 0 {
		globs = cfg.Generate.Files
	} else {
		// Default globs if no configuration available
		globs = []string{"src/**/*.{ts,js}"}
	}

	// Create callback that updates the registry with generated manifests
	callback := func(pkg *M.Package) error {
		helpers.SafeDebugLog("Received generated manifest with %d modules", len(pkg.Modules))

		// Debug: check the variant attribute type in the received manifest
		for _, module := range pkg.Modules {
			for _, decl := range module.Declarations {
				if customElement, ok := decl.(*M.CustomElementDeclaration); ok {
					if customElement.CustomElement.TagName == "test-button" {
						for _, attr := range customElement.CustomElement.Attributes {
							if attr.Name == "variant" && attr.Type != nil {
								fmt.Printf("[DEBUG] Received manifest - variant type: '%s'\n", attr.Type.Text)
							}
						}
					}
				}
			}
		}

		// Get package name from workspace package.json to preserve package scope
		var packageName string
		if packageJSON, err := r.localWorkspace.PackageJSON(); err == nil && packageJSON != nil {
			packageName = packageJSON.Name
			helpers.SafeDebugLog("[GENERATE_WATCHER] Using package name from workspace package.json: '%s'", packageName)
		} else {
			helpers.SafeDebugLog("[GENERATE_WATCHER] Could not read workspace package.json: %v", err)
		}

		// Add manifest with proper package name instead of empty string
		r.addManifest(pkg, packageName)
		return nil
	}

	// Create and start the in-process generate watcher
	helpers.SafeDebugLog("About to create InProcessGenerateWatcher with globs: %v", globs)
	watcher, err := NewInProcessGenerateWatcher(r.localWorkspace, globs, callback)
	if err != nil {
		helpers.SafeDebugLog("Failed to create InProcessGenerateWatcher: %v", err)
		return fmt.Errorf("failed to create generate watcher: %w", err)
	}

	helpers.SafeDebugLog("Starting in-process generate watcher for workspace: %s", workspaceRoot)

	err = watcher.Start()
	if err != nil {
		return fmt.Errorf("failed to start generate watcher: %w", err)
	}

	r.generateWatcher = watcher
	helpers.SafeDebugLog("In-process generate watcher started successfully")
	return nil
}

// StopGenerateWatcher stops the generate watcher if it's running
func (r *Registry) StopGenerateWatcher() error {
	r.generateMu.Lock()
	defer r.generateMu.Unlock()

	if r.generateWatcher == nil {
		return nil
	}

	helpers.SafeDebugLog("Stopping in-process generate watcher")

	// Stop the in-process watcher
	err := r.generateWatcher.Stop()
	if err != nil {
		helpers.SafeDebugLog("Warning: Error stopping generate watcher: %v", err)
	}

	r.generateWatcher = nil

	helpers.SafeDebugLog("In-process generate watcher stopped")
	return nil
}

// addManifestPathWithPackageName tracks a manifest file path with its package name for watching
func (r *Registry) addManifestPathWithPackageName(path string, packageName string) {
	// Resolve to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		helpers.SafeDebugLog("Warning: Could not resolve manifest path %s: %v", path, err)
		return
	}

	// Check if already tracked
	if slices.Contains(r.ManifestPaths, absPath) {
		// Update package name if already exists
		r.ManifestPackageNames[absPath] = packageName
		return
	}

	r.ManifestPaths = append(r.ManifestPaths, absPath)
	r.ManifestPackageNames[absPath] = packageName
	helpers.SafeDebugLog("Tracking manifest file: %s with package: %s", absPath, packageName)

	// If watcher is active, add this path
	r.watcherMu.RLock()
	if r.fileWatcher != nil {
		if err := r.fileWatcher.Add(absPath); err != nil {
			helpers.SafeDebugLog("Warning: Could not watch manifest file %s: %v", absPath, err)
		} else {
			helpers.SafeDebugLog("Now watching manifest file: %s", absPath)
		}
	}
	r.watcherMu.RUnlock()
}

// addManifestPath tracks a manifest file path for watching
func (r *Registry) addManifestPath(path string) {
	// Resolve to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		helpers.SafeDebugLog("Warning: Could not resolve manifest path %s: %v", path, err)
		return
	}

	// Check if already tracked
	if slices.Contains(r.ManifestPaths, absPath) {
		return
	}

	r.ManifestPaths = append(r.ManifestPaths, absPath)
	helpers.SafeDebugLog("Tracking manifest file: %s", absPath)

	// If watcher is active, add this path
	r.watcherMu.RLock()
	if r.fileWatcher != nil {
		if err := r.fileWatcher.Add(absPath); err != nil {
			helpers.SafeDebugLog("Warning: Could not watch manifest file %s: %v", absPath, err)
		} else {
			helpers.SafeDebugLog("Now watching manifest file: %s", absPath)
		}
	}
	r.watcherMu.RUnlock()
}
