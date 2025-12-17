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

package transform_test

import (
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// TestParseTsConfig_SrcDist tests basic rootDir/outDir extraction
func TestParseTsConfig_SrcDist(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")

	mappings, err := transform.ParseTsConfig("/test/tsconfig.json", fs)
	if err != nil {
		t.Fatalf("ParseTsConfig failed: %v", err)
	}

	// Should create mapping: /dist/ -> /src/
	expected := map[string]string{
		"/dist/": "/src/",
	}

	if len(mappings) != len(expected) {
		t.Errorf("Expected %d mappings, got %d", len(expected), len(mappings))
	}

	for key, expectedVal := range expected {
		if actualVal, ok := mappings[key]; !ok {
			t.Errorf("Missing mapping for %s", key)
		} else if actualVal != expectedVal {
			t.Errorf("For key %s: expected %s, got %s", key, expectedVal, actualVal)
		}
	}
}

// TestParseTsConfig_Extends tests tsconfig inheritance
func TestParseTsConfig_Extends(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/tsconfig-extends", "/test")

	mappings, err := transform.ParseTsConfig("/test/tsconfig.json", fs)
	if err != nil {
		t.Fatalf("ParseTsConfig failed: %v", err)
	}

	// Base config has rootDir: "./source"
	// Extending config has outDir: "./build"
	// Should create mapping: /build/ -> /source/
	expected := map[string]string{
		"/build/": "/source/",
	}

	if len(mappings) != len(expected) {
		t.Errorf("Expected %d mappings, got %d", len(expected), len(mappings))
	}

	for key, expectedVal := range expected {
		if actualVal, ok := mappings[key]; !ok {
			t.Errorf("Missing mapping for %s", key)
		} else if actualVal != expectedVal {
			t.Errorf("For key %s: expected %s, got %s", key, expectedVal, actualVal)
		}
	}
}

// TestParseTsConfig_DefaultValues tests that missing rootDir/outDir default to "."
func TestParseTsConfig_DefaultValues(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Create minimal tsconfig with no rootDir/outDir
	fs.AddFile("/test/tsconfig.json", `{"compilerOptions": {}}`, 0644)

	mappings, err := transform.ParseTsConfig("/test/tsconfig.json", fs)
	if err != nil {
		t.Fatalf("ParseTsConfig failed: %v", err)
	}

	// Both default to ".", so no mapping should be created
	// (no src/dist separation)
	if len(mappings) != 0 {
		t.Errorf("Expected empty mappings for in-place compilation, got %d mappings: %v", len(mappings), mappings)
	}
}

// TestParseTsConfig_RelativePaths tests normalization of relative paths
func TestParseTsConfig_RelativePaths(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Create tsconfig with various relative path formats
	fs.AddFile("/test/tsconfig.json", `{
		"compilerOptions": {
			"rootDir": "./src",
			"outDir": "./dist"
		}
	}`, 0644)

	mappings, err := transform.ParseTsConfig("/test/tsconfig.json", fs)
	if err != nil {
		t.Fatalf("ParseTsConfig failed: %v", err)
	}

	// Paths should be normalized to /dist/ and /src/
	expected := map[string]string{
		"/dist/": "/src/",
	}

	for key, expectedVal := range expected {
		if actualVal, ok := mappings[key]; !ok {
			t.Errorf("Missing mapping for %s", key)
		} else if actualVal != expectedVal {
			t.Errorf("For key %s: expected %s, got %s", key, expectedVal, actualVal)
		}
	}
}

// TestParseTsConfig_MissingFile tests error handling for non-existent tsconfig
func TestParseTsConfig_MissingFile(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	_, err := transform.ParseTsConfig("/test/nonexistent.json", fs)
	if err == nil {
		t.Error("Expected error for missing file, got nil")
	}
}

// TestParseTsConfig_InvalidJSON tests error handling for malformed JSON
func TestParseTsConfig_InvalidJSON(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Create invalid JSON
	fs.AddFile("/test/tsconfig.json", `{ invalid json }`, 0644)

	_, err := transform.ParseTsConfig("/test/tsconfig.json", fs)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

// TestParseTsConfig_CircularExtends tests detection of circular inheritance
func TestParseTsConfig_CircularExtends(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Create circular extends: a.json -> b.json -> a.json
	fs.AddFile("/test/a.json", `{"extends": "./b.json"}`, 0644)
	fs.AddFile("/test/b.json", `{"extends": "./a.json"}`, 0644)

	_, err := transform.ParseTsConfig("/test/a.json", fs)
	if err == nil {
		t.Error("Expected error for circular extends, got nil")
	}

	if err != nil && !strings.Contains(err.Error(), "circular") {
		t.Errorf("Expected 'circular' in error message, got: %v", err)
	}
}

// TestParseTsConfig_MaxDepthExceeded tests max depth protection
func TestParseTsConfig_MaxDepthExceeded(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Create chain longer than max depth (5)
	fs.AddFile("/test/1.json", `{"extends": "./2.json"}`, 0644)
	fs.AddFile("/test/2.json", `{"extends": "./3.json"}`, 0644)
	fs.AddFile("/test/3.json", `{"extends": "./4.json"}`, 0644)
	fs.AddFile("/test/4.json", `{"extends": "./5.json"}`, 0644)
	fs.AddFile("/test/5.json", `{"extends": "./6.json"}`, 0644)
	fs.AddFile("/test/6.json", `{"extends": "./7.json"}`, 0644)
	fs.AddFile("/test/7.json", `{"compilerOptions": {}}`, 0644)

	_, err := transform.ParseTsConfig("/test/1.json", fs)
	if err == nil {
		t.Error("Expected error for max depth exceeded, got nil")
	}

	if err != nil && !strings.Contains(err.Error(), "depth exceeded") {
		t.Errorf("Expected 'depth exceeded' in error message, got: %v", err)
	}
}

// TestParseTsConfig_ExtendsPathResolution tests that extends paths are resolved relative to config file
func TestParseTsConfig_ExtendsPathResolution(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Create nested directory structure
	fs.AddFile("/test/packages/app/tsconfig.json", `{"extends": "../../tsconfig.base.json"}`, 0644)
	fs.AddFile("/test/tsconfig.base.json", `{
		"compilerOptions": {
			"rootDir": "./src",
			"outDir": "./dist"
		}
	}`, 0644)

	mappings, err := transform.ParseTsConfig("/test/packages/app/tsconfig.json", fs)
	if err != nil {
		t.Fatalf("ParseTsConfig failed: %v", err)
	}

	// Should inherit rootDir and outDir from base
	expected := map[string]string{
		"/dist/": "/src/",
	}

	for key, expectedVal := range expected {
		if actualVal, ok := mappings[key]; !ok {
			t.Errorf("Missing mapping for %s", key)
		} else if actualVal != expectedVal {
			t.Errorf("For key %s: expected %s, got %s", key, expectedVal, actualVal)
		}
	}
}

// TestParseTsConfig_OverrideInheritedValues tests that child config overrides base config
func TestParseTsConfig_OverrideInheritedValues(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Base has both rootDir and outDir
	fs.AddFile("/test/tsconfig.base.json", `{
		"compilerOptions": {
			"rootDir": "./source",
			"outDir": "./output"
		}
	}`, 0644)

	// Child overrides outDir
	fs.AddFile("/test/tsconfig.json", `{
		"extends": "./tsconfig.base.json",
		"compilerOptions": {
			"outDir": "./build"
		}
	}`, 0644)

	mappings, err := transform.ParseTsConfig("/test/tsconfig.json", fs)
	if err != nil {
		t.Fatalf("ParseTsConfig failed: %v", err)
	}

	// Should use overridden outDir with inherited rootDir
	expected := map[string]string{
		"/build/": "/source/",
	}

	for key, expectedVal := range expected {
		if actualVal, ok := mappings[key]; !ok {
			t.Errorf("Missing mapping for %s", key)
		} else if actualVal != expectedVal {
			t.Errorf("For key %s: expected %s, got %s", key, expectedVal, actualVal)
		}
	}
}

// TestParseTsConfig_DefaultRootDirWithCustomOutDir tests edge case where rootDir="." with different outDir
func TestParseTsConfig_DefaultRootDirWithCustomOutDir(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Create tsconfig with rootDir="." and outDir="dist"
	fs.AddFile("/test/tsconfig.json", `{
		"compilerOptions": {
			"rootDir": ".",
			"outDir": "dist"
		}
	}`, 0644)

	mappings, err := transform.ParseTsConfig("/test/tsconfig.json", fs)
	if err != nil {
		t.Fatalf("ParseTsConfig failed: %v", err)
	}

	// Should create mapping from /dist/ to /./
	// PathResolver.fileExists() normalizes this correctly via filepath.Clean()
	expected := map[string]string{
		"/dist/": "/./",
	}

	if len(mappings) != len(expected) {
		t.Errorf("Expected %d mappings, got %d", len(expected), len(mappings))
	}

	for key, expectedVal := range expected {
		if actualVal, ok := mappings[key]; !ok {
			t.Errorf("Missing mapping for %s", key)
		} else if actualVal != expectedVal {
			t.Errorf("For key %s: expected %s, got %s", key, expectedVal, actualVal)
		}
	}
}
