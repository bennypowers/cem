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

package testutil

import (
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/textutil"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

// LSPFixture represents a single LSP test scenario loaded from fixtures
type LSPFixture struct {
	Name            string              // Test scenario name (directory name)
	InputHTML       string              // HTML content from input.html (deprecated: use InputContent)
	InputContent    string              // Content from input.html or input.ts
	InputType       string              // File type: "html" or "ts"
	Manifest        json.RawMessage     // Optional manifest data from manifest.json
	ExpectedMap     map[string]any      // Expected results from expected-*.json or expected.json
	AdditionalFiles map[string]string   // Additional source files in the fixture directory (filename → content)
	Cursor          *protocol.Position  // Cursor position extracted from ^cursor marker or YAML frontmatter
}

// RunLSPFixtures discovers and runs LSP tests from a testdata directory.
// Each subdirectory in testdata/ represents one test scenario with:
//   - input.html or input.ts (required): Content to test
//   - manifest.json (optional): Manifest data for custom elements
//   - expected.json or expected-*.json (optional): Expected results for assertions
//
// All fixture files are preloaded into MapFS at the start — no disk I/O
// during individual test execution.
func RunLSPFixtures(t *testing.T, testdataDir string, testFunc func(*testing.T, *LSPFixture)) {
	t.Helper()

	mfs := LoadTestdataFS(t, testdataDir, "/")

	entries, err := fs.ReadDir(mfs.GetMapFS(), ".")
	if err != nil {
		t.Fatalf("Failed to read testdata directory %s: %v", testdataDir, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		scenarioName := entry.Name()

		t.Run(scenarioName, func(t *testing.T) {
			fixture := loadLSPFixtureFromMapFS(t, mfs, scenarioName)
			testFunc(t, fixture)
		})
	}
}

// RunLSPFixture loads a single named fixture from testdataDir and runs testFunc.
// Use instead of RunLSPFixtures when a test targets exactly one fixture.
func RunLSPFixture(t *testing.T, testdataDir, fixtureName string, testFunc func(*testing.T, *LSPFixture)) {
	t.Helper()

	mfs := LoadTestdataFS(t, testdataDir, "/")

	t.Run(fixtureName, func(t *testing.T) {
		fixture := loadLSPFixtureFromMapFS(t, mfs, fixtureName)
		testFunc(t, fixture)
	})
}

// loadLSPFixtureFromMapFS loads a single LSP test fixture from MapFS
func loadLSPFixtureFromMapFS(t *testing.T, mfs *platform.MapFileSystem, scenarioName string) *LSPFixture {
	t.Helper()

	fixture := &LSPFixture{
		Name:            scenarioName,
		ExpectedMap:     make(map[string]any),
		AdditionalFiles: make(map[string]string),
	}

	// Try input.html first, then input.ts, then input.css
	found := false
	for _, candidate := range []struct {
		name     string
		fileType string
	}{
		{"input.html", "html"},
		{"input.ts", "ts"},
		{"input.css", "css"},
	} {
		path := filepath.Join(scenarioName, candidate.name)
		data, err := mfs.ReadFile(path)
		if err == nil {
			content := string(data)
			cleaned, cursor := extractCursor(content, candidate.fileType)
			fixture.InputContent = cleaned
			fixture.InputHTML = cleaned
			fixture.InputType = candidate.fileType
			fixture.Cursor = cursor
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("No input.html, input.ts, or input.css found for scenario %s", scenarioName)
	}

	// Load manifest.json (optional)
	manifestPath := filepath.Join(scenarioName, "manifest.json")
	if data, err := mfs.ReadFile(manifestPath); err == nil {
		fixture.Manifest = json.RawMessage(data)
	}

	// Load expected-*.json and additional source files
	scenarioEntries, err := fs.ReadDir(mfs.GetMapFS(), scenarioName)
	if err != nil {
		return fixture
	}

	primaryInput := "input." + fixture.InputType

	for _, entry := range scenarioEntries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		ext := filepath.Ext(name)
		filePath := filepath.Join(scenarioName, name)

		switch {
		case ext == ".json" && name != "manifest.json" && name != "package.json":
			data, err := mfs.ReadFile(filePath)
			if err != nil {
				continue
			}
			var expectedData any
			if err := json.Unmarshal(data, &expectedData); err != nil {
				t.Logf("Warning: failed to parse %s: %v", name, err)
				continue
			}
			var key string
			if name == "expected.json" {
				key = "expected"
			} else if strings.HasPrefix(name, "expected-") {
				key = name[len("expected-") : len(name)-len(".json")]
			} else {
				key = name[:len(name)-len(".json")]
			}
			fixture.ExpectedMap[key] = expectedData

		case (ext == ".ts" || ext == ".js" || ext == ".html" || ext == ".css") && name != primaryInput:
			data, err := mfs.ReadFile(filePath)
			if err != nil {
				t.Logf("Warning: failed to read additional file %s: %v", name, err)
				continue
			}
			fixture.AdditionalFiles[name] = string(data)
		}
	}

	return fixture
}

// GetExpected returns typed expected data from the fixture
func (f *LSPFixture) GetExpected(key string, v any) error {
	data, ok := f.ExpectedMap[key]
	if !ok {
		return nil // No expected data for this key
	}

	// Re-marshal and unmarshal to convert to desired type
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, v)
}

// CursorParser extracts a ^cursor marker from file content.
// Returns cleaned content (marker stripped) and cursor position (nil if not found).
type CursorParser func(content string) (string, *protocol.Position)

// ParseCursor extracts cursor position using the given parser.
// Skips if Cursor is already set (e.g., by frontmatter or HTML extraction during load).
func (f *LSPFixture) ParseCursor(parser CursorParser) {
	if f.Cursor != nil {
		return
	}
	cleaned, cursor := parser(f.InputContent)
	if cursor != nil {
		f.InputContent = cleaned
		f.InputHTML = cleaned
		f.Cursor = cursor
	}
}

// extractCursor checks content for a ^cursor comment marker or YAML frontmatter.
// Marker always takes precedence. Frontmatter is stripped when it contains a cursor key.
// For HTML, uses golang.org/x/net/html tokenizer to find comment nodes.
// For TS/CSS, only checks frontmatter; use ParseCursor with a tree-sitter parser.
func extractCursor(content string, fileType string) (string, *protocol.Position) {
	// Strip frontmatter first (if present) so it never appears in InputContent
	body, fmCursor := extractCursorFrontmatter(content)
	if fileType == "html" {
		cleaned, markerCursor := extractHTMLCursorMarker(body)
		if markerCursor != nil {
			return cleaned, markerCursor
		}
	}
	if fmCursor != nil {
		return body, fmCursor
	}
	return content, nil
}

// extractHTMLCursorMarker uses the Go HTML tokenizer to find <!-- ^cursor --> comments.
func extractHTMLCursorMarker(content string) (string, *protocol.Position) {
	tokenizer := html.NewTokenizer(strings.NewReader(content))
	byteOffset := 0
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			break
		}
		raw := string(tokenizer.Raw())
		if tt == html.CommentToken && strings.Contains(raw, "^cursor") {
			caretInRaw := strings.Index(raw, "^")
			if caretInRaw < 0 {
				byteOffset += len(raw)
				continue
			}
			caretGlobal := byteOffset + caretInRaw
			caretLine, caretChar := byteOffsetToPosition(content, caretGlobal)
			if caretLine == 0 {
				byteOffset += len(raw)
				continue
			}
			// Strip the entire line containing the comment (LF and CRLF)
			lineStart := byteOffset
			for lineStart > 0 && content[lineStart-1] != '\n' {
				lineStart--
			}
			lineEnd := byteOffset + len(raw)
			if lineEnd < len(content) && content[lineEnd] == '\r' {
				lineEnd++
			}
			if lineEnd < len(content) && content[lineEnd] == '\n' {
				lineEnd++
			}
			cleaned := content[:lineStart] + content[lineEnd:]
			return cleaned, &protocol.Position{
				Line:      uint32(caretLine - 1),
				Character: caretChar,
			}
		}
		byteOffset += len(raw)
	}
	return content, nil
}

