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

// DataFetcher represents a declarative data source for tools
type DataFetcher struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`     // Legacy: will be deprecated
	Source   string `yaml:"source"`   // New: data source name (registry, args, previous fetcher name)
	Path     string `yaml:"path"`     // JSONPath/gjson expression
	Filter   string `yaml:"filter"`   // Optional: first, count, exists
	Required bool   `yaml:"required,omitempty"`
}

// ToolDefinition represents a complete tool definition with metadata and handler
type ToolDefinition struct {
	Name         string                 `yaml:"name"`
	Description  string                 `yaml:"-"` // From markdown content
	InputSchema  map[string]interface{} `yaml:"inputSchema"`
	DataFetchers []DataFetcher          `yaml:"dataFetchers,omitempty"`
	Template     string                 `yaml:"template,omitempty"`
	ResponseType string                 `yaml:"responseType,omitempty"`
	Handler      mcp.ToolHandler        `yaml:"-"`
}

// Frontmatter represents the YAML frontmatter from tool markdown files
type Frontmatter struct {
	Name         string                 `yaml:"name"`
	InputSchema  map[string]interface{} `yaml:"inputSchema"`
	DataFetchers []DataFetcher          `yaml:"dataFetchers,omitempty"`
	Template     string                 `yaml:"template,omitempty"`
	ResponseType string                 `yaml:"responseType,omitempty"`
}

// ResourceDefinition represents a complete resource definition with metadata and handler
type ResourceDefinition struct {
	URI          string              `yaml:"uri"`
	Name         string              `yaml:"name"`
	MimeType     string              `yaml:"mimeType"`
	URITemplate  bool                `yaml:"uriTemplate,omitempty"`
	Description  string              `yaml:"-"` // From markdown content
	DataFetchers []DataFetcher       `yaml:"dataFetchers,omitempty"`
	Template     string              `yaml:"template,omitempty"`
	ResponseType string              `yaml:"responseType,omitempty"`
	Handler      mcp.ResourceHandler `yaml:"-"`
}

// ResourceFrontmatter represents the YAML frontmatter from resource markdown files
type ResourceFrontmatter struct {
	URI          string        `yaml:"uri"`
	Name         string        `yaml:"name"`
	MimeType     string        `yaml:"mimeType"`
	URITemplate  bool          `yaml:"uriTemplate,omitempty"`
	DataFetchers []DataFetcher `yaml:"dataFetchers,omitempty"`
	Template     string        `yaml:"template,omitempty"`
	ResponseType string        `yaml:"responseType,omitempty"`
}
