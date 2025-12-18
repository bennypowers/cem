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
	"context"
	"errors"
	"maps"
	"runtime"
	"sync"

	"bennypowers.dev/cem/internal/logging"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"

	ts "github.com/tree-sitter/go-tree-sitter"
)

// ModuleProcessorFunc represents a function that processes a single module
type ModuleProcessorFunc func(job processJob, qm *Q.QueryManager, parser *ts.Parser, depTracker *FileDependencyTracker, cssCache CssCache) (*M.Module, map[string]string, map[string]string, map[string]importInfo, *LogCtx, error)

// ModuleProcessorSimpleFunc represents a function that processes a single module (without dependency tracking)
type ModuleProcessorSimpleFunc func(job processJob, qm *Q.QueryManager, parser *ts.Parser, cssCache CssCache) (*M.Module, map[string]string, map[string]string, map[string]importInfo, *LogCtx, error)

// ModuleBatchProcessor handles batch/parallel processing of modules with common patterns.
// Abstracts the worker pool pattern used throughout the codebase for module processing.
// Supports both dependency-tracking and simple processing modes.
//
// Callsites:
// - session_core.go:251 (NewModuleBatchProcessor for full generation)
// - generate.go:115 (NewModuleBatchProcessor for simple processing)
// - session_incremental.go:143 (NewModuleBatchProcessor for incremental processing)
//
// Worker Management: Optimizes worker count based on job size (avoids over-allocation)
type ModuleBatchProcessor struct {
	numWorkers   int
	queryManager *Q.QueryManager
	depTracker   *FileDependencyTracker
	cssCache     CssCache
}

// ProcessingResult holds the results of parallel module processing.
// Aggregates modules, logs, aliases, and errors from all worker threads.
//
// Usage: Returned by all ModuleBatchProcessor.Process* methods
type ProcessingResult struct {
	Modules     []M.Module
	Logs        []*LogCtx
	Aliases     map[string]string
	TypeAliases moduleTypeAliasesMap
	Imports     moduleImportsMap
	Errors      error
}

// moduleTypeAliasesMap tracks type aliases for each module (key: module path)
type moduleTypeAliasesMap map[string]map[string]string

// NewModuleBatchProcessor creates a new batch/parallel processor.
//
// Callsites:
// - session_core.go:251 (with dependency tracking)
// - generate.go:115 (without dependency tracking, depTracker=nil)
// - session_incremental.go:143 (with dependency tracking)
//
// Parameters:
// - queryManager: Shared tree-sitter queries (expensive to create)
// - depTracker: Optional dependency tracker (nil for simple processing)
// - cssCache: CSS parsing cache for performance
func NewModuleBatchProcessor(queryManager *Q.QueryManager, depTracker *FileDependencyTracker, cssCache CssCache) *ModuleBatchProcessor {
	return &ModuleBatchProcessor{
		numWorkers:   runtime.NumCPU(),
		queryManager: queryManager,
		depTracker:   depTracker,
		cssCache:     cssCache,
	}
}

// SetWorkerCount allows overriding the number of workers
func (mbp *ModuleBatchProcessor) SetWorkerCount(count int) {
	if count > 0 {
		mbp.numWorkers = count
	}
}

// WorkerCount returns the configured number of workers
func (mbp *ModuleBatchProcessor) WorkerCount() int {
	return mbp.numWorkers
}

// ProcessModules processes modules in parallel using the provided processor function
func (mbp *ModuleBatchProcessor) ProcessModules(
	ctx context.Context,
	jobs []processJob,
	processor ModuleProcessorFunc,
) ProcessingResult {
	return mbp.processModulesInternal(ctx, jobs, func(job processJob, qm *Q.QueryManager, parser *ts.Parser) (*M.Module, map[string]string, map[string]string, map[string]importInfo, *LogCtx, error) {
		return processor(job, qm, parser, mbp.depTracker, mbp.cssCache)
	})
}

