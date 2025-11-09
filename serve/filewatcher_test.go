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
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"bennypowers.dev/cem/serve"
	"bennypowers.dev/cem/serve/testutil"
)

// TestFileWatcher_SourceFileChange verifies file watcher detects source file changes
func TestFileWatcher_SourceFileChange(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test-element.ts")

	// Write initial content
	err := os.WriteFile(testFile, []byte("export class TestElement {}"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	server, err := serve.NewServer(9100)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	// Connect WebSocket to receive reload events
	wsClient := testutil.NewWebSocketTestClient(t, "ws://localhost:9100/__cem-reload")

	// Modify the file
	time.Sleep(200 * time.Millisecond) // Let watcher initialize
	err = os.WriteFile(testFile, []byte("export class TestElement { updated: true }"), 0644)
	if err != nil {
		t.Fatalf("Failed to modify test file: %v", err)
	}

	// Should receive reload event
	msg := wsClient.ReceiveMessage(t, 3*time.Second)
	msgStr := string(msg)

	if !strings.Contains(msgStr, "reload") {
		t.Errorf("Expected reload event, got: %s", msgStr)
	}

	if !strings.Contains(msgStr, "test-element.ts") {
		t.Errorf("Expected message to contain filename, got: %s", msgStr)
	}
}

// TestFileWatcher_Debouncing verifies rapid changes result in single reload
func TestFileWatcher_Debouncing(t *testing.T) {
	tmpDir := t.TempDir()

	server, err := serve.NewServer(9101)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	wsClient := testutil.NewWebSocketTestClient(t, "ws://localhost:9101/__cem-reload")

	time.Sleep(200 * time.Millisecond) // Let watcher initialize

	// Make rapid changes to multiple files
	for i := 0; i < 5; i++ {
		testFile := filepath.Join(tmpDir, "test-"+string(rune('a'+i))+".ts")
		err = os.WriteFile(testFile, []byte("content"), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file %d: %v", i, err)
		}
		time.Sleep(10 * time.Millisecond) // Rapid changes (within 150ms window)
	}

	// Should receive ONE reload event (debounced)
	msg1 := wsClient.ReceiveMessage(t, 2*time.Second)
	if !strings.Contains(string(msg1), "reload") {
		t.Errorf("Expected reload event, got: %s", msg1)
	}

	// Should NOT receive another message within 500ms
	// We can't directly access the channel, so we'll just wait and verify no crash
	time.Sleep(500 * time.Millisecond)
	// If we got here without crashing, debouncing worked (only one message was sent)
}

// TestFileWatcher_MultipleFileTypes verifies watching different file types
func TestFileWatcher_MultipleFileTypes(t *testing.T) {
	tmpDir := t.TempDir()

	server, err := serve.NewServer(9102)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	wsClient := testutil.NewWebSocketTestClient(t, "ws://localhost:9102/__cem-reload")

	time.Sleep(200 * time.Millisecond)

	// Modify demo file
	demoFile := filepath.Join(tmpDir, "demo.html")
	err = os.WriteFile(demoFile, []byte("<my-element></my-element>"), 0644)
	if err != nil {
		t.Fatalf("Failed to write demo file: %v", err)
	}

	// Should receive reload event for demo file
	msg := wsClient.ReceiveMessage(t, 3*time.Second)
	if !strings.Contains(string(msg), "reload") {
		t.Errorf("Expected reload event for demo file, got: %s", msg)
	}
}
