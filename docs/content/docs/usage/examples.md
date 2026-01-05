---
title: Example Projects
weight: 30
---

The `cem` repository includes several example projects that demonstrate different approaches and complexity levels. Each example is a complete, working project you can explore, copy, and adapt to your needs. Don't overthink which one to start with—pick one that sounds interesting, run `npm install && npm run serve`, and start experimenting.

## Available Examples

**[minimal][minimal]** — The simplest possible custom element with one property, one slot, and basic JSDoc. Perfect for your first `cem` project or when you just want to see how the manifest workflow works. If you're brand new to `cem`, start here.

**[intermediate][intermediate]** — Multiple components that work together, showing realistic patterns like component composition and custom events. Use this as a template when building actual component libraries.

**[kitchen-sink][kitchensink]** — A production-ready button component demonstrating every `cem` feature: properties, slots, events, CSS parts, CSS custom properties, design tokens, form integration, and comprehensive documentation. Reference this when you need to see how a specific feature is documented.

**[ai-friendly-docs][aifriendlydocs]** — Comprehensive JSDoc documentation patterns for AI-friendly component descriptions, demonstrating RFC 2119 keywords, detailed event documentation, slot descriptions, and CSS property guidance. Use this as a reference for writing effective component documentation that works well with AI assistants.

**[vanilla][vanilla]** — Pure Web Components without Lit or any framework, using only native browser APIs. Choose this if you prefer framework-free development or want to understand how `cem` works with vanilla JavaScript.

**[typescript-paths][typescriptpaths]** — Demonstrates TypeScript path aliases and complex compiler configurations. Use this if your project has custom import resolution or monorepo setup.

## Quick Start

Clone the repository and run any example:

```shell
git clone https://github.com/bennypowers/cem.git
cd cem/examples/minimal  # or any other example
npm install
npm run serve            # Start dev server
```

Each example includes its own README with details about what it demonstrates and how to use it.

## Which One Should I Use?

If you're learning `cem`, start with **minimal**. If you're building something real, use **intermediate** as your template. If you need to see how a specific feature works, check **kitchen-sink**. If you want to write AI-friendly documentation, see **ai-friendly-docs**. If you prefer vanilla JavaScript, try **vanilla**. If you have complex TypeScript setup, see **typescript-paths**.

The examples are meant to be copied and modified. Take what works, remove what doesn't, and adapt them to your project's needs.

## See Also

- **[Getting Started][gettingstarted]** - Complete walkthrough for your first project
- **[Development Workflow][developmentworkflow]** - Understanding the write-generate-serve-test cycle
- **[Configuration Reference][configurationreference]** - All config options explained

[minimal]: https://github.com/bennypowers/cem/tree/main/examples/minimal
[intermediate]: https://github.com/bennypowers/cem/tree/main/examples/intermediate
[kitchensink]: https://github.com/bennypowers/cem/tree/main/examples/kitchen-sink
[aifriendlydocs]: https://github.com/bennypowers/cem/tree/main/examples/ai-friendly-docs
[vanilla]: https://github.com/bennypowers/cem/tree/main/examples/vanilla
[typescriptpaths]: https://github.com/bennypowers/cem/tree/main/examples/typescript-paths
[gettingstarted]: ../getting-started/
[developmentworkflow]: ../workflow/
[configurationreference]: /docs/reference/configuration/
