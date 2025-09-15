---
uri: cem://schema
name: schema
mimeType: application/json
dataFetchers:
  - name: schema
    type: manifest_schema
    path: ""
    required: true
template: schema
---

JSON Schema for Custom Elements Manifest format. This schema defines the structure and validation rules for custom elements manifest files, ensuring consistency and correctness across component definitions.

Provides comprehensive schema validation for:
- Manifest structure and required fields
- Module definitions and exports
- Custom element declarations with attributes, slots, events
- CSS custom properties, parts, and states
- Type definitions and inheritance
- Package metadata and dependencies
- Documentation and example formats

Use this resource to understand the manifest format, validate manifest files, or generate tooling that works with custom elements manifests.