---
title: Development Workflow
weight: 20
---

{{< tip >}}
**TL;DR**: The workflow is **Write → Generate → Serve → Test → Edit**. Write components with JSDoc, run `cem generate` when APIs change, use `cem serve` for live development. Skip regenerating for styling/implementation tweaks.
{{< /tip >}}

`cem` uses a manifest-driven development approach where you write custom
elements with JSDoc documentation, generate a JSON manifest from your source
code, and use that manifest to power developer tooling. The manifest enables 
[LSP features][lspfeatures] like autocomplete and validation in your editor, 
provides AI assistants with component information through [MCP][mcp], and drives 
the dev server's automatic documentation and [interactive 
controls][interactivecontrols]. This means your documentation isn't just 
comments—it becomes the foundation for your entire development experience.

The workflow follows a continuous cycle: write components with JSDoc, generate
the manifest to capture their APIs, serve them with hot reload for rapid
iteration, test with [interactive knobs][interactiveknobs] and manual 
validation, then edit based on feedback. For quick styling or implementation 
tweaks you can skip regenerating the manifest, but when you change public APIs 
(properties, slots, events, CSS parts) you'll regenerate to keep tooling in 
sync. The dev server provides [buildless TypeScript 
transformation][buildlesstypescripttransformation], automatic [import 
maps][importmaps] from package.json, and live reload, making the edit-test loop 
nearly instantaneous.

## The Core Cycle

`cem` follows a manifest-driven development workflow with five key phases:

{{< mermaid >}}
graph LR
    A[Write] --> B[Generate]
    B --> C[Serve]
    C --> D[Test]
    D --> E[Edit]
    E --> A
{{< /mermaid >}}

### 1. Write

Create your custom element with JSDoc documentation:

```typescript
import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * Displays a personalized greeting message
 *
 * @summary A simple greeting component
 */
@customElement('hello-world')
export class HelloWorld extends LitElement {
  static styles = css`
    :host {
      display: block;
      font-family: sans-serif;
    }

    #greeting {
      /** Text color (default: currentColor) */
      color: var(--greeting-color, currentColor);
      padding: 1rem;
    }
  `;

  /**
   * The name to greet
   */
  @property() name = 'World';

  render() {
    return html`
      <!-- The greeting text container -->
      <div id="greeting" part="greeting">
        Hello, ${this.name}!
        <!-- Default slot for custom content -->
        <slot></slot>
      </div>
    `;
  }
}
```

Use JSDoc to document all public APIs: attributes, properties, slots, events, CSS parts, and CSS custom properties. See [Documenting Components][documenting] for complete JSDoc tag reference and examples.

### 2. Generate

Run `cem generate` to create or update your manifest:

```bash
cem generate
```

This analyzes your code and produces `custom-elements.json` with metadata about your components. Regenerate after adding new components or changing public APIs (properties, attributes, slots, events, CSS APIs), and before committing to keep the manifest in sync. For quick styling or implementation tweaks during development, you can skip regeneration since `cem serve` works without it.

The manifest powers LSP features like autocomplete and validation, enables MCP integration for AI assistants, drives the dev server's component listing, and supports documentation generation.

### 3. Serve

Start the development server:

```bash
cem serve
```

The server provides a component listing to browse all elements, live demos to see components in action, hot reload for automatic refresh, buildless TypeScript development, and import maps for using npm packages without bundling. It runs at `http://localhost:8000` by default.

### 4. Test

Interact with your components in the browser by trying different property values, testing user interactions like clicks and typing, verifying responsive behavior, checking accessibility with screen readers, and testing across browsers. Use [interactive knobs][interactiveknobs] to tweak element attributes and properties in real-time.

### 5. Edit

Based on testing, make changes to your component. Quick fixes like styling changes, implementation details, internal logic, and performance optimizations don't need a manifest update. API changes—adding properties, attributes, slots, events, or CSS APIs, changing types or defaults, or updating documentation—require regenerating the manifest:

```bash
cem generate
```

The dev server will detect the change and reload automatically.

## Integration Points

### LSP Integration

The Language Server Protocol uses your manifest to provide autocomplete when using elements in HTML, hover documentation for attributes, validation of slot names and attribute values, and go-to-definition from usage to source. See [LSP Integration][lspintegration] for setup.

