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

