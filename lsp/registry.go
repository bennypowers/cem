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
	"bennypowers.dev/cem/modulegraph"
	"bennypowers.dev/cem/queries"
	"bennypowers.dev/cem/types"
	"gopkg.in/yaml.v3"
)

// ElementDefinition stores a custom element with its source information
type ElementDefinition struct {
	CustomElement *M.CustomElement
	className     string             // Class name from the declaration
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

// GetModulePath returns the module path (for module graph interface compatibility)
func (ed *ElementDefinition) GetModulePath() string {
	return ed.modulePath
}

// GetTagName returns the tag name
func (ed *ElementDefinition) GetTagName() string {
	if ed.CustomElement != nil {
		return ed.CustomElement.TagName
	}
	return ""
}

// GetClassName returns the class name
func (ed *ElementDefinition) GetClassName() string {
	return ed.className
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
	watcherDone chan struct{}  // Signal to stop file watching
	watcherWg   sync.WaitGroup // Wait for watcher goroutine to exit
	onReload    func()         // Callback when manifests are reloaded
	// Generate watching for local project
	generateWatcher platform.GenerateWatcher
	generateMu      sync.RWMutex
	localWorkspace  types.WorkspaceContext // Track the local workspace for generate watching
	// Module graph for tracking re-export relationships
	moduleGraph *modulegraph.ModuleGraph
}

// NewRegistry creates a new empty registry with the given file watcher.
// For production use, pass platform.NewFSNotifyFileWatcher().
// For testing, pass platform.NewMockFileWatcher().
func NewRegistry(fileWatcher platform.FileWatcher) *Registry {
	// Get QueryManager for dependency injection
	queryManager, err := queries.GetGlobalQueryManager()
	if err != nil {
		// For production, this should not happen, but handle gracefully
		queryManager = nil
	}

	moduleGraph := modulegraph.NewModuleGraph(queryManager)

	return &Registry{
		Elements:             make(map[string]*M.CustomElement),
		ElementDefinitions:   make(map[string]*ElementDefinition),
		attributes:           make(map[string]map[string]*M.Attribute),
		Manifests:            make([]*M.Package, 0),
		ManifestPaths:        make([]string, 0),
		ManifestPackageNames: make(map[string]string),
		fileWatcher:          fileWatcher,
		moduleGraph:          moduleGraph,
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
func (r *Registry) LoadFromWorkspace(workspace types.WorkspaceContext) error {
	helpers.SafeDebugLog("Loading manifests from workspace...")

	// Clear existing data
	r.clear()

	// 1. Load manifest from workspace itself (if available)
	if err := r.loadWorkspaceManifest(workspace); err != nil {
		helpers.SafeDebugLog("Warning: Could not load workspace manifest: %v", err)
	}

	// 2. Load manifests from workspace packages (npm/yarn/pnpm workspaces)
	if err := r.loadWorkspacePackageManifests(workspace); err != nil {
		helpers.SafeDebugLog("Warning: Could not load workspace package manifests: %v", err)
	}

	// 3. Load manifests from node_modules packages
	if err := r.loadNodeModulesManifests(workspace); err != nil {
		helpers.SafeDebugLog("Warning: Could not load node_modules manifests: %v", err)
	}

	// 4. Load manifests specified in config
	if err := r.loadConfigManifests(workspace); err != nil {
		helpers.SafeDebugLog("Warning: Could not load config manifests: %v", err)
	}

	// 5. Initialize module graph for lazy building
	// Instead of scanning the entire workspace upfront (which was 5000+ files),
	// we now build the module graph lazily as imports are discovered in open documents.
	// This provides fast startup while maintaining accurate re-export resolution.
	r.initializeLazyModuleGraph(workspace)

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
	// Get QueryManager for dependency injection
	queryManager, err := queries.GetGlobalQueryManager()
	if err != nil {
		// For production, this should not happen, but handle gracefully
		queryManager = nil
	}
	r.moduleGraph = modulegraph.NewModuleGraph(queryManager)
}

// clearDataOnly resets the registry data but preserves manifest paths for watching
func (r *Registry) clearDataOnly() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Elements = make(map[string]*M.CustomElement)
	r.ElementDefinitions = make(map[string]*ElementDefinition)
	r.attributes = make(map[string]map[string]*M.Attribute)
	r.Manifests = r.Manifests[:0]
	// Get QueryManager for dependency injection
	queryManager, err := queries.GetGlobalQueryManager()
	if err != nil {
		// For production, this should not happen, but handle gracefully
		queryManager = nil
	}
	r.moduleGraph = modulegraph.NewModuleGraph(queryManager)
	// Note: ManifestPaths are preserved for file watching
}

// loadWorkspaceManifest loads the manifest from the current workspace
func (r *Registry) loadWorkspaceManifest(workspace types.WorkspaceContext) error {
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

// loadWorkspacePackageManifests loads manifests from workspace packages (npm/yarn/pnpm workspaces)
func (r *Registry) loadWorkspacePackageManifests(workspace types.WorkspaceContext) error {
	helpers.SafeDebugLog("Attempting to load workspace package manifests...")

	// Detect package manager and get workspace configuration
	workspacePackages, err := r.discoverWorkspacePackages(workspace)
	if err != nil {
		return fmt.Errorf("could not discover workspace packages: %w", err)
	}

	if len(workspacePackages) == 0 {
		helpers.SafeDebugLog("No workspace packages found")
		return nil
	}

	helpers.SafeDebugLog("Found %d workspace packages", len(workspacePackages))

	// Load manifest from each workspace package
	for _, pkgPath := range workspacePackages {
		r.loadWorkspacePackage(pkgPath)
	}

	return nil
}

// discoverWorkspacePackages discovers workspace packages based on package manager config
func (r *Registry) discoverWorkspacePackages(workspace types.WorkspaceContext) ([]string, error) {
	root := workspace.Root()
	var workspacePatterns []string
	var packageManager string

	// Detect package manager and get workspace patterns
	// 1. Try pnpm (pnpm-workspace.yaml)
	pnpmWorkspaceFile := filepath.Join(root, "pnpm-workspace.yaml")
	if _, err := os.Stat(pnpmWorkspaceFile); err == nil {
		packageManager = "pnpm"
		patterns, err := r.parsePnpmWorkspace(pnpmWorkspaceFile)
		if err == nil {
			workspacePatterns = patterns
			helpers.SafeDebugLog("Detected pnpm workspace with %d patterns", len(patterns))
		}
	}

	// 2. Try npm/yarn (package.json workspaces field)
	if packageManager == "" {
		packageJSONPath := filepath.Join(root, "package.json")
		if _, err := os.Stat(packageJSONPath); err == nil {
			patterns, err := r.parseNpmYarnWorkspaces(packageJSONPath)
			if err == nil && len(patterns) > 0 {
				workspacePatterns = patterns
				// Determine if npm or yarn based on lock file
				if _, err := os.Stat(filepath.Join(root, "yarn.lock")); err == nil {
					packageManager = "yarn"
					helpers.SafeDebugLog("Detected yarn workspace with %d patterns", len(patterns))
				} else {
					packageManager = "npm"
					helpers.SafeDebugLog("Detected npm workspace with %d patterns", len(patterns))
				}
			}
		}
	}

	if len(workspacePatterns) == 0 {
		return nil, nil
	}

	// Separate positive and negative patterns
	var positivePatterns []string
	var negativePatterns []string
	for _, pattern := range workspacePatterns {
		if strings.HasPrefix(pattern, "!") {
			// Remove the "!" prefix for negation patterns
			negativePatterns = append(negativePatterns, strings.TrimPrefix(pattern, "!"))
		} else {
			positivePatterns = append(positivePatterns, pattern)
		}
	}

	// Expand positive glob patterns to find actual workspace package directories
	var packageDirs []string
	for _, pattern := range positivePatterns {
		matches, err := r.expandWorkspacePattern(workspace, pattern)
		if err != nil {
			helpers.SafeDebugLog("Warning: Could not expand pattern %s: %v", pattern, err)
			continue
		}
		packageDirs = append(packageDirs, matches...)
	}

	// Filter out packages matching negation patterns
	if len(negativePatterns) > 0 {
		packageDirs = r.filterNegatedPackages(workspace, packageDirs, negativePatterns)
	}

	helpers.SafeDebugLog("Expanded %d patterns to %d package directories", len(workspacePatterns), len(packageDirs))
	return packageDirs, nil
}

// parsePnpmWorkspace parses pnpm-workspace.yaml and returns workspace patterns
func (r *Registry) parsePnpmWorkspace(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read pnpm-workspace.yaml: %w", err)
	}

	var config struct {
		Packages []string `yaml:"packages"`
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse pnpm-workspace.yaml: %w", err)
	}

	return config.Packages, nil
}

// parseNpmYarnWorkspaces parses package.json workspaces field
func (r *Registry) parseNpmYarnWorkspaces(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read package.json: %w", err)
	}

	var pkg struct {
		Workspaces interface{} `json:"workspaces"`
	}

	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse package.json: %w", err)
	}

	if pkg.Workspaces == nil {
		return nil, nil
	}

	// Handle both array format and object format
	switch w := pkg.Workspaces.(type) {
	case []interface{}:
		// Array format: ["packages/*"]
		patterns := make([]string, 0, len(w))
		for _, p := range w {
			if str, ok := p.(string); ok {
				patterns = append(patterns, str)
			}
		}
		return patterns, nil
	case map[string]interface{}:
		// Object format: {"packages": ["packages/*"]}
		if pkgs, ok := w["packages"].([]interface{}); ok {
			patterns := make([]string, 0, len(pkgs))
			for _, p := range pkgs {
				if str, ok := p.(string); ok {
					patterns = append(patterns, str)
				}
			}
			return patterns, nil
		}
	}

	return nil, nil
}

