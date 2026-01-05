---
title: Element Demos
weight: 40
---

Demos are HTML files that showcase your custom elements in action. They serve as
living documentation, manual test cases, and interactive examples for users
exploring your components. The `cem` dev server discovers demos from `@demo`
JSDoc tags (always works) or by file pattern matching (requires `demoDiscovery`
configuration), then wraps them in minimal chrome with live reload and
navigation UI. You can organize demos [by component][bycomponent], [by
feature][byfeature], or use a [hybrid approach][hybridapproach]—whatever makes
sense for your project.

Demos are HTML partials, not full documents. You write just the content you want 
to showcase, and the dev server handles the document wrapper, import maps, 
TypeScript transformation, and error overlay. Demos fit naturally into the 
[development workflow][developmentworkflow]—create them as you build components, 
use them during the [test phase][testphase], and update them when you edit APIs. 
Use [HTML5 microdata][html5microdata] to add descriptions, control URLs, and 
explicitly associate demos with elements. For advanced URL generation and 
path-based discovery, configure [URLPattern matching][urlpatternmatching] and 
[URL templates][urltemplates] in your `cem.yaml`.

## Demo Discovery

`cem` discovers demos in two ways: via JSDoc, and by configuring automatic demo
discovery in [cem.yaml][configurationreference].

### Manifest Demos Field

Explicitly list demos in your manifest using the `@demo` JSDoc tag:

```typescript
/**
 * My awesome button component
 *
 * @demo elements/my-button/demos/basic.html Basic usage
 * @demo elements/my-button/demos/variants.html All variants
 */
@customElement('my-button')
export class MyButton extends LitElement {
  // ...
}
```

When you run `cem generate`, these demos are added to the manifest and automatically linked to your component. This approach gives you explicit control over demo ordering and descriptions, provides clear documentation of available demos, and works seamlessly in published packages.

**Important**: Paths in `@demo` tags are relative to your project root, not to the component file.

### File Pattern Matching

**Recommended approach**: Configure `cem` to discover demos by glob pattern.

**Required configuration** in `.config/cem.yaml`:

```yaml
generate:
  demoDiscovery:
    fileGlob: elements/**/demo/*.html
```

Without this configuration, file-based discovery won't happen—only `@demo` JSDoc tags will work.

**Common patterns**:

```yaml
# Demos in demos/ subdirectories (recommended)
fileGlob: elements/**/demo/*.html

# Demos named *.demo.html
fileGlob: elements/**/*.demo.html
```

**Example structure with `fileGlob: "elements/**/demo/*.html"`**:

```
elements/
└── my-button/
    ├── my-button.ts
    └── demos/
        ├── index.html        # ✅ Matched by fileGlob
        ├── variants.html     # ✅ Matched by fileGlob
        └── advanced.html     # ✅ Matched by fileGlob
```

See [Configuration Reference][configurationreference] for all options.

### URL Generation with URLPattern

For advanced demo URL generation, configure [URLPattern][urlpattern] matching 
and URL templates. The `urlTemplate` uses Go template syntax to generate URLs 
from those captured parameters. See the [example projects][examples] for more.

```yaml
generate:
  demoDiscovery:
    fileGlob: elements/**/demo/*.html
    urlPattern: /elements/:element/demo/:demo.html
    urlTemplate: https://example.com/{{.element | alias}}/{{.demo | slug}}/
```

**Available template functions**:

| Function | Description | Example |
|----------|-------------|---------|
| `alias` | Apply element alias mapping | `{{.tag \| alias}}` |
| `slug` | Convert to URL-friendly slug | `{{.demo \| slug}}` |
| `lower` | Convert to lowercase | `{{.component \| lower}}` |
| `upper` | Convert to uppercase | `{{.section \| upper}}` |

**Template examples**:

