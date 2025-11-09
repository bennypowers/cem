# Custom Knob Templates (OPTIONAL)

This document covers the custom template system that allows users to override default knob UI with their own components.

**Note**: This is an OPTIONAL feature. Basic knobs (50-KNOBS-CORE.md) and advanced knobs (51-KNOBS-ADVANCED.md) provide full functionality without custom templates.

## Overview

Custom templates allow users to:

**HIGH PRIORITY Use Cases**:
- Add specialized controls (icon pickers, color palette pickers) for specific enum types
- Create domain-specific knobs (e.g., breakpoint selector for responsive design)

**LOW PRIORITY Use Cases**:
- Bring their own design system components (`<rh-switch>`, `<pf-select>`)
- Match knobs UI to their brand/style guide

## Template Files

Users can override individual knob type templates via `--template-dir` flag:

```bash
cem serve --template-dir ./templates/
```

### Available Template Files

- `demo.html` - Main demo page chrome (rarely overridden)
- `knob-attribute-boolean.html` - Boolean attribute control (checkbox/switch)
- `knob-attribute-enum.html` - Enum/union type select (select/radio)
- `knob-attribute-string.html` - Text input
- `knob-attribute-number.html` - Number input

**Fallback**: If a template file is not provided, the default built-in template is used.

## Template Contract

Custom templates must manage their own internal state:

1. **Listen for relevant change events** from form inputs
2. **Derive new values** from the state of form inputs
3. **Construct and fire knob-change events** (see Event API below)

### Event API

Custom templates must fire `KnobChangeEvent` subtypes when values change:

```typescript
// Base event (abstract)
abstract class KnobChangeEvent extends Event {
  constructor(public value: unknown) {
    super('knob-change', { bubbles: true, cancelable: true });
  }
}

// Attribute knob event
class KnobAttributeChangeEvent extends KnobChangeEvent {
  constructor(
    public attribute: string,
    public value: string,
  ) {
    super(value);
  }
}

// Similar events for properties, CSS vars, etc.
```

**Requirements**:
- Event type: `'knob-change'`
- Event must bubble: `{ bubbles: true }`
- Event payload must include attribute/property name and new value

## Template Implementation Pattern

In most cases, users will wrap their custom control in a custom element:

### Default Result (Built-in)

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

### Custom Result

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

### Custom Element Example

```typescript
// User's custom knob wrapper
class UserCustomKnobAttribute extends HTMLElement {
  connectedCallback() {
    const input = this.shadowRoot.querySelector('my-switch');
    const attributeName = this.getAttribute('attribute');

    input.addEventListener('change', (e) => {
      const event = new KnobAttributeChangeEvent(
        attributeName,
        e.target.checked ? '' : null
      );
      this.dispatchEvent(event);
    });
  }
}
```

## Configuration

### Via Config File

```yaml
serve:
  knobs:
    enabled: true
    templates: ./templates/knobs/*.html
```

### Via CLI Flag

```bash
cem serve --template-dir ./templates/
```

**Precedence**: CLI flag overrides config file.

## Security Considerations

Custom templates are **trusted code**.

**Rationale**:
- In the majority case, the dev server runs locally
- Server doesn't have access to sensitive domains
- Templates are user's own code for their own development

**Warning for documentation**:
Users combining custom templates with proxies to sensitive domains should audit template code for security issues.

**Not a concern for initial release**: XSS in templates only affects the user's own dev environment.

## Template Data Model

Templates receive the same data as built-in templates:

```go
type KnobAttributeData struct {
    Name         string   // attribute name
    Type         string   // "boolean" | "string" | "number" | "enum"
    CurrentValue string   // current attribute value
    EnumValues   []string // for type=enum
    Description  string   // from CEM
    Default      string   // default value from CEM
}
```

**Template access**:
```html
<input type="checkbox"
       name="{{.Name}}"
       {{if .CurrentValue}}checked{{end}}>
```

## Template Testing

### For Template Authors

**Issue**: Users can provide arbitrary custom templates with no built-in testing.

**Requirements** (for template authors):
1. Template must fire correct events
2. Template must handle all CEM data fields
3. Template must not break chrome functionality

### Validation Strategy

**Server-side validation** (basic):
- Check template is valid HTML
- Check template doesn't contain obvious XSS vectors (if we add security model later)

**Runtime validation** (basic):
- Ensure template fires `knob-change` events
- Log warning if event format is incorrect

### Testing Harness (Long-Term Goal, Not Initial Release)

