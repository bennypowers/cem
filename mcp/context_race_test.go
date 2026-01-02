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
package mcp

import (
	"runtime"
	"sync"
	"testing"
	"time"

	W "bennypowers.dev/cem/workspace"
)

func TestMCPContext_ConcurrentCacheAccess(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	ctx, err := NewMCPContext(workspace)
	if err != nil {
		t.Fatalf("Failed to create MCP context: %v", err)
	}

	if err := ctx.LoadManifests(); err != nil {
		t.Fatalf("Failed to load manifests: %v", err)
	}

	const numGoroutines = 50
	const iterations = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Start multiple goroutines that access cached data concurrently
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				// Access both cached properties concurrently
				prefixes := ctx.CommonPrefixes()
				cssProps := ctx.AllCSSProperties()

				// Verify results are consistent (not nil)
				// Note: CommonPrefixes might be empty slice if no common prefixes exist
				if prefixes == nil {
					t.Errorf("Goroutine %d iteration %d: CommonPrefixes returned nil", goroutineID, j)
					return
				}
				if cssProps == nil {
					t.Errorf("Goroutine %d iteration %d: AllCSSProperties returned nil", goroutineID, j)
					return
				}
				// Verify CSS properties are actually loaded
				if len(cssProps) == 0 {
					t.Errorf("Goroutine %d iteration %d: AllCSSProperties returned empty slice", goroutineID, j)
					return
				}

				// Small delay to increase chance of race conditions
				if j%10 == 0 {
					runtime.Gosched()
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestMCPContext_ConcurrentLoadManifestsAndCacheAccess(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	ctx, err := NewMCPContext(workspace)
	if err != nil {
		t.Fatalf("Failed to create MCP context: %v", err)
	}

	if err := ctx.LoadManifests(); err != nil {
		t.Fatalf("Failed to load manifests: %v", err)
	}

	const duration = 2 * time.Second
	done := make(chan struct{})

	var wg sync.WaitGroup

	// Goroutine that repeatedly reloads manifests
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				_ = ctx.LoadManifests() // Ignore errors, focus on race conditions
			}
		}
	}()

	// Multiple goroutines accessing cached data
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			for {
				select {
				case <-done:
					return
				default:
					// Access cached properties
					prefixes := ctx.CommonPrefixes()
					cssProps := ctx.AllCSSProperties()

					// Verify consistency - should never be nil after initial load
					if prefixes == nil || cssProps == nil {
						t.Errorf("Goroutine %d: Got nil results during concurrent access", goroutineID)
						return
					}
					// Verify CSS properties contain data (CommonPrefixes can be empty)
					if len(cssProps) == 0 {
						t.Errorf("Goroutine %d: Got empty CSS properties during concurrent access", goroutineID)
						return
					}

					runtime.Gosched()
				}
			}
		}(i)
	}

	// Run test for specified duration
	time.Sleep(duration)
	close(done)
	wg.Wait()
}

func TestMCPContext_SnapshotConsistency(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	ctx, err := NewMCPContext(workspace)
	if err != nil {
		t.Fatalf("Failed to create MCP context: %v", err)
	}

	if err := ctx.LoadManifests(); err != nil {
		t.Fatalf("Failed to load manifests: %v", err)
	}

	// Test that snapshots are internally consistent
	// Even under concurrent access, a single snapshot should be atomic
	const numTests = 100

	for i := 0; i < numTests; i++ {
		ctx.mu.RLock()
		snapshot, err := ctx.takeRegistrySnapshot()
		ctx.mu.RUnlock()

		if err != nil {
			t.Fatalf("Failed to take snapshot: %v", err)
		}

		if snapshot == nil {
			t.Fatal("Snapshot is nil")
		}

		if snapshot.elements == nil {
			t.Fatal("Snapshot elements map is nil")
		}

		// Verify snapshot contains expected elements
		if len(snapshot.elements) == 0 {
			t.Error("Snapshot contains no elements")
		}

		// Verify all elements are properly converted
		for tagName, element := range snapshot.elements {
			if element == nil {
				t.Errorf("Element %s is nil in snapshot", tagName)
			}
			if element.TagName() != tagName {
				t.Errorf("Element tag name mismatch: expected %s, got %s", tagName, element.TagName())
			}
		}
	}
}

func TestMCPContext_CacheInvalidationRace(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	ctx, err := NewMCPContext(workspace)
	if err != nil {
		t.Fatalf("Failed to create MCP context: %v", err)
	}

	if err := ctx.LoadManifests(); err != nil {
		t.Fatalf("Failed to load manifests: %v", err)
	}

	const numGoroutines = 20
	var wg sync.WaitGroup

	// Multiple goroutines that invalidate cache and access data
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < 50; j++ {
				// Invalidate cache by reloading manifests
				if j%5 == 0 {
					_ = ctx.LoadManifests()
				}

				// Access cached data
				prefixes := ctx.CommonPrefixes()
				cssProps := ctx.AllCSSProperties()

				// Verify data is valid
				if prefixes == nil || cssProps == nil {
					t.Errorf("Goroutine %d iteration %d: Got nil cached data", goroutineID, j)
					return
				}
				// Verify CSS properties contain data
				if len(cssProps) == 0 {
					t.Errorf("Goroutine %d iteration %d: Got empty CSS properties", goroutineID, j)
					return
				}

				runtime.Gosched()
			}
		}(i)
	}

	wg.Wait()
}
