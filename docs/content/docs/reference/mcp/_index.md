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
| `cem://config`                                  | Current resolved configuration (merged from file + workspace cascade)                               |
| `cem://config/schema`                           | Overview of all config schema sections with links to per-section detail                             |
| `cem://config/schema/{section}`                 | JSON schema for a specific config section (generate, serve, health, mcp, export, warnings)          |

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

### `generate_config`

Generate or update CEM configuration for a project. Returns the current config, config file path, full JSON schema, and guidance for each config section. Use with an optional `focus` parameter to narrow to a specific section.

| Parameter | Type   | Required | Description                                                                                              |
| --------- | ------ | -------- | -------------------------------------------------------------------------------------------------------- |
| `focus`   | string |          | Config section to focus on: `generate`, `serve`, `health`, `mcp`, `export`, or `warnings`. Omit for all sections. |

### `validate_config`

Validate the current CEM configuration file. Runs schema validation and semantic checks (glob reachability, output path existence, URL rewrite validation). Returns structured JSON with errors and warnings, each with field path and source position.

No parameters required.

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
- `--package <path>` - Specify project directory or package specifier (npm:, jsr:, or URL)
- `--additional-packages <specs>` - Load additional packages alongside local project (repeatable)
- `--max-description-length <num>` - Override 2000 character description limit
- `--verbose`, `-v` - Increase verbosity (`-v` info, `-vv` debug, `-vvv` trace)
- `--quiet`, `-q` - Quiet output (warnings and errors only)

### Loading Additional Packages

Load elements from external packages that aren't in your local project:

```bash
# Single additional package
cem mcp --additional-packages npm:@rhds/elements@2.0.0

# Multiple packages
cem mcp --additional-packages npm:@vaadin/button@24.3.5 \
        --additional-packages https://cdn.jsdelivr.net/npm/@example/components/
```

Or configure in `.config/cem.yaml`:

```yaml
additionalPackages:
  - npm:@rhds/elements@2.0.0
  - https://cdn.jsdelivr.net/npm/@shortfuse/materialdesignweb/
```
