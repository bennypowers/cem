---
title: MCP Integration
weight: 30
---

Configure the CEM Model Context Protocol (MCP) server to give AI assistants intelligent access to your custom elements manifests.

## Prerequisites

[Install CEM](../setup/) first if you haven't already.

## What is MCP?

The Model Context Protocol enables AI assistants to access structured data about your custom elements. The CEM MCP server transforms your manifests into an AI-accessible interface for:

- Generating correct HTML with proper slots and attributes
- Understanding component APIs and relationships
- Providing CSS integration guidance
- Validating custom element usage patterns

## Quick Start

Test that MCP is working:

```bash
cem mcp
```

The server starts and communicates via stdio. Press Ctrl+C to stop.

## AI Client Configuration

### Claude Desktop

Add CEM to your Claude Desktop configuration file:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

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

Restart Claude Desktop for changes to take effect.

### Claude Code

Install the CEM plugin which includes both LSP and MCP support:

```text
/plugin marketplace add bennypowers/cem
/plugin install cem
```

The MCP server activates automatically when the plugin is installed.

### Cursor

Add CEM to your Cursor MCP configuration. See [Cursor MCP integration docs](https://docs.cursor.com/mcp) for configuration file location and format.

### Continue.dev

Add CEM to your Continue MCP setup. See [Continue MCP setup docs](https://docs.continue.dev/mcp) for configuration instructions.

### Other MCP Clients

The CEM MCP server works with any MCP-compatible AI client that supports stdio transport. Configure your client to run:

```bash
cem mcp
```

See the [MCP specification](https://spec.modelcontextprotocol.io/) for details on client implementation.

## Verify Installation

After configuring your AI client:

1. **Ask the AI**: "What custom elements are available in my project?"
2. **The AI should** list your components from the manifest
3. **Try generation**: "Generate a card component with my-card element"

If the AI can't access your components:
- Ensure you have a `custom-elements.json` in your project (run `cem generate`)
- Restart your AI client after configuration changes
- Check that `cem mcp` runs without errors when executed directly

## Debugging

### MCP Inspector

Use the [MCP Inspector](https://github.com/modelcontextprotocol/inspector) for interactive debugging:

```bash
# Install MCP Inspector
npm install -g @modelcontextprotocol/inspector

# Debug CEM MCP server
mcp-inspector cem mcp
```

The inspector provides a web UI for:
- Testing resources and tools interactively
- Viewing real-time server logs
- Validating resource URIs and tool parameters
- Debugging protocol communication

### Verbose Logging

Enable detailed logging to troubleshoot issues:

```bash
cem mcp --verbose
```

Or with the inspector:

```bash
mcp-inspector cem mcp --verbose
```

### Debug Specific Projects

Use the `--package` flag to debug different project directories:

```bash
mcp-inspector cem mcp --package /path/to/project
```

## Common Issues

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
- Ensure `cem mcp` runs without errors
- Check that manifest files exist in your workspace
- Verify your AI client is configured for stdio transport
- Restart your AI client after configuration changes

## Next Steps

- **[Effective MCP Descriptions](/docs/usage/effective-mcp-descriptions/)** - Write AI-friendly documentation
- **[MCP Protocol Reference](/docs/reference/mcp/)** - Technical API details
- **[Getting Started](/docs/usage/getting-started/)** - Create your first project
