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
package lsp

import (
	"bennypowers.dev/cem/internal/platform"
)

// NewTestRegistry creates a registry suitable for testing.
// It uses a mock file watcher that can be controlled for deterministic testing.
func NewTestRegistry() *Registry {
	mockWatcher := platform.NewMockFileWatcher()
	return NewRegistry(mockWatcher)
}

// NewTestRegistryWithMockWatcher creates a registry for testing and returns both
// the registry and the mock watcher for test control.
func NewTestRegistryWithMockWatcher() (*Registry, *platform.MockFileWatcher) {
	mockWatcher := platform.NewMockFileWatcher()
	registry := NewRegistry(mockWatcher)
	return registry, mockWatcher
}
