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
package hover_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type HoverTestConfig struct {
	Tests []struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		InputFile       string `json:"inputFile"`
		ElementPosition struct {
			Line      uint32 `json:"line"`
			Character uint32 `json:"character"`
		} `json:"elementPosition"`
		ElementMatch struct {
			TagName string `json:"tagName"`
			Range   struct {
				Start struct {
					Line      uint32 `json:"line"`
					Character uint32 `json:"character"`
				} `json:"start"`
				End struct {
					Line      uint32 `json:"line"`
					Character uint32 `json:"character"`
				} `json:"end"`
			} `json:"range"`
			Attributes map[string]struct {
				Name  string `json:"name"`
				Value string `json:"value"`
				Range struct {
					Start struct {
						Line      uint32 `json:"line"`
						Character uint32 `json:"character"`
					} `json:"start"`
					End struct {
						Line      uint32 `json:"line"`
						Character uint32 `json:"character"`
					} `json:"end"`
				} `json:"range"`
			} `json:"attributes,omitempty"`
		} `json:"elementMatch"`
		AttributePosition *struct {
			Line      uint32 `json:"line"`
			Character uint32 `json:"character"`
		} `json:"attributePosition,omitempty"`
	} `json:"tests"`
}

func TestHoverIntegrationWithDocumentChanges(t *testing.T) {
	fixturePath := filepath.Join("fixtures", "hover-integration")

	ctx := testhelpers.NewMockServerContext()

	// Load the manifest manually for testing
	manifestPath := filepath.Join("..", "..", "..", "test", "fixtures", "hover-integration", "custom-elements.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	ctx.AddManifest(&pkg)

	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Load test configuration
	configPath := filepath.Join(fixturePath, "test-config.json")
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read test config: %v", err)
	}

	var config HoverTestConfig
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		t.Fatalf("Failed to parse test config: %v", err)
	}

	for _, testCase := range config.Tests {
		t.Run(testCase.Description, func(t *testing.T) {
			// Load test fixture content
			contentPath := filepath.Join(fixturePath, testCase.InputFile)
			content, err := os.ReadFile(contentPath)
			if err != nil {
				t.Fatalf("Failed to read test fixture %s: %v", testCase.InputFile, err)
			}

			uri := "file://" + contentPath

			// Create element match from config
			elementMatch := types.CustomElementMatch{
				TagName: testCase.ElementMatch.TagName,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      testCase.ElementMatch.Range.Start.Line,
						Character: testCase.ElementMatch.Range.Start.Character,
					},
					End: protocol.Position{
						Line:      testCase.ElementMatch.Range.End.Line,
						Character: testCase.ElementMatch.Range.End.Character,
					},
				},
			}

			// Add attributes if present
			if len(testCase.ElementMatch.Attributes) > 0 {
				elementMatch.Attributes = make(map[string]types.AttributeMatch)
				for attrName, attrData := range testCase.ElementMatch.Attributes {
					elementMatch.Attributes[attrName] = types.AttributeMatch{
						Name:  attrData.Name,
						Value: attrData.Value,
						Range: protocol.Range{
							Start: protocol.Position{
								Line:      attrData.Range.Start.Line,
								Character: attrData.Range.Start.Character,
							},
							End: protocol.Position{
								Line:      attrData.Range.End.Line,
								Character: attrData.Range.End.Character,
							},
						},
					}
				}
			}

			// Add document to the mock context using DocumentManager
			dm, err := lsp.NewDocumentManager()
			if err != nil {
				t.Fatalf("Failed to create DocumentManager: %v", err)
			}
			defer dm.Close()
			ctx.SetDocumentManager(dm)

			doc := dm.OpenDocument(uri, string(content), 1)
			ctx.AddDocument(uri, doc)

			// Test element hover
			pos := protocol.Position{
				Line:      testCase.ElementPosition.Line,
				Character: testCase.ElementPosition.Character,
			}
			params := &protocol.HoverParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{URI: uri},
					Position:     pos,
				},
			}

			result, err := hover.Hover(ctx, nil, params)
			if err != nil {
				t.Fatalf("Hover failed: %v", err)
			}

			if result == nil {
				t.Fatal("Expected hover result, got nil")
			}

			hoverContent := result.Contents.(protocol.MarkupContent)
			if hoverContent.Value == "" {
				t.Error("Expected hover content, got empty string")
			}

			if !strings.Contains(hoverContent.Value, testCase.ElementMatch.TagName) {
				t.Errorf("Expected hover content to contain '%s', got: %s", testCase.ElementMatch.TagName, hoverContent.Value)
			}

			// Test attribute hover if position specified
			if testCase.AttributePosition != nil {
				attrPos := protocol.Position{
					Line:      testCase.AttributePosition.Line,
					Character: testCase.AttributePosition.Character,
				}
				attrParams := &protocol.HoverParams{
					TextDocumentPositionParams: protocol.TextDocumentPositionParams{
						TextDocument: protocol.TextDocumentIdentifier{URI: uri},
						Position:     attrPos,
					},
				}

				attrResult, err := hover.Hover(ctx, nil, attrParams)
				if err != nil {
					t.Fatalf("Hover on attribute failed: %v", err)
				}

				if attrResult == nil {
					t.Fatal("Expected hover result on attribute, got nil")
				}

				attrContent := attrResult.Contents.(protocol.MarkupContent)
				if attrContent.Value == "" {
					t.Error("Expected hover content on attribute, got empty string")
				}

				if !strings.Contains(attrContent.Value, "test-attr") {
					t.Errorf("Expected hover content to contain 'test-attr', got: %s", attrContent.Value)
				}
			}
		})
	}

}

