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
	"bennypowers.dev/cem/internal/platform"
	W "bennypowers.dev/cem/internal/workspace"
)

func TestCssCache_Interface(t *testing.T) {
	var cache CssCache = NewCssParseCache()

	testEntry := &CssCacheEntry{
		Props: CssPropsMap{
			"--test-color": M.CssCustomProperty{
				FullyQualified: M.FullyQualified{Name: "--test-color"},
			},
		},
		VarDefaults: map[string]bool{"--test-color": true},
	}

	cache.Set("/test/path.css", testEntry)

	retrieved, found := cache.Get("/test/path.css")
	if !found {
		t.Error("Should find cached entry")
	}

	if len(retrieved.Props) != 1 {
		t.Errorf("Expected 1 property, got %d", len(retrieved.Props))
	}

	if prop, exists := retrieved.Props["--test-color"]; !exists || prop.Name != "--test-color" {
		t.Error("Property not correctly stored/retrieved")
	}

	if !retrieved.VarDefaults["--test-color"] {
		t.Error("VarDefaults not correctly stored/retrieved")
	}

	cache.Invalidate([]string{"/test/path.css"})
	_, found = cache.Get("/test/path.css")
	if found {
		t.Error("Entry should be invalidated")
	}

	cache.Set("/test/path1.css", testEntry)
	cache.Set("/test/path2.css", testEntry)
	cache.Clear()

	_, found1 := cache.Get("/test/path1.css")
	_, found2 := cache.Get("/test/path2.css")
	if found1 || found2 {
		t.Error("All entries should be cleared")
	}
}

func TestCssCache_ThreadSafety(t *testing.T) {
	cache := NewCssParseCache()

	testEntry := &CssCacheEntry{
		Props: CssPropsMap{
			"--test": M.CssCustomProperty{
				FullyQualified: M.FullyQualified{Name: "--test"},
			},
		},
		VarDefaults: map[string]bool{"--test": true},
	}

	done := make(chan bool, 2)

	go func() {
		for range 100 {
			cache.Set("/test.css", testEntry)
		}
		done <- true
	}()

	go func() {
		for range 100 {
			cache.Get("/test.css")
		}
		done <- true
	}()

	<-done
	<-done
}

func TestGenerateSession_CssCache_Integration(t *testing.T) {
	ctx := W.NewFileSystemWorkspaceContext("testdata")
	setupCtx, err := NewGenerateContext(ctx, platform.NewOSFileSystem())
	if err != nil {
		t.Fatalf("Failed to create setup context: %v", err)
	}
	defer setupCtx.Close()

	gs := &GenerateSession{
		setupCtx: setupCtx,
	}

	cache := gs.CssCache()
	if cache == nil {
		t.Fatal("CSS cache should be available from GenerateSession")
	}

	testEntry := &CssCacheEntry{
		Props: CssPropsMap{
			"--session-test": M.CssCustomProperty{
				FullyQualified: M.FullyQualified{Name: "--session-test"},
			},
		},
		VarDefaults: map[string]bool{"--session-test": true},
	}

	cache.Set("/session/test.css", testEntry)
	retrieved, found := cache.Get("/session/test.css")

	if !found {
		t.Error("Should find cached entry through session")
	}

	if len(retrieved.Props) != 1 {
		t.Errorf("Expected 1 property, got %d", len(retrieved.Props))
	}
}

func TestCssCache_FullIntegration(t *testing.T) {
	cache1 := NewCssParseCache()
	cache2 := NewCssParseCache()

	testEntry := &CssCacheEntry{
		Props: CssPropsMap{
			"--test": M.CssCustomProperty{
				FullyQualified: M.FullyQualified{Name: "--test"},
			},
		},
		VarDefaults: map[string]bool{"--test": true},
	}

	cache1.Set("/test1.css", testEntry)
	cache2.Set("/test2.css", testEntry)

	_, found := cache1.Get("/test2.css")
	if found {
		t.Error("Cache1 should not have cache2's entries")
	}

	_, found = cache2.Get("/test1.css")
	if found {
		t.Error("Cache2 should not have cache1's entries")
	}
}
