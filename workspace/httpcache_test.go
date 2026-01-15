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
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPCache_CacheControlMaxAge(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Cache-Control", "max-age=3600") // 1 hour
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "test-package"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	// First request should fetch from server
	content, err := cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")
	assert.Equal(t, 1, requestCount, "Should have made 1 request")

	// Second request should use cache (within max-age)
	content, err = cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")
	assert.Equal(t, 1, requestCount, "Should still have 1 request (cached)")
}

func TestHTTPCache_ExpiresHeader(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		// Set Expires to 1 hour from now
		expires := time.Now().Add(1 * time.Hour).UTC().Format(http.TimeFormat)
		w.Header().Set("Expires", expires)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "test-package"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	// First request should fetch from server
	_, err := cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Equal(t, 1, requestCount)

	// Second request should use cache
	_, err = cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Equal(t, 1, requestCount, "Should use cache with valid Expires")
}

func TestHTTPCache_ETagConditionalRequest(t *testing.T) {
	requestCount := 0
	etag := `"abc123"`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		// Check for conditional request
		if r.Header.Get("If-None-Match") == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.Header().Set("ETag", etag)
		w.Header().Set("Cache-Control", "max-age=0") // Always revalidate
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "test-package"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	// First request should fetch from server
	content, err := cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")
	assert.Equal(t, 1, requestCount)

	// Second request should revalidate and get 304
	content, err = cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")
	assert.Equal(t, 2, requestCount, "Should have made conditional request")
}

func TestHTTPCache_LastModifiedConditionalRequest(t *testing.T) {
	requestCount := 0
	lastModified := time.Now().Add(-24 * time.Hour).UTC().Format(http.TimeFormat)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		// Check for conditional request
		if r.Header.Get("If-Modified-Since") != "" {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.Header().Set("Last-Modified", lastModified)
		w.Header().Set("Cache-Control", "max-age=0") // Always revalidate
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "test-package"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	// First request should fetch from server
	content, err := cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")
	assert.Equal(t, 1, requestCount)

	// Second request should revalidate and get 304
	content, err = cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")
	assert.Equal(t, 2, requestCount, "Should have made conditional request")
}

func TestHTTPCache_FollowRedirects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/npm/@scope/pkg/" {
			// Redirect to versioned URL
			http.Redirect(w, r, "/npm/@scope/pkg@2.0.0/", http.StatusMovedPermanently)
			return
		}
		w.Header().Set("Cache-Control", "max-age=3600")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "@scope/pkg", "version": "2.0.0"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	// First request to unversioned URL should follow redirect
	content, err := cache.Fetch(server.URL + "/npm/@scope/pkg/")
	require.NoError(t, err)
	assert.Contains(t, string(content), "2.0.0")

	// Second request to same URL should still work
	content, err = cache.Fetch(server.URL + "/npm/@scope/pkg/")
	require.NoError(t, err)
	assert.Contains(t, string(content), "2.0.0")
}

func TestHTTPCache_NoCacheHeaders(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		// No cache headers - RFC 7234 says this is not cacheable
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "test-package"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	// First request
	content, err := cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")

	// Content should still be returned on second request
	content, err = cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-package")
}

func TestHTTPCache_ForceRefresh(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Cache-Control", "max-age=3600")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "test-package"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	// First request
	_, err := cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Equal(t, 1, requestCount)

	// Second request with force refresh
	_, err = cache.FetchWithRefresh(server.URL + "/package.json")
	require.NoError(t, err)
	assert.Equal(t, 2, requestCount, "Should bypass cache with force refresh")
}

func TestHTTPCache_CacheDirectory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=3600")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name": "test-package"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	_, err := cache.Fetch(server.URL + "/package.json")
	require.NoError(t, err)

	// Verify cache directory has entries (diskcache creates files)
	entries, err := os.ReadDir(cacheDir)
	require.NoError(t, err)
	assert.NotEmpty(t, entries, "Cache directory should have entries")
}

func TestHTTPCache_404Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	_, err := cache.Fetch(server.URL + "/nonexistent")
	assert.Error(t, err, "Should return error for 404")
}

func TestHTTPCache_FinalURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/npm/@scope/pkg/" {
			http.Redirect(w, r, "/npm/@scope/pkg@2.0.0/", http.StatusMovedPermanently)
			return
		}
		w.Header().Set("Cache-Control", "max-age=3600")
		_, _ = w.Write([]byte(`{"name": "@scope/pkg"}`))
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	cache := workspace.NewHTTPCache(cacheDir)

	_, err := cache.Fetch(server.URL + "/npm/@scope/pkg/")
	require.NoError(t, err)

	// Get the final URL after redirects
	finalURL := cache.FinalURL(server.URL + "/npm/@scope/pkg/")
	assert.Contains(t, finalURL, "@2.0.0", "Should return final URL after redirects")
}
