# Basic Knobs

Knobs provide interactive controls for manipulating custom element attributes, properties, and CSS custom properties during development.

**Scope**: This document covers basic knobs for **single elements**. For multiple elements and complex compositions, see `51-KNOBS-ADVANCED.md`. For custom templates, see `52-KNOBS-CUSTOM.md`.

## Overview

Knobs should be a pure dev-server feature that reads from CEM data, rather than something embedded in users' demo HTML. This keeps demos simple.

**Key principle**: Server generates knob UI from manifest data, client-side code handles user interactions.

## Server-Side Rendering

### Knob Generation Process

1. **Read manifest** for element's attributes, properties, CSS custom properties
2. **Parse demo HTML** to extract current values
3. **Render appropriate control** for each attribute type:
   - Boolean → `<input type="checkbox">`
   - Union type (`'a' | 'b'`) → `<select>`
   - String → `<input type="text">`
   - Number → `<input type="number">`
   - Enum array → Multiple checkboxes

### Type Mapping

| CEM Type | HTML Control | Notes |
|----------|--------------|-------|
| `boolean` | `<input type="checkbox">` | Checked = attribute present |
| `string` | `<input type="text">` | Direct value binding |
| `number` | `<input type="number">` | Validates numeric input |
| `'a' \| 'b'` | `<select>` | Union types become options |
| `string[]` | Multiple `<input type="checkbox">` | Each value is checkbox |

### Data Model

The server provides knob templates with CEM data:

```go
type KnobGroupData struct {
    TagName     string
    Label       string
    Attributes  []AttributeKnobData
    Properties  []PropertyKnobData
    CSSProps    []CSSPropKnobData
}

type AttributeKnobData struct {
    Name         string   // attribute name
    Type         string   // "boolean" | "string" | "number" | "enum"
    CurrentValue string   // current attribute value from demo HTML
    EnumValues   []string // for type=enum
    Description  string   // from CEM
    Default      string   // default value from CEM
}
```

## Client-Side Behavior

Embedded `__cem-serve-chrome.js` provides event-based communication between knobs and demo elements.

### Event System

```typescript
abstract class KnobChangeEvent extends Event {
  constructor(public value: unknown) {
    super('knob-change', { bubbles: true, cancelable: true });
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

export class KnobPropertyChangeEvent extends KnobChangeEvent {
  constructor(
    public property: string,
    public value: unknown,
  ) {
    super(value);
  }
}

export class KnobCssCustomPropertyChangeEvent extends KnobChangeEvent {
  constructor(
    public cssProperty: string,
    public value: string,
  ) {
    super(value);
  }
}
```

### Chrome Element Integration

```typescript
class CemServeChrome extends HTMLElement {
  // ...

  constructor() {
    super();
    this.addEventListener('knob-change', this.#onKnobChange);
  }

  /**
   * Each knob type is a custom element e.g. cem-serve-knob-attribute
   * Each knob element class is responsible for constructing and
   * firing its particular change event
   */
  #onKnobChange(event: KnobChangeEvent) {
    switch (event.knob.type) {
      case 'attribute':
        return this.#demo.setAttribute(event.attribute, event.value);
      case 'property':
        return this.#demo.setProperty(event.property, event.value);
      case 'css-property':
        return this.#demo.setCssCustomProperty(event.cssProperty, event.value);
    }
  }

  // ...
}
```

### Demo Element Helper

```typescript
class CemServeDemo extends HTMLElement {
  // Helper methods for knobs to manipulate demo elements

  setAttribute(attribute: string, value: string): boolean {
    const target = this.querySelector(`[${attribute}]`);
    if (target) {
      target.setAttribute(attribute, value);
      return true;
    }
    return false;
  }

  setProperty(property: string, value: unknown): boolean {
    const target = this.querySelector('[data-knob-target]');
    if (target) {
      target[property] = value;
      return true;
    }
    return false;
  }

  setCssCustomProperty(cssProperty: string, value: string): boolean {
    const target = this.querySelector('[data-knob-target]');
    if (target) {
      target.style.setProperty(cssProperty, value);
      return true;
    }
    return false;
  }
}
```

