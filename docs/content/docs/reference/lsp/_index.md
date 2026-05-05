---
title: Language Server
weight: 30
---

Language Server Protocol implementation for custom elements. See [LSP Integration](/docs/installation/lsp/) for editor setup and [Using LSP Features](/docs/usage/using-lsp/) for usage guidance.

## Supported LSP Methods

### Text Document Features
- `textDocument/hover` - Show element and attribute documentation on hover
- `textDocument/completion` - Provide tag and attribute completion suggestions
- `textDocument/definition` - Jump to custom element source definitions
- `textDocument/diagnostic` - Pull-based diagnostics (LSP 3.17); falls back to push via `textDocument/publishDiagnostics` for older clients
- `textDocument/codeAction` - Provide one-click autofixes for validation errors
- `textDocument/inlayHint` - Show attribute type annotations and slot biscuits inline
- `textDocument/didOpen` - Track when documents are opened in the editor
- `textDocument/didChange` - Handle incremental document changes
- `textDocument/didClose` - Clean up resources when documents are closed

### Workspace Features
- `workspace/symbol` - Search and navigate custom elements across the entire workspace
- `workspace/diagnostic` - Workspace-wide pull diagnostics (LSP 3.17)
- `workspace/didChangeConfiguration` - Update server settings at runtime

### Server Lifecycle
- `initialize` - Establish server capabilities and workspace configuration
- `shutdown` - Gracefully terminate the language server
- `$/setTrace` - Control debug logging verbosity (LSP standard)

## Configuration

### Settings

The LSP server reads settings from the `cem` namespace in `initializationOptions` and `workspace/didChangeConfiguration`. Settings can be changed at runtime without restarting the server.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `inlayHints` | `boolean` | `true` | Show inline type annotations for attributes and slot names |

#### VS Code Example

```json
{
  "cem.inlayHints": true
}
```

#### Neovim Example

```lua
settings = {
  cem = {
    inlayHints = true,
  },
}
```

### Initialization Options

The LSP server accepts custom initialization options for loading additional packages:

```json
{
  "initializationOptions": {
    "additionalPackages": [
      "npm:@rhds/elements@2.0.0",
      "https://cdn.jsdelivr.net/npm/@shortfuse/materialdesignweb/",
      "jsr:@example/elements"
    ],
    "cem": {
      "inlayHints": true
    }
  }
}
```

### VS Code Example

In your VS Code `settings.json`:

```json
{
  "cem.additionalPackages": [
    "npm:@rhds/elements@2.0.0",
    "https://cdn.jsdelivr.net/npm/@example/components/"
  ]
}
```

The VS Code extension passes these as `initializationOptions.additionalPackages` to the LSP server.

### Supported Specifiers

| Format | Example |
|--------|---------|
| npm specifier | `npm:@scope/package@version` |
| jsr specifier | `jsr:@scope/package` |
| CDN URL | `https://cdn.jsdelivr.net/npm/@scope/package/` |

**Note**: URLs must point to the package root (where `package.json` lives). The server reads the `customElements` field from `package.json` to locate the manifest.

## Protocol Features

### Pull Diagnostics (LSP 3.17)

When the client supports LSP 3.17 pull diagnostics (`textDocument/diagnostic`), the server uses the pull model instead of push notifications. This reduces server-initiated traffic and enables caching. Clients that do not advertise diagnostic capability automatically fall back to push via `textDocument/publishDiagnostics`.

### Inlay Hints

Inlay hints display inline annotations:

- **Attribute types**: Shows `: Boolean`, `: String`, etc. after attribute values
- **Slot biscuits**: Shows `slot: header` after the closing tag of slotted elements

Disable inlay hints by setting `cem.inlayHints` to `false` in your editor settings.

### Deprecated Elements and Attributes

Elements, attributes, and slots marked as `deprecated` in the manifest are reported with `DiagnosticTag.Deprecated`, rendering as strikethrough in supporting editors. If the manifest includes a deprecation reason, it appears in the diagnostic message.
