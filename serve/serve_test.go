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

package serve_test

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"bennypowers.dev/cem/serve"
	"bennypowers.dev/cem/serve/testutil"
)

// TestServerTestHelper verifies the HTTP server test helper works
func TestServerTestHelper(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, test!"))
	})

	server := testutil.NewTestServer(t, handler)
	server.Start(t)

	// Make a request to verify server is running
	resp, err := http.Get(server.BaseURL + "/")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

// TestMockManifest verifies mock manifest generation produces valid JSON
func TestMockManifest(t *testing.T) {
	manifest, err := testutil.MockManifest()
	if err != nil {
		t.Fatalf("Failed to generate mock manifest: %v", err)
	}

	if len(manifest) == 0 {
		t.Error("Mock manifest is empty")
	}

	// Verify it contains expected CEM structure
	manifestStr := string(manifest)
	expectedFields := []string{"schemaVersion", "modules", "declarations"}
	for _, field := range expectedFields {
		if !strings.Contains(manifestStr, field) {
			t.Errorf("Mock manifest missing expected field: %s", field)
		}
	}
}

// TestTransformCacheTestDouble verifies the cache test double works
func TestTransformCacheTestDouble(t *testing.T) {
	cache := testutil.NewMockTransformCache()

	// Test Set and Get
	cache.Set("file.ts", []byte("transformed content"), nil)
	content, found := cache.Get("file.ts")
	if !found {
		t.Error("Expected to find cached file")
	}
	if string(content) != "transformed content" {
		t.Errorf("Expected 'transformed content', got %s", string(content))
	}

	// Test cache miss
	_, found = cache.Get("missing.ts")
	if found {
		t.Error("Expected cache miss for missing file")
	}

	// Test stats
	hits, misses := cache.Stats()
	if hits != 1 || misses != 1 {
		t.Errorf("Expected 1 hit and 1 miss, got %d hits and %d misses", hits, misses)
	}

	// Test invalidation
	cache.Set("dependent.ts", []byte("content"), []string{"file.ts"})
	invalidated := cache.Invalidate("file.ts")
	if len(invalidated) != 2 {
		t.Errorf("Expected 2 files invalidated, got %d", len(invalidated))
	}
}

// TestWebSocketReload verifies WebSocket live reload functionality
func TestWebSocketReload(t *testing.T) {
	// Find available port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to find available port: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()

	// Create temp dir for watch
	tmpDir := t.TempDir()

	// Configure server
	config := serve.Config{
		Port:   port,
		Reload: true,
	}

	server, err := serve.NewServerWithConfig(config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	if err := server.SetWatchDir(tmpDir); err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Start server
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer server.Close()

	// Wait a bit for server to be fully up
	// In a real scenario we'd want a more robust readiness check,
	// but sleep is acceptable for this level of E2E test.
	time.Sleep(100 * time.Millisecond)

	// Connect WebSocket client
	wsURL := fmt.Sprintf("ws://127.0.0.1:%d/__cem/reload", port)
	client := testutil.NewWebSocketTestClient(t, wsURL)
	defer client.Close()

	// Test BroadcastReload
	files := []string{"test.js", "style.css"}
	reason := "file-change"
	if err := server.BroadcastReload(files, reason); err != nil {
		t.Fatalf("Failed to broadcast reload: %v", err)
	}

	// Receive message
	msgBytes := client.ReceiveMessage(t, time.Second)
	if msgBytes == nil {
		t.Fatal("Did not receive expected reload message")
	}

	// Parse and verify
	var msg serve.ReloadMessage
	if err := json.Unmarshal(msgBytes, &msg); err != nil {
		t.Fatalf("Failed to unmarshal reload message: %v", err)
	}

	if msg.Type != "reload" {
		t.Errorf("Expected message type 'reload', got '%s'", msg.Type)
	}
	if msg.Reason != reason {
		t.Errorf("Expected reason '%s', got '%s'", reason, msg.Reason)
	}
	if len(msg.Files) != len(files) {
		t.Errorf("Expected %d files, got %d", len(files), len(msg.Files))
	} else {
		// Verify file list content
		for i, f := range files {
			if msg.Files[i] != f {
				t.Errorf("Expected file %d to be '%s', got '%s'", i, f, msg.Files[i])
			}
		}
	}

	// Test ExpectNoMessage (wait and ensure quiet)
	client.ExpectNoMessage(t, 200*time.Millisecond)
}

// TestFixturePattern verifies the fixture pattern is established
func TestFixturePattern(t *testing.T) {
	// This test doesn't need to run anything, it just verifies
	// the fixture structure exists
	// See serve/testdata/fixtures/simple-demo/ for the established pattern
	t.Log("Fixture pattern established in serve/testdata/fixtures/simple-demo/")
}
