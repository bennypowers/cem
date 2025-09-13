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
package types

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ToolDefinition represents a complete tool definition with metadata and handler
type ToolDefinition struct {
	Name        string                 `yaml:"name"`
	Description string                 `yaml:"-"` // From markdown content
	InputSchema map[string]interface{} `yaml:"inputSchema"`
	Handler     mcp.ToolHandler        `yaml:"-"`
}

// Frontmatter represents the YAML frontmatter from tool markdown files
type Frontmatter struct {
	Name        string                 `yaml:"name"`
	InputSchema map[string]interface{} `yaml:"inputSchema"`
}
