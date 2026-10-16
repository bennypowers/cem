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
	"unicode/utf8"

	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// generateSourceReference creates a SourceReference for a declaration if sourceControlRootUrl is configured
// Returns nil if node is nil, config is unavailable, or sourceControlRootUrl is not configured
func (mp *ModuleProcessor) generateSourceReference(node *ts.Node) (*M.SourceReference, error) {
	if node == nil {
		return nil, nil // No error, just no source reference
	}
	return mp.sourceReferenceAt(node.StartByte())
}

// generateTypeParameterSourceReference creates a SourceReference for a
// generic type parameter's declaration. The schema only requires an absolute
// URL; the anchor follows GitHub's line-and-column fragment convention
// (#L2C28) because several parameters usually share one line.
func (mp *ModuleProcessor) generateTypeParameterSourceReference(node *ts.Node) (*M.SourceReference, error) {
	if node == nil {
		return nil, nil // No error, just no source reference
	}
	line, column := mp.byteOffsetToLineColumn(node.StartByte())
	return mp.sourceReferenceAtLineColumn(line, column)
}

// sourceReferenceAt builds a line-anchored SourceReference for a byte offset,
// for declaration-level references.
func (mp *ModuleProcessor) sourceReferenceAt(offset uint) (*M.SourceReference, error) {
	line, _ := mp.byteOffsetToLineColumn(offset)
	return mp.sourceReferenceAtLineColumn(line, 0)
}

// sourceReferenceAtLineColumn builds a SourceReference anchored at the given
// line, plus a column when nonzero, if sourceControlRootUrl is configured.
// Returns nil without error when unconfigured or when no href can be built.
func (mp *ModuleProcessor) sourceReferenceAtLineColumn(line, column uint) (*M.SourceReference, error) {
	if mp.ctx == nil {
		return nil, nil // No error, just no source reference
	}

	cfg, err := mp.ctx.Config()
	if err != nil {
		return nil, fmt.Errorf("failed to get config for source reference: %w", err)
	}

	if cfg.SourceControlRootUrl == "" {
		return nil, nil // No error, just not configured
	}

	href, err := buildSourceHrefAt(cfg.SourceControlRootUrl, mp.file, line, column)
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
func (mp *ModuleProcessor) byteOffsetToLineNumber(offset uint) uint {
	line, _ := mp.byteOffsetToLineColumn(offset)
	return line
}

// byteOffsetToLineColumn converts a byte offset to a 1-based line number and
// a 1-based rune column within that line, for anchors that must address a
// position inside a line (e.g. one of several type parameters). Both
// lookups scan the cached newline offsets linearly.
func (mp *ModuleProcessor) byteOffsetToLineColumn(offset uint) (line, column uint) {
	if mp.lineOffsets == nil {
		mp.buildLineOffsetCache()
	}
	line, lineStart := uint(1), uint(0)
	for _, lineOffset := range mp.lineOffsets {
		if offset < lineOffset {
			break
		}
		line++
		lineStart = lineOffset + 1
	}
	if lineStart >= offset {
		return line, 1
	}
	return line, uint(utf8.RuneCount(mp.code[lineStart:offset])) + 1
}

// buildLineOffsetCache builds a cache of byte offsets for each newline
// so line lookups scan the cache instead of the whole source
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
	return buildSourceHrefAt(sourceControlRootUrl, filePath, lineNumber, 0)
}

// buildSourceHrefAt constructs a source href URL, additionally anchoring a
// 1-based column when column > 0 (GitHub-style #L2C28); column 0 emits a
// line-only anchor (#L2).
func buildSourceHrefAt(sourceControlRootUrl, filePath string, lineNumber, column uint) (string, error) {
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

	// Add the line (and column, when given) fragment
	if column > 0 {
		u.Fragment = fmt.Sprintf("L%dC%d", lineNumber, column)
	} else {
		u.Fragment = fmt.Sprintf("L%d", lineNumber)
	}

	return u.String(), nil
}
