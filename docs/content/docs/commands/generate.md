---
title: Generate
layout: page
---

- **Generates CEM files** from source code using syntax analysis powered by
[go][go] and [tree-sitter][treesitter].
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic style using Lit and TypeScript, with a
`@customElement` decorator, and `@property` decorators on class fields.

{{<tip "warning">}}

`cem generate` best supports LitElements written in idiomatic style with
TypeScript decorators. There is rudimentary support for `extends HTMLElement`,
but it is not a high priority for development. If you need something more
specific [open an issue](https://github.com/bennypowers/cem/issues/new).

{{</tip>}}

## JSDoc
Use JSDoc comments to add metadata to your element classes, similar to other
tools. Add a description by separating the name of the item with ` - `

- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts. See [Slots and Parts](#slots-and-parts)
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@demo` — Demo URL
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@slot` — Named or default slots. See [Slots and Parts](#slots-and-parts)
- `@summary` — Short summary for documentation

See the [generate test fixtures](https://github.com/bennypowers/cem/tree/main/generate/test/fixtures/) directory for examples

<a id="html-template-analysis-for-slots-and-parts"></a>
## Slots and Parts

`cem` automatically detects `<slot>` elements and `part` attributes in your element’s `render()` template, merging them with any information provided via JSDoc. You can also document slots and parts inline in your template HTML using HTML comments. This helps in documenting them for your users, but is not required to detect them.

- If the comment is a plain string, it will be used as the `description` for the element. In cases where an element is both a slot and a part, the description will only be applied to the slot.
- For more detailed documentation, you can use YAML inside the comment.
- Markdown is supported in all HTML comment documentation. Use a backslash (`\`) to escape backticks.

### Examples

#### Plain String Comment
A simple comment will be treated as the description. Markdown content is supported and encouraged.
```html
<!-- This is the **default** `slot`. -->
<slot></slot>
```

#### YAML Comment
For more complex metadata, use YAML syntax.
```html
<!--
  summary: The main slot for content
  description: |
    This slot displays user-provided content.
    Supports multiline **markdown**.
    e.g. `code`
  deprecated: true
-->
<slot></slot>
```

#### Documenting Slots and Parts together
When an element is both a slot and a part, you can document both in a single comment.
```html
<!-- slot:
       summary: The `info` slot
     part:
       summary: The `info-part` part
-->
<slot name="info" part="info-part"></slot>
```

{{<tip "warning">}}
When including inline markdown <code>\`code\`</code> in your comments in lit-html templates, 
you will need to escape the backticks in the comment.
{{</tip>}}

## CSS Custom Properties
Supports CSS Custom Properties by scanning css files and css tagged-template-literals. `cem` also discovers properties defined in `:host` rules.

- Custom properties beginning with `_` will be ignored (treated as "private")
e.g. `var(--_private)`
- If you provide a [Design Tokens Community Group][dtcg] format module (JSON) to
`cem` via the `--design-tokens` flag,
`cem` will add metadata from your design system to any matching css variables it
finds in your elements
- You can use jsdoc-like comment syntax before each var call to document your
variables
- When both user comments and design token descriptions exist for the same CSS 
custom property, `cem` uses the user's variable description first, then the 
design token's description, separated by two new lines.

### Example

```css
:host {
  /**
   * A property defined on the host
   * @summary The host's custom property
   */
  --host-property: red;

  color:
    /**
     * custom color for use in this element
     * @summary color
     * @deprecated just use the `color` property
     */
    var(--custom-color);
  border:
    1px
    solid
    /** Border color of the element */
    var(--border-color);
}
```

When you have a declaration where the LHS and RHS both contain CSS custom properties, 
you have to position your comments so that they target the correct variable:

Good:

```css
/** comment for --a */
color: var(--a);
/** comment for --b */
--b: blue;
/** comment for --c */
--c:
  /** comment for --d */
  var(--d);
```

Bad:
```css
/** comment for --d */
--c: var(--d);
```

### JSON Output Format

The generated manifest uses Go's JSON marshaling, which escapes HTML-sensitive characters for security reasons. Characters like `<` and `>` are converted to `\u003c` and `\u003e` respectively to prevent XSS attacks when JSON is embedded in HTML script tags.

This means CSS syntax values in your manifest will appear escaped:

```json
{
  "name": "--primary-color",
  "syntax": "\u003ccolor\u003e"
}
```

This is the expected behavior and doesn't affect the functionality of the manifest. Tools consuming the manifest should handle these standard JSON escape sequences correctly.

---

<a id="element-demos"></a>
## Demos

`cem generate` supports documenting your elements' demos by linking directly
from JSDoc, or by configurable file-system based discovery.

### JSDoc `@demo` Tag

Add demos directly to your element class or members with the `@demo` tag:

```ts
/**
 * @demo https://example.com/my-element-plain/
 * @demo https://example.com/my-element-fancy/ - A fancier demo with description
 */
@customElement('my-element')
class MyElement extends LitElement {
  // ...
}
```

Demos defined this way will always appear in your manifest for the element.

### Automatic Demo Discovery

`cem` can automatically discover demos from your codebase based on your
repository structure and configuration. Demos that are co-located with their
component's source module will be prioritized in the generated manifest.

<a id="demo-discovery-options"></a>
## Demo Discovery

The `urlPattern` uses the standard URLPattern syntax with named parameters (`:paramName`).
You can use it to match file paths and extract parameters to build your demo URLs.

For example, if your demos are in subdirectories like `src/my-element/demos/foo.html`,
you could use a pattern like this:

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  demoDiscovery:
    fileGlob: "src/**/demos/*.html"
    urlPattern: "/src/:tag/demos/:demo.html"
    urlTemplate: "https://example.com/elements/{{.tag | alias}}/{{.demo | slug}}/"
```

### URL Template Functions

The `urlTemplate` supports Go template syntax with built-in functions:

- `alias` - Apply element alias mapping
- `slug` - Convert to URL-friendly format  
- `lower` - Convert to lowercase
- `upper` - Convert to uppercase

Use pipeline syntax (`{{.param | function}}`) or function calls (`{{function .param}}`).

**Demo discovery options:**

| Option                 | Type   | Description                                                                                  |
| ---------------------- | ------ | -------------------------------------------------------------------------------------------- |
| `fileGlob`             | string | Glob pattern for discovering demo files.                                                     |
| `sourceControlRootUrl` | string | Canonical public source control URL for your repository root (on the main branch).           |
| `urlPattern`           | string | URLPattern with named parameters (`:param`) for matching demo file paths.                   |
| `urlTemplate`          | string | Go template with functions for generating canonical demo URLs.                               |

## Monorepos

If you are planning to use `cem` in an npm or yarn monorepo, the best way for now
is to create a new `.config/cem.yaml` file for each package you want to generate 
for, instead of using a top-level config file.

### Example

Root package.json:
```json
{
  "scripts": {
    "generate": "npm run generate --workspaces"
  },
  "workspaces": [
    "./core",
    "./elements"
  ]
}
```

core/.config/cem.yaml:
```yaml
generate:
  files:
    - './**/*.ts'
```

core/package.json
```json
{
  "scripts": {
    "generate": "cem generate"
  }
}
```

elements/.config/cem.yaml:
```yaml
generate:
  files:
    - './**/*.ts'
```

elements/package.json
```json
{
  "scripts": {
    "generate": "cem generate"
  }
}
```


## Usage

Generate a custom elements manifest from your files:

```sh
cem generate \
  "src/**/*.ts" \
  --design-tokens npm:@my-ds/tokens/tokens.json \
  --exclude "src/**/*.test.ts" \
  --output custom-elements.json
```

For npm projects you can use `npx @pwrs/cem generate ...`.

{{<tip "note">}}
The `generate` command does not support remote packages. To inspect a remote package's manifest, use the `cem list` command.
{{</tip>}}

<a id="command-line-arguments"></a>
### Arguments

| Argument                        | Type               | Description                                                                                       |
| ------------------------------- | ------------------ | ------------------------------------------------------------------------------------------------- |
| `<files or globs>`              | positional (array) | Files or glob patterns to include                                                                 |
| `--package, -p`                 | string             | Path to a package directory.                                                                      |
| `--output, -o`                  | string             | Write the manifest to this file instead of stdout                                                 |
| `--watch, -w`                   | bool               | Watch files for changes and regenerate automatically                                              |
| `--exclude, -e`                 | array              | Files or glob patterns to exclude                                                                 |
| `--no-default-excludes`         | bool               | Do not exclude files by default (e.g., `.d.ts` files will be included unless excluded explicitly) |
| `--design-tokens`               | string             | Path or npm specifier for DTCG-format design tokens                                               |
| `--design-tokens-prefix`        | string             | CSS custom property prefix for design tokens                                                      |
| `--demo-discovery-file-glob`    | string             | Glob pattern for discovering demo files                                                           |
| `--demo-discovery-url-pattern`  | string             | URLPattern with named parameters for matching demo file paths                                     |
| `--demo-discovery-url-template` | string             | Go template with functions for generating canonical demo URLs                                      |
| `--source-control-root-url`     | string             | Glob pattern for discovering demo files                                                           |
| `--project-dir`                 | string             | **Deprecated:** Use `--package` instead.                                                          |

By default, some files (like `.d.ts` TypeScript declaration files) are excluded from the manifest.
Use `--no-default-excludes` if you want to include all matching files and manage excludes yourself.

[cem]: https://github.com/webcomponents/custom-elements-manifest
[dtcg]: https://tr.designtokens.org/format/
[go]: https://go.dev
[treesitter]: https://tree-sitter.github.io/tree-sitter/
[gpl3]: https://www.gnu.org/licenses/gpl-3.0.html
[contributingmd]: ./CONTRIBUTING.md
[issuenew]: https://github.com/bennypowers/cem/issues/new
[installationdocs]: https://bennypowers.github.io/cem/installation/
[generatedocs]: https://bennypowers.github.io/cem/commands/generate/
