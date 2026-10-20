---
uri: cem://element/{tagName}/methods/{memberName}
name: element-methods-member
mimeType: application/json
uriTemplate: true
---

Single class method or non-attribute field of a custom element, returned as JSON. This is the deep-link target behind the member lists in `cem://element/{tagName}/methods`.

Each member object carries:

- Methods: name, parameters (name, type, default, optional, rest), return type, static, privacy, deprecated, summary, description, inheritedFrom, source
- Fields: name, type, default, static, privacy, readonly, deprecated, summary, description, inheritedFrom, source

When a method and a field share a name, the method wins. Use this resource to fetch one member's full signature without loading the whole markdown page.
