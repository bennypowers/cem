# CEM MCP Server Implementation Plan

## 🎯 Core Philosophy: Data + Context + LLM Decision Making

**The Central Approach**: CEM MCP operates on a fundamental principle:

1. **📊 Take User Data** - Raw manifest data (properties, parts, states, descriptions)
2. **📖 Add Rich Context** - Schema definitions, usage patterns, constraints, and guidance 
3. **🧠 Let LLM Decide** - Present enriched data to AI for intelligent decision making

**NOT** hardcoded outputs → **YES** templated context that guides LLM reasoning

### Template-Driven Intelligence

```
Manifest Data → Go Templates → Rich Context → LLM → Smart Decisions
     ↓              ↓              ↓           ↓           ↓
  Raw APIs    +  Usage Patterns  =  Guided   →  AI chooses  →  Proper
  Syntax info    Best practices     Context     appropriate    Usage
  Descriptions   Constraints                    values/styles
```

**Example in Practice:**
- **Don't generate**: `color: #007acc` (hardcoded)
- **Do provide**: `color: /* your color value (syntax: <color>) */` + property description
- **LLM decides**: Appropriate color based on element's purpose and context

**Why This Matters:**
- ✅ **Respects element authors** - Uses their documented constraints and intent
- ✅ **Scales to any manifest** - Works with any element, any property definition
- ✅ **Context-aware** - LLM considers element purpose, design system, accessibility
- ✅ **Maintainable** - Content updates don't require code changes
- ❌ **Avoids assumptions** - No hardcoded values that might be inappropriate

This approach respects element authors' documented intent while enabling intelligent AI assistance.

### 🌐 Template System Architecture

**Implementation**: All tool responses use Go templates with embedded manifest data:

```
mcp/resources/templates/
├── element.md             # Element overview and API summary
├── element-attributes.md  # Attribute documentation and constraints
├── element-events.md      # Event patterns and JavaScript integration
├── element-slots.md       # Slot usage and content guidelines
├── element-styling.md     # CSS customization with properties, parts, states
├── elements.md            # Elements discovery and capabilities overview
├── packages.md            # Package structure and organization
├── schema.md              # JSON schema definitions and validation
├── accessibility.md       # Accessibility patterns and requirements
└── guidelines.md          # Usage guidelines and best practices
```

**Template Features**:
- **Data-driven**: `{{.Name}}`, `{{.Description}}`, `{{.Syntax}}`
- **Conditional**: `{{if .Initial}}...{{end}}`
- **Iterative**: `{{range .CssProperties}}...{{end}}`
- **Smart logic**: `{{if gt (len .CssParts) 1}}` for plural handling

**Benefits**:
✓ **Maintainable**: Content in markdown, not Go code
✓ **Extensible**: New guidance = new template file
✓ **Consistent**: Same data model across all tools
✓ **Context-aware**: Templates adapt to actual manifest content
✓ **LLM-friendly**: Rich context without hardcoded assumptions

---

## Intelligent HTML Generation Platform

The CEM MCP server provides an AI-native component intelligence platform that enables intelligent HTML generation with:
- **Correct slot and attribute usage** based on manifest definitions
- **Proper CSS integration** with custom properties, parts, and states  
- **Guideline compliance** following element/attribute descriptions and best practices
- **Context-aware suggestions** that understand component relationships and design patterns

## Core Concept Analysis 

 MCP (Model Context Protocol) provides the perfect bridge between AI systems and your rich custom elements ecosystem. After analyzing the CEM codebase and MCP specification, here's my comprehensive plan with critical insights and creative enhancements.

## Critical Insights & Improvements

### 1. **Leverage Existing LSP Infrastructure** 
The existing LSP server (`cmd/lsp.go`) already solves 80% of the architecture challenges:
- Manifest discovery and registry system
- File watching and auto-reload
- Workspace context management  
- Performance optimizations (tree-sitter, incremental parsing)

**Strategy**: Refactor shared components into a common service layer that both LSP and MCP can consume.

### 2. **Template-Driven Context Enhancement**
Building on manifest data, the template system provides **rich, contextual guidance**:

**Resources (Read-only data)**:
- Live manifest registry with real-time updates
- Custom element definitions with inheritance chains
- Type information and validation rules
- **Template-generated usage patterns** from manifest descriptions
- CSS integration patterns (custom properties, parts, states)
- Design system guidelines and component relationships

