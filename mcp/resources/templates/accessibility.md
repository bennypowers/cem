# Accessibility Intelligence 🌟

## What's Built-In
**{{.FetchedData.accessibility.metadata.totalPatterns}} accessibility patterns** discovered in your element manifests

## Elements with Accessibility Smarts 🧠

{{range .FetchedData.accessibility.patterns}}
### `{{.element}}` knows about {{.pattern}}
{{.description}}

**✅ Follow the Element's Lead:**
- The element author documented specific accessibility requirements
- Trust their implementation and follow their guidelines
- Test the documented interaction patterns
- Verify the behavior matches the manifest descriptions

---
{{end}}

## Manifest-Driven Accessibility 📖

These patterns come directly from element authors who documented accessibility considerations in their manifests. Here's how to leverage them:

### 🎯 Trust the Documentation
- **Follow manifest guidelines** - Element authors know their accessibility requirements best
- **Use documented patterns** - Don't reinvent accessibility features that are built-in
- **Respect element semantics** - Custom elements often have non-obvious but intentional accessibility behavior
- **Check descriptions thoroughly** - Accessibility requirements are often embedded in element and attribute descriptions

### 🧪 Test the Right Things
- **Validate manifest claims** - Test that documented accessibility features actually work
- **Follow documented interactions** - Use elements as their authors intended
- **Test with real users** - Verify the experience matches documented expectations
- **Monitor consistency** - Ensure similar elements behave similarly

### 🚫 Manifest-First Mistakes to Avoid
- **Don't override documented behavior** - The element author likely tested their accessibility implementation
- **Don't add undocumented ARIA** - You might conflict with internal element logic
- **Don't ignore accessibility notes** - They're there for important reasons
- **Don't assume standard patterns** - Custom elements may implement accessibility differently

## Your Accessibility Action Plan 📋

When using custom elements, verify you're following their accessibility guidance:

### ✅ Manifest Compliance Check
1. **Read element descriptions** - Look for accessibility requirements and patterns
2. **Check attribute guidelines** - Some attributes may have accessibility implications
3. **Follow slot content rules** - Content structure affects screen reader navigation
4. **Respect event patterns** - Elements may emit accessibility-related events
5. **Use CSS states appropriately** - Visual states often correlate with accessibility states

### 🔍 Testing Your Implementation
- **Test documented behaviors** - Verify accessibility features work as described
- **Validate content structure** - Ensure slotted content maintains accessibility
- **Check interaction patterns** - Follow the element's documented interaction model
- **Monitor state changes** - Verify state transitions are accessible

## Manifest-Sourced Wisdom 💡

The best accessibility guidance comes from the element authors themselves. Look for:

### 📖 In Element Descriptions
- ARIA pattern implementations
- Keyboard interaction requirements
- Screen reader behavior notes
- Focus management details

### ⚙️ In Attribute Documentation
- Required accessibility attributes
- ARIA state mappings
- Semantic value constraints
- Interaction behavior modifiers

### 🎯 In Slot Guidelines
- Content structure requirements
- Heading hierarchy expectations
- Interactive element placement rules
- Screen reader navigation patterns

---

## Dive Deeper 🏊‍♀️

• **[Element Details](cem://element/{tagName})** — Complete accessibility documentation
• **[Attribute Guide](cem://element/{tagName}/attributes)** — Accessibility-related attributes
• **[Usage Guidelines](cem://guidelines)** — Manifest-sourced best practices
• **[Element Library](cem://elements)** — Browse elements by accessibility features