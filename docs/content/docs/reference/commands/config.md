---
title: Config
description: Inspect and manage CEM configuration
weight: 80
---

{{< tip >}}
**TL;DR**: Use `cem config init` to set up a new project, `cem config show` to see your resolved config, `cem config validate` to check for problems, and `cem config mcp` to generate MCP snippets for your AI tools.
{{< /tip >}}

The `cem config` command group helps you create, inspect, and validate CEM configuration files.

## Subcommands

### `cem config init`

Interactive wizard that detects your project's settings and generates a `.config/cem.yaml` file.

```bash
cem config init
```

The wizard inspects your project for TypeScript source files, existing manifests, demo HTML files, `tsconfig.json` settings, and git remote URLs, then pre-fills sensible defaults. You can review and adjust each setting before the config file is written.

| Flag | Description |
| ---- | ----------- |
| `--format` | Output format: `yaml` (default), `json`, or `jsonc` |
| `-y`, `--yes` | Accept detected defaults without prompting |

### `cem config show`

Print the fully resolved configuration, including workspace inheritance.

```bash
cem config show
cem config show --format json
```

This shows the merged result after config file loading and workspace cascade, so you can verify exactly what `cem` will use.

| Flag | Description |
| ---- | ----------- |
| `--format` | Output format: `yaml` (default) or `json` |

### `cem config validate`

Check the config file for invalid values, unreachable glob patterns, and missing files.

```bash
cem config validate
```

Runs both schema validation (correct field names and types) and semantic checks (do your glob patterns match any files? does your output directory exist? are URL rewrites valid?). Reports errors with file, line, and column positions.

| Flag | Description |
| ---- | ----------- |
| `--format` | Output format: `text` (default) or `json` |

### `cem config mcp`

Generate MCP configuration snippets for AI tools.

```bash
cem config mcp
```

In interactive mode (TTY), presents a menu to select your tools and configure options. In non-interactive mode, use `--tool` to specify tools directly:

```bash
cem config mcp --tool "Claude Code" --tool "Cursor"
```

Supported tools: Claude Code, Claude Desktop, Cursor, VS Code (Copilot), Zed, and Other.

| Flag | Description |
| ---- | ----------- |
| `--tool` | AI tool(s) to configure (repeatable) |
| `-a`, `--additional-packages` | Additional packages to include in MCP args (npm:, jsr:, or URL specifiers) |

### `cem config path`

Print the resolved config file path. Useful for scripting:

```bash
$EDITOR $(cem config path)
```
