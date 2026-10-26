# `@pwrs/cem`

Command-line tool for generating and querying
[Custom Elements Manifest](https://github.com/webcomponents/custom-elements-manifest)
files.

This package installs a prebuilt `cem` binary for your platform. Supported
platforms: Linux and macOS (x64 and arm64), and Windows (x64 and arm64).
Node.js 22 or later is required.

## Installation

```sh
npm install --save-dev @pwrs/cem
```

Or install it globally:

```sh
npm install --global @pwrs/cem
```

## Usage

Generate a manifest for your project:

```sh
npx cem generate
```

Inspect the available commands:

```sh
npx cem --help
```

Other commands include `serve` for a component development server, `lsp` for
editor integration, `mcp` for AI assistants, `list` and `search` for querying a
manifest, and `validate` for schema checks.

## Documentation

Full documentation is at [bennypowers.dev/cem](https://bennypowers.dev/cem/).

## Contributing

See [CONTRIBUTING.md](https://github.com/bennypowers/cem/blob/main/CONTRIBUTING.md).

## License

This program is free software: you can redistribute it and/or modify it under
the terms of the [GNU General Public License v3.0 or later](https://www.gnu.org/licenses/gpl-3.0.html).
