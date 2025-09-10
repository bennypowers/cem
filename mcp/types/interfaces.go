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
