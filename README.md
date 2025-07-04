# cem

**cem** is a command-line tool for generating and querying
[Custom Elements Manifest][cem] files. It can analyze your codebase and generate
rich metadata for your custom elements, facilitating documentation, tooling, and
integration. It can also query that manifest for information about your package

## Installation

```sh
npm install --save-dev @pwrs/cem
```

For more options, see [Installation docs][installationdocs]

## Features

### `cem generate`

- **Generates CEM files** from source code using syntax analysis powered by
[go][go] and [tree-sitter][treesitter].
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic Lit typescript style, with a
`@customElement` decorator, and `@property` decorators on class fields.

> [!NOTE]
> `cem generate` best supports LitElements written in idiomatic style with
> TypeScript decorators. There is rudimentary support for `extends HTMLElement`,
> but it is not a high priority for development. If you need something more
> specific [open an issue][issuenew].

#### JSDoc
Use JSDoc comments to add metadata to your element classes, similar to other
tools. Add a description by separating the name of the item with ` - `

- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@demo` — Demo URL
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@slot` — Named or default slots
- `@summary` — Short summary for documentation

See the [test-fixtures](tree/main/generate/test-fixtures/) directory for examples

#### HTML Template Analysis for Slots and Parts

- **Automatically detects `<slot>` elements and `part` attributes in your element’s
`render()` template.**
- Merges slot and part information found in templates with any provided
via JSDoc, ensuring comprehensive documentation in the generated manifest.
- **Deprecation and other metadata** for slots and parts can be specified via
YAML in HTML comments.
- **Supports documenting slots and parts inline in your template HTML** using
HTML comments with YAML blocks.
- YAML comments are not necessary to detect slots and parts, but help in
documenting them for your users.

##### Examples
```html
<!--
  summary: The main slot for content
  description: |
    This slot displays user-provided content.
    Supports multiline **markdown**.
  deprecated: true
-->
<slot></slot>
<!-- slot:
       summary: Named slot summary
     part:
       summary: Part summary
-->
<slot name="info" part="info-part"></slot>
```

#### CSS Custom Properties
Supports CSS Custom Properties by scanning css files and css tagged-template-literals

- Custom properties beginning with `_` will be ignored (treated as "private")
e.g. `var(--_private)`
- If you provide a [Design Tokens Community Group][dtcg] format module (JSON) to
`cem` via the `--design-tokens` flag,
`cem` will add metadata from your design system to any matching css variables it
finds in your elements
- You can use jsdoc-like comment syntax before each var call to document your
variables

##### Example

```css
:host {
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

---

#### Element Demos

`cem generate` supports documenting your elements' demos by linking directly
from JSDoc, or by configurable file-system based discovery.

##### 1. JSDoc `@demo` Tag

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

##### 2. Demo Discovery

`cem` can automatically discover demos from your codebase based on your
repository structure and configuration.

#### Demo Discovery Options

Configure demo discovery with the `demoDiscovery` key in your `.config/cem.yaml` file

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  demoDiscovery:
    fileGlob: "demos/**/*.html"
    urlPattern: "demos/(?P<tag>[\w-]+)/(?P<demo>[\w-]+).html"
    urlTemplate: "https://example.com/elements/{tag}/{demo}/"
```

**Demo discovery options:**

| Option                 | Type   | Description                                                                                  |
| ---------------------- | ------ | -------------------------------------------------------------------------------------------- |
| `fileGlob`             | string | Glob pattern for discovering demo files.                                                     |
| `sourceControlRootUrl` | string | Canonical public source control URL for your repository root (on the main branch).           |
| `urlPattern`           | string | Pattern for generating demo URLs, e.g. `"demos/{tag}.html"`. `{tag}` is replaced by tag name.|
| `urlTemplate`          | string | (optional) Alternative URL template for demo links.                                          |

#### Usage

Generate a custom elements manifest from your files:

```sh
cem generate \
  "src/**/*.ts" \
  --design-tokens npm:@my-ds/tokens/tokens.json \
  --exclude "src/**/*.test.ts" \
  --output custom-elements.json
```

For npm projects you can use `npx @pwrs/cem generate ...`.

##### Command Line Arguments

