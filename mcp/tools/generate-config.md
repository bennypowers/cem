---
name: generate_config
inputSchema:
  type: object
  properties:
    focus:
      type: string
      description: "Optional config section to focus guidance on (generate, serve, health, mcp, export). When omitted, returns guidance for all sections."
---

Generate or update CEM configuration for a project.

Returns the current config (if any), the config file path, the full JSON schema, and guidance prompts for what to investigate in the project. The agent should use its own tools to scan the project (check for TypeScript files, demo HTML files, package.json dependencies, etc.) and then write or update the config file based on its findings.

Read current resolved config via `cem://config`.
For detailed schema of a specific section, see `cem://config/schema/{section}`.
After writing config, run `validate_config` to check for errors.
