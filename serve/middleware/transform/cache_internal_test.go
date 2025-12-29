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
	"slices"
	"strings"
	"testing"
	"time"
)

// This file contains white-box tests that access unexported functions/types.
// Most cache tests are in cache_test.go using the transform_test package.

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

	found := slices.Contains(invalidated, absPath)

	if !found {
		t.Errorf("Expected %s to be invalidated when %s changes, got: %v", absPath, cssPath, invalidated)
	}
}
