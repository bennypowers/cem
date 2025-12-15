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
	"net/http"
)

// MarkdownRequest represents a request to convert markdown to HTML
type MarkdownRequest struct {
	Text string `json:"text"`
}

// MarkdownResponse represents the HTML output from markdown conversion
type MarkdownResponse struct {
	HTML string `json:"html"`
}

// serveMarkdownAPI handles POST requests to render markdown to HTML
// This endpoint allows client-side components to render markdown without
// bundling markdown libraries, reducing client bundle size
func serveMarkdownAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req MarkdownRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		// Empty text is valid - return empty HTML
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MarkdownResponse{HTML: ""})
		return
	}

	html, err := markdownToHTML(req.Text)
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
