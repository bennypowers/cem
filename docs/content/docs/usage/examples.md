---
title: Examples Overview
weight: 30
---

The CEM repository includes several example projects to help you get started. Each demonstrates different patterns and complexity levels.

## Decision Matrix

Choose the right example based on your needs:

| Example | Best For | Components | Complexity | Features |
|---------|----------|------------|------------|----------|
| **[Minimal](#minimal)** | First-time users, learning basics | 1 simple | ⭐ | Basic JSDoc, single property, slot, CSS part |
| **[Intermediate](#intermediate)** | Realistic projects, multiple components | 3 related | ⭐⭐ | Component composition, events, demo discovery |
| **[Kitchen Sink](#kitchen-sink)** | Reference implementation, all features | 1 comprehensive | ⭐⭐⭐⭐ | Forms, design tokens, all CSS APIs, microdata |
| **[Vanilla](#vanilla)** | Framework-free approach | 1 simple | ⭐⭐ | No Lit, pure Web Components API |
| **[TypeScript Paths](#typescript-paths)** | Advanced build setups | 1 simple | ⭐⭐⭐ | Path aliases, complex tsconfig |

## Minimal

**Location**: `examples/minimal/`
**Best for**: First-time CEM users

The simplest possible custom element demonstrating the absolute minimum needed.

### What's Included

- Single `hello-world` component
- One reactive property (`name`)
- Default slot for content projection
- One CSS custom property (`--hello-world-color`)
- One CSS part (`greeting`)
- Basic JSDoc documentation

### Project Structure

```
minimal/
├── .config/
│   └── cem.yaml              # CEM configuration
├── elements/
│   └── hello-world/
│       └── hello-world.ts    # The component
├── package.json
├── tsconfig.json
└── README.md
```

### Quick Start

```bash
cd examples/minimal
npm install
npm run analyze    # Generate manifest
npm run serve      # Start dev server
```

### When to Use

- ✅ Learning CEM for the first time
- ✅ Understanding manifest basics
- ✅ Quick prototyping
- ✅ Teaching custom elements

### Next Steps From Here

After mastering minimal, explore **[Intermediate](#intermediate)** for multi-component patterns.

## Intermediate

**Location**: `examples/intermediate/`
**Best for**: Realistic multi-component projects

Demonstrates realistic patterns with multiple related components that work together.

### What's Included

- Three related components (e.g., card, card-header, card-body)
- Component composition patterns
- Custom events between components
- Multiple properties per component
- Demo file discovery
- Practical JSDoc examples

### Project Structure

```
intermediate/
├── .config/
│   └── cem.yaml
├── elements/
│   ├── my-card/
│   │   ├── my-card.ts
│   │   └── demo/
│   │       └── index.html
│   ├── my-card-header/
│   │   └── my-card-header.ts
│   └── my-card-body/
│       └── my-card-body.ts
├── package.json
└── README.md
```

### Quick Start

```bash
cd examples/intermediate
npm install
npm run analyze
npm run serve
```

### When to Use

- ✅ Building component libraries
- ✅ Learning component composition
- ✅ Understanding event patterns
- ✅ Realistic project structure

### Next Steps From Here

For advanced features, see **[Kitchen Sink](#kitchen-sink)**. For framework-free approach, try **[Vanilla](#vanilla)**.

## Kitchen Sink

**Location**: `examples/kitchen-sink/`
**Best for**: Reference implementation, exploring all CEM features

Comprehensive demonstration of EVERY Custom Elements Manifest feature in one production-ready component.

### What's Included

The `demo-button` component showcases:

- **15+ properties** - All attribute types (string, boolean, enum, union)
- **3 slots** - Default, start, end with conditional rendering
- **4+ events** - Click, focus, blur, invalid
- **5 CSS parts** - Comprehensive styling API
- **8 CSS custom properties** - Complete theming system
- **3 public methods** - Click, focus, blur
- **Design tokens** - Integration with design token spec
- **Form integration** - Submit, reset, validation
- **Link mode** - Renders as `<a>` when href provided
- **Demo discovery** - HTML microdata patterns
- **Comprehensive docs** - Production-quality JSDoc

### Project Structure

```
kitchen-sink/
├── .config/
│   └── cem.yaml              # Config with tokens & demos
├── elements/
│   └── demo-button/
│       ├── demo-button.ts    # Comprehensive component
│       └── demo/
│           ├── variants.html # All color variants
│           └── states.html   # Interactive states
├── design-tokens.json        # Design token specification
├── package.json
└── README.md
```

### Quick Start

```bash
cd examples/kitchen-sink
npm install
npm run analyze
npm run serve
```

### When to Use

- ✅ Reference for production components
- ✅ Understanding all CEM features
- ✅ Design system integration
- ✅ Form components
- ✅ Complex component APIs

### Features Comparison

| Feature | Minimal | Intermediate | **Kitchen Sink** |
|---------|---------|--------------|------------------|
| Properties | 1 | 5-7 | **15+** |
| Slots | 1 | 2-3 | **3** |
| CSS Parts | 1 | 2-4 | **5** |
| Events | 0 | 1-2 | **4** |
| Design Tokens | ❌ | ❌ | **✅** |
| Demo Discovery | ❌ | ✅ | **✅** |
| Forms | ❌ | ❌ | **✅** |
| Public Methods | 0 | 0 | **3** |

## Vanilla

**Location**: `examples/vanilla/`
**Best for**: Framework-free Web Components

Pure Web Components without Lit or any framework, using only native browser APIs.

### What's Included

- `vanilla-counter` component
- Pure JavaScript (no framework)
- Native Shadow DOM API
- Direct template manipulation
- Event handling
- Reactive updates without decorators

### Project Structure

```
vanilla/
├── .config/
│   └── cem.yaml
├── elements/
│   └── vanilla-counter/
│       ├── vanilla-counter.js
│       └── demo/
│           └── index.html
├── package.json
└── README.md
```

### Quick Start

```bash
cd examples/vanilla
npm install
npm run analyze
npm run serve
```

### When to Use

- ✅ Learning native Web Components API
- ✅ Framework-free projects
- ✅ Minimal dependencies
- ✅ Maximum control over implementation

### Next Steps From Here

Compare with **[Minimal](#minimal)** to see how Lit simplifies component development.

## TypeScript Paths

**Location**: `examples/typescript-paths/`
**Best for**: Projects with complex TypeScript configurations

Demonstrates CEM's support for TypeScript path aliases and advanced compiler configurations.

### What's Included

- TypeScript path mapping (`@/elements/*`)
- Complex tsconfig.json setup
- Import alias resolution
- Build tool integration patterns

### Project Structure

```
typescript-paths/
├── .config/
│   └── cem.yaml
├── src/
│   └── elements/
│       └── path-demo/
│           └── path-demo.ts
├── tsconfig.json             # Advanced configuration
├── package.json
└── README.md
```

### Quick Start

```bash
cd examples/typescript-paths
npm install
npm run analyze
npm run serve
```

### When to Use

- ✅ Using TypeScript path aliases
- ✅ Complex monorepo setups
- ✅ Advanced build configurations
- ✅ Custom import resolution

## Running All Examples

From the repository root:

```bash
# Install dependencies for all examples
npm install

# Run an example
cd examples/minimal
npm run analyze
npm run serve
```

Each example is self-contained with its own `package.json` and configuration.

## Common Patterns Across Examples

All examples demonstrate:

- **CEM configuration** - `.config/cem.yaml`
- **Manifest generation** - `npm run analyze`
- **Dev server** - `npm run serve`
- **TypeScript** - Modern ES modules
- **JSDoc** - Component documentation

## Choosing Your Starting Point

**Never used CEM?**
→ Start with **[Minimal](#minimal)**

**Building a component library?**
→ Use **[Intermediate](#intermediate)** as template

**Need all features documented?**
→ Reference **[Kitchen Sink](#kitchen-sink)**

**Prefer vanilla JavaScript?**
→ Try **[Vanilla](#vanilla)**

**Complex build setup?**
→ See **[TypeScript Paths](#typescript-paths)**

## See Also

- **[Getting Started](../getting-started/)** - Step-by-step first project
- **[Development Workflow](../workflow/)** - Understanding the dev cycle
- **[Working with Demos](../demos/)** - Demo organization strategies
- **[Configuration Reference](/docs/reference/configuration/)** - Complete config options
