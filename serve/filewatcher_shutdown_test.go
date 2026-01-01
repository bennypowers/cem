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

/*
Regression test for file watcher shutdown hanging issue.

This test reproduces the bug where the file watcher's select statement
gets starved by a flood of events, preventing the shutdown signal from
being processed.

Original issue: cem serve process hangs with 800%+ CPU after tests complete,
launched by make test-frontend, PID 2300782 and PID 2601898.

Root cause: serve/filewatcher.go:164-227 - select statement doesn't prioritize
the done channel, so continuous events can prevent shutdown.
*/
package serve

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestServerShutdownUnderFileFlood tests that the server can shut down
// cleanly even when the file watcher is receiving a continuous flood of events.
//
// This is a regression test for the hanging serve process bug.
func TestServerShutdownUnderFileFlood(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping file flood test in short mode")
	}

	tmpDir := t.TempDir()

	// Create server
	server, err := NewServer(19900)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Set watch directory
	if err := server.SetWatchDir(tmpDir); err != nil {
		t.Fatalf("Failed to set watch directory: %v", err)
	}

	// Start server
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Give server time to initialize watcher
	time.Sleep(200 * time.Millisecond)

	// Start flooding the watched directory with file changes
	stopFlood := make(chan struct{})
	floodDone := make(chan struct{})

	go func() {
		defer close(floodDone)

		for i := 0; ; i++ {
			select {
			case <-stopFlood:
				return
			default:
				// Create/modify many files rapidly
				testFile := filepath.Join(tmpDir, fmt.Sprintf("test-%d.ts", i%10))
				content := fmt.Sprintf("// File %d\nexport const test%d = %d;\n", i, i, i)
				_ = os.WriteFile(testFile, []byte(content), 0644)

				// Minimal delay to flood the watcher
				if i%10 == 0 {
					time.Sleep(1 * time.Millisecond)
				}
			}
		}
	}()

	// Let the flood run for a bit to build up events
	time.Sleep(500 * time.Millisecond)

	// Now try to shut down the server while events are still flooding
	// This is the critical test - shutdown should succeed quickly
	shutdownStart := time.Now()

	// Close the server (this should close the file watcher)
	if err := server.Close(); err != nil {
		t.Errorf("Server close returned error: %v", err)
	}

	shutdownDuration := time.Since(shutdownStart)

	// Stop the flood
	close(stopFlood)
	<-floodDone

	// Shutdown should complete quickly (< 2 seconds)
	// If it takes longer, the file watcher's done channel is being starved
	maxShutdownTime := 2 * time.Second
	if shutdownDuration > maxShutdownTime {
		t.Errorf("Server shutdown took too long: %v (max %v) - file watcher done channel likely starved by events",
			shutdownDuration, maxShutdownTime)
		t.Error("This is the SELECT STATEMENT STARVATION BUG in serve/filewatcher.go:164-227")
	} else {
		t.Logf("Server shutdown completed in %v", shutdownDuration)
	}
}
