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
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"bennypowers.dev/cem/serve/internal/urlutil"
	"github.com/gorilla/websocket"
)

const (
	// maxWebSocketReadSize is the maximum size of messages the server will accept from clients
	// Clients shouldn't be sending us data, but if they do, limit it to prevent DoS attacks
	maxWebSocketReadSize = 64 * 1024 // 64 KB
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 64 * 1024, // 64KB to handle large error messages with stack traces
	CheckOrigin:     isLocalOrigin,
}

// isLocalOrigin checks if the WebSocket connection is from a localhost origin
// or from the same origin as the request (for Cloudflare tunnels, reverse proxies, etc.)
// This prevents cross-origin WebSocket connections from untrusted domains
func isLocalOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	if origin == "" {
		// No origin header - allow (some WebSocket clients don't send Origin)
		return true
	}

	originURL, err := url.Parse(origin)
	if err != nil {
		// Invalid origin - reject
		return false
	}

	// Extract hostname, removing port if present
	originHost := originURL.Hostname()

	// Check if origin matches the request's Host header
	// This allows connections through Cloudflare tunnels, reverse proxies, etc.
	// where Origin might be "https://my-app.example.com" and Host is "my-app.example.com"
	requestHost := r.Host
	if colonIndex := strings.IndexByte(requestHost, ':'); colonIndex != -1 {
		requestHost = requestHost[:colonIndex] // Strip port from Host header
	}
	if originHost == requestHost {
		return true
	}

	// Allow localhost, 127.0.0.0/8, and IPv6 localhost
	if originHost == "localhost" || originHost == "127.0.0.1" || originHost == "[::1]" || originHost == "::1" {
		return true
	}

	// Allow *.localhost subdomains (e.g., app.localhost)
	if strings.HasSuffix(originHost, ".localhost") {
		return true
	}

	// Allow 127.0.0.0/8 range (127.0.0.1 - 127.255.255.255)
	// Must check it's in the form 127.X.X.X, not just starts with "127"
	if strings.HasPrefix(originHost, "127.") {
		// Validate it's a proper 127.x.x.x address, not something like "127.evil.com"
		parts := strings.Split(originHost, ".")
		if len(parts) == 4 && parts[0] == "127" {
			// Basic validation that it looks like an IP (has 4 parts starting with 127)
			return true
		}
	}

	return false
}

// connWrapper wraps a WebSocket connection with a write mutex
type connWrapper struct {
	conn    *websocket.Conn
	mu      sync.Mutex
	pageURL string // The demo page URL this connection is viewing
}

// websocketManager implements WebSocketManager interface
type websocketManager struct {
	connections map[*websocket.Conn]*connWrapper
	mu          sync.RWMutex
	logger      Logger
}

// newWebSocketManager creates a new WebSocket manager
func newWebSocketManager() WebSocketManager {
	return &websocketManager{
		connections: make(map[*websocket.Conn]*connWrapper),
	}
}

// SetLogger sets the logger for the WebSocket manager
func (wm *websocketManager) SetLogger(logger Logger) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	wm.logger = logger
}

// ConnectionCount returns the number of active connections
func (wm *websocketManager) ConnectionCount() int {
	wm.mu.RLock()
	defer wm.mu.RUnlock()
	return len(wm.connections)
}

// Broadcast sends a message to all connected clients
func (wm *websocketManager) Broadcast(message []byte) error {
	// Snapshot connections while holding read lock to avoid blocking connects/disconnects
	wm.mu.RLock()
	snapshot := make([]*connWrapper, 0, len(wm.connections))
	for _, wrapper := range wm.connections {
		snapshot = append(snapshot, wrapper)
	}
	wm.mu.RUnlock()

	// Write to all connections without holding manager lock
	// so slow clients don't block connects/disconnects
	var failedConnections []*websocket.Conn
	for _, wrapper := range snapshot {
		// Lock this connection's write mutex to prevent concurrent writes
		wrapper.mu.Lock()
		err := wrapper.conn.WriteMessage(websocket.TextMessage, message)
		wrapper.mu.Unlock()

		if err != nil {
			// Connection is dead, mark for cleanup
			failedConnections = append(failedConnections, wrapper.conn)
			// Continue broadcasting to other clients
		}
	}

	// Clean up failed connections
	if len(failedConnections) > 0 {
		wm.mu.Lock()
		for _, conn := range failedConnections {
			delete(wm.connections, conn)
			_ = conn.Close()
		}
		wm.mu.Unlock()
	}

	return nil
}

// BroadcastShutdown sends a shutdown notification to all clients before closing
// Sets write deadlines to prevent blocking on unresponsive clients during shutdown
func (wm *websocketManager) BroadcastShutdown() error {
	shutdownMsg := []byte(`{"type":"shutdown","reason":"server-shutdown"}`)

	// Snapshot connections while holding read lock
	wm.mu.RLock()
	snapshot := make([]*connWrapper, 0, len(wm.connections))
	for _, wrapper := range wm.connections {
		snapshot = append(snapshot, wrapper)
	}
	wm.mu.RUnlock()

	// Write shutdown messages with timeouts to prevent hanging on unresponsive clients
	var failedConnections []*websocket.Conn
	for _, wrapper := range snapshot {
		wrapper.mu.Lock()
		// Set write deadline of 1 second - if client isn't reading, we'll time out
		_ = wrapper.conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		err := wrapper.conn.WriteMessage(websocket.TextMessage, shutdownMsg)
		wrapper.mu.Unlock()

		if err != nil {
			failedConnections = append(failedConnections, wrapper.conn)
		}
	}

	// Clean up failed connections
	if len(failedConnections) > 0 {
		wm.mu.Lock()
		for _, conn := range failedConnections {
			delete(wm.connections, conn)
			_ = conn.Close()
		}
		wm.mu.Unlock()
	}

	return nil
}

