# TypeScript Paths Example

Demonstrates how to use CEM with TypeScript path mappings and build outputs.

## Overview

This example shows a common TypeScript project structure:

- Source files in `src/`
- Compiled output in `dist/`
- TypeScript path mappings (`@components/*`, `@utils/*`)
- Dev server URL rewrites for development workflow

Perfect for:

- Projects using TypeScript compilation
- Build pipelines with src → dist flow
- Path-mapped imports
- Understanding dev server configuration

## Project Structure

```text
typescript-paths/
├── .config/
│   └── cem.yaml              # With URL rewrite config
├── src/
│   └── components/
│       └── ts-button.ts      # Source component
├── dist/                     # Compiled output (generated)
├── package.json
├── tsconfig.json             # With paths configuration
└── README.md
```

## Key Features

### 1. TypeScript Path Mappings

`tsconfig.json` includes path mappings:

```json
{
  "compilerOptions": {
    "paths": {
      "@components/*": ["./src/components/*"],
      "@utils/*": ["./src/utils/*"]
    }
  }
}
```

This allows imports like:

```typescript
import { TsButton } from '@components/ts-button.js';
import { capitalize, kebabCase } from '@utils/format.js';
```

**Real Example** (from `src/components/ts-button.ts`):

```typescript
import { capitalize } from '@utils/format.js';

// Use the utility in the component
const displayLabel = this.label ? capitalize(this.label) : '';
```

### 2. Separate Source and Output

- **Source**: `src/` directory
- **Output**: `dist/` directory (after `npm run build`)

CEM analyzes the source files in `src/`, not the compiled output.

### 3. Dev Server URL Rewrites

The dev server is configured to rewrite URLs from `/dist/` to `/src/`:

```yaml
serve:
  urlRewrites:
    - from: /dist/(.*)
      to: /src/$1
```

This allows:
- HTML files to reference `/dist/components/ts-button.js` (production path)
- Dev server to serve from `/src/components/ts-button.ts` (source)
- No build step needed during development

## Running the Example

### Install Dependencies

From the repository root:

```bash
npm install  # or pnpm install
```

### Generate the Manifest

```bash
cd examples/typescript-paths
npm run analyze
```

### Build TypeScript

```bash
npm run build
```

This compiles `src/**/*.ts` → `dist/**/*.js`

### Start Dev Server

```bash
npm run serve
```

The dev server will:
1. Serve files from the example directory
2. Rewrite `/dist/*` requests to `/src/*`
3. Allow development without building

## Using the Component

### In HTML (Development)

```html
<!DOCTYPE html>
<html>
<head>
  <!-- Dev server rewrites /dist/ to /src/ -->
  <script type="module" src="/dist/components/ts-button.js"></script>
</head>
<body>
  <ts-button>Click Me</ts-button>
  <ts-button disabled>Disabled</ts-button>
  <ts-button label="With Label Prop"></ts-button>
</body>
</html>
```

### In HTML (Production)

After building, the same HTML works in production because the files exist in `dist/`.

### With Path Mappings

Components and utilities use clean import paths:

```typescript
// src/index.ts - Entry point demonstrating path mappings
import { TsButton } from '@components/ts-button.js';
import { capitalize, kebabCase, truncate } from '@utils/format.js';

// TypeScript resolves:
// @components/* → src/components/*
// @utils/* → src/utils/*
```

**Available Utilities** (`src/utils/format.ts`):

- `capitalize(str)` - Capitalizes first letter
- `kebabCase(str)` - Converts to kebab-case
- `truncate(str, maxLength)` - Truncates with ellipsis

## Dev Server URL Rewrites Explained

### The Problem

- Production code imports from `/dist/components/button.js`
- During development, source is in `/src/components/button.ts`
- Don't want to build on every change

### The Solution

URL rewrites in `cem.yaml`:

```yaml
serve:
  urlRewrites:
    - from: /dist/(.*)
      to: /src/$1
    # TypeScript path mapping aliases
    - from: /@components/(.*)
      to: /src/components/$1
    - from: /@utils/(.*)
      to: /src/utils/$1
```

**How URL rewrites work:**

1. **Production paths** (`/dist/*`):
   - Browser requests: `/dist/components/ts-button.js`
   - Dev server rewrites to: `/src/components/ts-button.ts`
   - Serves the TypeScript source directly

2. **Path mapping aliases** (`@components/*`, `@utils/*`):
   - Browser requests: `/@components/ts-button.js`
   - Dev server rewrites to: `/src/components/ts-button.ts`
   - Allows using TypeScript aliases in development

Benefits:
- Edit source, see changes immediately
- No build step during development
- HTML files use production paths
- Smooth transition to production

## TypeScript Configuration

### Key Options

```json
{
  "compilerOptions": {
    "outDir": "./dist",           // Compiled output location
    "rootDir": "./src",            // Source root
    "paths": {                     // Path mappings
      "@components/*": ["./src/components/*"]
    },
    "declaration": true,           // Generate .d.ts files
    "sourceMap": true              // Generate .map files
  },
  "include": ["src/**/*"],         // Files to compile
  "exclude": ["node_modules", "dist"]
}
```

### Building

```bash
npm run build
```

Generates:
- `dist/components/ts-button.js` - Compiled JavaScript
- `dist/components/ts-button.d.ts` - Type definitions
- `dist/components/ts-button.js.map` - Source maps

## What's in the Manifest?

The generated `custom-elements.json` includes:

- Component metadata from source (`src/`)
- Source file paths (not dist paths)
- All JSDoc from TypeScript source
- Type information from TS types

CEM analyzes the TypeScript source directly, so you get:
- Accurate type information
- Original JSDoc comments
- Source code locations

## Comparison with Other Examples

| Feature | This Example | Others |
|---------|--------------|--------|
| Source directory | ✅ `src/` | ❌ `elements/` |
| Build output | ✅ `dist/` | ❌ None |
| Path mappings | ✅ `@components/*` | ❌ None |
| URL rewrites | ✅ Configured | ❌ Not needed |
| Build step | ✅ `npm run build` | ❌ Not needed |

Use this pattern when:
- Using TypeScript compilation
- Want separate source/output dirs
- Need path mappings
- Have a build pipeline

Use the elements/ pattern when:
- No build step
- Simpler project structure
- Direct file serving

## Workflow

### Development

1. Edit `src/components/ts-button.ts`
2. Run `npm run serve`
3. Open HTML file in browser
4. Changes appear immediately (no build needed)

### Production

1. Edit source files
2. Run `npm run build`
3. Deploy `dist/` directory
4. HTML files work unchanged

### Continuous Development

The typical workflow:

```bash
# Terminal 1: Dev server (with URL rewrites)
npm run serve

# Terminal 2: Watch mode (optional, for types)
tsc --watch

# Edit src/components/*.ts
# See changes immediately in browser
# TypeScript types update automatically
```

## Next Steps

Compare with other patterns:

- **[minimal](../minimal/)** - Simplest structure
- **[intermediate](../intermediate/)** - Multiple components
- **[kitchen-sink](../kitchen-sink/)** - All features
- **[vanilla](../vanilla/)** - No framework

## Learn More

- [TypeScript Path Mapping](https://www.typescriptlang.org/docs/handbook/module-resolution.html#path-mapping)
- [TypeScript Compiler Options](https://www.typescriptlang.org/tsconfig)
- [Module Resolution](https://www.typescriptlang.org/docs/handbook/module-resolution.html)
