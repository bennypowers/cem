# Component Context for Schema 2.1.1-speculative Test

Test context for schema version 2.1.1-speculative from fixture file with 1 elements.


*Using Custom Elements Manifest Schema 2.1.1-speculative (2.1.1-speculative)*


This is a speculative schema for 2.1.1 based on 2.1.0, with workarounds for known issues. See: https://github.com/webcomponents/custom-elements-manifest/issues/138 and https://github.com/vega/ts-json-schema-generator/pull/2323



## Custom Elements in Your Project
Custom elements are JavaScript classes that extend HTML with new functionality. A description of a custom element class.

Custom elements are JavaScript classes, so this extends from `ClassDeclaration` and adds custom-element-specific features like attributes, events, and slots.

Note that `tagName` in this interface is optional. Tag names are not necessarily part of a custom element class, but belong to the definition (often called the "registration") or the `customElements.define()` call.

Because classes and tag names can only be registered once, there's a one-to-one relationship between classes and tag names. For ease of use, we allow the tag name here.

Some packages define and register custom elements in separate modules. In these cases one `Module` should contain the `CustomElement` without a tagName, and another `Module` should contain the `CustomElementExport`.
*No custom elements found in your manifests.*

## Schema Reference

## How to Use This Context

Each element above shows its complete API surface. When working with custom elements:
- Use the exact tag names and attribute names shown
- Respect the slot structure for content placement
- Use CSS custom properties for theming and styling
- Follow the documented guidelines and constraints
- Reference the schema properties for understanding data types and requirements