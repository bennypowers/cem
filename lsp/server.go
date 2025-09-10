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
	"fmt"
	"os"

	"bennypowers.dev/cem/lsp/helpers"
	serverMethods "bennypowers.dev/cem/lsp/methods/server"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/methods/textDocument/codeAction"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/methods/textDocument/definition"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/methods/textDocument/references"
	"bennypowers.dev/cem/lsp/methods/workspace/symbol"
	"bennypowers.dev/cem/types"
	"github.com/pterm/pterm"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
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
	workspace types.WorkspaceContext
	registry  *Registry
	documents *DocumentManager
	server    *server.Server
	transport TransportKind
}

// NewServer creates a new CEM LSP server
func NewServer(workspace types.WorkspaceContext, transport TransportKind) (*Server, error) {
	// Configure pterm to output to stderr to avoid contaminating LSP stdout stream
	pterm.SetDefaultOutput(os.Stderr)

	documents, err := NewDocumentManager()
	if err != nil {
		return nil, fmt.Errorf("failed to create document manager: %w", err)
	}

	// Create registry with production defaults
	registry, err := NewRegistryWithDefaults()
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	s := &Server{
		workspace: workspace,
		registry:  registry,
		documents: documents,
		transport: transport,
	}

	// Server now directly implements all context interfaces

	// Create the GLSP server with our handler
	handler := protocol.Handler{
		Initialize:             s.initialize,
		Initialized:            s.initialized,
		Shutdown:               s.shutdown,
		SetTrace:               s.setTrace,
		TextDocumentHover:      s.hover,
		TextDocumentCompletion: s.completion,
		TextDocumentDefinition: s.definition,
		TextDocumentReferences: s.references,
		TextDocumentCodeAction: s.codeAction,
		TextDocumentDidOpen:    s.didOpen,
		TextDocumentDidChange:  s.didChange,
		TextDocumentDidClose:   s.didClose,
		WorkspaceSymbol:        s.workspaceSymbol,
	}

	// Enable debug mode when using stdio to help with VSCode troubleshooting
	debug := transport == TransportStdio
	s.server = server.NewServer(&handler, "cem-lsp", debug)

	return s, nil
}

// Run starts the LSP server using the configured transport
func (s *Server) Run() error {
	// Manifest initialization now happens in the LSP Initialized method
	// This ensures proper LSP protocol compliance and debug logging visibility

	helpers.SafeDebugLog("LSP: Running with transport: %s", s.transport)

	// Run the server with the appropriate transport
	switch s.transport {
	case TransportStdio:
		return s.server.RunStdio()
	case TransportTCP:
		// For now, use a default TCP address - this could be made configurable
		return s.server.RunTCP("localhost:8080")
	case TransportWebSocket:
		// For now, use a default WebSocket address - this could be made configurable
		return s.server.RunWebSocket("localhost:8081")
	case TransportNodeJS:
		return s.server.RunNodeJs()
	default:
		return fmt.Errorf("unsupported transport kind: %s", s.transport)
	}
}

// Close cleans up server resources
func (s *Server) Close() error {
	// Stop file watching
	if err := s.registry.StopFileWatching(); err != nil {
		helpers.SafeDebugLog("Warning: Error stopping file watcher: %v", err)
	}

	// Stop generate watcher
	if err := s.registry.StopGenerateWatcher(); err != nil {
		helpers.SafeDebugLog("Warning: Error stopping generate watcher: %v", err)
	}

	// Close document manager
	if s.documents != nil {
		s.documents.Close()
	}

	return nil
}

// handleManifestReload is called when manifest files change
func (s *Server) handleManifestReload() {
	helpers.SafeDebugLog("=== MANIFEST RELOAD TRIGGERED ===")
	helpers.SafeDebugLog("Reloading manifests due to file changes...")

	// Debug: Show current registry state before reload
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

	// The workspace context caches manifests and we can't clear that cache.
	// Instead, we'll bypass the workspace cache by directly reading manifest files
	// and reloading them into the registry.

	if err := s.reloadManifestsDirectly(); err != nil {
		helpers.SafeDebugLog("Error reloading manifests directly: %v", err)
		return
	}

	// Debug: Show registry state after reload
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

// InitializeForTesting initializes the server without running stdio - for testing only
func (s *Server) InitializeForTesting() error {
	// Initialize the manifest registry
	if err := s.registry.LoadFromWorkspace(s.workspace); err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	// Start file watching for manifest changes
	if err := s.registry.StartFileWatching(s.handleManifestReload); err != nil {
		helpers.SafeDebugLog("Warning: Could not start file watching: %v", err)
		// Don't fail startup if file watching fails
	}

	// Start generate watcher for local project source file changes
	if err := s.registry.StartGenerateWatcher(); err != nil {
		helpers.SafeDebugLog("Warning: Could not start generate watcher: %v", err)
		// Don't fail startup if generate watcher fails
	}

	return nil
}

// initialize handles the LSP initialize request
func (s *Server) initialize(
	context *glsp.Context,
	params *protocol.InitializeParams,
) (any, error) {
	return serverMethods.Initialize(s, context, params)
}

// initialized handles the LSP initialized notification
func (s *Server) initialized(
	context *glsp.Context,
	params *protocol.InitializedParams,
) error {
	return serverMethods.Initialized(s, context, params)
}

// shutdown handles the LSP shutdown request
func (s *Server) shutdown(
	context *glsp.Context,
) error {
	return serverMethods.Shutdown(s, context)
}

// setTrace handles the $/setTrace notification
func (s *Server) setTrace(
	context *glsp.Context,
	params *protocol.SetTraceParams,
) error {
	return serverMethods.SetTrace(s, context, params)
}

// hover handles textDocument/hover requests
func (s *Server) hover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	return hover.Hover(s, context, params)
}

// completion handles textDocument/completion requests
func (s *Server) completion(
	context *glsp.Context,
	params *protocol.CompletionParams,
) (any, error) {
	return completion.Completion(s, context, params)
}

// definition handles textDocument/definition requests
func (s *Server) definition(
	context *glsp.Context,
	params *protocol.DefinitionParams,
) (any, error) {
	return definition.Definition(s, context, params)
}

// references handles textDocument/references requests
func (s *Server) references(
	context *glsp.Context,
	params *protocol.ReferenceParams,
) ([]protocol.Location, error) {
	return references.References(s, context, params)
}

// codeAction handles textDocument/codeAction requests
func (s *Server) codeAction(
	context *glsp.Context,
	params *protocol.CodeActionParams,
) (any, error) {
	return codeAction.CodeAction(s, context, params)
}

// workspaceSymbol handles workspace/symbol requests
func (s *Server) workspaceSymbol(
	context *glsp.Context,
	params *protocol.WorkspaceSymbolParams,
) ([]protocol.SymbolInformation, error) {
	return symbol.Symbol(s, context, params)
}

// didOpen handles textDocument/didOpen notifications
func (s *Server) didOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	return textDocument.DidOpen(s, context, params)
}

// didChange handles textDocument/didChange notifications
func (s *Server) didChange(
	context *glsp.Context,
	params *protocol.DidChangeTextDocumentParams,
) error {
	return textDocument.DidChange(s, context, params)
}

// didClose handles textDocument/didClose notifications
func (s *Server) didClose(
	context *glsp.Context,
	params *protocol.DidCloseTextDocumentParams,
) error {
	return textDocument.DidClose(s, context, params)
}
