# Dev Mode for Chrome UI Development

This document describes the `cemdev` build tag feature that enables live-reload of embedded chrome UI assets during development.

## Overview

The cem dev server embeds JS/CSS/HTML templates at compile time via `embed.FS` (see `serve/middleware/routes/templates.go`). When iterating on the chrome UI (e.g., working on components in `serve/middleware/routes/templates/elements/`), developers normally have to:

1. Kill the server
2. Run `make build`
3. Restart the server

With the `cemdev` build tag, you can enable dev mode where:

1. The server reads `templates/` files from disk instead of `embed.FS`
2. A file watcher on `serve/middleware/routes/templates/elements/**/*.ts` triggers esbuild transpilation and sends a reload signal over WebSocket

## Building with Dev Mode

To build the binary with dev mode enabled:

```bash
make dev-serve
```

This creates `dist/cem` with the `cemdev` build tag.

## Usage

After building with dev mode, start the server as normal:

```bash
dist/cem serve -p examples/kitchen-sink
```

The server will log:
```
[INFO] Dev mode: reading internal modules from disk
[INFO] Dev mode: watching elements directory for changes
```

Now when you edit any TypeScript file in `serve/middleware/routes/templates/elements/`, the server will:

1. Automatically transpile the TypeScript to JavaScript using esbuild
2. Broadcast a reload message to all connected clients
3. The browser will automatically refresh to show your changes

## Architecture

The implementation uses Go build tags to conditionally compile different versions of the code:

### Production Mode (default)

- `serve/internal_modules_prod.go` (`//go:build !cemdev`)
- Reads internal modules from `embed.FS`
- `setupDevWatcher()` is a no-op

### Dev Mode (`-tags cemdev`)

- `serve/internal_modules_dev.go` (`//go:build cemdev`)
- Reads internal modules from disk using `runtime.Caller()` to find the source directory
- Sets up a file watcher on `serve/middleware/routes/templates/elements/**/*.ts`
- On file changes:
  - Calls `elements.TranspileElements()` to run esbuild
  - Broadcasts reload message to WebSocket clients

### Shared Code

- `serve/internal/elements/transpile.go` - Reusable esbuild transpilation logic
- `serve/middleware/routes/internal_modules.go` - Exports `ReadInternalModule` as a variable that can be overridden

## Testing

The feature includes comprehensive tests:

### Transpilation Tests

```bash
go test ./serve/internal/elements/ -v
```

Tests the esbuild transpilation logic in isolation.

### Production Mode Tests

```bash
go test ./serve/ -v -run TestReadInternalModule_Production
```

Verifies that production mode reads from embed.FS.

### Dev Mode Tests

```bash
go test -tags cemdev ./serve/ -v -run "TestReadInternalModule_Dev|TestSetupDevWatcher_Dev"
```

Verifies that dev mode reads from disk and sets up the file watcher correctly.

## Limitations

- Only TypeScript files in `serve/middleware/routes/templates/elements/` are watched
- Test files (`*.test.ts`) are not transpiled
- The file watcher has a 300ms debounce delay to avoid excessive rebuilds
- Dev mode is not recommended for production use (it's slower and reads from disk)

## Troubleshooting

If the watcher isn't working:

1. Check that you built with `make dev-serve` (not `make build`)
2. Verify the server logs show "Dev mode: watching elements directory for changes"
3. Make sure you're editing `.ts` files (not `.js` files)
4. Check that the file you're editing is in a subdirectory of `serve/middleware/routes/templates/elements/`
