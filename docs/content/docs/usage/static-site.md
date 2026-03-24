---
title: Static Site Generation
layout: docs
weight: 55
---

{{< tip >}}
**TL;DR**: Run `cem serve --build` to generate a static website from your component demos. The output includes server-rendered HTML with Declarative Shadow DOM, transformed source files, and vendored dependencies -- ready to deploy anywhere.
{{< /tip >}}

## Overview

The `--build` flag renders all your demo pages into a self-contained static site. Each page goes through the same middleware pipeline as the dev server: TypeScript transformation, CSS module wrapping, import map injection, and Lit SSR for Declarative Shadow DOM.

```sh
cem serve --build -o dist/
```

This produces a directory you can serve with any static file server:

```sh
python3 -m http.server -d dist/
# or
npx serve dist/
```

## What gets built

| Output | Description |
| ------ | ----------- |
| Demo pages | SSR-rendered HTML with Declarative Shadow DOM for all demos |
| Index page | Listing of all components with navigation |
| User sources | TypeScript transformed to JavaScript, CSS wrapped as modules |
| Dependencies | Vendored from `node_modules/` (or rewritten to CDN URLs) |
| Import map | Embedded in each page for bare specifier resolution |
| Chrome bundle | Single JS file with all dev UI components + Lit hydration |
| `sitemap.xml` | Lists all built pages |
| `custom-elements.json` | The generated manifest |
| Health data | Pre-rendered at `__cem/api/health` as static JSON |

## Dependency resolution

By default, `--build` vendors dependencies from `node_modules/` into the output directory. Use `--import` to resolve dependencies from a CDN instead:

```sh
# Vendor locally (default)
cem serve --build --import vendor

# Use esm.sh CDN
cem serve --build --import esm

# Use jspm CDN
cem serve --build --import jspm

# Use unpkg CDN
cem serve --build --import unpkg
```

CDN modes rewrite the import map entries from local paths to CDN URLs and skip copying `node_modules/`. This produces a smaller output that loads dependencies from the network.

## Deploying to a subdirectory

When the static site is hosted under a subdirectory (e.g., `https://example.com/docs/components/`), use `--base-path` to prefix all URLs:

```sh
cem serve --build -o dist/ --base-path /docs/components/
```

This rewrites asset links, import maps, navigation hrefs, and dynamic imports so they resolve correctly under the given path.

## Example

```sh
# Build the kitchen-sink example to /tmp/site
cem serve --build -p examples/kitchen-sink -o /tmp/site

# Serve it locally
python3 -m http.server -d /tmp/site
```

## Comparison with the dev server

| Feature | `cem serve` | `cem serve --build` |
| ------- | ----------- | ------------------- |
| Live reload | Yes | No |
| TypeScript transform | On-demand | At build time |
| SSR (Declarative Shadow DOM) | Yes | Yes |
| Import maps | Injected per-request | Embedded in HTML |
| WebSocket client | Active | Stubbed (no-op) |
| Health API | Live endpoint | Pre-rendered JSON |
| Output | HTTP responses | Static files on disk |
