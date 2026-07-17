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
	"bytes"
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

	"bennypowers.dev/cem/internal/platform/testutil"
	"go.lsp.dev/protocol"
	urilib "go.lsp.dev/uri"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	stdin     io.WriteCloser
	stdout    *bufio.Reader
	cmd       *exec.Cmd
	mu        sync.Mutex
	nextID    atomic.Int64
	pending   map[int64]chan jsonrpcResponse
	pendingMu sync.Mutex
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

	if result != nil && resp.Result != nil {
		if err := protocol.Unmarshal(resp.Result, result); err != nil {
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
	return setupLSPTestFromDir(t, filepath.Join("testdata", "fixtures", fixtureName), fixtureName)
}

// setupLSPTestFromExample creates a test workspace from an examples/ directory
func setupLSPTestFromExample(t *testing.T, exampleName string) (workDir string, client *lspClient, cleanup func()) {
	t.Helper()
	return setupLSPTestFromDir(t, filepath.Join("..", "examples", exampleName), exampleName)
}

// setupLSPTestFromDir copies srcDir into a temp directory and starts the LSP server there
func setupLSPTestFromDir(t *testing.T, srcDir, name string) (workDir string, client *lspClient, cleanup func()) {
	t.Helper()

	// Create temporary workspace
	tmpDir, err := os.MkdirTemp("", "cem-lsp-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Copy fixture to workspace
	workDir = filepath.Join(tmpDir, name)
	err = os.CopyFS(workDir, os.DirFS(srcDir))
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
	rootURI := urilib.URI("file://" + workDir)
	var initParams protocol.InitializeParams
	initParams.RootURI = &rootURI

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
			URI:        urilib.URI("file://" + jsFile),
			LanguageID: "javascript",
			Version:    1,
			Text:       string(content),
		},
	}

	if err := client.notify("textDocument/didOpen", didOpenParams); err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}

	// Request hover on "MyElement" class name (line 4, character 15)
	// The file content is:
	//   Line 1: /**
	//   Line 2:  * @customElement my-element
	//   Line 3:  */
	//   Line 4: export class MyElement extends HTMLElement {}
	hoverParams := protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: urilib.URI("file://" + jsFile),
			},
			Position: protocol.Position{
				Line:      3,  // 0-indexed, so line 4 in the file
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
			URI:        urilib.URI("file://" + jsFile),
			LanguageID: "javascript",
			Version:    1,
			Text:       string(content),
		},
	}

	if err := client.notify("textDocument/didOpen", didOpenParams); err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}

	// Compute safe EOF position from file contents
	lines := strings.Split(string(content), "\n")
	eofLine := len(lines) - 1
	if eofLine < 0 {
		eofLine = 0
	}
	eofChar := len(lines[eofLine])

	// Request completion at the end of the file
	completionParams := protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: urilib.URI("file://" + jsFile),
			},
			Position: protocol.Position{
				Line:      uint32(eofLine),
				Character: uint32(eofChar),
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
			URI:        urilib.URI("file://" + jsFile),
			LanguageID: "javascript",
			Version:    1,
			Text:       string(content),
		},
	}

	if err := client.notify("textDocument/didOpen", didOpenParams); err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}

	// Load the updated content from fixture
	updatedContentBytes, err := os.ReadFile(filepath.Join("testdata", "fixtures", "lsp-didchange", "my-element.updated.js"))
	if err != nil {
		t.Fatalf("Failed to read updated fixture: %v", err)
	}
	updatedContent := string(updatedContentBytes)

	didChangeParams := protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: urilib.URI("file://" + jsFile),
			},
			Version: 2,
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			&protocol.TextDocumentContentChangeWholeDocument{
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
				URI: urilib.URI("file://" + jsFile),
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

// openHTMLFile sends a didOpen notification for the given HTML file and returns its URI and content
func openHTMLFile(t *testing.T, client *lspClient, filePath string) (uri urilib.URI, content string) {
	t.Helper()
	raw, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", filePath, err)
	}
	content = string(raw)
	uri = urilib.URI("file://" + filePath)
	err = client.notify("textDocument/didOpen", protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        uri,
			LanguageID: "html",
			Version:    1,
			Text:       content,
		},
	})
	if err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}
	return uri, content
}

