# Example Projects for `cem serve`

**Status**: 📋 **PLANNED** (Post-Release)

This document outlines example projects that demonstrate `cem serve` features. These will be created after initial release to help users understand and adopt the dev server.

## Goals

- Demonstrate key features in realistic scenarios
- Provide starting points for common use cases
- Show best practices for manifest-driven development
- Give users reference implementations

## Example Project Ideas

### 1. Simple Button Component

**Scope**: Single element, basic features
**Features demonstrated**:
- Basic component with LitElement
- JSDoc documentation for knobs
- Single demo file
- TypeScript usage

**Structure**:
```
simple-button/
├── package.json
├── src/
│   └── my-button.ts
├── demo/
│   └── index.html
└── custom-elements.json
```

**Target users**: Beginners, "hello world" example

---

### 2. Card Component with Variants

**Scope**: Single element, multiple demos
**Features demonstrated**:
- Enum attributes (variant types)
- Multiple demo files
- Slot usage
- CSS custom properties for theming

**Structure**:
```
card-component/
├── package.json
├── src/
│   └── my-card.ts
├── demo/
│   ├── basic.html
│   ├── variants.html
│   └── custom-styling.html
└── custom-elements.json
```

**Target users**: Intermediate users learning demos and knobs

---

### 3. Form Components (Compositional)

**Scope**: Multiple related elements
**Features demonstrated**:
- Multiple elements (input, select, checkbox, form)
- Element composition
- Form-associated custom elements
- Multi-instance knobs

**Structure**:
```
form-components/
├── package.json
├── src/
│   ├── my-input.ts
│   ├── my-select.ts
│   ├── my-checkbox.ts
│   └── my-form.ts
├── demo/
│   ├── individual-controls.html
│   ├── complete-form.html
│   └── validation.html
└── custom-elements.json
```

**Target users**: Advanced users, compositional patterns

---

### 4. Design System Workspace

**Scope**: Monorepo with multiple packages
**Features demonstrated**:
- npm workspaces
- Cross-package imports
- Shared configuration
- Import map generation
- Multiple packages with dependencies

**Structure**:
```
design-system/
├── package.json          # Workspace root
├── .config/
│   └── cem.yaml         # Shared config
├── packages/
│   ├── button/
│   │   ├── package.json
│   │   ├── src/my-button.ts
│   │   └── demo/index.html
│   ├── card/
│   │   ├── package.json
│   │   ├── src/my-card.ts
│   │   └── demo/index.html  # Uses button from @ds/button
│   └── form/
│       ├── package.json
│       ├── src/my-form.ts
│       └── demo/index.html  # Uses button and card
```

**Target users**: Teams, design system developers

---

### 5. TypeScript Configuration Example

**Scope**: Advanced TypeScript usage
**Features demonstrated**:
- tsconfig.json configuration
- Path mappings
- Transform target configuration
- Decorators usage

**Structure**:
```
typescript-example/
├── package.json
├── tsconfig.json
├── .config/
│   └── cem.yaml
├── src/
│   ├── components/
│   │   └── my-element.ts
│   └── utils/
│       └── helpers.ts
└── demo/
    └── index.html
```

**Target users**: TypeScript users, advanced configuration

---

### 6. Themed Components

**Scope**: Components with extensive CSS theming
**Features demonstrated**:
- CSS custom properties
- Color knobs
- Design tokens integration
- Multiple themes

**Structure**:
```
themed-components/
├── package.json
├── src/
│   ├── my-themed-button.ts
│   └── themes/
│       ├── light.css
│       └── dark.css
├── demo/
│   ├── light-theme.html
│   ├── dark-theme.html
│   └── custom-theme.html
└── custom-elements.json
```

**Target users**: Design-focused developers, theming

---

## Implementation Plan

### Phase 1: Create Examples (Post-Release)
1. **Simple Button** - Start here, validate approach
2. **Card Component** - Add complexity (variants, slots)
3. **Form Components** - Show composition

**Timeline**: 1-2 weeks after release

### Phase 2: Advanced Examples
4. **Design System Workspace** - Demonstrate workspace mode
5. **TypeScript Configuration** - Show advanced config
6. **Themed Components** - Focus on styling

**Timeline**: 3-4 weeks after release

## Hosting Options

### Option A: Separate Repository
- Repo: `bennypowers/cem-serve-examples`
- Benefits: Easy to discover, can accept contributions, versioned
- Drawbacks: Separate from main repo

### Option B: Monorepo in Main Repo
- Location: `examples/` directory in main cem repo
- Benefits: Co-located with code, easier to keep in sync
- Drawbacks: Bloats main repo

### Option C: Documentation Examples
- Location: `docs/examples/serve/`
- Benefits: Part of docs site, discoverable
- Drawbacks: Limited by Hugo/docs tooling

**Recommendation**: Start with Option A (separate repo), link from docs.

## Example Quality Standards

Each example must have:
- [ ] Clear README.md explaining purpose and features
- [ ] Complete package.json with all dependencies
- [ ] Proper JSDoc documentation
- [ ] Multiple demos showing different aspects
- [ ] Generated custom-elements.json checked in
- [ ] TypeScript types (where applicable)
- [ ] License (MIT or GPL to match main project)
- [ ] Link back to main cem docs

## Documentation Integration

Once examples exist:
1. Add "Examples" section to docs/content/docs/serve/
2. Link from getting-started.md
3. Reference from commands/serve.md
4. Create docs/content/docs/serve/examples.md with:
   - List of all examples
   - What each demonstrates
   - Links to GitHub repos
   - Quick start for each

## Success Criteria

- [ ] At least 3 examples published (simple, intermediate, advanced)
- [ ] Examples linked from docs
- [ ] Each example has comprehensive README
- [ ] Examples used in tutorials/guides
- [ ] Community can contribute more examples

## Future Enhancements

- **Example templates** - CLI scaffolding (`cem init --template button`)
- **Interactive playground** - CodeSandbox/StackBlitz integrations
- **Video walkthroughs** - Screencast tutorials using examples
- **Community gallery** - Showcase user projects

## Notes

- Defer to post-release to avoid delaying main feature
- Focus on documentation completeness first
- Examples should be minimal but realistic
- Keep dependencies minimal (prefer standard web APIs)
- Show idiomatic patterns, not every possible feature