## Configuration

### Config File

Under `.serve.knobs`:

```yaml
serve:
  knobs:
    enabled: true
    types: [attributes, properties, css-properties]  # optional: filter which knob types to show
```

### Query Parameters

Control knobs via URL query parameters (see `10-URL-REWRITING.md`):

```
/components/pf-button/demo/variants/?disable-knobs=all
  → No knobs

/components/pf-button/demo/variants/?disable-knobs[]=properties
  → All knobs except properties
```

## Knob Types

### Attribute Knobs

Control HTML attributes on custom elements.

**Example**:
```html
<!-- Knob for boolean attribute -->
<label>
  <input type="checkbox" name="disabled" checked>
  disabled
</label>

<!-- Knob for enum attribute -->
<label>
  variant:
  <select name="variant">
    <option value="primary">primary</option>
    <option value="secondary">secondary</option>
  </select>
</label>
```

**Behavior**:
- Boolean: Checkbox checked = attribute present
- String/Number: Input value = attribute value
- Enum: Select option = attribute value

### Property Knobs

Control JavaScript properties on custom elements.

**Example**:
```html
<label>
  items (property):
  <input type="text" name="items" value='["a", "b"]'>
</label>
```

**Behavior**:
- Parse string as JSON for complex types
- Direct assignment for primitives
- Validate before setting (show error if invalid JSON)

### CSS Custom Property Knobs

Control CSS custom properties (CSS variables) on custom elements.

**Example**:
```html
<label>
  --button-color:
  <input type="color" name="--button-color" value="#0066cc">
</label>
```

**Behavior**:
- Set via `element.style.setProperty(name, value)`
- Supports all CSS value types
- Color type → color picker
- Length type → text input with unit suffix

## Testing Strategy

Fixture-based tests (per CLAUDE.md):

```
serve/knobs/test-fixtures/
  ├── simple/
  │   ├── demo.html (single element, few attributes)
  │   ├── manifest.json (element definition)
  │   ├── expected-knobs.html (rendered knob HTML)
  │   └── knobs-test.go (test knob generation)
  │
  ├── types/
  │   ├── boolean-demo.html
  │   ├── string-demo.html
  │   ├── number-demo.html
  │   ├── enum-demo.html
  │   └── types-test.go (test all type mappings)
  │
  └── integration/
      ├── demo.html
      ├── test.js (simulates knob interactions)
      └── integration-test.go (test client-server round trip)
```

### Test Coverage

**Server-side tests**:
- Knob generation from manifest data
- Type mapping (boolean, string, number, enum)
- Current value extraction from demo HTML
- Template rendering correctness

**Client-side tests** (if possible):
- KnobChangeEvent firing on input changes
- Event payload contains correct data
- Demo element receives attribute/property/CSS changes
- Invalid inputs handled gracefully

**Integration tests**:
- End-to-end: render knob → change value → update demo
- Config filtering (disable-knobs parameter)
- Query parameter handling

## Acceptance Criteria

- [ ] Knobs render for all element attributes in manifest
- [ ] Knob types match CEM attribute types (boolean, string, number, enum)
- [ ] Knob inputs reflect current demo values
- [ ] Changing knob updates demo element in real-time
- [ ] Boolean knobs: checkbox controls attribute presence
- [ ] Enum knobs: select shows all union type options
- [ ] Property knobs: can set JavaScript properties
- [ ] CSS property knobs: can set CSS custom properties
- [ ] Config `knobs.enabled` controls feature on/off
- [ ] Query param `?disable-knobs=all` hides knobs
- [ ] Knobs work with basic demos (single element)
- [ ] Tests pass (server-side, client-side, integration)
