---
title: Language Server
weight: 30
---

Language Server Protocol implementation for custom elements. See [Using LSP Features](/docs/usage/using-lsp/) for setup and usage.

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

## Configuration

### Debug Logging

Debug logging is controlled via the LSP standard `$/setTrace` notification. Most editors expose this through trace level settings:

- **`"off"`** - No debug logging (default)
- **`"messages"`** - Basic debug logging
- **`"verbose"`** - Detailed debug logging

### Neovim

```lua
return {
  cmd = { 'cem', 'lsp' },
  root_markers = { 'custom-elements.json', 'package.json', '.git' },
  filetypes = { 'html', 'typescript', 'javascript' },
  trace = 'off', -- 'off' | 'messages' | 'verbose'
}
```

### VS Code

```json
{
  "cem.lsp.executable": "",             // Custom path (empty = use bundled)
  "cem.lsp.trace.server": "off"        // LSP trace level: "off" | "messages" | "verbose"
}
```

### Emacs (lsp-mode)

```elisp
(lsp-register-client
 (make-lsp-client
  :new-connection (lsp-stdio-connection '("cem" "lsp"))
  :major-modes '(html-mode typescript-mode js-mode)
  :server-id 'cem-lsp))
```

### Emacs (eglot)

```elisp
(add-to-list 'eglot-server-programs
             '((html-mode typescript-mode js-mode) . ("cem" "lsp")))
```

### Other Editors

Configure your LSP client to run `cem lsp` for file types `html`, `typescript`, and `javascript`. The server communicates over stdio and follows the standard LSP specification.
