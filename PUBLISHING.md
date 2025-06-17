# Multi-Platform Publishing for cem

cem uses an esbuild-style optionalDependencies strategy for platform binaries.

## Releasing

1. Tag your release (e.g., `v1.2.3`) in the main branch.
2. The CI workflow will:
   - Cross-compile and publish each `@pwrs/cem-PLATFORM-ARCH` package with the correct binary.
   - Publish the main `@pwrs/cem` package with all subpackages as `optionalDependencies`.

## Local Testing

- Use `npm install --ignore-scripts` to test installation.
- The wrapper (`bin/cem.js`) will find and invoke the platform-specific binary.

## Requirements

- Node.js 22+ is required for all packages.
- All packages are ESM-only (`"type": "module"`).

## Adding a New Platform

- Add a new entry in `scripts/gen-platform-package-jsons.mjs` and build logic.
- Update the main package's `optionalDependencies`.

## How It Works

- The main package (`@pwrs/cem`) does not ship a binary.
- On install, npm will only pull in the appropriate platform package.
- The wrapper script detects your platform and runs the correct binary.