**Tools (Template-powered capabilities)**:
- Element validation with **context-aware suggestions**
- Import path resolution
- Component compatibility checking
- Real-time manifest querying with **templated responses**
- **HTML generation** using manifest data + template guidance
- **CSS integration guidance** via templates that respect property syntax/constraints

### 3. **Multi-Modal Context Delivery**
**Smart Context Adaptation**: Unlike static documentation, MCP can provide context that adapts to the AI's current task:
- **Code Generation**: Focus on attributes, slots, events, proper HTML structure
- **Documentation**: Emphasize descriptions, examples, demos  
- **Debugging**: Highlight validation rules, common issues
- **Architecture**: Show inheritance chains, dependencies
- **Styling**: CSS custom properties, parts, states, design tokens

## Implementation Architecture

### Core MCP Server
```
cmd/mcp.go              # MCP command implementation
mcp/
├── server.go           # MCP server implementation
├── resources/          # Declarative resource framework with YAML configs
│   ├── *.md            # Resource configurations with YAML frontmatter
│   ├── templates/      # Go template files for rich context rendering
│   └── declarative_framework.go  # Data fetchers and template engine
├── tools/              # Interactive tools for HTML generation and validation
├── context.go          # Manifest registry and workspace context
└── templates/          # Core template rendering engine
```

**Available Resources**:
- `cem://schema` - JSON schema for custom elements manifests
- `cem://packages` - Package discovery and manifest overview
- `cem://elements` - Summaries of available elements with capabilities
- `cem://element/{tagName}` - Individual element definitions with usage patterns
- `cem://guidelines` - Design system guidelines and component usage rules
- `cem://accessibility` - Accessibility patterns and validation rules

### Template-Powered Tools
**Interactive MCP Tools Using Template-Driven Guidance**:
- `validate_html` - Validate custom element usage based on manifest guidelines and best practices
- `suggest_attributes` - Present syntax definitions + usage guidance for LLM decisions
- `generate_html` - Template-driven HTML structure respecting manifest slot/attribute definitions
- `suggest_css_integration` - Template system presenting CSS APIs with rich context for LLM

**Template Philosophy**:
- **No hardcoded suggestions** → Rich manifest data + contextual templates
- **LLM-driven decisions** → AI chooses appropriate values based on provided constraints
- **Extensible guidance** → New templates = new capabilities without code changes

### HTML Generation Features
**Context Intelligence for HTML Generation**:
- **Semantic chunking**: Break large manifests into focused chunks for specific HTML contexts
- **Relationship mapping**: Show element dependencies and nesting patterns
- **Usage patterns**: Common attribute combinations, slot usage, CSS integration
- **Guideline enforcement**: Apply element/attribute description guidelines automatically
- **Template generation**: Create complete HTML templates with proper element usage

## Creative Enhancements for HTML Generation

### 1. **Live Development Context**
Beyond static manifests, provide **real-time development state**:
- Elements currently being edited (via file watching)
- Validation errors in current workspace
- Recently modified components with updated guidelines
- Development vs production manifest differences
- **CSS changes affecting component styling**

### 2. **AI-Native Documentation for HTML Generation**
Transform dry manifests into **AI-friendly HTML generation guides**:
- Natural language descriptions of complex types → HTML attribute values
- Usage examples with context about when/why to use specific slots
- Common pitfalls and solutions for proper HTML structure
- Design system integration notes for consistent styling
- **CSS custom property usage patterns**
- **Slot content guidelines and best practices**

### 3. **Cross-Package Intelligence for Component Ecosystems**
Leverage your multi-package discovery to provide **ecosystem insights**:
- Similar elements across packages (find alternatives for HTML generation)
- Breaking changes between versions affecting HTML usage
- Package compatibility matrices for consistent styling
- Migration guidance between similar components
- **CSS design token compatibility across packages**

### 4. **Generate-on-Demand Integration**
Extend beyond static manifests with **dynamic generation**:
- Trigger `cem generate` for specific files
- Return fresh manifest data for modified code
- Validate proposed HTML changes before writing
- Preview manifest changes from code diffs
- **Live CSS custom property validation**

## Architecture Innovations

