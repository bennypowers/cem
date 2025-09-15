---
uri: cem://packages/{package}
name: package
mimeType: text/markdown
uriTemplate: true
dataFetchers:
  - name: package
    type: package_collection
    path: ""
    required: true
template: package
---

Detailed information about a specific package in the workspace, including all elements within the package, module organization, and package-level metadata. This resource provides focused package exploration for understanding component organization and planning imports.

Provides comprehensive package details including:
- Complete element listings within the package
- Module structure and organization patterns
- Package-level metadata and configuration
- Element distribution and capabilities within the package
- Cross-package relationships and dependencies
- Import patterns and module resolution guidance
- Package-specific guidelines and conventions

Use this resource when working with a specific package to understand its organization, discover available elements, or plan component imports and usage within that package's ecosystem.

For related package information, explore these resources:

- **`cem://packages`** - Complete listing of all packages in the workspace
- **`cem://elements`** - Browse all elements across packages with package context
- **`cem://element/{tagName}`** - Detailed information about specific elements within packages
- **`cem://guidelines`** - Design system guidelines that may apply at package level