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
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestIndentationDetection tests the indentation detection functionality using fixtures
func TestIndentationDetection(t *testing.T) {
	fixturesPath := filepath.Join("fixtures", "indentation-detection")

	// Get all HTML fixture files
	entries, err := os.ReadDir(fixturesPath)
	if err != nil {
		t.Fatalf("Failed to read fixtures directory: %v", err)
	}

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".html") {
			continue
		}

		fixtureName := strings.TrimSuffix(entry.Name(), ".html")
		t.Run(fixtureName, func(t *testing.T) {
			// Read the HTML fixture
			htmlPath := filepath.Join(fixturesPath, entry.Name())
			htmlContent, err := os.ReadFile(htmlPath)
			if err != nil {
				t.Fatalf("Failed to read HTML fixture %s: %v", htmlPath, err)
			}

			// Read the golden file
			goldenPath := filepath.Join(fixturesPath, fixtureName+".golden")
			goldenContent, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("Failed to read golden file %s: %v", goldenPath, err)
			}

			// Parse expected values from golden file
			expected := parseGoldenFile(string(goldenContent))

			// Create mock server context and add document
			ctx := testhelpers.NewMockServerContext()

			// Create DocumentManager and add document
			dm, err := lsp.NewDocumentManager()
			if err != nil {
				t.Fatalf("Failed to create DocumentManager: %v", err)
			}
			defer dm.Close()
			ctx.SetDocumentManager(dm)

			doc := dm.OpenDocument("test://fixture.html", string(htmlContent), 1)
			ctx.AddDocument("test://fixture.html", doc)

			// Test base and script indentation detection
			baseIndent, scriptIndent := detectIndentation(doc)

			if baseIndent != expected.baseIndent {
				t.Errorf("Expected base indent %q (len=%d), got %q (len=%d)",
					expected.baseIndent, len(expected.baseIndent),
					baseIndent, len(baseIndent))
			}

			if scriptIndent != expected.scriptIndent {
				t.Errorf("Expected script indent %q (len=%d), got %q (len=%d)",
					expected.scriptIndent, len(expected.scriptIndent),
					scriptIndent, len(scriptIndent))
			}

			// Test script content indentation (position of <script> tag)
			scriptPosition := findScriptPosition(string(htmlContent))
			if scriptPosition != nil {
				scriptContentIndent := detectScriptTagIndentation(doc, *scriptPosition)
				if scriptContentIndent != expected.scriptContentIndent {
					t.Errorf("Expected script content indent %q (len=%d), got %q (len=%d)",
						expected.scriptContentIndent, len(expected.scriptContentIndent),
						scriptContentIndent, len(scriptContentIndent))
				}
			}

			t.Logf("✓ %s", fixtureName)
			t.Logf("  Base indent: %q (len=%d)", baseIndent, len(baseIndent))
			t.Logf("  Script indent: %q (len=%d)", scriptIndent, len(scriptIndent))
			if scriptPosition != nil {
				scriptContentIndent := detectScriptTagIndentation(doc, *scriptPosition)
				t.Logf("  Script content indent: %q (len=%d)", scriptContentIndent, len(scriptContentIndent))
			}
		})
	}
}

// expectedIndentation holds the expected indentation values from a golden file
type expectedIndentation struct {
	baseIndent          string
	scriptIndent        string
	scriptContentIndent string
}

// parseGoldenFile parses the golden file format into expected values
func parseGoldenFile(content string) expectedIndentation {
	var expected expectedIndentation

	// Parse lines like: baseIndent:"  "
	lines := strings.SplitSeq(strings.TrimSpace(content), "\n")
	for line := range lines {
		if strings.HasPrefix(line, "baseIndent:") {
			expected.baseIndent = parseQuotedString(line)
		} else if strings.HasPrefix(line, "scriptIndent:") {
			expected.scriptIndent = parseQuotedString(line)
		} else if strings.HasPrefix(line, "scriptContentIndent:") {
			expected.scriptContentIndent = parseQuotedString(line)
		}
	}

	return expected
}

