---
name: pf-component-porter
description: Ports patternfly v6 react components to cem's LitElement web components for the dev server chrome
---
# PatternFly Component Porter Agent

## Purpose
Port PatternFly React components to LitElement web components for the cem dev server chrome.
These are *minimal* ports - only implement what the dev server actually uses, not full-fidelity 1:1 React parity.

## References
- Steven Glanzer's full-fidelity ports: `~/Developer/pfv6-to-lit/elements/pfv6-*/`
- PatternFly React source: `~/Developer/patternfly/patternfly-react/packages/react-core/src/components/`
- PatternFly SCSS: `~/Developer/patternfly/patternfly/src/patternfly/components/`
- Existing ported components: `serve/elements/pf-v6-*/` (source) and `serve/middleware/routes/templates/elements/pf-v6-*/` (compiled output)
- API design advice: `extensions/claude-code/skills/design-api/ADVICE.md`

## Process

### 1. Research Phase
- Read the existing CemElement component in `serve/middleware/routes/templates/elements/pf-v6-{name}/`
- Check Steven's port at `~/Developer/pfv6-to-lit/elements/pfv6-{name}/` if it exists
- Read the PatternFly React component for additional context
- Identify which features the cem dev server actually uses (search for `<pf-v6-{name}` in templates)

### 2. Write the LitElement Component
Create `serve/elements/pf-v6-{name}/pf-v6-{name}.ts`

**Pattern:**
```typescript
import { LitElement, html, css } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

@customElement('pf-v6-{name}')
export class PfV6{Name} extends LitElement {
  static styles = css`
    :host {
      /* Copy CSS variables from the existing .css file */
      /* Use :host([attr]) for variant/state selectors */
    }
    /* Use ID selectors for unique shadow DOM elements */
  `;

  @property({ type: Boolean, reflect: true })
  someAttr = false;

  render() {
    return html`
      <slot></slot>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-{name}': PfV6{Name};
  }
}
```

### 3. Build & Verify
```bash
cd serve && go run generate_elements.go
# or
make build
```
This compiles .ts -> .js via esbuild, rewrites lit imports to /__cem/vendor/ paths.

### 4. Clean Up Old Files
After porting, remove the old CemElement files from `templates/elements/pf-v6-{name}/`:
- `pf-v6-{name}.html` (Go template - replaced by render())
- `pf-v6-{name}.css` (separate CSS - replaced by static styles)
Keep only the generated `.js` and `.js.map` output.

## Key Principles

### Minimal Surface Area
Only implement attributes, properties, events, and slots that the dev server chrome actually uses.
Search the codebase to find usage before adding features.

### Improve on PatternFly Where Possible
While maintaining visual compatibility, improve PF patterns where modern web platform,
accessibility, or performance allow:
- Use `color-scheme: light dark` on `:host` and `light-dark()` in CSS fallbacks
- Use native CSS nesting, `:has()`, `@layer`
- Use ElementInternals for ARIA instead of aria-* attributes on internal elements
- Use native `<dialog>`, `popover`, `<details>` where they replace JS-heavy PF patterns
- Prefer `<slot>` composition over imperative DOM manipulation
- Track improvements in `~/Developer/patternfly/MIGRATE.md` for the eventual full PFE port

### LitElement Idioms
- CSS in separate `.css` file, imported with `import styles from './{name}.css' with { type: 'css' };`
- `static styles = styles;`
- `render()` method returns `html\`...\`` template
- Standard (TC39) decorators with `accessor` keyword: `@property({ reflect: true }) accessor foo = false;`
- `@customElement('pf-v6-{name}')` decorator for registration
- `@state() accessor _internalState = ...;` for reactive internal state

### CSS Guidelines
- Always declare `color-scheme: light dark` on `:host`
- Copy CSS variables from the existing `.css` file into `static styles`
- Token references (e.g. `var(--pf-t--global--color--brand--default)`) already
  contain `light-dark()` in the token definitions - don't duplicate it
- Use `:host` for component-level styles
- Use `:host([attr])` for variant/state selectors (not classes on host)
- Use ID selectors for unique shadow DOM elements
- Native CSS nesting is fine
- Don't nest under `:host` except for host attribute selectors

### HTML Guidelines
- No wrapper div - the host element IS the component
- Prefer `<slot>` for content projection
- Use `id` attributes (not classes) for unique shadow elements
- Format attributes vertically aligned

### Events
- Create custom event classes extending `Event` with class fields (not `CustomEvent` with `detail`)
- Make events `bubbles: true`

### Accessibility
- Use ElementInternals for host-level ARIA (`this.attachInternals()`)
- Use ARIA IDL attributes where available

## Anti-Patterns to Avoid
- Don't wrap everything in a `<div class="pf-v6-c-component">` - use `:host`
- Don't use `class` on host for public API - use attributes
- Don't use `CustomEvent` with `detail` - create event classes with fields
- Don't add features the dev server doesn't use
- Don't import from `'lit'` bare specifier in output - the build handles rewriting

## Testing Checklist
- [ ] Component renders with default attributes
- [ ] All used variants render correctly
- [ ] Slots accept content properly
- [ ] Events fire with correct data
- [ ] Existing dev server chrome still works after swapping the component