// openHTMLFixture reads a fixture file from cmd/testdata/fixtures, parses and
// strips the ^cursor marker using testutil.ExtractHTMLCursor, copies the
// cleaned content into workDir, and sends a didOpen notification.
func openHTMLFixture(t *testing.T, client *lspClient, workDir, fixturePath string) (uri urilib.URI, cursor protocol.Position) {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("testdata", "fixtures", fixturePath))
	if err != nil {
		t.Fatalf("Failed to read fixture %s: %v", fixturePath, err)
	}
	cleaned, pos := testutil.ExtractHTMLCursor(string(raw))
	if pos == nil {
		t.Fatalf("no ^cursor marker found in fixture %s", fixturePath)
	}
	cursor = *pos

	destPath := filepath.Join(workDir, filepath.Base(fixturePath))
	if err := os.WriteFile(destPath, []byte(cleaned), 0644); err != nil {
		t.Fatalf("Failed to write fixture to workspace: %v", err)
	}

	uri = urilib.URI("file://" + destPath)
	err = client.notify("textDocument/didOpen", protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        uri,
			LanguageID: "html",
			Version:    1,
			Text:       cleaned,
		},
	})
	if err != nil {
		t.Fatalf("Failed to send didOpen: %v", err)
	}
	return uri, cursor
}

// normalizeJSON replaces workspace-specific paths with $WORKSPACE for golden comparison
func normalizeJSON(data json.RawMessage, workDir string) []byte {
	s := string(data)
	s = strings.ReplaceAll(s, workDir, "$WORKSPACE")
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, []byte(s), "", "  "); err != nil {
		return []byte(s)
	}
	return append(pretty.Bytes(), '\n')
}

// checkLSPGolden compares an LSP response against a golden file
func checkLSPGolden(t *testing.T, name string, data json.RawMessage, workDir string) {
	t.Helper()
	normalized := normalizeJSON(data, workDir)
	goldenPath := filepath.Join("testdata", "goldens", "lsp", name+".json")
	if *testutil.Update {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
			t.Fatalf("failed to create golden directory: %v", err)
		}
		if err := os.WriteFile(goldenPath, normalized, 0644); err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		t.Logf("Updated golden file: %s", goldenPath)
		return
	}
	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("golden file missing: %s (run with -update)\nerror: %v", goldenPath, err)
	}
	if string(normalized) != string(expected) {
		t.Errorf("Output differs from golden file %s.\n\nExpected:\n%s\nGot:\n%s",
			goldenPath, string(expected), string(normalized))
		t.Log("Run 'make test-e2e UPDATE=-update' to update golden files")
	}
}

// TestLSPDiagnostic tests pull diagnostics via textDocument/diagnostic
func TestLSPDiagnostic(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	// Use the errors fixture which has intentional typos to produce diagnostics
	htmlFile := filepath.Join(workDir, "elements", "demo-button", "demo", "incorrect-usage.html")
	uri, _ := openHTMLFile(t, client, htmlFile)

	var result json.RawMessage
	err := client.call("textDocument/diagnostic", protocol.DocumentDiagnosticParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	}, &result)
	require.NoError(t, err, "textDocument/diagnostic must not error")
	require.NotNil(t, result, "diagnostic result must not be nil (nil indicates adapter layer returned wrong type)")

	var report struct {
		Kind  string `json:"kind"`
		Items []any  `json:"items"`
	}
	require.NoError(t, json.Unmarshal(result, &report))
	assert.Equal(t, "full", report.Kind)
	assert.NotEmpty(t, report.Items, "should produce diagnostics for typo attributes")

	checkLSPGolden(t, "diagnostic", result, workDir)
}

// TestLSPDefinition tests go-to-definition on a custom element tag
func TestLSPDefinition(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	uri, cursor := openHTMLFixture(t, client, workDir, "lsp-positions/definition-target.html")

	var result json.RawMessage
	err := client.call("textDocument/definition", protocol.DefinitionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri},
			Position:     cursor,
		},
	}, &result)
	require.NoError(t, err, "textDocument/definition must not error")
	require.NotNil(t, result, "definition result must not be nil (nil indicates adapter layer returned wrong type)")

	var location protocol.Location
	require.NoError(t, json.Unmarshal(result, &location))
	assert.Contains(t, string(location.URI), "demo-button", "definition should point to demo-button source")

	checkLSPGolden(t, "definition", result, workDir)
}

// TestLSPReferences tests find-references on a custom element tag
func TestLSPReferences(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	uri, cursor := openHTMLFixture(t, client, workDir, "lsp-positions/definition-target.html")

	var result json.RawMessage
	err := client.call("textDocument/references", protocol.ReferenceParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri},
			Position:     cursor,
		},
		Context: protocol.ReferenceContext{IncludeDeclaration: true},
	}, &result)
	require.NoError(t, err, "textDocument/references must not error")
	require.NotNil(t, result, "references result must not be nil (nil indicates adapter layer returned wrong type)")

	var locations []protocol.Location
	require.NoError(t, json.Unmarshal(result, &locations))
	assert.NotEmpty(t, locations, "should find at least one reference to demo-button")

	checkLSPGolden(t, "references", result, workDir)
}

