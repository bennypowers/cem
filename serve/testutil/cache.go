//go:build e2e

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

package testutil

import (
	"sync"
)

// MockTransformCache is a test double for the transform cache
type MockTransformCache struct {
	mu      sync.RWMutex
	entries map[string][]byte
	deps    map[string][]string
	hits    int
	misses  int
}

// NewMockTransformCache creates a new mock transform cache
func NewMockTransformCache() *MockTransformCache {
	return &MockTransformCache{
		entries: make(map[string][]byte),
		deps:    make(map[string][]string),
	}
}

// Get retrieves a transformed file from the cache
func (c *MockTransformCache) Get(path string) ([]byte, bool) {
	// Read the entry under read lock
	c.mu.RLock()
	content, found := c.entries[path]
	c.mu.RUnlock()

	// Update counters under write lock
	c.mu.Lock()
	if found {
		c.hits++
	} else {
		c.misses++
	}
	c.mu.Unlock()

	return content, found
}

// Set stores a transformed file in the cache with its dependencies
func (c *MockTransformCache) Set(path string, content []byte, dependencies []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[path] = content
	c.deps[path] = dependencies
}

// Invalidate removes a file and its dependents from the cache
func (c *MockTransformCache) Invalidate(path string) []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	invalidated := []string{path}
	delete(c.entries, path)

	// Find all files that depend on this path
	for file, deps := range c.deps {
		for _, dep := range deps {
			if dep == path {
				delete(c.entries, file)
				invalidated = append(invalidated, file)
				break
			}
		}
	}

	return invalidated
}

// Stats returns cache hit/miss statistics
func (c *MockTransformCache) Stats() (hits, misses int) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.hits, c.misses
}

// Size returns the number of entries in the cache
func (c *MockTransformCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.entries)
}

// Clear removes all entries from the cache
func (c *MockTransformCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[string][]byte)
	c.deps = make(map[string][]string)
	c.hits = 0
	c.misses = 0
}