// byteOffsetToPosition converts a byte offset in content to an LSP-compatible
// (line, character) pair where character is in UTF-16 code units.
func byteOffsetToPosition(content string, offset int) (line int, character uint32) {
	lineStart := 0
	for i := range offset {
		if i >= len(content) {
			break
		}
		if content[i] == '\n' {
			line++
			lineStart = i + 1
		}
	}
	byteCol := uint(offset - lineStart)
	character = textutil.ByteOffsetToUTF16(content[lineStart:], byteCol)
	return line, character
}

// cursorFrontmatter is the YAML structure for cursor position in frontmatter.
type cursorFrontmatter struct {
	Cursor *struct {
		Line      uint32 `yaml:"line"`
		Character uint32 `yaml:"character"`
	} `yaml:"cursor"`
}

// extractCursorFrontmatter checks for YAML frontmatter with cursor position.
func extractCursorFrontmatter(content string) (string, *protocol.Position) {
	if !strings.HasPrefix(content, "---\n") {
		return content, nil
	}
	end := strings.Index(content[4:], "\n---\n")
	if end < 0 {
		return content, nil
	}
	fmContent := content[4 : 4+end]
	body := content[4+end+5:]

	var fm cursorFrontmatter
	if err := yaml.Unmarshal([]byte(fmContent), &fm); err != nil {
		return content, nil
	}
	if fm.Cursor == nil {
		return content, nil
	}
	return body, &protocol.Position{
		Line:      fm.Cursor.Line,
		Character: fm.Cursor.Character,
	}
}
