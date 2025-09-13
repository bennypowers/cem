# MCP Registry Usage Examples

This document provides comprehensive examples of using the MCP registry for various common scenarios.

## Table of Contents

- [Basic Setup](#basic-setup)
- [Working with Attributes](#working-with-attributes)
- [Working with Slots](#working-with-slots)
- [Working with Events](#working-with-events)
- [Working with CSS](#working-with-css)
- [Unified Item Processing](#unified-item-processing)
- [JSON API Integration](#json-api-integration)
- [Advanced Patterns](#advanced-patterns)
- [Error Handling](#error-handling)

## Basic Setup

### Creating a Registry

```go
package main

import (
    "fmt"
    "log"
    
    "bennypowers.dev/cem/mcp"
    "bennypowers.dev/cem/workspace"
)

func main() {
    // Create workspace context
    ws := workspace.NewFileSystemWorkspaceContext("/path/to/your/project")
    
    // Initialize workspace
    if err := ws.Init(); err != nil {
        log.Fatal("Failed to initialize workspace:", err)
    }
    
    // Create registry
    registry, err := mcp.NewRegistry(ws)
    if err != nil {
        log.Fatal("Failed to create registry:", err)
    }
    
    // Load all manifests from the workspace
    if err := registry.LoadManifests(); err != nil {
        log.Fatal("Failed to load manifests:", err)
    }
    
    fmt.Println("Registry initialized successfully!")
}
```

### Getting Element Information

```go
func getElementInfo(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        log.Printf("Failed to get element info for %s: %v", tagName, err)
        return
    }
    
    fmt.Printf("Element: %s\n", element.TagName)
    fmt.Printf("Description: %s\n", element.Description)
    fmt.Printf("Total items: %d\n", len(element.Items))
}
```

## Working with Attributes

### Basic Attribute Information

```go
func examineAttributes(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Attributes for %s:\n", tagName)
    for _, attr := range element.Attributes() {
        fmt.Printf("  %s: %s\n", attr.Name(), attr.Type())
        if attr.Description() != "" {
            fmt.Printf("    Description: %s\n", attr.Description())
        }
    }
}
```

### Working with Attribute Defaults and Values

```go
func analyzeAttributeDetails(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, attr := range element.Attributes() {
        fmt.Printf("Attribute: %s\n", attr.Name())
        fmt.Printf("  Type: %s\n", attr.Type())
        fmt.Printf("  Required: %t\n", attr.Required())
        
        // Check for default value
        if defaultVal := attr.Default(); defaultVal != "" {
            fmt.Printf("  Default: %s\n", defaultVal)
        }
        
        // Check for enumerated values (union types)
        if values := attr.Values(); len(values) > 0 {
            fmt.Printf("  Possible values: %v\n", values)
        }
        
        // Usage guidelines
        if guidelines := attr.Guidelines(); len(guidelines) > 0 {
            fmt.Printf("  Guidelines:\n")
            for _, guideline := range guidelines {
                fmt.Printf("    - %s\n", guideline)
            }
        }
        
        // Code examples
        if examples := attr.Examples(); len(examples) > 0 {
            fmt.Printf("  Examples:\n")
            for _, example := range examples {
                fmt.Printf("    - %s\n", example)
            }
        }
        
        fmt.Println()
    }
}
```

### Generating HTML with Proper Attributes

```go
func generateHTMLWithAttributes(registry *mcp.Registry, tagName string) string {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        return fmt.Sprintf("<!-- Error: %v -->", err)
    }
    
    var html strings.Builder
    html.WriteString(fmt.Sprintf("<%s", tagName))
    
    // Add attributes with defaults
    for _, attr := range element.Attributes() {
        if defaultVal := attr.Default(); defaultVal != "" {
            // Use default value, removing quotes if present
            value := strings.Trim(defaultVal, "\"'")
            html.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), value))
        }
    }
    
    html.WriteString(fmt.Sprintf("></%s>", tagName))
    return html.String()
}

// Example usage:
// html := generateHTMLWithAttributes(registry, "my-button")
// fmt.Println(html) // <my-button variant="primary" size="medium"></my-button>
```

## Working with Slots

### Examining Slot Structure

```go
func examineSlots(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Slots for %s:\n", tagName)
    for _, slot := range element.Slots() {
        if slot.Name() == "" {
            fmt.Printf("  Default slot: %s\n", slot.Description())
        } else {
            fmt.Printf("  Named slot '%s': %s\n", slot.Name(), slot.Description())
        }
        
        // Show usage examples
        if examples := slot.Examples(); len(examples) > 0 {
            fmt.Printf("    Usage: %s\n", examples[0])
        }
        
        fmt.Println()
    }
}
```

### Generating HTML with Slots

```go
func generateHTMLWithSlots(registry *mcp.Registry, tagName string, content map[string]string) string {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        return fmt.Sprintf("<!-- Error: %v -->", err)
    }
    
    var html strings.Builder
    html.WriteString(fmt.Sprintf("<%s>\n", tagName))
    
    // Add content for each slot
    for _, slot := range element.Slots() {
        slotName := slot.Name()
        
        if slotContent, exists := content[slotName]; exists {
            if slotName == "" {
                // Default slot content
                html.WriteString(fmt.Sprintf("  %s\n", slotContent))
            } else {
                // Named slot content
                html.WriteString(fmt.Sprintf("  <span slot=\"%s\">%s</span>\n", slotName, slotContent))
            }
        }
    }
    
    html.WriteString(fmt.Sprintf("</%s>", tagName))
    return html.String()
}

// Example usage:
// content := map[string]string{
//     "": "Click me!",           // default slot
//     "icon": "🚀",             // named slot
// }
// html := generateHTMLWithSlots(registry, "my-button", content)
```

## Working with Events

### Examining Events

```go
func examineEvents(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Events for %s:\n", tagName)
    for _, event := range element.Events() {
        fmt.Printf("  %s (%s)\n", event.Name(), event.Type())
        if event.Description() != "" {
            fmt.Printf("    %s\n", event.Description())
        }
        
        // Show usage examples
        if examples := event.Examples(); len(examples) > 0 {
            fmt.Printf("    Usage: %s\n", examples[0])
        }
        
        fmt.Println()
    }
}
```

### Generating Event Listeners

```go
func generateEventListeners(registry *mcp.Registry, tagName string) []string {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        return []string{fmt.Sprintf("// Error: %v", err)}
    }
    
    var listeners []string
    for _, event := range element.Events() {
        listener := fmt.Sprintf(
            "element.addEventListener('%s', (e) => {\n  console.log('%s event:', e.detail);\n});",
            event.Name(),
            event.Name(),
        )
        listeners = append(listeners, listener)
    }
    
    return listeners
}
```

## Working with CSS

### CSS Custom Properties

```go
func examineCSSProperties(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("CSS Properties for %s:\n", tagName)
    for _, prop := range element.CssProperties() {
        fmt.Printf("  %s\n", prop.Name())
        fmt.Printf("    Syntax: %s\n", prop.Syntax())
        fmt.Printf("    Initial: %s\n", prop.Initial())
        fmt.Printf("    Inherits: %t\n", prop.Inherits())
        
        if prop.Description() != "" {
            fmt.Printf("    Description: %s\n", prop.Description())
        }
        
        fmt.Println()
    }
}
```

### Generating CSS Rules

```go
func generateCSSRules(registry *mcp.Registry, tagName string) string {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        return fmt.Sprintf("/* Error: %v */", err)
    }
    
    var css strings.Builder
    css.WriteString(fmt.Sprintf("%s {\n", tagName))
    
    // Add CSS custom properties with their initial values
    for _, prop := range element.CssProperties() {
        if initial := prop.Initial(); initial != "" {
            css.WriteString(fmt.Sprintf("  %s: %s;\n", prop.Name(), initial))
        }
    }
    
    css.WriteString("}")
    return css.String()
}
```

### CSS Parts and States

```go
func generateCSSPartsAndStates(registry *mcp.Registry, tagName string) string {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        return fmt.Sprintf("/* Error: %v */", err)
    }
    
    var css strings.Builder
    
    // CSS Parts
    for _, part := range element.CssParts() {
        css.WriteString(fmt.Sprintf("/* %s */\n", part.Description()))
        css.WriteString(fmt.Sprintf("%s::%s {\n", tagName, part.Name()))
        css.WriteString("  /* Your styles here */\n")
        css.WriteString("}\n\n")
    }
    
    // CSS States
    for _, state := range element.CssStates() {
        css.WriteString(fmt.Sprintf("/* %s */\n", state.Description()))
        css.WriteString(fmt.Sprintf("%s:%s {\n", tagName, state.Name()))
        css.WriteString("  /* Your styles here */\n")
        css.WriteString("}\n\n")
    }
    
    return css.String()
}
```

## Unified Item Processing

### Processing All Items by Kind

```go
func processAllItems(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        log.Fatal(err)
    }
    
    itemCounts := make(map[mcp.ItemKind]int)
    
    fmt.Printf("All items for %s:\n", tagName)
    for _, item := range element.Items {
        kind := item.Kind()
        itemCounts[kind]++
        
        fmt.Printf("  %s: %s", kind, item.Name())
        
        // Add type-specific information
        switch kind {
        case mcp.KindAttribute:
            if attr, ok := item.(mcp.Attribute); ok {
                fmt.Printf(" (%s)", attr.Type())
                if attr.Required() {
                    fmt.Printf(" [REQUIRED]")
                }
            }
        case mcp.KindEvent:
            if event, ok := item.(mcp.Event); ok {
                fmt.Printf(" (%s)", event.Type())
            }
        case mcp.KindCssProperty:
            if prop, ok := item.(mcp.CssProperty); ok {
                fmt.Printf(" (%s)", prop.Syntax())
            }
        }
        
        fmt.Println()
        
        if description := item.Description(); description != "" {
            fmt.Printf("    %s\n", description)
        }
    }
    
    fmt.Printf("\nSummary:\n")
    for kind, count := range itemCounts {
        fmt.Printf("  %s: %d\n", kind, count)
    }
}
```

### Generic Item Filtering

```go
func filterItemsByKind(element *mcp.ElementInfo, kind mcp.ItemKind) []mcp.Item {
    var filtered []mcp.Item
    for _, item := range element.Items {
        if item.Kind() == kind {
            filtered = append(filtered, item)
        }
    }
    return filtered
}

func findItemByName(element *mcp.ElementInfo, name string) mcp.Item {
    for _, item := range element.Items {
        if item.Name() == name {
            return item
        }
    }
    return nil
}
```

## JSON API Integration

### Creating JSON API Endpoints

```go
import (
    "encoding/json"
    "net/http"
)

func createElementAPI(registry *mcp.Registry) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tagName := r.URL.Query().Get("tag")
        if tagName == "" {
            http.Error(w, "tag parameter required", http.StatusBadRequest)
            return
        }
        
        element, err := registry.GetElementInfo(tagName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(element); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func createElementListAPI(registry *mcp.Registry) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        elements := registry.GetAllElements()
        
        // Create summary response
        summary := make(map[string]interface{})
        for tagName, element := range elements {
            summary[tagName] = map[string]interface{}{
                "description": element.Description,
                "attributes":  len(element.Attributes()),
                "slots":       len(element.Slots()),
                "events":      len(element.Events()),
                "cssProps":    len(element.CssProperties()),
            }
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(summary)
    }
}
```

### Working with JSON Data

```go
func parseElementFromJSON(jsonData []byte) (*mcp.ElementInfo, error) {
    var element mcp.ElementInfo
    if err := json.Unmarshal(jsonData, &element); err != nil {
        return nil, err
    }
    return &element, nil
}

func extractSpecificItems(jsonData []byte, itemKind mcp.ItemKind) ([]map[string]interface{}, error) {
    var element mcp.ElementInfo
    if err := json.Unmarshal(jsonData, &element); err != nil {
        return nil, err
    }
    
    var items []map[string]interface{}
    for _, item := range element.Items {
        if item.Kind() == itemKind {
            // Convert to map for JSON manipulation
            itemData, _ := json.Marshal(item)
            var itemMap map[string]interface{}
            json.Unmarshal(itemData, &itemMap)
            items = append(items, itemMap)
        }
    }
    
    return items, nil
}
```

## Advanced Patterns

### Element Comparison

```go
func compareElements(registry *mcp.Registry, tagName1, tagName2 string) {
    element1, err1 := registry.GetElementInfo(tagName1)
    element2, err2 := registry.GetElementInfo(tagName2)
    
    if err1 != nil || err2 != nil {
        fmt.Printf("Error comparing elements: %v, %v\n", err1, err2)
        return
    }
    
    // Compare attributes
    attrs1 := make(map[string]mcp.Attribute)
    for _, attr := range element1.Attributes() {
        attrs1[attr.Name()] = attr
    }
    
    attrs2 := make(map[string]mcp.Attribute)
    for _, attr := range element2.Attributes() {
        attrs2[attr.Name()] = attr
    }
    
    fmt.Printf("Comparing %s vs %s:\n", tagName1, tagName2)
    
    // Common attributes
    fmt.Println("Common attributes:")
    for name := range attrs1 {
        if attr2, exists := attrs2[name]; exists {
            attr1 := attrs1[name]
            fmt.Printf("  %s: %s vs %s\n", name, attr1.Type(), attr2.Type())
        }
    }
    
    // Unique to element1
    fmt.Printf("Unique to %s:\n", tagName1)
    for name := range attrs1 {
        if _, exists := attrs2[name]; !exists {
            fmt.Printf("  %s (%s)\n", name, attrs1[name].Type())
        }
    }
    
    // Unique to element2
    fmt.Printf("Unique to %s:\n", tagName2)
    for name := range attrs2 {
        if _, exists := attrs1[name]; !exists {
            fmt.Printf("  %s (%s)\n", name, attrs2[name].Type())
        }
    }
}
```

### Validation Helpers

```go
func validateAttributeValue(attr mcp.Attribute, value string) error {
    // Check if attribute has constrained values
    if values := attr.Values(); len(values) > 0 {
        for _, validValue := range values {
            if value == validValue {
                return nil
            }
        }
        return fmt.Errorf("invalid value '%s' for attribute '%s', valid values: %v", 
            value, attr.Name(), values)
    }
    
    // Additional type-based validation could be added here
    return nil
}

func validateHTML(registry *mcp.Registry, html string) []error {
    // This is a simplified example - real implementation would parse HTML
    var errors []error
    
    // Extract tag names and attributes from HTML (simplified)
    // In practice, you'd use an HTML parser like golang.org/x/net/html
    
    return errors
}
```

## Error Handling

### Comprehensive Error Handling

```go
func robustElementAccess(registry *mcp.Registry, tagName string) {
    element, err := registry.GetElementInfo(tagName)
    if err != nil {
        switch {
        case strings.Contains(err.Error(), "not found"):
            fmt.Printf("Element '%s' is not available in the current workspace\n", tagName)
            
            // Suggest similar elements
            allElements := registry.GetAllElements()
            fmt.Println("Available elements:")
            for name := range allElements {
                fmt.Printf("  - %s\n", name)
            }
            
        default:
            fmt.Printf("Unexpected error accessing element '%s': %v\n", tagName, err)
        }
        return
    }
    
    // Safe access with nil checks
    if len(element.Attributes()) > 0 {
        fmt.Printf("Found %d attributes\n", len(element.Attributes()))
    } else {
        fmt.Println("No attributes defined")
    }
    
    if len(element.Slots()) > 0 {
        fmt.Printf("Found %d slots\n", len(element.Slots()))
    } else {
        fmt.Println("No slots defined")
    }
}
```

### Registry Health Checks

```go
func checkRegistryHealth(registry *mcp.Registry) error {
    // Check if registry can load manifests
    if err := registry.LoadManifests(); err != nil {
        return fmt.Errorf("failed to load manifests: %w", err)
    }
    
    // Check if any elements are available
    elements := registry.GetAllElements()
    if len(elements) == 0 {
        return fmt.Errorf("no elements found in registry")
    }
    
    // Validate a sample element
    for tagName := range elements {
        element, err := registry.GetElementInfo(tagName)
        if err != nil {
            return fmt.Errorf("failed to get info for element '%s': %w", tagName, err)
        }
        
        if element.TagName != tagName {
            return fmt.Errorf("tag name mismatch for element '%s'", tagName)
        }
        
        // Only check the first element
        break
    }
    
    return nil
}
```

These examples demonstrate the flexibility and power of the interface-based MCP registry design. The unified interface hierarchy enables type-safe access to all component information while maintaining the ability to work with items generically when needed.