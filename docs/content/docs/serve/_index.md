---
title: Development Server
weight: 25
---

`cem serve` is a specialized dev server for custom elements, offering live reload, interactive knobs, and component isolation based on your manifest demos.

{{< card-grid >}}
  {{< card title="Getting Started" href="getting-started/" icon="/images/sections/getting-started.svg" >}}
  Set up your first demo
  {{< /card >}}

  {{< card title="Buildless Development" href="buildless/" icon="/images/sections/buildless.svg" >}}
  TypeScript and CSS without build steps
  {{< /card >}}

  {{< card title="Knobs" href="knobs/" icon="/images/sections/knobs.svg" >}}
  Interactive testing controls
  {{< /card >}}

  {{< card title="Import Maps" href="import-maps/" icon="/images/sections/import-maps.svg" >}}
  Use npm packages without bundling
  {{< /card >}}

  {{< card title="Command Reference" href="reference/" icon="/images/sections/reference.svg" >}}
  Command flags and usage
  {{< /card >}}
{{< /card-grid >}}

## Quick Start

```sh
cem serve
```

Starts the server on `http://localhost:8000`.

{{<tip "warning">}}
`cem serve` is for developing custom elements in isolation, not a general-purpose dev server.
{{</tip>}}

## Features

- **Manifest-Driven:** Demos derive directly from your `custom-elements.json`.
- **Demo Chrome:** A PatternFly-based UI with a navigation drawer, component documentation, and theme toggles.
- **Interactive Knobs:** Auto-generated controls for attributes, properties, and CSS variables.
- **TypeScript Support:** On-the-fly compilation with source maps.
- **Workspace Support:** Seamlessly handles multi-package repositories.

## Navigation

The server lists all elements at the root (`/`) and provides a persistent drawer to jump between demos. You can also access demos directly via `/demo/{element-name}/` or `/pkg/{scope}/{package}/demo/{element-name}/` for workspaces.

## See Also

{{< card-grid >}}
  {{< card title="Generate Docs" href="/docs/commands/generate/" icon="/images/sections/generate.svg" >}}
  Document components
  {{< /card >}}

  {{< card title="Configuration" href="/docs/configuration/" icon="/images/sections/configuration.svg" >}}
  Config reference
  {{< /card >}}
{{< /card-grid >}}