// TestLSPCodeAction tests code actions driven by diagnostics
func TestLSPCodeAction(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	// Use the errors fixture which has intentional typos that produce fixable diagnostics
	htmlFile := filepath.Join(workDir, "elements", "demo-button", "demo", "incorrect-usage.html")
	uri, _ := openHTMLFile(t, client, htmlFile)

	// Get diagnostics first
	var diagResult json.RawMessage
	err := client.call("textDocument/diagnostic", protocol.DocumentDiagnosticParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	}, &diagResult)
	require.NoError(t, err, "diagnostic request should succeed")

	var report struct {
		Kind  string                `json:"kind"`
		Items []protocol.Diagnostic `json:"items"`
	}
	require.NoError(t, protocol.Unmarshal(diagResult, &report))

	var actions json.RawMessage
	err = client.call("textDocument/codeAction", protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   protocol.Position{Line: 99, Character: 0},
		},
		Context: protocol.CodeActionContext{
			Diagnostics: report.Items,
		},
	}, &actions)
	require.NoError(t, err, "textDocument/codeAction must not error")
	require.NotNil(t, actions, "codeAction result must not be nil (nil indicates adapter layer returned wrong type)")

	var actionList []json.RawMessage
	require.NoError(t, json.Unmarshal(actions, &actionList))
	assert.NotEmpty(t, actionList, "should produce code actions for typo diagnostics")

	checkLSPGolden(t, "code-action", actions, workDir)
}

// TestLSPInlayHint tests inlay hints on custom element attributes
func TestLSPInlayHint(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	htmlFile := filepath.Join(workDir, "elements", "demo-button", "demo", "variants.html")
	uri, _ := openHTMLFile(t, client, htmlFile)

	var result json.RawMessage
	err := client.call("textDocument/inlayHint", protocol.InlayHintParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   protocol.Position{Line: 56, Character: 0},
		},
	}, &result)
	require.NoError(t, err, "textDocument/inlayHint must not error")
	require.NotNil(t, result, "inlayHint result must not be nil (nil indicates adapter layer returned wrong type)")

	var hints []json.RawMessage
	require.NoError(t, json.Unmarshal(result, &hints))
	assert.NotEmpty(t, hints, "should produce inlay hints for demo-button attributes")

	checkLSPGolden(t, "inlay-hint", result, workDir)
}

// TestLSPWorkspaceSymbol tests workspace symbol search
func TestLSPWorkspaceSymbol(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	var result json.RawMessage
	err := client.call("workspace/symbol", protocol.WorkspaceSymbolParams{
		Query: "demo-button",
	}, &result)
	require.NoError(t, err, "workspace/symbol must not error")
	require.NotNil(t, result, "workspace/symbol result must not be nil (nil indicates adapter layer returned wrong type)")

	var symbols []protocol.SymbolInformation
	require.NoError(t, json.Unmarshal(result, &symbols))
	assert.NotEmpty(t, symbols, "should find demo-button symbol")

	found := false
	for _, sym := range symbols {
		if strings.Contains(sym.Name, "demo-button") {
			found = true
			break
		}
	}
	assert.True(t, found, "should contain a symbol matching 'demo-button'")

	checkLSPGolden(t, "workspace-symbol", result, workDir)
}

// TestLSPDidClose tests document close notification clears in-memory state.
// Proves didClose is not a no-op: modified content produces diagnostics,
// closing the document clears them, and inlay hints return nil.
func TestLSPDidClose(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	// Open the incorrect-usage file which has typo attributes -- produces diagnostics
	htmlFile := filepath.Join(workDir, "elements", "demo-button", "demo", "incorrect-usage.html")
	uri, _ := openHTMLFile(t, client, htmlFile)

	// Diagnostics on invalid content should be non-empty (proves document is tracked)
	var dirtyDiag json.RawMessage
	err := client.call("textDocument/diagnostic", protocol.DocumentDiagnosticParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	}, &dirtyDiag)
	require.NoError(t, err)
	var dirtyReport struct {
		Kind  string `json:"kind"`
		Items []any  `json:"items"`
	}
	require.NoError(t, json.Unmarshal(dirtyDiag, &dirtyReport))
	require.NotEmpty(t, dirtyReport.Items, "incorrect-usage content must produce diagnostics before close")

	// Close the document -- should clear in-memory state
	err = client.notify("textDocument/didClose", protocol.DidCloseTextDocumentParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
	require.NoError(t, err, "textDocument/didClose notification must not error")

	// Diagnostics after close: document is untracked, handler returns empty
	var closedDiag json.RawMessage
	err = client.call("textDocument/diagnostic", protocol.DocumentDiagnosticParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	}, &closedDiag)
	require.NoError(t, err, "diagnostic request must succeed after didClose")
	require.NotNil(t, closedDiag, "diagnostic result after didClose must not be nil")
	var closedReport struct {
		Kind  string `json:"kind"`
		Items []any  `json:"items"`
	}
	require.NoError(t, json.Unmarshal(closedDiag, &closedReport))
	assert.Empty(t, closedReport.Items, "diagnostics must be empty after didClose clears modified state")

	// Inlay hints after close: document is untracked, handler returns nil
	var hints json.RawMessage
	err = client.call("textDocument/inlayHint", protocol.InlayHintParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   protocol.Position{Line: 10, Character: 0},
		},
	}, &hints)
	require.NoError(t, err, "inlayHint must succeed after didClose")
	assert.Equal(t, "null", string(hints), "inlay hints must be nil after didClose")

	checkLSPGolden(t, "did-close-diagnostic", closedDiag, workDir)
}

