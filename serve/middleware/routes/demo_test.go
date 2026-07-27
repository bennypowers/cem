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
	"github.com/stretchr/testify/assert"
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
			routes, skipped, err := BuildDemoRoutingTable([]byte(manifestJSON), "", "")

			if err != nil {
				t.Errorf("Expected no error for %s, but got: %v", tt.description, err)
			}
			if tt.shouldReject {
				if len(routes) != 0 {
					t.Errorf("Expected traversal path to be skipped for %s, but got routes: %+v", tt.description, routes)
				}
				// Inline assertions: scalar field checks on security-critical rejection path
				if assert.Len(t, skipped, 1, "rejected demo should appear in skipped list") {
					assert.Equal(t, "my-element", skipped[0].TagName)
					assert.Equal(t, tt.demoURL, skipped[0].DemoURL)
					assert.Contains(t, skipped[0].Reason, "directory traversal")
				}
			} else {
				if len(routes) == 0 {
					t.Errorf("Expected routes to be created for %s", tt.description)
				}
				assert.Empty(t, skipped, "accepted demo should not appear in skipped list")
			}
		})
	}
}

func TestDemoURLPrefixFromTemplate(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{"full external URL", "https://example.com/docs/components/{{.tag}}/{{.demo}}/", "/docs/components"},
		{"path-only template", "/elements/{{.tag}}/demo/{{.demo}}/", "/elements"},
		{"no template markers", "https://example.com/static/path/", "/static/path"},
		{"root path", "https://example.com/{{.tag}}/", ""},
		{"empty", "", ""},
		{"relative template", "{{.tag}}/{{.demo}}/", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DemoURLPrefixFromTemplate(tt.template)
			if got != tt.want {
				t.Errorf("DemoURLPrefixFromTemplate(%q) = %q, want %q", tt.template, got, tt.want)
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

	_, _, err := BuildDemoRoutingTable([]byte(manifestJSON), "", "")

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

	_, _, err := buildPackageRoutingTable(pkg, "")

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

// Inline assertions: regression test for skipped-entry preservation on duplicate-route error
func TestBuildDemoRoutingTable_DuplicatePreservesSkipped(t *testing.T) {
	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [
			{
				"kind": "javascript-module",
				"path": "src/bad-element.js",
				"declarations": [
					{
						"kind": "class",
						"name": "BadElement",
						"tagName": "bad-element",
						"customElement": true,
						"demos": [
							{
								"url": "../../etc/passwd"
							}
						]
					}
				]
			},
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
								"url": "/demo/dup/"
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
								"url": "/demo/dup/"
							}
						]
					}
				]
			}
		]
	}`

	_, skipped, err := BuildDemoRoutingTable([]byte(manifestJSON), "", "")

	assert.Error(t, err, "should return duplicate-route error")
	assert.Contains(t, err.Error(), "duplicate demo route")
	if assert.Len(t, skipped, 1, "traversal skip before duplicate should be preserved") {
		assert.Equal(t, "bad-element", skipped[0].TagName)
	}
}

// Inline assertions: regression test for workspace aggregation preserving skipped across packages
func TestBuildWorkspaceRoutingTable_PreservesSkippedAcrossPackages(t *testing.T) {
	pkgWithTraversal := PackageContext{
		Name: "pkg-a",
		Path: "/path/to/pkg-a",
		Manifest: []byte(`{
			"schemaVersion": "1.0.0",
			"modules": [
				{
					"kind": "javascript-module",
					"path": "src/bad.js",
					"declarations": [
						{
							"kind": "class",
							"name": "Bad",
							"tagName": "bad-element",
							"customElement": true,
							"demos": [{
								"url": "https://example.com/demo/",
								"source": {"href": "../../etc/passwd"}
							}]
						}
					]
				}
			]
		}`),
	}
	pkgClean := PackageContext{
		Name: "pkg-b",
		Path: "/path/to/pkg-b",
		Manifest: []byte(`{
			"schemaVersion": "1.0.0",
			"modules": [
				{
					"kind": "javascript-module",
					"path": "src/good.js",
					"declarations": [
						{
							"kind": "class",
							"name": "Good",
							"tagName": "good-element",
							"customElement": true,
							"demos": [{"url": "./demo/index.html"}]
						}
					]
				}
			]
		}`),
	}

	routes, skipped, err := BuildWorkspaceRoutingTable(
		[]PackageContext{pkgWithTraversal, pkgClean}, "",
	)
	assert.NoError(t, err)
	assert.Len(t, routes, 1, "clean package should produce one route")
	if assert.Len(t, skipped, 1, "traversal skip from pkg-a should be preserved") {
		assert.Equal(t, "bad-element", skipped[0].TagName)
	}
}

// Inline assertions: pure function, scalar results
func TestIsExternalDemoURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{"https URL", "https://example.com/demo/", true},
		{"http URL", "http://example.com/demo/", true},
		{"localhost URL", "http://localhost:3000/demo/", true},
		{"relative path", "./demo/index.html", false},
		{"absolute path", "/demo/index.html", false},
		{"bare path", "demo/index.html", false},
		{"empty", "", false},
		{"file URI no host", "file:///local/path", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isExternalDemoURL(tt.url))
		})
	}
}

// Inline assertions: routing table skip behavior, scalar route map check
func TestBuildDemoRoutingTable_ExternalURLSkipped(t *testing.T) {
	manifestJSON := `{
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
								"url": "https://example.com/demo-field/",
								"description": "External demo"
							}
						]
					}
				]
			}
		]
	}`

	routes, _, err := BuildDemoRoutingTable([]byte(manifestJSON), "", "")
	assert.NoError(t, err)
	assert.Empty(t, routes, "external demo URL should not produce a route")
}

// Inline assertions: routing exception case, scalar route count check
func TestBuildDemoRoutingTable_ExternalURLWithSourceHrefStillRoutes(t *testing.T) {
	manifestJSON := `{
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
								"url": "https://example.com/demo-field/",
								"description": "External URL but local source",
								"source": {
									"href": "src/my-element/demo/index.html"
								}
							}
						]
					}
				]
			}
		]
	}`

	routes, _, err := BuildDemoRoutingTable([]byte(manifestJSON), "", "")
	assert.NoError(t, err)
	assert.Len(t, routes, 1, "external URL with Source.Href should still produce a route")
}

// Inline assertions: workspace-mode skip behavior, scalar route map check
func TestBuildPackageRoutingTable_ExternalURLSkipped(t *testing.T) {
	manifestJSON := `{
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
								"url": "https://example.com/demo-field/",
								"description": "External demo"
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

	routes, _, err := buildPackageRoutingTable(pkg, "")
	assert.NoError(t, err)
	assert.Empty(t, routes, "external demo URL should not produce a route in workspace mode")
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

			routes, _, err := BuildDemoRoutingTable([]byte(manifestJSON), "", "")
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
			name:        "exact route match returns empty",
			requestPath: "/elements/avatar/demo/sizes/",
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
			name:        "workspace with empty PackagePath falls back to watchDir",
			requestPath: "/elements/avatar/demo/sizes/perlman.jpg",
			watchDir:    "/test",
			isWorkspace: true,
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
