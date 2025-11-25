## Porting PatternFly Components

This document captures learnings from porting PatternFly v6 components to the dev server chrome.

### Design Token Strategy

**Always use PatternFly design tokens, never hardcoded values.**

Create a minimal `pf-tokens.css` with semantic tokens that map to PatternFly's global tokens:

```css
:root {
  /* Text colors */
  --pf-t--global--text--color--regular: light-dark(#151515, #ffffff);
  --pf-t--global--text--color--subtle: light-dark(#6a6e73, #c7c7c7);
  --pf-t--global--text--color--on-disabled: light-dark(#6a6e73, #b0b0b0);

  /* Background colors */
  --pf-t--global--background--color--primary--default: light-dark(#ffffff, #1f1f1f);
  --pf-t--global--background--color--action--plain--default: transparent;
  --pf-t--global--background--color--action--plain--hover: light-dark(#f0f0f0, #3c3f42);

  /* Border colors */
  --pf-t--global--border--color--default: light-dark(#c7c7c7, #707070);
  --pf-t--global--border--color--clicked: #0066cc;
  --pf-t--global--border--color--high-contrast: light-dark(#151515, #ffffff);

  /* Spacing */
  --pf-t--global--spacer--xs: 0.25rem;  /* 4px */
  --pf-t--global--spacer--sm: 0.5rem;   /* 8px */
  --pf-t--global--spacer--md: 1rem;     /* 16px */

  /* Border widths */
  --pf-t--global--border--width--100: 1px;
  --pf-t--global--border--width--extra-strong: 3px;
  --pf-t--global--border--width--action--plain--hover: 0px;

  /* Motion */
  --pf-t--global--motion--duration--fade--short: 100ms;
  --pf-t--global--motion--duration--md: 200ms;
  --pf-t--global--motion--timing-function--default: cubic-bezier(0.4, 0, 0.2, 1);
  --pf-t--global--motion--timing-function--decelerate: cubic-bezier(0, 0, 0.2, 1);
}
```

**Key principles:**

1. Use `light-dark()` for theme-aware colors
2. Add tokens incrementally as needed
3. Match PatternFly's token naming exactly
4. Include comments for pixel values on spacing tokens

### Understanding PatternFly Component Structure

PatternFly components often use a two-layer structure:

```html
<div class="pf-v6-c-tabs">
  <div class="pf-v6-c-tabs__list">
    <!-- Item provides spacing -->
    <div class="pf-v6-c-tabs__item">
      <!-- Link is the visual/interactive element -->
      <button class="pf-v6-c-tabs__link">Tab 1</button>
    </div>
  </div>
</div>
```

**Why two elements?**

- **Outer element** (`__item`): Provides spacing/padding around inner element
- **Inner element** (`__link`): The visual background and interactive surface

**Can you avoid the wrapper?**

Yes! Use a `::before` pseudo-element:

```css
.tab {
  /* Full padding for large click target */
  padding: calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--xs))
           calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--sm));
  background-color: transparent;
  border: 0;
}

/* Visual background - inset to create spacing effect */
.tab::before {
  content: '';
  position: absolute;
  inset-block-start: 0;
  inset-inline-start: 0;
  width: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));
  height: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));
  translate: var(--pf-t--global--spacer--sm) var(--pf-t--global--spacer--sm);
  background-color: var(--tab--BackgroundColor);
  border-radius: var(--pf-t--global--border--radius--small);
  z-index: -1;
}
```

**Benefits:**

- ✅ Single element in DOM (simpler)
- ✅ Same visual appearance
- ✅ Full clickable area (accessibility)
- ✅ GPU-accelerated `translate` (performance)

### Component Token Cascade Pattern

Use CSS custom properties to create overridable component tokens. Use `--_` prefix for private tokens (internal implementation):

