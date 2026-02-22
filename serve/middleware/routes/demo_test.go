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

package routes

import (
	"fmt"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware"
)

// TestBuildDemoRoutingTable_DirectoryTraversalPrevention tests that malicious
// file paths attempting directory traversal are rejected
func TestBuildDemoRoutingTable_DirectoryTraversalPrevention(t *testing.T) {
	tests := []struct {
		name         string
		demoURL      string
		shouldReject bool
		description  string
	}{
		{
			name:         "relative path with parent traversal",
			demoURL:      "../../etc/passwd",
			shouldReject: true,
			description:  "Should reject paths with .. attempting to escape root",
		},
		{
			name:         "absolute path",
			demoURL:      "/etc/passwd",
			shouldReject: false, // Absolute paths get leading slash stripped, then become relative
			description:  "Absolute paths should be converted to relative",
		},
		{
			name:         "mixed traversal",
			demoURL:      "valid/path/../../../escape.html",
			shouldReject: true,
			description:  "Should reject paths that resolve to parent directories",
		},
		{
			name:         "legitimate relative path",
			demoURL:      "./demo/index.html",
			shouldReject: false,
			description:  "Should accept normal relative paths",
		},
		{
			name:         "legitimate path without prefix",
			demoURL:      "demo/index.html",
			shouldReject: false,
			description:  "Should accept paths without ./ prefix",
		},
		{
			name:         "path with internal parent reference",
			demoURL:      "components/../demo/index.html",
			shouldReject: false,
			description:  "Should accept paths that resolve within the root",
		},
		{
			name:         "sneaky double-dot at start",
			demoURL:      "../demo/index.html",
			shouldReject: true,
			description:  "Should reject single-level parent traversal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a minimal manifest JSON directly with a demo using the test URL
			manifestJSON := fmt.Sprintf(`{
				"schemaVersion": "1.0.0",
				"modules": [
					{
						"kind": "javascript-module",
						"path": "src/my-element.js",
						"declarations": [
							{
								"kind": "class",
								"name": "MyElement",
								"tagName": "my-element",
								"customElement": true,
								"demos": [
									{
										"url": %q
									}
								]
							}
						]
					}
				]
			}`, tt.demoURL)

			// Call BuildDemoRoutingTable
			routes, err := BuildDemoRoutingTable([]byte(manifestJSON), "")

			if tt.shouldReject {
				if err == nil {
					t.Errorf("Expected error for %s, but got none. Routes: %+v", tt.description, routes)
				} else if !strings.Contains(err.Error(), "directory traversal") {
					t.Errorf("Expected directory traversal error for %s, got: %v", tt.description, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for %s, but got: %v", tt.description, err)
				}
				if len(routes) == 0 {
					t.Errorf("Expected routes to be created for %s", tt.description)
				}
			}
		})
	}
}

// TestBuildDemoRoutingTable_DuplicateDetection tests that duplicate routes are detected
func TestBuildDemoRoutingTable_DuplicateDetection(t *testing.T) {
	// Create manifest with two demos that have the same URL
	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [
			{
				"kind": "javascript-module",
				"path": "src/element-a.js",
				"declarations": [
					{
						"kind": "class",
						"name": "ElementA",
						"tagName": "element-a",
						"customElement": true,
						"demos": [
							{
								"url": "/demo/test/"
							}
						]
					}
				]
			},
			{
				"kind": "javascript-module",
				"path": "src/element-b.js",
				"declarations": [
					{
						"kind": "class",
						"name": "ElementB",
						"tagName": "element-b",
						"customElement": true,
						"demos": [
							{
								"url": "/demo/test/"
							}
						]
					}
				]
			}
		]
	}`

	_, err := BuildDemoRoutingTable([]byte(manifestJSON), "")

	if err == nil {
		t.Fatal("Expected error for duplicate routes, but got none")
	}

	expectedSubstrings := []string{
		"duplicate demo route",
		"/demo/test/",
		"element-a",
		"element-b",
	}

	for _, substr := range expectedSubstrings {
		if !strings.Contains(err.Error(), substr) {
			t.Errorf("Expected error to contain %q, got: %v", substr, err)
		}
	}
}

// TestBuildPackageRoutingTable_DuplicateDetection tests that duplicate routes
// within a single package are detected
func TestBuildPackageRoutingTable_DuplicateDetection(t *testing.T) {
	// Create manifest with two demos that have the same URL
	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [
			{
				"kind": "javascript-module",
				"path": "src/element-a.js",
				"declarations": [
					{
						"kind": "class",
						"name": "ElementA",
						"tagName": "element-a",
						"customElement": true,
						"demos": [
							{
								"url": "/demo/test/"
							}
						]
					}
				]
			},
			{
				"kind": "javascript-module",
				"path": "src/element-b.js",
				"declarations": [
					{
						"kind": "class",
						"name": "ElementB",
						"tagName": "element-b",
						"customElement": true,
						"demos": [
							{
								"url": "/demo/test/"
							}
						]
					}
				]
			}
		]
	}`

	pkg := PackageContext{
		Name:     "test-package",
		Path:     "/path/to/test-package",
		Manifest: []byte(manifestJSON),
	}

	_, err := buildPackageRoutingTable(pkg)

	if err == nil {
		t.Fatal("Expected error for duplicate routes, but got none")
	}

	expectedSubstrings := []string{
		"duplicate demo route",
		"/demo/test/",
		"test-package",
		"/path/to/test-package",
		"element-a",
		"element-b",
	}

	for _, substr := range expectedSubstrings {
		if !strings.Contains(err.Error(), substr) {
			t.Errorf("Expected error to contain %q, got: %v", substr, err)
		}
	}
}

