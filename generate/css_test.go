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
	"testing"

	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

func TestCssCache_Interface(t *testing.T) {
	// Test that CssParseCache implements CssCache interface
	var cache CssCache = NewCssParseCache()

	// Test basic operations
	testProps := CssPropsMap{
		"--test-color": M.CssCustomProperty{
			FullyQualified: M.FullyQualified{Name: "--test-color"},
		},
	}

	// Test Set and Get
	cache.Set("/test/path.css", testProps)

	retrieved, found := cache.Get("/test/path.css")
	if !found {
		t.Error("Should find cached entry")
	}

	if len(retrieved) != 1 {
		t.Errorf("Expected 1 property, got %d", len(retrieved))
	}

	if prop, exists := retrieved["--test-color"]; !exists || prop.Name != "--test-color" {
		t.Error("Property not correctly stored/retrieved")
	}

	// Test Invalidate
	cache.Invalidate([]string{"/test/path.css"})
	_, found = cache.Get("/test/path.css")
	if found {
		t.Error("Entry should be invalidated")
	}

	// Test Clear
	cache.Set("/test/path1.css", testProps)
	cache.Set("/test/path2.css", testProps)
	cache.Clear()

	_, found1 := cache.Get("/test/path1.css")
	_, found2 := cache.Get("/test/path2.css")
	if found1 || found2 {
		t.Error("All entries should be cleared")
	}
}

func TestCssCache_ThreadSafety(t *testing.T) {
	cache := NewCssParseCache()

	// Simple test for concurrent access (basic smoke test)
	// This doesn't guarantee thread safety but checks for obvious issues
	testProps := CssPropsMap{
		"--test": M.CssCustomProperty{
			FullyQualified: M.FullyQualified{Name: "--test"},
		},
	}

	done := make(chan bool, 2)

	// Writer goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Set("/test.css", testProps)
		}
		done <- true
	}()

	// Reader goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Get("/test.css")
		}
		done <- true
	}()

	// Wait for both goroutines
	<-done
	<-done
}

func TestGenerateSession_CssCache_Integration(t *testing.T) {
	// Test CSS cache initialization in GenerateSession
	ctx := W.NewFileSystemWorkspaceContext("testdata")
	setupCtx, err := NewGenerateContext(ctx)
	if err != nil {
		t.Fatalf("Failed to create setup context: %v", err)
	}
	defer setupCtx.Close()

	gs := &GenerateSession{
		setupCtx: setupCtx,
	}

	cache := gs.GetCssCache()
	if cache == nil {
		t.Fatal("CSS cache should be available from GenerateSession")
	}

	// Test basic functionality through session
	testProps := CssPropsMap{
		"--session-test": M.CssCustomProperty{
			FullyQualified: M.FullyQualified{Name: "--session-test"},
		},
	}

	cache.Set("/session/test.css", testProps)
	retrieved, found := cache.Get("/session/test.css")

	if !found {
		t.Error("Should find cached entry through session")
	}

	if len(retrieved) != 1 {
		t.Errorf("Expected 1 property, got %d", len(retrieved))
	}
}

func TestCssCache_FullIntegration(t *testing.T) {
	// Test that CSS cache abstraction is working properly

	// Test that different cache implementations can be used
	cache1 := NewCssParseCache()
	cache2 := NewCssParseCache()

	// Test that they behave independently
	testProps := CssPropsMap{
		"--test": M.CssCustomProperty{
			FullyQualified: M.FullyQualified{Name: "--test"},
		},
	}

	cache1.Set("/test1.css", testProps)
	cache2.Set("/test2.css", testProps)

	// cache1 should not have cache2's entry
	_, found := cache1.Get("/test2.css")
	if found {
		t.Error("Cache1 should not have cache2's entries")
	}

	// cache2 should not have cache1's entry
	_, found = cache2.Get("/test1.css")
	if found {
		t.Error("Cache2 should not have cache1's entries")
	}

	// This verifies that the CSS cache abstraction allows for proper isolation
	// between different sessions, which is essential for LSP integration
}

