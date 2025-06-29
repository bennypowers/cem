package generate

import (
	"sync"
)

// CssParseCache caches parsed CSS files by absolute path.
type CssParseCache struct {
	mu    sync.RWMutex
	cache map[string]CssPropsMap
}

func NewCssParseCache() *CssParseCache {
	return &CssParseCache{
		cache: make(map[string]CssPropsMap),
	}
}

// Get checks the cache for a parsed result.
func (c *CssParseCache) Get(path string) (CssPropsMap, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cache[path]
	return val, ok
}

// Set stores a parsed result.
func (c *CssParseCache) Set(path string, props CssPropsMap) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[path] = props
}
