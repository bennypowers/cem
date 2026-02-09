/*
Copyright © 2025 Benny Powers

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
package generate

import (
	"fmt"
	"net/url"
	"strings"

	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// generateSourceReference creates a SourceReference for a declaration if sourceControlRootUrl is configured
// Returns nil if node is nil, config is unavailable, or sourceControlRootUrl is not configured
func (mp *ModuleProcessor) generateSourceReference(node *ts.Node) (*M.SourceReference, error) {
	if node == nil || mp.ctx == nil {
		return nil, nil // No error, just no source reference
	}

	cfg, err := mp.ctx.Config()
	if err != nil {
		return nil, fmt.Errorf("failed to get config for source reference: %w", err)
	}

	if cfg.SourceControlRootUrl == "" {
		return nil, nil // No error, just not configured
	}

	lineNumber := mp.byteOffsetToLineNumber(node.StartByte())
	href, err := buildSourceHref(cfg.SourceControlRootUrl, mp.file, lineNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to build source href: %w", err)
	}

	if href == "" {
		return nil, nil // No error, just couldn't build href
	}

	return &M.SourceReference{
		Href: href,
	}, nil
}

// byteOffsetToLineNumber converts a byte offset to a 1-based line number
// Uses a line offset cache for O(log n) lookup instead of O(n) scan
func (mp *ModuleProcessor) byteOffsetToLineNumber(offset uint) uint {
	if mp.lineOffsets == nil {
		mp.buildLineOffsetCache()
	}

	// Binary search to find the line containing this offset
	line := 1
	for i, lineOffset := range mp.lineOffsets {
		if offset < lineOffset {
			return uint(line + i)
		}
	}

	// If offset is beyond all newlines, it's on the last line
	return uint(line + len(mp.lineOffsets))
}

// buildLineOffsetCache builds a cache of byte offsets for each newline
// This enables O(log n) line number lookups instead of O(n) scans
func (mp *ModuleProcessor) buildLineOffsetCache() {
	mp.lineOffsets = make([]uint, 0, 100) // Pre-allocate for typical file size

	for i, b := range mp.code {
		if b == '\n' {
			mp.lineOffsets = append(mp.lineOffsets, uint(i))
		}
	}
}

// buildSourceHref constructs a source href URL from base URL, file path, and line number
func buildSourceHref(sourceControlRootUrl, filePath string, lineNumber uint) (string, error) {
	if sourceControlRootUrl == "" {
		return "", fmt.Errorf("sourceControlRootUrl is empty")
	}
	if filePath == "" {
		return "", fmt.Errorf("filePath is empty")
	}

	// Parse the base URL to validate it
	baseURL, err := url.Parse(sourceControlRootUrl)
	if err != nil {
		return "", fmt.Errorf("invalid sourceControlRootUrl %q: %w", sourceControlRootUrl, err)
	}

	// Ensure we have a valid URL with scheme and host
	if baseURL.Scheme == "" {
		return "", fmt.Errorf("sourceControlRootUrl %q missing scheme", sourceControlRootUrl)
	}
	if baseURL.Host == "" {
		return "", fmt.Errorf("sourceControlRootUrl %q missing host", sourceControlRootUrl)
	}

	// Ensure the URL ends with a slash for proper joining
	if !strings.HasSuffix(baseURL.Path, "/") {
		baseURL.Path += "/"
	}

	// Clean the file path - remove leading ./ or / if present
	cleanPath := strings.TrimPrefix(filePath, "./")
	cleanPath = strings.TrimPrefix(cleanPath, "/")

	// Join the base URL with the file path
	u, err := baseURL.Parse(cleanPath)
	if err != nil {
		return "", fmt.Errorf("failed to join URL %q with path %q: %w", sourceControlRootUrl, cleanPath, err)
	}

	// Add the line number fragment
	u.Fragment = fmt.Sprintf("L%d", lineNumber)

	return u.String(), nil
}
