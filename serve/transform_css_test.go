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

package serve

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTransformCSS_Basic tests basic CSS to constructable stylesheet transformation
func TestTransformCSS_Basic(t *testing.T) {
	cssContent := `:host {
  display: block;
  color: var(--color, red);
}`

	result := TransformCSS([]byte(cssContent))

	// Should wrap in CSSStyleSheet
	if !strings.Contains(result, "new CSSStyleSheet()") {
		t.Error("Output missing CSSStyleSheet constructor")
	}

	// Should use replaceSync
	if !strings.Contains(result, "replaceSync") {
		t.Error("Output missing replaceSync call")
	}

	// Should export default
	if !strings.Contains(result, "export default") {
		t.Error("Output missing default export")
	}

	// Should contain the original CSS
	if !strings.Contains(result, ":host") {
		t.Error("Output missing original CSS content")
	}

	// Should escape backticks in CSS
	cssWithBacktick := "content: `test`;"
	escaped := TransformCSS([]byte(cssWithBacktick))
	if strings.Contains(escaped, "content: `test`;") {
		t.Error("Backticks in CSS not escaped")
	}
}

// TestServCSS_TransformsToModule tests HTTP serving of CSS as JavaScript module
func TestServeCSS_TransformsToModule(t *testing.T) {
	tmpDir := t.TempDir()

	// Write a CSS file
	cssContent := `:host { display: block; }`
	cssPath := filepath.Join(tmpDir, "styles.css")
	if err := os.WriteFile(cssPath, []byte(cssContent), 0644); err != nil {
		t.Fatalf("Failed to write CSS file: %v", err)
	}

	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	if err := server.SetWatchDir(tmpDir); err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Request the CSS file
	req := httptest.NewRequest("GET", "/styles.css", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	// Should return 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Should have JavaScript content type (not text/css)
	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/javascript") {
		t.Errorf("Expected JavaScript content type, got %s", contentType)
	}

	// Should be transformed to constructable stylesheet
	body := w.Body.String()
	if !strings.Contains(body, "CSSStyleSheet") {
		t.Error("CSS not transformed to CSSStyleSheet module")
	}

	// Should contain original CSS
	if !strings.Contains(body, ":host") {
		t.Error("Transformed output missing original CSS")
	}
}

// TestServeCSS_WarnsWithoutImportAttribute tests warning for CSS imports without 'with { type: "css" }'
func TestServeCSS_WarnsWithoutImportAttribute(t *testing.T) {
	t.Skip("Warning implementation TBD - need to detect import context")
	// This test verifies that when CSS is imported without `with { type: 'css' }`,
	// the server logs a warning. Implementation requires examining the referrer
	// to detect how the CSS is being imported.
}
