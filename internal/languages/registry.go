package languages

import (
	"fmt"
	"sync"
)

var (
	registry   = map[string]Language{}
	registryMu sync.RWMutex
)

// Register adds a language to the global registry. Called by language packages
// in their init() functions.
func Register(l Language) {
	registryMu.Lock()
	defer registryMu.Unlock()
	if _, exists := registry[l.Name()]; exists {
		panic(fmt.Sprintf("language %q already registered", l.Name()))
	}
	registry[l.Name()] = l
}

// Get returns the registered language by name, or nil if not registered.
func Get(name string) Language {
	registryMu.RLock()
	defer registryMu.RUnlock()
	return registry[name]
}

// All returns a copy of all registered languages.
func All() map[string]Language {
	registryMu.RLock()
	defer registryMu.RUnlock()
	result := make(map[string]Language, len(registry))
	for k, v := range registry {
		result[k] = v
	}
	return result
}
