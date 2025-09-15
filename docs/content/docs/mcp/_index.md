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

### Security and Limits

For security and optimal AI performance, the server enforces a **2000 character limit** on description fields. Descriptions exceeding this limit are automatically truncated. This limit can be customized via configuration file or command line flags.

For best practices on writing effective descriptions within this limit, see [Writing Effective Descriptions](/docs/mcp/writing-descriptions/).

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

The server provides interactive tools built on the declarative framework. All element-focused tools (`element_*`) are now fully declarative, using YAML configuration and template rendering for consistent, rich responses:

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

### Element Information Tools (Declarative)

These tools use the declarative framework to provide focused, template-driven information:

**Tool:** `element_details`
Complete element reference with all APIs, schema definitions, and usage guidance.

**Tool:** `element_attributes`
Focused attribute documentation with type constraints, valid values, and usage patterns.

**Tool:** `element_slots`
Content guidelines and accessibility considerations for slot usage.

**Tool:** `element_events`
Event triggers, data payloads, and JavaScript integration patterns.

**Tool:** `element_styling`
CSS customization guidance including custom properties, parts, and states.

| Parameter | Type   | Required | Description                                           |
| --------- | ------ | -------- | ----------------------------------------------------- |
| `tagName` | string | ✅       | Element to get information for                        |
| `context` | string |          | Usage context for additional focused guidance         |

All declarative tools automatically combine your manifest descriptions with schema information, type constraints, and usage patterns to provide comprehensive AI guidance.


## Usage

Once connected, AI assistants can access your component information through natural language:

- **"What custom elements are available?"** - Lists all components
- **"Tell me about the my-button component"** - Shows attributes, slots, CSS APIs
- **"Generate a form with my-input and my-button"** - Creates valid HTML
- **"How do I style the my-card component?"** - Provides CSS guidance

## AI Client Configuration

### Claude Desktop

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

### Other Clients

The CEM MCP server works with any MCP-compatible AI client:

- **Claude Code** - See [Claude Code MCP docs](https://docs.anthropic.com/en/docs/claude-code/mcp)
- **Cursor** - See [Cursor MCP integration](https://docs.cursor.com/mcp)
- **Continue.dev** - See [Continue MCP setup](https://docs.continue.dev/mcp)
- **Custom clients** - See [MCP specification](https://spec.modelcontextprotocol.io/)

## How It Works

The MCP server uses a **declarative tool framework** with a **Data + Context + LLM Decision Making** approach:

### Declarative Architecture

The server uses a data-driven architecture where tools are defined through YAML configuration rather than hardcoded logic:

1. **📊 Data Fetchers** - Extract specific manifest data using JSON path queries
2. **📋 Template Rendering** - Combine data with Go templates for rich, contextual responses
3. **📖 Schema Integration** - Provide JSON schema context for AI understanding
4. **🧠 Intelligent Context** - Present comprehensive information while letting AI make smart decisions

### Core Philosophy

1. **📊 Manifest Data** - Your component definitions, CSS APIs, constraints
2. **📖 Rich Context** - Usage patterns, guidelines, best practices, schema definitions
3. **🧠 AI Decisions** - AI chooses appropriate values based on provided context

This means no hardcoded suggestions - just rich context that enables intelligent decisions about component usage and styling.

### Benefits of the Declarative Approach

- **🔧 Drop-in Tool Support** - New tools can be added with just YAML configuration
- **📝 Template-Driven Content** - Rich responses generated from your manifest data
- **⚡ Reduced Maintenance** - 77% less Go code to maintain
- **🎯 Focused Tools** - Specialized tools for attributes, slots, events, styling
- **🔄 Consistent Output** - All tools use the same data processing pipeline

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

- [Writing Effective Descriptions](/docs/mcp/writing-descriptions/) - Guidelines for AI-friendly component documentation, including declarative framework integration
- [Registry Interface Documentation](/docs/mcp/interfaces/) - Detailed interface reference
- [Usage Examples](/docs/mcp/examples/) - Comprehensive usage examples
- [Architecture Overview](/docs/mcp/architecture/) - System design and patterns
- [API Reference](/docs/mcp/api/) - Complete API documentation
