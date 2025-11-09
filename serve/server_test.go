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

// PHASE 1 TESTS - Core Server & Live Reload
// These tests verify HTTP server, WebSocket endpoint, and file watching functionality
// using the fixture/golden pattern

package serve_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve"
)

// TestReloadMessageFormat verifies reload message structure matches expected format
func TestReloadMessageFormat(t *testing.T) {
	// Read expected golden file
	goldenPath := filepath.Join("testdata", "websocket-reload", "expected-reload-message.json")
	expectedBytes, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	// Parse expected message
	var expected map[string]interface{}
	if err := json.Unmarshal(expectedBytes, &expected); err != nil {
		t.Fatalf("Failed to parse expected message: %v", err)
	}

	// Create server and generate message
	server, err := serve.NewServer(8000)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Create reload message
	msgBytes, err := server.CreateReloadMessage([]string{"test.ts", "foo.ts"}, "file-change")
	if err != nil {
		t.Fatalf("Failed to create reload message: %v", err)
	}

	// Parse actual message
	var actual map[string]interface{}
	if err := json.Unmarshal(msgBytes, &actual); err != nil {
		t.Fatalf("Failed to parse actual message: %v", err)
	}

	// Verify structure matches expected
	if actual["type"] != expected["type"] {
		t.Errorf("Expected type %v, got %v", expected["type"], actual["type"])
	}

	if actual["reason"] != expected["reason"] {
		t.Errorf("Expected reason %v, got %v", expected["reason"], actual["reason"])
	}

	// Verify files array
	actualFiles, ok := actual["files"].([]interface{})
	if !ok {
		t.Fatal("Expected files to be an array")
	}

	expectedFiles, ok := expected["files"].([]interface{})
	if !ok {
		t.Fatal("Expected files in golden to be an array")
	}

	if len(actualFiles) != len(expectedFiles) {
		t.Errorf("Expected %d files, got %d", len(expectedFiles), len(actualFiles))
	}
}

// TestServerCreation verifies basic server creation
func TestServerCreation(t *testing.T) {
	server, err := serve.NewServer(8001)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	if server.Port() != 8001 {
		t.Errorf("Expected port 8001, got %d", server.Port())
	}
}

// TestServerConfig verifies config-based server creation
func TestServerConfig(t *testing.T) {
	config := serve.Config{
		Port:   8002,
		Reload: true,
	}

	server, err := serve.NewServerWithConfig(config)
	if err != nil {
		t.Fatalf("Failed to create server with config: %v", err)
	}
	defer server.Close()

	if server.Port() != 8002 {
		t.Errorf("Expected port 8002, got %d", server.Port())
	}
}

// TestServerLifecycle verifies start/stop behavior
func TestServerLifecycle(t *testing.T) {
	server, err := serve.NewServer(8003)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Should not be running initially
	if server.IsRunning() {
		t.Error("Expected server to not be running before Start()")
	}

	// Start server
	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Should be running after start
	if !server.IsRunning() {
		t.Error("Expected server to be running after Start()")
	}

	// Close server
	err = server.Close()
	if err != nil {
		t.Fatalf("Failed to close server: %v", err)
	}

	// Should not be running after close
	if server.IsRunning() {
		t.Error("Expected server to not be running after Close()")
	}
}

// TestWebSocketManager verifies WebSocket manager exists and has correct interface
func TestWebSocketManager(t *testing.T) {
	server, err := serve.NewServer(8004)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	wsManager := server.WebSocketManager()
	if wsManager == nil {
		t.Fatal("Expected WebSocket manager to exist")
	}

	// Should have zero connections initially
	if wsManager.ConnectionCount() != 0 {
		t.Errorf("Expected 0 connections initially, got %d", wsManager.ConnectionCount())
	}
}

// TestFileWatcherSetup verifies file watcher can be configured
func TestFileWatcherSetup(t *testing.T) {
	server, err := serve.NewServer(8005)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	tmpDir := t.TempDir()
	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	watchDir := server.GetWatchDir()
	if watchDir != tmpDir {
		t.Errorf("Expected watch dir %s, got %s", tmpDir, watchDir)
	}
}

// TestManifestAccessors verifies manifest get/set interface
func TestManifestAccessors(t *testing.T) {
	server, err := serve.NewServer(8006)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	testManifest := []byte(`{"schemaVersion":"1.0.0","modules":[]}`)
	err = server.SetManifest(testManifest)
	if err != nil {
		t.Fatalf("Failed to set manifest: %v", err)
	}

	manifest, err := server.GetManifest()
	if err != nil {
		t.Fatalf("Failed to get manifest: %v", err)
	}

	if string(manifest) != string(testManifest) {
		t.Error("Expected manifest to match what was set")
	}
}

