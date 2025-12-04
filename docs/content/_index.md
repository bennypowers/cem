---
title: cem
layout: home
---
# cem

![A hexagon with a spiderweb pattern spiraling into the bottom right corner][logo]

The blazing-fast **CLI** tool for generating and working with
**[Custom Elements Manifests][schema]**.
{.subheading}

Use `cem` to [generate][generate], [validate][validate], and [query][search]
custom elements manifests. Develop with the [dev server][serve] featuring
live reload, interactive knobs, and buildless TypeScript. Then, use the
built-in [LSP][lsp] and [MCP][mcp] servers to activate powerful, AI-native
editor features like hover documentation, autocomplete, and more.

```bash
npm install --save-dev @pwrs/cem
```

<div class="mt-3 grid-2">
  {{< cta link="/docs/installation" text="Get Started" >}}
  {{< cta link="/docs" type="secondary" text="Read the Docs" >}}
</div>

[logo]: /images/logo.svg
[schema]: https://github.com/webcomponents/custom-elements-manifest/
[generate]: ./docs/commands/generate/
[validate]: ./docs/commands/validate/
[search]: ./docs/commands/search/
[serve]: ./docs/serve/
[lsp]: ./docs/lsp/
[mcp]: ./docs/mcp/
