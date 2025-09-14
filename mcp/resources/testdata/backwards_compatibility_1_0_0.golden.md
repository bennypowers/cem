# Component Context for Backwards Compatibility Test

Testing schema 1.0.0 with minimal data.

## Schema Context:  (1.0.0)


The top-level interface of a custom elements manifest file.

Because custom elements are JavaScript classes, describing a custom element
may require describing arbitrary JavaScript concepts like modules, classes,
functions, etc. So custom elements manifests are capable of documenting
the elements in a package, as well as those JavaScript concepts.

The modules described in a package should be the public entrypoints that
other packages may import from. Multiple modules may export the same object
via re-exports, but in most cases a package should document the single
canonical export that should be used.

### Schema Field Definitions




#### modules
An array of the modules this package contains.





#### CustomElementDeclaration
A description of a custom element class.

Custom elements are JavaScript classes, so this extends from
`ClassDeclaration` and adds custom-element-specific features like
attributes, events, and slots.

Note that `tagName` in this interface is optional. Tag names are not
neccessarily part of a custom element class, but belong to the definition
(often called the &#34;registration&#34;) or the `customElements.define()` call.

Because classes and tag anmes can only be registered once, there&#39;s a
one-to-one relationship between classes and tag names. For ease of use,
we allow the tag name here.

Some packages define and register custom elements in separate modules. In
these cases one `Module` should contain the `CustomElement` without a
tagName, and another `Module` should contain the
`CustomElement`.

Core properties:


























This schema provides the semantic framework for understanding your specific component data below.


## Your Component Data















## How to Use This Context

This information helps AI understand:
- **Your naming conventions** ( prefixes)
- **Your component patterns** (common attributes and slots)
- **Your CSS architecture** (custom properties and design tokens)
- **Your documented constraints** (guidelines from descriptions)

When asking for component help, the AI can now reference your actual manifest data and the schema definitions that explain what each field means.