// ProcessModulesSimple processes modules in parallel without dependency tracking
func (mbp *ModuleBatchProcessor) ProcessModulesSimple(
	ctx context.Context,
	jobs []processJob,
	processor ModuleProcessorSimpleFunc,
) ProcessingResult {
	return mbp.processModulesInternal(ctx, jobs, func(job processJob, qm *Q.QueryManager, parser *ts.Parser) (*M.Module, map[string]string, map[string]string, map[string]importInfo, *LogCtx, error) {
		return processor(job, qm, parser, mbp.cssCache)
	})
}

// processModulesInternal handles the common parallel processing logic
func (mbp *ModuleBatchProcessor) processModulesInternal(
	ctx context.Context,
	jobs []processJob,
	processor func(processJob, *Q.QueryManager, *ts.Parser) (*M.Module, map[string]string, map[string]string, map[string]importInfo, *LogCtx, error),
) ProcessingResult {
	// Check for cancellation
	select {
	case <-ctx.Done():
		return ProcessingResult{Errors: ctx.Err()}
	default:
	}

	if len(jobs) == 0 {
		return ProcessingResult{
			Modules:     make([]M.Module, 0),
			Logs:        make([]*LogCtx, 0),
			Aliases:     make(map[string]string),
			TypeAliases: make(moduleTypeAliasesMap),
			Imports:     make(moduleImportsMap),
		}
	}

	// Optimize worker count for small job sets
	numWorkers := min(len(jobs), mbp.numWorkers)
	logging.Info("Starting generation with %d workers", numWorkers)

	// Initialize result collection
	var wg sync.WaitGroup
	var aliasesMu sync.Mutex
	var typeAliasesMu sync.Mutex
	var importsMu sync.Mutex
	var modulesMu sync.Mutex
	var errsMu sync.Mutex
	var logsMu sync.Mutex

	errsList := make([]error, 0)
	logs := make([]*LogCtx, 0, len(jobs))
	aliases := make(map[string]string)
	typeAliasesMap := make(moduleTypeAliasesMap)
	importsMap := make(moduleImportsMap)
	var modules []M.Module

	jobsChan := make(chan processJob, len(jobs))

	// Fill jobs channel
	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)

	// Start workers
	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			parser := Q.RetrieveTypeScriptParser()
			defer Q.PutTypeScriptParser(parser)

			for job := range jobsChan {
				// Check for cancellation
				select {
				case <-ctx.Done():
					return
				default:
				}

				module, tagAliases, typeAliases, imports, logger, err := processor(job, mbp.queryManager, parser)

				// Handle errors
				if err != nil {
					errsMu.Lock()
					errsList = append(errsList, err)
					errsMu.Unlock()
				}

				// Collect logs (always save duration for bar chart)
				logsMu.Lock()
				logs = append(logs, logger)
				logsMu.Unlock()

				// Merge tag aliases thread-safely
				if len(tagAliases) > 0 {
					aliasesMu.Lock()
					maps.Copy(aliases, tagAliases)
					aliasesMu.Unlock()
				}

				// Collect type aliases thread-safely
				if module != nil && len(typeAliases) > 0 {
					typeAliasesMu.Lock()
					typeAliasesMap[module.Path] = typeAliases
					typeAliasesMu.Unlock()
				}

				// Collect imports thread-safely
				if module != nil && len(imports) > 0 {
					importsMu.Lock()
					importsMap[module.Path] = imports
					importsMu.Unlock()
				}

				// Collect modules
				if module != nil {
					modulesMu.Lock()
					modules = append(modules, *module)
					modulesMu.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	// Combine errors
	var errs error
	if len(errsList) > 0 {
		errs = errors.Join(errsList...)
	}

	return ProcessingResult{
		Modules:     modules,
		Logs:        logs,
		Aliases:     aliases,
		TypeAliases: typeAliasesMap,
		Imports:     importsMap,
		Errors:      errs,
	}
}
