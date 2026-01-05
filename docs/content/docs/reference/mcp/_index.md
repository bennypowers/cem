---
title: MCP Server
weight: 40
---

Model Context Protocol server for custom elements. See [MCP Integration](/docs/installation/mcp/) for setup.

## MCP Resources

| URI                                             | Description                                                                                                                           |
| ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `cem://schema`                                  | JSON schema for custom elements manifests                   |
| `cem://packages`                                | Package discovery and overview of available manifest packages                                          |
| `cem://elements`                                | Summaries of all available elements with capabilities and metadata                                               |
| `cem://element/{tagName}`                       | Detailed element information including attributes, slots, events, CSS properties, parts, and states |
| `cem://element/{tagName}/attributes`            | Attribute documentation with type constraints, valid values, and usage patterns                                           |
| `cem://element/{tagName}/slots`                 | Content guidelines and accessibility considerations for slots                                                       |
| `cem://element/{tagName}/events`                | Event triggers, data payloads, and JavaScript integration patterns                                                                   |
| `cem://element/{tagName}/css/parts`             | CSS parts styling guidance                                                                    |
| `cem://element/{tagName}/css/custom-properties` | CSS custom properties documentation                                                           |
| `cem://element/{tagName}/css/states`            | CSS custom states documentation                                                                  |
| `cem://guidelines`                              | Design system guidelines and best practices                                                            |
| `cem://accessibility`                           | Accessibility patterns and validation rules                                                         |

## MCP Tools

### `generate_html`

Generate correct HTML structure with proper slots and attributes using manifest data.

| Parameter    | Type   | Required | Description                  |
| ------------ | ------ | -------- | ---------------------------- |
| `tagName`    | string | ✅       | Element to generate HTML for |
| `attributes` | object |          | Attribute values to include  |
| `content`    | object |          | Slot content mapping         |

### `validate_html`

Validates custom element usage based on manifest guidelines.

| Parameter | Type   | Required | Description                                            |
| --------- | ------ | -------- | ------------------------------------------------------ |
| `html`    | string | ✅       | HTML content to validate                               |
| `tagName` | string |          | Focus validation on specific element                   |
| `context` | string |          | Validation context for custom elements                 |

**Validation Types:**
- Slot content guidelines
- Attribute conflicts (e.g., `loading="eager"` + `lazy="true"`)
- Content/attribute redundancy
- Manifest compliance

## Configuration

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

Configure stdio transport with command `cem mcp`. See [MCP specification](https://spec.modelcontextprotocol.io/) for details.

## Server Options

```bash
cem mcp [flags]
```

**Flags:**
- `--package <path>` - Specify project directory
- `--max-description-length <num>` - Override 2000 character description limit
- `--verbose` - Enable detailed logging
