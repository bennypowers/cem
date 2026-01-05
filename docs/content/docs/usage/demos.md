---
title: Working with Demos
weight: 40
---

Demos are HTML files that showcase your custom elements in action. CEM provides powerful demo discovery and organization features.

## Demo Discovery

`cem serve` discovers demos in two ways:

### 1. Manifest Demos Field

**Recommended approach**: Explicitly list demos in your manifest using the `@demo` JSDoc tag:

```typescript
/**
 * My awesome button component
 *
 * @demo demos/basic.html Basic usage
 * @demo demos/variants.html All variants
 */
@customElement('my-button')
export class MyButton extends LitElement {
  // ...
}
```

When you run `cem generate`, these demos are added to the manifest and automatically linked to your component.

**Benefits**:
- Explicit control over demo ordering
- Custom descriptions for each demo
- Clear documentation of available demos
- Works in published packages

### 2. File Pattern Matching

**Automatic discovery**: CEM finds demo files matching these patterns:

- Files in `demo/` or `demos/` directories
- Files named `*.demo.html`

**Example structures**:

```
elements/
└── my-button/
    ├── my-button.ts
    └── demo/
        ├── index.html        # ✅ Discovered
        ├── variants.html     # ✅ Discovered
        └── advanced.html     # ✅ Discovered
```

```
elements/
└── my-card/
    ├── my-card.ts
    ├── basic.demo.html       # ✅ Discovered
    └── advanced.demo.html    # ✅ Discovered
```

**Configure patterns** in `.config/cem.yaml`:

```yaml
demoDiscovery:
  fileGlob: "elements/**/demo/*.html"  # Custom pattern
```

See [Configuration Reference](/docs/reference/configuration/) for all options.

## Demo File Format

Demos are **HTML partials** (not full HTML documents). The dev server wraps them in demo chrome automatically.

### Basic Demo

**`demo/index.html`**:
```html
<my-button>Click Me</my-button>

<script type="module">
  import '../my-button.js';
</script>
```

The server adds:
- HTML document wrapper
- Viewport meta tag
- Live reload script
- Error overlay
- Navigation UI

### Demo with Styles

```html
<style>
  my-button {
    --button-color: blue;
  }
</style>

<my-button>Styled Button</my-button>

<script type="module">
  import '../my-button.js';
</script>
```

### Demo with Multiple Elements

```html
<my-card>
  <my-card-header slot="header">
    Card Title
  </my-card-header>
  <my-card-body>
    Card content goes here
  </my-card-body>
</my-card>

<script type="module">
  import '../my-card.js';
  import '../my-card-header.js';
  import '../my-card-body.js';
</script>
```

## HTML5 Microdata

Use HTML5 microdata to control demo metadata and association.

### Explicit URLs

Declare the canonical URL for your demo:

```html
<meta itemprop="demo-url" content="/elements/my-button/demo/">

<my-button>Click Me</my-button>

<script type="module">
  import '../my-button.js';
</script>
```

### Rich Descriptions

Provide description with simple meta tag:

```html
<meta itemprop="description" content="Basic button demonstration">

<my-button>Click Me</my-button>
```

Or use markdown for rich descriptions:

```html
<script type="text/markdown" itemprop="description">
# Button Variants Demo

Showcases all button variants:
- **Primary** - Call to action buttons
- **Secondary** - Less prominent actions
- **Danger** - Destructive actions
</script>

<my-button variant="primary">Save</my-button>
<my-button variant="secondary">Cancel</my-button>
<my-button variant="danger">Delete</my-button>
```

### Explicit Element Association

Tell CEM which elements this demo showcases:

```html
<meta itemprop="demo-for" content="my-button pf-button">

<!-- Demo content -->
```

Useful when:
- Demo file path doesn't indicate element
- Demo showcases multiple elements
- Preventing incorrect auto-association

## Demo Association Logic

CEM uses this priority order to associate demos with elements:

1. **Explicit microdata** - `<meta itemprop="demo-for" content="element-name">`
2. **Path-based** - Elements whose aliases appear in demo file paths
3. **Content-based** - Custom elements found in the demo HTML

### Path-Based Association

When elements have configured aliases:

```yaml
aliases:
  my-button: "button"
  my-card: "card"
```

These paths match:

```
✅ elements/button/demo/basic.html  → my-button
✅ elements/card/demos/index.html   → my-card
❌ elements/btn/demo/index.html     → No match (alias is "button", not "btn")
```

### URLPattern-Based Matching

For precise path matching, configure a URL pattern:

```yaml
demoDiscovery:
  fileGlob: "shop/**/demos/*.html"
  urlPattern: "/shop/:element/:demo.html"
  urlTemplate: "https://mysite.com/shop/{{.element | alias}}/{{.demo}}/"
```

**Element Aliases**:
```yaml
aliases:
  my-shop-button: "shop-button"
  my-accordion: "accordion"
  my-accordion-header: "accordion-header"
```

**Matching behavior**:

```
✅ shop/shop-button/basic.html       → my-shop-button (alias in :element position)
❌ shop/my-shop-button/basic.html    → No match (literal name doesn't match alias)
✅ shop/accordion-header/demo.html   → my-accordion-header only
❌ shop/accordion-header/demo.html   → Does NOT match my-accordion (prevents false positives)
```

The URLPattern ensures aliases only match in their designated path positions, preventing shorter aliases from incorrectly matching longer path segments.

