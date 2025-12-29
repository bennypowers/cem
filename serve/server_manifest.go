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

package serve

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/routes"
	W "bennypowers.dev/cem/workspace"
)

// SetManifest sets the current manifest and builds the routing table
func (s *Server) SetManifest(manifest []byte) error {
	// Extract state under read lock
	s.mu.RLock()
	sourceControlURL := s.sourceControlRootURL
	s.mu.RUnlock()

	// Defensive copy to prevent caller from mutating our internal state
	manifestCopy := make([]byte, len(manifest))
	copy(manifestCopy, manifest)

	// Build routing table from manifest (expensive - no lock held)
	routingTable, err := routes.BuildDemoRoutingTable(manifestCopy, sourceControlURL)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table: %v", err)
		routingTable = nil
	} else {
		s.logger.Debug("Built routing table with %d demo routes", len(routingTable))
	}

	// Update state under write lock
	s.mu.Lock()
	defer s.mu.Unlock()

	s.manifest = manifestCopy
	s.demoRoutes = routingTable

	if err != nil {
		// Return error wrapped to indicate partial success
		// The manifest was stored successfully, but routing table construction failed
		return fmt.Errorf("manifest stored but routing table failed: %w", err)
	}
	return nil
}

// Manifest returns the current manifest
func (s *Server) Manifest() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Defensive copy to prevent caller from mutating our internal state
	if s.manifest == nil {
		return nil, nil
	}
	manifestCopy := make([]byte, len(s.manifest))
	copy(manifestCopy, s.manifest)
	return manifestCopy, nil
}

// PackageJSON returns parsed package.json from the watch directory
func (s *Server) PackageJSON() (*middleware.PackageJSON, error) {
	watchDir := s.WatchDir()
	if watchDir == "" {
		return nil, nil
	}

	packageJSONPath := filepath.Join(watchDir, "package.json")

	var data []byte
	var err error

	// Use injected filesystem if available, otherwise fall back to os.ReadFile
	if filesystem := s.FileSystem(); filesystem != nil {
		data, err = filesystem.ReadFile(packageJSONPath)
	} else {
		data, err = os.ReadFile(packageJSONPath)
	}

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	var pkg struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	return &middleware.PackageJSON{
		Name:    pkg.Name,
		Version: pkg.Version,
	}, nil
}

// TryLoadExistingManifest attempts to load a manifest from disk
// Returns the manifest size in bytes and any error
// Returns 0, nil if no manifest file exists (not an error condition)
func (s *Server) TryLoadExistingManifest() (int, error) {
	// Extract state under read lock to avoid blocking during I/O
	s.mu.RLock()
	watchDir := s.watchDir
	sourceControlURL := s.sourceControlRootURL
	filesystem := s.fs
	s.mu.RUnlock()

	if watchDir == "" {
		return 0, fmt.Errorf("no watch directory set")
	}

	// Perform I/O operations without holding lock
	// Determine manifest path
	workspace := W.NewFileSystemWorkspaceContext(watchDir)
	if err := workspace.Init(); err != nil {
		return 0, fmt.Errorf("initializing workspace: %w", err)
	}

	manifestPath := workspace.CustomElementsManifestPath()
	if manifestPath == "" {
		return 0, nil // No manifest path configured
	}

	// Check if manifest file exists
	if _, err := filesystem.Stat(manifestPath); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return 0, nil // File doesn't exist, not an error
		}
		return 0, fmt.Errorf("checking manifest file: %w", err)
	}

	// Read existing manifest
	manifestBytes, err := filesystem.ReadFile(manifestPath)
	if err != nil {
		return 0, fmt.Errorf("reading manifest file: %w", err)
	}

	// Validate it's valid JSON
	var pkg any
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return 0, fmt.Errorf("invalid manifest JSON: %w", err)
	}

	// Build routing table from manifest
	routingTable, err := routes.BuildDemoRoutingTable(manifestBytes, sourceControlURL)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table from cached manifest: %v", err)
		routingTable = nil
	} else {
		s.logger.Debug("Built routing table with %d demo routes from cached manifest", len(routingTable))
	}

	// Update state under write lock
	s.mu.Lock()
	defer s.mu.Unlock()

	// Store the manifest (defensive copy to prevent external mutation)
	s.manifest = make([]byte, len(manifestBytes))
	copy(s.manifest, manifestBytes)
	s.demoRoutes = routingTable

	return len(manifestBytes), nil
}

