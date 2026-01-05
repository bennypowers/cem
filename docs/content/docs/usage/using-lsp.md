---
title: Using LSP Features
weight: 80
---

{{< tip >}}
**TL;DR**: The LSP provides autocomplete for tag names and attributes, hover docs on elements, and go-to-definition (<kbd>F12</kbd>) to jump to source code. Run `cem generate` to update your manifest and refresh editor features.
{{< /tip >}}

The [Language Server Protocol][lspspec] integration provides editor features like autocomplete, hover documentation, and go-to-definition for custom elements in HTML and TypeScript files. After [setting up the LSP][lspsetup], your editor reads your [Custom Elements Manifest][customelementsjson] to power intelligent code completion for tag names, attributes, slot names, and enum values. This complements the [development workflow][workflow] by catching typos and invalid attributes as you write code, reducing the time spent switching between editor and browser during the [test phase][testphase].

The LSP uses your manifest as the source of truth, so [running `cem generate`][generate] after documenting components makes those APIs immediately available in your editor. Features work in plain HTML files, Lit template literals, and anywhere custom elements appear in your code.

## Autocomplete

Press <kbd>Ctrl</kbd>+<kbd>Space</kbd> (or your editor's autocomplete trigger) after typing `<my-bu` to see custom element suggestions like `my-button` and `my-button-group` with their descriptions. Type a space after a tag name to see available attributes with type information and descriptions. For attributes with enum values like `variant`, autocomplete suggests valid options like `primary`, `secondary`, and `danger`. When adding `slot=""` attributes, autocomplete suggests valid slot names based on the parent element's documented slots.

The LSP works in Lit template literals with special syntax support—use `@eventName` for events, `.propertyName` for properties, and `?booleanAttr` for boolean attributes. All completions include inline documentation from your manifest.

## Hover Documentation

Hover over tag names to see element summaries, complete API documentation (properties, attributes, slots, events), CSS custom properties and parts, and links to source code. Hover over attributes to see descriptions, type information, default values, and valid enum values. In CSS files, hover over `::part()` selectors to see styling guidance for that shadow part.

## Go-to-Definition

Position your cursor on a tag name like `<my-button>` and press <kbd>F12</kbd> (VS Code) or <kbd>ctrl</kbd>-<kbd>]</kbd> (Neovim) to jump to the component source file. Works from attributes too—trigger go-to-definition on `variant="primary"` to jump to the property definition in the component class.

## Find References

Position your cursor on a custom element tag and press <kbd>Shift</kbd>+<kbd>F12</kbd> (VS Code) or <kbd>gr</kbd> (Neovim) to see all usages across HTML, TypeScript, and JavaScript files. Results are filtered by `.gitignore` to exclude `node_modules/`, show only start tags to avoid duplicates, and work in template literals.

## Workspace Symbols

Press <kbd>Ctrl</kbd>+<kbd>T</kbd> (VS Code) or use `:Telescope lsp_workspace_symbols` (Neovim) to search for custom elements across your entire workspace with fuzzy matching. Typing `btn` finds `my-button`, `icon-button`, and `button-group`.

## Error Detection & Quick Fixes

The LSP validates HTML and provides one-click fixes for common errors. Position your cursor on red squiggles and press <kbd>Ctrl</kbd>+<kbd>.</kbd> (VS Code) or <kbd>&lt;leader&gt;ca</kbd> (Neovim) to see available fixes.

Detected errors include invalid slot names (`slot="heade"` suggests `"header"`), typos in tag names (`<my-buttom>` suggests `<my-button>`), invalid attribute names (`varient` suggests `variant`), invalid enum values (`variant="primar"` suggests `"primary"`), and missing imports (suggests adding `import` statements for undeclared elements).

## Troubleshooting

If autocomplete doesn't work, check that `custom-elements.json` exists, verify the LSP is running in your editor's status bar, regenerate the manifest with `cem generate`, and restart your editor if needed.

If suggestions are outdated, regenerate the manifest—the LSP watches for changes and reloads automatically. For validation errors that don't appear, check that diagnostics are enabled in your editor and that your manifest contains element schemas.

For large projects with performance issues, limit workspace scope, exclude build directories in `.gitignore`, or enable verbose logging (`"cem.lsp.trace.server": "verbose"` in VS Code, `trace = 'verbose'` in Neovim) to diagnose what's happening.

## See Also

- **[LSP Integration][lspsetup]** - Setup instructions for editors
- **[LSP Protocol Reference][lspprotocol]** - Technical implementation details
- **[Development Workflow][workflow]** - How LSP fits into the dev cycle

[lspspec]: https://microsoft.github.io/language-server-protocol/
[lspsetup]: /docs/installation/lsp/
[customelementsjson]: https://github.com/webcomponents/custom-elements-json
[workflow]: ../workflow/
[testphase]: ../workflow/#4-test
[generate]: /docs/reference/commands/generate/
[lspprotocol]: /docs/reference/lsp/