// TestLSPDidChangeConfiguration tests configuration change notification
func TestLSPDidChangeConfiguration(t *testing.T) {
	workDir, client, cleanup := setupLSPTestFromExample(t, "kitchen-sink")
	defer cleanup()

	htmlFile := filepath.Join(workDir, "elements", "demo-button", "demo", "variants.html")
	uri, _ := openHTMLFile(t, client, htmlFile)

	settingsJSON, _ := json.Marshal(map[string]any{
		"cem": map[string]any{
			"inlayHints": false,
		},
	})
	err := client.notify("workspace/didChangeConfiguration", protocol.DidChangeConfigurationParams{
		Settings: settingsJSON,
	})
	require.NoError(t, err, "workspace/didChangeConfiguration must not error")

	// Verify inlay hints are now empty after disabling
	var result json.RawMessage
	err = client.call("textDocument/inlayHint", protocol.InlayHintParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   protocol.Position{Line: 56, Character: 0},
		},
	}, &result)
	require.NoError(t, err, "inlayHint after config change must not error")

	checkLSPGolden(t, "inlay-hint-disabled", result, workDir)
}

// TestLSPCommandStdoutClean verifies the `cem lsp` command produces only
// valid JSON-RPC framed content on stdout, with no stray log output.
func TestLSPCommandStdoutClean(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, cemBinary, "lsp")
	cmd.Dir = filepath.Join("testdata", "fixtures", "generate-project")
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+coverDir)

	var stdout, stderr bytes.Buffer
	stdinPipe, err := cmd.StdinPipe()
	require.NoError(t, err)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	require.NoError(t, cmd.Start())

	initReq := `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"capabilities":{}}}`
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(initReq))
	_, err = stdinPipe.Write([]byte(header + initReq))
	require.NoError(t, err)

	// Close stdin to signal the server to exit, then wait for process
	// to finish. All stdout is captured after the process exits,
	// avoiding concurrent buffer access.
	require.NoError(t, stdinPipe.Close())
	if err := cmd.Wait(); err != nil {
		if ctx.Err() == nil {
			t.Fatalf("LSP process failed: %v\nstderr: %s", err, stderr.String())
		}
	}

	out := stdout.String()
	t.Logf("stdout length: %d bytes", len(out))
	t.Logf("stderr (expected): %s", stderr.String())

	contaminations := []string{
		"[REGISTRY]",
		"[IN-MEMORY]",
		"INFO",
		"DEBUG",
		"Loading manifests",
		"Successfully",
	}
	for _, pattern := range contaminations {
		assert.NotContains(t, out, pattern,
			"stdout contains log output %q which would corrupt JSON-RPC", pattern)
	}

	// Verify all stdout bytes parse as JSON-RPC framed messages
	remaining := out
	for len(remaining) > 0 {
		idx := strings.Index(remaining, "Content-Length:")
		if idx < 0 {
			if strings.TrimSpace(remaining) != "" {
				t.Errorf("trailing non-JSON-RPC content: %q", remaining)
			}
			break
		}
		if idx != 0 {
			nonRPC := remaining[:idx]
			if strings.TrimSpace(nonRPC) != "" {
				t.Errorf("non-JSON-RPC content before header: %q", nonRPC)
			}
			remaining = remaining[idx:]
		}
		headerEnd := strings.Index(remaining, "\r\n\r\n")
		if headerEnd < 0 {
			t.Errorf("incomplete JSON-RPC header: %q", remaining)
			break
		}
		clLine := remaining[:headerEnd]
		parts := strings.SplitN(clLine, ":", 2)
		if len(parts) != 2 {
			t.Errorf("malformed Content-Length header: %q", clLine)
			break
		}
		cl, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			t.Errorf("invalid Content-Length value: %v", err)
			break
		}
		bodyStart := headerEnd + 4
		if bodyStart+cl > len(remaining) {
			t.Errorf("body truncated: need %d bytes, have %d", cl, len(remaining)-bodyStart)
			break
		}
		body := remaining[bodyStart : bodyStart+cl]
		if !json.Valid([]byte(body)) {
			t.Errorf("JSON-RPC body is not valid JSON: %q", body[:min(100, len(body))])
		}
		remaining = remaining[bodyStart+cl:]
	}
}
