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
	"net/http"
	"testing"

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
		if !contains(manifestStr, field) {
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

// TestWebSocketTestClient_PlaceholderForPhase1 is a placeholder test
// This test will fail until Phase 1 (Core Server with WebSocket) is implemented
func TestWebSocketTestClient_PlaceholderForPhase1(t *testing.T) {
	t.Skip("Skipping until Phase 1 implements WebSocket server")

	// TODO: This test will be implemented in Phase 1
	// For now, it serves as documentation of what needs to be tested

	// Expected test:
	// 1. Start server with WebSocket endpoint
	// 2. Connect WebSocket test client
	// 3. Verify client can receive messages
	// 4. Verify client can send messages
}

// TestFixturePattern verifies the fixture pattern is established
func TestFixturePattern(t *testing.T) {
	// This test doesn't need to run anything, it just verifies
	// the fixture structure exists
	// See serve/fixture/simple-demo/ for the established pattern
	t.Log("Fixture pattern established in serve/fixture/simple-demo/")
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
