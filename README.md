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

> [!NOTE]
> `cem generate` best supports LitElements written in idiomatic style with
> TypeScript decorators. There is rudimentary support for `extends HTMLElement`,
> but it is not a high priority for development. If you need something more
> specific [open an issue][issuenew].

See more in the [Generate docs][generatedocs]

---

### `cem list`

The `cem list` command provides a fast, flexible way to inspect custom elements, their features, and their metadata directly from your manifest file.
With `cem list`, you can quickly explore and audit your custom elements API surface, making it easier to document, test, and share your components.

See more in the [List docs][listdocs]

See the [Configuration Reference][configdocs] for more information.

---

## Contributing

For information on building and testing, please see
[CONTRIBUTING.md][contributingmd].

## License

This program is free software: you can redistribute it and/or modify it under
the terms of the [GNU General Public License v3.0][gpl3].

&copy; 2025 Benny Powers

[cem]: https://github.com/webcomponents/custom-elements-manifest
[dtcg]: https://tr.designtokens.org/format/
[go]: https://go.dev
[treesitter]: https://tree-sitter.github.io/tree-sitter/
[gpl3]: https://www.gnu.org/licenses/gpl-3.0.html
[contributingmd]: https://bennypowers.github.io/cem/docs/contributing/
[issuenew]: https://github.com/bennypowers/cem/issues/new
[installationdocs]: https://bennypowers.github.io/cem/docs/installation/
[generatedocs]: https://bennypowers.github.io/cem/docs/commands/generate/
[listdocs]: https://bennypowers.github.io/cem/docs/commands/list/
[configdocs]: https://bennypowers.github.io/cem/docs/configuration/
