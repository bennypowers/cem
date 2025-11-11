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

package transform

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCacheStats_TracksHitsAndMisses(t *testing.T) {
	cache := NewCache(1024 * 1024) // 1MB cache

	key := CacheKey{Path: "test.ts", ModTime: time.Now(), Size: 100}

	// Miss on first get
	_, found := cache.Get(key)
	if found {
		t.Fatal("Expected cache miss, got hit")
	}

	// Add entry
	cache.Set(key, []byte("code"), nil)

	// Hit on second get
	_, found = cache.Get(key)
	if !found {
		t.Fatal("Expected cache hit, got miss")
	}

	// Check stats
	stats := cache.Stats()
	if stats.Hits != 1 {
		t.Errorf("Expected 1 hit, got %d", stats.Hits)
	}
	if stats.Misses != 1 {
		t.Errorf("Expected 1 miss, got %d", stats.Misses)
	}
	if stats.HitRate < 49.9 || stats.HitRate > 50.1 {
		t.Errorf("Expected ~50%% hit rate, got %.2f%%", stats.HitRate)
	}
}

func TestCacheStats_TracksEvictions(t *testing.T) {
	cache := NewCache(100) // Tiny cache to force evictions

	// Add entries that exceed cache size
	for i := 0; i < 5; i++ {
		key := CacheKey{Path: "test.ts", ModTime: time.Now().Add(time.Duration(i) * time.Second), Size: 50}
		cache.Set(key, make([]byte, 50), nil)
	}

	// Check that evictions occurred
	stats := cache.Stats()
	if stats.Evictions == 0 {
		t.Fatal("Expected evictions to be tracked, got 0")
	}
	if stats.SizeBytes > 100 {
		t.Errorf("Cache size %d exceeds max size 100", stats.SizeBytes)
	}
}

func TestCache_LogsStatsWhenCalled(t *testing.T) {
	// Create mock logger
	logged := false
	var logMessage string
	mockLogger := &mockLogger{
		infoFunc: func(msg string, args ...any) {
			logged = true
			logMessage = msg
		},
	}

	cache := NewCacheWithLogger(1024, mockLogger)

	// Add some entries
	key := CacheKey{Path: "test.ts", ModTime: time.Now(), Size: 100}
	cache.Set(key, []byte("code"), nil)
	cache.Get(key)

	// Log stats
	cache.LogStats()

	// Verify logger was called
	if !logged {
		t.Fatal("Expected stats to be logged")
	}
	if logMessage == "" {
		t.Fatal("Expected log message to be set")
	}
}

// mockLogger implements logger.Logger interface for testing
type mockLogger struct {
	infoFunc    func(msg string, args ...any)
	warningFunc func(msg string, args ...any)
	errorFunc   func(msg string, args ...any)
	debugFunc   func(msg string, args ...any)
}

func (m *mockLogger) Info(msg string, args ...any) {
	if m.infoFunc != nil {
		m.infoFunc(msg, args...)
	}
}

func (m *mockLogger) Warning(msg string, args ...any) {
	if m.warningFunc != nil {
		m.warningFunc(msg, args...)
	}
}

func (m *mockLogger) Error(msg string, args ...any) {
	if m.errorFunc != nil {
		m.errorFunc(msg, args...)
	}
}

func (m *mockLogger) Debug(msg string, args ...any) {
	if m.debugFunc != nil {
		m.debugFunc(msg, args...)
	}
}

func TestCache_CSSImportDependencies(t *testing.T) {
	cache := NewCache(10 * 1024 * 1024) // 10MB cache

	// Simulate a TypeScript file that imports a CSS file
	tsPath := "/test/component.ts"
	cssPath := "/test/component.css"
	tsCode := []byte(`import './component.css';\nexport class Component {}`)
	
	// Set the TS file with CSS as a dependency
	key := CacheKey{Path: tsPath, ModTime: time.Now(), Size: int64(len(tsCode))}
	cache.Set(key, tsCode, []string{cssPath})

	// Now when component.css changes, it should invalidate component.ts
	invalidated := cache.Invalidate(cssPath)
	
	if len(invalidated) == 0 {
		t.Fatal("Expected CSS file change to invalidate TypeScript file, got 0 invalidations")
	}
	
	found := false
	for _, path := range invalidated {
		if path == tsPath {
			found = true
			break
		}
	}
	
	if !found {
		t.Errorf("Expected %s to be invalidated when %s changes, got: %v", tsPath, cssPath, invalidated)
	}
}

func TestExtractDependencies_CSSImports(t *testing.T) {
	// Read fixture file with CSS import
	inputPath := filepath.Join("..", "..", "testdata", "transforms", "css-import", "input.ts")
	source, err := os.ReadFile(inputPath)
	if err != nil {
		t.Fatalf("Failed to read test fixture: %v", err)
	}

	// Extract dependencies from the source
	absPath, _ := filepath.Abs(inputPath)

	deps := extractDependencies(source, absPath)

	if len(deps) == 0 {
		t.Fatal("Expected CSS import to be extracted as dependency, got 0 dependencies")
	}

	foundCSS := false
	for _, dep := range deps {
		if strings.HasSuffix(dep, "component.css") {
			foundCSS = true
			break
		}
	}

	if !foundCSS {
		t.Errorf("Expected to find component.css in dependencies, got: %v", deps)
	}
}

func TestCache_CSSImportDependenciesExtracted(t *testing.T) {
	// Integration test: Verify that CSS imports are automatically extracted and tracked
	cache := NewCache(10 * 1024 * 1024)

	// Read fixture file with CSS import
	inputPath := filepath.Join("..", "..", "testdata", "transforms", "css-import", "input.ts")
	source, err := os.ReadFile(inputPath)
	if err != nil {
		t.Fatalf("Failed to read test fixture: %v", err)
	}

	absPath, _ := filepath.Abs(inputPath)
	cssPath, _ := filepath.Abs(filepath.Join(filepath.Dir(inputPath), "component.css"))

	// Extract dependencies and cache with them
	deps := extractDependencies(source, absPath)
	key := CacheKey{Path: absPath, ModTime: time.Now(), Size: int64(len(source))}
	cache.Set(key, source, deps)

	// Verify CSS file change invalidates TypeScript file
	invalidated := cache.Invalidate(cssPath)

	if len(invalidated) == 0 {
		t.Fatal("Expected CSS file change to invalidate TypeScript file via extracted dependencies")
	}

	found := false
	for _, path := range invalidated {
		if path == absPath {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected %s to be invalidated when %s changes, got: %v", absPath, cssPath, invalidated)
	}
}
