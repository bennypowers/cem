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

package testutil

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
)

// NewFixtureFS loads fixture files from testdata and returns a MapFileSystem
// with files mapped to the specified root path (usually "/test").
// The fixtureDir should be relative to serve/testdata/ (e.g., "transforms/config-test").
// Go tests run from the module root, so we check both possible locations.
func NewFixtureFS(t *testing.T, fixtureDir string, rootPath string) *platform.MapFileSystem {
	t.Helper()

	mfs := platform.NewMapFileSystem(nil)

	// Go test changes working directory based on which package is being tested.
	// Try multiple possible paths:
	// 1. testdata/ - for tests in serve package
	// 2. ../../testdata/ - for tests in serve/middleware/* packages
	// 3. serve/testdata/ - fallback if running from module root
	possiblePaths := []string{
		filepath.Join("testdata", fixtureDir),
		filepath.Join("..", "..", "testdata", fixtureDir),
		filepath.Join("serve", "testdata", fixtureDir),
	}

	var fixturePath string
	var statErr error
	for _, path := range possiblePaths {
		if _, statErr = os.Stat(path); statErr == nil {
			fixturePath = path
			break
		}
	}
	if fixturePath == "" {
		t.Fatalf("Could not find fixtures at %s (tried all paths)", fixtureDir)
	}

	// Walk fixture directory and load all files into memory
	// Callback returns errors immediately to terminate on first failure
	err := filepath.WalkDir(fixturePath, func(path string, d fs.DirEntry, err error) error {
		// Propagate any errors from WalkDir (e.g., permission denied)
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// Read fixture file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Get relative path from fixture directory
		relPath, err := filepath.Rel(fixturePath, path)
		if err != nil {
			return err
		}

		// Map to virtual filesystem at rootPath
		virtualPath := filepath.Join(rootPath, relPath)
		mfs.AddFile(virtualPath, string(content), 0644)

		return nil
	})

	if err != nil {
		t.Fatalf("Failed to load fixtures from %s: %v", fixtureDir, err)
	}

	return mfs
}

// LoadFixtureFile reads a single fixture file and returns its content.
// The fixturePath should be relative to serve/testdata/ (e.g., "demo-routing/manifest.json").
// Go tests run from the module root, so we check both possible locations.
func LoadFixtureFile(t *testing.T, fixturePath string) []byte {
	t.Helper()

	// Go test changes working directory based on which package is being tested.
	// Try multiple possible paths:
	// 1. testdata/ - for tests in serve package
	// 2. ../../testdata/ - for tests in serve/middleware/* packages
	// 3. serve/testdata/ - fallback if running from module root
	possiblePaths := []string{
		filepath.Join("testdata", fixturePath),
		filepath.Join("..", "..", "testdata", fixturePath),
		filepath.Join("serve", "testdata", fixturePath),
	}

	var content []byte
	var err error
	for _, path := range possiblePaths {
		content, err = os.ReadFile(path)
		if err == nil {
			return content
		}
	}
	t.Fatalf("Failed to read fixture %s (tried all paths): %v", fixturePath, err)
	return nil
}
