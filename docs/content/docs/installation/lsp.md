---
title: LSP Integration
weight: 20
---

{{< tip >}}
**TL;DR**: Install CEM, then configure your editor to run `cem lsp`. VS Code users can install the extension. Neovim/Zed users configure their LSP client to start `cem lsp` for HTML/TypeScript files.
{{< /tip >}}

Configure the CEM Language Server Protocol integration for your editor to get intelligent autocomplete, hover documentation, and validation for custom elements.

## Prerequisites

[Install CEM](/docs/installation/) first if you haven't already.

## What is LSP?

The Language Server Protocol provides a standard way to add language-specific features to any editor. The CEM language server analyzes your custom elements manifests to offer contextual autocomplete, hover documentation, go-to-definition, and other IDE enhancements for HTML and TypeScript files.

## Features

All editors with LSP support get the same core features:

- ✅ **Autocomplete** - Tag names, attributes, slot names, and values
- ✅ **Hover Documentation** - Element and attribute documentation on hover
- ✅ **Go-to-Definition** - Jump to element source code
- ✅ **Go-to-References** - Find all usages of elements across your workspace
- ✅ **Validation** - Real-time error detection for slots, attributes, and tag names
- ✅ **Quick Fixes** - One-click typo corrections and import suggestions
- ✅ **Workspace Symbols** - Search and navigate elements project-wide

See [Using LSP Features](/docs/usage/using-lsp/) for detailed usage guides.

## Editor Configuration

### Neovim

For Neovim 0.12+'s native LSP configuration support, create `~/.config/nvim/lsp/cem.lua`:

~~~lua
---https://bennypowers.dev/cem/
---
---`cem lsp` provides editor features for custom-elements a.k.a. web components
---
---Install with go
---```sh
---go install bennypowers.dev/cem
---```
---Or with NPM
---```sh
---npm install -g @pwrs/cem
---```
---
---@type vim.lsp.ClientConfig
return {
  cmd = { 'cem', 'lsp' },
  root_markers = { 'custom-elements.json', 'package.json', '.git' },
  filetypes = { 'html', 'typescript', 'javascript' },
  -- Control debug logging via LSP trace levels
  trace = 'off', -- 'off' | 'messages' | 'verbose'
}
~~~

### VS Code

Install the [Custom Elements Language Server extension](https://marketplace.visualstudio.com/items?itemName=pwrs.cem-language-server-vscode) from the VS Code marketplace. The extension bundles the language server and provides configuration options.

Configuration options in `settings.json`:

```json
{
  "cem.lsp.executable": "",             // Custom path (empty = use bundled)
  "cem.lsp.trace.server": "off"        // LSP trace level: "off" | "messages" | "verbose"
}
```

### Emacs

Depending on which LSP plugin you use, configure Emacs to run `cem` for HTML, JavaScript, and TypeScript files:

**lsp-mode:**
```elisp
(lsp-register-client
 (make-lsp-client
  :new-connection (lsp-stdio-connection '("cem" "lsp"))
  :major-modes '(html-mode typescript-mode js-mode)
  :server-id 'cem-lsp))
```

**eglot:**
```elisp
(add-to-list 'eglot-server-programs
             '((html-mode typescript-mode js-mode) . ("cem" "lsp")))
```

### Claude Code

Install CEM first if not already installed:

```bash
npm install -g @pwrs/cem # or
go install bennypowers.dev/cem@latest
```

Then install the plugin:

```text
/plugin marketplace add bennypowers/cem
/plugin install cem
```

The LSP activates automatically for HTML, TypeScript, and JavaScript files.

**Bonus**: The plugin also includes MCP server support. See [MCP Integration](../mcp/) for details.

### Other Editors

Configure your LSP client to run `cem lsp` for file types `html`, `typescript`, and `javascript`. The server communicates over stdio and follows the standard LSP specification.

Typical configuration elements:
- **Command**: `cem lsp`
- **File types**: `html`, `typescript`, `javascript`
- **Root markers**: `custom-elements.json`, `package.json`, `.git`
- **Transport**: stdio

## Verify Installation

After configuring your editor:

1. **Open an HTML file** in your project
2. **Start typing** a custom element tag name
3. **You should see** autocomplete suggestions for your custom elements

If autocomplete doesn't appear:
- Ensure you have a `custom-elements.json` in your project (run `cem generate`)
- Check your editor's LSP logs for errors
- Try enabling verbose logging (see configuration section)

## Debug Logging

Debug logging is controlled via the LSP standard `$/setTrace` notification. Most editors expose this through trace level settings:

- **`"off"`** - No debug logging (default)
- **`"messages"`** - Basic debug logging
- **`"verbose"`** - Detailed debug logging

Consult your editor's LSP plugin documentation for how to set trace levels.

## Next Steps

- **[Using LSP Features](/docs/usage/using-lsp/)** - Learn how to use autocomplete, hover, and validation
- **[LSP Protocol Reference](/docs/reference/lsp/)** - Technical API details
- **[Getting Started](/docs/usage/getting-started/)** - Create your first project
