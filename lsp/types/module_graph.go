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
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/queries"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// DefaultMaxTransitiveDepth is the default maximum depth for transitive closure computation
// to prevent performance issues with deeply nested dependency chains.
//
// This limit of 5 levels was chosen based on practical analysis of real-world projects:
// - Most legitimate dependency chains are 2-3 levels deep
// - 5 levels covers complex library architectures while preventing pathological cases
// - Protects against circular dependencies and runaway calculations
// - Balances comprehensive element discovery with acceptable performance
const DefaultMaxTransitiveDepth = 5

// FileParser interface abstracts file reading and workspace traversal operations
// for improved testability and dependency injection
type FileParser interface {
	// ReadFile reads the content of a file at the given path
	ReadFile(path string) ([]byte, error)
	// WalkWorkspace walks through a workspace directory calling the walkFn for each file
	WalkWorkspace(workspaceRoot string, walkFn filepath.WalkFunc) error
}

// ExportParser interface abstracts the parsing of export statements from content
// This provides a higher-level abstraction over tree-sitter operations
type ExportParser interface {
	// ParseExportsFromContent parses export statements from file content
	// Returns any parsing errors, but implementations should be resilient
	ParseExportsFromContent(modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker) error
}

// ManifestResolver interface abstracts manifest-based path resolution
// This allows the module graph to integrate with existing manifest path resolution
type ManifestResolver interface {
	// FindManifestModulesForImportPath finds manifest modules that match an import path
	// Returns module paths from the manifest (e.g., "rh-tabs/rh-tab.js")
	FindManifestModulesForImportPath(importPath string) []string

	// GetManifestModulePath converts a file path to its corresponding manifest module path
	// Returns empty string if no manifest module exists for this file
	GetManifestModulePath(filePath string) string

	// GetElementsFromManifestModule returns all custom element tag names available from a manifest module
	// This is the key method that maps manifest modules to their available custom elements
	GetElementsFromManifestModule(manifestModulePath string) []string
}

// MetricsCollector interface abstracts metrics collection for observability
type MetricsCollector interface {
	// IncrementCounter increments a named counter metric
	IncrementCounter(name string)
	// RecordDuration records a duration metric in milliseconds
	RecordDuration(name string, duration time.Duration)
	// SetGauge sets a gauge metric value
	SetGauge(name string, value int64)
	// AddHistogramValue adds a value to a histogram metric
	AddHistogramValue(name string, value float64)
}

// ModuleGraphMetrics tracks performance and usage metrics for the module graph
type ModuleGraphMetrics struct {
	// Counter metrics (atomic for thread safety)
	FilesProcessed         int64 // Total files processed
	ExportsFound           int64 // Total exports found
	DependenciesTracked    int64 // Total dependencies tracked
	TransitiveComputations int64 // Total transitive closure computations
	CacheHits              int64 // Cache hit count
	CacheMisses            int64 // Cache miss count
	ParseErrors            int64 // Parse error count

	// Performance timing metrics
	TotalBuildTime    time.Duration // Total time for BuildFromWorkspace
	AvgFileParseTime  time.Duration // Average time per file parse
	AvgTransitiveTime time.Duration // Average time for transitive computation

	// Memory usage metrics
	MaxMemoryUsage int64 // Peak memory usage (if measurable)

	// State metrics
	ActiveModules  int64 // Number of modules currently tracked
	ActiveElements int64 // Number of elements currently tracked
}

// DefaultMetricsCollector implements MetricsCollector with in-memory storage
type DefaultMetricsCollector struct {
	counters   sync.Map // map[string]*int64
	gauges     sync.Map // map[string]*int64
	durations  sync.Map // map[string]*durationList
	histograms sync.Map // map[string]*histogramList
}

// durationList provides thread-safe duration slice operations
type durationList struct {
	mu        sync.Mutex
	durations []time.Duration
}

// add appends a duration to the list in a thread-safe manner
func (dl *durationList) add(duration time.Duration) {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	dl.durations = append(dl.durations, duration)
}

// histogramList provides thread-safe histogram slice operations
type histogramList struct {
	mu     sync.Mutex
	values []float64
}