// RegenerateManifest performs a full manifest regeneration
// Returns the manifest size in bytes and any error
func (s *Server) RegenerateManifest() (int, error) {
	// Extract state under read lock to avoid blocking during I/O
	s.mu.RLock()
	watchDir := s.watchDir
	sourceControlURL := s.sourceControlRootURL
	s.mu.RUnlock()

	if watchDir == "" {
		return 0, fmt.Errorf("no watch directory set")
	}

	// Perform expensive operations without holding lock
	// Create fresh workspace context and session for live reload
	// This ensures we always read the latest file contents
	workspace := W.NewFileSystemWorkspaceContext(watchDir)
	if err := workspace.Init(); err != nil {
		return 0, fmt.Errorf("initializing workspace: %w", err)
	}

	session, err := G.NewGenerateSession(workspace)
	if err != nil {
		return 0, fmt.Errorf("creating generate session: %w", err)
	}

	// Configure adaptive worker count to prevent goroutine explosion under concurrent load
	// Use same formula as transform pool for consistency
	maxWorkers := min(max(runtime.NumCPU()/2, 2), 8)
	session.SetMaxWorkers(maxWorkers)

	// Generate manifest (expensive - no lock held)
	ctx := context.Background()
	pkg, err := session.GenerateFullManifest(ctx)
	if err != nil {
		return 0, fmt.Errorf("generating manifest: %w", err)
	}

	// Extract source files from manifest for targeted file watching
	sourceFiles := make(map[string]bool)
	for _, module := range pkg.Modules {
		if module.Path != "" {
			sourceFiles[module.Path] = true
		}
	}

	// Marshal to JSON (expensive - no lock held)
	manifestBytes, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return 0, fmt.Errorf("marshaling manifest: %w", err)
	}

	// Build routing table from manifest (expensive - no lock held)
	routingTable, err := routes.BuildDemoRoutingTable(manifestBytes, sourceControlURL)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table: %v", err)
		routingTable = nil
	} else {
		s.logger.Debug("Built routing table with %d demo routes", len(routingTable))
	}

	s.logger.Debug("Tracking %d source files from manifest", len(sourceFiles))
	// Update server state under write lock
	s.mu.Lock()
	defer s.mu.Unlock()

	// Close old session if it exists
	if s.generateSession != nil {
		s.generateSession.Close()
	}

	s.generateSession = session
	s.sourceFiles = sourceFiles

	// Defensive copy (though json.MarshalIndent already returns a new slice)
	s.manifest = make([]byte, len(manifestBytes))
	copy(s.manifest, manifestBytes)
	s.demoRoutes = routingTable

	return len(manifestBytes), nil
}

// RegenerateManifestIncremental incrementally updates the manifest for changed files
// Returns the manifest size in bytes and any error
func (s *Server) RegenerateManifestIncremental(changedFiles []string) (int, error) {
	// Read required state under read lock
	s.mu.RLock()
	watchDir := s.watchDir
	session := s.generateSession
	sourceControlURL := s.sourceControlRootURL
	s.mu.RUnlock()

	if watchDir == "" {
		return 0, fmt.Errorf("no watch directory set")
	}

	// If no generate session exists yet, do a full regeneration
	if session == nil {
		return s.RegenerateManifest()
	}

	// Use incremental processing with existing session (no locks held)
	ctx := context.Background()
	pkg, err := session.ProcessChangedFiles(ctx, changedFiles)
	if err != nil {
		// If incremental processing fails, fall back to full regeneration
		s.logger.Warning("Incremental manifest generation failed, falling back to full regeneration: %v", err)
		return s.RegenerateManifest()
	}

	// Extract source files from manifest for targeted file watching
	sourceFiles := make(map[string]bool)
	for _, module := range pkg.Modules {
		if module.Path != "" {
			sourceFiles[module.Path] = true
		}
	}
	s.logger.Debug("Tracking %d source files from manifest", len(sourceFiles))

	// Marshal to JSON
	manifestBytes, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return 0, fmt.Errorf("marshaling manifest: %w", err)
	}

	// Build routing table from manifest
	routingTable, err := routes.BuildDemoRoutingTable(manifestBytes, sourceControlURL)
	if err != nil {
		s.logger.Warning("Failed to build demo routing table: %v", err)
		routingTable = nil
	} else {
		s.logger.Debug("Built routing table with %d demo routes", len(routingTable))
	}

	// Update server state under write lock
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sourceFiles = sourceFiles
	// Defensive copy (though json.MarshalIndent already returns a new slice)
	s.manifest = make([]byte, len(manifestBytes))
	copy(s.manifest, manifestBytes)
	s.demoRoutes = routingTable

	return len(manifestBytes), nil
}
