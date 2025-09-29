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
package document

import (
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// extractDocumentManager safely extracts a types.Manager from any type,
// handling both real types.Manager and mock implementations
func ExtractDocumentManager(dm any) types.Manager {
	if dm == nil {
		return nil
	}

	// Try direct cast first
	if directDM, ok := dm.(types.Manager); ok {
		return directDM
	}

	// Try to extract from mock implementations
	if mockDM, ok := dm.(interface{ GetRealDocumentManager() types.Manager }); ok {
		return mockDM.GetRealDocumentManager()
	}

	// Last resort: handle other types
	helpers.SafeDebugLog("[DOCUMENT] Warning: DocumentManager type %T not recognized, using nil", dm)
	return nil
}

// getLanguageFromURI determines the language from file extension
func getLanguageFromURI(uri string) string {
	ext := strings.ToLower(filepath.Ext(uri))
	switch ext {
	case ".html", ".htm":
		return "html"
	case ".ts":
		return "typescript"
	case ".js":
		return "typescript" // Use TypeScript parser for JS too
	case ".tsx", ".jsx":
		return "tsx" // TSX files get their own language type
	default:
		return "html" // Default to HTML
	}
}

// Helper function to check if position is within range
func isPositionInRange(pos protocol.Position, r protocol.Range) bool {
	if pos.Line < r.Start.Line || pos.Line > r.End.Line {
		return false
	}

	if pos.Line == r.Start.Line && pos.Character < r.Start.Character {
		return false
	}

	if pos.Line == r.End.Line && pos.Character > r.End.Character {
		return false
	}

	return true
}