| Argument                        | Type               | Description                                                                                       |
| ------------------------------- | ------------------ | ------------------------------------------------------------------------------------------------- |
| `<files or globs>`              | positional (array) | Files or glob patterns to include                                                                 |
| `--output, -o`                  | string             | Write the manifest to this file instead of stdout                                                 |
| `--exclude, -e`                 | array              | Files or glob patterns to exclude                                                                 |
| `--no-default-excludes`         | bool               | Do not exclude files by default (e.g., `.d.ts` files will be included unless excluded explicitly) |
| `--design-tokens, -t`           | string             | Path or npm specifier for DTCG-format design tokens                                               |
| `--design-tokens-prefix, -p`    | string             | CSS custom property prefix for design tokens                                                      |
| `--demo-discovery-file-glob`    | string             | Glob pattern for discovering demo files                                                           |
| `--demo-discovery-url-pattern`  | string             | Go Regexp pattern with named capture groups for generating canonical demo urls                    |
| `--demo-discovery-url-template` | string             | URL pattern string using {groupName} syntax to interpolate named captures from the URL pattern    |
| `--source-control-root-url`     | string             | Glob pattern for discovering demo files                                                           |
| `--project-dir`                 | string             | Specify the project root directory to use for resolving relative paths and configuration.         |

By default, some files (like `.d.ts` TypeScript declaration files) are excluded from the manifest.
Use `--no-default-excludes` if you want to include all matching files and manage excludes yourself.

---

### `cem list`

The `cem list` command provides a fast, flexible way to inspect custom elements, their features, and their metadata directly from your manifest file.
With `cem list`, you can quickly explore and audit your custom elements API surface, making it easier to document, test, and share your components.

#### Available Subcommands

- `cem list tags` — Lists all custom element tag names in the project.
- `cem list modules` — Lists all module paths in the project.
- `cem list -t <tag> attrs` — Lists all attributes for a given custom element tag.
- `cem list -t <tag> slots` — Lists all named and default slots for a tag.
- `cem list -t <tag> events` — Lists all custom events fired by a tag, including their types and descriptions.
- `cem list -t <tag> css-properties` — Lists CSS custom properties (CSS variables) for a tag.
- `cem list -t <tag> css-states` — Lists CSS custom states for a tag.
- `cem list -t <tag> css-parts` — Lists CSS shadow parts for a tag.
- `cem list -t <tag> methods` — Lists methods for a tag's DOM object.

#### Column Filtering and Output

- Use the `--columns,-c` flag to specify which columns to include in the output, e.g.:

  ```sh
  cem list events my-element -c name -c summary -c type
  cem list attrs my-element -c description --columns default
  ```
Note that the name column is always included, and that if a column is specified but contains only empty values for all rows, it is automatically omitted from the output for clarity.

#### Output Formats

- By default, tables are shown in a human-readable table format.
- [ ] TODO: json, markdown, flat lists, etc

## Configuration Reference

You can configure CEM via `.config/cem.yaml`, relative to your project root,
or via CLI flags.

### Example Configuration

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  files:
    - "src/**/*.ts"
  exclude:
    - "src/**/*.test.ts"
  output: "custom-elements.json"
  noDefaultExcludes: false
  designTokens:
    spec: "npm:@my-ds/tokens/tokens.json"
    prefix: "--my-ds"
  demoDiscovery:
    fileGlob: "demos/**/*.html"
    urlPattern: "demos/(?P<tag>[\w-]+)/(?P<demo>[\w-]+).html"
    urlTemplate: "https://example.com/elements/{tag}/{demo}/"
```

---

## Contributing

For information on building and testing, please see
[CONTRIBUTING.md][contributingmd].

## License

This program is free software: you can redistribute it and/or modify it under
the terms of the [GNU General Public License v3.0][gpl3].

&copy; 2025 Benny Powers <web@bennypowers.com>

[cem]: https://github.com/webcomponents/custom-elements-manifest
[dtcg]: https://tr.designtokens.org/format/
[go]: https://go.dev
[treesitter]: https://tree-sitter.github.io/tree-sitter/
[gpl3]: https://www.gnu.org/licenses/gpl-3.0.html
[contributingmd]: ./CONTRIBUTING.md
[issuenew]: https://github.com/bennypowers/cem/issues/new
[installationdocs]: https://bennypowers.github.io/cem/installation/
