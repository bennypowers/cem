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
	"slices"
	"sync"
)

// ModuleExport represents an export statement in a module
type ModuleExport struct {
	ElementName  string // Name of the exported element (e.g., "RhTab")
	TagName      string // Custom element tag name (e.g., "rh-tab")
	ExportType   string // "default", "named", "namespace"
	SourceModule string // Module where this is originally defined (if re-export)
}

// ExportTracker manages module exports and element source mappings
type ExportTracker struct {
	mu sync.RWMutex

	// ModuleExports maps module paths to the elements they export
	// Key: module path (e.g., "@rhds/elements/rh-tabs/rh-tabs.js")
	// Value: slice of ModuleExport representing what this module exports
	ModuleExports map[string][]ModuleExport

	// ElementSources maps custom element tag names to all modules that export them
	// Key: tag name (e.g., "rh-tab")
	// Value: slice of module paths that export this element
	ElementSources map[string][]string
}

// NewExportTracker creates a new export tracker
func NewExportTracker() *ExportTracker {
	return &ExportTracker{
		ModuleExports:  make(map[string][]ModuleExport),
		ElementSources: make(map[string][]string),
	}
}

// AddDirectExport registers a direct export (element defined in this module)
func (et *ExportTracker) AddDirectExport(modulePath, elementName, tagName string) {
	// Input validation
	if modulePath == "" || elementName == "" || tagName == "" {
		return // Skip invalid parameters
	}

	et.mu.Lock()
	defer et.mu.Unlock()

	export := ModuleExport{
		ElementName: elementName,
		TagName:     tagName,
		ExportType:  "named",
	}

	et.ModuleExports[modulePath] = append(et.ModuleExports[modulePath], export)
	et.ElementSources[tagName] = append(et.ElementSources[tagName], modulePath)
}

// AddReExport registers a re-export relationship
func (et *ExportTracker) AddReExport(reExportingModule, sourceModule, elementName, tagName string) {
	// Input validation
	if reExportingModule == "" || sourceModule == "" || elementName == "" || tagName == "" {
		return // Skip invalid parameters
	}

	et.mu.Lock()
	defer et.mu.Unlock()

	export := ModuleExport{
		ElementName:  elementName,
		TagName:      tagName,
		ExportType:   "named",
		SourceModule: sourceModule,
	}

	et.ModuleExports[reExportingModule] = append(et.ModuleExports[reExportingModule], export)
	et.ElementSources[tagName] = append(et.ElementSources[tagName], reExportingModule)
}

// GetElementSources returns all module paths that export the given custom element tag
func (et *ExportTracker) GetElementSources(tagName string) []string {
	et.mu.RLock()
	defer et.mu.RUnlock()

	sources, exists := et.ElementSources[tagName]
	if !exists {
		return nil
	}

	// Return a copy to avoid concurrent modification
	result := make([]string, len(sources))
	copy(result, sources)
	return result
}

// GetAllTagNames returns all custom element tag names tracked in the export tracker
func (et *ExportTracker) GetAllTagNames() []string {
	et.mu.RLock()
	defer et.mu.RUnlock()

	tagNames := make([]string, 0, len(et.ElementSources))
	for tagName := range et.ElementSources {
		tagNames = append(tagNames, tagName)
	}

	return tagNames
}

// GetModuleExports returns all exports from the given module
func (et *ExportTracker) GetModuleExports(modulePath string) []ModuleExport {
	et.mu.RLock()
	defer et.mu.RUnlock()

	exports, exists := et.ModuleExports[modulePath]
	if !exists {
		return nil
	}

	// Return a copy to avoid concurrent modification
	result := make([]ModuleExport, len(exports))
	copy(result, exports)
	return result
}

// GetAllModulePaths returns all module paths that have exports
func (et *ExportTracker) GetAllModulePaths() []string {
	et.mu.RLock()
	defer et.mu.RUnlock()

	var modules []string
	for modulePath := range et.ModuleExports {
		modules = append(modules, modulePath)
	}
	return modules
}

// AddElementSource adds a module as a source for an element (used by re-export resolution)
func (et *ExportTracker) AddElementSource(tagName, modulePath string) {
	if tagName == "" || modulePath == "" {
		return
	}

	et.mu.Lock()
	defer et.mu.Unlock()

	// Check if this module is already listed as a source
	for _, existingSource := range et.ElementSources[tagName] {
		if existingSource == modulePath {
			return // Already exists
		}
	}

	et.ElementSources[tagName] = append(et.ElementSources[tagName], modulePath)
}

// DependencyTracker manages module dependencies and re-export chains
type DependencyTracker struct {
	mu sync.RWMutex

	// ReExportChains tracks transitive re-export relationships
	// Key: module path that re-exports
	// Value: slice of source module paths it ultimately depends on
	ReExportChains map[string][]string

	// ModuleDependencies tracks direct import relationships between modules
	// Key: importing module path (e.g., "rh-tabs.js")
	// Value: slice of imported module paths (e.g., ["rh-tab.js", "rh-icon.js"])
	ModuleDependencies map[string][]string
}

// NewDependencyTracker creates a new dependency tracker
func NewDependencyTracker() *DependencyTracker {
	return &DependencyTracker{
		ReExportChains:     make(map[string][]string),
		ModuleDependencies: make(map[string][]string),
	}
}

// AddModuleDependency registers that one module imports another
func (dt *DependencyTracker) AddModuleDependency(importingModule, importedModule string) {
	// Input validation
	if importingModule == "" || importedModule == "" {
		return // Skip invalid module paths
	}

	dt.mu.Lock()
	defer dt.mu.Unlock()

	// Add to dependency list if not already present
	if !slices.Contains(dt.ModuleDependencies[importingModule], importedModule) {
		dt.ModuleDependencies[importingModule] = append(dt.ModuleDependencies[importingModule], importedModule)
	}
}

// AddReExportChain adds a re-export relationship to the chain tracker
func (dt *DependencyTracker) AddReExportChain(reExportingModule, sourceModule string) {
	if reExportingModule == "" || sourceModule == "" {
		return
	}

	dt.mu.Lock()
	defer dt.mu.Unlock()

	if !slices.Contains(dt.ReExportChains[reExportingModule], sourceModule) {
		dt.ReExportChains[reExportingModule] = append(dt.ReExportChains[reExportingModule], sourceModule)
	}
}

// GetModuleDependencies returns the direct dependencies of a module
func (dt *DependencyTracker) GetModuleDependencies(modulePath string) []string {
	dt.mu.RLock()
	defer dt.mu.RUnlock()

	deps, exists := dt.ModuleDependencies[modulePath]
	if !exists {
		return nil
	}

	// Return a copy to avoid concurrent modification
	result := make([]string, len(deps))
	copy(result, deps)
	return result
}

// GetReExportChains returns all re-export chains
func (dt *DependencyTracker) GetReExportChains() map[string][]string {
	dt.mu.RLock()
	defer dt.mu.RUnlock()

	// Return a deep copy to avoid concurrent modification
	result := make(map[string][]string, len(dt.ReExportChains))
	for k, v := range dt.ReExportChains {
		chains := make([]string, len(v))
		copy(chains, v)
		result[k] = chains
	}
	return result
}
