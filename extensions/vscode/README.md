# CEM Language Server for VS Code

IDE features for custom elements with intelligent autocomplete and hover documentation.

## Features

### ✨ Intelligent Autocomplete
- **Custom element tag names** with snippets including opening and closing tags
- **Element attributes** with type information and descriptions
- **Lit-specific bindings**: Event bindings (`@event`), property bindings (`.property`), boolean attributes (`?attribute`)
- **Context-aware completions** based on cursor position

### 🔍 Rich Hover Information
- **Element documentation** with comprehensive API details including attributes, events, and slots
- **Attribute documentation** with type information, descriptions, and default values
- **Stability during edits** - hover information remains accurate after document changes

### 📁 Automatic Manifest Discovery
- **Package.json integration** - automatically finds manifests from `customElements` field
- **Workspace manifests** from current project
- **Node modules scanning** for third-party custom elements
- **Real-time updates** when manifests change

## Quick Start

1. **Install this extension** from the VS Code Marketplace
2. **Open a project** with custom elements and a `custom-elements.json` manifest
3. **Start coding** - you'll get autocomplete and hover info for custom elements in HTML and TypeScript files

**That's it!** The extension includes the CEM language server, so no additional installation is required.

## Extension Settings

This extension contributes the following settings:

* `cem.lsp.debugLogging`: Enable debug logging for troubleshooting (default: `false`)
* `cem.lsp.executable`: Custom path to the CEM executable. Leave empty to use the bundled version (default: `""`)
* `cem.lsp.trace.server`: Trace communication with the language server (default: `"off"`)

## Commands

* `CEM: Restart Language Server`: Restart the CEM language server
* `CEM: Show Output Channel`: Show the language server output for debugging

## Supported File Types

- **HTML** (`.html`, `.htm`) - Full custom element support
- **TypeScript** (`.ts`) - Template literal support for `html\`\``, `element.innerHTML = \`\``, etc.
- **JavaScript** (`.js`, `.mjs`) - Same template literal support as TypeScript

## Configuration Example

```json
{
  "cem.lsp.debugLogging": true,
  "cem.lsp.executable": "",  // Use bundled version (default)
  "cem.lsp.trace.server": "verbose"
}
```

For advanced users who want to use a specific CEM version:
```json
{
  "cem.lsp.executable": "/usr/local/bin/cem"  // Custom installation
}
```

## Template Literal Support

The extension automatically detects custom elements in TypeScript/JavaScript template literals:

```typescript
// Supported patterns
const template = html`<my-element></my-element>`;
const generic = html<MyElement>`<my-element></my-element>`;
element.innerHTML = `<my-element></my-element>`;
element.outerHTML = `<my-element></my-element>`;
```

## Troubleshooting

### No Completions or Hover Information

1. **Verify your project has manifests**:
   ```bash
   # If you have CEM installed globally
   cem list tags  # Should show your custom elements
   ```

2. **Enable debug logging** in VS Code settings:
   ```json
   {
     "cem.lsp.debugLogging": true
   }
   ```

3. **Check the output panel** (View → Output → "CEM Language Server") for debug messages

### Language Server Won't Start

The extension uses a bundled version of CEM, so this should rarely happen. If it does:

1. **Check the output panel** for error messages
2. **Try restarting** with the "CEM: Restart Language Server" command
3. **For custom installations**: Ensure the `cem.lsp.executable` path is correct
4. **Report issues** on the [GitHub repository](https://github.com/bennypowers/cem/issues)

## Performance

The language server is optimized for performance with:
- Parser pooling for fast startup
- Incremental parsing for document updates
- Document caching with automatic cleanup
- Selective manifest loading

## Related Links

- [CEM Documentation](https://bennypowers.dev/cem/)
- [Custom Elements Manifest Specification](https://github.com/webcomponents/custom-elements-manifest)
- [GitHub Repository](https://github.com/bennypowers/cem/)

## License

GPL-3.0-or-later
