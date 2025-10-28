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

package cmd_test

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

// jsonrpcRequest represents a JSON-RPC 2.0 request
type jsonrpcRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int64  `json:"id,omitempty"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
}

// jsonrpcResponse represents a JSON-RPC 2.0 response
type jsonrpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int64           `json:"id,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *jsonrpcError   `json:"error,omitempty"`
}

// jsonrpcError represents a JSON-RPC 2.0 error
type jsonrpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// lspClient implements a minimal LSP client over stdin/stdout
type lspClient struct {
	stdin      io.WriteCloser
	stdout     *bufio.Reader
	cmd        *exec.Cmd
	mu         sync.Mutex
	nextID     atomic.Int64
	pending    map[int64]chan jsonrpcResponse
	pendingMu  sync.Mutex
}

func (c *lspClient) readMessage() (json.RawMessage, error) {
	// Read headers
	var contentLength int
	for {
		line, err := c.stdout.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("reading header: %w", err)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break // Empty line signals end of headers
		}

		if strings.HasPrefix(line, "Content-Length:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid Content-Length header: %s", line)
			}
			contentLength, err = strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return nil, fmt.Errorf("parsing Content-Length: %w", err)
			}
		}
	}

	if contentLength == 0 {
		return nil, fmt.Errorf("missing Content-Length header")
	}

	// Read body
	body := make([]byte, contentLength)
	if _, err := io.ReadFull(c.stdout, body); err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	return body, nil
}

func (c *lspClient) writeMessage(msg any) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshaling message: %w", err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body))
	if _, err := c.stdin.Write([]byte(header)); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	if _, err := c.stdin.Write(body); err != nil {
		return fmt.Errorf("writing body: %w", err)
	}

	return nil
}

func (c *lspClient) handleResponses() {
	for {
		msg, err := c.readMessage()
		if err != nil {
			// Connection closed or error
			return
		}

		// Peek to detect server-side requests/notifications
		var meta struct {
			ID     *int64 `json:"id,omitempty"`
			Method string `json:"method,omitempty"`
		}
		if err := json.Unmarshal(msg, &meta); err == nil && meta.Method != "" {
			// If it's a request, reply with "method not found" to avoid server blocking.
			if meta.ID != nil {
				_ = c.writeMessage(jsonrpcResponse{
					JSONRPC: "2.0",
					ID:      *meta.ID,
					Error:   &jsonrpcError{Code: -32601, Message: "method not supported by test client"},
				})
			}
			continue
		}

		// Otherwise it's a response
		var resp jsonrpcResponse
		if err := json.Unmarshal(msg, &resp); err != nil {
			continue // Skip malformed
		}
		c.pendingMu.Lock()
		if ch, ok := c.pending[resp.ID]; ok {
			ch <- resp
			delete(c.pending, resp.ID)
		}
		c.pendingMu.Unlock()
	}
}

func (c *lspClient) call(method string, params, result any) error {
	id := c.nextID.Add(1)

	// Register response channel
	respChan := make(chan jsonrpcResponse, 1)
	c.pendingMu.Lock()
	c.pending[id] = respChan
	c.pendingMu.Unlock()

	// Send request
	req := jsonrpcRequest{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}

	if err := c.writeMessage(req); err != nil {
		c.pendingMu.Lock()
		delete(c.pending, id)
		c.pendingMu.Unlock()
		return err
	}

	// Wait for response with timeout
	var resp jsonrpcResponse
	select {
	case resp = <-respChan:
		// Response received
	case <-time.After(30 * time.Second):
		// Timeout - clean up pending map entry
		c.pendingMu.Lock()
		delete(c.pending, id)
		c.pendingMu.Unlock()
		return fmt.Errorf("timeout waiting for response to %s (id=%d)", method, id)
	}

	// Check for error in response
	if resp.Error != nil {
		return fmt.Errorf("jsonrpc error %d: %s", resp.Error.Code, resp.Error.Message)
	}

	// Unmarshal result
	if result != nil && resp.Result != nil {
		if err := json.Unmarshal(resp.Result, result); err != nil {
			return fmt.Errorf("unmarshaling result: %w", err)
		}
	}

	return nil
}

func (c *lspClient) notify(method string, params any) error {
	req := jsonrpcRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	return c.writeMessage(req)
}

func (c *lspClient) close() error {
	c.stdin.Close()
	c.cmd.Wait()
	return nil
}

// startLSPServer starts the cem lsp subprocess and returns a client
func startLSPServer(ctx context.Context, workDir string) (*lspClient, error) {
	cmd := exec.CommandContext(ctx, cemBinary, "lsp")
	cmd.Dir = workDir
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+coverDir)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	// Capture stderr for debugging
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start cem lsp: %w", err)
	}

	client := &lspClient{
		stdin:   stdin,
		stdout:  bufio.NewReader(stdoutPipe),
		cmd:     cmd,
		pending: make(map[int64]chan jsonrpcResponse),
	}

	// Start response handler
	go client.handleResponses()

	return client, nil
}

