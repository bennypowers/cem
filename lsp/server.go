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
	"net/http"
	"os"
	"strconv"
	"sync"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/ephemeral"
	"bennypowers.dev/cem/lsp/helpers"
	lspTypes "bennypowers.dev/cem/lsp/types"
	"bennypowers.dev/cem/types"
	"github.com/gorilla/websocket"
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

	switch s.transport {
	case TransportStdio:
		return s.serve(&stdioRWC{})
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
		return s.serve(conn)
	case TransportWebSocket:
		return s.runWebSocket("localhost:8081")
	case TransportNodeJS:
		return s.runNodeJS()
	default:
		return fmt.Errorf("unsupported transport kind: %s", s.transport)
	}
}

// serve wires a ReadWriteCloser to the LSP protocol server.
// Client is set before conn.Go starts dispatching to avoid a data race
// between handler dispatch and s.client assignment.
func (s *Server) serve(rwc io.ReadWriteCloser) error {
	stream := jsonrpc2.NewStream(rwc)
	conn := jsonrpc2.NewConn(stream, jsonrpc2.WithCodec(protocolCodec{}))
	s.client = protocol.ClientDispatcher(conn)

	ctx := context.Background()
	ctx = protocol.WithClient(ctx, s.client)

	helpers.SetGlobalLoggerClient(s.client, ctx)

	conn.Go(ctx, protocol.Handlers(protocol.ServerHandler(s, jsonrpc2.MethodNotFoundHandler)))

	<-conn.Done()
	return conn.Err()
}

func (s *Server) runWebSocket(address string) error {
	upgrader := websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen for WebSocket on %s: %w", address, err)
	}
	defer func() { _ = listener.Close() }()

	helpers.SafeDebugLog("LSP: listening for WebSocket connection on %s", address)

	// Accept a single connection -- Server is single-session by design
	mux := http.NewServeMux()
	connCh := make(chan error, 1)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("websocket upgrade failed: %v", err), http.StatusBadRequest)
			return
		}
		defer func() { _ = wsConn.Close() }()
		connCh <- s.serve(newWSReadWriteCloser(wsConn))
	})

	srv := &http.Server{Handler: mux}
	go func() { _ = srv.Serve(listener) }()

	serveErr := <-connCh
	_ = srv.Close()
	return serveErr
}

func (s *Server) runNodeJS() error {
	fdStr := os.Getenv("NODE_CHANNEL_FD")
	if fdStr == "" {
		return fmt.Errorf("NODE_CHANNEL_FD not set in environment")
	}
	fd, err := strconv.Atoi(fdStr)
	if err != nil {
		return fmt.Errorf("invalid NODE_CHANNEL_FD %q: %w", fdStr, err)
	}
	if fd < 0 {
		return fmt.Errorf("NODE_CHANNEL_FD must be non-negative, got %d", fd)
	}
	file := os.NewFile(uintptr(fd), "/lsp/NODE_CHANNEL_FD")
	if file == nil {
		return fmt.Errorf("NODE_CHANNEL_FD %d is not a valid file descriptor", fd)
	}
	helpers.SafeDebugLog("LSP: serving via Node.js IPC on fd %d", fd)
	return s.serve(file)
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


// protocolCodec implements jsonrpc2.Codec using the protocol library's
// reflection-free marshal/unmarshal, matching the internal lspCodec
// used by protocol.NewServer.
type protocolCodec struct{}

func (protocolCodec) Marshal(v any) ([]byte, error)     { return protocol.Marshal(v) }
func (protocolCodec) Unmarshal(data []byte, v any) error { return protocol.Unmarshal(data, v) }

type stdioRWC struct{}

func (*stdioRWC) Read(p []byte) (int, error)  { return os.Stdin.Read(p) }
func (*stdioRWC) Write(p []byte) (int, error) { return os.Stdout.Write(p) }
func (*stdioRWC) Close() error                { return os.Stdin.Close() }

// wsReadWriteCloser adapts a gorilla/websocket.Conn to io.ReadWriteCloser
// for use with jsonrpc2.NewStream.
type wsReadWriteCloser struct {
	conn   *websocket.Conn
	reader io.Reader
}

func newWSReadWriteCloser(conn *websocket.Conn) *wsReadWriteCloser {
	return &wsReadWriteCloser{conn: conn}
}

func (ws *wsReadWriteCloser) Read(p []byte) (int, error) {
	for {
		if ws.reader == nil {
			_, reader, err := ws.conn.NextReader()
			if err != nil {
				return 0, err
			}
			ws.reader = reader
		}
		n, err := ws.reader.Read(p)
		if err == io.EOF {
			ws.reader = nil
			if n > 0 {
				return n, nil
			}
			continue
		}
		return n, err
	}
}

func (ws *wsReadWriteCloser) Write(p []byte) (int, error) {
	w, err := ws.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	n, err := w.Write(p)
	if closeErr := w.Close(); err == nil {
		err = closeErr
	}
	return n, err
}

func (ws *wsReadWriteCloser) Close() error {
	return ws.conn.Close()
}

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
