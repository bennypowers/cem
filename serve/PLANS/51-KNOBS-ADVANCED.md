# Advanced Knobs: Multiple Elements & Complex Compositions

This document covers advanced knob features for handling complex demos with multiple custom elements, compositional components, and dynamic DOM manipulation.

**Prerequisites**: Basic knobs (50-KNOBS-CORE.md) must be implemented first.

## Multiple Elements on the Page

Demos may contain multiple instances of custom elements. Each needs its own knob group.

**Strategy**: Use a Mutation Observer to watch for DOM changes and add/remove knobs dynamically.

## The Complex Case: Compositional Components

Many component libraries use compositional patterns where multiple related elements work together.

### Example Scenario

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

This demo contains:
- 1 `<pf-tabs>` (container)
- 3 `<pf-tab>` instances
- 3 `<pf-tab-panel>` instances

Each element should have independently controllable knobs.

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

## Server-Side Implementation

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

### Shadow DOM Considerations

- **Elements in shadow DOM?** Traverse shadow roots to find elements
- **Dynamically created elements?** Mutation observer detects and adds knobs
- **Closed shadow roots?** Unsupported (can't access)

### Element Labeling

To distinguish multiple instances of the same element, use a priority-based labeling system:

**Priority for labeling child instances:**

1. **`id` attribute** → `"Tab: #home-tab"`
2. **Text content (trimmed)** → `"Tab: Users"`
3. **`aria-label`** → `"Tab: User Settings"`
4. **Slot name + index** → `"Tab (slot: tab) No. 2"`
5. **Fallback** → `"pf-tab No. 3"`

### Edge Cases

- **Empty text content**: Element has `textContent` but it's all whitespace → no label
- **Long text content**: 1000 character text node → max 20 chars (via CSS `text-overflow: ellipsis`)
- **Multiple aria-labels**: Nested elements with different labels → take most relevant to the element in question
- **Special characters**: Label contains `<`, `>`, `"` (XSS risk) → escaped by template engine

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

**Template behavior**:
- Primary element: `open` attribute (expanded by default)
- Child elements: collapsed by default
- User can expand/collapse any group

### UI Templates

Rather than async fetch calls and `innerHTML` every time we need to add knobs, the server should pre-compute `<template>` elements for each knob type.

**Strategy**:
- Server computes template elements during initial render
- Templates contain Declarative Shadow DOM
- CEM content applied via attributes (preferred) or slotted content (descriptions)
- Client-side code clones and appends templates synchronously

## Client-Side Implementation

### Mutation Observer Setup

```ts
function isCustomElement(node: Node): node is HTMLElement {
  // is it a CE known to the registry?
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
    // determine according to new source order which knob-set-for-element precedes the new one
    // source order:
    // - Walk demo DOM in tree order, collect all known custom elements
    // - Build ordered list of knob elements matching demo order
    // - Insert new knobs at correct position in list
    const beforeKnob = /* ... */
    beforeKnob.after(knob);
    this.#elementKnobsMap.set(element, knob);
  }

  /**
   * Constructs a client-side instance of the knob for this element
   * Based on tag name and attributes, using <template> elements generated
   * server side for each knob type.
   * knobs don't reset their targets on connect, rather they adopt the value of
   * their targets. This prevents loops when adding new knobs
   */
  #createKnobGroupForElement(element: HTMLElement): HTMLElement {
    // ...
  }
}
```

### Preventing Infinite Loops

**Problem**: Knob changes attribute → mutation observer fires → may recreate knobs → infinite loop.

**Solution**: Knobs adopt target element values on creation instead of resetting them.

```typescript
// WRONG: May cause loops
knob.connectedCallback() {
  this.targetElement.setAttribute('color', this.defaultValue);
}

// CORRECT: Read existing value
knob.connectedCallback() {
  this.value = this.targetElement.getAttribute('color') || this.defaultValue;
}
```

## Knob State Synchronization

Each element gets a mutation observer for attributes to keep knobs in sync.

### Synchronization Scenarios

**Scenario 1: Attribute changes via script (not knob)**
- Solution: Mutation Observer for attributes updates knob UI

**Scenario 2: Property changes internally**
- Solution: Mutation Observer catches reflected changes
- Optional: Read DOM properties at each mutation for deeper sync

**Scenario 3: Multiple knobs control same element**
- Non-issue: Each knob controls a single API (attribute, property, CSS var)
- No conflicts since APIs are independent

## Server-Side Template Data Flow

### Data Model

**Primary source**: Manifest (includes attribute types, defaults, etc.)

**Type parsing**: Knob code parses TypeScript type strings for enum select widgets.

See RHDS docs `<uxdot-knob>` for prior art.

### Complex Type Handling

**Simple union types**: `'a' | 'b'` → `<select>` with options

**Compound union types**: `'a' | 'b' | OtherUnion`
- Can implement rudimentary type computation
- Don't recreate full `tsc` (out of scope)

**Complex types**: `'a' | 'b' | {some: 'object'}`
- Fall back to `<input type="text">`
- Document this limitation

**Future enhancement** (after initial release):
- Consider bringing in TypeScript's Go library to compute unions
- Not a goal for this initial project

## Performance with Many Elements

### The Problem

Complex demos may have many elements:
- Example: Data table with 100 rows
- 100 elements × 10 knobs = 1000 controls
- Concerns: DOM size, memory, interaction lag

### Solutions

#### Required for Initial Release

**Collapse knobs by default** (only primary expanded)
- Reduces initial DOM size
- User expands on demand
- Implemented via `<details>` element

#### Stretch Goals (Not for Initial Release)

**Search/filter for knobs**
- Filter knob groups by tag name or label
- Helpful when many elements present

**Virtualize knob list if >50 elements**
- Only render visible knob groups
- Scroll → render more groups
- Reduces DOM size significantly

## Testing Strategy

**Follow established patterns from Phase 5a and Phases 1-4:**

### Fixture Structure
```
serve/middleware/knobs/testdata/
  ├── advanced/
  │   ├── complex-composition/
  │   │   ├── manifest.json (pf-tabs + pf-tab + pf-tab-panel)
  │   │   ├── demo.html (nested elements)
  │   │   └── expected-knobs.html (golden: all knob groups rendered)
  │   ├── mutation-observer/
  │   │   ├── demo-initial.html
  │   │   ├── expected-initial.html (golden)
  │   │   ├── expected-after-add.html (golden)
  │   │   └── expected-after-remove.html (golden)
  │   └── label-generation/
  │       ├── manifest.json
  │       ├── demo-with-ids.html
  │       ├── demo-with-text.html
  │       ├── demo-with-aria.html
  │       └── expected-*.html (golden files for each labeling scenario)
```

### Test Pattern
```go
func TestAdvancedKnobs_ComplexComposition(t *testing.T) {
    mfs := testutil.NewFixtureFS()
    // Load manifest with multiple elements
    // Load demo with nested composition
    // Generate knobs
    // Verify against golden file
    // Verify ordering matches source order
    // Verify primary element expanded, others collapsed
}
```

### Client-Side Testing
If feasible with Go's JavaScript testing support:
- Test MutationObserver behavior
- Test element addition/removal
- Test no-loop guarantee (knobs don't reset targets)
```

### Test Coverage

**Server-side tests**:
- Element discovery (primary vs. secondary classification)
- Label generation (all priority levels)
- Template rendering for multiple elements
- Knob ordering (matches demo source order)

**Client-side tests** (if possible):
- Mutation observer adds knobs when elements added
- Mutation observer removes knobs when elements removed
- Knob state syncs when attributes change externally
- No infinite loops when knobs change attributes

**Integration tests**:
- Complex compositions render correctly
- All element instances get knobs
- Performance acceptable with 50+ elements

## Acceptance Criteria

- [ ] Multiple instances of same element each have knob group
- [ ] Knob groups ordered to match demo source order
- [ ] Primary element expanded by default, others collapsed
- [ ] Element labels distinguish instances (using priority system)
- [ ] Mutation observer adds knobs when elements dynamically created
- [ ] Mutation observer removes knobs when elements removed
- [ ] Knob state syncs when attributes change via script
- [ ] No infinite loops (knobs don't reset target values on creation)
- [ ] Performance acceptable with 50 elements (no lag)
- [ ] Compositional components (like pf-tabs example) work correctly
