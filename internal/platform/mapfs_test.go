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

// This file demonstrates Go 1.25's enhanced filesystem mocking capabilities
// using testing/fstest.MapFS integrated with synctest for concurrent testing.

package platform_test

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/synctest"
	"time"

	"bennypowers.dev/cem/internal/platform"
)

func TestMapFileSystem_BasicOperations(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mockTime := platform.NewMockTimeProvider(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
		mfs := platform.NewMapFileSystem(mockTime)

		// Test writing a file
		content := []byte("Hello, fstest.MapFS!")
		err := mfs.WriteFile("/test.txt", content, 0644)
		if err != nil {
			t.Fatalf("WriteFile failed: %v", err)
		}

		// Test reading the file
		readContent, err := mfs.ReadFile("/test.txt")
		if err != nil {
			t.Fatalf("ReadFile failed: %v", err)
		}

		if string(readContent) != string(content) {
			t.Errorf("File content mismatch. Expected %q, got %q", content, readContent)
		}

		// Test file exists
		if !mfs.Exists("/test.txt") {
			t.Error("File should exist")
		}

		// Test file stat
		info, err := mfs.Stat("/test.txt")
		if err != nil {
			t.Fatalf("Stat failed: %v", err)
		}

		if info.Size() != int64(len(content)) {
			t.Errorf("File size mismatch. Expected %d, got %d", len(content), info.Size())
		}

		if info.Mode().Perm() != 0644 {
			t.Errorf("File mode mismatch. Expected 0644, got %o", info.Mode().Perm())
		}

		if info.IsDir() {
			t.Error("File should not be a directory")
		}
	})
}

func TestMapFileSystem_TestFSValidation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		// Set up some test files using AddFile convenience method
		mfs.AddFile("hello.txt", "Hello, World!", 0644)
		mfs.AddFile("dir/nested.txt", "Nested content", 0644)
		mfs.AddFile("another_dir/file.txt", "Another file", 0644)

		// Use Go's fstest.TestFS to validate the filesystem
		err := mfs.TestFS("hello.txt", "dir/nested.txt", "another_dir/file.txt")
		if err != nil {
			t.Fatalf("TestFS validation failed: %v", err)
		}

		// Test that TestFS catches missing files
		err = mfs.TestFS("hello.txt", "nonexistent.txt")
		if err == nil {
			t.Error("TestFS should have failed for nonexistent file")
		}
	})
}

func TestMapFileSystem_FileWatcherIntegration(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)
		watcher := mfs.GetWatcher()

		// Debug: Check what paths are being watched
		t.Logf("Initial watched paths: %v", watcher.GetWatchedPaths())

		// Add watch for the root directory to catch all events
		err := watcher.Add("/")
		if err != nil {
			t.Fatalf("Failed to add watcher: %v", err)
		}

		t.Logf("Watched paths after adding '/': %v", watcher.GetWatchedPaths())

		// Perform filesystem operations that should trigger events
		_ = mfs.WriteFile("/test.txt", []byte("watched content"), 0644)

		// Check for file watcher events
		events := make([]platform.FileWatchEvent, 0)

		// Collect events with a reasonable limit
	eventLoop:
		for i := 0; i < 5; i++ {
			select {
			case event := <-watcher.Events():
				events = append(events, event)
				t.Logf("Received event: %s %v", event.Name, event.Op)
			default:
				// No more immediate events
				break eventLoop
			}
		}

		if len(events) < 1 {
			t.Errorf("Expected at least 1 event, got %d", len(events))
		}

		// Remove the file
		_ = mfs.Remove("/test.txt")

		// Check for remove event
		select {
		case event := <-watcher.Events():
			t.Logf("Received remove event: %s %v", event.Name, event.Op)
			if event.Op != platform.Remove {
				t.Errorf("Expected Remove event, got %v", event.Op)
			}
		default:
			t.Error("Expected Remove event")
		}
	})
}

func TestMapFileSystem_ConcurrentOperations(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		// Pre-create base directory
		mfs.AddDir("concurrent", 0755)

		const numGoroutines = 10
		const numFiles = 20

		done := make(chan bool, numGoroutines)

		// Start concurrent file operations
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < numFiles; j++ {
					filename := fmt.Sprintf("/concurrent/file_%d_%d.txt", id, j)
					content := fmt.Sprintf("Content from goroutine %d, file %d", id, j)

					// Write file
					err := mfs.WriteFile(filename, []byte(content), 0644)
					if err != nil {
						t.Errorf("WriteFile failed: %v", err)
						return
					}

					// Read file back
					readContent, err := mfs.ReadFile(filename)
					if err != nil {
						t.Errorf("ReadFile failed: %v", err)
						return
					}

					if string(readContent) != content {
						t.Errorf("Content mismatch for %s", filename)
						return
					}

					// Verify file exists
					if !mfs.Exists(filename) {
						t.Errorf("File %s should exist", filename)
						return
					}
				}
			}(i)
		}

		// Wait for all goroutines to complete
		for i := 0; i < numGoroutines; i++ {
			<-done
		}

		// Verify all files were created
		files := mfs.ListFiles()
		expectedFiles := numGoroutines*numFiles + 1 // +1 for the concurrent directory

		if len(files) != expectedFiles {
			t.Errorf("Expected %d files/dirs, got %d", expectedFiles, len(files))
		}
	})
}

