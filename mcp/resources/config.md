---
uri: cem://config
name: config
mimeType: application/json
dataFetchers:
  - name: config
    source: config
    path: ""
    required: true
responseType: json
---

Resolved CEM project configuration as JSON.

Includes all config sections:
- **generate**: source file globs, output path, design tokens, demo discovery
- **serve**: dev server port, import map (auto-generation and overrides), TypeScript/CSS transforms, URL rewrites, demo rendering mode
- **health**: score threshold, disabled checks
- **mcp**: MCP server settings (description length limits)
- **export**: framework-specific wrapper generation

Also includes top-level fields: projectDir, packageName, sourceControlRootUrl, additionalPackages.

Use `cem://config/schema` for field descriptions, valid values, and defaults.
Use the `validate_config` tool to check config for errors.
Use the `generate_config` tool for guidance on creating or updating config.
