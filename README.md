# `cem` - Custom Elements Manifest Multitool

**cem** is a command-line tool for generating and querying
[Custom Elements Manifest][cem] files. It can analyze your codebase and generate
rich metadata for your custom elements, facilitating documentation, tooling, and
integration. It can also query that manifest for information about your package

## Installation

```sh
npm install --save-dev @pwrs/cem
```

For more options, see [Installation docs][installationdocs]

## Features

### `cem generate`

> [!NOTE]
> `cem generate` best supports LitElements written in idiomatic style with
> TypeScript decorators. There is rudimentary support for `extends HTMLElement`,
> but it is not a high priority for development. If you need something more
> specific [open an issue][issuenew].

See more in the [Generate docs][generatedocs]

---

### `cem serve`

The `cem serve` command starts a development server specifically for custom element development. It provides an opinionated, manifest-driven workflow with live reload, interactive knobs for testing attributes and properties, and a PatternFly-designed UI for component isolation.

**Features:**
- Live reload with WebSocket-based updates
- Interactive knobs auto-generated from manifest
- TypeScript transformation on-the-fly
- Automatic import map generation
- npm workspaces support for monorepos
- Demo discovery from manifest `demos` field
- Multiple rendering modes: full UI, shadow DOM, or chromeless (for testing/embedding)

**Rendering Modes:**
- **Light** (default): Full PatternFly UI with sidebar, knobs, logs, and event monitoring
- **Shadow**: Same UI as light, but renders demos in Shadow DOM for testing shadow root behavior
- **Chromeless**: Minimal HTML with live reload only - no UI chrome. Perfect for Playwright tests, isolated development, embedding in docs, or capturing clean screenshots

```bash
# Full UI with development tools
cem serve

# Shadow DOM testing
cem serve --rendering=shadow

# Chromeless for automated testing
cem serve --rendering=chromeless
```

See more in the [Serve docs][servedocs]

---

### `cem lsp`

The `cem lsp` command starts a Language Server Protocol (LSP) server that provides intelligent IDE features for custom elements in HTML and TypeScript files. It offers contextual autocomplete, hover documentation, and other editor enhancements by analyzing your custom elements manifests.

**Features:**
- Tag name and attribute completion for custom elements
- Slot attribute value completion for direct children of slotted elements  
- Hover documentation with type information
- **Error detection with autofixes** - Real-time validation with one-click corrections:
  - Slot validation with smart suggestions
  - Tag name validation with typo detection and missing import suggestions
  - Attribute validation against HTML global attributes and custom element schemas
  - Attribute value validation against manifest type definitions (union types, literals, numbers, booleans)
- Go-to-definition support for jumping to element source code
- Go-to-references to find all usages of custom elements across your workspace
- Support for HTML files and TypeScript template literals
- Automatic manifest discovery and live reloading

**IDE Extensions:**
- [VSCode](extensions/vscode/)
- [Zed](extensions/zed/)
- [Claude Code](extensions/claude-code/)

See more in the [LSP docs][lspdocs]

---

### `cem mcp`

The `cem mcp` command starts a Model Context Protocol (MCP) server that provides AI-native access to your custom elements manifest data. This enables intelligent HTML generation, component understanding, and design system compliance for AI assistants.

**Features:**
- **Resources**: Access to schemas, package discovery, element summaries, and accessibility patterns
- **Tools**: HTML validation, attribute suggestions, HTML generation, and CSS integration guidance  
- **Cross-package discovery**: Multi-manifest support for complex design systems

The server transforms your custom elements manifests into structured, actionable context that AI systems can use to generate proper HTML with correct slot usage, appropriate attributes, and design system compliance.

See more in the [MCP docs][mcpdocs]

---

### `cem list`

The `cem list` command provides a fast, flexible way to inspect custom elements, their features, and their metadata directly from your manifest file.
With `cem list`, you can quickly explore and audit your custom elements API surface, making it easier to document, test, and share your components.

See more in the [List docs][listdocs]

---

### `cem search`

The `cem search` command allows you to search through your custom elements manifest for any element by keyword or regex pattern. Search through names, descriptions, summaries, and labels of all manifest items including tags, modules, attributes, slots, CSS properties, CSS states, CSS parts, events, methods, demos, functions, variables, and more.

See more in the [Search docs][searchdocs]

---

### `cem validate`

The `cem validate` command validates your `custom-elements.json` file against the official JSON schema and provides intelligent warnings for potentially inappropriate manifest content. Beyond basic schema validation, it analyzes your manifest for lifecycle methods, private implementations, and other patterns that shouldn't be part of your public API documentation.

See more in the [Validate docs][validatedocs]

See the [Configuration Reference][configdocs] for more information.

---

## Examples

The `examples/` directory contains fully-runnable example projects that demonstrate CEM features at different complexity levels. Each example serves as both test fixtures and user-facing documentation.

- [Minimal][minimal]: the simplest possible `<hello-world>` custom element - perfect for getting started. **Start here** if you're new to CEM
- [Vanilla][vanilla]: web components without frameworks - pure JavaScript and DOM.
- [Intermediate][intermediate]: real-world UI component library; multiple components, demo discovery with HTML microdata
- [Kitchen Sink][kitchensink]: comprehensive feature showcase with demos and design tokens integration
- [TypeScript Paths][typescriptpaths]: various build configuration patterns like `src/` → `dist/`

See each example project's README for more info.

---

## Contributing

For information on building and testing, please see
[CONTRIBUTING.md][contributingmd].

## License

This program is free software: you can redistribute it and/or modify it under
the terms of the [GNU General Public License v3.0][gpl3].

&copy; 2025 Benny Powers

[cem]:              https://github.com/webcomponents/custom-elements-manifest
[dtcg]:             https://tr.designtokens.org/format/
[go]:               https://go.dev
[treesitter]:       https://tree-sitter.github.io/tree-sitter/
[gpl3]:             https://www.gnu.org/licenses/gpl-3.0.html
[issuenew]:         https://github.com/bennypowers/cem/issues/new
[contributingmd]:   https://bennypowers.dev/cem/docs/contributing/
[installationdocs]: https://bennypowers.dev/cem/docs/installation/
[generatedocs]:     https://bennypowers.dev/cem/docs/commands/generate/
[listdocs]:         https://bennypowers.dev/cem/docs/commands/list/
[searchdocs]:       https://bennypowers.dev/cem/docs/commands/search/
[validatedocs]:     https://bennypowers.dev/cem/docs/commands/validate/
[servedocs]:        https://bennypowers.dev/cem/docs/commands/serve/
[lspdocs]:          https://bennypowers.dev/cem/docs/lsp/
[mcpdocs]:          https://bennypowers.dev/cem/docs/mcp/
[configdocs]:       https://bennypowers.dev/cem/docs/configuration/
[minimal]:          examples/minimal/
[vanilla]:          examples/vanilla/
[intermediate]:     examples/intermediate/
[kitchensink]:      examples/kitchen-sink/
[typescriptpaths]:  examples/typescript-paths/
