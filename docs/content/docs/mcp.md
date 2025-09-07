---
title: MCP Server
weight: 30
---

The `cem mcp` command starts a Model Context Protocol (MCP) server that provides AI-native access to your custom elements manifest data. This enables intelligent HTML generation, component understanding, and design system compliance for AI assistants.

## Overview

The MCP server transforms your custom elements manifests into a structured, type-safe interface that AI systems can use to:

- Generate correct HTML with proper slot and attribute usage
- Understand component APIs and relationships
- Provide CSS integration guidance for custom properties and parts
- Enforce design system guidelines and best practices
- Validate component usage in real-time

## Quick Start

```bash
# Start MCP server for current workspace
cem mcp

# Start with specific transport
cem mcp --transport stdio
cem mcp --transport tcp --port 8080
cem mcp --transport websocket --port 9090
```

## Command Options

### Transport Options

```bash
--transport string    Transport type (stdio, tcp, websocket) (default "stdio")
--port int           Port for TCP/WebSocket transports (default 8080)
--host string        Host for TCP/WebSocket transports (default "localhost")
```

### Server Options

```bash
--cache              Enable element info caching (default true)
--watch              Watch manifest files for changes (default true)
--debug              Enable debug logging
```

## MCP Resources

The server provides several resource types that AI systems can access:

### Schema Resource
```
cem://schema
```
Returns the JSON schema for custom elements manifests, enabling validation and understanding of manifest structure.

### Registry Resource
```
cem://registry
```
Provides the complete registry of all available elements with their metadata, including counts and basic information.

### Element Resources
```
cem://element/{tagName}
```
Returns detailed information about a specific element, including all attributes, slots, events, CSS properties, parts, and states.

### Package Resources
```
cem://package/{packageName}
```
Provides package-specific manifest information for multi-package workspaces.

## MCP Tools

The server provides interactive tools for AI assistance:

### Element Validation
```json
{
  "name": "validate_element",
  "description": "Validate custom element usage",
  "inputSchema": {
    "type": "object",
    "properties": {
      "tagName": {"type": "string"},
      "attributes": {"type": "object"},
      "slots": {"type": "object"}
    }
  }
}
```

### Attribute Suggestions
```json
{
  "name": "suggest_attributes",
  "description": "Get valid attributes with type-aware suggestions",
  "inputSchema": {
    "type": "object", 
    "properties": {
      "tagName": {"type": "string"},
      "context": {"type": "string"}
    }
  }
}
```

### HTML Generation
```json
{
  "name": "generate_html",
  "description": "Generate correct HTML structure with proper slots and attributes",
  "inputSchema": {
    "type": "object",
    "properties": {
      "tagName": {"type": "string"},
      "content": {"type": "object"},
      "attributes": {"type": "object"}
    }
  }
}
```

### CSS Integration
```json
{
  "name": "suggest_css",
  "description": "Recommend CSS patterns for styling elements",
  "inputSchema": {
    "type": "object",
    "properties": {
      "tagName": {"type": "string"},
      "styleTarget": {"type": "string"}
    }
  }
}
```

## Usage Examples

### Basic Server Setup

```typescript
// Connect to MCP server
import { MCPClient } from '@modelcontextprotocol/client';

const client = new MCPClient({
  transport: 'stdio',
  command: 'cem',
  args: ['mcp']
});

await client.connect();
```

### Getting Element Information

```typescript
// Get schema information
const schema = await client.readResource('cem://schema');

// Get all available elements
const registry = await client.readResource('cem://registry');

// Get specific element details
const button = await client.readResource('cem://element/my-button');
```

### Generating HTML with AI Assistance

```typescript
// Generate HTML for a button component
const result = await client.callTool('generate_html', {
  tagName: 'my-button',
  attributes: {
    variant: 'primary',
    size: 'large'
  },
  content: {
    '': 'Click me!',
    'icon': '🚀'
  }
});

console.log(result.html);
// Output: <my-button variant="primary" size="large">
//           Click me!
//           <span slot="icon">🚀</span>
//         </my-button>
```

### Validating Component Usage

