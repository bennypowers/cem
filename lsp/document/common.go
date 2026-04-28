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
)

// ExtractDocumentManager safely extracts a types.Manager from any type,
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

// getLanguageFromURI determines the language from file extension.
// Template files (Nunjucks, Jinja2, Twig, Liquid, Handlebars) use injection
// handlers that strip template syntax before HTML parsing. Blade uses a
// dedicated handler that parses with tree-sitter-blade directly.
func getLanguageFromURI(uri string) string {
	if strings.HasSuffix(strings.ToLower(uri), ".blade.php") {
		return "blade"
	}
	ext := strings.ToLower(filepath.Ext(uri))
	switch ext {
	case ".html", ".htm":
		return "html"
	case ".php":
		return "php"
	case ".ts":
		return "typescript"
	case ".js":
		return "typescript"
	case ".tsx", ".jsx":
		return "tsx"
	case ".njk", ".j2", ".jinja", ".jinja2", ".liquid", ".hbs", ".twig", ".erb", ".ejs":
		return "template"
	default:
		return "html"
	}
}