### MCP Integration

The Model Context Protocol gives AI assistants access to your components, allowing them to understand component APIs, generate correct HTML with proper slots, suggest appropriate attribute values, and validate component usage. See [MCP Integration][mcpintegration] for setup.

### CI/CD Integration

Automate manifest generation in your build pipeline:

```yaml
# GitHub Actions example
- name: Generate manifest
  run: |
    npm run analyze  # or: cem generate
    git diff --exit-code custom-elements.json
```

This ensures the manifest stays in sync with code changes.

## Best Practices

### Documentation-First Development

Write JSDoc **before** implementing features:

1. **Document the API** - Define properties, slots, events
2. **Generate manifest** - See how it looks
3. **Implement** - Write the component code
4. **Test** - Verify behavior matches documentation

This ensures documentation drives implementation, not the other way around.

### Incremental Regeneration

You don't need to regenerate the manifest after every change:

Skip regeneration when tweaking styles, fixing implementation bugs, refactoring internal code, or adding private methods. Regenerate for public API changes, documentation updates, before committing, and before publishing.

The `--watch` flag can help during active development:

```bash
cem generate --watch
```

{{< tip "info">}} you don't need to run `cem generate --watch` alongside the dev server, it 
rebuilds the manifest itself in-memory when your sources change.{{</ tip >}}

### Demo-Driven Development

Create demos as you build:

1. **Start with demo** - Write the HTML you want to work
2. **Implement component** - Make the demo functional
3. **Refine demo** - Add edge cases and variations
4. **Document patterns** - Add descriptions and guidance

Demos become living documentation and test cases.

### Workspace Organization

Structure your project with our recommended layout for efficient workflows:

```text
my-components/
├── elements/
│   ├── hello-world/
│   │   ├── hello-world.ts      # Component
│   │   ├── hello-world.css     # Styles
│   │   └── demo/
│   │       ├── basic.html      # Basic demo
│   │       └── variants.html   # Variations
│   └── my-card/
│       ├── my-card.ts
│       └── demo/
│           └── index.html
├── .config/
│   └── cem.yaml               # `cem` configuration
├── custom-elements.json       # Generated manifest
└── package.json
```

This structure keeps components self-contained, co-locates demos with components, makes navigation easy, and works well with monorepos.

## Common Workflows

### Adding a New Component

```bash
# 1. Create component file
mkdir -p elements/my-button/demo/
touch elements/my-button/my-button.ts

# 2. Write component with JSDoc

# 3. Generate manifest
cem generate

# 4. Create demo
touch elements/my-button/demo/index.html

# 5. Start dev server
cem serve

# 6. Test and iterate
```

### Updating Component APIs

```bash
# 1. Update JSDoc and code

# 2. Regenerate manifest
cem generate

# 3. Verify in dev server (refreshes automatically)
# Already running: cem serve

# 4. Update demos if needed

# 5. Test changes
```

### Publishing Components

```bash
# 1. Ensure manifest is current
cem generate

# 2. Run tests
npm test

# 3. Build if needed
npm run build

# 4. Publish
npm publish
```

The manifest is included in your package and enables LSP/MCP for consumers.

## See Also

- **[Getting Started](../getting-started/)** - First project walkthrough
- **[Documenting Components][documenting]** - JSDoc usage guide and examples
- **[Examples Overview](../examples/)** - Starter project templates
- **[Working with Demos](../demos/)** - Demo organization strategies
- **[Buildless Development](../buildless-development/)** - TypeScript without builds
- **[Using LSP Features](../using-lsp/)** - Editor integration tips
- **[Troubleshooting](../troubleshooting/)** - Common issues and solutions

[documenting]: ../documenting-components/
[interactiveknobs]: ../knobs/
[lspintegration]: /docs/installation/lsp/
[mcpintegration]: /docs/installation/mcp/
[lspfeatures]: /docs/installation/lsp/
[mcp]: /docs/installation/mcp/
[interactivecontrols]: ../knobs/
[buildlesstypescripttransformation]: ../buildless-development/
[importmaps]: ../import-maps/
