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
package platform_test

import (
	"testing"

	"bennypowers.dev/cem/internal/platform"
)

// MockTimeProvider and MockFileWatcher tests moved to platform_race_test.go
//
// These tests trigger Go's race detector "hole in findfunctab" issue due to
// channel operations in the mock implementations. The tests are preserved with
// the `race_unsafe` build tag and can be run manually:
//
//   go test -tags=race_unsafe ./internal/platform/
//
// This approach allows us to:
// 1. Keep the tests for future use (when Go fixes the race detector)
// 2. Run them manually for verification
// 3. Maintain race detection on the main codebase
// 4. Document the limitation clearly
//
// The mocks work correctly in production and their value is tested through
// integration tests in other packages that use them.

func TestFileSystemIsolationEffects(t *testing.T) {
	// Test that TempDirFileSystem enables isolated filesystem testing
	// by providing real filesystem operations within a controlled, isolated environment
	
	fs, err := platform.NewTempDirFileSystem()
	if err != nil {
		t.Fatalf("Failed to create isolated filesystem: %v", err)
	}
	defer func() {
		if err := fs.Cleanup(); err != nil {
			t.Logf("Warning: error during filesystem cleanup: %v", err)
		}
	}()

	// Test that operations work with real filesystem semantics but without interference
	testData := []byte("isolated test content")
	err = fs.WriteFile("manifest.json", testData, 0644)
	if err != nil {
		t.Fatalf("Isolated filesystem should support real file operations: %v", err)
	}

	// Verify isolation - files exist within the temp environment
	if !fs.Exists("manifest.json") {
		t.Error("Isolated filesystem should support real file existence checks")
	}

	// Verify data integrity with real filesystem operations
	data, err := fs.ReadFile("manifest.json")
	if err != nil {
		t.Fatalf("Isolated filesystem should support real file reading: %v", err)
	}

	if string(data) != string(testData) {
		t.Errorf("Isolated filesystem should preserve data integrity")
	}

	// Test that complex operations work (directory creation, nested paths)
	err = fs.MkdirAll("components/shared", 0755)
	if err != nil {
		t.Fatalf("Isolated filesystem should support directory operations: %v", err)
	}

	// Verify the isolation effect - operations are contained within temp directory
	tempDir := fs.TempDir()
	if tempDir == "" {
		t.Error("TempDirFileSystem should provide access to isolation boundary")
	}

	// Test path resolution works for debugging/verification
	realPath := fs.RealPath("manifest.json")
	if realPath == "manifest.json" {
		t.Error("TempDirFileSystem should resolve paths within isolation boundary")
	}
	
	// The key effect: enables testing with real filesystem semantics
	// but without polluting the actual filesystem or interfering with other tests
}

func TestInterfaceCompliance(t *testing.T) {
	// Ensure our implementations satisfy the interfaces
	var _ platform.TimeProvider = (*platform.RealTimeProvider)(nil)
	var _ platform.TimeProvider = (*platform.MockTimeProvider)(nil)
	var _ platform.FileWatcher = (*platform.FSNotifyFileWatcher)(nil)
	var _ platform.FileWatcher = (*platform.MockFileWatcher)(nil)
	var _ platform.FileSystem = (*platform.OSFileSystem)(nil)
	var _ platform.FileSystem = (*platform.TempDirFileSystem)(nil)
}
