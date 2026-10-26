/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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

// Package markup validates HTML and template markup against a custom elements
// manifest.
//
// Unlike the parent [bennypowers.dev/cem/validate] package, which validates
// manifests themselves, this package needs tree-sitter and therefore cgo. It
// links only the HTML grammar and the Jinja dialect grammar used for Twig and
// Nunjucks, so it can be used without an LSP document manager, MCP server, or
// workspace.
//
// Callers opt in to attribute and slot checks:
//
//	pkg := &manifest.Package{ /* ... */ }
//	issues, err := markup.Validate(pkg, `<my-button variant="primary">Save</my-button>`, markup.Options{
//		Format:        "html",
//		ValidateAttrs: true,
//		ValidateSlots: true,
//	})
//	if err != nil {
//		return err
//	}
//	for _, issue := range issues {
//		log.Printf("%s:%d:%d %s: %s", issue.Element, issue.Line, issue.Column, issue.Rule, issue.Message)
//	}
//
// Manifest-backed findings (unknown elements, attributes, slot names, invalid
// enum values, deprecated APIs) carry [ConfidenceHigh]. Findings inferred from
// prose or heuristics carry [ConfidenceLow] and are only reported when
// [Options.Strict] is set.
package markup