func TestHoverWithTypeScriptTemplates(t *testing.T) {
	fixturePath := filepath.Join("fixtures", "typescript-templates")

	ctx := testhelpers.NewMockServerContext()

	// Load the manifest manually for testing
	manifestPath := filepath.Join("test", "fixtures", "typescript-templates", "manifest.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	ctx.AddManifest(&pkg)

	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Load test configuration
	configPath := filepath.Join(fixturePath, "test-config.json")
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read test config: %v", err)
	}

	var config HoverTestConfig
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		t.Fatalf("Failed to parse test config: %v", err)
	}

	for _, testCase := range config.Tests {
		t.Run(testCase.Description, func(t *testing.T) {
			// Load test fixture content
			contentPath := filepath.Join(fixturePath, testCase.InputFile)
			content, err := os.ReadFile(contentPath)
			if err != nil {
				t.Fatalf("Failed to read test fixture %s: %v", testCase.InputFile, err)
			}

			uri := "file:///test.ts"

			// Add document to the mock context using DocumentManager
			dm, err := lsp.NewDocumentManager()
			if err != nil {
				t.Fatalf("Failed to create DocumentManager: %v", err)
			}
			defer dm.Close()
			ctx.SetDocumentManager(dm)

			doc := dm.OpenDocument(uri, string(content), 1)
			ctx.AddDocument(uri, doc)

			// Test hover
			pos := protocol.Position{
				Line:      testCase.ElementPosition.Line,
				Character: testCase.ElementPosition.Character,
			}
			params := &protocol.HoverParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{URI: uri},
					Position:     pos,
				},
			}

			result, err := hover.Hover(ctx, nil, params)
			if err != nil {
				t.Fatalf("Hover failed: %v", err)
			}

			if result == nil {
				t.Fatal("Expected hover result, got nil")
			}

			hoverContent := result.Contents.(protocol.MarkupContent)
			if !strings.Contains(hoverContent.Value, testCase.ElementMatch.TagName) {
				t.Errorf("Expected hover content to contain '%s', got: %s", testCase.ElementMatch.TagName, hoverContent.Value)
			}
		})
	}
}