### Unified Service Layer
```go
// Shared between LSP and MCP
type ManifestService interface {
    GetRegistry() *Registry
    QueryElements(criteria) []Element
    ValidateUsage(element, context) ValidationResult
    ResolveImports(element) ImportInfo
    // HTML Generation specific
    GenerateHTML(element, options) HTMLGenerationResult
    ValidateHTML(html, context) HTMLValidationResult
    GetCSSIntegration(element) CSSIntegrationInfo
}
```

### Context-Aware Resource Delivery for HTML Generation
```go
type MCPResource struct {
    URI         string
    Content     string
    MimeType    string
    Metadata    map[string]any
    // AI optimization for HTML generation
    Summary     string           // Key points for AI
    Context     ResourceContext  // When this is most relevant
    Related     []string         // URIs of related resources
    HTMLPatterns []HTMLPattern   // Common HTML usage patterns
    CSSGuidance CSSIntegration   // CSS custom properties and styling
}

type HTMLPattern struct {
    Description string
    Example     string
    Guidelines  []string
    CommonSlots []SlotUsage
    CSSClasses  []string
}
```

## Strategic Benefits

### For CEM Project
- **New distribution channel**: MCP servers can be shared across AI platforms
- **Enhanced adoption**: Lower barrier to custom element discovery and proper usage
- **Ecosystem growth**: Better tooling attracts more users
- **Data insights**: Understanding how AI uses custom element data for HTML generation

### For Developers  
- **Seamless integration**: Custom elements become "native" to AI workflows
- **Reduced context switching**: AI understands project's component library
- **Better code generation**: AI suggests correct attributes, slots, imports, and CSS
- **Documentation automation**: AI can generate usage examples with proper HTML structure
- **Design system compliance**: Automatic adherence to component guidelines

### For AI Systems
- **Structured knowledge**: Well-defined schemas vs unstructured docs
- **Live updates**: Always current information via file watching
- **Actionable data**: Can validate and suggest, not just inform
- **Domain expertise**: Deep understanding of web components ecosystem
- **HTML generation intelligence**: Understands proper element usage patterns

## HTML Generation Intelligence Features

### 1. **Slot Usage Intelligence**
- Analyze manifest slot definitions to suggest proper content placement
- Understand slot inheritance and nesting patterns
- Provide examples of correct slot content based on element descriptions
- Validate slot usage against manifest definitions

### 2. **Attribute Intelligence**
- Type-aware attribute value suggestions (union types, enums, booleans)
- Default value handling and recommendation
- Required vs optional attribute guidance
- Proper attribute naming (kebab-case, data attributes)

### 3. **CSS Integration Intelligence**
- CSS custom property discovery and usage guidance
- CSS parts styling recommendations
- CSS custom states integration
- Design token compatibility and usage

### 4. **Guideline Compliance**
- Parse element and attribute descriptions for usage guidelines
- Enforce design system rules through intelligent suggestions
- Provide accessibility guidance based on manifest metadata
- Suggest semantic HTML patterns based on element purpose

## Implementation Status

The CEM MCP server implementation is complete with the following features:
- ✅ Core MCP server with resource delivery
- ✅ Tool integration and workspace context sharing
- ✅ HTML generation tools and validation
- ✅ CSS integration and guideline enforcement
- ✅ Template-driven HTML generation intelligence
- ✅ Cross-package ecosystem intelligence and live development features

## Success Metrics

### Technical Metrics
- Manifest discovery speed and accuracy
- HTML generation correctness (slots, attributes, CSS)
- Guideline compliance percentage
- Real-time update latency

### User Experience Metrics  
- Reduction in HTML validation errors
- Improved CSS integration consistency
- Faster component adoption rates
- Developer satisfaction with AI-generated HTML

## Future Enhancements

### Advanced HTML Generation
- **Component composition patterns**: Intelligent nesting and layout suggestions
- **Accessibility enhancement**: Automatic ARIA attribute suggestions
- **Performance optimization**: Lazy loading and bundling guidance
- **Framework integration**: Lit, React, Vue-specific HTML patterns

### Ecosystem Intelligence
- **Design system evolution**: Track component changes and migration paths
- **Usage analytics**: Popular attribute combinations and slot patterns
- **Community patterns**: Crowdsourced best practices and examples
- **Integration testing**: Validate HTML across different environments

This implementation has transformed CEM into an **AI-native component intelligence platform for intelligent HTML generation** - positioning the project at the forefront of AI-assisted web development with proper component usage, styling, and guideline compliance.

## Accessibility Validation Capabilities

