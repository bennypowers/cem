---
title: Language Server
weight: 20
---

The CEM Language Server provides intelligent IDE features for custom elements in
HTML and TypeScript files. It analyzes your [custom elements manifest][manifest]
files to offer contextual autocomplete, hover documentation, and other editor
enhancements. CEM discovers all the custom elements manifests available to your
project (in your package.json), and watches for changes to your element sources
to keep an in-memory version of your local element manifest up to date.

## Features

### Autocomplete

Get up-to-date HTML element API completions in HTML files and Lit template
literals.

- **Tag name** completion for custom elements
- **Attribute** name completion based on element schemas
- **Attribute value** completion for known value sets i.e. TypeScript unions
- **Slot attribute** value completion for direct children of elements with slots
- **Event**, **DOM property**, and **Boolean attribute** support in lit templates.

### Hover Documentation

Get element API documentation close to where you use them with Hover docs.

- Element descriptions, properties, and usage examples
- Attribute documentation with type information
- Links to source code and external documentation

### Go-to-Definition

Quickly jump to your element, attribute, event, etc. definitions.

- Jump to custom element source definitions
- Support for tag names and attributes (extensible for slots and events)

### Go-to-References

Find all usages of custom elements across your entire workspace.

- **Place cursor on element**: Position cursor anywhere within a custom element tag name
- **Workspace-wide search**: Finds all references across HTML, TypeScript, and JavaScript files
- **Gitignore filtering**: Respects your project's `.gitignore` file to exclude irrelevant files like `node_modules/`
- **Clean results**: Shows only start tags to avoid duplicate entries for opening and closing tags
- **Template literal support**: Works within TypeScript `html` template literals and similar constructs

### Workspace Symbol Search

Search and navigate custom elements across your entire workspace.

- **Global element search**: Use your IDE's symbol search (Ctrl+T in VS Code) to find custom elements
- **Fuzzy matching**: Quickly locate elements with partial names and case-insensitive search
- **Direct navigation**: Click search results to jump immediately to element source files
- **Workspace-wide coverage**: Searches all custom elements from your manifests and dependencies

### Error Detection & Autofixes

Real-time validation with intelligent error correction.

- **Slot validation**: Detects invalid slot attribute values with smart suggestions
- **Tag name validation**: Validates custom element tag names with two intelligent error classes:
  - **Typo detection**: Suggests corrections for misspelled tag names (e.g., "my-buttom" → "my-button")
  - **Missing imports**: Detects elements that exist but aren't imported, with package-aware import suggestions
- **Attribute validation**: Validates HTML attributes using authoritative MDN browser-compat-data:
  - **Global attributes**: Validates against official HTML global attributes (class, id, data-*, aria-*, event handlers)
  - **Custom element attributes**: Validates against custom element manifest schemas with typo suggestions
  - **Standards-based**: Uses MDN browser-compat-data, automatically updated in CI/CD
- **One-click autofixes**: Automatically correct typos and add missing imports using your editor's quick fix feature
- **Smart import suggestions**: Resolves proper package names from package.json for npm packages vs. local modules
- **Intelligent error messaging**: Helpful guidance instead of overwhelming lists for large projects

## Quick Start

Install using Go:
```bash
go install bennypowers.dev/cem@latest
```
Or, using NPM:
```bash
npm install --global @pwrs/cem@latest
```

The language server communicates via the Language Server Protocol (LSP) over
standard input/output. Most modern editors support LSP through extensions or
built-in features.

## IDE Setup

### Neovim

For Neovim 0.12+'s native LSP configuration support, create
`~/.config/nvim/lsp/cem.lua`:

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

Install the [Custom Elements Language Server extension][vscode] from the VS Code
marketplace. The extension bundles the language server and provides
configuration options.

Configuration options in `settings.json`:

```json
{
  "cem.lsp.executable": "",             // Custom path (empty = use bundled)
  "cem.lsp.trace.server": "off"        // LSP trace level: "off" | "messages" | "verbose"
}
```

### Emacs

Depending on which LSP plugin you use, configure emacs to run cem for HTML,
javascript, and typescript files:

For **lsp-mode**:
```elisp
(lsp-register-client
 (make-lsp-client
  :new-connection (lsp-stdio-connection '("cem" "lsp"))
  :major-modes '(html-mode typescript-mode js-mode)
  :server-id 'cem-lsp))
```

For **eglot**:
```elisp
(add-to-list 'eglot-server-programs
             '((html-mode typescript-mode js-mode) . ("cem" "lsp")))
```

### Other Editors

Configure your LSP client to run `cem lsp` for file types `html`, `typescript`,
and `javascript`. The server communicates over stdio and follows the standard
LSP specification.

## Configuration

### Debug Logging

Debug logging is controlled via the LSP standard `$/setTrace` notification. Most editors expose this through trace level settings:

- **`"off"`** - No debug logging (default)
- **`"messages"`** - Basic debug logging
- **`"verbose"`** - Detailed debug logging

### Supported LSP Methods

**Text Document Features**
- `textDocument/hover` - Show element and attribute documentation on hover
- `textDocument/completion` - Provide tag and attribute completion suggestions
- `textDocument/definition` - Jump to custom element source definitions
- `textDocument/publishDiagnostics` - Report validation errors with intelligent suggestions
- `textDocument/codeAction` - Provide one-click autofixes for validation errors
- `textDocument/didOpen` - Track when documents are opened in the editor
- `textDocument/didChange` - Handle incremental document changes
- `textDocument/didClose` - Clean up resources when documents are closed

**Workspace Features**
- `workspace/symbol` - Search and navigate custom elements across the entire workspace

**Server Lifecycle**
- `initialize` - Establish server capabilities and workspace configuration
- `shutdown` - Gracefully terminate the language server
- `$/setTrace` - Control debug logging verbosity (LSP standard)

## Example: Slot Validation

The language server provides real-time validation for slot attributes with intelligent autofixes:

```html
<!-- ❌ Invalid slot name -->
<my-card>
  <div slot="heade">Title</div>  <!-- Red squiggles appear -->
</my-card>

<!-- ✅ After autofix (Ctrl+. → "Change 'heade' to 'header'") -->
<my-card>
  <div slot="header">Title</div>  <!-- Automatically corrected -->
</my-card>
```

When you type an invalid slot name, the language server:
1. **Detects the error** and shows red squiggles
2. **Suggests corrections** based on available slots from your manifest
3. **Provides one-click fixes** through your editor's quick fix menu

## Architecture
The server uses tree-sitter for robust parsing and maintains an in-memory index
of custom elements for fast completion and hover responses. LSP methods are
arranged into dedicated go packages for ease of maintenance.

[manifest]: https://github.com/webcomponents/custom-elements-manifest/
[vscode]: https://marketplace.visualstudio.com/items?itemName=pwrs.cem-language-server-vscode
