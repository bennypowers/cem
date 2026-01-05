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
- `textDocument/publishDiagnostics` - Report validation errors with intelligent suggestions
- `textDocument/codeAction` - Provide one-click autofixes for validation errors
- `textDocument/didOpen` - Track when documents are opened in the editor
- `textDocument/didChange` - Handle incremental document changes
- `textDocument/didClose` - Clean up resources when documents are closed

### Workspace Features
- `workspace/symbol` - Search and navigate custom elements across the entire workspace

### Server Lifecycle
- `initialize` - Establish server capabilities and workspace configuration
- `shutdown` - Gracefully terminate the language server
- `$/setTrace` - Control debug logging verbosity (LSP standard)
