# cem

**cem** is a command-line tool for generating [Custom Elements Manifest](https://github.com/webcomponents/custom-elements-manifest) (CEM) files from TypeScript sources. It analyzes your codebase and generates rich metadata for your custom elements, facilitating documentation, tooling, and integration.

## Features

### `cem generate`

- **Generates CEM files** from source code using syntax analysis powered by [go](https://go.dev) and [tree-sitter](https://tree-sitter.github.io/tree-sitter/).
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic Lit typescript style, with a `@customElement` decorator, and `@property` decorators on class fields.

#### HTML Template Analysis for Slots and Parts

- **Automatically detects `<slot>` and `part` attributes in your element’s `render()` template.**
- Extracts slot names and CSS shadow parts directly from HTML templates returned by the `render()` method.
- **Supports documenting slots and parts inline in your template HTML** using HTML comments with YAML blocks. For example:
  ```html
  <!--
    summary: The main slot for content
    description: |
      This slot displays user-provided content.
      Supports multiline **markdown**.
    deprecated: true
  -->
  <slot></slot>
  ```
- You can document named slots, default slots, and CSS parts:
  ```html
  <!-- slot:
        summary: Named slot summary
     part:
        summary: Part summary
  -->
  <slot name="info" part="info-part"></slot>
  ```
- The tool merges slot and part information found in templates with any provided via JSDoc, ensuring comprehensive documentation in the generated manifest.
- **Deprecation and other metadata** for slots and parts can be specified via YAML in HTML comments.

#### JSDoc
Use JSDoc comments to add metadata to your element classes, similar to other tools

- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@slot` — Named or default slots
- `@summary` — Short summary for documentation

#### CSS Custom Properties
Supports CSS Custom Properties by scanning css files and css tagged-template-literals

- Custom properties beginning with `_` will be ignored (treated as "private") e.g. `var(--_private)`
- If you provide a [Design Tokens Community 
Group](https://tr.designtokens.org/format/) format module (JSON) to `cem` via the `--design-tokens` flag,
`cem` will add metadata from your design system to any matching css variables it finds in your elements
- You can use jsdoc-like comment syntax before each var call to document your 
  variables
  ```css
  :host {
    color:
      /**
       * custom color for use in this element
       * @summary color
       * @deprecated just use the `color` property
       */
      var(--custom-color);
  }
  ```

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

for npm projects you can use `npx @pwrs/cem generate ...`, just be sure to
install the package first.

- `generate`: Command to start manifest generation.
- Accepts file paths and glob patterns.
- `--design-tokens`: path to tokens file or npm:package specifier
- `--exclude` / `-e`: Specify patterns to exclude from the manifest.
- `--output` / `-o`: Write the manifest to a file instead of stdout.


## Local Windows Build Using Podman

Running podman will output a cem.exe file in the root directory

```sh
podman build -t cem-windows .
podman run --rm -v $(pwd):/output:Z cem-windows cp cem.exe /output/
```

## License

This program is free software: you can redistribute it and/or modify it under the terms of the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html).

&copy; 2025 Benny Powers <web@bennypowers.com>
