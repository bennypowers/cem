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
package codeAction_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/codeAction"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestConfig represents the configuration for a missing import test case
type TestConfig struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	InputFile    string `json:"inputFile"`
	GoldenFile   string `json:"goldenFile"`
	DocumentURI  string `json:"documentURI"`
	ImportPath   string `json:"importPath"`
	TagName      string `json:"tagName"`
	ExpectAction bool   `json:"expectAction"`
}

func TestCreateMissingImportAction(t *testing.T) {
	fixturesPath := filepath.Join("testdata", "missing-import")

	// Read test configuration
	configPath := filepath.Join(fixturesPath, "test-config.json")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read test config: %v", err)
	}

	var testConfigs []TestConfig
	err = json.Unmarshal(configData, &testConfigs)
	if err != nil {
		t.Fatalf("Failed to parse test config: %v", err)
	}

	for _, config := range testConfigs {
		t.Run(config.Name, func(t *testing.T) {
			// Read input file
			inputPath := filepath.Join(fixturesPath, config.InputFile)
			inputContent, err := os.ReadFile(inputPath)
			if err != nil {
				t.Fatalf("Failed to read input file %s: %v", inputPath, err)
			}

			// Read expected output (golden file)
			goldenPath := filepath.Join(fixturesPath, config.GoldenFile)
			expectedContent, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("Failed to read golden file %s: %v", goldenPath, err)
			}

			// Create mock context with the document
			mockCtx := testhelpers.NewMockServerContext()
			dm, err := document.NewDocumentManager()
			if err != nil {
				t.Fatalf("Failed to create DocumentManager: %v", err)
			}
			defer dm.Close()
			mockCtx.SetDocumentManager(dm)

			// Create document from input content using DocumentManager
			doc := dm.OpenDocument(config.DocumentURI, string(inputContent), 1)
			mockCtx.AddDocument(config.DocumentURI, doc)

			// Create diagnostic data for missing import
			diagnosticData := map[string]any{
				"type":       string(types.DiagnosticTypeMissingImport),
				"original":   config.TagName,
				"suggestion": "import '" + config.ImportPath + "'",
				"importPath": config.ImportPath,
				"tagName":    config.TagName,
				"range": map[string]any{
					"start": map[string]any{"line": float64(5), "character": float64(10)},
					"end":   map[string]any{"line": float64(5), "character": float64(20)},
				},
			}

			diagnostic := &protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{Line: 5, Character: 10},
					End:   protocol.Position{Line: 5, Character: 20},
				},
				Message: "Test diagnostic",
			}

			// Call the function under test
			action, err := codeAction.CreateMissingImportAction(mockCtx, diagnostic, diagnosticData, config.DocumentURI)
			if err != nil {
				t.Fatalf("createMissingImportAction failed: %v", err)
			}

			if !config.ExpectAction {
				if action != nil {
					t.Errorf("Expected no action, but got: %+v", action)
				}
				return
			}

			// Verify action was created
			if action == nil {
				t.Fatal("Expected action to be created, but got nil")
			}

			// Verify action properties
			expectedTitle := "Add import for '" + config.TagName + "'"
			if action.Title != expectedTitle {
				t.Errorf("Expected title '%s', got '%s'", expectedTitle, action.Title)
			}

			if action.Kind == nil || *action.Kind != protocol.CodeActionKindQuickFix {
				t.Errorf("Expected kind 'quickfix', got %v", action.Kind)
			}

			if action.Edit == nil || action.Edit.Changes == nil {
				t.Fatal("Expected edit with changes, got nil")
			}

			changes, exists := action.Edit.Changes[config.DocumentURI]
			if !exists {
				t.Fatal("Expected changes for document URI")
			}

			if len(changes) != 1 {
				t.Fatalf("Expected 1 text edit, got %d", len(changes))
			}

			// Apply the text edit to the original content
			actualContent := applyTextEdit(string(inputContent), changes[0])
			expectedContentStr := strings.TrimSpace(string(expectedContent))
			actualContentStr := strings.TrimSpace(actualContent)

			if actualContentStr != expectedContentStr {
				t.Errorf("Content mismatch for %s:\nExpected:\n%s\n\nActual:\n%s", config.Name, expectedContentStr, actualContentStr)
			}

			// Verify diagnostic is attached
			if len(action.Diagnostics) != 1 {
				t.Errorf("Expected 1 diagnostic, got %d", len(action.Diagnostics))
			}

			t.Logf("✓ %s: %s", config.Name, config.Description)
		})
	}
}

// TestCreateMissingImportActionErrors tests error cases
func TestCreateMissingImportActionErrors(t *testing.T) {
	tests := []struct {
		name         string
		data         map[string]any
		expectAction bool
	}{
		{
			name: "Wrong diagnostic type",
			data: map[string]any{
				"type":       string(types.DiagnosticTypeTagSuggestion),
				"original":   "my-element",
				"suggestion": "my-other-element",
			},
			expectAction: false,
		},
		{
			name: "Invalid data structure",
			data: map[string]any{
				"type": "invalid-type",
			},
			expectAction: false,
		},
		{
			name:         "Empty data",
			data:         map[string]any{},
			expectAction: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagnostic := &protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{Line: 5, Character: 10},
					End:   protocol.Position{Line: 5, Character: 20},
				},
				Message: "Test diagnostic",
			}

			mockCtx := testhelpers.NewMockServerContext()
			documentURI := "file:///test.ts"

			action, err := codeAction.CreateMissingImportAction(mockCtx, diagnostic, tt.data, documentURI)
			if err != nil {
				t.Fatalf("%v", err)
			}

			if tt.expectAction {
				if action == nil {
					t.Fatal("Expected action to be created, but got nil")
				}
			} else {
				if action != nil {
					t.Errorf("Expected no action, but got: %+v", action)
				}
			}
		})
	}
}

// applyTextEdit applies a single text edit to content and returns the result
func applyTextEdit(content string, edit protocol.TextEdit) string {
	lines := strings.Split(content, "\n")

	// Insert at the specified position
	insertLine := int(edit.Range.Start.Line)
	insertChar := int(edit.Range.Start.Character)

	if insertLine >= len(lines) {
		// Insert at end of file
		return content + edit.NewText
	}

	if insertLine == 0 && insertChar == 0 {
		// Insert at beginning of file
		return edit.NewText + content
	}

	// Insert at specific line/character
	line := lines[insertLine]
	if insertChar > len(line) {
		insertChar = len(line)
	}

	newLine := line[:insertChar] + edit.NewText + line[insertChar:]
	lines[insertLine] = newLine

	return strings.Join(lines, "\n")
}
