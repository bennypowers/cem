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

package workspace

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

// HTTPCache provides HTTP caching with support for redirects and RFC 7234 cache headers.
// It wraps gregjones/httpcache for RFC 7234 compliant caching.
type HTTPCache struct {
	client *http.Client
	mu     sync.RWMutex
	// Track final URLs after redirects for cache key lookup
	finalURLs map[string]string
}

// NewHTTPCache creates a new HTTP cache backed by the given cache directory.
// The cache respects RFC 7234 headers (Cache-Control, ETag, Last-Modified, etc.).
func NewHTTPCache(cacheDir string) *HTTPCache {
	// Create disk-backed cache
	cache := diskcache.New(cacheDir)

	// Create transport with caching
	transport := httpcache.NewTransport(cache)

	return &HTTPCache{
		client:    transport.Client(),
		finalURLs: make(map[string]string),
	}
}

// Fetch retrieves content from a URL, using cache when valid per RFC 7234.
func (c *HTTPCache) Fetch(url string) ([]byte, error) {
	return c.fetch(url, false)
}

// FetchWithRefresh retrieves content from a URL, bypassing the cache.
func (c *HTTPCache) FetchWithRefresh(url string) ([]byte, error) {
	return c.fetch(url, true)
}

// FinalURL returns the final URL after redirects for a given URL.
// Returns the original URL if no redirect was tracked.
func (c *HTTPCache) FinalURL(url string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if final, ok := c.finalURLs[url]; ok {
		return final
	}
	return url
}

func (c *HTTPCache) fetch(url string, forceRefresh bool) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Force refresh by setting Cache-Control header
	if forceRefresh {
		req.Header.Set("Cache-Control", "no-cache")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	// Track final URL after redirects
	finalURL := resp.Request.URL.String()
	c.trackFinalURL(url, finalURL)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (c *HTTPCache) trackFinalURL(originalURL, finalURL string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.finalURLs[originalURL] = finalURL
}

var ErrHTTPNotFound = errors.New("resource not found")