// add appends a value to the histogram in a thread-safe manner
func (hl *histogramList) add(value float64) {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	hl.values = append(hl.values, value)
}

// NewDefaultMetricsCollector creates a new in-memory metrics collector
func NewDefaultMetricsCollector() *DefaultMetricsCollector {
	return &DefaultMetricsCollector{}
}

// IncrementCounter implements MetricsCollector
func (m *DefaultMetricsCollector) IncrementCounter(name string) {
	counterInterface, _ := m.counters.LoadOrStore(name, new(int64))
	counter := counterInterface.(*int64)
	atomic.AddInt64(counter, 1)
}

// RecordDuration implements MetricsCollector
func (m *DefaultMetricsCollector) RecordDuration(name string, duration time.Duration) {
	durationsInterface, _ := m.durations.LoadOrStore(name, &durationList{durations: make([]time.Duration, 0)})
	durationsList := durationsInterface.(*durationList)
	durationsList.add(duration)
}

// SetGauge implements MetricsCollector
func (m *DefaultMetricsCollector) SetGauge(name string, value int64) {
	gaugeInterface, _ := m.gauges.LoadOrStore(name, new(int64))
	gauge := gaugeInterface.(*int64)
	atomic.StoreInt64(gauge, value)
}

// AddHistogramValue implements MetricsCollector
func (m *DefaultMetricsCollector) AddHistogramValue(name string, value float64) {
	histogramInterface, _ := m.histograms.LoadOrStore(name, &histogramList{values: make([]float64, 0)})
	histogramList := histogramInterface.(*histogramList)
	histogramList.add(value)
}

// GetCounterValue returns the current value of a counter
func (m *DefaultMetricsCollector) GetCounterValue(name string) int64 {
	if counterInterface, exists := m.counters.Load(name); exists {
		counter := counterInterface.(*int64)
		return atomic.LoadInt64(counter)
	}
	return 0
}

// GetGaugeValue returns the current value of a gauge
func (m *DefaultMetricsCollector) GetGaugeValue(name string) int64 {
	if gaugeInterface, exists := m.gauges.Load(name); exists {
		gauge := gaugeInterface.(*int64)
		return atomic.LoadInt64(gauge)
	}
	return 0
}

// NoOpMetricsCollector implements MetricsCollector but discards all metrics
type NoOpMetricsCollector struct{}

// IncrementCounter implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) IncrementCounter(name string) {}

// RecordDuration implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) RecordDuration(name string, duration time.Duration) {}

// SetGauge implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) SetGauge(name string, value int64) {}

// AddHistogramValue implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) AddHistogramValue(name string, value float64) {}

// OSFileParser implements FileParser using standard OS operations
type OSFileParser struct{}

// ReadFile implements FileParser using os.ReadFile
func (p *OSFileParser) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WalkWorkspace implements FileParser using filepath.Walk
func (p *OSFileParser) WalkWorkspace(workspaceRoot string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(workspaceRoot, walkFn)
}

// MockFileParser implements FileParser for testing with in-memory file systems
type MockFileParser struct {
	Files map[string][]byte // Map of file paths to their contents
}

// ReadFile implements FileParser by returning content from the in-memory map
func (p *MockFileParser) ReadFile(path string) ([]byte, error) {
	content, exists := p.Files[path]
	if !exists {
		return nil, os.ErrNotExist
	}
	return content, nil
}

// WalkWorkspace implements FileParser by iterating through the in-memory file map
func (p *MockFileParser) WalkWorkspace(workspaceRoot string, walkFn filepath.WalkFunc) error {
	// Create a simplified mock directory structure
	// In real usage, this would simulate proper filesystem traversal
	for path := range p.Files {
		// Create a mock FileInfo for each file
		info := &mockFileInfo{name: filepath.Base(path), isDir: false}
		if err := walkFn(path, info, nil); err != nil {
			if err == filepath.SkipDir {
				continue // Skip this path
			}
			return err
		}
	}
	return nil
}

// mockFileInfo implements os.FileInfo for testing
type mockFileInfo struct {
	name  string
	isDir bool
}

