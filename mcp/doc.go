/*
Package mcp provides an interface-based registry for Custom Elements Manifest data,
designed for AI-native component intelligence and HTML generation.

# Overview

The mcp package transforms raw custom elements manifest data into a structured,
type-safe interface hierarchy that enables intelligent HTML generation and
component understanding. It's specifically designed to work with Model Context
Protocol (MCP) servers for AI assistance.

# Architecture

The package uses a carefully designed interface hierarchy that provides both
flexibility and type safety:

	Item (base interface)
	├── Typed (adds Type() method)
	│   ├── Attribute (adds Default(), Required(), Values())
	│   └── Event (adds isEvent() marker)
	├── Slot (adds isSlot() marker)
	├── CssProperty (adds Syntax(), Inherits(), Initial())
	├── CssPart (adds isCssPart() marker)
	└── CssState (adds isCssState() marker)

# Basic Usage

	import "bennypowers.dev/cem/mcp"
	import "bennypowers.dev/cem/workspace"

	// Create registry
	ws := workspace.NewFileSystemWorkspaceContext("/path/to/project")
	registry, err := mcp.NewRegistry(ws)
	if err != nil {
		return err
	}

	// Load manifests
	err = registry.LoadManifests()
	if err != nil {
		return err
	}

	// Get element information
	element, err := registry.GetElementInfo("my-button")
	if err != nil {
		return err
	}

	// Access different item types
	for _, attr := range element.Attributes() {
		fmt.Printf("Attribute: %s (%s)\n", attr.Name(), attr.Type())
	}

# Interface Design

All component items implement the base Item interface, which provides common
functionality like Name(), Description(), Guidelines(), and Examples().
Specialized interfaces extend this base to add type-specific methods.

The design follows Go best practices:
- No "Get" prefixes on method names
- Interface segregation principle
- Marker methods for type discrimination
- Proper error handling and documentation

# Thread Safety

The registry is designed for concurrent use:
- All public methods are thread-safe
- Read/write locks protect shared state
- Caching layer prevents race conditions
- Deadlock prevention in complex operations

# JSON Marshaling

All item types support JSON marshaling for API compatibility. The JSON output
includes all relevant fields based on the item type, making it suitable for
external consumption by AI systems or web APIs.

# Performance

The registry includes several performance optimizations:
- Thread-safe caching of converted elements
- Lazy loading of element information
- Efficient interface-based filtering
- Memory-efficient shared implementations

# Error Handling

The package provides clear, actionable error messages:
- Element not found errors include the requested tag name
- Manifest loading errors include workspace context
- All errors are wrapped with sufficient context for debugging

# Extension Points

The interface-based design makes it easy to extend with new item types:
1. Define the interface extending Item or a specialized interface
2. Create an implementation struct
3. Add factory functions following naming conventions
4. Implement JSON marshaling
5. Add tests and documentation

# Example: Working with Attributes

	element, err := registry.GetElementInfo("my-button")
	if err != nil {
		return err
	}

	for _, attr := range element.Attributes() {
		fmt.Printf("Attribute: %s\n", attr.Name())
		fmt.Printf("  Type: %s\n", attr.Type())
		fmt.Printf("  Required: %t\n", attr.Required())

		if attr.Default() != "" {
			fmt.Printf("  Default: %s\n", attr.Default())
		}

		if len(attr.Values()) > 0 {
			fmt.Printf("  Enum values: %v\n", attr.Values())
		}
	}

# Example: Processing All Items

	element, err := registry.GetElementInfo("my-element")
	if err != nil {
		return err
	}

	// Process all items regardless of type
	for _, item := range element.Items {
		fmt.Printf("%s: %s\n", item.Kind(), item.Name())

		// Type-specific processing using type assertions
		switch item.Kind() {
		case KindAttribute:
			if attr, ok := item.(Attribute); ok {
				fmt.Printf("  Type: %s\n", attr.Type())
			}
		case KindCssProperty:
			if prop, ok := item.(CssProperty); ok {
				fmt.Printf("  Syntax: %s\n", prop.Syntax())
			}
		}
	}

# Example: JSON API Response

	element, err := registry.GetElementInfo("my-card")
	if err != nil {
		return err
	}

	// Marshal to JSON for API response
	data, err := json.Marshal(element)
	if err != nil {
		return err
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

This package is designed to be the foundation for AI-assisted web component
development, providing rich, structured data that enables intelligent HTML
generation with proper slot usage, attribute validation, and CSS integration.
*/
package mcp
