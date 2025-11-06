- Always use Makefile targets for running tests or builds, since they export the necessary env vars

## Performance Optimizations

### Attribute/Element Lookup Maps

The codebase uses map-based lookups for performance-critical operations in the manifest package:

- **Attribute field lookups**: `CustomElementDeclaration.LookupAttributeField(attrName)`
  - O(1) map lookup vs O(n) linear search through Members
  - Lazy-initialized cache in `attributeFieldMap` field
  - 1.2x-17.7x faster depending on member count

- **Export lookups**: `Module.LookupCustomElementExport(name)` / `LookupJavaScriptExport(name)`
  - O(1) map lookup vs O(n) linear search through Exports
  - Lazy-initialized caches in `customElementExportMap` and `jsExportMap` fields
  - 4.4x-11.4x faster for medium/large export counts

**Implementation details**:
- Maps are `json:"-"` tagged (won't affect JSON serialization)
- Built on-demand via lazy initialization with `sync.Once` (minimal memory overhead)
- Thread-safe for concurrent reads and initialization
- Helper functions exported for testing: `BuildAttributeFieldMap()`, `BuildExportMaps()`

**Benchmarks**: Run `make bench-lookup` to see A/B comparison
**Memory impact**: Negligible for small structures, beneficial for large
**Thread safety**: Uses sync.Once to ensure safe concurrent initialization