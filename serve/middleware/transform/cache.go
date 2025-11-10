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
	"container/list"
	"sync"
	"time"
)

// lruEntry wraps a cache key for the LRU list
type lruEntry struct {
	key CacheKey
}

// Cache provides thread-safe caching of transformed files with LRU eviction
type Cache struct {
	mu sync.RWMutex

	// Cache storage: key -> entry
	entries map[CacheKey]*CacheEntry

	// LRU tracking: most recently used at front
	lru *list.List
	// Map cache key to LRU element for O(1) removal
	lruMap map[CacheKey]*list.Element

	// Dependency graph: file path -> files that import it
	dependents map[string][]string

	// Metrics
	hits   int64
	misses int64

	// Configuration
	maxSize int64 // Max cache size in bytes
	curSize int64 // Current cache size in bytes
}

// NewCache creates a new transform cache with the given max size
func NewCache(maxSizeBytes int64) *Cache {
	return &Cache{
		entries:    make(map[CacheKey]*CacheEntry),
		lru:        list.New(),
		lruMap:     make(map[CacheKey]*list.Element),
		dependents: make(map[string][]string),
		maxSize:    maxSizeBytes,
	}
}

// Get retrieves a cached entry if it exists
func (c *Cache) Get(key CacheKey) (*CacheEntry, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.entries[key]
	if !found {
		c.misses++
		return nil, false
	}

	// Update access time and move to front of LRU
	entry.AccessTime = time.Now()
	if elem, ok := c.lruMap[key]; ok {
		c.lru.MoveToFront(elem)
	}

	c.hits++
	return entry, true
}

// Set adds or updates a cache entry
func (c *Cache) Set(key CacheKey, code []byte, dependencies []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Calculate entry size
	entrySize := int64(len(code))

	// Check if entry already exists
	if existing, found := c.entries[key]; found {
		// Update size accounting
		c.curSize -= existing.Size
		c.curSize += entrySize

		// Update entry
		existing.Code = code
		existing.Dependencies = dependencies
		existing.Size = entrySize
		existing.AccessTime = time.Now()

		// Move to front of LRU
		if elem, ok := c.lruMap[key]; ok {
			c.lru.MoveToFront(elem)
		}
	} else {
		// Add new entry
		entry := &CacheEntry{
			Code:         code,
			Dependencies: dependencies,
			Size:         entrySize,
			AccessTime:   time.Now(),
		}

		c.entries[key] = entry
		c.curSize += entrySize

		// Add to LRU
		elem := c.lru.PushFront(lruEntry{key: key})
		c.lruMap[key] = elem
	}

	// Update dependency graph (file -> dependents)
	for _, dep := range dependencies {
		c.addDependent(dep, key.Path)
	}

	// Evict if over size limit
	c.evictIfNeeded()
}

// addDependent adds a dependent relationship (dep is imported by dependent)
func (c *Cache) addDependent(dep, dependent string) {
	// Find existing dependents for this dependency
	deps := c.dependents[dep]

	// Check if already present
	for _, d := range deps {
		if d == dependent {
			return
		}
	}

	// Add to list
	c.dependents[dep] = append(deps, dependent)
}

// evictIfNeeded evicts LRU entries until under size limit
func (c *Cache) evictIfNeeded() {
	for c.curSize > c.maxSize && c.lru.Len() > 0 {
		// Get least recently used (back of list)
		elem := c.lru.Back()
		if elem == nil {
			break
		}

		entry := elem.Value.(lruEntry)
		c.evict(entry.key)
	}
}

// evict removes an entry from the cache
func (c *Cache) evict(key CacheKey) {
	entry, found := c.entries[key]
	if !found {
		return
	}

	// Remove from entries
	delete(c.entries, key)
	c.curSize -= entry.Size

	// Remove from LRU
	if elem, ok := c.lruMap[key]; ok {
		c.lru.Remove(elem)
		delete(c.lruMap, key)
	}

	// Remove from dependency graph
	c.removeDependents(key.Path)
}

// removeDependents removes all dependent relationships for a file
func (c *Cache) removeDependents(path string) {
	delete(c.dependents, path)

	// Also remove this path from other files' dependent lists
	for dep, dependentList := range c.dependents {
		filtered := make([]string, 0, len(dependentList))
		for _, d := range dependentList {
			if d != path {
				filtered = append(filtered, d)
			}
		}
		if len(filtered) > 0 {
			c.dependents[dep] = filtered
		} else {
			delete(c.dependents, dep)
		}
	}
}

// Invalidate removes a file and all its transitive dependents from cache
// Returns list of invalidated paths
func (c *Cache) Invalidate(path string) []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	invalidated := make([]string, 0)
	visited := make(map[string]bool)

	c.invalidateRecursive(path, visited, &invalidated)

	return invalidated
}

// invalidateRecursive recursively invalidates a file and its dependents
func (c *Cache) invalidateRecursive(path string, visited map[string]bool, invalidated *[]string) {
	if visited[path] {
		return
	}
	visited[path] = true

	// Collect dependents BEFORE evicting (evict removes from dependents map)
	var dependentsList []string
	if deps, ok := c.dependents[path]; ok {
		dependentsList = make([]string, len(deps))
		copy(dependentsList, deps)
	}

	// Find all cache entries for this path (may have multiple with different mod times)
	for key := range c.entries {
		if key.Path == path {
			c.evict(key)
			*invalidated = append(*invalidated, path)
			break // Only one entry per path should exist at a time
		}
	}

	// Recursively invalidate dependents (using the copy we made before eviction)
	for _, dependent := range dependentsList {
		c.invalidateRecursive(dependent, visited, invalidated)
	}
}

// Stats returns cache statistics
func (c *Cache) Stats() CacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return CacheStats{
		Hits:      c.hits,
		Misses:    c.misses,
		Entries:   len(c.entries),
		SizeBytes: c.curSize,
		MaxSize:   c.maxSize,
		HitRate:   c.hitRate(),
	}
}

// hitRate calculates the cache hit rate as a percentage
func (c *Cache) hitRate() float64 {
	total := c.hits + c.misses
	if total == 0 {
		return 0
	}
	return float64(c.hits) / float64(total) * 100
}

// Clear removes all entries from the cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[CacheKey]*CacheEntry)
	c.lru.Init()
	c.lruMap = make(map[CacheKey]*list.Element)
	c.dependents = make(map[string][]string)
	c.curSize = 0
}

// CacheStats represents cache metrics
type CacheStats struct {
	Hits      int64
	Misses    int64
	Entries   int
	SizeBytes int64
	MaxSize   int64
	HitRate   float64
}
