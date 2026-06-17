---
uri: cem://config/schema/{section}
name: config-schema-section
mimeType: application/json
uriTemplate: true
---

Detailed JSON Schema for a specific CEM config section.

Valid sections: generate, serve, health, mcp, export.

Returns the full JSON Schema object for the requested section, including property types, descriptions, enum constraints, and format validators. Use this to understand available fields and valid values when configuring that section.
