---
name: validate_config
inputSchema:
  type: object
  properties: {}
---

Validate the current CEM configuration file.

Runs schema validation and semantic checks against the project's config file, including filesystem I/O checks (glob pattern reachability, output path existence, URL rewrite validation, demo discovery validation).

Returns structured JSON with config file path, validity status, and arrays of errors and warnings. Each error includes the field path, message, severity, and source position (line/column) when available.