// CloseAll gracefully closes all WebSocket connections
// Sets write deadlines to prevent blocking on unresponsive clients during shutdown
func (wm *websocketManager) CloseAll() error {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	for conn, wrapper := range wm.connections {
		// Send close frame with normal closure code and 500ms timeout
		// If client isn't reading, we'll time out and force close anyway
		wrapper.mu.Lock()
		_ = conn.SetWriteDeadline(time.Now().Add(500 * time.Millisecond))
		_ = conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server shutting down"))
		wrapper.mu.Unlock()

		// Close the underlying connection (this will make read loops exit)
		_ = conn.Close()
	}

	// Clear the connections map
	wm.connections = make(map[*websocket.Conn]*connWrapper)

	return nil
}

// BroadcastToPages sends a message only to clients viewing specific page URLs
// pageURLs can contain partial matches (e.g., "/elements/accordion/demo/" will match all accordion demos)
func (wm *websocketManager) BroadcastToPages(message []byte, pageURLs []string) error {
	if len(pageURLs) == 0 {
		return nil
	}

	// Snapshot connections while holding read lock
	wm.mu.RLock()

	snapshot := make([]*connWrapper, 0, len(wm.connections))
	for _, wrapper := range wm.connections {
		// Check if this connection's page matches any of the target URLs
		for _, targetURL := range pageURLs {
			if wrapper.pageURL == targetURL || urlutil.ContainsPath(wrapper.pageURL, targetURL) {
				snapshot = append(snapshot, wrapper)
				break
			}
		}
	}
	wm.mu.RUnlock()

	// Write to matched connections without holding manager lock
	var failedConnections []*websocket.Conn
	for _, wrapper := range snapshot {
		wrapper.mu.Lock()
		err := wrapper.conn.WriteMessage(websocket.TextMessage, message)
		wrapper.mu.Unlock()

		if err != nil {
			// Connection is dead, mark for cleanup
			failedConnections = append(failedConnections, wrapper.conn)
			// Continue broadcasting to other clients
		}
	}

	// Clean up failed connections
	if len(failedConnections) > 0 {
		wm.mu.Lock()
		for _, conn := range failedConnections {
			delete(wm.connections, conn)
			_ = conn.Close()
		}
		wm.mu.Unlock()
	}

	if wm.logger != nil && len(snapshot) > 0 {
		wm.logger.Debug("Broadcast to %d/%d connections (targeted pages)", len(snapshot), wm.ConnectionCount())
	}

	return nil
}

// HandleConnection handles a new WebSocket connection
func (wm *websocketManager) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if wm.logger != nil {
			wm.logger.Error("Failed to upgrade WebSocket: %v", err)
		}
		return
	}

	// Set maximum read size to prevent DoS attacks via large messages
	// Clients shouldn't be sending us data, but if they do, limit it
	conn.SetReadLimit(maxWebSocketReadSize)

	// Remove HTTP server timeouts for WebSocket connections
	// WebSockets are long-lived connections that should not be subject to write/read timeouts
	if err := conn.UnderlyingConn().SetDeadline(time.Time{}); err != nil {
		if wm.logger != nil {
			wm.logger.Warning("Failed to clear deadline on WebSocket connection: %v", err)
		}
	}

	// Extract page URL from query parameter (preferred) or Referer header (fallback)
	// Browsers don't send Referer header for WebSocket connections, so client passes it as ?page=...
	pageURL := r.URL.Query().Get("page")
	// Fallback to Referer header if query param not present
	if pageURL == "" {
		referer := r.Header.Get("Referer")
		if referer != "" {
			// Parse the Referer URL to extract just the path
			// Referer will be something like "http://localhost:8000/elements/accordion/demo/"
			// We need to extract just "/elements/accordion/demo/"
			if parsedURL, err := url.Parse(referer); err == nil {
				pageURL = parsedURL.Path
				if wm.logger != nil {
					wm.logger.Info("Parsed Referer to path: %s", pageURL)
				}
			} else {
				if wm.logger != nil {
					wm.logger.Warning("Failed to parse Referer URL: %v", err)
				}
			}
		}
	}

	// Last resort: use request URL path
	if pageURL == "" {
		pageURL = r.URL.Path
		if wm.logger != nil {
			wm.logger.Warning("No page URL from query or Referer, using request path: %s", pageURL)
		}
	}

	// Register connection with wrapper
	wrapper := &connWrapper{
		conn:    conn,
		pageURL: pageURL,
	}
	wm.mu.Lock()
	wm.connections[conn] = wrapper
	count := len(wm.connections)
	wm.mu.Unlock()

	if wm.logger != nil {
		wm.logger.Debug("WebSocket client connected from %s (total: %d)", pageURL, count)
	}

	// Handle disconnection
	defer func() {
		wm.mu.Lock()
		delete(wm.connections, conn)
		wm.mu.Unlock()

		_ = conn.Close()

		if wm.logger != nil {
			wm.logger.Debug("WebSocket client disconnected (total: %d)", wm.ConnectionCount())
		}
	}()

	// Read loop (to detect disconnects)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
