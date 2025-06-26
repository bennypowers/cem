# cem

**cem** is a command-line tool for generating [Custom Elements Manifest][cem]
(CEM) files from TypeScript sources. It analyzes your codebase and generates
rich metadata for your custom elements, facilitating documentation, tooling, and
integration.

> [!NOTE]
> `cem` best supports LitElements written in idiomatic style with
> TypeScript decorators. There is rudimentary support for `extends HTMLElement`,
> but it is not a high priority for development. If you need something more
> specific [open an issue][issuenew].

## Features

### `cem generate`

- **Generates CEM files** from source code using syntax analysis powered by
[go][go] and [tree-sitter][treesitter].
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic Lit typescript style, with a
`@customElement` decorator, and `@property` decorators on class fields.

#### JSDoc
Use JSDoc comments to add metadata to your element classes, similar to other
tools.

- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@slot` — Named or default slots
- `@summary` — Short summary for documentation

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

## Element Demos

`cem generate` supports documenting your elements' demos by linking directly
from JSDoc, or by configurable file-system based discovery.

### 1. JSDoc `@demo` Tag

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

### 2. Demo Discovery

`cem` can automatically discover demos from your codebase based on your
repository structure and configuration.

#### Demo Discovery Options

Configure demo discovery with the `demoDiscovery` key in your `.config/cem.yaml` file

```yaml
generate:
  demoDiscovery:
    filePattern: "demos/**/*.html"
    sourceControlUrl: "https://github.com/your/repo/tree/main/"
    urlPattern: "demos/{tag}.html"
    urlTemplate: ""
```

**Demo discovery options:**

| Option              | Type     | Description                                                                                 |
|---------------------|----------|---------------------------------------------------------------------------------------------|
| `filePattern`       | string   | Glob pattern for discovering demo files.                                                     |
| `sourceControlUrl`  | string   | Canonical public source control URL for your repository root (on the main branch).           |
| `urlPattern`        | string   | Pattern for generating demo URLs, e.g. `"demos/{tag}.html"`. `{tag}` is replaced by tag name.|
| `urlTemplate`       | string   | (optional) Alternative URL template for demo links.                                         |

---

## Configuration Reference

You can configure CEM via `.config/cem.yaml`, relative to your project root,
or via CLI flags.

### Example Configuration

```yaml
generate:
  files:
    - "src/**/*.ts"
  exclude:
    - "src/**/*.test.ts"
  designTokensSpec: "npm:@my-ds/tokens/tokens.json"
  designTokensPrefix: "--my-ds"
  noDefaultExcludes: false
  output: "custom-elements.json"
  demoDiscovery:
    filePattern: "demos/**/*.html"
    sourceControlUrl: "https://github.com/your/repo/tree/main/"
    urlPattern: "demos/{tag}.html"
    urlTemplate: ""
tagPrefix: ""
```

### All Available Flags

| Option                     | Type     | Description                                                                                       |
|----------------------------|----------|---------------------------------------------------------------------------------------------------|
| `files`                    | list     | Files or glob patterns to include.                                                                |
| `exclude`                  | list     | Files or glob patterns to exclude.                                                                |
| `designTokensSpec`         | string   | Path or npm specifier for DTCG-format design tokens.                                              |
| `designTokensPrefix`       | string   | CSS custom property prefix for design tokens.                                                     |
| `noDefaultExcludes`        | bool     | Do not exclude files by default (e.g., `.d.ts` files will be included unless excluded explicitly).|
| `output`                   | string   | Write the manifest to this file instead of stdout.                                                |
| `demoDiscovery`            | object   | See above for all demo discovery options.                                                         |
| `tagPrefix`                | string   | (optional) Prefix all custom element tag names.                                                   |

---

## Installation

For go binaries:
```sh
go install bennypowers.dev/cem@latest
```

For NPM projects:

```sh
npm install --save-dev @pwrs/cem
```

Or clone this repository and build from source:

```sh
git clone https://github.com/bennypowers/cem.git
cd cem
make
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
### Arguments
| Argument                     | Type               | Description                                                                                       |
| ---------------------------- | ------------------ | ------------------------------------------------------------------------------------------------- |
| `<files or globs>`           | positional (array) | Files or glob patterns to include                                                                 |
| `--exclude, -e`              | array              | Files or glob patterns to exclude                                                                 |
| `--design-tokens, -t`        | string             | Path or npm specifier for DTCG-format design tokens                                               |
| `--design-tokens-prefix, -p` | string             | CSS custom property prefix for design tokens                                                      |
| `--output, -o`               | string             | Write the manifest to this file instead of stdout                                                 |
| `--no-default-excludes`      | bool               | Do not exclude files by default (e.g., `.d.ts` files will be included unless excluded explicitly) |

By default, some files (like `.d.ts` TypeScript declaration files) are excluded from the manifest.
Use `--no-default-excludes` if you want to include all matching files and manage excludes yourself.

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
