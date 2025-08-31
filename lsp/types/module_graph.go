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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/queries"
)

// ModuleGraph tracks the import/export relationships between modules
// and their custom elements for accurate re-export resolution
type ModuleGraph struct {
	// Focused component responsibilities
	exportTracker     *ExportTracker
	dependencyTracker *DependencyTracker

	// TransitiveElementsCache stores computed transitive elements using lock-free sync.Map for better performance
	// Key: module path (string), Value: []string (slice of all transitively available element tag names)
	TransitiveElementsCache sync.Map

	// Workspace root for lazy building
	workspaceRoot string

	// MaxTransitiveDepth is the maximum depth for transitive closure computation
	// to prevent performance issues with deeply nested dependency chains
	MaxTransitiveDepth int

	// Injected dependencies for improved testability
	fileParser       FileParser
	exportParser     ExportParser
	manifestResolver ManifestResolver
	metrics          MetricsCollector
	queryManager     *queries.QueryManager
}

// NewModuleGraph creates a new empty module graph with default dependencies
func NewModuleGraph(queryManager *queries.QueryManager) *ModuleGraph {
	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         &OSFileParser{},
		exportParser:       &DefaultExportParser{},
		manifestResolver:   &NoOpManifestResolver{},
		metrics:            &NoOpMetricsCollector{}, // Default to no-op for performance
		queryManager:       queryManager,            // Required parameter for dependency injection
		MaxTransitiveDepth: DefaultMaxTransitiveDepth,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithMetrics creates a new module graph with metrics collection enabled
func NewModuleGraphWithMetrics(queryManager *queries.QueryManager) *ModuleGraph {
	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         &OSFileParser{},
		exportParser:       &DefaultExportParser{},
		manifestResolver:   &NoOpManifestResolver{},
		metrics:            NewDefaultMetricsCollector(),
		queryManager:       queryManager, // Required parameter for dependency injection
		MaxTransitiveDepth: DefaultMaxTransitiveDepth,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithDependencies creates a new module graph with custom dependencies
// This is useful for testing and dependency injection
func NewModuleGraphWithDependencies(fileParser FileParser, exportParser ExportParser, manifestResolver ManifestResolver, metrics MetricsCollector, queryManager *queries.QueryManager) *ModuleGraph {
	// Validate all parameters and provide safe defaults
	if fileParser == nil {
		fileParser = &OSFileParser{}
	}
	if exportParser == nil {
		exportParser = &DefaultExportParser{}
	}
	if manifestResolver == nil {
		manifestResolver = &NoOpManifestResolver{}
	}
	if metrics == nil {
		metrics = &NoOpMetricsCollector{}
	}
	// Note: queryManager can be nil for graceful degradation

	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         fileParser,
		exportParser:       exportParser,
		manifestResolver:   manifestResolver,
		metrics:            metrics,
		queryManager:       queryManager,
		MaxTransitiveDepth: DefaultMaxTransitiveDepth,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithFileParser creates a new module graph with a custom file parser
// This is useful for testing with mock file systems (backwards compatibility)
func NewModuleGraphWithFileParser(fileParser FileParser, queryManager *queries.QueryManager) *ModuleGraph {
	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         fileParser,
		exportParser:       &DefaultExportParser{},
		manifestResolver:   &NoOpManifestResolver{},
		metrics:            &NoOpMetricsCollector{},
		queryManager:       queryManager, // Required parameter for dependency injection
		MaxTransitiveDepth: DefaultMaxTransitiveDepth,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// AddDirectExport registers a direct export (element defined in this module)
func (mg *ModuleGraph) AddDirectExport(modulePath, elementName, tagName string) {
	mg.exportTracker.AddDirectExport(modulePath, elementName, tagName)
	mg.metrics.IncrementCounter("exports_added")
	mg.updateGauges()
}

// AddReExport registers a re-export relationship
func (mg *ModuleGraph) AddReExport(reExportingModule, sourceModule, elementName, tagName string) {
	mg.exportTracker.AddReExport(reExportingModule, sourceModule, elementName, tagName)
	mg.dependencyTracker.AddReExportChain(reExportingModule, sourceModule)
}

// GetElementSources returns all module paths that export the given custom element tag
func (mg *ModuleGraph) GetElementSources(tagName string) []string {
	return mg.exportTracker.GetElementSources(tagName)
}

// GetAllTagNames returns all custom element tag names tracked in the module graph
func (mg *ModuleGraph) GetAllTagNames() []string {
	return mg.exportTracker.GetAllTagNames()
}

// GetAllModulePaths returns all module paths that have exports
func (mg *ModuleGraph) GetAllModulePaths() []string {
	return mg.exportTracker.GetAllModulePaths()
}

// GetModuleDependencies returns all direct dependencies of a module (for debugging)
func (mg *ModuleGraph) GetModuleDependencies(modulePath string) []string {
	return mg.dependencyTracker.GetModuleDependencies(modulePath)
}

// GetModuleExports returns all exports from the given module
func (mg *ModuleGraph) GetModuleExports(modulePath string) []ModuleExport {
	return mg.exportTracker.GetModuleExports(modulePath)
}

// AddModuleDependency registers that one module imports another
func (mg *ModuleGraph) AddModuleDependency(importingModule, importedModule string) {
	mg.dependencyTracker.AddModuleDependency(importingModule, importedModule)
	// Invalidate transitive closure cache for this module (lock-free)
	mg.TransitiveElementsCache.Delete(importingModule)
	mg.metrics.IncrementCounter("dependencies_added")
}

// GetTransitiveElements returns all custom elements available when importing a module
// This includes elements defined in the module itself and all transitively imported modules
// The module graph only tracks import relationships; element definitions come from the manifest
func (mg *ModuleGraph) GetTransitiveElements(modulePath string) []string {
	start := time.Now()
	defer func() {
		mg.metrics.RecordDuration("transitive_computation_time", time.Since(start))
	}()

	// Input validation
	if modulePath == "" {
		return nil
	}

	// Check if already computed (lock-free)
	if cached, exists := mg.TransitiveElementsCache.Load(modulePath); exists {
		mg.metrics.IncrementCounter("cache_hits")
		// Return a copy to avoid concurrent modification
		cachedElements := cached.([]string)
		result := make([]string, len(cachedElements))
		copy(result, cachedElements)
		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 CACHE HIT: Returning cached result for '%s': %v", modulePath, result)
		return result
	}

	mg.metrics.IncrementCounter("cache_misses")
	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 CACHE MISS: Computing fresh result for '%s'", modulePath)

	// Double-check after acquiring lock (another goroutine might have computed it)
	if cached, exists := mg.TransitiveElementsCache.Load(modulePath); exists {
		mg.metrics.IncrementCounter("cache_hits")
		cachedElements := cached.([]string)
		result := make([]string, len(cachedElements))
		copy(result, cachedElements)
		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 CACHE HIT (double-check): Returning cached result for '%s': %v", modulePath, result)
		return result
	}

	// Compute transitive closure using manifest resolver
	mg.metrics.IncrementCounter("transitive_computations")
	elements := mg.calculateTransitiveElementsFromManifest(modulePath)

	// Cache the result (lock-free)
	mg.TransitiveElementsCache.Store(modulePath, elements)
	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 CACHE STORE: Caching result for '%s': %v", modulePath, elements)

	// Return a copy
	result := make([]string, len(elements))
	copy(result, elements)
	return result
}

// calculateTransitiveElementsFromManifest computes all elements transitively available from a module
// using the manifest resolver to determine which elements are available from each manifest module
func (mg *ModuleGraph) calculateTransitiveElementsFromManifest(modulePath string) []string {
	// The module graph should only work with manifest resolvers
	// Unit tests that need direct export tracking should use calculateTransitiveElementsDirect directly
	if mg.manifestResolver == nil {
		helpers.SafeDebugLog("[MODULE_GRAPH] No manifest resolver available for '%s'", modulePath)
		return nil
	}

	visited := make(map[string]bool)
	elements := make(map[string]bool) // Use map to avoid duplicates
	maxDepth := mg.MaxTransitiveDepth // Configurable depth limit to prevent performance issues

	// Breadth-first traversal to collect all transitive manifest modules
	queue := []string{modulePath}
	depth := 0

	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: Starting traversal for '%s' with maxDepth=%d", modulePath, maxDepth)

	for len(queue) > 0 && depth < maxDepth {
		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: Depth %d - Queue: %v", depth, queue)
		nextQueue := []string{}

		for _, currentModule := range queue {
			if visited[currentModule] {
				helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG:   Skipping already visited: %s", currentModule)
				continue // Skip already visited modules (handles cycles)
			}
			visited[currentModule] = true
			helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG:   Processing module: %s", currentModule)

			// The module graph only tracks import relationships at the manifest module level.
			// Elements are determined by the manifest resolver, which has access to the registry
			// and knows which custom elements are defined in each manifest module.
			// This module is now "imported" so any elements it defines should be available.
			// Note: We don't determine the elements here - that's the manifest resolver's job.

			// Add dependencies to next level queue
			dependencies := mg.dependencyTracker.GetModuleDependencies(currentModule)
			helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG:     Dependencies: %v", dependencies)
			for _, dependency := range dependencies {
				if !visited[dependency] {
					nextQueue = append(nextQueue, dependency)
					helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG:     Adding to next queue: %s", dependency)
				} else {
					helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG:     Skipping already visited dependency: %s", dependency)
				}
			}
		}

		queue = nextQueue
		depth++
		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: End of depth %d, next queue: %v", depth-1, queue)
	}

	// Convert visited map to slice for logging
	visitedKeys := make([]string, 0, len(visited))
	for module := range visited {
		visitedKeys = append(visitedKeys, module)
	}
	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: Traversal complete. Final visited: %v", visitedKeys)

	// Get the list of all transitively imported manifest modules
	allImportedModules := make([]string, 0, len(visited))
	for module := range visited {
		allImportedModules = append(allImportedModules, module)
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Transitively imported manifest modules for %s: %v", modulePath, allImportedModules)

	// For each transitively imported manifest module, get its available custom elements
	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: Processing %d visited modules for element collection", len(allImportedModules))
	for _, manifestModule := range allImportedModules {
		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: Getting elements from manifest module: %s", manifestModule)
		// Use the manifest resolver to determine which custom elements are available from this module
		moduleElements := mg.manifestResolver.GetElementsFromManifestModule(manifestModule)
		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG:   Raw elements: %v", moduleElements)

		// If no elements found, try with normalized path (add ./ prefix if missing)
		if len(moduleElements) == 0 && !strings.HasPrefix(manifestModule, "./") {
			normalizedModule := "./" + manifestModule
			moduleElements = mg.manifestResolver.GetElementsFromManifestModule(normalizedModule)
			helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: Trying normalized path '%s' -> '%s', found %d elements", manifestModule, normalizedModule, len(moduleElements))
		}

		for _, element := range moduleElements {
			elements[element] = true
			helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG:   Adding element to result: %s", element)
		}
		helpers.SafeDebugLog("[MODULE_GRAPH] Manifest module '%s' provides %d elements: %v", manifestModule, len(moduleElements), moduleElements)
	}

	// Convert elements map to slice
	result := make([]string, 0, len(elements))
	for element := range elements {
		result = append(result, element)
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 BFS DEBUG: Final elements map contains %d elements: %v", len(elements), elements)
	helpers.SafeDebugLog("[MODULE_GRAPH] Total transitive elements for '%s': %d elements: %v", modulePath, len(result), result)
	return result
}

// GetTransitiveElementsDirect computes transitive elements using direct element tracking
// This method is exposed for testing scenarios where no manifest resolver is available
func (mg *ModuleGraph) GetTransitiveElementsDirect(modulePath string) []string {
	return mg.calculateTransitiveElementsDirect(modulePath)
}

// calculateTransitiveElementsDirect computes transitive elements using direct element tracking (for tests)
func (mg *ModuleGraph) calculateTransitiveElementsDirect(modulePath string) []string {
	visited := make(map[string]bool)
	elements := make(map[string]bool)
	maxDepth := mg.MaxTransitiveDepth

	// Debug: Check what exports we have for the starting module
	startingExports := mg.exportTracker.GetModuleExports(modulePath)
	helpers.SafeDebugLog("[MODULE_GRAPH] Starting module '%s' has %d exports: %v", modulePath, len(startingExports), startingExports)

	// Breadth-first traversal to collect all transitively imported modules
	queue := []string{modulePath}
	depth := 0

	for len(queue) > 0 && depth < maxDepth {
		nextQueue := []string{}

		for _, currentModule := range queue {
			if visited[currentModule] {
				continue
			}
			visited[currentModule] = true

			// Add elements directly tracked in this module
			moduleExports := mg.exportTracker.GetModuleExports(currentModule)
			helpers.SafeDebugLog("[MODULE_GRAPH] Module '%s' has %d exports", currentModule, len(moduleExports))
			for _, export := range moduleExports {
				elements[export.TagName] = true
				helpers.SafeDebugLog("[MODULE_GRAPH] Adding element '%s' from module '%s'", export.TagName, currentModule)
			}

			// Add dependencies to next level queue
			dependencies := mg.dependencyTracker.GetModuleDependencies(currentModule)
			helpers.SafeDebugLog("[MODULE_GRAPH] Module '%s' has %d dependencies: %v", currentModule, len(dependencies), dependencies)
			for _, dependency := range dependencies {
				if !visited[dependency] {
					nextQueue = append(nextQueue, dependency)
				}
			}
		}

		queue = nextQueue
		depth++
	}

	// Convert elements map to slice
	result := make([]string, 0, len(elements))
	for element := range elements {
		result = append(result, element)
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Total transitive elements for '%s': %d elements: %v", modulePath, len(result), result)
	return result
}

// PopulateFromManifests populates the module graph with custom element definitions from manifests
func (mg *ModuleGraph) PopulateFromManifests(elements map[string]interface{}) {
	// elements is expected to be a map[string]*ElementDefinition from registry
	// but we use interface{} to avoid circular imports
	for tagName, element := range elements {
		if elementDef, ok := element.(interface {
			GetModulePath() string
			GetTagName() string
			GetClassName() string
		}); ok {
			modulePath := elementDef.GetModulePath()
			className := elementDef.GetClassName()
			if modulePath != "" && tagName != "" {
				mg.AddDirectExport(modulePath, className, tagName)
			}
		}
	}
}

// BuildFromWorkspace analyzes the workspace to build the module graph
// by parsing TypeScript/JavaScript files for import/export statements only
func (mg *ModuleGraph) BuildFromWorkspace(workspaceRoot string) error {
	start := time.Now()
	defer func() {
		mg.metrics.RecordDuration("build_workspace_time", time.Since(start))
		mg.updateGauges()
	}()

	mg.metrics.IncrementCounter("build_workspace_calls")

	// Walk the workspace to find TypeScript/JavaScript files
	err := mg.fileParser.WalkWorkspace(workspaceRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			mg.metrics.IncrementCounter("file_walk_errors")
			return err
		}

		// Skip node_modules and other common directories
		if info.IsDir() {
			dirName := info.Name()
			if dirName == "node_modules" || dirName == ".git" || dirName == "dist" || dirName == "build" {
				return filepath.SkipDir
			}
			return nil
		}

		// Only process TypeScript and JavaScript files
		ext := filepath.Ext(path)
		if ext != ".ts" && ext != ".js" && ext != ".mjs" {
			return nil
		}

		mg.metrics.IncrementCounter("files_processed")

		// Parse exports from this file
		parseStart := time.Now()
		parseErr := mg.parseFileExports(path, workspaceRoot)
		mg.metrics.RecordDuration("file_parse_time", time.Since(parseStart))

		if parseErr != nil {
			mg.metrics.IncrementCounter("parse_errors")
		}

		return parseErr
	})

	if err != nil {
		mg.metrics.IncrementCounter("build_failures")
		return fmt.Errorf("failed to build module graph from workspace %s: %w", workspaceRoot, err)
	}

	mg.resolveReExportChains()
	mg.logSummary()
	mg.metrics.IncrementCounter("build_successes")
	return nil
}

// parseFileExports parses a single file to extract export statements
func (mg *ModuleGraph) parseFileExports(filePath, workspaceRoot string) error {
	content, err := mg.fileParser.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Convert absolute path to relative module path
	relPath, err := filepath.Rel(workspaceRoot, filePath)
	if err != nil {
		// If relative path conversion fails, use absolute path as fallback
		// This ensures parsing continues even with path resolution issues
		relPath = filePath
	}

	// Normalize path separators for consistent module paths
	modulePath := filepath.ToSlash(relPath)

	// Parse exports using injected parser
	return mg.exportParser.ParseExportsFromContent(modulePath, content, mg.exportTracker, mg.dependencyTracker, mg.queryManager)
}

// resolveReExportChains resolves transitive re-export relationships
func (mg *ModuleGraph) resolveReExportChains() {
	// Get re-export chains from dependency tracker
	reExportChains := mg.dependencyTracker.GetReExportChains()

	// Resolve re-export chains to properly track which modules make elements available
	// For each re-exporting module, find the elements from its source modules
	for reExportingModule, sourceModules := range reExportChains {
		for _, sourceModule := range sourceModules {
			// Find all elements exported by the source module
			exports := mg.exportTracker.GetModuleExports(sourceModule)
			for _, export := range exports {
				// Add the re-exporting module as a source for this element
				tagName := export.TagName
				if tagName != "" {
					mg.exportTracker.AddElementSource(tagName, reExportingModule)
				}
			}
		}
	}
}

// logSummary logs a summary of the built module graph
func (mg *ModuleGraph) logSummary() {
	// Log summary - using plain format to avoid circular dependencies
	// In practice, the calling site would use proper logging
}

// updateGauges updates gauge metrics with current state information
func (mg *ModuleGraph) updateGauges() {
	// Update active modules gauge
	allModules := make(map[string]bool)

	// Count modules from export tracker
	exportTracker := mg.exportTracker
	if exportTracker != nil {
		allTagNames := exportTracker.GetAllTagNames()
		mg.metrics.SetGauge("active_elements", int64(len(allTagNames)))
	}

	// For now, we'll use a simple approximation
	mg.metrics.SetGauge("active_modules", int64(len(allModules)))
}

// GetMetrics returns the current metrics collector (useful for testing)
func (mg *ModuleGraph) GetMetrics() MetricsCollector {
	return mg.metrics
}

// SetManifestResolver updates the manifest resolver for this module graph
func (mg *ModuleGraph) SetManifestResolver(manifestResolver ManifestResolver) {
	mg.manifestResolver = manifestResolver
}

// SetMaxTransitiveDepth updates the maximum transitive depth for this module graph
func (mg *ModuleGraph) SetMaxTransitiveDepth(depth int) {
	if depth <= 0 {
		depth = DefaultMaxTransitiveDepth
	}
	mg.MaxTransitiveDepth = depth
	// Clear cache since depth limit affects results
	mg.TransitiveElementsCache = sync.Map{}
}
