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
mcp/tools/templates/
├── css_properties.md      # CSS custom properties guidance
├── css_parts.md           # CSS parts styling patterns  
├── css_states.md          # CSS custom states usage
├── theming_guidance.md    # Theme-aware styling advice
├── basic_styling.md       # Fallback for elements without CSS APIs
└── responsive_guidance.md # Responsive design patterns
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

## Vision: Intelligent HTML Generation Platform

**End Goal**: Create an AI-native component intelligence platform that enables intelligent HTML generation with:
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

## Implementation Plan

### Phase 1: Core MCP Server (Weeks 1-2)
```
cmd/mcp.go              # New MCP command following existing patterns
mcp/
├── server.go           # MCP server implementation
├── resources.go        # Resource providers (manifests, schemas)
├── tools.go           # Interactive tools for HTML generation
├── context.go         # Shared with LSP via interface
├── html_generator.go  # Intelligent HTML generation logic
└── css_integration.go # CSS custom properties and parts handling
```

**Key Resources**:
- `cem://schema` - JSON schema for custom elements manifests
- `cem://registry` - Current manifest registry state with CSS metadata
- `cem://element/{tagName}` - Individual element definitions with usage patterns
- `cem://package/{packageName}` - Package-specific manifests
- `cem://css/{tagName}` - CSS integration guide (custom properties, parts, states)
- `cem://guidelines` - Design system guidelines and component usage rules

### Phase 2: Template-Powered Tools (Weeks 3-4)
**Interactive MCP Tools Using Template-Driven Guidance**:
- `validate_element` - Validate using **manifest constraints + contextual feedback**
- `suggest_attributes` - Present **syntax definitions + usage guidance** for LLM decisions
- `resolve_import` - Find correct import path for element
- `query_registry` - Search elements by criteria
- `generate_html` - **Template-driven HTML structure** respecting manifest slot/attribute definitions
- `validate_css_usage` - **Template-based validation** using property syntax and constraints
- `suggest_css_integration` - **Template system** presenting CSS APIs with rich context for LLM

**Template Philosophy Applied**:
- **No hardcoded suggestions** → Rich manifest data + contextual templates
- **LLM-driven decisions** → AI chooses appropriate values based on provided constraints
- **Extensible guidance** → New templates = new capabilities without code changes

### Phase 3: Intelligent HTML Generation Features (Week 5+)
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

## Implementation Timeline

**Week 1**: Core MCP server with basic resource delivery
**Week 2**: Tool integration and workspace context sharing
**Week 3**: HTML generation tools and validation
**Week 4**: CSS integration and guideline enforcement
**Week 5**: Advanced HTML generation intelligence
**Week 6+**: Cross-package ecosystem intelligence and live development features

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

This implementation transforms CEM from a development tool into an **AI-native component intelligence platform for intelligent HTML generation** - a significant evolution that positions your project at the forefront of AI-assisted web development with proper component usage, styling, and guideline compliance.

## Comprehensive Accessibility Validation TODOs

### Phase 1: Core Accessibility Validation Infrastructure

#### HTML Structure Analysis
- [ ] **Semantic HTML validation**: Parse HTML and verify proper semantic element usage (main, nav, section, article, aside, header, footer)
- [ ] **Heading hierarchy validation**: Ensure proper h1-h6 nesting and logical document outline
- [ ] **Landmark structure verification**: Check for appropriate landmark roles and ARIA regions
- [ ] **List structure validation**: Verify proper ul/ol/li nesting and semantic list usage

#### ARIA Implementation Validation  
- [ ] **ARIA role validation**: Verify roles are appropriate for element types and contexts
- [ ] **ARIA property validation**: Check aria-label, aria-describedby, aria-labelledby usage and relationships
- [ ] **ARIA state validation**: Validate aria-expanded, aria-checked, aria-selected states
- [ ] **ARIA relationship validation**: Ensure aria-controls, aria-owns relationships point to valid elements
- [ ] **Required ARIA attributes**: Check for required ARIA attributes based on roles (e.g., aria-valuemin/max for sliders)

#### Keyboard Navigation Validation
- [ ] **Tabindex validation**: Ensure proper tabindex usage (0 for focusable, -1 for programmatic focus only)
- [ ] **Focus management validation**: Verify focus trapping in modals and focus restoration patterns
- [ ] **Keyboard event validation**: Check for proper Enter/Space key handling on custom interactive elements
- [ ] **Skip link validation**: Verify skip links are present and functional for navigation-heavy pages

### Phase 2: Advanced Accessibility Compliance

