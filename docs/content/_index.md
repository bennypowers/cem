---
title: cem
layout: home
---
# cem

![A hexagon with a spiderweb pattern spiraling into the bottom right corner](/images/logo.svg)

The blazing-fast **CLI** tool for generating and working with **[Custom Elements Manifests][schema]**. 
{.subheading}

Use `cem` to [generate][generate], [validate][validate], and [query][search] 
custom elements manifests. Then, use the built-in [LSP server][lsp] to activate powerful 
editor features like hover documentation, autocomplete, and more.

```bash
npm install --save-dev @pwrs/cem
```

<div class="mt-3 grid-2">
  {{< cta link="/docs/installation" text="Get Started" >}}
  {{< cta link="/docs" type="secondary" text="Read the Docs" >}}
</div>

[schema]: https://github.com/webcomponents/custom-elements-manifest/
[generate]: ./docs/commands/generate/
[validate]: ./docs/commands/validate/
[search]: ./docs/commands/search/
[lsp]: ./docs/lsp/
