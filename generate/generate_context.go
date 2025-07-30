/*
Copyright © 2025 Benny Powers

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
package generate

import (
	"fmt"

	Q "bennypowers.dev/cem/generate/queries"
	W "bennypowers.dev/cem/workspace"
)

// GenerateContext wraps WorkspaceContext with setup objects for generation.
// This wrapper provides clean separation of concerns and enables dependency injection
// for testing and future extensibility. It composes setup objects needed for
// manifest generation while maintaining backward compatibility.
//
// Design rationale (from DEVELOPMENT.md):
// - Type safety for setup object management
// - Clean separation of concerns between workspace and generation setup
// - Composability for testing and future features
// - Dependency injection support for different cache strategies
//
// Usage:
// - Primary context object for GenerateSession initialization
// - Testing with mock implementations
// - Future extension points for additional setup objects
type GenerateContext struct {
	W.WorkspaceContext                        // Embedded workspace context
	cssCache           CssCache               // CSS parsing cache for performance
	queryManager       *Q.QueryManager        // Tree-sitter query manager (expensive to initialize)
	depTracker         *FileDependencyTracker // File dependency tracker for incremental builds
}

// NewGenerateContext creates a new generate context with initialized components.
// This is expensive due to QueryManager initialization (tree-sitter query loading)
// and should be done once per generation session.
//
// Performance: Expensive operation (~10-50ms) due to tree-sitter query compilation
func NewGenerateContext(ctx W.WorkspaceContext) (*GenerateContext, error) {
	qm, err := Q.NewQueryManager()
	if err != nil {
		return nil, fmt.Errorf("initialize QueryManager: %w", err)
	}

	return &GenerateContext{
		WorkspaceContext: ctx,
		cssCache:         NewCssParseCache(),
		queryManager:     qm,
		depTracker:       NewFileDependencyTracker(ctx),
	}, nil
}

// Close releases resources held by the generate context
func (gsc *GenerateContext) Close() {
	if gsc.queryManager != nil {
		gsc.queryManager.Close()
	}
}

// CssCache returns the CSS cache for dependency injection
func (gsc *GenerateContext) CssCache() CssCache {
	return gsc.cssCache
}

// QueryManager returns the query manager for tree-sitter operations
func (gsc *GenerateContext) QueryManager() *Q.QueryManager {
	return gsc.queryManager
}

// DependencyTracker returns the file dependency tracker for incremental builds
func (gsc *GenerateContext) DependencyTracker() *FileDependencyTracker {
	return gsc.depTracker
}

// WithCssCache returns a new generate context with a different CSS cache.
// This enables dependency injection and testing with mock implementations.
func (gsc *GenerateContext) WithCssCache(cache CssCache) *GenerateContext {
	return &GenerateContext{
		WorkspaceContext: gsc.WorkspaceContext,
		cssCache:         cache,
		queryManager:     gsc.queryManager,
		depTracker:       gsc.depTracker,
	}
}

// WithDependencyTracker returns a new generate context with a different dependency tracker.
// This enables testing with mock implementations.
func (gsc *GenerateContext) WithDependencyTracker(tracker *FileDependencyTracker) *GenerateContext {
	return &GenerateContext{
		WorkspaceContext: gsc.WorkspaceContext,
		cssCache:         gsc.cssCache,
		queryManager:     gsc.queryManager,
		depTracker:       tracker,
	}
}