#### Color and Visual Accessibility
- [ ] **Color contrast validation**: Implement WCAG AA contrast ratio checking (4.5:1 normal, 3:1 large text)
- [ ] **Color independence validation**: Ensure information isn't conveyed through color alone
- [ ] **Focus indicator validation**: Verify visible focus indicators meet contrast requirements
- [ ] **CSS custom property accessibility**: Validate that custom properties maintain accessibility when overridden

#### Form Accessibility Validation
- [ ] **Form labeling validation**: Ensure all inputs have associated labels (explicit or implicit)
- [ ] **Required field indication**: Verify required fields are properly marked and announced
- [ ] **Error message association**: Check aria-describedby relationships for validation messages
- [ ] **Fieldset and legend validation**: Ensure proper grouping for related form controls
- [ ] **Input type validation**: Verify appropriate input types for better mobile/assistive technology support

#### Dynamic Content Accessibility
- [ ] **Live region validation**: Check for proper aria-live usage for dynamic content updates
- [ ] **Screen reader announcement validation**: Ensure status changes are properly announced
- [ ] **Loading state accessibility**: Verify loading indicators are accessible to screen readers
- [ ] **Progressive enhancement validation**: Check that functionality works without JavaScript

### Phase 3: Custom Element Specific Accessibility

#### Component-Level Accessibility
- [ ] **Custom element role validation**: Ensure custom elements have appropriate implicit or explicit roles
- [ ] **Slot content accessibility**: Validate that slotted content maintains semantic meaning
- [ ] **CSS Shadow DOM accessibility**: Verify CSS parts and custom properties don't break accessibility
- [ ] **Event accessibility**: Ensure custom events provide appropriate accessibility information

#### Design System Compliance
- [ ] **Consistent interaction patterns**: Validate that similar elements behave consistently
- [ ] **Accessibility pattern compliance**: Check adherence to established accessibility patterns (dialogs, menus, etc.)
- [ ] **Component accessibility documentation**: Verify elements have proper accessibility usage guidelines
- [ ] **Cross-component accessibility**: Ensure components work together accessibly

### Phase 4: Testing and Validation Tools Integration

#### Automated Testing Integration
- [ ] **axe-core integration**: Integrate axe-core accessibility testing engine
- [ ] **Pa11y integration**: Add command-line accessibility testing capabilities  
- [ ] **Lighthouse accessibility integration**: Include Lighthouse accessibility audit results
- [ ] **Custom accessibility rules**: Develop custom element specific accessibility rules

#### Manual Testing Guidance
- [ ] **Screen reader testing patterns**: Provide specific testing patterns for different screen readers
- [ ] **Keyboard testing procedures**: Generate comprehensive keyboard testing procedures
- [ ] **User testing integration**: Framework for accessibility user testing with disabled users
- [ ] **Accessibility reporting**: Generate comprehensive accessibility compliance reports

### Phase 5: Real-time Accessibility Assistance

#### AI-Powered Accessibility Suggestions
- [ ] **Context-aware accessibility suggestions**: Provide accessibility improvements based on element context
- [ ] **Alternative text generation**: AI-assisted alt text suggestions for images and icons
- [ ] **ARIA label suggestions**: Context-aware aria-label and description suggestions
- [ ] **Accessibility anti-pattern detection**: Identify and suggest fixes for common accessibility mistakes

#### Accessibility-First HTML Generation
- [ ] **Semantic element selection**: Choose most appropriate semantic elements for content type
- [ ] **Automatic ARIA attribute addition**: Add required ARIA attributes based on element roles and context
- [ ] **Accessibility-compliant component composition**: Ensure generated component combinations are accessible
- [ ] **Progressive enhancement generation**: Generate accessible fallback content for custom elements

### Implementation Strategy

#### Phase 1 Priority (Core Infrastructure)
1. HTML structure and ARIA validation (highest impact, foundational)
2. Keyboard navigation validation (critical for usability)
3. Color contrast validation (WCAG compliance requirement)

#### Phase 2 Priority (Form and Dynamic Content)
1. Form accessibility validation (high user impact)
2. Dynamic content accessibility (modern web requirement)
3. Focus management validation (complex but essential)

#### Phase 3 Priority (Custom Element Specific)
1. Component-level accessibility validation
2. Design system compliance checking
3. Cross-component accessibility validation

#### Integration with Existing Tools
- Leverage existing LSP infrastructure for real-time validation
- Integrate with MCP tools for AI-assisted accessibility improvements
- Use interface-based registry for type-safe accessibility metadata
- Build on manifest data for component-specific accessibility patterns

### Success Metrics
- **WCAG AA compliance**: Automated detection of WCAG AA violations
- **Screen reader compatibility**: Components work correctly with major screen readers
- **Keyboard accessibility**: All functionality available via keyboard
- **AI assistance quality**: AI suggestions improve accessibility compliance rates
- **Developer experience**: Accessibility validation integrated seamlessly into development workflow