// expandWorkspacePattern expands a glob pattern to actual directories using workspace.Glob
func (r *Registry) expandWorkspacePattern(workspace types.WorkspaceContext, pattern string) ([]string, error) {
	// Use workspace.Glob which supports ** doublestar patterns
	matches, err := workspace.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to expand pattern %s: %w", pattern, err)
	}

	root := workspace.Root()

	// Filter to only directories that contain package.json
	var packageDirs []string
	for _, relPath := range matches {
		absPath := filepath.Join(root, relPath)
		info, err := os.Stat(absPath)
		if err != nil || !info.IsDir() {
			continue
		}

		// Check if this directory has a package.json
		packageJSONPath := filepath.Join(absPath, "package.json")
		if _, err := os.Stat(packageJSONPath); err == nil {
			packageDirs = append(packageDirs, absPath)
		}
	}

	return packageDirs, nil
}

// filterNegatedPackages removes packages matching negation patterns
func (r *Registry) filterNegatedPackages(workspace types.WorkspaceContext, packageDirs []string, negativePatterns []string) []string {
	root := workspace.Root()
	var filtered []string

	for _, pkgDir := range packageDirs {
		// Convert to relative path for pattern matching
		relPath, err := filepath.Rel(root, pkgDir)
		if err != nil {
			// If we can't get relative path, skip this package
			continue
		}

		excluded := false
		for _, negPattern := range negativePatterns {
			// Try to match the negation pattern
			matches, err := workspace.Glob(negPattern)
			if err != nil {
				continue
			}

			// Check if this package matches any negation pattern
			for _, match := range matches {
				if filepath.Clean(match) == filepath.Clean(relPath) {
					excluded = true
					helpers.SafeDebugLog("Excluding package %s (matched negation pattern %s)", relPath, negPattern)
					break
				}
			}
			if excluded {
				break
			}
		}

		if !excluded {
			filtered = append(filtered, pkgDir)
		}
	}

	return filtered
}