## URL Generation

Demo URLs are generated using this priority:

1. **Explicit microdata** - `<meta itemprop="demo-url">`
2. **URLPattern fallback** - Using `urlPattern` and `urlTemplate`
3. **No URL** - Demo is skipped if no pattern matches

### URL Templates

The `urlTemplate` uses Go template syntax with transformation functions:

```yaml
demoDiscovery:
  urlPattern: "/elements/:element/demo/:demo.html"
  urlTemplate: "https://example.com/{{.element | alias}}/{{.demo | slug}}/"
```

**Available functions**:

| Function | Description | Example |
|----------|-------------|---------|
| `alias` | Apply element alias mapping | `{{.tag \| alias}}` |
| `slug` | Convert to URL-friendly slug | `{{.demo \| slug}}` |
| `lower` | Convert to lowercase | `{{.component \| lower}}` |
| `upper` | Convert to uppercase | `{{.section \| upper}}` |

**Template examples**:

```yaml
# Basic interpolation
urlTemplate: "https://example.com/{{.component}}/{{.demo}}/"

# With alias transformation
urlTemplate: "https://example.com/{{.component | alias}}/{{.demo}}/"

# Chain multiple functions
urlTemplate: "https://example.com/{{.component | alias | slug}}/{{.demo | lower}}/"

# Function call syntax (alternative)
urlTemplate: "https://example.com/{{alias .component}}/{{slug .demo}}/"
```

**Important**: All transformations must be explicit. No automatic aliasing or slugification is applied unless requested in the template.

## Configuration Examples

### Minimal (Microdata-Driven)

```yaml
demoDiscovery:
  fileGlob: elements/**/demo/*.html
```

All URLs and descriptions come from microdata in demo files.

### URLPattern with Explicit Configuration

```yaml
demoDiscovery:
  fileGlob: elements/**/demo/*.html
  urlPattern: "/elements/:element/demo/:demo.html"
  urlTemplate: "https://site.com/components/{{.element | alias}}/demo/{{.demo | slug}}/"
```

Generates URLs from file paths using the template.

### Complex Multi-Site Example

```yaml
demoDiscovery:
  fileGlob: src/components/**/demos/*.html
  urlPattern: "/src/components/:component/demos/:variant.html"
  urlTemplate: "https://{{.component | alias | lower}}.examples.com/{{.variant | slug}}/"
```

Supports complex URL generation with multiple transformations.

## Demo Organization Strategies

### By Component (Recommended)

```
elements/
├── my-button/
│   ├── my-button.ts
│   └── demo/
│       ├── basic.html
│       ├── variants.html
│       └── states.html
└── my-card/
    ├── my-card.ts
    └── demo/
        ├── simple.html
        └── complex.html
```

**Benefits**:
- Co-located with component
- Easy to find and maintain
- Works well with automatic discovery

### By Feature

```
demos/
├── buttons/
│   ├── primary.html
│   ├── secondary.html
│   └── icon-buttons.html
└── forms/
    ├── login.html
    ├── registration.html
    └── validation.html
```

**Benefits**:
- Groups related demos
- Good for showcasing patterns
- Useful for design system documentation

Requires explicit association via microdata.

### Hybrid Approach

```
elements/
└── my-button/
    ├── my-button.ts
    └── demo/
        └── basic.html        # Component-specific

demos/
└── patterns/
    └── cta-buttons.html      # Cross-component pattern
```

**Benefits**:
- Component demos co-located
- Pattern demos separate
- Flexibility for different use cases

## Best Practices

### Write Focused Demos

Each demo should showcase one thing:

```html
<!-- ✅ Good: Focused on variants -->
<h2>Button Variants</h2>
<my-button variant="primary">Primary</my-button>
<my-button variant="secondary">Secondary</my-button>
<my-button variant="danger">Danger</my-button>
```

```html
<!-- ❌ Too much: Variants, sizes, states all mixed -->
<my-button variant="primary" size="large">Large Primary</my-button>
<my-button variant="secondary" size="small" disabled>Small Disabled</my-button>
<!-- Confusing: too many variables -->
```

### Use Descriptive Filenames

```
✅ variants.html
✅ loading-states.html
✅ form-integration.html

❌ demo1.html
❌ test.html
❌ example.html
```

### Add Context with Descriptions

```html
<script type="text/markdown" itemprop="description">
## When to Use This Pattern

Use primary buttons for the main call-to-action on a page.
Only use one primary button per section to avoid confusion.
</script>
```

### Include Usage Instructions

```html
<script type="text/markdown" itemprop="description">
## Interactive Demo

Try clicking the button to see the loading state.
The button is disabled while loading.
</script>

<my-button id="demo">Click Me</my-button>

<script type="module">
  import '../my-button.js';

  const button = document.getElementById('demo');
  button.addEventListener('click', async () => {
    button.loading = true;
    await new Promise(resolve => setTimeout(resolve, 2000));
    button.loading = false;
  });
</script>
```

## See Also

- **[Rendering Modes](../rendering-modes/)** - Light DOM, shadow DOM, iframe options
- **[Interactive Knobs](../knobs/)** - Add controls to demos
- **[Configuration Reference](/docs/reference/configuration/)** - Complete demo discovery config
- **[Development Workflow](../workflow/)** - How demos fit into the dev cycle
