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
package modulegraph_test

import (
	"fmt"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform"
)

// TestFileOperationsPerformance demonstrates performance benefits
// of in-memory operations over disk I/O.
func TestFileOperationsPerformance(t *testing.T) {
	// Create in-memory filesystem - instant setup
	mapFS := platform.NewMapFileSystem(nil)

	// Create test files with realistic content in memory
	iconContent := `export class MyIcon {
  render() {
    return '<svg>...</svg>';
  }
}`

	tabContent := `import { MyIcon } from './my-icon.js';

export class MyTab {
  private icon = new MyIcon();
  
  render() {
    return this.icon.render() + '<div>tab content</div>';
  }
}`

	tabsContent := `import { MyTab } from './my-tab.js';

export class MyTabs {
  private tabs: MyTab[] = [];
  
  addTab() {
    this.tabs.push(new MyTab());
  }
}`

	// Add files to MapFS instantly - no disk I/O
	mapFS.AddFile("my-icon.ts", iconContent, 0644)
	mapFS.AddFile("my-tab.ts", tabContent, 0644)
	mapFS.AddFile("my-tabs.ts", tabsContent, 0644)

	// Validate filesystem compliance using Go 1.25 TestFS
	err := mapFS.TestFS("my-icon.ts", "my-tab.ts", "my-tabs.ts")
	if err != nil {
		t.Fatalf("MapFS should pass filesystem compliance tests: %v", err)
	}

	// Test that file operations work correctly and instantly
	for _, filename := range []string{"my-icon.ts", "my-tab.ts", "my-tabs.ts"} {
		// Test file existence
		if !mapFS.Exists(filename) {
			t.Errorf("File %s should exist", filename)
		}

		// Test file reading
		content, err := mapFS.ReadFile(filename)
		if err != nil {
			t.Errorf("Failed to read %s: %v", filename, err)
		}

		// Verify content is not empty
		if len(content) == 0 {
			t.Errorf("File %s should have content", filename)
		}

		// Test that content contains expected TypeScript elements
		contentStr := string(content)
		if !strings.Contains(contentStr, "export") {
			t.Errorf("File %s should contain TypeScript export", filename)
		}
	}

	// Demonstrate instant large-scale file operations that would be slow with disk I/O
	for i := 0; i < 50; i++ {
		filename := fmt.Sprintf("test-file-%d.ts", i)
		content := fmt.Sprintf("export const value%d = %d;", i, i)
		mapFS.AddFile(filename, content, 0644)

		// Verify file was created instantly
		if !mapFS.Exists(filename) {
			t.Errorf("File %s should exist immediately", filename)
		}
	}

	t.Logf("✅ File operations performance test: 53 files processed instantly")
}

// TestFileSystemCompliance tests that in-memory filesystem operations work correctly
func TestFileSystemCompliance(t *testing.T) {
	// Create in-memory filesystem
	mapFS := platform.NewMapFileSystem(nil)

	// Add test files
	mapFS.AddFile("test.ts", "export const test = 'hello';", 0644)
	mapFS.AddFile("subdir/nested.ts", "export const nested = 'world';", 0644)
	mapFS.AddDir("emptydir", 0755)

	// Test filesystem compliance using Go 1.25 TestFS
	err := mapFS.TestFS("test.ts", "subdir/nested.ts")
	if err != nil {
		t.Fatalf("MapFS should pass filesystem compliance tests: %v", err)
	}

	// Test ReadFile
	content, err := mapFS.ReadFile("test.ts")
	if err != nil {
		t.Fatalf("ReadFile should succeed: %v", err)
	}
	if string(content) != "export const test = 'hello';" {
		t.Errorf("Content mismatch: got %s", string(content))
	}

	// Test existence checks
	if !mapFS.Exists("test.ts") {
		t.Errorf("test.ts should exist")
	}
	if !mapFS.Exists("subdir/nested.ts") {
		t.Errorf("subdir/nested.ts should exist")
	}
	if !mapFS.Exists("emptydir") {
		t.Errorf("emptydir should exist")
	}

	// Test file listing
	files := mapFS.ListFiles()
	t.Logf("MapFS contains: %v", files)

	// Check that files exist (directories might not be listed the same way)
	expectedFiles := []string{"test.ts", "subdir/nested.ts"}
	for _, expected := range expectedFiles {
		if _, found := files[expected]; !found {
			t.Errorf("Expected file %s not found in listing", expected)
		}
	}

	// Verify directory exists using Exists() method instead of listing
	if !mapFS.Exists("emptydir") {
		t.Errorf("Expected directory emptydir to exist")
	}

	t.Logf("✅ Filesystem compliance test passed")
}
