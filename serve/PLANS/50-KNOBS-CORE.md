# Basic Knobs

**Status**: ✅ **COMPLETE** (with enhancements beyond original plan)

## TDD First Steps

Phase 5 has no stubs in existing tests. Follow standard TDD: write tests first, implement to make them pass.

**Prerequisites**: Complete Phase 4 remaining work (CSS glob filtering, YAML config) before starting Phase 5a. ✅ Done

**Approach**:
1. Write failing tests using fixture/golden patterns
2. Implement server-side knob generation
3. Implement client-side knob behavior
4. Verify integration tests pass

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

**Follow established patterns from Phases 1-4:**

### Use testutil Package
- Use `testutil.NewFixtureFS()` for in-memory filesystems
- Use `testutil.LoadFixtureFile()` to load test data
- Place fixtures in `internal/platform/testutil/fixtures/knobs/`
- Use golden file pattern for expected outputs

### Test Structure
Follow the patterns from `serve/middleware/routes/`, `serve/middleware/transform/`:

```
serve/middleware/knobs/
  ├── knobs.go (main implementation)
  ├── knobs_test.go (unit tests)
  ├── testdata/
  │   ├── simple-element/
  │   │   ├── manifest.json (element definition)
  │   │   └── expected-basic.html (golden file for rendered knobs)
  │   ├── type-mappings/
  │   │   ├── manifest-boolean.json
  │   │   ├── manifest-enum.json
  │   │   └── expected-*.html (golden files)
  │   └── chrome-rendering/
  │       └── expected-*.html (rendered chrome with knobs)
```

### Testing Fixtures
Use the fixture pattern established in Phase 4:

```go
// Example test
func TestKnobRendering_SimpleElement(t *testing.T) {
    mfs := testutil.NewFixtureFS()
    mfs.AddFile("/test/manifest.json", manifestJSON, 0644)
    mfs.AddFile("/test/demo.html", demoHTML, 0644)

    // Test knob generation
    knobs := GenerateKnobs(manifest, demo)

    // Compare against golden file
    expected := testutil.LoadFixtureFile(t, "knobs/simple-element/expected-basic.html")
    if diff := cmp.Diff(expected, knobs); diff != "" {
        t.Errorf("Knobs mismatch (-want +got):\n%s", diff)
    }
}
```

### Update Flag Support
Use `--update` flag to regenerate golden files (like Phase 3/4 tests):

```bash
go test ./serve/middleware/knobs/... --update
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

- [x] Knobs render for all element attributes in manifest ✅
- [x] Knob types match CEM attribute types (boolean, string, number, enum) ✅
- [x] Knob inputs reflect current demo values ✅
- [x] Changing knob updates demo element in real-time ✅
- [x] Boolean knobs: checkbox controls attribute presence ✅
- [x] Enum knobs: select shows all union type options ✅
- [x] Property knobs: can set JavaScript properties ✅
- [x] CSS property knobs: can set CSS custom properties ✅
- [x] Config `knobs.enabled` controls feature on/off ✅
- [x] Query param `?disable-knobs=all` hides knobs ✅
- [x] Knobs work with basic demos (single element) ✅
- [x] Tests pass (server-side, client-side, integration) ✅

## Features Beyond Original Plan

The implementation includes several enhancements not in the original specification:

### Type Badges
Each knob displays a type badge (e.g., `string`, `enum`, `boolean`, `color`) next to the knob name, providing immediate visual feedback about the expected value type.

### Dual Color Input
Color-type knobs feature both a color picker and a parallel text input field, allowing users to:
- Use the visual color picker for standard hex colors
- Enter CSS variable names (design tokens) like `var(--color-primary)`
- Input any valid CSS color value (rgb, hsl, named colors)

### Automatic Deduplication
Server-side logic prevents duplicate knobs when inheritance creates multiple manifest entries:
- **Property-over-attribute**: Skips attributes when same-named properties exist
- **Type specificity**: Prefers attributes with specific types over `null` types

### State Synchronization
Knobs intelligently sync with demo element state:
- Read actual element values on page load (not just defaults)
- Prevent browser form restoration conflicts via `autocomplete="off"`
- No `name` attributes to avoid browser state persistence
- Initialize to current state, then allow user control

### Markdown Descriptions with Syntax Highlighting
Knob descriptions support full GitHub-Flavored Markdown with:
- Code blocks with syntax highlighting (chroma)
- Themable highlighting via CSS custom properties
- Inline code, lists, emphasis, links
- Proper typography and spacing
