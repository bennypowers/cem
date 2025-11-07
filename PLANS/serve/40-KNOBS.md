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
  
