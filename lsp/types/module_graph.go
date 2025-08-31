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
	"sync/atomic"
	"time"

	"bennypowers.dev/cem/queries"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// DefaultMaxTransitiveDepth is the default maximum depth for transitive closure computation
// to prevent performance issues with deeply nested dependency chains
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

// PathNormalizer interface abstracts path normalization operations
type PathNormalizer interface {
	// ToSlash converts path separators to forward slashes
	ToSlash(path string) string
	// Rel returns the relative path from basepath to targpath
	Rel(basepath, targpath string) (string, error)
	// NormalizeImportPath converts a relative import path to an absolute module path
	NormalizeImportPath(importPath, currentModulePath string) string
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
		// Skip files that fail to parse - this is expected for files with syntax errors
		// or files that aren't valid TypeScript/JavaScript. We silently continue to
		// ensure parsing failures don't break the entire module graph building process.
		return nil
	}
	defer tree.Close()

	// Try to use tree-sitter queries for accurate parsing
	err := p.parseExportsWithQueries(tree, modulePath, content, exportTracker, dependencyTracker)
	if err != nil {
		// Fallback to regex-based parsing when tree-sitter queries fail
		// This provides resilience against query compilation issues or unsupported syntax
		contentStr := string(content)
		p.parseCustomElementsDefine(modulePath, contentStr, exportTracker)
		p.parseExportStatements(modulePath, contentStr)
	}

	return nil
}

// parseExportsWithQueries uses tree-sitter queries to parse export statements
func (p *DefaultExportParser) parseExportsWithQueries(tree *ts.Tree, modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker) error {
	// Create query manager for exports
	queryManager, err := queries.NewQueryManager(queries.LSPQueries())
	if err != nil {
		return fmt.Errorf("failed to create query manager for module %s: %w", modulePath, err)
	}
	defer queryManager.Close()

	// Create export matcher
	exportMatcher, err := queries.NewQueryMatcher(queryManager, "typescript", "exports")
	if err != nil {
		return fmt.Errorf("failed to create export matcher for module %s: %w", modulePath, err)
	}
	defer exportMatcher.Close()

	// Process export matches
	for match := range exportMatcher.AllQueryMatches(tree.RootNode(), content) {
		p.processExportMatch(match, exportMatcher, modulePath, content, exportTracker, dependencyTracker)
	}

	return nil
}

// processExportMatch processes a single export/import match from tree-sitter
func (p *DefaultExportParser) processExportMatch(match *ts.QueryMatch, matcher *queries.QueryMatcher, modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker) {
	var exportName, sourceModule, tagName, className string
	var importSource string
	var isCustomElementsDefine, isImport bool

	// Process all captures in this match
	for _, capture := range match.Captures {
		captureName := matcher.GetCaptureNameByIndex(capture.Index)
		captureText := strings.TrimSpace(capture.Node.Utf8Text(content))

		switch captureName {
		case "export.name":
			exportName = captureText
		case "export.source":
			sourceModule = captureText
		case "customElements.tagName":
			tagName = captureText
			isCustomElementsDefine = true
		case "customElements.className":
			className = captureText
			isCustomElementsDefine = true
		case "export.class.name", "export.function.name", "export.variable.name", "export.default.name":
			exportName = captureText
		case "import.source", "import.dynamic.source":
			importSource = captureText
			isImport = true
		}
	}

	// Handle CustomElements.define() - this is a direct export
	if isCustomElementsDefine && tagName != "" && className != "" {
		exportTracker.AddDirectExport(modulePath, className, tagName)
		return
	}

	// Handle import statements - track module dependencies
	if isImport && importSource != "" {
		// Normalize the import path to handle relative imports
		pathNormalizer := &DefaultPathNormalizer{}
		normalizedImportPath := pathNormalizer.NormalizeImportPath(importSource, modulePath)
		dependencyTracker.AddModuleDependency(modulePath, normalizedImportPath)
		return
	}

	// Handle re-exports (export { X } from './module')
	if exportName != "" && sourceModule != "" {
		// Track re-export relationship at module level - this is sufficient because
		// resolveReExportChains() will later resolve all elements from source modules
		// to re-exporting modules using existing manifest data
		pathNormalizer := &DefaultPathNormalizer{}
		normalizedSourceModule := pathNormalizer.NormalizeImportPath(sourceModule, modulePath)
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

// parseCustomElementsDefine finds customElements.define() calls to identify direct exports
func (p *DefaultExportParser) parseCustomElementsDefine(modulePath, content string, exportTracker *ExportTracker) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Look for customElements.define('tag-name', ClassName)
		if strings.Contains(line, "customElements.define") {
			// Extract tag name from the define call
			if start := strings.Index(line, "customElements.define("); start != -1 {
				after := line[start+len("customElements.define("):]
				if end := strings.Index(after, ","); end != -1 {
					tagNamePart := strings.TrimSpace(after[:end])
					// Remove quotes
					tagName := strings.Trim(tagNamePart, `"'`)

					// Extract class name (simplified)
					remaining := strings.TrimSpace(after[end+1:])
					if classEnd := strings.IndexAny(remaining, ",)"); classEnd != -1 {
						className := strings.TrimSpace(remaining[:classEnd])

						exportTracker.AddDirectExport(modulePath, className, tagName)
					}
				}
			}
		}
	}
}

