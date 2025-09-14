# Component Context from Your Manifests

Your workspace contains **2 custom elements** with real constraints and patterns.

## Your Component Library Statistics
- **2 elements** across **0 common prefixes** ()
- **3 CSS custom properties** defined
- **Schema versions**: 2.1.1

This context is derived from your actual manifest data, not generic guidelines.


*Using Custom Elements Manifest Schema 2.1.1-speculative (2.1.1)*


This is a speculative schema for 2.1.1 based on 2.1.0, with workarounds for known issues. See: https://github.com/webcomponents/custom-elements-manifest/issues/138 and https://github.com/vega/ts-json-schema-generator/pull/2323



## Custom Elements in Your Project
Custom elements are JavaScript classes that extend HTML with new functionality. A description of a custom element class.

Custom elements are JavaScript classes, so this extends from `ClassDeclaration` and adds custom-element-specific features like attributes, events, and slots.

Note that `tagName` in this interface is optional. Tag names are not neccessarily part of a custom element class, but belong to the definition (often called the &#34;registration&#34;) or the `customElements.define()` call.

Because classes and tag names can only be registered once, there&#39;s a one-to-one relationship between classes and tag names. For ease of use, we allow the tag name here.

Some packages define and register custom elements in separate modules. In these cases one `Module` should contain the `CustomElement` without a tagName, and another `Module` should contain the `CustomElementExport`.

### `card-element`
**Attributes:**
- `elevation` (number) - Card elevation level
**Slots:**
- `default` - Card content
- `header` - Card header
- `footer` - Card footer
**CSS Custom Properties:**
- `--card-background` (&lt;color&gt;) - Card background color
**CSS Parts:**
- `container` - Card container
- `header` - Card header part
- `content` - Card content part
- `footer` - Card footer part

### `button-element`
**Attributes:**
- `variant` (&#34;primary&#34; | &#34;secondary&#34; | &#34;ghost&#34;) - Button variant
- `size` (&#34;small&#34; | &#34;medium&#34; | &#34;large&#34;) - Button size
- `disabled` (boolean) - Whether button is disabled
**Slots:**
- `default` - Button content
- `icon` - Button icon
**CSS Custom Properties:**
- `--button-color` (&lt;color&gt;) - Button color
- `--button-padding` (&lt;length&gt;) - Button padding
**CSS Parts:**
- `button` - The button element
**Events:**
- `button-click` - Button click event

## Schema Reference

## How to Use This Context

Each element above shows its complete API surface. When working with custom elements:
- Use the exact tag names and attribute names shown
- Respect the slot structure for content placement
- Use CSS custom properties for theming and styling
- Follow the documented guidelines and constraints
- Reference the schema properties for understanding data types and requirements