// loadWorkspacePackage loads a manifest from a single workspace package directory
func (r *Registry) loadWorkspacePackage(pkgPath string) {
	helpers.SafeDebugLog("Loading workspace package from: %s", pkgPath)
	r.loadPackageManifest(pkgPath)
}

// loadNodeModulesManifests loads manifests from node_modules packages
func (r *Registry) loadNodeModulesManifests(workspace types.WorkspaceContext) error {
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
	} else {
		helpers.SafeDebugLog("Failed to load manifest from %s (%s): %v", packageJSON.Name, manifestPath, err)
	}
}

// loadConfigManifests loads manifests specified in the config
func (r *Registry) loadConfigManifests(workspace types.WorkspaceContext) error {
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
						className:     customElementDecl.Name, // Store the class name from the declaration
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

	// First check manifest-based definitions
	definition, exists := r.ElementDefinitions[tagName]
	if exists {
		helpers.SafeDebugLog("[REGISTRY] GetElementDefinition('%s'): found in manifests", tagName)
		helpers.SafeDebugLog("[REGISTRY] Element '%s' module path: %s", tagName, definition.modulePath)
		if definition.Source != nil {
			helpers.SafeDebugLog("[REGISTRY] Element '%s' source href: %s", tagName, definition.Source.Href)
		}
		return definition, true
	}

	// If not found in manifests, check module graph
	if r.moduleGraph != nil {
		elementSources := r.moduleGraph.GetElementSources(tagName)
		if len(elementSources) > 0 {
			helpers.SafeDebugLog("[REGISTRY] GetElementDefinition('%s'): found in module graph with sources: %v", tagName, elementSources)
			// Create a minimal ElementDefinition for module graph elements
			// Use the first source as the primary module path
			modulePath := elementSources[0]
			return &ElementDefinition{
				CustomElement: &M.CustomElement{
					TagName: tagName,
				},
				className:   "", // Class name not available from module graph
				modulePath:  modulePath,
				packageName: "", // Package names not available for module graph elements - they come from file scanning, not manifests
			}, true
		}
	}

	helpers.SafeDebugLog("[REGISTRY] GetElementDefinition('%s'): not found", tagName)
	return nil, false
}

