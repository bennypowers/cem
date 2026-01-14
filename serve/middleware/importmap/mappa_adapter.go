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

package importmap

import (
	"maps"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/middleware/types"

	mappafs "bennypowers.dev/mappa/fs"
	"bennypowers.dev/mappa/importmap"
	"bennypowers.dev/mappa/packagejson"
	"bennypowers.dev/mappa/resolve"
	"bennypowers.dev/mappa/resolve/local"
)

// fsAdapter wraps cem's platform.FileSystem to satisfy mappa's fs.FileSystem interface.
// Both interfaces are designed to be duck-type compatible, so this adapter just
// provides the type assertion bridge.
type fsAdapter struct {
	platform.FileSystem
}

// Ensure fsAdapter implements mappa's FileSystem interface at compile time.
var _ mappafs.FileSystem = (*fsAdapter)(nil)

// mappaLogger adapts cem's Logger to mappa's resolve.Logger interface.
type mappaLogger struct {
	logger types.Logger
}

func (l *mappaLogger) Warning(format string, args ...any) {
	if l.logger != nil {
		l.logger.Warning(format, args...)
	}
}

func (l *mappaLogger) Debug(format string, args ...any) {
	if l.logger != nil {
		l.logger.Debug(format, args...)
	}
}

// wrapFS wraps a cem FileSystem for use with mappa.
func wrapFS(fs platform.FileSystem) mappafs.FileSystem {
	return &fsAdapter{fs}
}

// wrapLogger wraps a cem Logger for use with mappa.
func wrapLogger(logger types.Logger) resolve.Logger {
	if logger == nil {
		return nil
	}
	return &mappaLogger{logger: logger}
}

// convertMappaImportMap converts a mappa ImportMap to cem's ImportMap type.
// Both types have the same structure, so this is a field-by-field copy.
func convertMappaImportMap(m *importmap.ImportMap) *ImportMap {
	if m == nil {
		return nil
	}
	result := &ImportMap{
		Imports: make(map[string]string, len(m.Imports)),
	}
	maps.Copy(result.Imports, m.Imports)
	if m.Scopes != nil {
		result.Scopes = make(map[string]map[string]string, len(m.Scopes))
		for scopeKey, scopeMap := range m.Scopes {
			result.Scopes[scopeKey] = make(map[string]string, len(scopeMap))
			maps.Copy(result.Scopes[scopeKey], scopeMap)
		}
	}
	return result
}

// convertToMappaImportMap converts a cem ImportMap to mappa's ImportMap type.
func convertToMappaImportMap(m *ImportMap) *importmap.ImportMap {
	if m == nil {
		return nil
	}
	result := &importmap.ImportMap{
		Imports: make(map[string]string, len(m.Imports)),
	}
	maps.Copy(result.Imports, m.Imports)
	if m.Scopes != nil {
		result.Scopes = make(map[string]map[string]string, len(m.Scopes))
		for scopeKey, scopeMap := range m.Scopes {
			result.Scopes[scopeKey] = make(map[string]string, len(scopeMap))
			maps.Copy(result.Scopes[scopeKey], scopeMap)
		}
	}
	return result
}

// convertWorkspacePackages converts cem's middleware.WorkspacePackage to mappa's resolve.WorkspacePackage.
func convertWorkspacePackages(packages []workspacePackageInfo) []resolve.WorkspacePackage {
	result := make([]resolve.WorkspacePackage, len(packages))
	for i, pkg := range packages {
		result[i] = resolve.WorkspacePackage{
			Name: pkg.Name,
			Path: pkg.Path,
		}
	}
	return result
}

// workspacePackageInfo holds the minimal info needed for workspace package conversion.
// This is used internally to bridge between cem's middleware.WorkspacePackage and mappa's resolve.WorkspacePackage.
type workspacePackageInfo struct {
	Name string
	Path string
}

// newMappaResolver creates a new mappa resolver configured for cem's use case.
// Uses an in-memory cache for parsed package.json files to avoid repeated parsing.
func newMappaResolver(fs platform.FileSystem, logger types.Logger, cache packagejson.Cache) *local.Resolver {
	resolver := local.New(wrapFS(fs), wrapLogger(logger))
	if cache != nil {
		resolver = resolver.WithPackageCache(cache)
	}
	return resolver
}

