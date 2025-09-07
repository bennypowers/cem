# Interface Hierarchy Reference

This document provides detailed information about the interface hierarchy used in the MCP registry for type-safe custom elements manifest processing.

## Design Principles

### Interface Segregation
Each interface provides only the methods needed for its specific purpose, following the Interface Segregation Principle.

### Marker Methods
Some interfaces use marker methods (like `isSlot()`) to enable proper type discrimination when multiple interfaces might otherwise be structurally equivalent.

### Composition over Inheritance
Interfaces are composed to build more complex contracts rather than using deep inheritance hierarchies.

### Go Naming Conventions
All method names follow Go conventions:
- No "Get" prefixes (`Name()` not `GetName()`)
- Boolean methods use "Is" prefix when appropriate
- Methods are concise and descriptive

## Interface Tree

```
Item
├── Typed
│   ├── Attribute (+ Defaultable + Enumerable)
│   └── Event
├── Slot  
├── CssProperty
├── CssPart
└── CssState
```

## Base Interfaces

### Item

The foundational interface that all component items implement.

```go
type Item interface {
    Name() string        // The name/identifier of the item
    Description() string // Human-readable description
    Guidelines() []string // Usage guidelines extracted from descriptions
    Examples() []string  // Code examples for usage
    Kind() ItemKind     // The type of item (attribute, slot, etc.)
}
```

**Purpose**: Provides common functionality for all manifest items.

**Implementation Notes**:
- `Name()` returns the primary identifier (attribute name, slot name, etc.)
- `Description()` comes from manifest descriptions or JSDoc comments
- `Guidelines()` extracts usage rules from descriptions using keywords like "should", "must"
- `Examples()` provides generated or extracted code examples
- `Kind()` enables runtime type identification

### Typed

Extends Item to add type information for items that have TypeScript/JavaScript type annotations.

```go
type Typed interface {
    Item
    Type() string // The TypeScript/JavaScript type annotation
}
```

**Purpose**: Represents items that have type information in the manifest.

**Used by**: Attributes and Events (both have type annotations in manifests).

**Implementation Notes**:
- `Type()` returns the raw type string from manifest (e.g., `"string | number"`)
- Type parsing and validation should be done by consumers
- Empty string indicates no type information available

### Defaultable

Adds default value functionality for items that can have defaults.

```go
type Defaultable interface {
    Item
    Default() string // The default value as a string
}
```

**Purpose**: Represents items that can have default values.

**Used by**: Attributes (which can have default values in manifests).

**Implementation Notes**:
- Returns empty string if no default is specified
- Default values are stored as strings regardless of type
- Consumers should parse based on the item's type information

### Enumerable

Adds enumeration/union value functionality for items with constrained values.

```go
type Enumerable interface {
    Item
    Values() []string // Available values for union/enum types
}
```

**Purpose**: Represents items that have a constrained set of possible values.

**Used by**: Attributes with union types (e.g., `"small" | "medium" | "large"`).

**Implementation Notes**:
- Extracts values from union type strings automatically
- Returns empty slice if no constrained values exist
- Values are unquoted (e.g., `["small", "medium", "large"]`)

## Specialized Interfaces

### Attribute

The most complex interface, combining multiple capabilities for HTML attributes.

```go
type Attribute interface {
    Typed
    Defaultable
    Enumerable
    Required() bool // Whether the attribute is required
}
```

**Purpose**: Represents HTML attributes with full type, default, and constraint information.

**Capabilities**:
- Type information via `Typed`
- Default values via `Defaultable`  
- Enum/union values via `Enumerable`
- Requirement status via `Required()`

**Usage Example**:
```go
for _, attr := range element.Attributes() {
    fmt.Printf("Attribute: %s (%s)\n", attr.Name(), attr.Type())
    
    if attr.Required() {
        fmt.Printf("  REQUIRED\n")
    }
    
    if attr.Default() != "" {
        fmt.Printf("  Default: %s\n", attr.Default())
    }
    
    if len(attr.Values()) > 0 {
        fmt.Printf("  Values: %v\n", attr.Values())
    }
}
```

### Event

Represents custom events with type information and discrimination.

```go
type Event interface {
    Typed
    isEvent() // Marker method for type discrimination
}
```

**Purpose**: Represents custom events fired by elements.

**Marker Method**: `isEvent()` distinguishes events from attributes (both implement `Typed`).

**Usage Example**:
```go
for _, event := range element.Events() {
    fmt.Printf("Event: %s (%s)\n", event.Name(), event.Type())
    fmt.Printf("  Listen with: element.addEventListener('%s', handler)\n", event.Name())
}
```

### Slot

Represents content slots for element composition.

```go
type Slot interface {
    Item
    isSlot() // Marker method for type discrimination
}
```

**Purpose**: Represents named and default slots for content projection.

**Marker Method**: `isSlot()` enables discrimination from other simple Item implementations.

