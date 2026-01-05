---
title: Editor Configuration
weight: 40
---

Editor-specific integration guides for using CEM with your development environment.

## Supported Editors

CEM integrates with most modern code editors through the Language Server Protocol (LSP). The following editors have been tested and documented:

### VS Code

**Extension**: [Custom Elements Language Server](https://marketplace.visualstudio.com/items?itemName=pwrs.cem-language-server-vscode)

1. Install the extension from the VS Code marketplace
2. The extension bundles the language server automatically
3. Configure in `settings.json` if needed:

```json
{
  "cem.lsp.executable": "",             // Custom path (empty = use bundled)
  "cem.lsp.trace.server": "off"        // LSP trace level
}
```

See [LSP Integration](../lsp/) for full setup details.

### Neovim

**Requirements**: Neovim 0.12+ with native LSP support

1. [Install CEM](../setup/) globally
2. Create `~/.config/nvim/lsp/cem.lua` with LSP configuration
3. The LSP client activates automatically for HTML, TypeScript, and JavaScript files

See [LSP Integration](../lsp/) for the complete configuration file.

### Emacs

**Supported plugins**: lsp-mode, eglot

1. [Install CEM](../setup/) globally
2. Configure your LSP plugin to run `cem lsp`
3. Register the server for `html-mode`, `typescript-mode`, and `js-mode`

See [LSP Integration](../lsp/) for plugin-specific configuration examples.

### Claude Code

**Plugin**: bennypowers/cem

1. [Install CEM](../setup/) globally
2. Install the plugin via Claude Code's plugin system
3. Both LSP and MCP activate automatically

See [LSP Integration](../lsp/) and [MCP Integration](../mcp/) for details.

### Other LSP-Compatible Editors

Any editor with LSP support can use CEM:

**Configuration requirements**:
- Command: `cem lsp`
- File types: `html`, `typescript`, `javascript`
- Root markers: `custom-elements.json`, `package.json`, `.git`
- Transport: stdio

Consult your editor's LSP plugin documentation for configuration instructions.

## Features by Editor

All editors with LSP support get the same core features:

- ✅ **Autocomplete** - Tag names, attributes, slot names
- ✅ **Hover Documentation** - Element and attribute docs
- ✅ **Go-to-Definition** - Jump to element source
- ✅ **Go-to-References** - Find element usage
- ✅ **Validation** - Real-time error detection
- ✅ **Quick Fixes** - One-click typo corrections
- ✅ **Workspace Symbols** - Search elements project-wide

## Configuration Tips

### Performance

For large projects, you may want to adjust LSP settings:

- Disable features you don't use
- Limit workspace symbol search scope
- Adjust debounce times for validation

### Debugging

Enable verbose logging to troubleshoot issues:

**VS Code**: Set `"cem.lsp.trace.server": "verbose"` in settings

**Neovim**: Set `trace = 'verbose'` in your LSP config

**Other editors**: Consult your LSP plugin docs for trace configuration

### Multiple Manifests

CEM automatically discovers all manifests in your workspace from package.json dependencies. No special configuration needed for monorepos or multi-package projects.

## Need Help?

If your editor isn't listed or you encounter issues:

1. Check that your editor supports LSP
2. Verify CEM is installed (`cem --version`)
3. Test the language server directly: `cem lsp`
4. Enable verbose logging and check for errors
5. Open an issue on [GitHub](https://github.com/bennypowers/cem/issues)

## See Also

- **[LSP Integration](../lsp/)** - Language server setup details
- **[MCP Integration](../mcp/)** - AI assistant integration
- **[Using LSP Features](/docs/usage/using-lsp/)** - How to use LSP features effectively