The MCP server provides comprehensive accessibility validation through its resource and tool systems:

### Current Accessibility Features

#### Custom Element Content Analysis
- **Slot content guidelines**: Parse slot descriptions and validate slotted content against manifest guidelines
- **Attribute combination validation**: Detect contradictory attribute combinations based on manifest constraints
- **Content/attribute redundancy**: Identify when slot content overrides or conflicts with attribute values
- **Manifest compliance**: Ensure custom elements are used according to their documented requirements

#### Custom Element Accessibility Validation  
- **Component role validation**: Verify custom elements have appropriate implicit or explicit ARIA roles
- **Slot accessibility**: Ensure slotted content maintains accessibility semantics and proper ARIA relationships
- **Component state validation**: Validate accessibility states specific to custom element behavior  
- **Shadow DOM accessibility**: Check that CSS parts and custom properties don't break accessibility
- **Manifest accessibility patterns**: Apply accessibility guidelines from manifest descriptions and documentation


### Integration with Existing Tools
- Leverages existing LSP infrastructure for real-time validation
- Integrates with MCP tools for AI-assisted accessibility improvements
- Uses interface-based registry for type-safe accessibility metadata
- Builds on manifest data for component-specific accessibility patterns

### Success Metrics
- **WCAG AA compliance**: Automated detection of WCAG AA violations
- **Screen reader compatibility**: Components work correctly with major screen readers
- **Keyboard accessibility**: All functionality available via keyboard
- **AI assistance quality**: AI suggestions improve accessibility compliance rates
- **Developer experience**: Accessibility validation integrated seamlessly into development workflow

## Current Implementation: Declarative Architecture Details

### How It Works: Technical Implementation

The MCP server uses a **declarative tool framework** with a **Data + Context + LLM Decision Making** approach:

#### Declarative Architecture

The server uses a data-driven architecture where tools are defined through YAML configuration rather than hardcoded logic:

1. **📊 Data Fetchers** - Extract specific manifest data using JSON path queries
2. **📋 Template Rendering** - Combine data with Go templates for rich, contextual responses
3. **📖 Schema Integration** - Provide JSON schema context for AI understanding
4. **🧠 Intelligent Context** - Present comprehensive information while letting AI make smart decisions

#### Core Philosophy

1. **📊 Manifest Data** - Your component definitions, CSS APIs, constraints
2. **📖 Rich Context** - Usage patterns, guidelines, best practices, schema definitions
3. **🧠 AI Decisions** - AI chooses appropriate values based on provided context

This means no hardcoded suggestions - just rich context that enables intelligent decisions about component usage and styling.

#### Benefits of the Declarative Approach

- **🔧 Drop-in Resource Support** - New resources can be added with just YAML configuration
- **📝 Template-Driven Content** - Rich responses generated from your manifest data
- **⚡ Reduced Maintenance** - 77% less Go code to maintain
- **🎯 Focused Resources** - Specialized resources for attributes, slots, events, CSS APIs
- **🔄 Consistent Output** - All resources use the same data processing pipeline
- **🚀 RESTful Architecture** - Resources follow web standards for caching and access
- **📊 Granular Access** - Optional sub-resource paths for specific items (e.g., `/attributes/{name}`)

### Architecture: Resources as Database

The MCP server follows a **"manifests as database, resources as API"** principle:

#### 📊 Manifest Database
Your custom elements manifests serve as a structured database of component information, providing:
- Type definitions and constraints
- Usage guidelines and patterns
- Design system integration rules
- Accessibility requirements

#### 🔗 Declarative Resources
Declarative resources provide comprehensive access to this database through RESTful endpoints:

**Element Information:**
- `cem://element/{tagName}/attributes` - Attribute documentation and constraints
- `cem://element/{tagName}/slots` - Content guidelines and accessibility
- `cem://element/{tagName}/events` - Event patterns and integration

**CSS Styling:**
- `cem://element/{tagName}/css/parts` - CSS parts for targeted styling
- `cem://element/{tagName}/css/custom-properties` - Theming and customization
- `cem://element/{tagName}/css/states` - Interactive state styling

#### ⚡ Action Tools
Core tools perform concrete actions using the resource data:
- **`generate_html`** - Creates proper HTML structure
- **`validate_html`** - Ensures manifest compliance

This architecture separates **information access** (resources) from **actions** (tools), providing better performance, caching, and a cleaner API surface.
