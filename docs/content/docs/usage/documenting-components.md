---
title: Documenting Components
weight: 25
---

{{< tip >}}
**TL;DR**: Use JSDoc tags (`@slot`, `@csspart`, `@cssprop`, `@fires`, `@attr`) to document your components. These power autocomplete, validation, AI assistants, and interactive controls. See examples below.
{{< /tip >}}

Use JSDoc comments to document your custom elements for the [manifest][workflow]. The manifest powers [LSP features][lsp] like autocomplete and validation, enables [AI assistants][mcp] to understand your components, and drives the dev server's [interactive controls][knobs].

## JSDoc Tags

Use these tags in your element class and member JSDoc comments. See the [generate command reference][generateref] for the complete list.

### Basic Example

```typescript
import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * Displays a personalized greeting message
 *
 * @summary A simple greeting component
 * @slot - Default slot for custom content
 * @csspart greeting - The greeting text container
 * @cssprop --greeting-color - Text color (default: currentColor)
 */
@customElement('hello-world')
export class HelloWorld extends LitElement {
  /**
   * The name to greet
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

### Specifying Tag Names

When the tag name can't be detected automatically, use `@customElement`, `@element`, or `@tagName`:

```typescript
/**
 * A vanilla custom element
 *
 * @customElement vanilla-element
 */
class MyElement extends HTMLElement {
  static is = 'vanilla-element';

  static {
    customElements.define(this.is, this);
  }
}
```

All three tags are aliases and work identically:
- `@customElement vanilla-element` (recommended)
- `@element vanilla-element`
- `@tagName vanilla-element`

The generator automatically detects tag names from the `@customElement` 
decorator and `customElements.define('tag-name', Class)` calls with static 
strings. However, when using dynamic patterns like 
`customElements.define(this.is, this)` where the tag name is stored in a 
variable, you must use one of the JSDoc tags above to specify the tag name 
explicitly.

## Documenting Slots and Parts

`cem` automatically detects `<slot>` elements and `part` attributes in your 
template. You can add descriptions using JSDoc tags or inline HTML comments.

### HTML Comment Documentation

Use plain comments for descriptions:

```html
<!-- This is the **default** `slot`. -->
<slot></slot>
```

Or use YAML for detailed metadata:

```html
<!--
  summary: The main slot for content
  description: |
    This slot displays user-provided content.
    Supports multiline **markdown**.
  deprecated: true
-->
<slot></slot>
```

When an element has both a slot and part attribute:

```html
<!-- slot:
       summary: The `info` slot
     part:
       summary: The `info-part` part
-->
<slot name="info" part="info-part"></slot>
```

{{<tip "warning">}}
When including inline markdown <code>\`code\`</code> in lit-html templates, escape the backticks in the comment.
{{</tip>}}

## Documenting CSS Custom Properties

Document CSS variables using JSDoc-style comments in your CSS:

```css
:host {
  /**
   * A property defined on the host
   * @summary The host's custom property
   */
  --host-property: red;

  color:
    /**
     * Custom color for use in this element
     * @summary color
     * @deprecated Use the `color` property instead
     */
    var(--custom-color);

  border:
    1px solid
    /** Border color of the element */
    var(--border-color);
}
```

**Position comments correctly** when both LHS and RHS contain CSS custom properties:

```css
/** Comment for --a */
color: var(--a);

/** Comment for --b */
--b: blue;

/** Comment for --c */
--c:
  /** Comment for --d */
  var(--d);
```

### Design Token Integration

Use the `--design-tokens` flag to integrate [DTCG-format][dtcg] design tokens:

```bash
cem generate --design-tokens npm:@my-ds/tokens/tokens.json --design-tokens-prefix my-ds
```

`npm:` and `jsr:` specifiers resolve from `node_modules` first. If the package
isn't installed locally, cem fetches it from [esm.sh](https://esm.sh)
automatically.

The prefix should not include leading dashes — use `my-ds`, not `--my-ds`.

When both user comments and design token descriptions exist for the same property, both are included (user description first, then design token description).

## Documenting Demos

### JSDoc `@demo` Tag

Link to demos directly from your element class:

```typescript
/**
 * @demo https://example.com/my-element-plain/
 * @demo https://example.com/my-element-fancy/ - A fancier demo
 */
@customElement('my-element')
class MyElement extends LitElement { }
```

### Automatic Demo Discovery

Configure automatic discovery in `.config/cem.yaml`:

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  demoDiscovery:
    fileGlob: "src/**/demos/*.html"
    urlPattern: "/src/:tag/demos/:demo.html"
    urlTemplate: "https://example.com/elements/{{.tag | alias}}/{{.demo | slug}}/"
```

**Template functions:**
- `alias` - Apply element alias mapping
- `slug` - Convert to URL-friendly format
- `lower` - Convert to lowercase
- `upper` - Convert to uppercase

See [Working with Demos][demos] for organization strategies.

## Code Examples

Use the `@example` tag for code examples:

```typescript
/**
 * @example Basic usage
 *          ```html
 *          <my-element></my-element>
 *          ```
 */
@customElement('my-element')
class MyElement extends LitElement { }
```

With explicit caption:

```typescript
/**
 * @example
 * <caption>Advanced usage with properties</caption>
 * ```html
 * <my-element color="primary" size="large"></my-element>
 * ```
 */
```

Multiple examples:

```typescript
/**
 * A flexible element
 * @example Simple case
 *          ```html
 *          <my-element></my-element>
 *          ```
 * @example With attributes
 *          ```html
 *          <my-element foo="bar"></my-element>
 *          ```
 */
```

Examples with captions are wrapped in `<figure>/<figcaption>` elements in the generated manifest.

## Documenting Methods

Document public methods using `@param` and `@returns` tags:

```typescript
/**
 * Updates the element's theme
 *
 * @summary Apply a new theme
 * @param {string} themeName - Name of the theme to apply
 * @param {boolean} [persist=false] - Whether to save theme preference
 * @returns {boolean} True if theme was applied successfully
 * @example
 * ```typescript
 * element.setTheme('dark', true);
 * ```
 */
setTheme(themeName: string, persist = false): boolean {
  // implementation
}
```

**Parameter syntax:**
- Required: `@param {type} name - description`
- Optional: `@param {type} [name] - description`
- With default: `@param {type} [name=default] - description`

## Monorepo Setup

For npm or yarn workspaces, create a `.config/cem.yaml` file in each package:

**Root package.json:**
```json
{
  "scripts": {
    "generate": "npm run generate --workspaces"
  },
  "workspaces": ["./core", "./elements"]
}
```

**core/.config/cem.yaml:**
```yaml
generate:
  files:
    - './**/*.ts'
```

**core/package.json:**
```json
{
  "scripts": {
    "generate": "cem generate"
  }
}
```

Repeat for each workspace package.

## See Also

- **[Development Workflow][workflow]** - When to regenerate the manifest
- **[Generate Command][generateref]** - Complete JSDoc tag reference and command options
- **[Working with Demos][demos]** - Demo organization strategies
- **[Effective Writing for AI][mcpwriting]** - AI-friendly documentation patterns

[workflow]: ../workflow/
[lsp]: /docs/installation/lsp/
[mcp]: /docs/installation/mcp/
[knobs]: ../knobs/
[generateref]: /docs/reference/commands/generate/
[demos]: ../demos/
[mcpwriting]: ../effective-mcp-descriptions/
[dtcg]: https://tr.designtokens.org/format/
