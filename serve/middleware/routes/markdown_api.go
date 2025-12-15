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

package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

const (
	// maxMarkdownSize is the maximum allowed size for markdown text (1MB)
	// This prevents DoS attacks from excessively large payloads
	maxMarkdownSize = 1 * 1024 * 1024 // 1MB
)

// MarkdownRequest represents a request to convert markdown to HTML from a manifest field
// Only markdown fields from the Custom Elements Manifest can be rendered.
type MarkdownRequest struct {
	// Path is a gjson path to the markdown field in the manifest
	// Examples:
	//   "modules.0.declarations.0.summary"
	//   "modules.#(path==\"src/button.ts\").declarations.#(name==\"MyButton\").description"
	Path string `json:"path"`
}

// MarkdownResponse represents the HTML output from markdown conversion
type MarkdownResponse struct {
	HTML string `json:"html"`
}

// serveMarkdownAPI handles POST requests to render markdown to HTML from manifest fields
// This endpoint allows client-side components to render markdown without
// bundling markdown libraries, reducing client bundle size. For security,
// only markdown from the Custom Elements Manifest can be rendered.
func serveMarkdownAPI(w http.ResponseWriter, r *http.Request, config Config) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit request body size to prevent DoS attacks
	r.Body = http.MaxBytesReader(w, r.Body, maxMarkdownSize)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("r.Body.Close: %v", err)
		}
	}()

	var req MarkdownRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Path == "" {
		// Empty path is valid - return empty HTML
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(MarkdownResponse{HTML: ""}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// Get manifest from context
	manifestBytes, err := config.Context.Manifest()
	if err != nil {
		config.Context.Logger().Error("Failed to get manifest: %v", err)
		http.Error(w, "Manifest not available", http.StatusInternalServerError)
		return
	}

	// Query the manifest using the provided gjson path
	result := gjson.GetBytes(manifestBytes, req.Path)
	if !result.Exists() {
		http.Error(w, "Path not found in manifest", http.StatusNotFound)
		return
	}

	// Ensure the result is a string (markdown field)
	if result.Type != gjson.String {
		http.Error(w, "Path does not point to a string field", http.StatusBadRequest)
		return
	}

	markdownText := result.String()
	if markdownText == "" {
		// Empty markdown is valid - return empty HTML
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(MarkdownResponse{HTML: ""}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	html, err := markdownToHTML(markdownText)
	if err != nil {
		http.Error(w, "Markdown conversion failed", http.StatusInternalServerError)
		return
	}

	resp := MarkdownResponse{HTML: html}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