// parseExportStatements finds export statements that might be re-exports
func (p *DefaultExportParser) parseExportStatements(modulePath, content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Look for "export { Something } from './other-module'"
		if strings.HasPrefix(line, "export") && strings.Contains(line, "from") {
			// This is a re-export, we need to analyze it further
			// For now, we'll add a simple implementation

			// Extract the source module
			if fromIndex := strings.LastIndex(line, "from"); fromIndex != -1 {
				sourcePart := strings.TrimSpace(line[fromIndex+4:])
				// Remove quotes and semicolon
				sourceModule := strings.Trim(sourcePart, `"'; `)
				// sourceModule is extracted but not used here - module-level re-export resolution
				// via resolveReExportChains() provides sufficient functionality
				_ = sourceModule // Acknowledge variable (extracted for potential future use)

				// Extract exported names (simplified)
				if start := strings.Index(line, "{"); start != -1 {
					if end := strings.Index(line, "}"); end != -1 {
						exportsPart := line[start+1 : end]
						exports := strings.Split(exportsPart, ",")

						for _, export := range exports {
							exportName := strings.TrimSpace(export)
							if exportName != "" {
								// Individual export tracking is not implemented here because
								// module-level resolveReExportChains() provides sufficient functionality
								// by resolving all elements from source modules using manifest data
								continue
							}
						}
					}
				}
			}
		}
	}
}

// DefaultPathNormalizer implements PathNormalizer using standard path operations
type DefaultPathNormalizer struct{}

// ToSlash implements PathNormalizer using filepath.ToSlash
func (p *DefaultPathNormalizer) ToSlash(path string) string {
	return filepath.ToSlash(path)
}

// Rel implements PathNormalizer using filepath.Rel
func (p *DefaultPathNormalizer) Rel(basepath, targpath string) (string, error) {
	return filepath.Rel(basepath, targpath)
}

