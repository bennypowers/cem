---
title: Knobs
layout: docs
weight: 30
---

Knobs are interactive controls that let you test your custom elements by manipulating their attributes, properties, and CSS custom properties in real-time.

## How Knobs Work

Knobs are automatically generated from your **Custom Elements Manifest**. The dev server reads your component's API documentation and creates appropriate controls for each attribute, property, and CSS variable.

### The Workflow

```goat
  Component            Manifest
.-----------.       .----------.
|  @attr    |       | attrs[]  |
|  @prop    | ----> | members  |
|  @cssprop |       | cssProps |
'-----------'       '----+-----'
                         |
 cem generate            | cem serve
                         v
                    .---------.
                    | Switch  |
                    | Input   |
                    | Select  |
                    '---------'
                      Knobs
```

Everything derives from your component documentation - the manifest is the single source of truth.

## Manifest to Knobs Mapping

Knobs are generated from the manifest entries for each element. The dev server looks at:

**Attributes** (`attributes` array in manifest):
- Each attribute becomes a knob based on its `type` field
- `type: { text: "boolean" }` → Switch control
- `type: { text: "string" }` → Text input
- `type: { text: "'a' | 'b' | 'c'" }` → Select dropdown

**Properties** (`members` array where `kind: "field"`):
- JavaScript properties not reflected to attributes
- Same type-to-control mapping as attributes
- Useful for complex data that doesn't serialize to strings

**CSS Custom Properties** (`cssProperties` array):
- CSS variables defined on the element
- Color values get a color picker
- Other values get text input

To document your components for knobs, see the **[Generate documentation](/docs/commands/generate/)** which covers JSDoc syntax for `@attr`, `@property`, `@cssprop`, and other tags.

## Knob Types

The dev server generates different controls based on the type information in your manifest:

| Type | Control | Notes |
| ---- | ------- | ----- |
| `boolean` | Switch | Toggle on/off for boolean values |
| `string` | Text input | Single-line text entry |
| `number` | Number input | Numeric entry with increment/decrement |
| `'a' \| 'b' \| 'c'` | Select dropdown | Dropdown with defined options |
| CSS color value | Color picker | Visual color selector with hex input |

### Boolean Knobs

For boolean attributes/properties:

```js
/**
 * @attr {boolean} disabled
 */
```

Generates a checkbox. When checked, the attribute is added; unchecked, it's removed.

### Enum Knobs

For union types (enums):

```js
/**
 * @attr {'primary' | 'secondary' | 'danger'} variant
 */
```

Generates a select dropdown with options: primary, secondary, danger.

### Number Knobs

For numeric types:

```js
/**
 * @property {number} count
 */
```

Generates a number input with increment/decrement controls.

### Color Knobs

CSS custom properties with color values:

```js
/**
 * @cssprop --bg-color - Background color (default: #ff0000)
 */
```

Generates a dual control: color picker + text input for precise hex values.

## Knob Categories

Knobs are organized into three collapsible sections:

1. **Attributes** - DOM attributes (e.g., `disabled`, `variant`)
2. **Properties** - JavaScript properties (e.g., `value`, `data`)
3. **CSS Properties** - CSS custom properties (e.g., `--bg-color`)

You can expand/collapse each section in the UI.

## Knobs for All Elements

The dev server generates knobs for every custom element in your demo that's documented in the manifest. Each element gets its own labeled group of controls.

**Element detection:**
1. The dev server scans your demo HTML for custom elements
2. For each element found, it looks up the manifest documentation
3. If documented, a knob group is created with controls for that element's API

**Instance labeling:**
When you have multiple instances of the same element, each gets a unique label using:
- `#id` (if element has an `id` attribute)
- Text content (first 20 characters)
- `aria-label`
- Fallback: "tag-name No. N"

**Example:**

```html
<my-card id="card-1">
  <h2 slot="title">First Card</h2>
</my-card>

<my-card id="card-2">
  <h2 slot="title">Second Card</h2>
</my-card>

<my-button variant="primary">Click Me</my-button>
```

The sidebar shows:
- **my-card** section:
  - **#card-1** (knobs for first card)
  - **#card-2** (knobs for second card)
- **my-button** section:
  - **my-button No. 1** (knobs for button)

Each element's knobs are independent - changing a knob updates only that specific instance.

## State Synchronization

Knobs stay synchronized with your elements:

- **Initial values** come from demo HTML attributes
- **Changing a knob** updates the element immediately
- **Manual changes** (via DevTools) don't update knobs (one-way binding)

## How to Use Knobs

1. **Run your demo** in `cem serve`
2. **Open the sidebar** (desktop) or menu (mobile)
3. **Find your element** in the knobs section
4. **Adjust controls** to test different states
5. **See live updates** in the demo

## Troubleshooting

### No knobs appearing

Check:
1. **Manifest has API docs** - Run `cem generate` and verify `custom-elements.json` has attributes/properties/cssProperties
2. **Component is in demo** - The element must be present in your demo HTML
3. **Documentation syntax** - Use proper JSDoc tags (`@attr`, `@property`, `@cssprop`)

### Wrong control type

If a string shows a checkbox or vice versa:
1. **Check type annotation** - Ensure `@attr {boolean}` or `@attr {string}` is correct
2. **Regenerate manifest** - Run `cem generate` after fixing JSDoc
3. **Restart serve** - Changes to manifest require server restart

### Knob doesn't update element

If changing a knob has no effect:
1. **Check attribute reflection** - Ensure your component responds to attribute changes
2. **Use attributeChangedCallback** - For reactive attributes
3. **Check console** - Look for JavaScript errors

### Multiple instances not working

If multi-instance detection fails:
1. **Add IDs** - Give elements unique `id` attributes for better labels
2. **Check HTML structure** - Elements must be at same level or nested consistently
3. **Verify tag names** - All instances must have the same tag name

## Example: Complete Component

Here's a well-documented component that gets full knobs support:

```js
/**
 * A customizable button component
 *
 * @attr {string} variant - Visual style: 'primary', 'secondary', 'danger'
 * @attr {boolean} disabled - Disables the button
 * @attr {string} label - Accessible label for the button
 *
 * @property {number} clickCount - Number of times button was clicked
 *
 * @cssprop --button-bg - Background color (default: #0066cc)
 * @cssprop --button-padding - Internal padding (default: 8px 16px)
 *
 * @customElement my-button
 */
class MyButton extends HTMLElement {
  // Implementation...
}
```

This generates:
- **Attributes**: variant (select), disabled (checkbox), label (text)
- **Properties**: clickCount (number)
- **CSS Properties**: --button-bg (color), --button-padding (text)

## What's Next?

- **[Generate Docs](/docs/commands/generate/)** - Learn JSDoc syntax for documentation
- **[Buildless Development](/docs/serve/buildless/)** - Write TypeScript and import CSS without build steps
- **[Getting Started](/docs/serve/getting-started/)** - Set up your first demo
