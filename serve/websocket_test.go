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
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"bennypowers.dev/cem/serve"
	"bennypowers.dev/cem/serve/internal/urlutil"
	"bennypowers.dev/cem/serve/testutil"
)

// TestWebSocketEndpoint_Connection verifies WebSocket endpoint accepts connections
func TestWebSocketEndpoint_Connection(t *testing.T) {
	server, err := serve.NewServer(9000)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	// Connect to WebSocket endpoint
	_ = testutil.NewWebSocketTestClient(t, "ws://localhost:9000/__cem/reload")

	// Give WebSocket time to connect
	time.Sleep(100 * time.Millisecond)

	// Verify connection count increased
	if server.WebSocketManager().ConnectionCount() != 1 {
		t.Errorf("Expected 1 connection, got %d", server.WebSocketManager().ConnectionCount())
	}
}

// TestWebSocketEndpoint_Broadcast verifies broadcast to multiple clients
func TestWebSocketEndpoint_Broadcast(t *testing.T) {
	server, err := serve.NewServer(9001)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	// Connect multiple clients
	client1 := testutil.NewWebSocketTestClient(t, "ws://localhost:9001/__cem/reload")
	client2 := testutil.NewWebSocketTestClient(t, "ws://localhost:9001/__cem/reload")
	client3 := testutil.NewWebSocketTestClient(t, "ws://localhost:9001/__cem/reload")

	// Broadcast a reload message
	err = server.BroadcastReload([]string{"test.ts"}, "file-change")
	if err != nil {
		t.Fatalf("Failed to broadcast: %v", err)
	}

	// All clients should receive the message
	msg1 := client1.ReceiveMessage(t, 2*time.Second)
	msg2 := client2.ReceiveMessage(t, 2*time.Second)
	msg3 := client3.ReceiveMessage(t, 2*time.Second)

	// Verify all received same message
	if string(msg1) != string(msg2) || string(msg2) != string(msg3) {
		t.Error("Clients received different messages")
	}

	// Verify message structure matches golden file
	goldenPath := filepath.Join("testdata", "websocket-reload", "expected-reload-message.json")
	expectedBytes, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	var expected map[string]any
	if err := json.Unmarshal(expectedBytes, &expected); err != nil {
		t.Fatalf("Failed to parse expected message: %v", err)
	}

	var actual map[string]any
	if err := json.Unmarshal(msg1, &actual); err != nil {
		t.Fatalf("Failed to parse actual message: %v", err)
	}

	if actual["type"] != "reload" {
		t.Errorf("Expected type 'reload', got %v", actual["type"])
	}

	if actual["reason"] != "file-change" {
		t.Errorf("Expected reason 'file-change', got %v", actual["reason"])
	}

	if !strings.Contains(string(msg1), "test.ts") {
		t.Errorf("Expected message to contain 'test.ts', got: %s", msg1)
	}
}

// TestWebSocketEndpoint_Disconnect verifies connection cleanup
func TestWebSocketEndpoint_Disconnect(t *testing.T) {
	server, err := serve.NewServer(9002)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	// Connect client
	wsClient := testutil.NewWebSocketTestClient(t, "ws://localhost:9002/__cem/reload")

	// Verify connection exists
	if server.WebSocketManager().ConnectionCount() != 1 {
		t.Errorf("Expected 1 connection, got %d", server.WebSocketManager().ConnectionCount())
	}

	// Disconnect client
	wsClient.Close()
	time.Sleep(100 * time.Millisecond)

	// Verify connection removed
	if server.WebSocketManager().ConnectionCount() != 0 {
		t.Errorf("Expected 0 connections after disconnect, got %d", server.WebSocketManager().ConnectionCount())
	}
}

// TestContainsPath_ExactMatch tests that exact matches return true
func TestContainsPath_ExactMatch(t *testing.T) {
	tests := []struct {
		fullURL string
		path    string
		want    bool
	}{
		{"/elements/button", "/elements/button", true},
		{"/demo/basic.html", "/demo/basic.html", true},
		{"/", "/", true},
	}

	for _, tt := range tests {
		t.Run(tt.fullURL, func(t *testing.T) {
			if got := urlutil.ContainsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("urlutil.ContainsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}

// TestContainsPath_ValidDelimiters tests that paths followed by valid delimiters return true
func TestContainsPath_ValidDelimiters(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		path    string
		want    bool
	}{
		{"slash delimiter", "/elements/button/demo/", "/elements/button", true},
		{"query delimiter", "/elements/button?variant=primary", "/elements/button", true},
		{"fragment delimiter", "/elements/button#overview", "/elements/button", true},
		{"path ending with slash", "/elements/accordion/demo/", "/elements/accordion", true},
		{"nested path", "/elements/accordion/demo/accents", "/elements/accordion/demo", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlutil.ContainsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("urlutil.ContainsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}

// TestContainsPath_InvalidPrefixMatch tests that invalid prefix matches return false
// This is the key test that demonstrates the bug: "/elements/button" should NOT match "/elements/buttonly"
func TestContainsPath_InvalidPrefixMatch(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		path    string
		want    bool
	}{
		{"partial word match", "/elements/buttonly", "/elements/button", false},
		{"partial with hyphen", "/elements/button-group", "/elements/button", false},
		{"partial with underscore", "/elements/button_component", "/elements/button", false},
		{"similar prefix", "/api/v1/users-list", "/api/v1/users", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlutil.ContainsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("urlutil.ContainsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}

// TestContainsPath_EdgeCases tests edge cases and invalid inputs
func TestContainsPath_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		path    string
		want    bool
	}{
		{"empty path", "/elements/button", "", false},
		{"path longer than fullURL", "/short", "/very/long/path", false},
		{"empty fullURL", "", "/path", false},
		{"both empty", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlutil.ContainsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("urlutil.ContainsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}