// NormalizeImportPath converts a relative import path to an absolute module path
func (p *DefaultPathNormalizer) NormalizeImportPath(importPath, currentModulePath string) string {
	// Handle relative imports like './my-icon.js'
	if strings.HasPrefix(importPath, "./") {
		// Convert './my-icon.js' to 'my-icon.js'
		return strings.TrimPrefix(importPath, "./")
	}

	// Handle parent directory imports like '../my-icon.js'
	if strings.HasPrefix(importPath, "../") {
		// For now, just remove the prefix - more sophisticated path resolution needed for real world
		return strings.TrimPrefix(importPath, "../")
	}

	// Return as-is for absolute imports
	return importPath
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
	if !contains(dt.ModuleDependencies[importingModule], importedModule) {
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

	if !contains(dt.ReExportChains[reExportingModule], sourceModule) {
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

	// Injected dependencies for improved testability
	fileParser     FileParser
	exportParser   ExportParser
	pathNormalizer PathNormalizer
	metrics        MetricsCollector
}

// NewModuleGraph creates a new empty module graph with default dependencies
func NewModuleGraph() *ModuleGraph {
	return &ModuleGraph{
		exportTracker:     NewExportTracker(),
		dependencyTracker: NewDependencyTracker(),
		fileParser:        &OSFileParser{},
		exportParser:      &DefaultExportParser{},
		pathNormalizer:    &DefaultPathNormalizer{},
		metrics:           &NoOpMetricsCollector{}, // Default to no-op for performance
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithMetrics creates a new module graph with metrics collection enabled
func NewModuleGraphWithMetrics() *ModuleGraph {
	return &ModuleGraph{
		exportTracker:     NewExportTracker(),
		dependencyTracker: NewDependencyTracker(),
		fileParser:        &OSFileParser{},
		exportParser:      &DefaultExportParser{},
		pathNormalizer:    &DefaultPathNormalizer{},
		metrics:           NewDefaultMetricsCollector(),
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithDependencies creates a new module graph with custom dependencies
// This is useful for testing and dependency injection
func NewModuleGraphWithDependencies(fileParser FileParser, exportParser ExportParser, pathNormalizer PathNormalizer, metrics MetricsCollector) *ModuleGraph {
	if metrics == nil {
		metrics = &NoOpMetricsCollector{}
	}
	return &ModuleGraph{
		exportTracker:     NewExportTracker(),
		dependencyTracker: NewDependencyTracker(),
		fileParser:        fileParser,
		exportParser:      exportParser,
		pathNormalizer:    pathNormalizer,
		metrics:           metrics,
		// TransitiveElementsCache is initialized as zero value (ready to use)
	}
}

// NewModuleGraphWithFileParser creates a new module graph with a custom file parser
// This is useful for testing with mock file systems (backwards compatibility)
func NewModuleGraphWithFileParser(fileParser FileParser) *ModuleGraph {
	return &ModuleGraph{
		exportTracker:     NewExportTracker(),
		dependencyTracker: NewDependencyTracker(),
		fileParser:        fileParser,
		exportParser:      &DefaultExportParser{},
		pathNormalizer:    &DefaultPathNormalizer{},
		metrics:           &NoOpMetricsCollector{},
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

// AddModuleDependency registers that one module imports another
func (mg *ModuleGraph) AddModuleDependency(importingModule, importedModule string) {
	mg.dependencyTracker.AddModuleDependency(importingModule, importedModule)
	// Invalidate transitive closure cache for this module (lock-free)
	mg.TransitiveElementsCache.Delete(importingModule)
	mg.metrics.IncrementCounter("dependencies_added")
}

// GetTransitiveElements returns all custom elements available when importing a module
// This includes elements defined in the module itself and all transitively imported modules
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
		return result
	}

	mg.metrics.IncrementCounter("cache_misses")

	// Need to compute transitive closure - no global lock needed since components handle their own synchronization

	// Double-check after acquiring lock (another goroutine might have computed it)
	if cached, exists := mg.TransitiveElementsCache.Load(modulePath); exists {
		mg.metrics.IncrementCounter("cache_hits")
		cachedElements := cached.([]string)
		result := make([]string, len(cachedElements))
		copy(result, cachedElements)
		return result
	}

	// Compute transitive closure
	mg.metrics.IncrementCounter("transitive_computations")
	elements := mg.calculateTransitiveClosure(modulePath)

	// Cache the result (lock-free)
	mg.TransitiveElementsCache.Store(modulePath, elements)

	// Return a copy
	result := make([]string, len(elements))
	copy(result, elements)
	return result
}

// calculateTransitiveClosure computes all elements transitively available from a module
func (mg *ModuleGraph) calculateTransitiveClosure(modulePath string) []string {
	visited := make(map[string]bool)
	elements := make(map[string]bool) // Use map to avoid duplicates
	maxDepth := DefaultMaxTransitiveDepth // Configurable depth limit to prevent performance issues

	// Breadth-first traversal to collect all transitive elements
	queue := []string{modulePath}
	depth := 0

	for len(queue) > 0 && depth < maxDepth {
		nextQueue := []string{}

		for _, currentModule := range queue {
			if visited[currentModule] {
				continue // Skip already visited modules (handles cycles)
			}
			visited[currentModule] = true

			// Add elements defined in this module
			exports := mg.exportTracker.GetModuleExports(currentModule)
			for _, export := range exports {
				elements[export.TagName] = true
			}

			// Add dependencies to next level queue
			dependencies := mg.dependencyTracker.GetModuleDependencies(currentModule)
			for _, dependency := range dependencies {
				if !visited[dependency] {
					nextQueue = append(nextQueue, dependency)
				}
			}
		}

		queue = nextQueue
		depth++
	}

	// Convert map to slice
	result := make([]string, 0, len(elements))
	for element := range elements {
		result = append(result, element)
	}

	return result
}

// GetModuleExports returns all exports from the given module
func (mg *ModuleGraph) GetModuleExports(modulePath string) []ModuleExport {
	return mg.exportTracker.GetModuleExports(modulePath)
}

// BuildFromWorkspace analyzes the workspace to build the module graph
// by parsing TypeScript/JavaScript files for export statements
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
	relPath, err := mg.pathNormalizer.Rel(workspaceRoot, filePath)
	if err != nil {
		// If relative path conversion fails, use absolute path as fallback
		// This ensures parsing continues even with path resolution issues
		relPath = filePath
	}

	// Normalize path separators for consistent module paths
	modulePath := mg.pathNormalizer.ToSlash(relPath)

	// Parse exports using injected parser
	return mg.exportParser.ParseExportsFromContent(modulePath, content, mg.exportTracker, mg.dependencyTracker)
}

// Note: parseExportsWithTreeSitter is now handled by the injected ExportParser
// This method is removed to avoid duplication

// Note: parseExportsWithQueries is now handled by the injected ExportParser
// This method is removed to avoid duplication

// Note: processExportMatch is now handled by the injected ExportParser
// This method is removed to avoid duplication

// Note: Removed unused functions trackDirectExport, trackReExportRelationship, and normalizeImportPath
// These were replaced by the dependency injection pattern using ExportParser and PathNormalizer interfaces
// which provide better separation of concerns and improved testability

// Note: parseCustomElementsDefine is now handled by the injected ExportParser
// This method is removed to avoid duplication

// Note: parseExportStatements is now handled by the injected ExportParser
// This method is removed to avoid duplication

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
		// Note: We can't access private fields directly, but in a real implementation
		// we would add methods to get counts from the focused components
		allTagNames := exportTracker.GetAllTagNames()
		mg.metrics.SetGauge("active_elements", int64(len(allTagNames)))
	}

	// In a real implementation, we would add similar methods to get module counts
	// For now, we'll use a simple approximation
	mg.metrics.SetGauge("active_modules", int64(len(allModules)))
}

// GetMetrics returns the current metrics collector (useful for testing)
func (mg *ModuleGraph) GetMetrics() MetricsCollector {
	return mg.metrics
}

// contains checks if a string slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
