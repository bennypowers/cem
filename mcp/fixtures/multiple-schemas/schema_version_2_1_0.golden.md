# Component Context for Schema 2.1.0 Test

Test context for schema version 2.1.0 from fixture file with 1 elements.


*Using  (2.1.0)*


The top-level interface of a custom elements manifest file.

Because custom elements are JavaScript classes, describing a custom element
may require describing arbitrary JavaScript concepts like modules, classes,
functions, etc. So custom elements manifests are capable of documenting
the elements in a package, as well as those JavaScript concepts.

The modules described in a package should be the public entrypoints that
other packages may import from. Multiple modules may export the same object
via re-exports, but in most cases a package should document the single
canonical export that should be used.



## Custom Elements in Your Project
Custom elements are JavaScript classes that extend HTML with new functionality. A description of a custom element class.

Custom elements are JavaScript classes, so this extends from
`ClassDeclaration` and adds custom-element-specific features like
attributes, events, and slots.

Note that `tagName` in this interface is optional. Tag names are not
neccessarily part of a custom element class, but belong to the definition
(often called the &#34;registration&#34;) or the `customElements.define()` call.

Because classes and tag names can only be registered once, there&#39;s a
one-to-one relationship between classes and tag names. For ease of use,
we allow the tag name here.

Some packages define and register custom elements in separate modules. In
these cases one `Module` should contain the `CustomElement` without a
tagName, and another `Module` should contain the
`CustomElementExport`.
*No custom elements found in your manifests.*

## Schema Reference

## How to Use This Context

Each element above shows its complete API surface. When working with custom elements:
- Use the exact tag names and attribute names shown
- Respect the slot structure for content placement
- Use CSS custom properties for theming and styling
- Follow the documented guidelines and constraints
- Reference the schema properties for understanding data types and requirements