// TestBuildDemoRoutingTable_SourceHrefPaths tests that demo source.href paths
// are correctly resolved, especially when they are file paths rather than URLs
func TestBuildDemoRoutingTable_SourceHrefPaths(t *testing.T) {
	tests := []struct {
		name         string
		sourceHref   string
		expectedPath string
		description  string
	}{
		{
			name:         "absolute path with leading slash",
			sourceHref:   "/src/components/button/demos/index.html",
			expectedPath: "src/components/button/demos/index.html",
			description:  "Leading slash should be stripped from file paths",
		},
		{
			name:         "relative path without leading slash",
			sourceHref:   "src/components/button/demos/index.html",
			expectedPath: "src/components/button/demos/index.html",
			description:  "Relative paths should be used as-is",
		},
		{
			name:         "path with dot-slash prefix",
			sourceHref:   "./demos/index.html",
			expectedPath: "demos/index.html",
			description:  "Dot-slash prefix should be handled",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create manifest with source.href
			manifestJSON := fmt.Sprintf(`{
				"schemaVersion": "1.0.0",
				"modules": [
					{
						"kind": "javascript-module",
						"path": "src/components/button/button.js",
						"declarations": [
							{
								"kind": "class",
								"name": "WaButton",
								"tagName": "wa-button",
								"customElement": true,
								"demos": [
									{
										"url": "https://example.com/demos/button/",
										"source": {
											"href": %q
										}
									}
								]
							}
						]
					}
				]
			}`, tt.sourceHref)

			routes, err := BuildDemoRoutingTable([]byte(manifestJSON), "")
			if err != nil {
				t.Fatalf("BuildDemoRoutingTable failed: %v", err)
			}

			// Find the route
			var foundPath string
			for _, entry := range routes {
				if entry != nil {
					foundPath = entry.FilePath
					break
				}
			}

			if foundPath != tt.expectedPath {
				t.Errorf("%s: expected path %q, got %q", tt.description, tt.expectedPath, foundPath)
			}
		})
	}
}

// TestResolveViaDemoSubresource tests that subresources referenced from
// directory-style demo URLs are resolved to the demo file's parent directory
func TestResolveViaDemoSubresource(t *testing.T) {
	tests := []struct {
		name        string
		requestPath string
		watchDir    string
		isWorkspace bool
		demoRoutes  map[string]*middleware.DemoRouteEntry
		expected    string
	}{
		{
			name:        "flat-file subresource single-package",
			requestPath: "/elements/avatar/demo/sizes/perlman.jpg",
			watchDir:    "/test",
			demoRoutes: map[string]*middleware.DemoRouteEntry{
				"/elements/avatar/demo/sizes/": {
					LocalRoute: "/elements/avatar/demo/sizes/",
					TagName:    "rh-avatar",
					FilePath:   "elements/rh-avatar/demo/sizes.html",
				},
			},
			expected: "/test/elements/rh-avatar/demo/perlman.jpg",
		},
		{
			name:        "nested subresource",
			requestPath: "/elements/avatar/demo/sizes/images/foo.jpg",
			watchDir:    "/test",
			demoRoutes: map[string]*middleware.DemoRouteEntry{
				"/elements/avatar/demo/sizes/": {
					LocalRoute: "/elements/avatar/demo/sizes/",
					TagName:    "rh-avatar",
					FilePath:   "elements/rh-avatar/demo/sizes.html",
				},
			},
			expected: "/test/elements/rh-avatar/demo/images/foo.jpg",
		},
		{
			name:        "path traversal rejection",
			requestPath: "/elements/avatar/demo/sizes/../../../etc/passwd",
			watchDir:    "/test",
			demoRoutes: map[string]*middleware.DemoRouteEntry{
				"/elements/avatar/demo/sizes/": {
					LocalRoute: "/elements/avatar/demo/sizes/",
					TagName:    "rh-avatar",
					FilePath:   "elements/rh-avatar/demo/sizes.html",
				},
			},
			expected: "",
		},
		{
			name:        "no match",
			requestPath: "/unrelated/path/image.jpg",
			watchDir:    "/test",
			demoRoutes: map[string]*middleware.DemoRouteEntry{
				"/elements/avatar/demo/sizes/": {
					LocalRoute: "/elements/avatar/demo/sizes/",
					TagName:    "rh-avatar",
					FilePath:   "elements/rh-avatar/demo/sizes.html",
				},
			},
			expected: "",
		},
		{
			name:        "workspace mode uses PackagePath",
			requestPath: "/elements/avatar/demo/sizes/perlman.jpg",
			isWorkspace: true,
			demoRoutes: map[string]*middleware.DemoRouteEntry{
				"/elements/avatar/demo/sizes/": {
					LocalRoute:  "/elements/avatar/demo/sizes/",
					TagName:     "rh-avatar",
					FilePath:    "elements/rh-avatar/demo/sizes.html",
					PackagePath: "/workspace/packages/avatars",
				},
			},
			expected: "/workspace/packages/avatars/elements/rh-avatar/demo/perlman.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Context: &mockContext{
					watchDir:    tt.watchDir,
					demoRoutes:  tt.demoRoutes,
					isWorkspace: tt.isWorkspace,
				},
			}
			got := resolveViaDemoSubresource(tt.requestPath, cfg)
			if got != tt.expected {
				t.Errorf("resolveViaDemoSubresource(%q) = %q, want %q", tt.requestPath, got, tt.expected)
			}
		})
	}
}
