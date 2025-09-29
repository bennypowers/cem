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
package typescript

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TemplateContext represents an HTML template context in TypeScript
type TemplateContext struct {
	Range   protocol.Range
	content string
	Type    string // "html", "innerHTML", "outerHTML"
}

// Content returns the template content
func (tc *TemplateContext) Content() (string, error) {
	return tc.content, nil
}