```typescript
// Validate attribute usage
const validation = await client.callTool('validate_element', {
  tagName: 'my-card',
  attributes: {
    elevation: '5',
    variant: 'outlined'
  },
  slots: {
    header: 'Card Title',
    '': 'Card content here'
  }
});

if (!validation.valid) {
  console.log('Validation errors:', validation.errors);
}
```

## Integration with AI Systems

### Claude Desktop

Add to your Claude Desktop configuration:

```json
{
  "mcpServers": {
    "cem": {
      "command": "cem",
      "args": ["mcp"],
      "env": {
        "CEM_WORKSPACE": "/path/to/your/project"
      }
    }
  }
}
```

### VSCode with Continue

```json
{
  "mcp": {
    "servers": [
      {
        "name": "cem",
        "command": "cem mcp",
        "workspaceRoot": true
      }
    ]
  }
}
```

### Custom Integration

```typescript
import { MCPServer } from '@modelcontextprotocol/server';
import { spawn } from 'child_process';

class CEMIntegration {
  private mcpProcess: any;
  
  async start(workspacePath: string) {
    this.mcpProcess = spawn('cem', ['mcp'], {
      cwd: workspacePath,
      stdio: ['pipe', 'pipe', 'pipe']
    });
    
    // Handle MCP communication
    return new MCPServer({
      transport: {
        read: this.mcpProcess.stdout,
        write: this.mcpProcess.stdin
      }
    });
  }
}
```

## Interface-Based Registry

The MCP server uses a sophisticated interface-based registry that provides type-safe access to manifest data:

### Interface Hierarchy

```
Item (base interface)
├── Typed (adds Type() method)
│   ├── Attribute (adds Default(), Required(), Values())
│   └── Event (adds isEvent() marker)
├── Slot (adds isSlot() marker)  
├── CssProperty (adds Syntax(), Inherits(), Initial())
├── CssPart (adds isCssPart() marker)
└── CssState (adds isCssState() marker)
```

### Benefits

- **Type Safety**: Compile-time guarantees for component data
- **Polymorphism**: Unified access with specialized interfaces
- **Extensibility**: Easy to add new item types
- **Performance**: Efficient caching and filtering
- **JSON Compatible**: Full API serialization support

## Advanced Configuration

### Workspace Detection

The MCP server automatically detects workspace boundaries and loads manifests from:

- `custom-elements.json` files
- `package.json` files with `customElements` field
- Nested component packages
- Monorepo structures

### Caching Strategy

```bash
# Enable persistent caching
cem mcp --cache-file .cem-cache.json

# Memory-only caching (default)
cem mcp --cache

# Disable caching for development
cem mcp --no-cache
```

### File Watching

```bash
# Watch for manifest changes (default)
cem mcp --watch

# Disable watching for production
cem mcp --no-watch

# Custom watch patterns
cem mcp --watch-pattern "**/*.json"
```

## Troubleshooting

### Common Issues

**Server won't start:**
```bash
# Check workspace has manifests
cem list

# Verify workspace structure
cem validate
```

**No elements found:**
```bash
# Generate manifests first
cem generate

# Check workspace context
ls custom-elements.json package.json
```

**Connection issues:**
```bash
# Test with debug mode
cem mcp --debug

# Try different transport
cem mcp --transport tcp --port 8080
```

### Debug Mode

Enable comprehensive logging:

```bash
cem mcp --debug
```

This provides:
- Manifest loading details
- Element conversion process
- Cache hit/miss information
- Tool invocation logs
- Resource access tracking

### Performance Monitoring

The server includes built-in performance monitoring:

```bash
# Enable metrics
cem mcp --metrics

# Custom metrics endpoint
cem mcp --metrics-port 9091
```

Metrics include:
- Element access patterns
- Cache efficiency
- Tool usage statistics
- Resource request timing
- Memory usage tracking

## See Also

- [Registry Interface Documentation](/docs/mcp/interfaces/) - Detailed interface reference
- [Usage Examples](/docs/mcp/examples/) - Comprehensive usage examples  
- [Architecture Overview](/docs/mcp/architecture/) - System design and patterns
- [API Reference](/docs/mcp/api/) - Complete API documentation