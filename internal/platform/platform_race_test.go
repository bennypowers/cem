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

// This file contains tests for concurrent operations using Go 1.25's testing/synctest package.
//
// These tests previously had issues with Go's race detector "hole in findfunctab" limitation,
// but Go 1.25's synctest package provides proper isolation for concurrent testing scenarios.
//
// The synctest package provides:
// - Isolated "bubble" environments for concurrent tests
// - Virtual time that advances when goroutines block
// - Proper synchronization primitives for testing concurrent code
//
// This resolves the previous race detector issues with channel operations in mocks.

package platform_test

import (
	"testing"
	"testing/synctest"
	"time"

	"bennypowers.dev/cem/internal/platform"
)

func TestMockTimeProvider(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
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

		// Test After channel - now works correctly with synctest isolation
		ch := mockTime.After(1 * time.Second)
		select {
		case receivedTime := <-ch:
			expectedAfterTime := expectedTime.Add(1 * time.Second)
			if receivedTime != expectedAfterTime {
				t.Errorf("Expected After time %v, got %v", expectedAfterTime, receivedTime)
			}
		default:
			t.Error("After channel should have delivered time immediately")
		}
	})
}

func TestMockFileWatcher(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
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

		// Test removing watch paths
		err = watcher.Remove("/test/path")
		if err != nil {
			t.Fatalf("Failed to remove watch path: %v", err)
		}

		watchedPaths = watcher.GetWatchedPaths()
		if len(watchedPaths) != 0 {
			t.Errorf("Expected no watched paths, got %v", watchedPaths)
		}

		// Test triggering events - now works correctly with synctest isolation
		watcher.Add("/test/path") // Re-add for event testing
		watcher.TriggerEvent("/test/path/file.txt", platform.Write)

		// Read from Events() channel - synctest provides proper isolation
		select {
		case event := <-watcher.Events():
			if event.Name != "/test/path/file.txt" {
				t.Errorf("Expected event name /test/path/file.txt, got %s", event.Name)
			}
			if event.Op != platform.Write {
				t.Errorf("Expected Write operation, got %v", event.Op)
			}
		default:
			t.Error("Expected to receive file watcher event")
		}
	})
}