// parseQuotedString extracts the quoted string value from a line like: key:"value"
func parseQuotedString(line string) string {
	// Extract everything after the colon and parse as Go string literal
	colonIndex := strings.Index(line, ":")
	if colonIndex == -1 {
		return ""
	}

	quotedValue := strings.TrimSpace(line[colonIndex+1:])

	// Handle Go string literals with escape sequences
	if len(quotedValue) >= 2 && quotedValue[0] == '"' && quotedValue[len(quotedValue)-1] == '"' {
		// Remove quotes and handle escape sequences
		unquoted := quotedValue[1 : len(quotedValue)-1]
		// Replace \t with actual tab character
		unquoted = strings.ReplaceAll(unquoted, "\\t", "\t")
		// Replace \n with actual newline (if needed)
		unquoted = strings.ReplaceAll(unquoted, "\\n", "\n")
		return unquoted
	}

	return quotedValue
}

// findScriptPosition finds the line number of the <script> tag in HTML content
func findScriptPosition(content string) *protocol.Position {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.Contains(line, "<script") {
			return &protocol.Position{Line: uint32(i), Character: 0}
		}
	}
	return nil
}

// Test helper functions that mirror the private functions from missingImportCodeActions.go

// detectIndentation mimics the private detectIndentation function
func detectIndentation(doc types.Document) (string, string) {
	if doc == nil {
		return "  ", "  "
	}

	content, err := doc.Content()
	if err != nil {
		return "  ", "  "
	}

	lines := strings.Split(content, "\n")

	// Collect indentation samples from non-empty lines
	var indentLevels []int
	var hasTab bool

	// Regex to capture leading whitespace
	indentRegex := regexp.MustCompile(`^(\s+)`)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		}

		matches := indentRegex.FindStringSubmatch(line)
		if len(matches) > 1 && matches[1] != "" {
			indent := matches[1]

			// Check if file uses tabs
			if strings.Contains(indent, "\t") {
				hasTab = true
				indentLevels = append(indentLevels, strings.Count(indent, "\t"))
			} else {
				// Count spaces
				spaceCount := len(indent)
				if spaceCount > 0 {
					indentLevels = append(indentLevels, spaceCount)
				}
			}
		}
	}

	if hasTab {
		// File uses tabs
		return "\t", "\t\t"
	}

	if len(indentLevels) > 0 {
		// Find the smallest non-zero indentation (likely base indentation unit)
		minIndent := indentLevels[0]
		for _, level := range indentLevels {
			if level > 0 && level < minIndent {
				minIndent = level
			}
		}

		// Use the detected minimum indentation as the base unit
		baseIndentStr := strings.Repeat(" ", minIndent)
		scriptIndentStr := strings.Repeat(" ", minIndent*2) // Double for script content

		return baseIndentStr, scriptIndentStr
	}

	// No indentation found, default to 2 spaces (common in HTML)
	return "  ", "  "
}

// detectScriptTagIndentation mimics the private detectScriptTagIndentation function
func detectScriptTagIndentation(doc types.Document, scriptPosition protocol.Position) string {
	if doc == nil {
		return "  " // Default to 2 spaces
	}

	content, err := doc.Content()
	if err != nil {
		return "  " // Default to 2 spaces
	}

	lines := strings.Split(content, "\n")

	// Collect indentation from all existing import statements to find the most consistent pattern
	var importIndents []string
	scriptStartLine := int(scriptPosition.Line)
	indentRegex := regexp.MustCompile(`^(\s*)`)

	for i := scriptStartLine; i < len(lines); i++ {
		line := lines[i]

		// Stop if we hit the closing script tag
		if strings.Contains(line, "</script>") {
			break
		}

		// Look for existing import statements
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "import ") {
			// Extract the indentation from this import line
			matches := indentRegex.FindStringSubmatch(line)
			if len(matches) > 1 && matches[1] != "" {
				importIndents = append(importIndents, matches[1])
			}
		}
	}

	// If we found import statements, check if they're consistent
	if len(importIndents) > 0 {
		// Count occurrences of each indentation pattern
		indentCounts := make(map[string]int)
		for _, indent := range importIndents {
			indentCounts[indent]++
		}

		// If all imports use the same indentation, use it
		if len(indentCounts) == 1 {
			for indent := range indentCounts {
				return indent
			}
		}

		// If imports are inconsistent, prefer file's general indentation for consistency
		// This standardizes mixed indentation rather than perpetuating it
	}

	// If imports are inconsistent, prefer file's general indentation for consistency
	if len(importIndents) > 0 {
		baseIndent, _ := detectIndentation(doc)
		return baseIndent
	}

	// If no existing imports found, use the file's script indentation pattern
	_, scriptIndent := detectIndentation(doc)
	return scriptIndent
}