func (m *mockFileInfo) Name() string       { return m.name }
func (m *mockFileInfo) Size() int64        { return 0 }
func (m *mockFileInfo) Mode() os.FileMode  { return 0644 }
func (m *mockFileInfo) ModTime() time.Time { return time.Time{} }
func (m *mockFileInfo) IsDir() bool        { return m.isDir }
func (m *mockFileInfo) Sys() interface{}   { return nil }

// DefaultExportParser implements ExportParser using tree-sitter queries
type DefaultExportParser struct{}

// ParseExportsFromContent implements ExportParser using tree-sitter
func (p *DefaultExportParser) ParseExportsFromContent(modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker) error {
	// Get TypeScript parser from pool
	parser := queries.RetrieveTypeScriptParser()
	defer queries.PutTypeScriptParser(parser)

	// Parse the content
	tree := parser.Parse(content, nil)
	if tree == nil {
		// Log parsing failure and return error for proper error handling
		helpers.SafeDebugLog("[MODULE_GRAPH] Failed to parse file '%s' - syntax errors or invalid TS/JS", modulePath)
		return fmt.Errorf("failed to parse file '%s': syntax errors or invalid TypeScript/JavaScript", modulePath)
	}
	defer tree.Close()

	// Try to use tree-sitter queries for accurate parsing
	return p.parseExportsWithQueries(tree, modulePath, content, exportTracker, dependencyTracker)
}

// parseExportsWithQueries uses tree-sitter queries to parse export statements
func (p *DefaultExportParser) parseExportsWithQueries(tree *ts.Tree, modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker) error {
	// PERFORMANCE OPTIMIZATION: Use thread-safe singleton QueryManager and cached matchers
	// There can be only one QueryManager, but we cache matchers for different query types
	queryManager, err := queries.GetGlobalQueryManager()
	if err != nil {
		return fmt.Errorf("failed to get global query manager for module %s: %w", modulePath, err)
	}

	// Get cached export matcher (thread-safe)
	exportMatcher, err := queries.GetCachedQueryMatcher(queryManager, "typescript", "exports")
	if err != nil {
		return fmt.Errorf("failed to get cached export matcher for module %s: %w", modulePath, err)
	}
	// Note: Don't defer Close() on cached singleton instances

	// Process export matches
	for match := range exportMatcher.AllQueryMatches(tree.RootNode(), content) {
		p.processExportMatch(match, exportMatcher, modulePath, content, exportTracker, dependencyTracker)
	}

	return nil
}

// processExportMatch processes a single export/import match from tree-sitter
func (p *DefaultExportParser) processExportMatch(match *ts.QueryMatch, matcher *queries.QueryMatcher, modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker) {
	var exportName, sourceModule string
	var importSource string
	var isImport bool

	// Process all captures in this match
	for _, capture := range match.Captures {
		captureName := matcher.GetCaptureNameByIndex(capture.Index)
		captureText := strings.TrimSpace(capture.Node.Utf8Text(content))

		switch captureName {
		case "export.name":
			exportName = captureText
		case "export.source":
			sourceModule = captureText
		case "export.class.name", "export.function.name", "export.variable.name", "export.default.name":
			exportName = captureText
		case "import.source", "import.dynamic.source":
			importSource = captureText
			isImport = true
		}
	}

	// Handle import statements - track module dependencies
	if isImport && importSource != "" {
		// Normalize the import path to handle relative imports
		normalizedImportPath := helpers.NormalizeImportPath(importSource)
		dependencyTracker.AddModuleDependency(modulePath, normalizedImportPath)
		return
	}

	// Handle re-exports (export { X } from './module')
	if exportName != "" && sourceModule != "" {
		// Track re-export relationship at module level - this is sufficient because
		// resolveReExportChains() will later resolve all elements from source modules
		// to re-exporting modules using existing manifest data
		normalizedSourceModule := helpers.NormalizeImportPath(sourceModule)
		dependencyTracker.AddModuleDependency(modulePath, normalizedSourceModule)
		dependencyTracker.AddReExportChain(modulePath, normalizedSourceModule)
		return
	}

	// Handle direct exports (export class MyElement)
	if exportName != "" && sourceModule == "" {
		// Direct exports are handled by the existing manifest registry integration.
		// The registry already contains element definitions from manifests by module path,
		// so individual export tracking here would be redundant. The LSP system uses
		// the manifest registry as the authoritative source for element definitions.
		return
	}
}

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
}

