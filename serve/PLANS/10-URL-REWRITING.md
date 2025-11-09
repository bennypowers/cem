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
- How are demo URLs updated when demo files are added/removed during watch? 
Answer: This is handled in `generate/demodiscovery`
- Does URL routing table regenerate on every manifest update? Answer: If 
manifest changes are relevant to demo urls (e.g. new or removed demos), yes.
- What happens to open browser tabs when demo file is deleted? Nothing, 404 on 
manual reload.

### URL Collision Handling
- Two demos with same slug for same element? error in overlay, error logged in 
  console. Possibly exit 1 and log error on startup
- Demo slug contains invalid URL characters? they get slugified
- URLPattern parameters produce duplicate URLs? log error in console, possibly 
exit 1

**Example collision:**
```html
<!-- demo-1.html -->
<meta itemprop="name" content="Example">

<!-- demo-2.html -->
<meta itemprop="name" content="Example">
```
Both produce: `/components/my-el/demo/example/`

**Recommendation:** Add collision detection, and error out.

### Non-Existent Element References
**Issue:** What happens when demo references elements not in manifest?
- Demo contains `<unknown-element>`
- Manifest doesn't have `unknown-element`
- Knobs panel tries to discover controls

Answer: LSP catches this and reports in editor. Skip knobs for that element

### Query Parameter Validation
- Invalid values for `chrome` (not boolean)
- Malformed `disable-knobs[]` arrays
- Conflicting parameters (`chrome=false&disable-knobs=all`)
- values are valid when they conform to the config schema

**Recommendation:** return 400 for invalid inputs, and log validation errors to console (and to error overlay)
