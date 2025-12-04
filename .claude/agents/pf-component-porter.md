# PatternFly Component Porter Agent

## Purpose
Port PatternFly React components to native web components following established patterns in the codebase.

## Process

### 1. Research Phase
- Read the PatternFly React component from `../patternfly/patternfly-react/packages/react-core/src/components/`
- Identify props, variants, and component structure
- Check for subcomponents and their relationships
- Note any PatternFly SCSS files in `../patternfly/patternfly/src/patternfly/components/`

### 2. Component Design
Translate React patterns to web component patterns:

**React Props → Web Component Attributes/Properties:**
- Boolean props → Boolean attributes (e.g., `isInline` → `inline` attribute)
- Enum props → String attributes (e.g., `variant="success"`)
- Complex props → Properties with getters/setters
- Children → Default slot
- Named children → Named slots

**React Events → Custom Events:**
- Create custom event classes extending `Event`
- Use class fields for event data (not `detail`)
- Make events `bubbles: true`

**React State → Web Component State:**
- Use private fields for internal state
- Use `observedAttributes` for reactive attributes
- Implement `attributeChangedCallback` for reactions

### 3. File Structure
Create component directory: `serve/middleware/routes/templates/elements/pf-v6-{component-name}/`

**Required files:**
1. `pf-v6-{component-name}.html` - Shadow DOM template (Go template)
2. `pf-v6-{component-name}.js` - Component logic
3. `pf-v6-{component-name}.css` - Styles

### 4. HTML Template (`*.html`)
- **No wrapper div** - The host element IS the component
- Use Go template conditionals for variant-based rendering
- Use simple selectors (IDs preferred) for shadow DOM elements
- Grid areas and semantic structure match PatternFly HTML
- Example:
  ```html
  {{if hasAttr . "Expandable"}}
  <div id="toggle">...</div>
  {{end}}

  <div id="icon">
    <slot name="icon"><!-- default icon --></slot>
  </div>

  <h4 id="title">
    <slot name="title"></slot>
  </h4>
  ```

### 5. CSS Styles (`*.css`)
**Host styles:**
```css
:host {
  display: grid; /* or appropriate display */

  /* Copy all CSS variables from PatternFly */
  --pf-v6-c-{component}--Property: value;
}

/* Variant modifiers use :host() selectors */
:host([variant="success"]) {
  --pf-v6-c-{component}--BorderColor: var(--pf-v6-c-{component}--m-success--BorderColor);
}

/* Boolean modifiers */
:host([inline]) {
  --pf-v6-c-{component}--BoxShadow: none;
}
```

**Internal elements:**
```css
#element-id {
  grid-area: element;
  /* PatternFly properties */
}

.pf-v6-c-{component}__element {
  /* Only use classes when needed for complex selectors */
}
```

**RTL support:**
```css
:host(:dir(rtl)) .element {
  /* RTL-specific styles */
}
```

### 6. JavaScript (`*.js`)
```javascript
import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom events
 */
export class PfComponentChangeEvent extends Event {
  constructor(value) {
    super('change', { bubbles: true });
    this.value = value;
  }
}

/**
 * Component documentation
 *
 * @fires change - Description
 * @slot - Default slot description
 * @slot name - Named slot description
 */
export class PfComponent extends CemElement {
  static is = 'pf-v6-{component}';

  static observedAttributes = ['variant', 'inline'];

  #$(id) {
    return this.shadowRoot?.getElementById(id);
  }

  // Getters/setters for attributes
  get variant() {
    return this.getAttribute('variant') || 'default';
  }

  set variant(value) {
    this.setAttribute('variant', value);
  }

  async afterTemplateLoaded() {
    // Set up event listeners
    // Initialize default values
    // Connect slots
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.shadowRoot?.firstChild) return;

    // React to attribute changes
  }

  // Public methods

  static {
    customElements.define(this.is, this);
  }
}
```

### 7. Integration
1. Add import to element's which use it, likely `cem-serve-chrome.js`:
   ```javascript
   import '/__cem/elements/pf-v6-{component}/pf-v6-{component}.js';
   ```

2. If needed, copy lightdom CSS:
   ```bash
   mv /path/to/patternfly/dist/css/components/{Component}/{component}-lightdom.css \
      serve/middleware/routes/templates/css/pf-v6-c-{component}-lightdom.css
   ```

### 8. Key Principles
- **Host as container**: The custom element itself gets the `pf-v6-c-{component}` styling
- **Modifiers as attributes**: Use `:host([attr])` selectors, not class modifiers on host
- **Shadow DOM friendly**: Prefer IDs over classes for unique internal elements
- **PatternFly CSS fidelity**: Copy all CSS variables to maintain theming
- **Web component idioms**: Slots for content projection, custom events for communication
- **SSR compatible**: Component works when shadow DOM is pre-rendered

### 9. Testing Checklist
- [ ] Component renders with default attributes
- [ ] All variants render correctly
- [ ] Boolean modifiers toggle correctly
- [ ] Slots accept content properly
- [ ] Events fire with correct data
- [ ] Public methods work as expected
- [ ] SSR rendering works
- [ ] RTL support (if applicable)
- [ ] Keyboard navigation (if applicable)
- [ ] ARIA attributes correct

### 10. Common Patterns

**Expandable/Collapsible:**
```javascript
toggle() {
  this.expanded = !this.expanded;
  this.dispatchEvent(this.expanded ? new ExpandEvent() : new CollapseEvent());
}
```

**Item Management (Groups):**
```javascript
addItem(item) {
  if (this.toast) {
    const wrapper = document.createElement('div');
    wrapper.className = 'pf-v6-c-{group}__item';
    wrapper.appendChild(item);
    this.appendChild(wrapper);
  } else {
    this.appendChild(item);
  }
}
```

**Close Actions:**
```javascript
// Listen for close events from slotted content
actionSlot.addEventListener('slotchange', () => {
  const actions = actionSlot.assignedElements();
  actions.forEach(action => {
    action.addEventListener('click', () => {
      if (action.hasAttribute('close')) {
        this.dispatchEvent(new CloseEvent());
      }
    });
  });
});
```

## Example: Porting pf-v6-alert

See the implementation in `serve/middleware/routes/templates/elements/pf-v6-alert/` for a complete example following all these patterns.

## Anti-Patterns to Avoid

❌ Don't wrap everything in a `<div class="pf-v6-c-component">` - use `:host`
❌ Don't use `class` on host for public API - use attributes
❌ Don't use `CustomEvent` with `detail` - create event classes with fields
❌ Don't import CSS in CSS files - use direct CSS variable definitions
❌ Don't use `:where(.pf-v6-m-dir-rtl)` - use `:host(:dir(rtl))`
❌ Don't create new files unless necessary - prefer editing existing
❌ Don't use aria-attrs as public APIs. Implement default semantics using `ElementInternals`.

## Accessibility
- Use element internals and ARIA IDL attributes for cross-root ARIA support 
where available e.g. `#internals.ariaLabelledByElements = [...]`;
## Resources

- PatternFly React: `~/Developer/patternfly/patternfly-react/`
- PatternFly SCSS: `~/Developer/patternfly/patternfly/`
- Existing components: `serve/middleware/routes/templates/elements/pf-v6-*/`
- Base class: `serve/middleware/routes/templates/js/cem-element.js`
