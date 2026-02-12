---
title: Export
layout: page
---

Generates framework-specific wrapper components (React, Vue, Angular) from a custom-elements manifest. The generated wrappers provide native framework DX with compile-time type checking and IDE autocomplete.

## Syntax

```bash
cem export [flags]
```

## Arguments

| Argument            | Type   | Description                                                                  |
| ------------------- | ------ | ---------------------------------------------------------------------------- |
| `--format`          | string | Framework to export: `react`, `vue`, or `angular` (exports all configured frameworks if omitted) |
| `--output, -o`      | string | Output directory for generated wrapper files                                 |
| `--strip-prefix`    | string | Prefix to strip from tag names (e.g., `demo-`)                              |
| `--package, -p`     | string | Package specifier: local path, `npm:@scope/package`, or URL                 |

## Use Cases

### Library author

As a web component library author, generate wrappers directly from your project source:

```bash
cem export --format react -o react-wrappers/ -p .
```

This reads the manifest from the current project and writes React wrapper components to `react-wrappers/`.

### App developer consuming a library

As an app developer using a web component library installed via npm, generate local wrappers using the `npm:` specifier:

```bash
npm install @example/web-components
cem export --format react -o src/generated/ -p npm:@example/web-components
```

The `npm:` specifier resolves the package through `node_modules/`, reads its custom-elements manifest, and generates typed wrappers in your project.

## Configuration

Configure export in `.config/cem.yaml` to avoid repeating flags:

```yaml
export:
  react:
    output: react-wrappers
    stripPrefix: "my-"
  vue:
    output: vue-wrappers
  angular:
    output: angular-wrappers
    moduleName: MyComponentsModule
```

Each key under `export` is a framework name. Per-framework options:

| Option        | Type   | Description                                                |
| ------------- | ------ | ---------------------------------------------------------- |
| `output`      | string | Output directory for generated wrapper files                |
| `stripPrefix` | string | Prefix to strip from tag names when generating component names |
| `packageName` | string | Override the npm package name used in import paths         |
| `moduleName`  | string | Angular NgModule name (defaults to `ComponentsModule`)     |

Running `cem export` with no `--format` flag exports all frameworks configured in the YAML file.

## Supported Frameworks

### React

Generates React components that wrap custom elements using `ref` callbacks to set properties and forward events. Each component accepts typed props corresponding to the element's attributes, properties, and events.

### Vue

Generates Vue single-file components (`.vue`) with typed props and `v-model` support where applicable.

### Angular

Generates standalone Angular components that wrap custom elements, binding inputs to properties/attributes and outputs to events. An optional NgModule re-exports all components for convenience. Use the `moduleName` option to customize the NgModule name.

## See Also

- **[Configuration Reference][config]** — All configuration options including `export`

[config]: /docs/reference/configuration/
