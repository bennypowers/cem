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
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/types"
)

// newTestFS creates an in-memory filesystem with manifest-regen fixture data
func newTestFS(t *testing.T) platform.FileSystem {
	t.Helper()

	mfs := platform.NewMapFileSystem(nil)

	// Load fixture data from testdata/manifest-regen
	mfs.AddFile("/test-package/package.json", `{
  "name": "test-package",
  "customElements": "custom-elements.json"
}`, 0644)

	mfs.AddFile("/test-package/.config/cem.yaml", `generate:
  files:
    - "src/*.ts"
`, 0644)

	mfs.AddFile("/test-package/src/test.ts", `export class TestElement extends HTMLElement {}`, 0644)

	return mfs
}

// TestReloadMessageFormat verifies reload message structure matches expected format
func TestReloadMessageFormat(t *testing.T) {
	// Read expected golden file
	goldenPath := filepath.Join("testdata", "websocket-reload", "expected-reload-message.json")
	expectedBytes, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	// Parse expected message
	var expected map[string]any
	if err := json.Unmarshal(expectedBytes, &expected); err != nil {
		t.Fatalf("Failed to parse expected message: %v", err)
	}

	// Create server and generate message
	server, err := serve.NewServer(8000)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Create reload message
	msgBytes, err := server.CreateReloadMessage([]string{"test.ts", "foo.ts"}, "file-change")
	if err != nil {
		t.Fatalf("Failed to create reload message: %v", err)
	}

	// Parse actual message
	var actual map[string]any
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
	actualFiles, ok := actual["files"].([]any)
	if !ok {
		t.Fatal("Expected files to be an array")
	}

	expectedFiles, ok := expected["files"].([]any)
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
	defer func() { _ = server.Close() }()

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
	defer func() { _ = server.Close() }()

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
	defer func() { _ = server.Close() }()

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
	mfs := newTestFS(t)
	server, err := serve.NewServerWithConfig(serve.Config{
		Port:   8005,
		Reload: true,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	testDir := "/test-package"
	err = server.SetWatchDir(testDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	watchDir := server.WatchDir()
	if watchDir != testDir {
		t.Errorf("Expected watch dir %s, got %s", testDir, watchDir)
	}
}

// TestManifestAccessors verifies manifest get/set interface
func TestManifestAccessors(t *testing.T) {
	server, err := serve.NewServer(8006)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	testManifest := []byte(`{"schemaVersion":"1.0.0","modules":[]}`)
	err = server.SetManifest(testManifest)
	if err != nil {
		t.Fatalf("Failed to set manifest: %v", err)
	}

	manifest, err := server.Manifest()
	if err != nil {
		t.Fatalf("Failed to get manifest: %v", err)
	}

	if string(manifest) != string(testManifest) {
		t.Error("Expected manifest to match what was set")
	}
}

// TestManifestDefensiveCopy verifies that manifest uses defensive copying
func TestManifestDefensiveCopy(t *testing.T) {
	server, err := serve.NewServer(8014)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Create a manifest and modify it after setting
	original := []byte(`{"schemaVersion":"1.0.0"}`)
	err = server.SetManifest(original)
	if err != nil {
		t.Fatalf("Failed to set manifest: %v", err)
	}

	// Modify the original slice
	original[0] = 'X'

	// Get manifest and verify it wasn't affected
	retrieved, err := server.Manifest()
	if err != nil {
		t.Fatalf("Failed to get manifest: %v", err)
	}

	if retrieved[0] != '{' {
		t.Error("SetManifest didn't make defensive copy - caller mutation affected internal state")
	}

	// Modify the retrieved slice
	retrieved[0] = 'Y'

	// Get manifest again and verify it wasn't affected
	retrieved2, err := server.Manifest()
	if err != nil {
		t.Fatalf("Failed to get manifest second time: %v", err)
	}

	if retrieved2[0] != '{' {
		t.Error("Manifest didn't make defensive copy - caller mutation affected internal state")
	}
}

// TestDebounceDuration verifies 50ms debounce duration
func TestDebounceDuration(t *testing.T) {
	server, err := serve.NewServer(8007)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	duration := server.DebounceDuration()
	expected := "50ms"

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
	defer func() { _ = server.Close() }()

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
	defer func() { _ = server.Close() }()

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
	defer func() { _ = server.Close() }()

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
	defer func() { _ = server.Close() }()

	// Should accept files and reason
	err = server.BroadcastReload([]string{"test.ts"}, "file-change")
	if err != nil {
		t.Fatalf("BroadcastReload failed: %v", err)
	}
}

// TestManifestRegenerationTrigger verifies regeneration can be triggered
func TestManifestRegenerationTrigger(t *testing.T) {
	mfs := newTestFS(t)
	server, err := serve.NewServerWithConfig(serve.Config{
		Port:   8012,
		Reload: true,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	err = server.SetWatchDir("/test-package")
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	// Should be able to trigger regeneration
	_, err = server.RegenerateManifest()
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
	defer func() { _ = server.Close() }()

	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	// Regenerate manifest
	_, err = server.RegenerateManifest()
	if err != nil {
		t.Fatalf("Failed to regenerate manifest: %v", err)
	}

	// Get manifest
	manifest, err := server.Manifest()
	if err != nil {
		t.Fatalf("Failed to get manifest: %v", err)
	}

	// Verify manifest is not empty
	if len(manifest) == 0 {
		t.Error("Expected manifest to be generated, got empty")
	}

	// Verify manifest is valid JSON
	var manifestObj map[string]any
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

// TestPortBindingError verifies that port binding errors are detected immediately
func TestPortBindingError(t *testing.T) {
	// Create and start first server
	server1, err := serve.NewServer(8888)
	if err != nil {
		t.Fatalf("Failed to create first server: %v", err)
	}

	err = server1.Start()
	if err != nil {
		t.Fatalf("Failed to start first server: %v", err)
	}
	defer func() { _ = server1.Close() }()

	// Try to start second server on same port - should fail immediately
	server2, err := serve.NewServer(8888)
	if err != nil {
		t.Fatalf("Failed to create second server: %v", err)
	}

	err = server2.Start()
	if err == nil {
		_ = server2.Close()
		t.Fatal("Expected port binding error, but Start() succeeded")
	}

	// Verify error message mentions port binding
	if !strings.Contains(err.Error(), "bind") && !strings.Contains(err.Error(), "address already in use") {
		t.Errorf("Expected error about port binding, got: %v", err)
	}

	// Verify second server is not running
	if server2.IsRunning() {
		t.Error("Expected second server to not be running after bind failure")
	}
}

// TestImportResolution verifies import map resolution with prefix matching
func TestImportResolution(t *testing.T) {
	// Create in-memory filesystem with test package
	mfs := platform.NewMapFileSystem(nil)
	packageJSON := `{
  "name": "@test/elements",
  "exports": {
    "./*": "./src/*"
  }
}`
	mfs.AddFile("/test-package/package.json", packageJSON, 0644)

	server, err := serve.NewServerWithConfig(serve.Config{
		Port:   8015,
		Reload: true,
		FS:     mfs,
		ImportMap: types.ImportMapConfig{
			Generate: true,
		},
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	err = server.SetWatchDir("/test-package")
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	// The server should have generated an import map with the prefix entry
	// @test/elements/ -> /src/
	// This would allow @test/elements/foo/bar.js to resolve to /src/foo/bar.js

	importMap := server.ImportMap()
	if importMap == nil {
		t.Fatal("Expected import map to be generated, got nil")
	}

	// Type assert to concrete import map type to inspect entries
	im, ok := importMap.(*importmappkg.ImportMap)
	if !ok {
		t.Fatalf("Expected *importmap.ImportMap, got %T", importMap)
	}

	// Verify the prefix entry exists
	expectedPrefix := "@test/elements/"
	expectedTarget := "/src/"

	actualTarget, exists := im.Imports[expectedPrefix]
	if !exists {
		t.Errorf("Expected import map to contain prefix entry %q, but it was missing. Imports: %v", expectedPrefix, im.Imports)
	}

	if actualTarget != expectedTarget {
		t.Errorf("Expected prefix entry %q to map to %q, got %q", expectedPrefix, expectedTarget, actualTarget)
	}
}

// TestRegenerateManifestIncremental_NoDoubleLockPanic verifies that calling
// RegenerateManifestIncremental when generateSession is nil doesn't cause
// a double-unlock panic
func TestRegenerateManifestIncremental_NoDoubleLockPanic(t *testing.T) {
	// Create in-memory filesystem with fixture data
	mfs := newTestFS(t)

	server, err := serve.NewServerWithConfig(serve.Config{
		Port:   0,
		Reload: true,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	err = server.SetWatchDir("/test-package")
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	// Call RegenerateManifestIncremental WITHOUT calling RegenerateManifest first
	// This means generateSession will be nil, triggering the fallback path
	// With the bug, this will cause a double-unlock panic
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Double-unlock panic detected: %v", r)
		}
	}()

	srcFile := "/test-package/src/test.ts"
	_, err = server.RegenerateManifestIncremental([]string{srcFile})
	// Error is expected (since we're calling incremental without a session),
	// but no panic should occur
	if err != nil {
		// This is acceptable - we're just checking for panic
		t.Logf("RegenerateManifestIncremental returned error (expected): %v", err)
	}
}

// TestStaticFileMIMETypes_JSModules verifies that JavaScript module files are served with correct MIME types
func TestStaticFileMIMETypes_JSModules(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		fileContent  string
		expectedMIME string
	}{
		{
			name:         ".mjs file",
			path:         "/test.mjs",
			fileContent:  "export const foo = 'bar';",
			expectedMIME: "application/javascript; charset=utf-8",
		},
		{
			name:         ".cjs file",
			path:         "/test.cjs",
			fileContent:  "module.exports = { foo: 'bar' };",
			expectedMIME: "application/javascript; charset=utf-8",
		},
		{
			name:         ".json file",
			path:         "/test.json",
			fileContent:  `{"foo": "bar"}`,
			expectedMIME: "application/json; charset=utf-8",
		},
		{
			name:         ".svg file",
			path:         "/test.svg",
			fileContent:  `<svg xmlns="http://www.w3.org/2000/svg"></svg>`,
			expectedMIME: "image/svg+xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create in-memory filesystem with test file
			mfs := platform.NewMapFileSystem(nil)
			mfs.AddFile("/test-package"+tt.path, tt.fileContent, 0644)
			mfs.AddFile("/test-package/package.json", `{"name": "test"}`, 0644)

			server, err := serve.NewServerWithConfig(serve.Config{
				Port: 0,
				FS:   mfs,
			})
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}
			defer func() { _ = server.Close() }()

			err = server.SetWatchDir("/test-package")
			if err != nil {
				t.Fatalf("Failed to set watch directory: %v", err)
			}

			// Make request for static file (no need to start the server, just use the handler)
			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()

			server.Handler().ServeHTTP(rec, req)

			// Verify status code
			if rec.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", rec.Code)
			}

			// Verify MIME type
			contentType := rec.Header().Get("Content-Type")
			if contentType != tt.expectedMIME {
				t.Errorf("Expected Content-Type %q, got %q", tt.expectedMIME, contentType)
			}
		})
	}
}

// TestStaticFileMIMETypes_NodeModules verifies MIME types for node_modules files (regression test)
func TestStaticFileMIMETypes_NodeModules(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		fileContent  string
		expectedMIME string
	}{
		{
			name:         "tslib .mjs in node_modules",
			path:         "/node_modules/tslib/tslib.es6.mjs",
			fileContent:  "export function __extends() {}",
			expectedMIME: "application/javascript; charset=utf-8",
		},
		{
			name:         "package .mjs in node_modules",
			path:         "/node_modules/some-package/index.mjs",
			fileContent:  "export default {}",
			expectedMIME: "application/javascript; charset=utf-8",
		},
		{
			name:         "package .cjs in node_modules",
			path:         "/node_modules/some-package/index.cjs",
			fileContent:  "module.exports = {}",
			expectedMIME: "application/javascript; charset=utf-8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create in-memory filesystem with test file
			mfs := platform.NewMapFileSystem(nil)
			mfs.AddFile("/test-package"+tt.path, tt.fileContent, 0644)
			mfs.AddFile("/test-package/package.json", `{"name": "test"}`, 0644)

			server, err := serve.NewServerWithConfig(serve.Config{
				Port: 0,
				FS:   mfs,
			})
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}
			defer func() { _ = server.Close() }()

			err = server.SetWatchDir("/test-package")
			if err != nil {
				t.Fatalf("Failed to set watch directory: %v", err)
			}

			// Make request for static file (no need to start the server, just use the handler)
			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()

			server.Handler().ServeHTTP(rec, req)

			// Verify status code
			if rec.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", rec.Code)
			}

			// Verify MIME type - this is the critical regression test
			contentType := rec.Header().Get("Content-Type")
			if contentType != tt.expectedMIME {
				t.Errorf("Expected Content-Type %q, got %q (regression: .mjs/.cjs should not be text/plain)", tt.expectedMIME, contentType)
			}
		})
	}
}
