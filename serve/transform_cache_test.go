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

package serve

import (
	"testing"
	"time"
)

// TestTransformCache_BasicOperations tests basic cache operations
func TestTransformCache_BasicOperations(t *testing.T) {
	cache := NewTransformCache(1024 * 1024) // 1MB cache

	key := CacheKey{
		Path:    "/test/file.ts",
		ModTime: time.Now(),
		Size:    100,
	}

	code := []byte("console.log('test');")
	deps := []string{"/test/dep.ts"}

	// Test cache miss
	if _, found := cache.Get(key); found {
		t.Error("Expected cache miss, got hit")
	}

	// Test cache set
	cache.Set(key, code, deps)

	// Test cache hit
	entry, found := cache.Get(key)
	if !found {
		t.Fatal("Expected cache hit, got miss")
	}

	if string(entry.Code) != string(code) {
		t.Errorf("Expected code %q, got %q", string(code), string(entry.Code))
	}

	if len(entry.Dependencies) != 1 || entry.Dependencies[0] != "/test/dep.ts" {
		t.Errorf("Expected dependencies [/test/dep.ts], got %v", entry.Dependencies)
	}
}

// TestTransformCache_Invalidation tests cache invalidation
func TestTransformCache_Invalidation(t *testing.T) {
	cache := NewTransformCache(1024 * 1024)

	// Create cache entry
	key := CacheKey{
		Path:    "/test/file.ts",
		ModTime: time.Now(),
		Size:    100,
	}

	cache.Set(key, []byte("code"), []string{})

	// Verify it's cached
	if _, found := cache.Get(key); !found {
		t.Fatal("Expected cache hit before invalidation")
	}

	// Invalidate
	invalidated := cache.Invalidate("/test/file.ts")
	if len(invalidated) != 1 || invalidated[0] != "/test/file.ts" {
		t.Errorf("Expected invalidated [/test/file.ts], got %v", invalidated)
	}

	// Verify it's no longer cached
	if _, found := cache.Get(key); found {
		t.Error("Expected cache miss after invalidation")
	}
}

// TestTransformCache_TransitiveInvalidation tests transitive dependency invalidation
func TestTransformCache_TransitiveInvalidation(t *testing.T) {
	cache := NewTransformCache(1024 * 1024)

	// Create dependency chain: a.ts imports b.ts imports c.ts
	keyA := CacheKey{Path: "/test/a.ts", ModTime: time.Now(), Size: 100}
	keyB := CacheKey{Path: "/test/b.ts", ModTime: time.Now(), Size: 100}
	keyC := CacheKey{Path: "/test/c.ts", ModTime: time.Now(), Size: 100}

	cache.Set(keyC, []byte("// c.ts"), []string{})              // c has no deps
	cache.Set(keyB, []byte("// b.ts"), []string{"/test/c.ts"}) // b imports c
	cache.Set(keyA, []byte("// a.ts"), []string{"/test/b.ts"}) // a imports b

	// Verify all are cached
	if _, found := cache.Get(keyA); !found {
		t.Fatal("Expected a.ts to be cached")
	}
	if _, found := cache.Get(keyB); !found {
		t.Fatal("Expected b.ts to be cached")
	}
	if _, found := cache.Get(keyC); !found {
		t.Fatal("Expected c.ts to be cached")
	}

	// Invalidate c.ts - should also invalidate b.ts and a.ts
	invalidated := cache.Invalidate("/test/c.ts")

	// Check that all three files were invalidated
	expectedInvalidated := map[string]bool{
		"/test/a.ts": true,
		"/test/b.ts": true,
		"/test/c.ts": true,
	}

	if len(invalidated) != 3 {
		t.Errorf("Expected 3 invalidated files, got %d: %v", len(invalidated), invalidated)
	}

	for _, path := range invalidated {
		if !expectedInvalidated[path] {
			t.Errorf("Unexpected invalidated path: %s", path)
		}
		delete(expectedInvalidated, path)
	}

	if len(expectedInvalidated) > 0 {
		t.Errorf("Missing invalidated paths: %v", expectedInvalidated)
	}

	// Verify all entries were removed from cache
	if _, found := cache.Get(keyA); found {
		t.Error("Expected a.ts to be invalidated")
	}
	if _, found := cache.Get(keyB); found {
		t.Error("Expected b.ts to be invalidated")
	}
	if _, found := cache.Get(keyC); found {
		t.Error("Expected c.ts to be invalidated")
	}
}

// TestTransformCache_LRUEviction tests LRU cache eviction
func TestTransformCache_LRUEviction(t *testing.T) {
	// Create small cache (100 bytes)
	cache := NewTransformCache(100)

	// Add entries that exceed cache size
	key1 := CacheKey{Path: "/test/1.ts", ModTime: time.Now(), Size: 50}
	key2 := CacheKey{Path: "/test/2.ts", ModTime: time.Now(), Size: 50}
	key3 := CacheKey{Path: "/test/3.ts", ModTime: time.Now(), Size: 50}

	code := make([]byte, 50) // 50 bytes each

	cache.Set(key1, code, []string{})
	cache.Set(key2, code, []string{})

	// key1 should be evicted when we add key3 (exceeds 100 bytes)
	cache.Set(key3, code, []string{})

	// Check that key1 was evicted
	if _, found := cache.Get(key1); found {
		t.Error("Expected key1 to be evicted")
	}

	// Check that key2 and key3 are still cached
	if _, found := cache.Get(key2); !found {
		t.Error("Expected key2 to be cached")
	}
	if _, found := cache.Get(key3); !found {
		t.Error("Expected key3 to be cached")
	}
}

