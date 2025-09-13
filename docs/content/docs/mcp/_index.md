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
- Validate custom element usage patterns and manifest compliance

## Quick Start

```bash
# Start MCP server for current workspace
cem mcp
```

The server uses stdio transport and automatically discovers manifests from your workspace.

## MCP Resources

The server provides several resource types that AI systems can access:

| URI                       | Description                                                                                                                        |
| ------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `cem://schema`            | Returns the JSON schema for custom elements manifests, enabling validation and understanding of manifest structure.                |
| `cem://packages`          | Provides package discovery and overview of available manifest packages in the workspace.                                          |
| `cem://elements`          | Returns summaries of all available elements with their capabilities and basic metadata.                                           |
| `cem://element/{tagName}` | Returns detailed information about a specific element, including all attributes, slots, events, CSS properties, parts, and states. |
| `cem://guidelines`        | Provides design system guidelines and best practices for component usage.                                                          |
| `cem://accessibility`     | Returns accessibility patterns and validation rules for component compliance.                                                      |

## MCP Tools

The server provides interactive tools for AI assistance:

### HTML Validation
**Tool:** `validate_html`

Validates custom element usage based on manifest guidelines and best practices. Focuses on custom element-specific issues like slot content guidelines, attribute conflicts, and content/attribute redundancy.

| Parameter | Type   | Required | Description                                            |
| --------- | ------ | -------- | ------------------------------------------------------ |
| `html`    | string | ✅       | HTML content to validate                               |
| `tagName` | string |          | Focus validation on specific element                   |
| `context` | string |          | Validation context for custom elements                 |

**Validation Types:**
- **Slot Content Guidelines**: Validates slotted content against manifest slot descriptions
- **Attribute Conflicts**: Detects contradictory attribute combinations (e.g., `loading="eager"` + `lazy="true"`)
- **Content/Attribute Redundancy**: Identifies when slot content overrides attribute values
- **Manifest Compliance**: Ensures custom elements are used according to their documented constraints

### Attribute Suggestions
**Tool:** `suggest_attributes`

Get intelligent attribute suggestions with manifest context.

| Parameter | Type   | Required | Description                                                       |
| --------- | ------ | -------- | ----------------------------------------------------------------- |
| `tagName` | string | ✅       | Element to get attribute suggestions for                          |
| `context` | string |          | Usage context (`accessibility`, `form`, `interactive`, `styling`) |

### HTML Generation
**Tool:** `generate_html`

Generate correct HTML structure with proper slots and attributes.

| Parameter    | Type   | Required | Description                  |
| ------------ | ------ | -------- | ---------------------------- |
| `tagName`    | string | ✅       | Element to generate HTML for |
| `attributes` | object |          | Attribute values to include  |
| `content`    | object |          | Slot content mapping         |

### CSS Integration
**Tool:** `suggest_css_integration`

Provides CSS integration guidance using manifest-defined CSS APIs.

| Parameter     | Type   | Required | Description                                                |
| ------------- | ------ | -------- | ---------------------------------------------------------- |
| `tagName`     | string | ✅       | Element to get CSS guidance for                            |
| `styleTarget` | string |          | What to style (`element`, `parts`, `states`, `properties`) |
| `context`     | string |          | Styling context (`theme`, `responsive`, `dark-mode`)       |


## Usage Examples

### Basic Server Setup

```typescript
// Connect to MCP server
import { MCPClient } from '@modelcontextprotocol/client';

const client = new MCPClient({
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
const elements = await client.readResource('cem://elements');

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

console.log(result.content);
// Output: <my-button variant="primary" size="large">
//           Click me!
//           <span slot="icon">🚀</span>
//         </my-button>
```

### Validating HTML Usage

```typescript
// Validate HTML against manifests
const validation = await client.callTool('validate_html', {
  html: '<my-card elevation="5" variant="outlined">Card content</my-card>',
  tagName: 'my-card',
  context: 'manifest-compliance'
});

console.log('Validation results:', validation.content);
```

## Integration with AI Systems

### Claude Desktop

Add to your Claude Desktop configuration:

```json
{
  "mcpServers": {
    "cem": {
      "command": "cem",
      "args": ["mcp"]
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
        "command": "cem mcp"
      }
    ]
  }
}
```

### Custom Integration

```typescript
import { spawn } from 'child_process';

class CEMIntegration {
  private mcpProcess: any;
  
  async start(workspacePath: string) {
    this.mcpProcess = spawn('cem', ['mcp'], {
      cwd: workspacePath,
      stdio: ['pipe', 'pipe', 'pipe']
    });
    
    // Handle MCP communication via stdio
    return this.mcpProcess;
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

## Template-Driven Intelligence

The MCP server uses a unique **Data + Context + LLM Decision Making** approach:

1. **📊 Manifest Data** - Raw element definitions, CSS APIs, constraints
2. **📖 Rich Context** - Usage patterns, guidelines, best practices  
3. **🧠 LLM Decisions** - AI chooses appropriate values based on provided context

This means the server doesn't provide hardcoded suggestions but instead gives AI systems the rich context needed to make intelligent decisions about element usage, styling, and validation.

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
- Ensure the `cem mcp` command runs without errors
- Check that manifest files exist in your workspace
- Verify your AI client is configured for stdio transport

## See Also

- [Registry Interface Documentation](/docs/mcp/interfaces/) - Detailed interface reference
- [Usage Examples](/docs/mcp/examples/) - Comprehensive usage examples  
- [Architecture Overview](/docs/mcp/architecture/) - System design and patterns
- [API Reference](/docs/mcp/api/) - Complete API documentation