// AllTagNames returns all registered custom element tag names
func (r *Registry) AllTagNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Collect tags from manifests
	tags := make([]string, 0, len(r.Elements))
	for tag := range r.Elements {
		tags = append(tags, tag)
	}

	// Also include tags from module graph (for elements discovered from source files)
	if r.moduleGraph != nil {
		moduleGraphTags := r.moduleGraph.GetAllTagNames()

		// Use map for O(1) duplicate detection instead of O(n²) nested loops
		seen := make(map[string]bool, len(tags))
		for _, tag := range tags {
			seen[tag] = true
		}

		for _, tag := range moduleGraphTags {
			if !seen[tag] {
				tags = append(tags, tag)
				seen[tag] = true
			}
		}
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
// Thread-safety: All operations in this function are protected by watcherMu.Lock()
func (r *Registry) StartFileWatching(onReload func()) error {
	r.watcherMu.Lock()
	defer r.watcherMu.Unlock()

	if r.fileWatcher == nil {
		return fmt.Errorf("no file watcher configured")
	}

	// Prevent starting the watcher multiple times
	if r.watcherDone != nil {
		return fmt.Errorf("file watcher already running")
	}

	r.onReload = onReload
	// PROTECTED BY LOCK: Channel initialization is thread-safe with StopFileWatching
	r.watcherDone = make(chan struct{})

	// Add all known manifest paths to the watcher
	for _, path := range r.ManifestPaths {
		if err := r.fileWatcher.Add(path); err != nil {
			helpers.SafeDebugLog("Warning: Could not watch manifest file %s: %v", path, err)
		} else {
			helpers.SafeDebugLog("Watching manifest file: %s", path)
		}
	}

	// Start watching in a goroutine
	r.watcherWg.Go(func() {
		r.watchFiles()
		helpers.SafeDebugLog("File watcher goroutine exiting")
	})

	return nil
}

// StopFileWatching stops file watching
// Thread-safety: All operations in this function are protected by watcherMu.Lock()
func (r *Registry) StopFileWatching() error {
	r.watcherMu.Lock()
	defer r.watcherMu.Unlock()

	if r.fileWatcher == nil {
		return nil
	}

	// PROTECTED BY LOCK: This nil check and channel close are atomic with respect
	// to StartFileWatching, preventing races between concurrent Start/Stop calls
	if r.watcherDone != nil {
		close(r.watcherDone)
		// Closing is sufficient for shutdown signaling
	}

	// Release lock before waiting to avoid deadlock
	r.watcherMu.Unlock()

	// Wait for goroutine to exit
	r.watcherWg.Wait()
	helpers.SafeDebugLog("File watcher goroutine has exited")

	// Reacquire lock for final cleanup
	r.watcherMu.Lock()

	// Now it's safe to set watcherDone to nil - goroutine has exited
	r.watcherDone = nil
	err := r.fileWatcher.Close()
	r.onReload = nil
	return err
}

// watchFiles handles file system events in a goroutine
func (r *Registry) watchFiles() {
	// Get channel references under lock to avoid race condition
	// We copy the channel references once at startup to avoid holding the lock
	// during the entire watch loop. The channels are immutable once copied,
	// so this is safe even if StopFileWatching is called concurrently.
	r.watcherMu.RLock()
	if r.fileWatcher == nil {
		r.watcherMu.RUnlock()
		return
	}
	events := r.fileWatcher.Events()
	errors := r.fileWatcher.Errors()
	done := r.watcherDone
	r.watcherMu.RUnlock()

	// If done is nil, the select will simply ignore that case (nil channels never trigger)
	// This shouldn't happen in practice, but is safe if it does
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
		case <-done:
			// This case triggers when StopFileWatching closes the done channel
			helpers.SafeDebugLog("File watcher stopped due to shutdown signal")
			return
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

	// Trigger reload callback if available (with proper synchronization)
	r.watcherMu.RLock()
	callback := r.onReload
	r.watcherMu.RUnlock()

	if callback != nil {
		callback()
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

	// Only start if we have a local workspace
	if r.localWorkspace == nil {
		return nil
	}

	// If a generate watcher is already set (e.g., for testing), just start it
	if r.generateWatcher != nil {
		if r.generateWatcher.IsRunning() {
			return nil // Already running
		}
		return r.generateWatcher.Start()
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
					if customElement.TagName == "test-button" {
						for _, attr := range customElement.Attributes {
							if attr.Name == "variant" && attr.Type != nil {
								helpers.SafeDebugLog("Received manifest - variant type: '%s'", attr.Type.Text)
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

// GetModuleGraph returns the module graph for re-export analysis
func (r *Registry) GetModuleGraph() *modulegraph.ModuleGraph {
	return r.moduleGraph
}

// GetFileWatcher returns the file watcher for testing purposes
func (r *Registry) GetFileWatcher() platform.FileWatcher {
	r.watcherMu.RLock()
	defer r.watcherMu.RUnlock()
	return r.fileWatcher
}

// SetGenerateWatcher sets a custom generate watcher for testing purposes
func (r *Registry) SetGenerateWatcher(watcher platform.GenerateWatcher) {
	r.generateMu.Lock()
	defer r.generateMu.Unlock()
	r.generateWatcher = watcher
}

// RegistryManifestResolver implements ManifestResolver using the registry's manifest data
type RegistryManifestResolver struct {
	registry *Registry
}

// NewRegistryManifestResolver creates a new manifest resolver using the registry
func NewRegistryManifestResolver(registry *Registry) *RegistryManifestResolver {
	return &RegistryManifestResolver{registry: registry}
}

// FindManifestModulesForImportPath finds manifest modules that match an import path
func (r *RegistryManifestResolver) FindManifestModulesForImportPath(importPath string) []string {
	if r.registry == nil {
		return nil
	}

	var matchingModules []string

	r.registry.mu.RLock()
	defer r.registry.mu.RUnlock()

	// Search through all element definitions to find matches
	for _, elementDef := range r.registry.ElementDefinitions {
		modulePath := elementDef.GetModulePath()
		if modulePath == "" {
			continue
		}

		// Use the same path matching logic as tagDiagnostics.go
		if r.pathsMatch(importPath, modulePath) {
			// Avoid duplicates
			found := false
			for _, existing := range matchingModules {
				if existing == modulePath {
					found = true
					break
				}
			}
			if !found {
				matchingModules = append(matchingModules, modulePath)
			}
		}
	}

	return matchingModules
}

// GetManifestModulePath converts a file path to its corresponding manifest module path
func (r *RegistryManifestResolver) GetManifestModulePath(filePath string) string {
	if r.registry == nil {
		return ""
	}

	r.registry.mu.RLock()
	defer r.registry.mu.RUnlock()

	// Search through all element definitions to find one that matches this file path
	// This is a reverse lookup from file path to manifest module path
	for _, elementDef := range r.registry.ElementDefinitions {
		modulePath := elementDef.GetModulePath()
		if modulePath == "" {
			continue
		}

		// Try direct path matching first
		if r.pathsMatch(filePath, modulePath) {
			return modulePath
		}

		// For Red Hat Design System pattern: convert TypeScript source paths to manifest paths
		// e.g., "elements/rh-tabs/rh-tab-panel.ts" -> "rh-tabs/rh-tab-panel.js"
		if strings.HasSuffix(filePath, ".ts") {
			// Convert .ts to .js for comparison
			jsFilePath := strings.TrimSuffix(filePath, ".ts") + ".js"
			if r.pathsMatch(jsFilePath, modulePath) {
				return modulePath
			}

			// Also try without elements/ prefix
			// "elements/rh-tabs/rh-tab-panel.ts" -> "rh-tabs/rh-tab-panel.js"
			if strings.HasPrefix(filePath, "elements/") {
				relativeJsPath := strings.TrimPrefix(jsFilePath, "elements/")
				if r.pathsMatch(relativeJsPath, modulePath) {
					return modulePath
				}
			}
		}
	}

	return ""
}

// GetElementsFromManifestModule returns all custom element tag names available from a manifest module
func (r *RegistryManifestResolver) GetElementsFromManifestModule(manifestModulePath string) []string {
	if r.registry == nil {
		return nil
	}

	var elements []string

	r.registry.mu.RLock()
	defer r.registry.mu.RUnlock()

	// Search through all element definitions to find those from this manifest module
	for tagName, elementDef := range r.registry.ElementDefinitions {
		modulePath := elementDef.GetModulePath()
		if modulePath == manifestModulePath {
			elements = append(elements, tagName)
		}
	}

	return elements
}

// pathsMatch checks if an import path matches an element source path
// This duplicates the logic from tagDiagnostics.go for now, but should ideally be unified
func (r *RegistryManifestResolver) pathsMatch(importPath, elementSource string) bool {
	// Direct match first (for exact package imports)
	if importPath == elementSource {
		return true
	}

	// Normalize paths for comparison
	normalizedImport := r.normalizePath(importPath)
	normalizedSource := r.normalizePath(elementSource)

	// Direct match on normalized paths
	if normalizedImport == normalizedSource {
		return true
	}

	// Check if import path ends with the element source (relative imports)
	if strings.HasSuffix(importPath, elementSource) {
		return true
	}

	// Check if element source ends with import path (package imports)
	if strings.HasSuffix(elementSource, importPath) {
		return true
	}

	// Extract just the filename and compare
	importFile := filepath.Base(importPath)
	elementFile := filepath.Base(elementSource)
	return importFile == elementFile
}

// normalizePath normalizes a file path for comparison
// This duplicates the logic from tagDiagnostics.go for now, but should ideally be unified
func (r *RegistryManifestResolver) normalizePath(path string) string {
	// Remove common prefixes/suffixes
	path = strings.TrimPrefix(path, "./")
	path = strings.TrimPrefix(path, "../")
	path = strings.TrimPrefix(path, "/")

	// Handle npm package paths like @rhds/elements/rh-card/rh-card.js
	// vs manifest paths like ./dist/rh-card.js
	if strings.Contains(path, "/") {
		// Keep the last two segments for better matching
		parts := strings.Split(path, "/")
		if len(parts) >= 2 {
			return strings.Join(parts[len(parts)-2:], "/")
		}
	}

	return path
}

// initializeLazyModuleGraph initializes the module graph with manifest data but defers file scanning
func (r *Registry) initializeLazyModuleGraph(workspace types.WorkspaceContext) {
	if workspace == nil {
		return
	}

	workspaceRoot := workspace.Root()
	if workspaceRoot == "" {
		return
	}

	helpers.SafeDebugLog("[REGISTRY] Initializing lazy module graph for workspace: %s", workspaceRoot)

	// Create manifest resolver and update module graph to use it
	manifestResolver := NewRegistryManifestResolver(r)
	r.moduleGraph.SetManifestResolver(manifestResolver)

	// Populate the module graph with custom element definitions from manifests
	// This is fast since it only uses already-loaded manifest data
	elementMap := make(map[string]interface{})
	for tagName, elementDef := range r.ElementDefinitions {
		elementMap[tagName] = elementDef
	}
	r.moduleGraph.PopulateFromManifests(elementMap)

	// Store workspace root for lazy file parsing
	r.moduleGraph.SetWorkspaceRoot(workspaceRoot)

	helpers.SafeDebugLog("[REGISTRY] Module graph initialized with %d manifest elements and manifest resolver, ready for lazy building", len(elementMap))
}