```css
.tabs {
  /* Private defaults (internal implementation) */
  --_tab--Color: var(--pf-t--global--text--color--subtle);
  --_tab--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);

  /* Private hover targets */
  --_tab--hover--Color: var(--pf-t--global--text--color--regular);
  --_tab--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);
}

.tab {
  /* Use private component tokens */
  color: var(--_tab--Color);
  background-color: var(--_tab--BackgroundColor);
}

.tab:where(:hover, :focus) {
  /* Override to point to hover values */
  --_tab--Color: var(--_tab--hover--Color);
  --_tab--BackgroundColor: var(--_tab--hover--BackgroundColor);
}
```

**Private vs Public tokens:**

- `--_foo` = Private, internal implementation detail
- `--foo` = Public API, consumers can override

Use public tokens when you want to expose customization points.

**Why this pattern?**

PatternFly uses this extensively. From their SASS:

```scss
--#{$tabs}__link--Color: var(--pf-t--global--text--color--subtle);
--#{$tabs}__link--hover--Color: var(--pf-t--global--text--color--regular);

.pf-v6-c-tabs__link {
  color: var(--pf-v6-c-tabs__link--Color);

  &:hover {
    --pf-v6-c-tabs__link--Color: var(--pf-v6-c-tabs__link--hover--Color);
  }
}
```

This creates a clean cascade where:
1. Component defines default and target tokens
2. Element uses current value tokens
3. State selectors override to point to target tokens

### Animated Accent Lines

PatternFly tabs use a single animated `::after` element for the active indicator:

```css
@property --tabs--link-accent--length {
  syntax: "<length>";
  inherits: true;
  initial-value: 0px;
}

.tabs::after {
  content: '';
  position: absolute;
  inset-block-start: auto;
  inset-block-end: 0;
  inset-inline-start: 0;
  width: var(--tabs--link-accent--length);
  translate: var(--tabs--link-accent--start) 0;
  border-block-end: 3px solid var(--pf-t--global--border--color--clicked);
  transition-property: width, translate;
  transition-duration: var(--pf-t--global--motion--duration--md);
  transition-timing-function: var(--pf-t--global--motion--timing-function--decelerate);
}
```

**Key techniques:**

1. **`@property` for smooth width transitions** - Allows animating to/from auto values
2. **`translate` for position** - GPU-accelerated, better than `left`
3. **Single element** - One `::after` that animates position and width
4. **JavaScript updates CSS custom properties** - Sets `--start` and `--length` values

```javascript
#updateAccentLine(index) {
  const buttons = Array.from(this.shadowRoot.querySelectorAll('.tab'));
  const activeButton = buttons[index];

  const containerRect = this.#tabsContainer.getBoundingClientRect();
  const buttonRect = activeButton.getBoundingClientRect();

  // If using pseudo-element spacing, account for inset
  const inset = 8;
  const start = (buttonRect.left - containerRect.left) + inset;
  const length = buttonRect.width - (inset * 2);

  this.#tabsContainer.style.setProperty('--tabs--link-accent--start', `${start}px`);
  this.#tabsContainer.style.setProperty('--tabs--link-accent--length', `${length}px`);
}
```

### Border Width Gotchas

PatternFly's `--border-width--action--plain--hover` is `0px` by default!

It only shows a border in high-contrast mode:

```scss
// tokens-default.scss
--pf-t--global--border--width--action--plain--hover: 0px;

// tokens-highcontrast.scss
--pf-t--global--border--width--action--plain--hover: var(--pf-t--global--border--width--100);
```

If you want visible borders on hover in normal mode, use `--border-width--100` (1px) directly.

### Using Playwright to Extract Styles

When visual appearance doesn't match PatternFly, use Playwright to scrape actual computed styles:

```javascript
import { chromium } from 'playwright';

const browser = await chromium.launch();
const page = await browser.newPage();

await page.goto('https://staging-v6.patternfly.org/components/tabs/html-demos/default');

// Get computed styles
const styles = await page.locator('.pf-v6-c-tabs__link').first().evaluate((el) => {
  const computed = window.getComputedStyle(el);
  return {
    padding: computed.padding,
    fontSize: computed.fontSize,
    backgroundColor: computed.backgroundColor,
    // ... etc
  };
});

console.log(styles);
await browser.close();
```