```yaml
# Basic interpolation
urlTemplate: https://example.com/{{.component}}/{{.demo}}/

# With alias transformation
urlTemplate: https://example.com/{{.component | alias}}/{{.demo}}/

# Chain multiple functions
urlTemplate: https://example.com/{{.component | alias | slug}}/{{.demo | lower}}/
```

URLPattern also enables precise path-based demo association, preventing false matches when element aliases are substrings of each other (e.g., preventing "accordion" from matching "accordion-header").

## Writing Demos

Demos are **HTML partials** (not full HTML documents). The dev server wraps them in demo chrome automatically.

**`demos/index.html`**:

```html
<my-button>Click Me</my-button>

<script type="module">
  import '../my-button.js';
</script>
```

The server automatically wraps this partial in a full HTML document with viewport meta tag, live reload script, error overlay, and navigation UI. Include styles in your demo with an inline `<style>` element. We recommend scoping styles using CSS nesting to avoid your styles leaking out into the rest of the demo page.

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

You can include as many elements as you want in your demos:

```html
<my-card>
  <my-icon slot="header" icon="star"></my-icon>
  <h2 slot="header">Card Title</h2>
  <p>Card content goes here</p>
  <my-button slot="footer"></my-button>
</my-card>

<script type="module">
  import '../my-card.js';
  import '../my-button.js';
  import '../my-icon.js';
</script>
```

## Documenting Your Demos

Use HTML5 microdata to control demo metadata and association. You can modify the demo's URL, add a description, or explicitly associate that demo with a tag name.


```html
<meta itemprop="demo-url" content="/elements/my-button/demo/">
<meta itemprop="demo-for" content="my-button my-card">
<meta itemprop="description" content="Basic button demonstration">
```

The `demo-for` property is useful when the demo file path doesn't indicate which element it's for, when a demo showcases multiple elements, or when you need to prevent incorrect auto-association.

### Rich Descriptions

You can also add Markdown scripts to add rich content to your demo's description
```html
<script type="text/markdown" itemprop="description">
Showcases all button variants:
- **Primary** - Call to action buttons
- **Secondary** - Less prominent actions
- **Danger** - Destructive actions
</script>
```

## Demo Association Logic

CEM uses this priority order to associate demos with elements:

1. **Explicit microdata** - `<meta itemprop="demo-for" content="element-name">`
2. **Path-based** - Elements whose aliases appear in demo file paths
3. **Content-based** - Custom elements found in the demo HTML

### Path-Based Association

When elements use the `@alias` JSDoc tag, or have configured aliases in `cem.yaml`:

```yaml
aliases:
  my-button: button
  my-card: card
```

These paths match:

```
✅ elements/button/demo/basic.html → my-button
✅ elements/card/demo/index.html   → my-card
❌ elements/btn/demo/index.html    → No match (alias is "button", not "btn")
```

## See Also

- **[Rendering Modes][renderingmodes]** - Light DOM, shadow DOM, iframe options
- **[Interactive Knobs][interactiveknobs]** - Add controls to demos
- **[Configuration Reference][configurationreference]** - Complete demo discovery config workflow
- **[Development Workflow][developmentworkflow]** - How demos fit into the dev cycle

[devserver]: ../workflow/#3-serve
[bycomponent]: #by-component-recommended
[hybridapproach]: #hybrid-approach
[byfeature]: #by-feature
[developmentworkflow]: ../workflow/
[testphase]: ../workflow/#4-test
[html5microdata]: #html5-microdata
[urlpatternmatching]: #urlpattern-based-matching
[urltemplates]: #url-templates
[configurationreference]: /docs/reference/configuration/
[renderingmodes]: ../rendering-modes/
[interactiveknobs]: ../knobs/
[configurationreference]: /docs/reference/configuration/
[developmentworkflow]: ../workflow/
[urlpattern]: https://developer.mozilla.org/en-US/docs/Web/API/URL_Pattern_API
[examples]: https://github.com/bennypowers/cem/tree/main/examples/