// NoOpManifestResolver implements ManifestResolver but returns empty results
type NoOpManifestResolver struct{}

func (r *NoOpManifestResolver) FindManifestModulesForImportPath(importPath string) []string {
	return nil
}

func (r *NoOpManifestResolver) GetManifestModulePath(filePath string) string {
	return ""
}

func (r *NoOpManifestResolver) GetElementsFromManifestModule(manifestModulePath string) []string {
	return nil
}

// NewModuleGraph creates a new empty module graph with default dependencies
func NewModuleGraph() *ModuleGraph {
	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         &OSFileParser{},
		exportParser:       &DefaultExportParser{},
		manifestResolver:   &NoOpManifestResolver{},
		metrics:            &NoOpMetricsCollector{}, // Default to no-op for performance
		MaxTransitiveDepth: DefaultMaxTransitiveDepth,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithMetrics creates a new module graph with metrics collection enabled
func NewModuleGraphWithMetrics() *ModuleGraph {
	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         &OSFileParser{},
		exportParser:       &DefaultExportParser{},
		manifestResolver:   &NoOpManifestResolver{},
		metrics:            NewDefaultMetricsCollector(),
		MaxTransitiveDepth: DefaultMaxTransitiveDepth,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithDependencies creates a new module graph with custom dependencies
// This is useful for testing and dependency injection
func NewModuleGraphWithDependencies(fileParser FileParser, exportParser ExportParser, manifestResolver ManifestResolver, metrics MetricsCollector) *ModuleGraph {
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
	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         fileParser,
		exportParser:       exportParser,
		manifestResolver:   manifestResolver,
		metrics:            metrics,
		MaxTransitiveDepth: DefaultMaxTransitiveDepth,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithFileParser creates a new module graph with a custom file parser
// This is useful for testing with mock file systems (backwards compatibility)
func NewModuleGraphWithFileParser(fileParser FileParser) *ModuleGraph {
	return &ModuleGraph{
		exportTracker:      NewExportTracker(),
		dependencyTracker:  NewDependencyTracker(),
		fileParser:         fileParser,
		exportParser:       &DefaultExportParser{},
		manifestResolver:   &NoOpManifestResolver{},
		metrics:            &NoOpMetricsCollector{},
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
	return mg.exportParser.ParseExportsFromContent(modulePath, content, mg.exportTracker, mg.dependencyTracker)
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

// Lazy Module Graph Building Methods

// SetWorkspaceRoot stores the workspace root for lazy file parsing
func (mg *ModuleGraph) SetWorkspaceRoot(workspaceRoot string) {
	mg.workspaceRoot = workspaceRoot
}

// BuildForImportPath builds module graph entries for a specific import path
func (mg *ModuleGraph) BuildForImportPath(importPath string) error {
	if mg.workspaceRoot == "" {
		return fmt.Errorf("workspace root not set for lazy module graph building")
	}

	start := time.Now()
	defer func() {
		mg.metrics.RecordDuration("lazy_build_time", time.Since(start))
	}()

	mg.metrics.IncrementCounter("lazy_build_calls")

	// First, check if this import path corresponds to any known manifest modules
	// The key insight: instead of resolving to files and parsing,
	// we should discover dependencies between manifest modules
	manifestModule := mg.findManifestModuleForImportPath(importPath)
	if manifestModule != "" {
		helpers.SafeDebugLog("[MODULE_GRAPH] Found manifest module '%s' for import '%s'", manifestModule, importPath)
		// Parse the source file for this manifest module to discover its dependencies
		if err := mg.buildDependenciesForManifestModule(manifestModule); err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to build dependencies for manifest module %s: %v", manifestModule, err)
		}
	}

	// Fallback: resolve to file paths and parse (for non-manifest imports)
	filePaths := mg.resolveImportPathToFiles(importPath)
	if len(filePaths) > 0 {
		helpers.SafeDebugLog("[MODULE_GRAPH] Building graph for import %s -> %d files", importPath, len(filePaths))

		// Parse only the specific files related to this import
		// Use a set to track already processed files to avoid infinite recursion
		processedFiles := make(map[string]bool)

		for _, filePath := range filePaths {
			if err := mg.processFileWithDependencies(filePath, processedFiles); err != nil {
				mg.metrics.IncrementCounter("lazy_parse_errors")
				helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to process file %s: %v", filePath, err)
				// Continue with other files
			}
		}
	}

	// CRITICAL: Resolve re-export chains after parsing files
	// This is essential for making re-exported elements available
	// (e.g., my-tab available through my-tabs.js re-export)
	mg.resolveReExportChains()
	helpers.SafeDebugLog("[MODULE_GRAPH] Resolved re-export chains for lazy build of %s", importPath)

	// Cache invalidation is handled automatically by AddModuleDependency() calls
	// during the recursive dependency building process, so no additional cache clearing needed

	mg.metrics.IncrementCounter("lazy_imports_built")
	return nil
}

// resolveImportPathToFiles converts an import path to actual file paths that need to be parsed
// This function resolves relative imports relative to the workspace root (legacy behavior)
func (mg *ModuleGraph) resolveImportPathToFiles(importPath string) []string {
	return mg.resolveImportPathToFilesRelativeTo(importPath, "")
}

// resolveImportPathToFilesRelativeTo converts an import path to actual file paths,
// resolving relative imports relative to the given file path
func (mg *ModuleGraph) resolveImportPathToFilesRelativeTo(importPath, currentFilePath string) []string {
	var filePaths []string

	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 Resolving import path '%s' to files in workspace '%s' (relative to '%s')", importPath, mg.workspaceRoot, currentFilePath)

	// Handle different import path formats
	if strings.HasPrefix(importPath, "./") || strings.HasPrefix(importPath, "../") {
		// Relative import - resolve relative to current file's directory
		var baseDir string
		if currentFilePath != "" {
			// Get the directory of the current file
			baseDir = filepath.Dir(currentFilePath)
		} else {
			// Fallback to workspace root for backward compatibility
			baseDir = mg.workspaceRoot
		}

		fullPath := filepath.Join(baseDir, importPath)
		helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Trying relative path: %s (base dir: %s)", fullPath, baseDir)
		if mg.fileExists(fullPath) {
			filePaths = append(filePaths, fullPath)
			helpers.SafeDebugLog("[MODULE_GRAPH] ✅ Found relative file: %s", fullPath)
		}

		// For imports like './rh-tab.js', also try corresponding TypeScript source
		// This handles TypeScript projects where imports reference .js but sources are .ts
		if strings.HasSuffix(importPath, ".js") {
			tsPath := strings.TrimSuffix(fullPath, ".js") + ".ts"
			if mg.fileExists(tsPath) {
				filePaths = append(filePaths, tsPath)
				helpers.SafeDebugLog("[MODULE_GRAPH] ✅ Found TypeScript source: %s", tsPath)
			}
		}

		// Also try with common extensions
		for _, ext := range []string{".js", ".ts", ".mjs"} {
			extPath := fullPath + ext
			if mg.fileExists(extPath) {
				// Avoid duplicates
				isDuplicate := false
				for _, existing := range filePaths {
					if existing == extPath {
						isDuplicate = true
						break
					}
				}
				if !isDuplicate {
					filePaths = append(filePaths, extPath)
					helpers.SafeDebugLog("[MODULE_GRAPH] ✅ Found file with extension: %s", extPath)
				}
			}
		}
	} else if strings.Contains(importPath, "/") {
		// Package import like '@rhds/elements/rh-tabs/rh-tabs.js' - look in workspace for matching files
		// This handles cases where packages are developed locally

		// Extract the meaningful part after the package name
		// For '@rhds/elements/rh-tabs/rh-tabs.js' -> 'rh-tabs/rh-tabs.js'
		var relativePath string
		if strings.HasPrefix(importPath, "@") {
			// Handle scoped packages like @rhds/elements/path/file.js
			parts := strings.SplitN(importPath, "/", 3)
			helpers.SafeDebugLog("[MODULE_GRAPH] 📦 Scoped package parts: %v", parts)
			if len(parts) >= 3 {
				relativePath = parts[2] // Get 'rh-tabs/rh-tabs.js'
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Extracted relative path: %s", relativePath)
			} else {
				relativePath = filepath.Base(importPath)
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Fallback to base path: %s", relativePath)
			}
		} else {
			// Handle regular packages like some-package/path/file.js
			parts := strings.SplitN(importPath, "/", 2)
			helpers.SafeDebugLog("[MODULE_GRAPH] 📦 Regular package parts: %v", parts)
			if len(parts) >= 2 {
				relativePath = parts[1]
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Extracted relative path: %s", relativePath)
			} else {
				relativePath = filepath.Base(importPath)
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Fallback to base path: %s", relativePath)
			}
		}

		possiblePaths := []string{
			filepath.Join(mg.workspaceRoot, importPath),               // Try full import path
			filepath.Join(mg.workspaceRoot, relativePath),             // Try relative path
			filepath.Join(mg.workspaceRoot, "elements", relativePath), // Try in elements/ dir
			filepath.Join(mg.workspaceRoot, "src", relativePath),      // Try in src/ dir
			filepath.Join(mg.workspaceRoot, "lib", relativePath),      // Try in lib/ dir
		}

		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 Checking %d possible paths for import '%s'", len(possiblePaths), importPath)
		for i, path := range possiblePaths {
			helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Path %d: %s", i+1, path)
			if mg.fileExists(path) {
				filePaths = append(filePaths, path)
				helpers.SafeDebugLog("[MODULE_GRAPH] ✅ FOUND: %s", path)
			} else {
				helpers.SafeDebugLog("[MODULE_GRAPH] ❌ Not found: %s", path)
			}
		}
	}

	return filePaths
}

// fileExists checks if a file exists
func (mg *ModuleGraph) fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// processFile processes a single file and adds its exports to the module graph
func (mg *ModuleGraph) processFile(filePath string) error {
	helpers.SafeDebugLog("[MODULE_GRAPH] Processing file: %s", filePath)

	// Convert to relative path for the module graph
	relPath, err := filepath.Rel(mg.workspaceRoot, filePath)
	if err != nil {
		relPath = filePath // Fallback to absolute path
	}

	// Parse exports from this file
	parseStart := time.Now()
	parseErr := mg.parseFileExports(filePath, mg.workspaceRoot)
	mg.metrics.RecordDuration("lazy_file_parse_time", time.Since(parseStart))

	if parseErr != nil {
		mg.metrics.IncrementCounter("lazy_parse_errors")
		return parseErr
	}

	mg.metrics.IncrementCounter("lazy_files_processed")
	helpers.SafeDebugLog("[MODULE_GRAPH] Successfully processed file: %s", relPath)
	return nil
}

// processFileWithDependencies processes a file and recursively processes its dependencies
func (mg *ModuleGraph) processFileWithDependencies(filePath string, processedFiles map[string]bool) error {
	// Check if already processed to avoid infinite recursion
	if processedFiles[filePath] {
		helpers.SafeDebugLog("[MODULE_GRAPH] Skipping already processed file: %s", filePath)
		return nil
	}

	// Mark as processed
	processedFiles[filePath] = true
	helpers.SafeDebugLog("[MODULE_GRAPH] Processing file with dependencies: %s", filePath)

	// Process the file itself
	if err := mg.processFile(filePath); err != nil {
		return err
	}

	// Get dependencies (imports) from this file and process them recursively
	dependencies, err := mg.extractFileDependencies(filePath)
	if err != nil {
		helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to extract dependencies from %s: %v", filePath, err)
		// Continue without dependencies rather than failing completely
		return nil
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Found %d dependencies for %s: %v", len(dependencies), filePath, dependencies)

	// Recursively process each dependency
	for _, depPath := range dependencies {
		if err := mg.processFileWithDependencies(depPath, processedFiles); err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to process dependency %s: %v", depPath, err)
			// Continue with other dependencies
		}
	}

	return nil
}

// extractFileDependencies extracts the import dependencies from a file
func (mg *ModuleGraph) extractFileDependencies(filePath string) ([]string, error) {
	content, err := mg.fileParser.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Get TypeScript parser from pool
	parser := queries.RetrieveTypeScriptParser()
	defer queries.PutTypeScriptParser(parser)

	// Parse the content
	tree := parser.Parse(content, nil)
	if tree == nil {
		helpers.SafeDebugLog("[MODULE_GRAPH] DEBUG: Tree parsing failed for %s", filePath)
		return nil, nil // Skip files that fail to parse
	}
	defer tree.Close()
	helpers.SafeDebugLog("[MODULE_GRAPH] DEBUG: Successfully parsed tree for %s", filePath)

	// Get cached export matcher (which includes import queries)
	queryManager, err := queries.GetGlobalQueryManager()
	if err != nil {
		return nil, fmt.Errorf("failed to get global query manager for file %s: %w", filePath, err)
	}

	exportMatcher, err := queries.GetCachedQueryMatcher(queryManager, "typescript", "exports")
	if err != nil {
		return nil, fmt.Errorf("failed to get cached export matcher for file %s: %w", filePath, err)
	}

	var dependencies []string
	dependencySet := make(map[string]bool) // Use set to avoid duplicates

	// Process matches to extract import sources
	matchCount := 0
	for match := range exportMatcher.AllQueryMatches(tree.RootNode(), content) {
		matchCount++
		for _, capture := range match.Captures {
			captureName := exportMatcher.GetCaptureNameByIndex(capture.Index)
			captureText := strings.TrimSpace(capture.Node.Utf8Text(content))

			// Look for import sources
			if captureName == "import.source" || captureName == "import.dynamic.source" {
				if captureText != "" && !dependencySet[captureText] {
					helpers.SafeDebugLog("[MODULE_GRAPH] Processing import source: %s", captureText)
					// Resolve import path to file paths relative to current file
					resolvedPaths := mg.resolveImportPathToFilesRelativeTo(captureText, filePath)
					helpers.SafeDebugLog("[MODULE_GRAPH] Resolved import '%s' to %d paths: %v", captureText, len(resolvedPaths), resolvedPaths)

					// Add all resolved file paths as dependencies
					for _, resolvedPath := range resolvedPaths {
						if !dependencySet[resolvedPath] {
							dependencies = append(dependencies, resolvedPath)
							dependencySet[resolvedPath] = true
						}
					}
				}
			}
		}
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Extracted %d dependencies from %s", len(dependencies), filePath)
	return dependencies, nil
}

// findManifestModuleForImportPath finds a manifest module that matches an import path
// Returns the manifest module path if found, empty string otherwise
func (mg *ModuleGraph) findManifestModuleForImportPath(importPath string) string {
	if mg.manifestResolver == nil {
		return ""
	}

	// Use the ManifestResolver to find matching manifest modules
	manifestModules := mg.manifestResolver.FindManifestModulesForImportPath(importPath)
	if len(manifestModules) > 0 {
		// Return the first matching module - in practice, there should usually be only one
		return manifestModules[0]
	}

	return ""
}

// buildDependenciesForManifestModule builds the complete dependency tree for a manifest module
// by recursively tracking all imports and mapping them to manifest modules.
// This ensures transitive dependency resolution works correctly (e.g., importing my-tabs.js
// provides access to transitively imported elements like my-icon from my-tab.js).
func (mg *ModuleGraph) buildDependenciesForManifestModule(manifestModule string) error {
	visited := make(map[string]bool)
	return mg.buildDependenciesRecursive(manifestModule, visited)
}

// buildDependenciesRecursive recursively builds dependencies for a manifest module and all
// its transitive dependencies. The visited map prevents infinite recursion in circular
// dependency scenarios while ensuring complete dependency tree construction.
func (mg *ModuleGraph) buildDependenciesRecursive(manifestModule string, visited map[string]bool) error {
	if mg.workspaceRoot == "" {
		helpers.SafeDebugLog("[MODULE_GRAPH] No workspace root set, cannot build dependencies for manifest module: %s", manifestModule)
		return nil
	}

	// Check if we've already processed this module to prevent infinite recursion
	if visited[manifestModule] {
		helpers.SafeDebugLog("[MODULE_GRAPH] Skipping already processed manifest module: %s", manifestModule)
		return nil
	}
	visited[manifestModule] = true

	helpers.SafeDebugLog("[MODULE_GRAPH] Building dependencies for manifest module: %s", manifestModule)

	// Find the source file for this manifest module to parse its imports
	sourceFileCandidates := mg.findSourceFileForManifestModule(manifestModule)
	if len(sourceFileCandidates) == 0 {
		helpers.SafeDebugLog("[MODULE_GRAPH] No source file candidates found for manifest module: %s", manifestModule)
		return nil
	}

	// Try each source file candidate until we find one that exists and can be parsed
	for _, sourceFile := range sourceFileCandidates {
		if !mg.fileExists(sourceFile) {
			continue
		}

		helpers.SafeDebugLog("[MODULE_GRAPH] Found source file for manifest module '%s': %s", manifestModule, sourceFile)

		// Parse the source file to extract its import dependencies
		dependencies, err := mg.extractFileDependencies(sourceFile)
		if err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to extract dependencies from %s: %v", sourceFile, err)
			continue // Try next candidate
		}

		// Process each dependency and recursively build its dependency tree
		mg.processDependencies(manifestModule, dependencies, visited)

		helpers.SafeDebugLog("[MODULE_GRAPH] Successfully built dependencies for manifest module '%s' with %d dependencies", manifestModule, len(dependencies))
		return nil
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Could not find or parse source file for manifest module: %s", manifestModule)
	return nil
}

// processDependencies processes all file dependencies for a manifest module, mapping them
// to manifest modules and recursively building their dependency trees.
func (mg *ModuleGraph) processDependencies(manifestModule string, dependencies []string, visited map[string]bool) {
	for _, depFilePath := range dependencies {
		// Convert file path to manifest module path using ManifestResolver
		depManifestModule := mg.manifestResolver.GetManifestModulePath(depFilePath)
		if depManifestModule == "" {
			helpers.SafeDebugLog("[MODULE_GRAPH] No manifest module found for dependency file: %s", depFilePath)
			continue
		}

		// Track the module dependency relationship at manifest level
		mg.dependencyTracker.AddModuleDependency(manifestModule, depManifestModule)
		helpers.SafeDebugLog("[MODULE_GRAPH] Manifest module dependency: %s -> %s (via file %s)", manifestModule, depManifestModule, depFilePath)

		// Recursively build dependencies for the dependency module
		if err := mg.buildDependenciesRecursive(depManifestModule, visited); err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to recursively build dependencies for %s: %v", depManifestModule, err)
			// Continue with other dependencies rather than failing completely
		}
	}
}

// findSourceFileForManifestModule attempts to find the source file for a given manifest module
func (mg *ModuleGraph) findSourceFileForManifestModule(manifestModule string) []string {
	var candidates []string

	// Common patterns for finding source files from manifest modules:
	// 1. "rh-tabs/rh-tabs.js" -> "elements/rh-tabs/rh-tabs.ts" (Red Hat Design System)
	// 2. "rh-tabs/rh-tabs.js" -> "rh-tabs/rh-tabs.ts" (direct mapping)
	// 3. "rh-tabs/rh-tabs.js" -> "src/rh-tabs/rh-tabs.ts" (src directory)
	// 4. "rh-tabs/rh-tabs.js" -> "lib/rh-tabs/rh-tabs.ts" (lib directory)

	// Extract path without extension
	pathWithoutExt := strings.TrimSuffix(manifestModule, filepath.Ext(manifestModule))

	// Try different directory structures and file extensions
	prefixes := []string{"", "elements/", "src/", "lib/", "components/"}
	extensions := []string{".ts", ".js", ".mjs"}

	for _, prefix := range prefixes {
		for _, ext := range extensions {
			candidate := filepath.Join(mg.workspaceRoot, prefix+pathWithoutExt+ext)
			candidates = append(candidates, candidate)
		}
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Generated %d source file candidates for manifest module '%s': %v", len(candidates), manifestModule, candidates)
	return candidates
}