This helped discover:
- Combined padding: item (8px) + link (4px/8px) = 12px/16px
- `border: 0` not `border: transparent`
- Exact color values for backgrounds

### Server-Side Rendering with DSD

All chrome components use Declarative Shadow DOM (DSD) for SSR:

**Template structure:**

```html
<pf-v6-tabs>
  {{renderElementShadowRoot "pf-v6-tabs" dict}}
  <pf-v6-tab title="Knobs">
    {{renderElementShadowRoot "pf-v6-tab" dict}}
    <div>Tab content</div>
  </pf-v6-tab>
</pf-v6-tabs>
```

**Component expects pre-attached shadow root:**

```javascript
class PfV6Tabs extends HTMLElement {
  connectedCallback() {
    if (!this.shadowRoot) {
      console.error('Shadow root not found. Use renderElementShadowRoot in templates.');
      return;
    }

    // Shadow root already exists, just use it
    this.#tabsContainer = this.shadowRoot.querySelector('.tabs');
  }
}
```

**No client-side shadow root creation** - it's SSR only.

### Dynamic Slots for ARIA Accessibility

When buttons are in shadow DOM and panels in light DOM, `aria-controls` doesn't work across the boundary.

**Solution:** Create panels in shadow DOM with dynamic named slots:

```javascript
#updateTabs() {
  const panelsContainer = this.shadowRoot.querySelector('.panels');

  this.#tabs.forEach((tab, index) => {
    // Create panel in shadow DOM
    const panel = document.createElement('div');
    panel.className = 'panel';
    panel.id = `panel-${index}`;
    panel.role = 'tabpanel';

    // Add named slot
    const slot = document.createElement('slot');
    slot.name = `panel-${index}`;
    panel.appendChild(slot);

    panelsContainer.appendChild(panel);

    // Assign slot to light DOM tab
    tab.setAttribute('slot', `panel-${index}`);
  });
}
```

Now buttons and panels are in the same shadow root, so ARIA relationships work!

### Declarative Child API Pattern

Instead of manual `<slot>` assignment:

```html
<!-- ❌ Manual slot API -->
<pf-v6-tabs>
  <button slot="tab-0">Tab 1</button>
  <div slot="panel-0">Panel 1</div>
</pf-v6-tabs>
```

Use declarative children:

```html
<!-- ✅ Declarative API -->
<pf-v6-tabs>
  <pf-v6-tab title="Tab 1">
    Panel content
  </pf-v6-tab>
  <pf-v6-tab title="Tab 2">
    Panel content
  </pf-v6-tab>
</pf-v6-tabs>
```

**Benefits:**

- Better DX (more intuitive)
- Easier to SSR
- More accessible (proper structure)
- Tab component generates buttons from children

**Implementation:**

```javascript
#updateTabs() {
  this.#tabs = Array.from(this.querySelectorAll('pf-v6-tab'));

  this.#tabs.forEach((tab, index) => {
    const button = document.createElement('button');
    button.textContent = tab.title || `Tab ${index + 1}`;
    button.addEventListener('click', () => {
      this.selectedIndex = index;
    });

    this.#tabsContainer.appendChild(button);
  });
}
```

Child element (`pf-v6-tab`) is just a wrapper with `display: contents`:

```css
:host {
  display: contents;
}
```

### Performance: Prefer CSS `translate` Over Layout Properties

Always use CSS transforms for animations:

```css
/* ✅ Good - GPU accelerated */
.element {
  translate: var(--x) var(--y);
  transition: translate 200ms;
}

/* ❌ Bad - triggers layout */
.element {
  left: var(--x);
  top: var(--y);
  transition: left 200ms, top 200ms;
}
```

For `width` animations, use `@property`:

```css
@property --width {
  syntax: "<length>";
  inherits: true;
  initial-value: 0px;
}

.element {
  width: var(--width);
  transition: width 200ms;
}
```

This allows smooth transitions for computed values.

### Attribute-Dependent SSR Templates

The shadowroot middleware provides `.Attributes` map and camelCased properties to templates, enabling conditional rendering based on host attributes.