Provide testing utilities for template authors:

```bash
cem serve test-template ./templates/knob-attribute-boolean.html
```

Would validate:
- Template renders without errors
- Template fires correct events
- Template handles edge cases (empty values, long strings, etc.)

**Not a priority for initial release** - document template contract clearly instead.

## Use Case Examples

### Example 1: Color Palette Picker

**Scenario**: Component has `color` attribute with predefined palette.

**Custom template** (`knob-attribute-enum.html`):
```html
<color-palette-picker attribute="{{.Name}}" value="{{.CurrentValue}}">
  {{range .EnumValues}}
  <color-swatch value="{{.}}"></color-swatch>
  {{end}}
</color-palette-picker>

<script>
  class ColorPalettePicker extends HTMLElement {
    connectedCallback() {
      this.addEventListener('click', (e) => {
        if (e.target.matches('color-swatch')) {
          const event = new KnobAttributeChangeEvent(
            this.getAttribute('attribute'),
            e.target.getAttribute('value')
          );
          this.dispatchEvent(event);
        }
      });
    }
  }
  customElements.define('color-palette-picker', ColorPalettePicker);
</script>
```

### Example 2: Icon Picker

**Scenario**: Component has `icon` attribute with icon library.

**Custom template**:
```html
<icon-picker attribute="icon" value="{{.CurrentValue}}">
  <!-- Grid of icons from library -->
</icon-picker>
```

### Example 3: Design System Components

**Scenario**: User wants knobs to match their design system.

**Custom template** (`knob-attribute-boolean.html`):
```html
<rh-switch attribute="{{.Name}}" {{if .CurrentValue}}checked{{end}}>
  <span slot="label">{{.Name}}</span>
</rh-switch>

<script type="module">
  import '@rhds/elements/rh-switch/rh-switch.js';

  document.querySelector('rh-switch').addEventListener('change', (e) => {
    const event = new KnobAttributeChangeEvent(
      e.target.getAttribute('attribute'),
      e.target.checked ? '' : null
    );
    e.target.dispatchEvent(event);
  });
</script>
```

## Limitations

### Known Limitations (Document Clearly)

1. **Templates must provide scripts** - No declarative-only templates
2. **No validation** - Server doesn't validate template correctness
3. **No sandboxing** - Templates run in same context as chrome
4. **Breaking changes** - Template API may change between versions

### Future Enhancements (Not Initial Release)

1. **Template validation** - Check event contract at runtime
2. **Template testing harness** - Automated testing for custom templates
3. **Template gallery** - Community-contributed templates
4. **Template versioning** - Compatibility across cem versions

## Acceptance Criteria

### Template Loading and Fallback
- [ ] Users can override templates via `--template-dir` flag
- [ ] Server loads custom templates from specified directory
- [ ] Fallback to default built-in templates when custom not provided
- [ ] Template file naming convention documented (e.g., `knob-attribute-boolean.html`)
- [ ] Config file supports `serve.knobs.templates` path
- [ ] CLI flag `--template-dir` overrides config file setting

### Data Model Contract
- [ ] Custom templates receive correct `KnobGroupData` structure from server
- [ ] Templates can access all CEM data fields (name, type, current value, enum values, description, default)
- [ ] Template variables work correctly with Go html/template syntax
- [ ] Data model documented with complete type definitions

### Event Contract Validation
- [ ] Custom templates can fire `KnobChangeEvent` subtypes
- [ ] Server validates event type is `'knob-change'`
- [ ] Server validates event bubbles (`{ bubbles: true }`)
- [ ] Server validates event payload includes attribute/property name and value
- [ ] Runtime validation logs warnings for malformed events
- [ ] Event contract fully documented with TypeScript definitions

### Documentation and Examples
- [ ] Documentation clearly explains template contract
- [ ] Example templates provided for color picker use case
- [ ] Example templates provided for icon picker use case
- [ ] Example templates provided for design system integration (rh-switch, pf-select)
- [ ] Troubleshooting guide covers common issues (events not firing, data missing)
- [ ] Security warning documented for proxy scenarios

## Documentation Requirements

Must document:
1. **Event contract** - `KnobChangeEvent` types and payloads
2. **Data model** - What data templates receive
3. **Template naming** - Which files override which knobs
4. **Security model** - Templates are trusted, warnings for proxies
5. **Example templates** - Color picker, icon picker, design system
6. **Troubleshooting** - Common issues (events not firing, data missing)
