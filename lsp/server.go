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
package lsp

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/ephemeral"
	"bennypowers.dev/cem/lsp/helpers"
	lspTypes "bennypowers.dev/cem/lsp/types"
	"bennypowers.dev/cem/types"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

// TransportKind represents different LSP transport methods
type TransportKind string

const (
	TransportStdio     TransportKind = "stdio"
	TransportTCP       TransportKind = "tcp"
	TransportWebSocket TransportKind = "websocket"
	TransportNodeJS    TransportKind = "nodejs"
)

// Server represents the CEM LSP server
type Server struct {
	protocol.UnimplementedServer

	workspace          types.WorkspaceContext
	registry           *Registry
	ephemeralRegistry  *ephemeral.Registry
	documents          lspTypes.Manager
	client             protocol.Client
	transport          TransportKind
	additionalPackages []string
	config             lspTypes.ServerConfig
	configMu           sync.RWMutex
	usePullDiagnostics bool
}

// NewServer creates a new CEM LSP server
func NewServer(workspace types.WorkspaceContext, transport TransportKind) (*Server, error) {
	documents, err := document.NewDocumentManager()
	if err != nil {
		return nil, fmt.Errorf("failed to create document manager: %w", err)
	}

	registry, err := NewRegistryWithDefaults()
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	s := &Server{
		workspace:         workspace,
		registry:          registry,
		ephemeralRegistry: ephemeral.NewRegistry(),
		documents:         documents,
		transport:         transport,
		config:            lspTypes.DefaultConfig(),
	}

	return s, nil
}

// Run starts the LSP server using the configured transport
func (s *Server) Run() error {
	helpers.SafeDebugLog("LSP: Running with transport: %s", s.transport)

	var rwc io.ReadWriteCloser

	switch s.transport {
	case TransportStdio:
		rwc = &stdioRWC{}
	case TransportTCP:
		listener, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			return fmt.Errorf("failed to listen on TCP: %w", err)
		}
		defer func() { _ = listener.Close() }()
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept TCP connection: %w", err)
		}
		rwc = conn
	case TransportWebSocket:
		return fmt.Errorf("websocket transport not yet implemented for new protocol library")
	case TransportNodeJS:
		return fmt.Errorf("nodejs transport not yet implemented for new protocol library")
	default:
		return fmt.Errorf("unsupported transport kind: %s", s.transport)
	}

	stream := jsonrpc2.NewStream(rwc)
	ctx := context.Background()
	ctx, conn, client := protocol.NewServer(ctx, s, stream)
	s.client = client

	<-conn.Done()
	_ = ctx
	return nil
}

// Close cleans up server resources
func (s *Server) Close() error {
	if err := s.registry.StopFileWatching(); err != nil {
		helpers.SafeDebugLog("Warning: Error stopping file watcher: %v", err)
	}

	if err := s.registry.StopGenerateWatcher(); err != nil {
		helpers.SafeDebugLog("Warning: Error stopping generate watcher: %v", err)
	}

	if s.documents != nil {
		s.documents.Close()
	}

	return nil
}

// handleManifestReload is called when manifest files change
func (s *Server) handleManifestReload() {
	helpers.SafeDebugLog("=== MANIFEST RELOAD TRIGGERED ===")
	helpers.SafeDebugLog("Reloading manifests due to file changes...")

	helpers.SafeDebugLog("Before reload: %d elements in registry", len(s.registry.Elements))
	for tagName := range s.registry.Elements {
		helpers.SafeDebugLog("  - Element: %s", tagName)
		if attrs, exists := s.registry.Attributes(tagName); exists {
			for attrName, attr := range attrs {
				typeText := ""
				if attr.Type != nil {
					typeText = attr.Type.Text
				}
				helpers.SafeDebugLog("    - Attribute %s: %s", attrName, typeText)
			}
		}
	}

	if err := s.reloadManifestsDirectly(); err != nil {
		helpers.SafeDebugLog("Error reloading manifests directly: %v", err)
		return
	}

	helpers.SafeDebugLog("After reload: %d elements in registry", len(s.registry.Elements))
	for tagName := range s.registry.Elements {
		helpers.SafeDebugLog("  - Element: %s", tagName)
		if attrs, exists := s.registry.Attributes(tagName); exists {
			for attrName, attr := range attrs {
				typeText := ""
				if attr.Type != nil {
					typeText = attr.Type.Text
				}
				helpers.SafeDebugLog("    - Attribute %s: %s", attrName, typeText)
			}
		}
	}

	helpers.SafeDebugLog("Successfully reloaded manifests: %d elements available", len(s.registry.Elements))
	helpers.SafeDebugLog("=== MANIFEST RELOAD COMPLETE ===")
}

// reloadManifestsDirectly bypasses workspace caching by directly reading manifest files
func (s *Server) reloadManifestsDirectly() error {
	return s.registry.ReloadManifestsDirectly()
}

// Registry returns the registry for testing purposes
func (s *Server) Registry() *Registry {
	return s.registry
}

type stdioRWC struct{}

func (*stdioRWC) Read(p []byte) (int, error)  { return os.Stdin.Read(p) }
func (*stdioRWC) Write(p []byte) (int, error) { return os.Stdout.Write(p) }
func (*stdioRWC) Close() error                { return os.Stdin.Close() }

// InitializeForTesting initializes the server without running stdio - for testing only
func (s *Server) InitializeForTesting() error {
	if err := s.registry.LoadFromWorkspace(s.workspace); err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	if err := s.registry.StartFileWatching(s.handleManifestReload); err != nil {
		helpers.SafeDebugLog("Warning: Could not start file watching: %v", err)
	}

	if err := s.registry.StartGenerateWatcher(); err != nil {
		helpers.SafeDebugLog("Warning: Could not start generate watcher: %v", err)
	}

	return nil
}
