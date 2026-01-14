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
package workspace_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestURLWorkspaceContext_Init(t *testing.T) {
	// Create a mock server with package.json and manifest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/package.json":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":           "@test/components",
				"version":        "1.0.0",
				"customElements": "./custom-elements.json",
			})
		case "/custom-elements.json":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"schemaVersion": "1.0.0",
				"modules": []map[string]any{
					{
						"path": "./src/my-button.ts",
						"declarations": []map[string]any{
							{
								"kind":    "class",
								"name":    "MyButton",
								"tagName": "my-button",
							},
						},
					},
				},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	err := ctx.Init()
	require.NoError(t, err, "Init should succeed")

	// Verify package.json is loaded
	pkgJSON, err := ctx.PackageJSON()
	require.NoError(t, err)
	assert.Equal(t, "@test/components", pkgJSON.Name)
	assert.Equal(t, "1.0.0", pkgJSON.Version)
	assert.Equal(t, "./custom-elements.json", pkgJSON.CustomElements)

	// Verify manifest is loaded
	manifest, err := ctx.Manifest()
	require.NoError(t, err)
	assert.NotNil(t, manifest)
	assert.Len(t, manifest.Modules, 1)
}

func TestURLWorkspaceContext_Root(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"name":           "@test/pkg",
			"customElements": "./custom-elements.json",
		})
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	// Root should return the URL
	assert.Equal(t, server.URL+"/", ctx.Root())
}

func TestURLWorkspaceContext_FollowsRedirects(t *testing.T) {
	// Simulate CDN redirect from unversioned to versioned URL
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Redirect unversioned to versioned
		if r.URL.Path == "/" || r.URL.Path == "/package.json" {
			if r.URL.Query().Get("redirected") != "true" {
				http.Redirect(w, r, r.URL.Path+"?redirected=true", http.StatusMovedPermanently)
				return
			}
		}

		switch r.URL.Path {
		case "/package.json":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":           "@test/pkg",
				"version":        "2.0.0",
				"customElements": "./custom-elements.json",
			})
		case "/custom-elements.json":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"schemaVersion": "1.0.0",
				"modules":       []any{},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	err := ctx.Init()
	require.NoError(t, err, "Init should succeed even with redirects")

	pkgJSON, err := ctx.PackageJSON()
	require.NoError(t, err)
	assert.Equal(t, "2.0.0", pkgJSON.Version)
}

func TestURLWorkspaceContext_MissingPackageJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	err := ctx.Init()
	assert.Error(t, err, "Init should fail when package.json is missing")
}

func TestURLWorkspaceContext_MissingCustomElements(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/package.json" {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":    "@test/pkg",
				"version": "1.0.0",
				// No customElements field
			})
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	err := ctx.Init()
	assert.Error(t, err, "Init should fail when customElements is not specified")
}

func TestURLWorkspaceContext_ManifestNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/package.json" {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":           "@test/pkg",
				"version":        "1.0.0",
				"customElements": "./custom-elements.json",
			})
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	err := ctx.Init()
	assert.Error(t, err, "Init should fail when manifest file is not found")
}

func TestURLWorkspaceContext_CustomElementsManifestPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/package.json":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":           "@test/pkg",
				"customElements": "./dist/custom-elements.json",
			})
		case "/dist/custom-elements.json":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"schemaVersion": "1.0.0",
				"modules":       []any{},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	err := ctx.Init()
	require.NoError(t, err)

	// Should return the path to the manifest
	manifestPath := ctx.CustomElementsManifestPath()
	assert.Contains(t, manifestPath, "custom-elements.json")
}

func TestURLWorkspaceContext_Caching(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Cache-Control", "max-age=3600")
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/package.json":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":           "@test/pkg",
				"customElements": "./custom-elements.json",
			})
		case "/custom-elements.json":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"schemaVersion": "1.0.0",
				"modules":       []any{},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	cacheDir := t.TempDir()

	// First context
	ctx1 := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)
	err := ctx1.Init()
	require.NoError(t, err)
	initialRequests := requestCount

	// Second context with same cache dir should use cache
	ctx2 := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)
	err = ctx2.Init()
	require.NoError(t, err)

	assert.Equal(t, initialRequests, requestCount, "Second Init should use cache")
}

func TestURLWorkspaceContext_Cleanup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/package.json":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":           "@test/pkg",
				"customElements": "./custom-elements.json",
			})
		case "/custom-elements.json":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"schemaVersion": "1.0.0",
				"modules":       []any{},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	ctx := workspace.NewURLWorkspaceContext(server.URL+"/", cacheDir)

	err := ctx.Init()
	require.NoError(t, err)

	// Cleanup should not error
	err = ctx.Cleanup()
	assert.NoError(t, err)
}
