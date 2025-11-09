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
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketTestClient is a test client for WebSocket connections.
// Note: Messages may be dropped when the message buffer (cap 100) is full.
// The error channel uses blocking sends to ensure connection errors are not lost.
type WebSocketTestClient struct {
	conn     *websocket.Conn
	messages chan []byte
	errors   chan error
	done     chan struct{}
}

// NewWebSocketTestClient creates a new WebSocket test client
func NewWebSocketTestClient(t *testing.T, url string) *WebSocketTestClient {
	t.Helper()

	dialer := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}

	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}

	client := &WebSocketTestClient{
		conn:     conn,
		messages: make(chan []byte, 100),
		errors:   make(chan error, 10),
		done:     make(chan struct{}),
	}

	// Start reading messages in background
	go client.readLoop()

	t.Cleanup(func() {
		client.Close()
	})

	return client
}

// readLoop reads messages from the WebSocket connection
func (c *WebSocketTestClient) readLoop() {
	defer close(c.done)

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			// Block to ensure error is not lost
			c.errors <- err
			return
		}

		select {
		case c.messages <- message:
		default:
			// Buffer full, message dropped (documented in type comment)
		}
	}
}

// ReceiveMessage waits for a message with timeout
func (c *WebSocketTestClient) ReceiveMessage(t *testing.T, timeout time.Duration) []byte {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case msg := <-c.messages:
		return msg
	case err := <-c.errors:
		t.Fatalf("WebSocket error: %v", err)
	case <-ctx.Done():
		t.Fatal("Timeout waiting for WebSocket message")
	}
	return nil
}

// SendMessage sends a message to the WebSocket
func (c *WebSocketTestClient) SendMessage(t *testing.T, message []byte) {
	t.Helper()

	if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
		t.Fatalf("Failed to send WebSocket message: %v", err)
	}
}

// Close closes the WebSocket connection
func (c *WebSocketTestClient) Close() {
	c.conn.Close()
	<-c.done
}

// ExpectNoMessage ensures no message is received within the timeout
func (c *WebSocketTestClient) ExpectNoMessage(t *testing.T, timeout time.Duration) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case msg := <-c.messages:
		t.Fatalf("Unexpected WebSocket message received: %s", msg)
	case err := <-c.errors:
		t.Fatalf("WebSocket error: %v", err)
	case <-ctx.Done():
		// Expected - no message received
		return
	}
}