// resolveWithMappa uses mappa's resolver for import map generation.
// This is the main integration point between cem and mappa.
//
// For hot-reload performance, consider using ResolveWithGraph() and ResolveIncremental()
// which only re-resolve changed packages and their dependents.
func resolveWithMappa(
	rootDir string,
	fs platform.FileSystem,
	logger types.Logger,
	inputMap *ImportMap,
	workspacePackages []workspacePackageInfo,
	includeRootExports bool,
) (*ImportMap, error) {
	// Create a memory cache for this resolution to avoid repeated parsing
	// of the same package.json files (especially important for workspace mode)
	cache := packagejson.NewMemoryCache()
	resolver := newMappaResolver(fs, logger, cache)

	// Configure workspace packages if provided
	if len(workspacePackages) > 0 {
		resolver = resolver.WithWorkspacePackages(convertWorkspacePackages(workspacePackages))
	}

	// Include root package exports for dev server self-imports
	if includeRootExports {
		resolver = resolver.WithIncludeRootExports()
	}

	// Merge with input map if provided
	if inputMap != nil {
		resolver = resolver.WithInputMap(convertToMappaImportMap(inputMap))
	}

	// Resolve the import map
	result, err := resolver.Resolve(rootDir)
	if err != nil {
		return nil, err
	}

	return convertMappaImportMap(result), nil
}

// resolveWithMappaGraph uses mappa's resolver and returns both import map and dependency graph.
// This is used for initial resolution when incremental updates will be needed.
func resolveWithMappaGraph(
	rootDir string,
	fs platform.FileSystem,
	logger types.Logger,
	inputMap *ImportMap,
	workspacePackages []workspacePackageInfo,
	includeRootExports bool,
	cache packagejson.Cache,
) (*IncrementalResult, error) {
	resolver := newMappaResolver(fs, logger, cache)

	// Configure workspace packages if provided (overrides auto-discovery)
	if len(workspacePackages) > 0 {
		resolver = resolver.WithWorkspacePackages(convertWorkspacePackages(workspacePackages))
	}

	// Include root package exports for dev server self-imports
	if includeRootExports {
		resolver = resolver.WithIncludeRootExports()
	}

	// Merge with input map if provided
	if inputMap != nil {
		resolver = resolver.WithInputMap(convertToMappaImportMap(inputMap))
	}

	// Resolve with graph for incremental updates
	result, err := resolver.ResolveWithGraph(rootDir)
	if err != nil {
		return nil, err
	}

	return &IncrementalResult{
		ImportMap: convertMappaImportMap(result.ImportMap),
		Graph:     &DependencyGraph{graph: result.DependencyGraph},
	}, nil
}

// resolveIncrementalWithMappa performs an incremental update using mappa's resolver.
// Only changed packages and their dependents are re-resolved.
func resolveIncrementalWithMappa(
	rootDir string,
	fs platform.FileSystem,
	logger types.Logger,
	update IncrementalUpdate,
	cache packagejson.Cache,
) (*IncrementalResult, error) {
	resolver := newMappaResolver(fs, logger, cache)

	// Extract mappa's dependency graph from our wrapper
	var mappaGraph *resolve.DependencyGraph
	if update.PreviousGraph != nil && update.PreviousGraph.graph != nil {
		mappaGraph = update.PreviousGraph.graph.(*resolve.DependencyGraph)
	}

	// Build mappa's incremental update
	mappaUpdate := resolve.IncrementalUpdate{
		ChangedPackages: update.ChangedPackages,
		PreviousMap:     convertToMappaImportMap(update.PreviousMap),
		PreviousGraph:   mappaGraph,
	}

	// Perform incremental resolution
	result, err := resolver.ResolveIncremental(rootDir, mappaUpdate)
	if err != nil {
		return nil, err
	}

	return &IncrementalResult{
		ImportMap: convertMappaImportMap(result.ImportMap),
		Graph:     &DependencyGraph{graph: result.DependencyGraph},
	}, nil
}
