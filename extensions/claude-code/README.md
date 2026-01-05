# CEM for Claude Code

Custom Elements Manifest tooling for Claude Code, providing both Language Server Protocol (LSP) and Model Context Protocol (MCP) support for custom elements.

## Features

### LSP Features (Editor Intelligence)

- **Tag Completion**: Auto-complete custom element tag names in HTML
- **Attribute Completion**: Intelligent attribute suggestions based on element definitions
- **Hover Documentation**: View element and attribute documentation on hover
- **Go-to-Definition**: Navigate to element class definitions
- **Diagnostics**: Real-time validation and error reporting
- **Code Actions**: Quick fixes for missing imports and common issues

### MCP Features (AI-Native Component Understanding)

- **Component Discovery**: Find available custom elements in your project
- **Attribute Validation**: Get accurate attribute names, types, and values
- **HTML Generation**: AI-generated valid HTML with proper component usage
- **Documentation Access**: Access component descriptions and usage guidelines
- **Design System Compliance**: Automatic adherence to your design tokens

## Installation

### 1. Install CEM Binary

Choose your preferred method:

**Via npm:**
```bash
npm install -g @pwrs/cem
```

**Via Go:**
```bash
go install bennypowers.dev/cem@latest
```

Verify installation:
```bash
cem version
```

### 2. Install Claude Code Plugin

Add this marketplace to Claude Code:
```
/plugin marketplace add bennypowers/cem
```

Then install the plugin:
```
/plugin install cem
```

## Usage

### LSP (Editor Intelligence)

The LSP activates automatically for:
- HTML files (`.html`, `.htm`)
- TypeScript files (`.ts`, `.tsx`)
- JavaScript files (`.js`, `.mjs`, `.cjs`)

Simply open a file with custom elements and start typing!

### MCP (AI Understanding)

The MCP server activates automatically when the plugin is installed. Use it by asking questions like:

- "What custom elements are available in this project?"
- "Tell me about the `my-button` component"
- "Create a form with my-input and my-button components"

The AI will have direct access to your component manifests and can generate accurate HTML using your design system.

## Configuration

The plugin works out-of-the-box with zero configuration. CEM automatically:
- Discovers `custom-elements.json` manifests in your project
- Scans `node_modules` for component libraries
- Provides intelligent completions and diagnostics

### Custom Configuration

Create a `.config/cem.yaml` in your project root for advanced options:

```yaml
# Exclude patterns
exclude:
  - "**/*.spec.ts"
  - "dist/**"

# Additional manifest paths
manifests:
  - "./custom-manifests/components.json"
```

## Troubleshooting

### LSP Not Starting

1. Verify CEM is installed and in PATH:
   ```bash
   which cem
   cem lsp --help
   ```

2. Check Claude Code logs for errors:
   ```
   /logs
   ```

3. Restart the LSP server:
   ```
   /lsp restart
   ```

### No Completions Appearing

1. Ensure your project has a `custom-elements.json` manifest:
   ```bash
   cem generate
   ```

2. Check that the manifest is valid:
   ```bash
   cem validate
   ```

3. Verify component libraries in `node_modules` have manifests

### Performance Issues

For large projects, exclude unnecessary files:

```yaml
# .config/cem.yaml
exclude:
  - "node_modules/**"
  - "dist/**"
  - "**/*.test.ts"
```

### MCP Not Working

1. Verify the MCP server is loaded:
   ```
   /mcp
   ```

2. Check that CEM is installed:
   ```bash
   which cem
   cem mcp --help
   ```

3. Restart Claude Code to reload the MCP server

## Support

- **LSP Documentation**: https://bennypowers.dev/cem/docs/reference/lsp
- **MCP Documentation**: https://bennypowers.dev/cem/docs/reference/mcp
- **Issues**: https://github.com/bennypowers/cem/issues
- **Discussions**: https://github.com/bennypowers/cem/discussions

## License

GPL-3.0 - See [LICENSE](../../LICENSE) for details
