# cem

**cem** is a command-line tool for generating [Custom Elements Manifest](https://github.com/webcomponents/custom-elements-manifest) (CEM) files from TypeScript sources. It analyzes your codebase and outputs a manifest describing your custom elements, classes, functions, variables, and exports, suitable for documentation tools and web component discovery.

## Features

- **Generates CEM files** from source code using syntax analysis powered by [go](https://go.dev) and [tree-sitter](https://tree-sitter.github.io/tree-sitter/).
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic Lit typescript style, with a `@customElement` decorator, and `@property` decorators on class fields.
- Supports JSDoc annotations for custom elements.
- **Supported JSDoc annotations include:**
  - `@attr` / `@attribute` — Custom element attributes
  - `@csspart` — CSS shadow parts
  - `@cssprop` / `@cssproperty` — Custom CSS properties
  - `@cssstate` — Custom CSS states
  - `@deprecated` — Marks a feature or member as deprecated
  - `@event` — Custom events dispatched by the element
  - `@slot` — Named or default slots
  - `@summary` — Short summary for documentation
- Supports scanning css files and css tagged-template-literals for css custom 
properties
  - custom properties beginning with `_` will be ignored (treated as "private")
  e.g. `var(--_private)`
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

```sh
go install bennypowers.dev/cem@latest
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
cem generate "src/**/*.ts" --exclude "src/**/*.test.ts" --output custom-elements.json
```

- `generate`: Command to start manifest generation.
- Accepts file paths and glob patterns.
- `--exclude` / `-e`: Specify patterns to exclude from the manifest.
- `--output` / `-o`: Write the manifest to a file instead of stdout.

## License

This program is free software: you can redistribute it and/or modify it under the terms of the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html).

&copy; 2025 Benny Powers <web@bennypowers.com>
