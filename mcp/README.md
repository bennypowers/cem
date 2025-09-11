# CEM MCP Registry

The MCP (Model Context Protocol) registry provides an interface-based system for accessing and working with Custom Elements Manifest data in an AI-native way.

## Overview

The MCP registry transforms raw custom elements manifest data into a structured, type-safe interface hierarchy that enables intelligent HTML generation and component understanding.

## Key Features

- **Interface-based polymorphism**: Unified `Item` interface with specialized implementations
- **Type-safe accessors**: Compile-time guarantees for different item types
- **JSON marshaling**: Full API compatibility for external consumption
- **Thread-safe operations**: Concurrent access with proper locking
- **Caching layer**: Performance optimization for repeated queries
- **Go best practices**: No "Get" prefixes, proper naming conventions

## Quick Start

```go
package main

import (
    "fmt"
    "bennypowers.dev/cem/mcp"
    "bennypowers.dev/cem/workspace"
)

func main() {
    // Create workspace and registry
    ws := workspace.NewFileSystemWorkspaceContext("/path/to/project")
    registry, err := mcp.NewRegistry(ws)
    if err != nil {
        panic(err)
    }
    
    // Load manifests
    err = registry.LoadManifests()
    if err != nil {
        panic(err)
    }
    
    // Get element information
    element, err := registry.GetElementInfo("my-button")
    if err != nil {
        panic(err)
    }
    
    // Access different item types
    fmt.Printf("Attributes: %d\n", len(element.Attributes()))
    fmt.Printf("Slots: %d\n", len(element.Slots()))
    fmt.Printf("CSS Properties: %d\n", len(element.CssProperties()))
    
    // Work with specific items
    for _, attr := range element.Attributes() {
        fmt.Printf("Attribute: %s (%s)\n", attr.Name(), attr.Type())
        if attr.Default() != "" {
            fmt.Printf("  Default: %s\n", attr.Default())
        }
    }
}
```

## Architecture

### Interface Hierarchy

The registry uses a carefully designed interface hierarchy that provides both flexibility and type safety:

```
Item (base interface)
├── Typed (adds Type() method)
│   ├── Attribute (adds Default(), Required(), Values())
│   └── Event (adds isEvent() marker)
├── Slot (adds isSlot() marker)
├── CssProperty (adds Syntax(), Inherits(), Initial())
├── CssPart (adds isCssPart() marker)
└── CssState (adds isCssState() marker)
```

### Core Types

- **Registry**: Main entry point, manages manifest loading and caching
- **ElementInfo**: Container for all element-related information
- **Item interfaces**: Type-safe representations of manifest components
- **Implementation structs**: Concrete types that implement the interfaces

## Interface Reference

### Base Interfaces

#### Item
```go
type Item interface {
    Name() string
    Description() string
    Guidelines() []string
    Examples() []string
    Kind() ItemKind
}
```

#### Typed
```go
type Typed interface {
    Item
    Type() string
}
```

#### Defaultable
```go
type Defaultable interface {
    Item
    Default() string
}
```

#### Enumerable
```go
type Enumerable interface {
    Item
    Values() []string
}
```

### Specialized Interfaces

#### Attribute
```go
type Attribute interface {
    Typed
    Defaultable
    Enumerable
    Required() bool
}
```

#### Event
```go
type Event interface {
    Typed
    isEvent() // Marker method
}
```

#### Slot
```go
type Slot interface {
    Item
    isSlot() // Marker method
}
```

#### CssProperty
```go
type CssProperty interface {
    Item
    Syntax() string
    Inherits() bool
    Initial() string
}
```

#### CssPart
```go
type CssPart interface {
    Item
    isCssPart() // Marker method
}
```

#### CssState
```go
type CssState interface {
    Item
    isCssState() // Marker method
}
```

## Usage Examples

### Working with Attributes

