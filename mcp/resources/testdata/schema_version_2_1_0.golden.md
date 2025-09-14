# Component Context for Schema 2.1.0 Test

Test context for schema version 2.1.0.

## Schema Context:  (2.1.0)


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

Because classes and tag names can only be registered once, there&#39;s a
one-to-one relationship between classes and tag names. For ease of use,
we allow the tag name here.

Some packages define and register custom elements in separate modules. In
these cases one `Module` should contain the `CustomElement` without a
tagName, and another `Module` should contain the
`CustomElementExport`.

Core properties:

- **tagName**: An optional tag name that should be specified if this is a
self-registering element.

Self-registering elements must also include a CustomElementExport
in the module&#39;s exports.

- **attributes**: The attributes that this element is known to understand.

- **slots**: The shadow dom content slots that this element accepts.

- **events**: The events that this element fires.











#### Slot
Content slots define where child content can be placed within your custom element's shadow DOM. Slots enable flexible, composable component designs.

Key properties:

- **name**: The slot name, or the empty string for an unnamed slot.



#### CssCustomProperty
CSS custom properties (CSS variables) provide a styling API for your custom elements. They allow theme customization and design system integration.

Properties:

- **name**: The name of the property, including leading `--`.

- **syntax**: The expected syntax of the defined property. Defaults to &#34;*&#34;.

The syntax must be a valid CSS [syntax string](https://developer.mozilla.org/en-US/docs/Web/CSS/@property/syntax)
as defined in the CSS Properties and Values API.

Examples:

&#34;&lt;color&gt;&#34;: accepts a color
&#34;&lt;length&gt; | &lt;percentage&gt;&#34;: accepts lengths or percentages but not calc expressions with a combination of the two
&#34;small | medium | large&#34;: accepts one of these values set as custom idents.
&#34;*&#34;: any valid token










This schema provides the semantic framework for understanding your specific component data below.


## Your Component Data





### Common Slots

- `default` - Used in 1 elements



### Your CSS Custom Properties

Your components define **1 CSS custom properties**:


- `--test-color`









## How to Use This Context

This information helps AI understand:
- **Your naming conventions** ( prefixes)
- **Your component patterns** (common attributes and slots)
- **Your CSS architecture** (custom properties and design tokens)
- **Your documented constraints** (guidelines from descriptions)

When asking for component help, the AI can now reference your actual manifest data and the schema definitions that explain what each field means.