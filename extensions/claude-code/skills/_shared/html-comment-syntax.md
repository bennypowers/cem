# Inline HTML Comment Syntax for Slots and Parts

`cem` extracts documentation from HTML comments immediately before `<slot>`
elements and elements with `part` attributes. Prefer inline comments over
JSDoc `@slot`/`@csspart` tags -- they stay co-located with the markup.

## Plain string comment

A plain (non-YAML) comment becomes the `description`:

```html
<!-- Button label content. SHOULD contain text. -->
<slot></slot>

<!-- The native button element -->
<button part="button">
```

## Single-key YAML

Use `summary:` for a short label shown in tooling previews:

```html
<!-- summary: Button label content -->
<slot></slot>
```

## Multi-key YAML

Combine `summary`, `description`, and `deprecated`:

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

## Combined slot and part

When an element has both a `slot` name and a `part` attribute, use nested
`slot:` and `part:` keys. Each accepts either a scalar string (shorthand
for `description`) or the full object form. Mix them freely:

```html
<!-- slot: Sub navigation links
     part:
       summary: Nav container
       description: The scrollable link list container -->
<slot name="nav" part="nav-container"></slot>

<!-- part: The overlay container
     slot:
       summary: Overlay
       description: Content shown in the overlay -->
<slot name="overlay" part="overlay"></slot>

<!-- slot: Info content
     part: Info container -->
<slot name="info" part="info"></slot>
```

## Backtick escaping

Inside lit-html `html` tagged template literals, escape backticks in
comments with `\``:

```html
<!-- summary: Uses \`currentColor\` by default -->
<slot></slot>
```

## CSS custom properties

Document CSS variables using JSDoc-style comments in CSS source files
(preferred over `@cssprop` tags on the class):

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
