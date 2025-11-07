## Demo Rendering & URL Rewriting


### Demo Discovery (Reuse Existing)
- Use `generate/demodiscovery` package as-is
- Extract metadata from `<meta itemprop>` tags
- Generate URLs via URLPattern templates
- Store in manifest

### URL Routing
```
/components/{tag-name}/demo/{demo-slug}/
  → Renders demo with chrome

/components/{tag-name}/demo/{demo-slug}/?chrome=false
  → Renders demo HTML wrapped in a minimal document (no chrome, no knobs)

/components/{tag-name}/demo/{demo-slug}/?disable-knobs=all
  → Renders demo HTML with no knobs

/components/{tag-name}/demo/{demo-slug}/?disable-knobs[]=attributes&disable-knobs[]=slots
  → Renders demo HTML with all knobs except attributes && slots
```

### Demo URL Synchronization with Watch Mode
**Issue:** Integration between demo discovery and watch mode is underspecified.
- How are demo URLs updated when demo files are added/removed during watch?
- Does URL routing table regenerate on every manifest update?
- What happens to open browser tabs when demo file is deleted?

**Recommendation:** Document demo URL lifecycle during watch mode.

### URL Collision Handling
**Missing:** Strategy for handling URL conflicts:
- Two demos with same slug for same element
- Demo slug contains invalid URL characters
- URLPattern parameters produce duplicate URLs

**Example collision:**
```html
<!-- demo-1.html -->
<meta itemprop="name" content="Example">

<!-- demo-2.html -->
<meta itemprop="name" content="Example">
```
Both produce: `/components/my-el/demo/example/`

**Recommendation:** Add collision detection, suffix with index or hash.

### Non-Existent Element References
**Issue:** What happens when demo references elements not in manifest?
- Demo contains `<unknown-element>`
- Manifest doesn't have `unknown-element`
- Knobs panel tries to discover controls

**Options:**
1. Show demo with warning overlay
2. Return 404
3. Show demo, skip knobs for unknown elements

**Recommendation:** Option 3 for developer flexibility.

### Query Parameter Validation
**Missing:** Validation for query parameters:
- Invalid values for `chrome` (not boolean)
- Malformed `disable-knobs[]` arrays
- Conflicting parameters (`chrome=false&disable-knobs=all`)

**Recommendation:** Document parameter validation rules, return 400 for invalid inputs.
