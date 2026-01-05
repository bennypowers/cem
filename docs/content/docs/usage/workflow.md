---
title: Development Workflow
weight: 20
---

Understanding the complete CEM development workflow will help you work efficiently with custom elements.

## The Core Cycle

CEM follows a manifest-driven development workflow with five key phases:

```
┌─────────────────────────────────────────────────────┐
│                                                     │
│  Write → Generate → Serve → Test → Edit            │
│    ↑                                       ↓        │
│    └───────────────────────────────────────┘        │
│                                                     │
└─────────────────────────────────────────────────────┘
```

### 1. Write

Create your custom element with JSDoc documentation:

```typescript
/**
 * A simple greeting component
 *
 * @summary Displays a personalized greeting message
 * @slot - Default slot for custom content
 * @csspart greeting - The greeting text container
 * @cssprop --greeting-color - Text color
 */
@customElement('hello-world')
export class HelloWorld extends LitElement {
  /**
   * The name to greet
   * @type {string}
   * @attr name
   */
  @property() name = 'World';

  render() {
    return html`
      <div part="greeting">
        Hello, ${this.name}!
        <slot></slot>
      </div>
    `;
  }
}
```

**Key points**:
- Use JSDoc for all public APIs
- Document attributes, properties, slots, events, CSS parts, and CSS custom properties
- Include examples and usage guidance
- Follow your design system's documentation standards

### 2. Generate

Run `cem generate` to create or update your manifest:

```bash
cem generate
```

This analyzes your code and produces `custom-elements.json` with metadata about your components.

**When to generate**:
- ✅ **After adding new components** - So they appear in tooling
- ✅ **After changing APIs** - When you add/remove properties, attributes, slots, etc.
- ✅ **Before committing** - Keep manifest in sync with code
- ⚠️ **Optional during dev** - `cem serve` can work without regenerating if you're just tweaking implementation

The manifest enables:
- LSP features (autocomplete, hover, validation)
- MCP integration (AI assistant access)
- Dev server component listing
- Documentation generation

### 3. Serve

Start the development server:

```bash
cem serve
```

The server provides:
- **Component listing** - Browse all elements
- **Live demos** - See components in action
- **Hot reload** - Changes refresh automatically
- **Buildless dev** - Write TypeScript without build steps
- **Import maps** - Use npm packages without bundling

The server runs at `http://localhost:8000` by default.

### 4. Test

Interact with your components in the browser:

**Manual Testing**:
- Try different property values
- Test user interactions (clicks, typing, etc.)
- Verify responsive behavior
- Check accessibility with screen readers
- Test in different browsers

**Interactive Knobs**:
Add knobs to your demos for easy property testing:

```html
<hello-world name="{{name}}"></hello-world>

<script type="module">
  import '@pwrs/cem/knobs';
</script>
```

See [Interactive Knobs](../knobs/) for details.

**Automated Testing**:
While the dev server is great for manual testing, also write automated tests:

```typescript
import { fixture, expect } from '@open-wc/testing';
import './hello-world.js';

it('displays greeting', async () => {
  const el = await fixture('<hello-world name="Test"></hello-world>');
  expect(el.shadowRoot.textContent).to.include('Hello, Test!');
});
```

### 5. Edit

Based on testing, make changes to your component:

**Quick fixes** (no manifest update needed):
- Styling changes
- Implementation details
- Internal logic
- Performance optimizations

**API changes** (requires manifest update):
- New properties, attributes, or slots
- Changed property types or defaults
- New events or CSS APIs
- Updated documentation

After API changes, regenerate the manifest:

```bash
cem generate
```

The dev server will detect the change and reload automatically.

## Integration Points

### LSP Integration

The Language Server Protocol uses your manifest to provide editor features:

**During Write phase**:
- Autocomplete when using your elements in HTML
- Hover documentation for attributes
- Validation of slot names and attribute values
- Go-to-definition from usage to source

**Setup**: See [LSP Integration](/docs/installation/lsp/)

### MCP Integration

The Model Context Protocol gives AI assistants access to your components:

**AI can**:
- Understand your component APIs
- Generate correct HTML with proper slots
- Suggest appropriate attribute values
- Validate component usage

**Setup**: See [MCP Integration](/docs/installation/mcp/)

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

**Skip regeneration when**:
- Tweaking styles
- Fixing implementation bugs
- Refactoring internal code
- Adding private methods

**Regenerate when**:
- Public API changes
- Documentation updates
- Before committing
- Before publishing

The `--watch` flag can help during active development:

```bash
cem generate --watch
```

### Demo-Driven Development

Create demos as you build:

1. **Start with demo** - Write the HTML you want to work
2. **Implement component** - Make the demo functional
3. **Refine demo** - Add edge cases and variations
4. **Document patterns** - Add descriptions and guidance

Demos become living documentation and test cases.

### Workspace Organization

Structure your project for efficient workflows:

```
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
│   └── cem.yaml               # CEM configuration
├── custom-elements.json       # Generated manifest
└── package.json
```

This structure:
- Keeps components self-contained
- Co-locates demos with components
- Makes navigation easy
- Works well with monorepos

## Common Workflows

### Adding a New Component

```bash
# 1. Create component file
mkdir -p elements/my-button
touch elements/my-button/my-button.ts

# 2. Write component with JSDoc

# 3. Generate manifest
cem generate

# 4. Create demo
mkdir elements/my-button/demo
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

## Troubleshooting

### LSP Features Not Working

1. **Check manifest exists**: `ls custom-elements.json`
2. **Verify manifest content**: `cem list`
3. **Regenerate if stale**: `cem generate`
4. **Check LSP logs** in your editor

### Dev Server Not Showing Components

1. **Regenerate manifest**: `cem generate`
2. **Check file patterns**: Verify `fileGlob` in config
3. **Verify package.json**: Ensure `customElements` field exists
4. **Check demos**: Ensure demo files exist and are discoverable

### Changes Not Refreshing

1. **Check live reload** is enabled (default)
2. **Look for errors** in browser console
3. **Verify source maps** for TypeScript
4. **Clear browser cache** if needed

## See Also

- **[Getting Started](../getting-started/)** - First project walkthrough
- **[Examples Overview](../examples/)** - Starter project templates
- **[Working with Demos](../demos/)** - Demo organization strategies
- **[Buildless Development](../buildless-development/)** - TypeScript without builds
- **[Using LSP Features](../using-lsp/)** - Editor integration tips