// setupLSPTest creates a test workspace with the given fixture and initializes LSP
func setupLSPTest(t *testing.T, fixtureName string) (workDir string, client *lspClient, cleanup func()) {
	t.Helper()

	// Create temporary workspace
	tmpDir, err := os.MkdirTemp("", "cem-lsp-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Copy fixture to workspace
	workDir = filepath.Join(tmpDir, fixtureName)
	err = os.CopyFS(workDir, os.DirFS(filepath.Join(".", "fixture", fixtureName)))
	if err != nil {
		os.RemoveAll(tmpDir)
		t.Fatalf("Failed to copy fixture: %v", err)
	}

	// Start LSP server
	ctx := context.Background()
	client, err = startLSPServer(ctx, workDir)
	if err != nil {
		os.RemoveAll(tmpDir)
		t.Fatalf("Failed to start LSP server: %v", err)
	}

	// Initialize the server
	var initResult protocol.InitializeResult
	rootURI := protocol.DocumentUri("file://" + workDir)
	initParams := protocol.InitializeParams{
		RootURI: &rootURI,
		Capabilities: protocol.ClientCapabilities{
			TextDocument: &protocol.TextDocumentClientCapabilities{},
		},
	}

	if err := client.call("initialize", initParams, &initResult); err != nil {
		client.close()
		os.RemoveAll(tmpDir)
		t.Fatalf("Failed to initialize LSP server: %v", err)
	}

	// Send initialized notification
	if err := client.notify("initialized", protocol.InitializedParams{}); err != nil {
		client.close()
		os.RemoveAll(tmpDir)
		t.Fatalf("Failed to send initialized notification: %v", err)
	}

	cleanup = func() {
		// Try to shutdown gracefully
		client.call("shutdown", nil, nil)
		client.notify("exit", nil)
		client.close()
		os.RemoveAll(tmpDir)
	}

	return workDir, client, cleanup
}

// TestLSPInitializeShutdown tests the basic LSP lifecycle
func TestLSPInitializeShutdown(t *testing.T) {
	_, _, cleanup := setupLSPTest(t, "generate-project")
	defer cleanup()

	// Server is already initialized in setupLSPTest
	// The cleanup function will handle shutdown and exit
	// This test just verifies that initialization succeeded
	t.Log("LSP server initialized and ready for shutdown")
}

// TestLSPHover tests the hover functionality
func TestLSPHover(t *testing.T) {
	workDir, client, cleanup := setupLSPTest(t, "generate-project")
	defer cleanup()

	// Open the JavaScript file with didOpen notification
	jsFile := filepath.Join(workDir, "my-element.js")
	content, err := os.ReadFile(jsFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	didOpenParams := protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        protocol.DocumentUri("file://" + jsFile),
			LanguageID: "javascript",
			Version:    1,
			Text:       string(content),
		},
	}

	if err := client.notify("textDocument/didOpen", didOpenParams); err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}

	// Request hover on "MyElement" class name (line 4, character 13)
	// The file content is:
	//   Line 1: /**
	//   Line 2:  * @customElement my-element
	//   Line 3:  */
	//   Line 4: export class MyElement extends HTMLElement {}
	hoverParams := protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.DocumentUri("file://" + jsFile),
			},
			Position: protocol.Position{
				Line:      3, // 0-indexed, so line 4 in the file
				Character: 15, // Position on "MyElement" in "export class MyElement"
			},
		},
	}

	var hover *protocol.Hover
	if err := client.call("textDocument/hover", hoverParams, &hover); err != nil {
		t.Fatalf("Hover request failed: %v", err)
	}

	// Hover might be nil for simple class names - that's OK for this E2E test
	// The important part is that the LSP server handled the request without error
	t.Logf("Hover result: %+v", hover)
}

// TestLSPCompletion tests code completion
func TestLSPCompletion(t *testing.T) {
	workDir, client, cleanup := setupLSPTest(t, "generate-project")
	defer cleanup()

	// Open the JavaScript file
	jsFile := filepath.Join(workDir, "my-element.js")
	content, err := os.ReadFile(jsFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	didOpenParams := protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        protocol.DocumentUri("file://" + jsFile),
			LanguageID: "javascript",
			Version:    1,
			Text:       string(content),
		},
	}

	if err := client.notify("textDocument/didOpen", didOpenParams); err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}

	// Request completion at the end of the file
	completionParams := protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.DocumentUri("file://" + jsFile),
			},
			Position: protocol.Position{
				Line:      4, // After the class definition
				Character: 0,
			},
		},
	}

	var completions any
	if err := client.call("textDocument/completion", completionParams, &completions); err != nil {
		t.Fatalf("Completion request failed: %v", err)
	}

	t.Logf("Completion result: %+v", completions)
}

// TestLSPDidChange tests file change notifications
func TestLSPDidChange(t *testing.T) {
	workDir, client, cleanup := setupLSPTest(t, "generate-project")
	defer cleanup()

	// Open the JavaScript file
	jsFile := filepath.Join(workDir, "my-element.js")
	content, err := os.ReadFile(jsFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	didOpenParams := protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        protocol.DocumentUri("file://" + jsFile),
			LanguageID: "javascript",
			Version:    1,
			Text:       string(content),
		},
	}

	if err := client.notify("textDocument/didOpen", didOpenParams); err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}

	// Load the updated content from fixture
	updatedContentBytes, err := os.ReadFile(filepath.Join("fixture", "lsp-didchange", "my-element.updated.js"))
	if err != nil {
		t.Fatalf("Failed to read updated fixture: %v", err)
	}
	updatedContent := string(updatedContentBytes)

	didChangeParams := protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: protocol.DocumentUri("file://" + jsFile),
			},
			Version: 2,
		},
		ContentChanges: []interface{}{
			protocol.TextDocumentContentChangeEvent{
				Text: updatedContent,
			},
		},
	}

	if err := client.notify("textDocument/didChange", didChangeParams); err != nil {
		t.Fatalf("Failed to send didChange: %v", err)
	}

	// Request hover to verify the change was processed
	hoverParams := protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.DocumentUri("file://" + jsFile),
			},
			Position: protocol.Position{
				Line:      5, // On the observedAttributes line
				Character: 10,
			},
		},
	}

	var hover *protocol.Hover
	if err := client.call("textDocument/hover", hoverParams, &hover); err != nil {
		t.Fatalf("Hover after change failed: %v", err)
	}

	t.Logf("Hover after change: %+v", hover)
}
