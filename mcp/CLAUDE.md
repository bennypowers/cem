# CEM MCP Server Instructions

## 🎯 Core Philosophy: Data + Context + LLM Decision Making

CEM MCP operates on this fundamental principle:

1. **📊 Take User Data** - Raw manifest data (properties, parts, states, descriptions)
2. **📖 Add Rich Context** - Schema definitions, usage patterns, constraints, and guidance
3. **🧠 Let LLM Decide** - Present enriched data to AI for intelligent decision making

**NOT** hardcoded outputs → **YES** templated context that guides LLM reasoning

### Key Rules

- **No hardcoded suggestions** → Rich manifest data + contextual templates
- **LLM-driven decisions** → AI chooses appropriate values based on provided constraints
- **Respect element authors** → Use their documented constraints and intent
- **Manifest-driven accessibility** → Follow documented accessibility patterns, don't add generic ARIA advice

## Template System

All responses use Go templates with embedded manifest data:

```
mcp/resources/templates/
├── element.md             # Element overview and API summary
├── element-attributes.md  # Attribute documentation and constraints
├── element-events.md      # Event patterns and JavaScript integration
├── element-slots.md       # Slot usage and content guidelines
├── element-styling.md     # CSS customization with properties, parts, states
├── elements.md            # Elements discovery and capabilities overview
├── accessibility.md       # Accessibility patterns from manifests
└── guidelines.md          # Usage guidelines and best practices
```

**Template naming**: Use dashes (e.g., `element-attributes.md`), not underscores.

## Architecture Principles

### Resources vs Tools
- **Resources** (read-only): Information access via `cem://` URIs
- **Tools** (actions): Concrete actions like `generate_html`, `validate_html`

### Resource Structure
- `cem://element/{tagName}` - Complete element reference
- `cem://element/{tagName}/attributes` - Attribute documentation
- `cem://element/{tagName}/slots` - Slot usage patterns
- `cem://element/{tagName}/events` - Event patterns
- `cem://element/{tagName}/css` - CSS customization

### Element Descriptions
- Reference **subresources** not tools: `cem://element/{tagName}/attributes`
- Don't reference tool names like `element_attributes`

## Development Guidelines

- Use Makefile targets for tests/builds (exports necessary env vars)
- Never create temporary files for testing - use fixture patterns
- Add regression tests for all features using established fixture patterns
- Never introduce regressions just to meet performance goals