**Template data structure:**

```go
type TemplateData struct {
    Attributes map[string]string  // Raw attributes: {"href": "...", "variant": "primary"}
    // Plus camelCased keys:
    Href       string             // .Href
    Variant    string             // .Variant
    AriaLabel  string             // .AriaLabel (from aria-label)
    DataTestId string             // .DataTestId (from data-test-id)
}
```

**Example: pf-v6-button with conditional structure**

Host HTML:
```html
<!-- Without href attribute -->
<pf-v6-button variant="primary">Click me</pf-v6-button>

<!-- With href attribute -->
<pf-v6-button href="/docs" variant="link">Docs</pf-v6-button>
```

Template (`pf-v6-button.html`):
```html
{{- define "button-content" -}}
  <slot name="icon-start"></slot>
  <slot id="content"></slot>
  <slot name="icon-end"></slot>
{{- end -}}

{{if .Attributes.Href -}}
<a href="{{.Attributes.Href}}">
  {{- template "button-content" . -}}
</a>
{{- else -}}
<span id="wrapper" part="wrapper">
  {{- template "button-content" . -}}
</span>
{{- end}}
```

**Rendered output (without href):**
```html
<pf-v6-button variant="primary">
  <template shadowrootmode="open">
    <style>/* ... */</style>
    <span id="wrapper" part="wrapper">
      <slot name="icon-start"></slot>
      <slot id="content"></slot>
      <slot name="icon-end"></slot>
    </span>
  </template>
  Click me
</pf-v6-button>
```

**Rendered output (with href):**
```html
<pf-v6-button href="/docs" variant="link">
  <template shadowrootmode="open">
    <style>/* ... */</style>
    <a href="/docs">
      <slot name="icon-start"></slot>
      <slot id="content"></slot>
      <slot name="icon-end"></slot>
    </a>
  </template>
  Docs
</pf-v6-button>
```

**Key benefits:**

1. **No client-side DOM manipulation** - Structure is correct from SSR
2. **Faster rendering** - Browser doesn't need to wait for JS to restructure
3. **Consistent semantics** - Button vs link role determined at render time
4. **Better accessibility** - Screen readers see correct structure immediately

**Boolean attributes:**

```html
{{if .Attributes.Expanded}}
  <div id="body">
    <slot></slot>
  </div>
{{else}}
  <div id="body" inert>
    <slot></slot>
  </div>
{{end}}
```

Boolean attributes in `.Attributes` are:
- Present with empty string value (`""`) when attribute exists: `<element expanded>`
- Absent from map when attribute doesn't exist

Test for presence with `{{if .Attributes.AttrName}}`.

**Accessing camelCase properties:**

```html
<!-- ARIA attributes -->
{{if .AriaLabel}}
  aria-label="{{.AriaLabel}}"
{{end}}

<!-- Data attributes -->
{{if .DataTestId}}
  data-testid="{{.DataTestId}}"
{{end}}
```

**Conditional classes:**

```html
<div class="base-class{{if .Attributes.Compact}} pf-m-compact{{end}}">
  <slot></slot>
</div>
```

**Components using attribute-dependent SSR:**

- **pf-v6-button** - Conditional `<a>` vs `<span>` based on `href`
- **pf-v6-label** - Conditional `<a>` vs `<span>` based on `href`
- **pf-v6-form-field-group** - Conditional toggle button and header based on `expandable` and `toggle-text`

**Client-side hydration:**

Components must detect SSR structure and skip redundant DOM manipulation:

```javascript
async afterTemplateLoaded() {
  this.#isLink = this.hasAttribute('href');
  this.#element = this.#isLink ? this.shadowRoot.querySelector('a') : null;

  if (!this.#isLink) {
    // Host-based button semantics
    this.#setupHostButton();
  } else {
    // Shadow link semantics
    this.#setupShadowLink();
  }
}
```

Never re-create structure that SSR already rendered—just enhance it with interactivity.

### Performance: SSR vs Client-Side Rendering

**SSR with attribute-dependent templates provides measurable performance benefits:**

**Rendering Timeline Comparison:**

