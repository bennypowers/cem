---
title: Getting Started
weight: 10
---

This guide will walk you through creating your first custom element project with 
`cem`.

## Prerequisites

- Node.js 24+ or Go 1.24+
- A code editor (VS Code, Neovim, etc.)
- Basic familiarity with custom elements / web components

## Quick Start

1. **[Install `cem`][installation]** - Get the tool installed
2. **Write** your component with JSDoc documentation
3. **Generate** the manifest: `cem generate`
4. **Create** a demo HTML file
5. **Serve** and test: `cem serve`

The **[Development Workflow][developmentworkflow]** guide provides detailed 
examples and explanations for each step.

### Enhance Your Setup

Use `cem`'s LSP and MCP servers to enhance your development environment.

- **[LSP Integration][lspintegration]** - Get autocomplete and validation in your editor
- **[MCP Integration][mcpintegration]** - Enable AI assistant access to your components

## Choose Your Starting Point

`cem` provides several example projects to help you get started. Choose based on 
your needs:

- **[minimal example][minimalexample]** with just a single, simple element,
- **[intermediate example][intermediateexample]** with multiple elements & demos,
- **["kitchen sink" example][kitchensinkexample]** with automatic design token 
  discovery, demos, etc.

See the full **[Examples Overview][examplesoverview]** for all options.

## The Development Workflow

`cem` uses a manifest-driven workflow: **Write → Generate → Serve → Test → Edit**. 
You write custom elements with JSDoc documentation, generate a manifest file, 
and use the dev server to view and test your components with live reload. The 
manifest powers LSP features in your editor and enables the dev server's 
automatic documentation and interactive controls.

For a complete walkthrough of each phase, see the **[Development 
Workflow][developmentworkflow]** guide.

## Features

The dev server provides:

- **[Live Reload][developmentworkflow]** - Changes automatically refresh your browser
- **[Buildless Development][buildlessdevelopment]** - Write TypeScript and import CSS without build steps
- **[Import Maps][importmaps]** - Use npm packages without bundling
- **[Interactive Knobs][interactiveknobs]** - Automatically generated controls for testing component variations
- **[Error Overlay][developmentworkflow]** - Developer-friendly error messages with source maps
- **[Rendering Modes][renderingmodes]** - Test in light DOM, shadow DOM, or chromeless mode

## Next Steps

### Reference Documentation

- **[Commands][commands]** - CLI command reference
- **[Configuration][configuration]** - Complete config options
- **[Serve Command][servecommand]** - Dev server options

Need help? See **[Troubleshooting][troubleshooting]** for common issues and solutions.

[minimalexample]: /docs/usage/examples/#minimal
[intermediateexample]: /docs/usage/examples/#intermediate
[kitchensinkexample]: /docs/usage/examples/#kitchen-sink
[examplesoverview]: /docs/usage/examples/
[installation]: /docs/installation/
[developmentworkflow]: /docs/usage/workflow/
[buildlessdevelopment]: /docs/usage/buildless-development/
[importmaps]: /docs/usage/import-maps/
[interactiveknobs]: /docs/usage/knobs/
[renderingmodes]: /docs/usage/rendering-modes/
[lspintegration]: /docs/installation/lsp/
[mcpintegration]: /docs/installation/mcp/
[commands]: /docs/reference/commands/
[configuration]: /docs/reference/configuration/
[servecommand]: /docs/reference/commands/serve/
[troubleshooting]: /docs/usage/troubleshooting/