func TestMapFileSystem_DirectoryOperations(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		// Test MkdirAll
		err := mfs.MkdirAll("/deep/nested/directory/structure", 0755)
		if err != nil {
			t.Fatalf("MkdirAll failed: %v", err)
		}

		// Verify all intermediate directories exist
		paths := []string{"/deep", "/deep/nested", "/deep/nested/directory", "/deep/nested/directory/structure"}
		for _, path := range paths {
			if !mfs.Exists(path) {
				t.Errorf("Directory %s should exist", path)
			}

			info, err := mfs.Stat(path)
			if err != nil {
				t.Errorf("Stat failed for %s: %v", path, err)
				continue
			}

			if !info.IsDir() {
				t.Errorf("Path %s should be a directory", path)
			}
		}

		// Test writing file in deep directory
		err = mfs.WriteFile("/deep/nested/directory/structure/file.txt", []byte("deep content"), 0644)
		if err != nil {
			t.Fatalf("WriteFile in deep directory failed: %v", err)
		}

		// Verify file content
		content, err := mfs.ReadFile("/deep/nested/directory/structure/file.txt")
		if err != nil {
			t.Fatalf("ReadFile failed: %v", err)
		}

		if string(content) != "deep content" {
			t.Errorf("Content mismatch. Expected 'deep content', got %q", content)
		}
	})
}

func TestMapFileSystem_ErrorHandling(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		// Test reading non-existent file
		_, err := mfs.ReadFile("/nonexistent.txt")
		if err == nil {
			t.Error("ReadFile should fail for non-existent file")
		}

		// Test removing non-existent file
		err = mfs.Remove("/nonexistent.txt")
		if err == nil {
			t.Error("Remove should fail for non-existent file")
		}

		// Test stat non-existent file
		_, err = mfs.Stat("/nonexistent.txt")
		if err == nil {
			t.Error("Stat should fail for non-existent file")
		}

		// Test that errors have the correct type
		_, err = mfs.ReadFile("/missing")
		if pathErr, ok := err.(*fs.PathError); ok {
			// fs.ReadFile typically returns "open" operation for missing files
			if pathErr.Op != "open" && pathErr.Op != "read" {
				t.Errorf("Expected operation 'open' or 'read', got %q", pathErr.Op)
			}
		} else {
			t.Errorf("Expected *fs.PathError, got %T", err)
		}
	})
}

func TestMapFileSystem_InterfaceCompliance(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Verify that MapFileSystem implements the FileSystem interface
		var fs platform.FileSystem = platform.NewMapFileSystem(nil)

		// Test all interface methods work
		err := fs.WriteFile("/interface_test.txt", []byte("interface test"), 0644)
		if err != nil {
			t.Fatalf("WriteFile failed: %v", err)
		}

		content, err := fs.ReadFile("/interface_test.txt")
		if err != nil {
			t.Fatalf("ReadFile failed: %v", err)
		}

		if string(content) != "interface test" {
			t.Errorf("Content mismatch")
		}

		if !fs.Exists("/interface_test.txt") {
			t.Error("File should exist")
		}

		info, err := fs.Stat("/interface_test.txt")
		if err != nil {
			t.Fatalf("Stat failed: %v", err)
		}

		if info.Size() != 14 { // len("interface test")
			t.Errorf("Size mismatch")
		}

		err = fs.MkdirAll("/test/dir", 0755)
		if err != nil {
			t.Fatalf("MkdirAll failed: %v", err)
		}

		tempDir := fs.TempDir()
		if tempDir == "" {
			t.Error("TempDir should return a non-empty string")
		}

		err = fs.Remove("/interface_test.txt")
		if err != nil {
			t.Fatalf("Remove failed: %v", err)
		}

		if fs.Exists("/interface_test.txt") {
			t.Error("File should not exist after removal")
		}
	})
}

// Demonstrate advanced fstest features
func TestMapFileSystem_AdvancedFeatures(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		// Create a complex filesystem structure
		mfs.AddDir("project", 0755)
		mfs.AddDir("project/src", 0755)
		mfs.AddDir("project/tests", 0755)
		mfs.AddFile("project/README.md", "# My Project\n\nA great project!", 0644)
		mfs.AddFile("project/src/main.go", "package main\n\nfunc main() {\n\tprintln(\"Hello!\")\n}", 0644)
		mfs.AddFile("project/tests/main_test.go", "package main\n\nimport \"testing\"\n\nfunc TestMain(t *testing.T) {}", 0644)

		// Use TestFS to validate the entire structure
		expectedFiles := []string{
			"project",
			"project/README.md",
			"project/src",
			"project/src/main.go",
			"project/tests",
			"project/tests/main_test.go",
		}

		err := mfs.TestFS(expectedFiles...)
		if err != nil {
			t.Fatalf("TestFS failed: %v", err)
		}

		// Test file operations on the complex structure
		content, err := mfs.ReadFile("/project/src/main.go")
		if err != nil {
			t.Fatalf("ReadFile failed: %v", err)
		}

		expectedContent := "package main\n\nfunc main() {\n\tprintln(\"Hello!\")\n}"
		if string(content) != expectedContent {
			t.Errorf("Content mismatch for main.go")
		}

		// List all files for debugging
		files := mfs.ListFiles()
		t.Logf("MapFS contains %d entries:", len(files))
		for path, info := range files {
			t.Logf("  %s: %v", path, info)
		}
	})
}
