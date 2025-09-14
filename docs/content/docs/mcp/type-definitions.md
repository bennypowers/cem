---
title: Type Definition Guidelines
weight: 50
---

When using the `cem mcp` server, your TypeScript type definitions are passed directly to AI systems for intelligent code generation and validation. This guide explains the current type parsing approach and best practices for writing AI-friendly type definitions.

## Current Type Parsing Approach

The CEM MCP server uses a simple and effective approach to type parsing that prioritizes compatibility with AI language models:

### Union Type Processing

For union types (the most common pattern in component APIs), the type parser:

1. **Splits on `|` separator**: `"primary" | "secondary" | "danger"` becomes `["\"primary\"", "\"secondary\"", "\"danger\""]`
2. **Preserves original syntax**: All TypeScript syntax including quotes, numbers, and type references are passed verbatim
3. **Lets AI interpret meaning**: Language models excel at understanding TypeScript syntax and can handle complex patterns

### What This Means for Your Code

```typescript
// ✅ All of these work excellently with AI systems
variant: 'primary' | 'secondary' | 'danger' | 'ghost'
size: "small" | "medium" | "large"
priority: 1 | 2 | 3 | 4 | 5
theme: Color.Light | Color.Dark | 'custom'
mode: 'strict' | 'loose' | `${string}-mode`
```

The AI receives the exact TypeScript syntax and can intelligently:
- Understand string literal values
- Recognize numeric constants
- Handle type references and enum values
- Process template literal types
- Suggest appropriate values in context

## Supported Patterns

### ✅ Union Types (Recommended)

Union types work perfectly with the current approach and provide the best AI experience:

```typescript
// String literals - most common and effective
variant: 'primary' | 'secondary' | 'danger'

// Mixed types - numbers and strings
priority: 1 | 2 | 3 | 'high' | 'critical'

// Type references - preserved for AI interpretation
status: Status.Active | Status.Inactive | 'pending'

// Template literals - AI understands the pattern
className: `${string}__${string}` | 'default'
```

### ✅ Simple Types

Basic TypeScript types are passed through unchanged:

```typescript
// Boolean attributes
disabled: boolean
required: boolean

// String types (though less informative than unions)
label: string
placeholder: string

// Number types
tabIndex: number
maxLength: number
```

### ⚠️ Non-Union Types

Types without `|` separators are still passed to the AI, but provide less specific guidance:

```typescript
// AI receives "ButtonVariant" - less specific than union
variant: ButtonVariant

// AI receives "Record<string, boolean>" - complex but interpretable
classes: Record<string, boolean>

// AI receives the full type signature
handler: (event: CustomEvent) => void
```

**Recommendation**: Use union types when possible for the most specific AI guidance, but complex types will still work.

## Best Practices for AI-Friendly Types

### 1. Prefer Explicit Union Types

```typescript
// ✅ AI-friendly: Clear, specific options
size: 'small' | 'medium' | 'large'

// ❌ Less AI-friendly: Requires type definition lookup
size: ComponentSize
```

### 2. Use Meaningful String Literals

```typescript
// ✅ Self-documenting values
variant: 'primary' | 'secondary' | 'danger' | 'ghost'

// ❌ Unclear meaning requires additional context
variant: 'a' | 'b' | 'c' | 'd'
```

### 3. Combine Types with Rich Descriptions

The most effective approach combines clear types with comprehensive descriptions:

```typescript
/**
 * Controls button prominence and semantic meaning.
 * Use 'primary' for main call-to-action buttons.
 * Use 'secondary' for supporting actions.
 * Use 'danger' for destructive actions requiring confirmation.
 * Each variant ensures proper color contrast for accessibility.
 */
@property()
variant: 'primary' | 'secondary' | 'danger' | 'ghost' = 'primary'
```

This gives AI systems:
- **Specific valid values** from the union type
- **Usage guidance** from RFC 2119 keywords in the description
- **Context and constraints** for intelligent suggestions

### 4. Document Complex Types in Descriptions

For types that can't be simplified to unions, provide context in descriptions:

```typescript
/**
 * Event handler for selection changes.
 * Receives CustomEvent with detail: { selectedItems: string[], previousSelection: string[] }
 * Called when user selects/deselects items via mouse, keyboard, or programmatic changes.
 */
@property()
onSelectionChange: (event: CustomEvent<SelectionChangeDetail>) => void
```

## Testing Your Type Definitions

Use the `cem mcp` server to verify how well your types work with AI:

1. **Ask for attribute suggestions**: "What values can I use for the variant attribute?"
2. **Request examples**: "Show me how to use the size property"
3. **Test understanding**: "What's the difference between primary and secondary variants?"

If AI responses are unclear or miss important details:
- Consider simplifying complex types to unions when possible
- Add more descriptive content using RFC 2119 keywords
- Include usage examples and constraints in descriptions

## Migration from Complex Type Systems

If you're using complex type systems and want better AI integration:

### Converting Enums to Unions

```typescript
// Before: Enum-based (works but less AI-friendly)
enum ButtonSize {
  Small = 'small',
  Medium = 'medium',
  Large = 'large'
}
@property() size: ButtonSize

// After: Union-based (more AI-friendly)
@property() size: 'small' | 'medium' | 'large'
```

### Handling Imported Types

```typescript
// If you must use imported types, consider also providing unions
import { ComponentVariant } from './types'

// Option 1: Use union directly (most AI-friendly)
@property() variant: 'primary' | 'secondary' | 'danger'

// Option 2: Use both with type assertion (maintains type safety)
@property() variant: ComponentVariant & ('primary' | 'secondary' | 'danger')

// Option 3: Document valid values in description
/**
 * Component style variant. Valid values: 'primary', 'secondary', 'danger', 'ghost'
 * Use 'primary' for main actions, 'secondary' for supporting actions...
 */
@property() variant: ComponentVariant
```

## Summary

The current type parsing approach optimizes for AI effectiveness by:
- Preserving original TypeScript syntax for AI interpretation
- Focusing on union types as the most informative pattern
- Letting language models handle complex type understanding
- Encouraging rich descriptions to supplement type information

This approach ensures your components work excellently with AI assistants while maintaining type safety and developer experience. Focus on clear union types and comprehensive descriptions for the best results.