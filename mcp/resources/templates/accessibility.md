# Accessibility Patterns

## Summary
{{.FetchedData.accessibility.metadata.totalPatterns}} accessibility patterns identified in manifest documentation

## Accessibility-Aware Elements

{{range .FetchedData.accessibility.patterns}}
### `{{.element}}` - {{.pattern}}
{{.description}}

**Implementation Notes:**
- Follow WCAG 2.1 AA guidelines for this element
- Test with screen readers and keyboard navigation
- Ensure proper ARIA attributes when required
- Validate color contrast and focus indicators

---
{{end}}

## Accessibility Implementation Guide

These patterns are identified from elements that include accessibility-related information in their documentation. Use them to:

### WCAG Compliance
- **Semantic markup** - Use elements according to their semantic meaning
- **Keyboard navigation** - Ensure all interactive elements are keyboard accessible
- **Screen reader support** - Implement proper ARIA labels and descriptions
- **Focus management** - Provide clear focus indicators and logical tab order

### Testing Requirements
- **Automated testing** - Use accessibility testing tools during development
- **Manual testing** - Test with actual assistive technologies
- **User testing** - Include users with disabilities in testing processes
- **Continuous monitoring** - Regularly audit accessibility in production

### Implementation Patterns
- **Progressive enhancement** - Build accessible foundations first
- **Semantic HTML** - Use appropriate HTML elements as the base
- **ARIA enhancement** - Add ARIA attributes only when necessary
- **Custom element accessibility** - Ensure custom elements maintain accessibility semantics

## Accessibility Checklist

When implementing custom elements, verify:

1. **Semantic structure** - Element roles and relationships are clear
2. **Keyboard interaction** - All functionality available via keyboard
3. **Screen reader compatibility** - Content and state changes are announced
4. **Focus management** - Focus moves logically and is visible
5. **Color and contrast** - Text meets WCAG contrast requirements
6. **Animation and motion** - Respects user motion preferences
7. **Error handling** - Errors are clearly communicated
8. **Form accessibility** - Form elements have proper labels and validation

## Common Accessibility Patterns

### Interactive Elements
- Ensure keyboard event handlers complement mouse events
- Provide clear focus indicators
- Implement proper ARIA states (expanded, pressed, selected)
- Support expected keyboard shortcuts

### Content Elements
- Use semantic heading structure
- Provide alternative text for images
- Ensure link text is descriptive
- Structure content with proper landmarks

### Form Elements
- Associate labels with form controls
- Provide clear error messages
- Group related fields appropriately
- Support keyboard navigation between fields

---

For element-specific accessibility guidance, use:
- **`cem://element/{tagName}`** - Element documentation including accessibility notes
- **`cem://element/{tagName}/attributes`** - ARIA and accessibility-related attributes
- **`cem://guidelines`** - General usage guidelines including accessibility requirements
- **`cem://elements`** - Browse elements with accessibility capabilities