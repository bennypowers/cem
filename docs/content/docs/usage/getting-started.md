---
title: Getting Started
weight: 10
---

This guide will walk you through creating your first custom element project with CEM.

## Prerequisites

- Node.js 16+ or Go 1.21+
- A code editor (VS Code, Neovim, etc.)
- Basic familiarity with custom elements / web components

## Choose Your Starting Point

CEM provides several example projects to help you get started. Choose based on your needs:

### For First-Time Users: Minimal Example

The **[minimal example](/docs/usage/examples/#minimal)** is the best starting point. It demonstrates:
- Single simple component
- Basic JSDoc documentation
- One property, slot, and CSS part
- Clear project structure

### For Realistic Projects: Intermediate Example

The **[intermediate example](/docs/usage/examples/#intermediate)** shows:
- Multiple related components
- Component composition patterns
- Custom events
- Demo discovery

### For Learning All Features: Kitchen Sink

The **[kitchen sink example](/docs/usage/examples/#kitchen-sink)** is a comprehensive reference:
- All CEM features in one component
- Design token integration
- Form integration
- Complete styling API

See the full **[Examples Overview](/docs/usage/examples/)** for all options.

## Quick Start from Scratch

If you prefer to start from scratch instead of using an example:

### 1. Install CEM

**For NPM projects**:
```sh
npm install --save-dev @pwrs/cem
```

**For Go projects**:
```sh
go install bennypowers.dev/cem@latest
```

See [Installation](/docs/installation/setup/) for more options.

### 2. Create Your First Component

Create a simple component with JSDoc documentation:

**`elements/hello-world/hello-world.ts`**:
```typescript
import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A simple greeting component
 *
 * @summary Displays a personalized greeting message
 * @slot - Default slot for custom content
 * @csspart greeting - The greeting text container
 * @cssprop --greeting-color - Text color (default: currentColor)
 */
@customElement('hello-world')
export class HelloWorld extends LitElement {
  static styles = css`
    :host {
      display: block;
      font-family: sans-serif;
    }

    #greeting {
      color: var(--greeting-color, currentColor);
      padding: 1rem;
    }
  `;

  /**
   * The name to greet
   * @type {string}
   * @attr name
   */
  @property() name = 'World';

  render() {
    return html`
      <div id="greeting" part="greeting">
        Hello, ${this.name}!
        <slot></slot>
      </div>
    `;
  }
}
```

### 3. Configure CEM

Create a configuration file:

**`.config/cem.yaml`**:
```yaml
input:
  glob: "elements/**/*.ts"

output:
  file: "custom-elements.json"

demoDiscovery:
  fileGlob: "elements/**/demo/*.html"
```

See [Configuration Reference](/docs/reference/configuration/) for all options.

### 4. Generate the Manifest

```sh
cem generate
```

This creates `custom-elements.json` with metadata about your component.

**Verify the output**:
```sh
cem list
```

You should see your `hello-world` component listed.

### 5. Create a Demo

Demos are HTML partials that showcase your component. The dev server wraps them in UI chrome automatically.

**`elements/hello-world/demo/index.html`**:
```html
<h2>Basic Usage</h2>
<hello-world></hello-world>

<h2>With Name</h2>
<hello-world name="CEM"></hello-world>

<h2>With Slot Content</h2>
<hello-world name="Developer">
  <p>Welcome to custom elements!</p>
</hello-world>

<h2>With Styling</h2>
<hello-world name="Styled" style="--greeting-color: blue;"></hello-world>

<script type="module">
  import '../../hello-world.js';
</script>
```

The dev server automatically discovers files in `demo/` directories.

### 6. Start the Dev Server

```sh
cem serve
```

The server starts on `http://localhost:8000` and opens your browser to the element listing page.

### 7. Test Your Component

In the browser:
1. **Click your component** in the listing
2. **View the demo** - See your component in action
3. **Inspect the docs** - Browse generated API documentation
4. **Make changes** - Edit code and see live reload

## Development Workflow

Now that you have a working project, follow the write → generate → serve → test → edit cycle:

1. **Write** - Add features to your component
2. **Generate** - Update the manifest: `cem generate`
3. **Serve** - View in browser: `cem serve`
4. **Test** - Interact with demos
5. **Edit** - Refine based on testing

See [Development Workflow](/docs/usage/workflow/) for detailed guidance.

## Features You Get

### Live Reload

Changes to your source files, demos, or manifest automatically refresh your browser. No manual reloads needed.

**Disable if needed**:
```sh
cem serve --no-reload
```

### Buildless Development

Write TypeScript and import CSS without build steps. The dev server transforms files on-demand.

See [Buildless Development](/docs/usage/buildless-development/) for details.

### Error Overlay

When errors occur, `cem serve` displays a developer-friendly overlay with:
- Full error message and stack trace
- Source-mapped locations (TypeScript)
- Syntax-highlighted code context

Click outside or press Escape to dismiss.

### Import Maps

Use npm packages without bundling:

```html
<script type="module">
  import { html } from 'lit';
  import '@shoelace-style/shoelace/dist/components/button/button.js';
</script>
```

See [Import Maps](/docs/usage/import-maps/) for details.

### Interactive Knobs

Add controls to test component variations:

```html
<hello-world name="{{name}}"></hello-world>

<script type="module">
  import '@pwrs/cem/knobs';
</script>
```

See [Interactive Knobs](/docs/usage/knobs/) for details.

## Next Steps

### Enhance Your Setup

- **[LSP Integration](/docs/installation/lsp/)** - Get autocomplete and validation in your editor
- **[MCP Integration](/docs/installation/mcp/)** - Enable AI assistant access to your components

### Learn More

- **[Development Workflow](/docs/usage/workflow/)** - Understanding the complete dev cycle
- **[Working with Demos](/docs/usage/demos/)** - Demo organization strategies
- **[Examples Overview](/docs/usage/examples/)** - Explore all starter projects

### Reference Documentation

- **[Commands](/docs/reference/commands/)** - CLI command reference
- **[Configuration](/docs/reference/configuration/)** - Complete config options
- **[Serve Command](/docs/reference/commands/serve/)** - Dev server options

## Troubleshooting

### Manifest is empty

Make sure your glob pattern matches your files:
```sh
cem generate --verbose
```

Check the `.config/cem.yaml` input glob pattern.

### Component not showing in dev server

1. Verify the manifest includes your component: `cem list`
2. Regenerate if needed: `cem generate`
3. Restart the server: `cem serve`

### Demos not discovered

Check demo file location matches discovery patterns:
- Files in `demo/` or `demos/` directories
- Files named `*.demo.html`
- Or configured in `.config/cem.yaml`

See [Working with Demos](/docs/usage/demos/) for details.

### TypeScript errors

Ensure you have a `tsconfig.json`:
```json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "ES2022",
    "moduleResolution": "bundler",
    "experimentalDecorators": true
  }
}
```

## Get Help

- **[Documentation](/docs/)** - Browse all docs
- **[GitHub Issues](https://github.com/bennypowers/cem/issues)** - Report bugs or request features
- **[Examples](https://github.com/bennypowers/cem/tree/main/examples)** - See working examples
