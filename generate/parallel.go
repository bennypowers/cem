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

	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
)

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

// ModuleProcessorFunc represents a function that processes a single module
type ModuleProcessorFunc func(job processJob, qm *Q.QueryManager, parser *ts.Parser, depTracker *FileDependencyTracker, cssCache CssCache) (*M.Module, map[string]string, *LogCtx, error)

// ModuleProcessorSimpleFunc represents a function that processes a single module (without dependency tracking)
type ModuleProcessorSimpleFunc func(job processJob, qm *Q.QueryManager, parser *ts.Parser, cssCache CssCache) (*M.Module, map[string]string, *LogCtx, error)

// ParallelModuleProcessor handles parallel processing of modules with common patterns
type ParallelModuleProcessor struct {
	numWorkers   int
	queryManager *Q.QueryManager
	depTracker   *FileDependencyTracker
	cssCache     CssCache
}

// ProcessingResult holds the results of parallel module processing
type ProcessingResult struct {
	Modules []M.Module
	Logs    []*LogCtx
	Aliases map[string]string
	Errors  error
}

// NewParallelModuleProcessor creates a new parallel processor
func NewParallelModuleProcessor(queryManager *Q.QueryManager, depTracker *FileDependencyTracker, cssCache CssCache) *ParallelModuleProcessor {
	return &ParallelModuleProcessor{
		numWorkers:   runtime.NumCPU(),
		queryManager: queryManager,
		depTracker:   depTracker,
		cssCache:     cssCache,
	}
}

// SetWorkerCount allows overriding the number of workers
func (pmp *ParallelModuleProcessor) SetWorkerCount(count int) {
	if count > 0 {
		pmp.numWorkers = count
	}
}

// ProcessModules processes modules in parallel using the provided processor function
func (pmp *ParallelModuleProcessor) ProcessModules(
	ctx context.Context,
	jobs []processJob,
	processor ModuleProcessorFunc,
) ProcessingResult {
	return pmp.processModulesInternal(ctx, jobs, func(job processJob, qm *Q.QueryManager, parser *ts.Parser) (*M.Module, map[string]string, *LogCtx, error) {
		return processor(job, qm, parser, pmp.depTracker, pmp.cssCache)
	})
}

// ProcessModulesSimple processes modules in parallel without dependency tracking
func (pmp *ParallelModuleProcessor) ProcessModulesSimple(
	ctx context.Context,
	jobs []processJob,
	processor ModuleProcessorSimpleFunc,
) ProcessingResult {
	return pmp.processModulesInternal(ctx, jobs, func(job processJob, qm *Q.QueryManager, parser *ts.Parser) (*M.Module, map[string]string, *LogCtx, error) {
		return processor(job, qm, parser, pmp.cssCache)
	})
}

// processModulesInternal handles the common parallel processing logic
func (pmp *ParallelModuleProcessor) processModulesInternal(
	ctx context.Context,
	jobs []processJob,
	processor func(processJob, *Q.QueryManager, *ts.Parser) (*M.Module, map[string]string, *LogCtx, error),
) ProcessingResult {
	// Check for cancellation
	select {
	case <-ctx.Done():
		return ProcessingResult{Errors: ctx.Err()}
	default:
	}

	if len(jobs) == 0 {
		return ProcessingResult{
			Modules: make([]M.Module, 0),
			Logs:    make([]*LogCtx, 0),
			Aliases: make(map[string]string),
		}
	}

	// Optimize worker count for small job sets
	numWorkers := pmp.numWorkers
	if len(jobs) < numWorkers {
		numWorkers = len(jobs)
	}

	pterm.Info.Printf("Starting Generation with %d workers\n", numWorkers)

	// Initialize result collection
	var wg sync.WaitGroup
	var aliasesMu sync.Mutex
	var modulesMu sync.Mutex
	var errsMu sync.Mutex
	var logsMu sync.Mutex
	
	errsList := make([]error, 0)
	logs := make([]*LogCtx, 0, len(jobs))
	aliases := make(map[string]string)
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
			parser := Q.GetTypeScriptParser()
			defer Q.PutTypeScriptParser(parser)
			
			for job := range jobsChan {
				// Check for cancellation
				select {
				case <-ctx.Done():
					return
				default:
				}

				module, tagAliases, logger, err := processor(job, pmp.queryManager, parser)
				
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

				// Merge aliases thread-safely
				if len(tagAliases) > 0 {
					aliasesMu.Lock()
					maps.Copy(aliases, tagAliases)
					aliasesMu.Unlock()
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
		Modules: modules,
		Logs:    logs,
		Aliases: aliases,
		Errors:  errs,
	}
}