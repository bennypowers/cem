---
uri: cem://packages
name: packages
mimeType: application/json
dataFetchers:
  - name: packages
    type: package_overview
    path: ""
    required: true
responseType: json
---

Listing of all packages containing custom elements manifests.

Provides:
- Package names and element counts
- Module organization
- Package dependencies
- Export patterns

Use to discover packages and plan imports.