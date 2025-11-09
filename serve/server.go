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
	"fmt"
	"net/http"
)

// Server represents the development server
type Server struct {
	Port    int
	Handler http.Handler
}

// NewServer creates a new development server
// TODO: Implement in Phase 1
func NewServer(port int) (*Server, error) {
	return nil, fmt.Errorf("not implemented: Phase 1 - Core Server")
}

// Start starts the development server
// TODO: Implement in Phase 1
func (s *Server) Start() error {
	return fmt.Errorf("not implemented: Phase 1 - Core Server")
}

// Stop stops the development server
// TODO: Implement in Phase 1
func (s *Server) Stop() error {
	return fmt.Errorf("not implemented: Phase 1 - Core Server")
}

// WebSocketEndpoint handles WebSocket connections for live reload
// TODO: Implement in Phase 1
func (s *Server) WebSocketEndpoint(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented: Phase 1 - Core Server", http.StatusNotImplemented)
}
