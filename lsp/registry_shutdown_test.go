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
package lsp_test

import (
	"testing"
	"testing/synctest"
	"time"

	"bennypowers.dev/cem/internal/platform"
	LSP "bennypowers.dev/cem/lsp"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRegistryFileWatcherShutdown verifies that the file watcher goroutine
// exits properly when StopFileWatching is called.
// This test prevents the bug where LSP server hangs at 100% CPU after client disconnects.
func TestRegistryFileWatcherShutdown(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a mock file watcher
		mockWatcher := platform.NewMockFileWatcher()
		registry := LSP.NewRegistry(mockWatcher)

		// Track if reload callback was called
		reloadCalled := false
		reloadCallback := func() {
			reloadCalled = true
		}

		// Start file watching
		err := registry.StartFileWatching(reloadCallback)
		require.NoError(t, err)

		// Give the goroutine time to start
		synctest.Wait()

		// Verify watcher is running by triggering an event
		mockWatcher.TriggerEvent("/test/manifest.json", platform.Write)
		synctest.Wait()

		// Note: The callback won't be called because the path isn't in ManifestPaths
		// but this verifies the goroutine is listening

		// Now stop file watching - this should signal the goroutine to exit
		err = registry.StopFileWatching()
		require.NoError(t, err)

		// Give time for goroutine to exit
		synctest.Wait()

		// Verify that the watcher is closed
		// Attempting to trigger events on a closed watcher should be safe (no-op)
		mockWatcher.TriggerEvent("/test/manifest.json", platform.Write)

		// If we get here without hanging, the goroutine exited successfully
		assert.False(t, reloadCalled, "Callback should not have been called for unwatched path")
	})
}

// TestRegistryFileWatcherShutdownWithManifests tests shutdown with actual watched paths
func TestRegistryFileWatcherShutdownWithManifests(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a mock file watcher
		mockWatcher := platform.NewMockFileWatcher()
		registry := LSP.NewRegistry(mockWatcher)

		// Add a manifest path manually for testing
		registry.AddManifestPath("/test/manifest.json")

		// Track if reload callback was called
		reloadCalled := false
		reloadCallback := func() {
			reloadCalled = true
		}

		// Start file watching
		err := registry.StartFileWatching(reloadCallback)
		require.NoError(t, err)

		// Give the goroutine time to start
		synctest.Wait()

		// Verify watcher is running and callback works
		mockWatcher.TriggerEvent("/test/manifest.json", platform.Write)
		synctest.Wait()
		assert.True(t, reloadCalled, "Callback should have been called for watched path")

		// Reset flag
		reloadCalled = false

		// Now stop file watching
		err = registry.StopFileWatching()
		require.NoError(t, err)

		// Give time for goroutine to exit
		synctest.Wait()

		// After stopping, trigger another event - callback should NOT be called
		mockWatcher.TriggerEvent("/test/manifest.json", platform.Write)
		synctest.Wait()

		// Callback should not be called because watcher is stopped
		assert.False(t, reloadCalled, "Callback should not be called after stopping watcher")
	})
}

// TestRegistryFileWatcherStopWithTimeout verifies shutdown completes within reasonable time
// This is the critical test for the hanging LSP server bug
func TestRegistryFileWatcherStopWithTimeout(t *testing.T) {
	// Create a mock file watcher
	mockWatcher := platform.NewMockFileWatcher()
	registry := LSP.NewRegistry(mockWatcher)

	// Start file watching
	err := registry.StartFileWatching(func() {})
	require.NoError(t, err)

	// Add a small delay to ensure goroutine is running
	time.Sleep(10 * time.Millisecond)

	// Stop file watching with a timeout
	done := make(chan error, 1)
	go func() {
		done <- registry.StopFileWatching()
	}()

	// Wait for StopFileWatching to complete with a timeout
	select {
	case err := <-done:
		require.NoError(t, err, "StopFileWatching should complete without error")
	case <-time.After(1 * time.Second):
		t.Fatal("StopFileWatching did not complete within 1 second - goroutine is hanging!")
	}
}

// TestGenerateWatcherShutdown verifies that the generate watcher goroutine
// exits properly when Stop is called.
// This test prevents the bug where LSP server hangs at 100% CPU from the generate watcher.
func TestGenerateWatcherShutdown(t *testing.T) {
	// Create a test workspace context
	tempDir := t.TempDir()
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err := workspace.Init()
	require.NoError(t, err)

	// Create callback
	callbackCalled := false
	callback := func(pkg *M.Package) error {
		callbackCalled = true
		return nil
	}

	// Create generate watcher
	watcher, err := LSP.NewInProcessGenerateWatcher(workspace, []string{"**/*.ts"}, callback)
	require.NoError(t, err)

	// Start the watcher
	err = watcher.Start()
	require.NoError(t, err)

	// Give it time to start
	time.Sleep(10 * time.Millisecond)

	// Stop the watcher with a timeout
	done := make(chan error, 1)
	go func() {
		done <- watcher.Stop()
	}()

	select {
	case err := <-done:
		require.NoError(t, err, "Stop should complete without error")
		assert.True(t, callbackCalled, "Initial callback should have been called")
	case <-time.After(1 * time.Second):
		t.Fatal("Generate watcher Stop did not complete within 1 second - goroutine is hanging!")
	}
}

// TestGenerateWatcherDoubleStart verifies that starting the watcher twice fails appropriately
func TestGenerateWatcherDoubleStart(t *testing.T) {
	tempDir := t.TempDir()
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err := workspace.Init()
	require.NoError(t, err)

	watcher, err := LSP.NewInProcessGenerateWatcher(workspace, []string{"**/*.ts"}, func(pkg *M.Package) error {
		return nil
	})
	require.NoError(t, err)

	// First start should succeed
	err1 := watcher.Start()
	require.NoError(t, err1, "First Start should succeed")

	// Second start should fail
	err2 := watcher.Start()
	require.Error(t, err2, "Second Start should fail")
	assert.Contains(t, err2.Error(), "already running", "Error should mention already running")

	// Cleanup
	_ = watcher.Stop()
}