// TestTransformCache_Stats tests cache statistics
func TestTransformCache_Stats(t *testing.T) {
	cache := NewTransformCache(1024 * 1024)

	key := CacheKey{Path: "/test/file.ts", ModTime: time.Now(), Size: 100}
	code := []byte("console.log('test');")

	// Initial stats
	stats := cache.Stats()
	if stats.Hits != 0 || stats.Misses != 0 {
		t.Errorf("Expected 0 hits/misses initially, got %d/%d", stats.Hits, stats.Misses)
	}

	// Cache miss
	cache.Get(key)
	stats = cache.Stats()
	if stats.Misses != 1 {
		t.Errorf("Expected 1 miss, got %d", stats.Misses)
	}

	// Cache set
	cache.Set(key, code, []string{})

	// Cache hit
	cache.Get(key)
	stats = cache.Stats()
	if stats.Hits != 1 {
		t.Errorf("Expected 1 hit, got %d", stats.Hits)
	}

	// Hit rate should be 50% (1 hit, 1 miss)
	if stats.HitRate < 49.0 || stats.HitRate > 51.0 {
		t.Errorf("Expected ~50%% hit rate, got %.1f%%", stats.HitRate)
	}

	// Check entry count
	if stats.Entries != 1 {
		t.Errorf("Expected 1 entry, got %d", stats.Entries)
	}
}

// TestTransformCache_CircularDependency tests circular dependency handling
func TestTransformCache_CircularDependency(t *testing.T) {
	cache := NewTransformCache(1024 * 1024)

	// Create circular dependency: a.ts -> b.ts -> a.ts
	keyA := CacheKey{Path: "/test/a.ts", ModTime: time.Now(), Size: 100}
	keyB := CacheKey{Path: "/test/b.ts", ModTime: time.Now(), Size: 100}

	cache.Set(keyA, []byte("// a.ts"), []string{"/test/b.ts"}) // a imports b
	cache.Set(keyB, []byte("// b.ts"), []string{"/test/a.ts"}) // b imports a (circular!)

	// Invalidate a.ts - should handle circular dependency gracefully
	invalidated := cache.Invalidate("/test/a.ts")

	// Both files should be invalidated exactly once
	if len(invalidated) != 2 {
		t.Errorf("Expected 2 invalidated files, got %d: %v", len(invalidated), invalidated)
	}

	// Verify both entries were removed from cache
	if _, found := cache.Get(keyA); found {
		t.Error("Expected a.ts to be invalidated")
	}
	if _, found := cache.Get(keyB); found {
		t.Error("Expected b.ts to be invalidated")
	}
}

// TestTransformCache_MultipleEntriesSamePath tests that all cache entries with the same path are invalidated
func TestTransformCache_MultipleEntriesSamePath(t *testing.T) {
	cache := NewTransformCache(1024 * 1024)

	// Create multiple cache entries for same path with different mod times
	// This simulates rapid file changes where old entries haven't been evicted yet
	path := "/test/file.ts"
	baseTime := time.Now()

	key1 := CacheKey{Path: path, ModTime: baseTime, Size: 100}
	key2 := CacheKey{Path: path, ModTime: baseTime.Add(1 * time.Second), Size: 100}
	key3 := CacheKey{Path: path, ModTime: baseTime.Add(2 * time.Second), Size: 100}

	// Set all three entries (different keys, same path)
	cache.Set(key1, []byte("version 1"), []string{})
	cache.Set(key2, []byte("version 2"), []string{})
	cache.Set(key3, []byte("version 3"), []string{})

	// Verify all are in cache
	if _, found := cache.Get(key1); !found {
		t.Error("Expected key1 to be cached")
	}
	if _, found := cache.Get(key2); !found {
		t.Error("Expected key2 to be cached")
	}
	if _, found := cache.Get(key3); !found {
		t.Error("Expected key3 to be cached")
	}

	// Invalidate the path - should remove ALL entries with this path
	invalidated := cache.Invalidate(path)

	// Should report path only once in invalidated list
	if len(invalidated) != 1 {
		t.Errorf("Expected path to appear once in invalidated list, got %d times: %v", len(invalidated), invalidated)
	}
	if invalidated[0] != path {
		t.Errorf("Expected invalidated path %q, got %q", path, invalidated[0])
	}

	// Verify ALL entries were removed from cache
	if _, found := cache.Get(key1); found {
		t.Error("Expected key1 to be invalidated")
	}
	if _, found := cache.Get(key2); found {
		t.Error("Expected key2 to be invalidated")
	}
	if _, found := cache.Get(key3); found {
		t.Error("Expected key3 to be invalidated")
	}

	// Verify cache is empty
	stats := cache.Stats()
	if stats.Entries != 0 {
		t.Errorf("Expected 0 cache entries after invalidation, got %d", stats.Entries)
	}
}