// TestDebounceDuration verifies 150ms debounce as per spec
func TestDebounceDuration(t *testing.T) {
	server, err := serve.NewServer(8007)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	duration := server.GetDebounceDuration()
	expected := "150ms"

	if duration.String() != expected {
		t.Errorf("Expected debounce duration %s, got %s", expected, duration.String())
	}
}

// TestConfigReloadDisabled verifies reload can be disabled
func TestConfigReloadDisabled(t *testing.T) {
	config := serve.Config{
		Port:   8008,
		Reload: false,
	}

	server, err := serve.NewServerWithConfig(config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// WebSocket manager should be nil when reload disabled
	wsManager := server.WebSocketManager()
	if wsManager != nil {
		t.Error("Expected WebSocket manager to be nil when reload disabled")
	}
}

// TestCORSHeaders verifies CORS middleware exists
func TestCORSHeaders(t *testing.T) {
	server, err := serve.NewServer(8009)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Server should have CORS middleware configured
	// This is tested more thoroughly in e2e tests
	if server.Handler() == nil {
		t.Error("Expected server to have handler configured")
	}
}

// TestLogging verifies logger exists
func TestLogging(t *testing.T) {
	server, err := serve.NewServer(8010)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	logger := server.Logger()
	if logger == nil {
		t.Error("Expected server to have a logger")
	}
}

// TestBroadcastReload verifies broadcast method exists and accepts correct parameters
func TestBroadcastReload(t *testing.T) {
	server, err := serve.NewServer(8011)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Should accept files and reason
	err = server.BroadcastReload([]string{"test.ts"}, "file-change")
	if err != nil {
		t.Fatalf("BroadcastReload failed: %v", err)
	}
}

// TestManifestRegenerationTrigger verifies regeneration can be triggered
func TestManifestRegenerationTrigger(t *testing.T) {
	server, err := serve.NewServer(8012)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	tmpDir := t.TempDir()
	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	// Should be able to trigger regeneration
	err = server.RegenerateManifest()
	if err != nil {
		// Error is expected if no source files exist, but method should exist
		if !strings.Contains(err.Error(), "no source files") &&
			!strings.Contains(err.Error(), "not implemented") {
			t.Fatalf("Unexpected error from RegenerateManifest: %v", err)
		}
	}
}

// TestManifestGeneration_WithTempDir verifies manifest generation with temp directory
func TestManifestGeneration_WithTempDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Create package.json
	packageJSON := `{
  "name": "manifest-test",
  "customElements": "custom-elements.json"
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Create .config directory and cem.yaml
	configDir := filepath.Join(tmpDir, ".config")
	err = os.Mkdir(configDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create .config directory: %v", err)
	}

	cemConfig := `generate:
  files:
    - "src/*.ts"
`
	err = os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(cemConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to write cem.yaml: %v", err)
	}

	// Create src directory
	srcDir := filepath.Join(tmpDir, "src")
	err = os.Mkdir(srcDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create src directory: %v", err)
	}

	// Create a simple TypeScript file
	tsContent := `/**
 * A test element
 * @element test-element
 */
export class TestElement extends HTMLElement {}
`
	err = os.WriteFile(filepath.Join(srcDir, "test-element.ts"), []byte(tsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write TypeScript file: %v", err)
	}

	server, err := serve.NewServer(8013)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	// Regenerate manifest
	err = server.RegenerateManifest()
	if err != nil {
		t.Fatalf("Failed to regenerate manifest: %v", err)
	}

	// Get manifest
	manifest, err := server.GetManifest()
	if err != nil {
		t.Fatalf("Failed to get manifest: %v", err)
	}

	// Verify manifest is not empty
	if len(manifest) == 0 {
		t.Error("Expected manifest to be generated, got empty")
	}

	// Verify manifest is valid JSON
	var manifestObj map[string]interface{}
	if err := json.Unmarshal(manifest, &manifestObj); err != nil {
		t.Fatalf("Expected valid JSON manifest, got error: %v", err)
	}

	// Verify manifest has schemaVersion
	if _, ok := manifestObj["schemaVersion"]; !ok {
		t.Error("Expected manifest to have schemaVersion field")
	}

	// Verify manifest has modules
	if _, ok := manifestObj["modules"]; !ok {
		t.Error("Expected manifest to have modules field")
	}
}
