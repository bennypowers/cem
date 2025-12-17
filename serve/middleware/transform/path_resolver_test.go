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
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// TestPathResolver_CoLocated tests backward compatibility with in-place compilation
// where .ts and .js files are in the same directory
func TestPathResolver_CoLocated(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	resolver := transform.NewPathResolver("/test", nil, fs, nil)

	// Test co-located file resolution
	result := resolver.ResolveTsSource("/components/button.js")
	expected := "/components/button.ts"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestPathResolver_ExplicitSrcDistMapping tests src/dist with explicit path mapping
func TestPathResolver_ExplicitSrcDistMapping(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")

	// Explicit mapping (from tsconfig.json parsing)
	resolver := transform.NewPathResolver("/test", map[string]string{
		"/dist/": "/src/",
	}, fs, nil)

	// Test explicit mapping: /dist/ -> /src/
	result := resolver.ResolveTsSource("/dist/components/widget.js")
	expected := "/src/components/widget.ts"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestPathResolver_WorkspaceMode tests path resolution with package prefixes
func TestPathResolver_WorkspaceMode(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/workspace-mode", "/test")

	// Explicit mapping handles workspace paths
	resolver := transform.NewPathResolver("/test", map[string]string{
		"/packages/webawesome/dist/": "/packages/webawesome/src/",
	}, fs, nil)

	// Test workspace mode with package prefix
	// Request: /packages/webawesome/dist/components/alert.js
	// Should resolve to: /packages/webawesome/src/components/alert.ts
	result := resolver.ResolveTsSource("/packages/webawesome/dist/components/alert.js")
	expected := "/packages/webawesome/src/components/alert.ts"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestPathResolver_ExplicitMappings tests user-configured path mappings
func TestPathResolver_ExplicitMappings(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/explicit-mappings", "/test")

	resolver := transform.NewPathResolver("/test", map[string]string{
		"/lib/": "/sources/",
	}, fs, nil)

	// Test explicit mapping: /lib/ -> /sources/
	result := resolver.ResolveTsSource("/lib/helpers/formatter.js")
	expected := "/sources/helpers/formatter.ts"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestPathResolver_MultipleOutputDirs tests multiple dist directories mapping to same source
func TestPathResolver_MultipleOutputDirs(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")

	// Explicit mappings for both output directories
	resolver := transform.NewPathResolver("/test", map[string]string{
		"/dist/":     "/src/",
		"/dist-cdn/": "/src/",
	}, fs, nil)

	// Test both /dist/ and /dist-cdn/ resolve to /src/
	testCases := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "dist directory",
			request:  "/dist/components/widget.js",
			expected: "/src/components/widget.ts",
		},
		{
			name:     "dist-cdn directory",
			request:  "/dist-cdn/components/widget.js",
			expected: "/src/components/widget.ts",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := resolver.ResolveTsSource(tc.request)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

// TestPathResolver_NotFound tests that non-existent files return empty string
func TestPathResolver_NotFound(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	resolver := transform.NewPathResolver("/test", nil, fs, nil)

	// Test non-existent file
	result := resolver.ResolveTsSource("/components/nonexistent.js")

	if result != "" {
		t.Errorf("Expected empty string for non-existent file, got %s", result)
	}
}

// TestPathResolver_SecurityPathTraversal tests path traversal protection
func TestPathResolver_SecurityPathTraversal(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	resolver := transform.NewPathResolver("/test", nil, fs, nil)

	// Test path traversal attempts
	maliciousPaths := []string{
		"/../../../etc/passwd.js",
		"/components/../../../../../../etc/passwd.js",
		"/components/../../../button.js",
	}

	for _, path := range maliciousPaths {
		t.Run("traversal_"+path, func(t *testing.T) {
			result := resolver.ResolveTsSource(path)
			// Path traversal should be rejected, returning empty string
			if result != "" {
				t.Errorf("Path traversal not blocked for %s, got %s", path, result)
			}
		})
	}
}

// TestPathResolver_NonJsExtension tests that non-.js requests are ignored
func TestPathResolver_NonJsExtension(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	resolver := transform.NewPathResolver("/test", nil, fs, nil)

	// Test non-.js extensions
	testCases := []string{
		"/components/button.ts",  // Already TypeScript
		"/components/button.css", // CSS file
		"/components/button.html", // HTML file
		"/components/button",      // No extension
	}

	for _, path := range testCases {
		t.Run("non_js_"+path, func(t *testing.T) {
			result := resolver.ResolveTsSource(path)
			if result != "" {
				t.Errorf("Expected empty string for non-.js request %s, got %s", path, result)
			}
		})
	}
}

// TestPathResolver_PriorityOrder tests that resolution tries strategies in correct order
func TestPathResolver_PriorityOrder(t *testing.T) {
	// Create a fixture where both co-located and src/dist files exist
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")
	// Add a file in src/ to test priority
	fs.AddFile("/test/src/components/button.ts", "// src version", 0644)
	fs.AddFile("/test/dist/components/button.js", "// dist version", 0644)

	resolver := transform.NewPathResolver("/test", nil, fs, nil)

	// Request /components/button.js
	// Should find co-located /components/button.ts (Strategy 1)
	// NOT /src/components/button.ts (Strategy 3)
	result := resolver.ResolveTsSource("/components/button.js")
	expected := "/components/button.ts"

	if result != expected {
		t.Errorf("Co-located file should have priority. Expected %s, got %s", expected, result)
	}
}

// TestPathResolver_ExplicitMappingPriorityOverFallback tests explicit mappings override fallbacks
func TestPathResolver_ExplicitMappingPriorityOverFallback(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/explicit-mappings", "/test")
	// Also add a file that would match fallback pattern
	fs.AddFile("/test/src/helpers/formatter.ts", "// fallback version", 0644)

	resolver := transform.NewPathResolver("/test", map[string]string{
		"/lib/": "/sources/",
	}, fs, nil)

	// Request /lib/helpers/formatter.js
	// Should use explicit mapping to /sources/ (Strategy 2)
	// NOT fallback to /src/ (Strategy 3)
	result := resolver.ResolveTsSource("/lib/helpers/formatter.js")
	expected := "/sources/helpers/formatter.ts"

	if result != expected {
		t.Errorf("Explicit mapping should have priority. Expected %s, got %s", expected, result)
	}
}

// TestPathResolver_DefaultRootDirMapping tests edge case where rootDir="." produces "/./" prefix
// This tests that PathResolver correctly handles the "/./" mapping through filepath.Join() normalization
func TestPathResolver_DefaultRootDirMapping(t *testing.T) {
	// Create empty fixture to avoid co-located files interfering
	fs := testutil.NewFixtureFS(t, "path-mappings/explicit-mappings", "/test")
	// Add a source file at root level that only the mapping can find
	fs.AddFile("/test/button.ts", "// root level source", 0644)
	// Add a dist output file to make the request realistic
	fs.AddFile("/test/dist/button.js", "// compiled output", 0644)

	// Create resolver with mapping that would be generated from tsconfig with rootDir="." and outDir="dist"
	// This produces the edge case mapping: "/dist/" -> "/./"
	// filepath.Join() in PathResolver normalizes "/./" + "button.js" to "button.js"
	resolver := transform.NewPathResolver("/test", map[string]string{
		"/dist/": "/./",
	}, fs, nil)

	// Request /dist/button.js
	// The mapping "/dist/" -> "/./" combined with filepath.Join() normalization
	// produces /button.ts (filepath.Join("/./", "button") normalizes to "button")
	result := resolver.ResolveTsSource("/dist/button.js")
	expected := "/button.ts"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
