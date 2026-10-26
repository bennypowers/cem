//go:build cemdev

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
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestReadInternalModule_Dev tests that in dev mode,
// readInternalModule reads from disk instead of embed.FS
func TestReadInternalModule_Dev(t *testing.T) {
	// Get the actual templates directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Failed to get caller information")
	}
	serveDir := filepath.Dir(filename)
	templatesDir := filepath.Join(serveDir, "middleware", "routes", "templates")

	// Verify the templates directory exists
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		t.Skipf("Templates directory does not exist: %s", templatesDir)
	}

	// Try reading a known file
	data, err := readInternalModule("templates/js/cem-element.js")
	if err != nil {
		t.Fatalf("readInternalModule() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("readInternalModule() returned empty data")
	}

	// Verify it's reading from disk by checking the actual file
	actualPath := filepath.Join(templatesDir, "js", "cem-element.js")
	diskData, err := os.ReadFile(actualPath)
	if err != nil {
		t.Fatalf("Failed to read actual file: %v", err)
	}

	// In dev mode, the data should match the disk version
	if string(data) != string(diskData) {
		t.Error("Dev mode should read from disk, but data doesn't match")
	}
}

// TestSetupDevWatcher_Dev tests that setupDevWatcher properly configures
// the file watcher in dev mode
func TestSetupDevWatcher_Dev(t *testing.T) {
	// Get the elements directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Failed to get caller information")
	}
	serveDir := filepath.Dir(filename)
	elementsDir := filepath.Join(serveDir, "middleware", "routes", "templates", "elements")

	// Skip if elements directory doesn't exist
	if _, err := os.Stat(elementsDir); os.IsNotExist(err) {
		t.Skipf("Elements directory does not exist: %s", elementsDir)
	}

	// Create a minimal server with WebSocket manager for testing
	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("NewServer() error = %v", err)
	}

	// Ensure WebSocket manager is set
	if server.wsManager == nil {
		server.wsManager = newWebSocketManager()
		server.wsManager.SetLogger(server.logger)
	}

	// Set up dev watcher
	err = setupDevWatcher(server)
	if err != nil {
		t.Fatalf("setupDevWatcher() error = %v", err)
	}

	// Verify that the file is read from disk
	data, err := readInternalModule("templates/js/cem-element.js")
	if err != nil {
		t.Fatalf("After setupDevWatcher, readInternalModule() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("After setupDevWatcher, readInternalModule() returned empty data")
	}

	// Verify it contains valid JavaScript content
	content := string(data)
	if !strings.Contains(content, "class") && !strings.Contains(content, "export") {
		t.Error("Content doesn't appear to be valid JavaScript")
	}
}