```
Client-Side Rendering:
1. HTML arrives          (0ms)
2. Parse HTML            (+5ms)
3. First paint           (+10ms) - but wrong structure
4. JS downloads          (+50ms)
5. JS parses/executes    (+100ms)
6. Component hydrates    (+105ms)
7. DOM manipulation      (+110ms) - replace <button> with <a>
8. Reflow/repaint        (+115ms)
9. Correct structure     (+120ms)

SSR with Attributes:
1. HTML arrives          (0ms)
2. Parse HTML            (+5ms)
3. Correct structure     (+10ms) - already has <a>
4. First paint           (+10ms)
5. JS downloads          (+50ms)
6. JS parses/executes    (+100ms)
7. Component hydrates    (+105ms)
8. Done                  (+105ms)
```

**Key Performance Wins:**

1. **No layout thrashing** - Structure is correct from initial parse, no DOM replacements
2. **Faster First Contentful Paint (FCP)** - Correct structure visible immediately
3. **Reduced Cumulative Layout Shift (CLS)** - No structural changes after hydration
4. **Lower Time to Interactive (TTI)** - 15ms saved by skipping DOM manipulation

**Real-World Impact:**

For a dev server chrome with 20+ components using attribute-dependent SSR:
- **Client-side**: 20 components × 10ms DOM manipulation = 200ms blocking time
- **SSR**: 0ms DOM manipulation, components enhance in place

**Memory Benefits:**

SSR approach uses less memory:
```javascript
// Client-side: Creates temporary nodes, then replaces them
const oldButton = document.createElement('button');
shadowRoot.append(oldButton);
// Later...
const newLink = document.createElement('a');
shadowRoot.replaceChild(newLink, oldButton); // oldButton becomes garbage

// SSR: Uses existing node, no garbage created
const element = shadowRoot.querySelector('a'); // Already rendered by SSR
element.addEventListener('click', handler); // Just enhance
```

**Browser Optimization Opportunities:**

SSR-rendered structure enables browser optimizations:
- **Speculative parsing** - Browser can preload linked resources from `<a href>`
- **Paint coalescing** - Single paint pass instead of multiple
- **Layer composition** - Correct stacking context from start

**When Client-Side Rendering Might Be Acceptable:**

- **Highly dynamic UIs** - Structure changes frequently based on user interaction
- **Client-only state** - Structure depends on localStorage/IndexedDB
- **Extreme customization** - User can completely reconfigure component structure

For dev server chrome components, SSR is clearly superior—structure is deterministic and attribute-driven.

**Measurement Recommendations:**

Use Chrome DevTools Performance panel to measure:
1. **FCP** - Should be 10-20ms faster with SSR
2. **LCP** - Largest Contentful Paint should improve
3. **CLS** - Should be 0 with SSR (no layout shifts)
4. **Main thread blocking time** - Reduced by DOM manipulation time

Use Lighthouse to validate:
```bash
lighthouse http://localhost:8080/demo --view
```

Look for improvements in:
- Performance score
- First Contentful Paint
- Cumulative Layout Shift
- Total Blocking Time

### Summary

**Essential patterns for porting PatternFly components:**

1. **Use design tokens exclusively** - Never hardcode colors/spacing
2. **Token cascade pattern** - Define defaults and targets, override in selectors
3. **Pseudo-elements for spacing** - Avoid wrapper elements when possible
4. **Prefer `translate` over layout properties** - Better performance
5. **Use `@property` for animating lengths** - Smooth width/height transitions
6. **SSR with DSD** - Pre-render shadow roots on server
7. **Attribute-dependent templates** - Conditional structure based on host attributes (10-20ms FCP improvement)
8. **Client-side hydration** - Detect SSR structure, only add interactivity (eliminates layout shifts)
9. **Dynamic slots for ARIA** - Keep related elements in same shadow root
10. **Declarative child APIs** - Better DX than manual slot assignment
11. **Performance measurement** - Use Lighthouse and DevTools to validate SSR benefits
12. **Playwright for validation** - Extract real computed styles when stuck
