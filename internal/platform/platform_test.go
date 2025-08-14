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
	"time"

	"bennypowers.dev/cem/internal/platform"
)

func TestMockTimeProvider(t *testing.T) {
	startTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	mockTime := platform.NewMockTimeProvider(startTime)

	// Test initial time
	if mockTime.Now() != startTime {
		t.Errorf("Expected initial time %v, got %v", startTime, mockTime.Now())
	}

	// Test Sleep - should advance time instantly
	mockTime.Sleep(5 * time.Second)
	expectedTime := startTime.Add(5 * time.Second)
	if mockTime.Now() != expectedTime {
		t.Errorf("Expected time after sleep %v, got %v", expectedTime, mockTime.Now())
	}

	// Test sleep calls tracking
	sleepCalls := mockTime.GetSleepCalls()
	if len(sleepCalls) != 1 || sleepCalls[0] != 5*time.Second {
		t.Errorf("Expected sleep calls [5s], got %v", sleepCalls)
	}

	// Test After channel
	ch := mockTime.After(1 * time.Second)
	select {
	case receivedTime := <-ch:
		expectedAfterTime := expectedTime.Add(1 * time.Second)
		if receivedTime != expectedAfterTime {
			t.Errorf("Expected After time %v, got %v", expectedAfterTime, receivedTime)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("After channel should have delivered time immediately")
	}
}

func TestMockFileWatcher(t *testing.T) {
	watcher := platform.NewMockFileWatcher()
	defer watcher.Close()

	// Test adding watch paths
	err := watcher.Add("/test/path")
	if err != nil {
		t.Fatalf("Failed to add watch path: %v", err)
	}

	// Verify path is watched
	watchedPaths := watcher.GetWatchedPaths()
	if len(watchedPaths) != 1 || watchedPaths[0] != "/test/path" {
		t.Errorf("Expected watched paths [/test/path], got %v", watchedPaths)
	}

	// Test triggering events
	go watcher.TriggerEvent("/test/path/file.txt", platform.Write)

	select {
	case event := <-watcher.Events():
		if event.Name != "/test/path/file.txt" {
			t.Errorf("Expected event name /test/path/file.txt, got %s", event.Name)
		}
		if event.Op != platform.Write {
			t.Errorf("Expected Write operation, got %v", event.Op)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should have received event immediately")
	}

	// Test removing watch paths
	err = watcher.Remove("/test/path")
	if err != nil {
		t.Fatalf("Failed to remove watch path: %v", err)
	}

	watchedPaths = watcher.GetWatchedPaths()
	if len(watchedPaths) != 0 {
		t.Errorf("Expected no watched paths, got %v", watchedPaths)
	}
}

func TestTempDirFileSystem(t *testing.T) {
	fs, err := platform.NewTempDirFileSystem()
	if err != nil {
		t.Fatalf("Failed to create temp dir filesystem: %v", err)
	}
	defer fs.Cleanup()

	// Test file operations
	testData := []byte("hello world")
	err = fs.WriteFile("test.txt", testData, 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Test file exists
	if !fs.Exists("test.txt") {
		t.Error("File should exist after writing")
	}

	// Test reading file
	data, err := fs.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(data) != string(testData) {
		t.Errorf("Expected file content %q, got %q", testData, data)
	}

	// Test directory creation
	err = fs.MkdirAll("subdir/nested", 0755)
	if err != nil {
		t.Fatalf("Failed to create directories: %v", err)
	}

	// Test temp dir isolation
	tempDir := fs.TempDir()
	if tempDir == "" {
		t.Error("TempDir should return non-empty path")
	}

	// Test real path resolution
	realPath := fs.RealPath("test.txt")
	if realPath == "test.txt" {
		t.Error("RealPath should return absolute path within temp directory")
	}
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
