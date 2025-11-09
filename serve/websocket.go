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
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for dev server
	},
}

// websocketManager implements WebSocketManager interface
type websocketManager struct {
	connections map[*websocket.Conn]bool
	mu          sync.RWMutex
	logger      Logger
}

// newWebSocketManager creates a new WebSocket manager
func newWebSocketManager() WebSocketManager {
	return &websocketManager{
		connections: make(map[*websocket.Conn]bool),
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
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	for conn := range wm.connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			if wm.logger != nil {
				wm.logger.Error("Failed to send WebSocket message: %v", err)
			}
			// Continue broadcasting to other clients even if one fails
		}
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

	// Register connection
	wm.mu.Lock()
	wm.connections[conn] = true
	count := len(wm.connections)
	wm.mu.Unlock()

	if wm.logger != nil {
		wm.logger.Debug("WebSocket client connected (total: %d)", count)
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
