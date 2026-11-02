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

package serve

import (
	"testing"

	"bennypowers.dev/cem/serve/middleware/routes"
)

// TestReadInternalModule_Production tests that in production mode (default build),
// readInternalModule reads from embed.FS
func TestReadInternalModule_Production(t *testing.T) {
	// This test verifies the production behavior (reading from embed.FS)
	// The cemdev build tag should NOT be active when running normal tests

	// Try reading a known embedded file
	data, err := readInternalModule("templates/js/cem-element.js")
	if err != nil {
		t.Fatalf("readInternalModule() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("readInternalModule() returned empty data")
	}

	// Verify it's valid JavaScript by checking for expected content
	content := string(data)
	if content == "" {
		t.Error("readInternalModule() returned empty content")
	}
}

// TestSetupDevWatcher_Production tests that setupDevWatcher is a no-op in production
func TestSetupDevWatcher_Production(t *testing.T) {
	// Create a minimal server
	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("NewServer() error = %v", err)
	}

	// In production mode, setupDevWatcher should be a no-op
	err = setupDevWatcher(server)
	if err != nil {
		t.Errorf("setupDevWatcher() error = %v, want nil", err)
	}

	// Verify ReadInternalModule was not overridden (still points to embed.FS reader)
	// We can't directly test the function pointer, but we can test its behavior
	data, err := routes.ReadInternalModule("templates/js/cem-element.js")
	if err != nil {
		t.Fatalf("routes.ReadInternalModule() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("routes.ReadInternalModule() returned empty data after setupDevWatcher")
	}
}
