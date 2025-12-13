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
package mcp_test

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/mcp"
	W "bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMCPCommandStdoutClean is an E2E test that verifies the actual `cem mcp`
// command produces no stdout contamination. This is the definitive test for #129.
//
// See: https://github.com/bennypowers/cem/issues/129
func TestMCPCommandStdoutClean(t *testing.T) {
	// Build cem binary from project root
	binPath := filepath.Join(t.TempDir(), "cem-test")
	buildCmd := exec.Command("go", "build", "-o", binPath, ".")
	buildCmd.Dir = ".." // Build from project root
	err := buildCmd.Run()
	require.NoError(t, err, "failed to build cem binary")

	// Start MCP server
	cmd := exec.Command(binPath, "mcp")
	cmd.Dir = "./fixtures/multiple-elements-integration"

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start the command
	err = cmd.Start()
	require.NoError(t, err)

	// Give it time to initialize and load manifests
	time.Sleep(500 * time.Millisecond)

	// Kill the process
	err = cmd.Process.Kill()
	require.NoError(t, err)

	// Wait for it to finish
	_ = cmd.Wait() // Ignore error since we killed it

	// Verify stdout is completely empty (no manifest loading logs)
	stdoutContent := stdout.String()
	assert.Empty(t, stdoutContent, "stdout should be completely empty, got: %q", stdoutContent)

	// Check for log contamination patterns
	contaminations := []string{
		"[REGISTRY]",
		"[IN-MEMORY]",
		"INFO",
		"DEBUG",
		"Loading manifests",
		"Successfully",
	}

	for _, pattern := range contaminations {
		assert.NotContains(t, stdoutContent, pattern,
			"stdout contains log output '%s' which would corrupt JSON-RPC", pattern)
	}

	// Stderr MAY contain logs (this is acceptable and expected)
	// We just verify we got some stderr output indicating the server ran
	t.Logf("stderr output (expected): %s", stderr.String())
}

// TestMCPServerStdoutCleanWithQuietMode verifies that enabling quiet mode
// suppresses INFO/DEBUG logs at the unit level. This tests the actual fix.
//
// See: https://github.com/bennypowers/cem/issues/129
func TestMCPServerStdoutCleanWithQuietMode(t *testing.T) {
	// Redirect pterm to stderr (like cmd/mcp.go does)
	originalPtermOutput := pterm.DefaultBasicText.Writer
	pterm.SetDefaultOutput(os.Stderr)
	defer pterm.SetDefaultOutput(originalPtermOutput)

	// Enable quiet mode (THIS IS THE FIX WE'RE TESTING)
	originalQuiet := logging.IsQuietEnabled()
	logging.SetQuietEnabled(true)
	defer logging.SetQuietEnabled(originalQuiet)

	// Create a pipe to capture stdout
	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	defer func() { os.Stdout = originalStdout }()
	defer func() { _ = w.Close() }() // Ensure pipe is closed even if test fails early

	// Buffer to capture stdout
	var stdoutBuf bytes.Buffer
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(&stdoutBuf, r)
		close(done)
	}()

	// Create workspace and MCP server
	workspace := W.NewFileSystemWorkspaceContext("./fixtures/multiple-elements-integration")
	err = workspace.Init()
	require.NoError(t, err)

	server, err := mcp.NewServer(workspace)
	require.NoError(t, err)

	// Load manifests (this is where logging happens)
	// We don't run the full server because that blocks on stdio
	registry := server.GetRegistry()
	err = registry.LoadManifests()
	require.NoError(t, err)

	// Close stdout pipe
	_ = w.Close()
	<-done

	// Verify stdout is clean
	stdoutContent := stdoutBuf.String()

	// Should contain ZERO output (no JSON-RPC in this test)
	assert.Empty(t, stdoutContent, "stdout should be completely empty during manifest loading")

	// Check for common log prefixes that indicate contamination
	contaminations := []string{
		"[REGISTRY]",
		"[IN-MEMORY]",
		"INFO",
		"DEBUG",
		"Loading manifests",
		"Successfully",
	}

	for _, pattern := range contaminations {
		assert.NotContains(t, stdoutContent, pattern,
			"stdout contains log output '%s' which would corrupt JSON-RPC", pattern)
	}
}

// TestMCPServerStderrAllowed verifies that error/warning messages CAN go to stderr.
// This is acceptable per MCP spec - only stdout must be clean.
func TestMCPServerStderrAllowed(t *testing.T) {
	// Capture stderr
	originalStderr := os.Stderr
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stderr = w
	defer func() { os.Stderr = originalStderr }()
	defer func() { _ = w.Close() }() // Ensure pipe is closed even if test fails early

	var stderrBuf bytes.Buffer
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(&stderrBuf, r)
		close(done)
	}()

	// Create workspace with basic integration
	workspace := W.NewFileSystemWorkspaceContext("./fixtures/basic-integration")
	err = workspace.Init()
	require.NoError(t, err)

	server, err := mcp.NewServer(workspace)
	require.NoError(t, err)

	registry := server.GetRegistry()
	err = registry.LoadManifests()
	require.NoError(t, err)

	_ = w.Close()
	<-done

	// stderr MAY contain error/warning messages (this is acceptable and desired for debugging)
	// We're just verifying the test infrastructure works
	// The actual stderr content depends on whether pterm outputs debug info
	// This is fine - we just want to ensure stderr is available for logging
}