**Usage Example**:
```go
for _, slot := range element.Slots() {
    if slot.Name() == "" {
        fmt.Printf("Default slot: %s\n", slot.Description())
    } else {
        fmt.Printf("Named slot '%s': %s\n", slot.Name(), slot.Description())
        fmt.Printf("  Usage: <span slot=\"%s\">content</span>\n", slot.Name())
    }
}
```

### CssProperty

Represents CSS custom properties with syntax and inheritance information.

```go
type CssProperty interface {
    Item
    Syntax() string   // CSS syntax descriptor (e.g., "<color>", "<length>")
    Inherits() bool   // Whether the property inherits from parent elements
    Initial() string  // Initial/default value
}
```

**Purpose**: Represents CSS custom properties (CSS variables) exposed by elements.

**CSS Integration**: Provides metadata for proper CSS custom property usage.

**Usage Example**:
```go
for _, prop := range element.CssProperties() {
    fmt.Printf("CSS Property: %s\n", prop.Name())
    fmt.Printf("  Syntax: %s\n", prop.Syntax())
    fmt.Printf("  Initial: %s\n", prop.Initial())
    fmt.Printf("  Inherits: %t\n", prop.Inherits())
    
    // Generate CSS usage example
    fmt.Printf("  Usage: my-element { %s: value; }\n", prop.Name())
}
```

### CssPart

Represents CSS parts for styling element internals.

```go
type CssPart interface {
    Item
    isCssPart() // Marker method for type discrimination
}
```

**Purpose**: Represents CSS parts that can be styled from outside the element.

**CSS Integration**: Enables external styling of element internals via `::part()`.

**Usage Example**:
```go
for _, part := range element.CssParts() {
    fmt.Printf("CSS Part: %s\n", part.Name())
    fmt.Printf("  Description: %s\n", part.Description())
    fmt.Printf("  Usage: my-element::%s { styles }\n", part.Name())
}
```

### CssState

Represents CSS custom states for element state styling.

```go
type CssState interface {
    Item
    isCssState() // Marker method for type discrimination
}
```

**Purpose**: Represents custom CSS states that elements can expose.

**CSS Integration**: Enables state-based styling via `:--state-name`.

**Usage Example**:
```go
for _, state := range element.CssStates() {
    fmt.Printf("CSS State: %s\n", state.Name())
    fmt.Printf("  Description: %s\n", state.Description())
    fmt.Printf("  Usage: my-element:%s { styles }\n", state.Name())
}
```

## Implementation Patterns

### Factory Functions

Each interface has a corresponding factory function that creates instances from manifest data:

```go
func NewAttributeItem(attr M.Attribute, guidelines []string, examples []string) Attribute
func NewSlotItem(slot M.Slot, guidelines []string, examples []string) Slot
func NewEventItem(event M.Event, guidelines []string, examples []string) Event
func NewCssPropertyItem(prop M.CssCustomProperty, guidelines []string, examples []string) CssProperty
func NewCssPartItem(part M.CssPart, guidelines []string, examples []string) CssPart
func NewCssStateItem(state M.CssCustomState, guidelines []string, examples []string) CssState
```

### Implementation Structs

Each interface is implemented by a corresponding struct:

- `AttributeItem` implements `Attribute`
- `SlotItem` implements `Slot`
- `EventItem` implements `Event`
- `CssPropertyItem` implements `CssProperty`
- `CssPartItem` implements `CssPart`
- `CssStateItem` implements `CssState`

All implementation structs embed `BaseItem` for common functionality.

### Type Assertions

When working with the unified `Items` slice, use type assertions for specific functionality:

```go
for _, item := range element.Items {
    switch item.Kind() {
    case KindAttribute:
        if attr, ok := item.(Attribute); ok {
            // Work with attribute-specific methods
            fmt.Printf("Type: %s, Required: %t\n", attr.Type(), attr.Required())
        }
    case KindCssProperty:
        if prop, ok := item.(CssProperty); ok {
            // Work with CSS property-specific methods
            fmt.Printf("Syntax: %s, Inherits: %t\n", prop.Syntax(), prop.Inherits())
        }
    }
}
```

### JSON Marshaling

All items support JSON marshaling with type-specific fields:

```go
// Attribute JSON includes type, default, required, values
{
  "kind": "attribute",
  "name": "variant",
  "type": "\"primary\" | \"secondary\"",
  "default": "\"primary\"",
  "required": false,
  "values": ["primary", "secondary"]
}

// CSS Property JSON includes syntax, inherits, initial
{
  "kind": "css-property", 
  "name": "--button-color",
  "syntax": "<color>",
  "inherits": false,
  "initial": "blue"
}
```

## Best Practices

### Interface Usage
- Always use the most specific interface for your needs
- Prefer type assertions over runtime reflection
- Check `Kind()` before type assertions for safety

### Error Handling
- Check for nil returns from finder functions
- Handle type assertion failures gracefully
- Provide context in error messages

### Performance
- Cache interface instances when possible
- Use type-specific accessor methods for filtering
- Avoid unnecessary type assertions in hot paths

### Extension
- Add marker methods when creating new interfaces
- Follow naming conventions (no "Get" prefixes)
- Implement JSON marshaling for new types
- Add comprehensive tests for new interfaces