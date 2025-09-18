# Your Component Library 📚

## What's Available
**{{.FetchedData.elements.metadata.totalElements}} custom elements** ready to use in your project!

{{if .FetchedData.elements.metadata.categories}}
## Element Categories

{{range $category, $count := .FetchedData.elements.metadata.categories}}{{if gt $count 0}}**{{$category | title}}:** {{$count}}{{if eq $category "form-elements"}} elements{{else if eq $category "layout-elements"}} elements{{else if eq $category "navigation-elements"}} elements{{else}} elements{{end}}
{{end}}{{end}}{{end}}

## Available Elements

| Tag Name       | Package        | Module       | Summary                            |
| -------------- | -------------- | ------------ | ---------------------------------- |
{{range .FetchedData.elements.elements}}
| `{{.tagName | sanitize}}` | `{{.package | sanitize}}` | `{{.module | sanitize}}`| {{if .summary}}{{.summary | sanitize}}{{end}} |
{{end}}

## Element Capabilities

Use this overview to find elements by their functional capabilities:

- **Configurable** - Elements with custom attributes for behavior control
- **Content-slots** - Elements that accept slotted content
- **Interactive** - Elements that emit custom events
- **Themeable** - Elements with CSS custom properties for styling
- **Styleable** - Elements with CSS parts for targeted styling
- **Stateful** - Elements with CSS custom states for conditional styling
- **Form-elements** - Components for form construction and input handling
- **Layout-elements** - Components for page structure and content organization
- **Navigation-elements** - Components for site navigation and user guidance

## Smart Element Selection 🧠

### ✅ Choose Wisely
1. **Match functionality to needs** - Don't over-engineer with complex elements for simple tasks
2. **Stick to one design system** - Mix-and-match can create inconsistent experiences
3. **Consider maintenance** - Simpler elements = fewer breaking changes

### 🚫 Common Selection Mistakes
- **Don't** pick elements just because they have more features
- **Don't** mix components from different packages without checking compatibility
- **Don't** ignore accessibility requirements when choosing interactive elements
- **Don't** assume all elements work well together without testing

### Integration Wisdom 💡
- **Configurable elements** need thoughtful default values
- **Interactive elements** require proper event handling
- **Themeable elements** work best with design system integration
- **Layout elements** should complement your existing CSS architecture

---

For detailed element information, use:
- **`cem://element/{tagName}`** - Complete element reference
- **`cem://element/{tagName}/attributes`** - Focused attribute documentation
- **`cem://packages`** - Package-level organization and structure
- **`cem://guidelines`** - Usage guidelines and best practices