```go
element, _ := registry.GetElementInfo("my-button")

for _, attr := range element.Attributes() {
    fmt.Printf("Attribute: %s\n", attr.Name())
    fmt.Printf("  Type: %s\n", attr.Type())
    fmt.Printf("  Required: %t\n", attr.Required())
    
    if attr.Default() != "" {
        fmt.Printf("  Default: %s\n", attr.Default())
    }
    
    if len(attr.Values()) > 0 {
        fmt.Printf("  Values: %v\n", attr.Values())
    }
}
```

### Working with CSS Properties

```go
element, _ := registry.GetElementInfo("my-card")

for _, prop := range element.CssProperties() {
    fmt.Printf("CSS Property: %s\n", prop.Name())
    fmt.Printf("  Syntax: %s\n", prop.Syntax())
    fmt.Printf("  Initial: %s\n", prop.Initial())
    fmt.Printf("  Inherits: %t\n", prop.Inherits())
}
```

### Working with Slots

```go
element, _ := registry.GetElementInfo("my-dialog")

for _, slot := range element.Slots() {
    if slot.Name() == "" {
        fmt.Println("Default slot:", slot.Description())
    } else {
        fmt.Printf("Named slot '%s': %s\n", slot.Name(), slot.Description())
    }
}
```

### Unified Item Processing

```go
element, _ := registry.GetElementInfo("my-element")

// Process all items regardless of type
for _, item := range element.Items {
    fmt.Printf("%s: %s (%s)\n", item.Kind(), item.Name(), item.Description())
    
    // Type-specific processing
    switch item.Kind() {
    case mcp.KindAttribute:
        if attr, ok := item.(mcp.Attribute); ok {
            fmt.Printf("  Type: %s\n", attr.Type())
        }
    case mcp.KindCssProperty:
        if prop, ok := item.(mcp.CssProperty); ok {
            fmt.Printf("  Syntax: %s\n", prop.Syntax())
        }
    }
}
```

## JSON Marshaling

All item types support JSON marshaling for API compatibility:

```go
element, _ := registry.GetElementInfo("my-button")
data, err := json.Marshal(element)
if err != nil {
    panic(err)
}

// The JSON will include all items with their type-specific fields
fmt.Println(string(data))
```

Example JSON output:
```json
{
  "tagName": "my-button",
  "name": "my-button",
  "items": [
    {
      "kind": "attribute",
      "name": "variant",
      "description": "Button variant",
      "type": "\"primary\" | \"secondary\"",
      "default": "\"primary\"",
      "required": false,
      "values": ["primary", "secondary"]
    },
    {
      "kind": "slot",
      "name": "",
      "description": "Button content"
    },
    {
      "kind": "css-property",
      "name": "--button-color",
      "description": "Button text color",
      "syntax": "<color>",
      "inherits": false,
      "initial": "currentColor"
    }
  ]
}
```

## Performance Considerations

### Caching

The registry implements a thread-safe caching layer:
- Element info is cached after first access
- Cache is invalidated when manifests are reloaded
- Concurrent access is protected with read/write mutexes

### Memory Usage

- Interface-based design minimizes memory overhead
- Shared base implementations reduce duplication
- Lazy loading of element information

### Thread Safety

- All public methods are thread-safe
- Read/write locks protect concurrent access
- Deadlock prevention in `GetAllElements()`

## Error Handling

The registry provides clear error messages for common scenarios:

```go
// Element not found
element, err := registry.GetElementInfo("non-existent")
if err != nil {
    // Error: "failed to get element info for \"non-existent\": element not found in registry"
}

// Manifest loading failures
err = registry.LoadManifests()
if err != nil {
    // Error includes workspace path and underlying cause
}
```

## Contributing

When extending the registry with new item types:

1. Add the interface to the hierarchy
2. Create a corresponding implementation struct
3. Add a factory function following naming conventions
4. Implement JSON marshaling
5. Add marker methods if needed for type discrimination
6. Update tests and documentation

## See Also

- [MCP CLAUDE.md](./CLAUDE.md) - Complete implementation plan
- [Registry Tests](./registry_test.go) - Usage examples and test cases
- [Interface Implementation](./registry.go) - Full source code