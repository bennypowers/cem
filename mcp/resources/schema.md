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

JSON Schema for Custom Elements Manifest format validation.

Provides schema definitions for:
- Manifest structure and fields
- Element declarations
- CSS APIs and type definitions
- Package metadata

Use to validate manifests or build tooling.