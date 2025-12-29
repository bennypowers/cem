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

package testutil

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"
)

// TestServer wraps an HTTP server for testing
type TestServer struct {
	Server   *http.Server
	Port     int
	BaseURL  string
	started  bool
	stopChan chan struct{}
}

// NewTestServer creates a new test server with an available port
func NewTestServer(t *testing.T, handler http.Handler) *TestServer {
	t.Helper()

	// Find an available port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to find available port: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	ts := &TestServer{
		Server:   server,
		Port:     port,
		BaseURL:  fmt.Sprintf("http://127.0.0.1:%d", port),
		stopChan: make(chan struct{}),
	}

	t.Cleanup(func() {
		ts.Stop(t)
	})

	return ts
}

// Start starts the test server
func (ts *TestServer) Start(t *testing.T) {
	t.Helper()

	if ts.started {
		t.Fatal("Server already started")
	}

	go func() {
		if err := ts.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Logf("Server error: %v", err)
		}
		close(ts.stopChan)
	}()

	// Wait for server to be ready
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			t.Fatal("Timeout waiting for server to start")
		case <-ts.stopChan:
			t.Fatal("Server stopped unexpectedly during startup")
		default:
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", ts.Port), 100*time.Millisecond)
			if err == nil {
				conn.Close()
				ts.started = true
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Stop stops the test server
func (ts *TestServer) Stop(t *testing.T) {
	t.Helper()

	if !ts.started {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ts.Server.Shutdown(ctx); err != nil {
		t.Logf("Server shutdown error: %v", err)
	}

	<-ts.stopChan
	ts.started = false
}
