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
	"crypto/sha256"
	"io"
	"os"
	"sync"
	"time"

	W "bennypowers.dev/cem/workspace"
)

// Type aliases for dependency tracking maps
type FileHashMap map[string][32]byte
type ModuleDependencyMap map[string]*ModuleDependencies
type CssReverseDepMap map[string][]string

// FileDependencyTracker tracks file dependencies and content hashes for incremental updates.
// Maintains three interconnected dependency maps for efficient change detection:
//
// Dependency Relationships:
// - fileHashes: Maps filesystem paths to SHA256 hashes for change detection
// - moduleDeps: Maps module paths to their ModuleDependencies (what each module imports)
// - cssDepReverse: Reverse mapping from CSS filesystem paths to modules that depend on them
//
// Data Flow:
// 1. RecordModuleDependencies() populates moduleDeps and builds cssDepReverse
// 2. UpdateFileHash() and HasFileChanged() use fileHashes for change detection
// 3. GetModulesAffectedByFiles() uses both moduleDeps and cssDepReverse to find rebuild targets
//
// Path Conventions:
// - fileHashes and cssDepReverse use filesystem paths (absolute)
// - moduleDeps uses module paths (relative to workspace root)
// - Workspace context handles conversions between path types
//
// Callsites:
// - session_core.go:55 (NewGenerateSession creation)
// - session_core.go:251 (dependency tracking in processing)
// - session_incremental.go:53,178 (incremental rebuild logic)
//
// Thread Safety: Protected by sync.RWMutex for concurrent read/write access
type FileDependencyTracker struct {
	mu            sync.RWMutex
	fileHashes    FileHashMap         // FS path -> SHA256 hash for change detection
	moduleDeps    ModuleDependencyMap // Module path -> dependencies (forward mapping)
	cssDepReverse CssReverseDepMap    // CSS FS path -> module paths that depend on it (reverse mapping)
	lastScanTime  time.Time
	ctx           W.WorkspaceContext
}

// ModuleDependencies tracks dependencies for a specific module
type ModuleDependencies struct {
	ModulePath    string
	StyleImports  []string
	ImportedFiles []string
	LastModified  time.Time
}

// NewFileDependencyTracker creates a new dependency tracker.
//
// Callsites:
// - session_core.go:55 (GenerateSession initialization)
//
// Returns: Initialized tracker with empty dependency maps
func NewFileDependencyTracker(ctx W.WorkspaceContext) *FileDependencyTracker {
	return &FileDependencyTracker{
		fileHashes:    make(FileHashMap),
		moduleDeps:    make(ModuleDependencyMap),
		cssDepReverse: make(CssReverseDepMap),
		ctx:           ctx,
	}
}

// UpdateFileHash computes and stores the hash for a file (expects FS path)
func (fdt *FileDependencyTracker) UpdateFileHash(fsPath string) ([32]byte, error) {
	// Check the file's modification time
	fileInfo, err := os.Stat(fsPath)
	if err != nil {
		return [32]byte{}, err
	}
	// Acquire read lock to check if the file hash is up-to-date
	fdt.mu.RLock()
	lastHash, exists := fdt.fileHashes[fsPath]
	fdt.mu.RUnlock()
	if exists && fileInfo.ModTime().Before(fdt.lastScanTime) {
		// File hasn't changed since the last scan; return the existing hash
		return lastHash, nil
	}
	// Compute the new hash
	fdt.mu.Lock()
	defer fdt.mu.Unlock()

	file, err := os.Open(fsPath)
	if err != nil {
		return [32]byte{}, err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return [32]byte{}, err
	}

	hash := [32]byte(hasher.Sum(nil))
	fdt.fileHashes[fsPath] = hash
	return hash, nil
}

// HasFileChanged checks if a file has changed since last scan (expects FS path)
func (fdt *FileDependencyTracker) HasFileChanged(fsPath string) (bool, error) {
	fdt.mu.RLock()
	lastHash, exists := fdt.fileHashes[fsPath]
	fdt.mu.RUnlock()

	if !exists {
		return true, nil // New file
	}

	currentHash, err := fdt.UpdateFileHash(fsPath)
	if err != nil {
		return true, err // Assume changed if we can't read it
	}

	return currentHash != lastHash, nil
}

// GetModulesAffectedByFiles returns modules that need rebuilding due to file changes.
// Expects module paths as input, returns module paths.
//
// Callsites:
// - session_incremental.go:53 (ProcessChangedFiles)
// - session_incremental.go:275 (ProcessChangedFilesIncremental)
//
// Algorithm: Direct module changes + CSS reverse dependency lookup
func (fdt *FileDependencyTracker) GetModulesAffectedByFiles(changedModulePaths []string) []string {
	fdt.mu.RLock()
	defer fdt.mu.RUnlock()

	affectedModules := make(map[string]bool)

	for _, modulePath := range changedModulePaths {
		// Direct module file changes
		if deps, exists := fdt.moduleDeps[modulePath]; exists {
			affectedModules[deps.ModulePath] = true
		}

		// CSS dependency changes - convert module path to FS path for CSS reverse lookup
		fsPath := fdt.ctx.ModulePathToFS(modulePath)
		if modules, exists := fdt.cssDepReverse[fsPath]; exists {
			for _, module := range modules {
				affectedModules[module] = true
			}
		}
	}

	result := make([]string, 0, len(affectedModules))
	for module := range affectedModules {
		result = append(result, module)
	}
	return result
}

// RecordModuleDependencies stores dependencies for a module.
//
// Callsites:
// - generate.go:178 (processModuleWithDeps)
//
// Purpose: Records style imports and file imports to enable incremental rebuilds
func (fdt *FileDependencyTracker) RecordModuleDependencies(modulePath string, styleImports, importedFiles []string) {
	fdt.mu.Lock()
	defer fdt.mu.Unlock()

	// Resolve all dependency paths using workspace context
	resolvedStyleImports := make([]string, 0, len(styleImports))
	resolvedImportedFiles := make([]string, 0, len(importedFiles))

	for _, styleImport := range styleImports {
		resolved := fdt.ctx.ResolveModuleDependency(modulePath, styleImport)
		resolvedStyleImports = append(resolvedStyleImports, resolved)
	}

	for _, importFile := range importedFiles {
		resolved := fdt.ctx.ResolveModuleDependency(modulePath, importFile)
		resolvedImportedFiles = append(resolvedImportedFiles, resolved)
	}

	deps := &ModuleDependencies{
		ModulePath:    modulePath,
		StyleImports:  resolvedStyleImports,
		ImportedFiles: resolvedImportedFiles,
		LastModified:  time.Now(),
	}

	fdt.moduleDeps[modulePath] = deps

	// Update reverse CSS dependencies using FS paths for CSS files
	for _, cssModulePath := range resolvedStyleImports {
		cssFS := fdt.ctx.ModulePathToFS(cssModulePath)
		if fdt.cssDepReverse[cssFS] == nil {
			fdt.cssDepReverse[cssFS] = make([]string, 0)
		}
		// Avoid duplicates
		found := false
		for _, existing := range fdt.cssDepReverse[cssFS] {
			if existing == modulePath {
				found = true
				break
			}
		}
		if !found {
			fdt.cssDepReverse[cssFS] = append(fdt.cssDepReverse[cssFS], modulePath)
		}
	}
}
