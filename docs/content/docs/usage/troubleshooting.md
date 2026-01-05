---
title: Troubleshooting
weight: 110
---

Common issues and solutions when working with CEM.

## Manifest is empty

Make sure your glob pattern matches your files:
```sh
cem generate --verbose
```

Check the `.config/cem.yaml` `generate.files` glob pattern matches your source files.

## Component not showing in dev server

1. **Regenerate manifest**: `cem generate`
2. **Check file patterns**: Verify `generate.files` glob in `.config/cem.yaml`
3. **Verify package.json**: Ensure `customElements` field exists
4. **Verify manifest content**: `cem list`
5. **Check demos**: Ensure demo files exist and are discoverable
6. **Restart the server**: `cem serve`

## Demos not discovered

Check demo file location matches discovery patterns:
- Files in `demo/` or `demos/` directories
- Files named `*.demo.html`
- Or configured in `.config/cem.yaml`

See [Working with Demos](../demos/) for details on demo discovery configuration.

## TypeScript errors

Ensure you have a `tsconfig.json`:
```json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "ES2022",
    "moduleResolution": "bundler",
    "experimentalDecorators": true
  }
}
```

## LSP features not working

1. **Check manifest exists**: `ls custom-elements.json`
2. **Verify manifest content**: `cem list`
3. **Verify LSP is running**: Check editor's LSP status
4. **Regenerate if stale**: `cem generate`
5. **Check LSP logs** in your editor
6. **Restart editor**: Reload window

See [Using LSP Features](../using-lsp/) for more details.

## Dev server won't start

**Check for port conflicts**:
```sh
# Try a different port
cem serve --port 3000
```

**Check for syntax errors**:
```sh
# Validate your manifest
cem validate
```

## Import maps not working

1. **Check package.json**: Ensure dependencies are listed
2. **Restart server**: `cem serve` picks up package.json on startup
3. **Clear browser cache**: Hard reload (Ctrl+Shift+R / Cmd+Shift+R)

See [Import Maps](../import-maps/) for configuration details.

## Changes not refreshing

1. **Check live reload** is enabled (default)
2. **Look for errors** in browser console
3. **Verify source maps** for TypeScript
4. **Clear browser cache** if needed

## Get Help

- **[Documentation](/docs/)** - Browse all docs
- **[GitHub Issues](https://github.com/bennypowers/cem/issues)** - Report bugs or request features
- **[Examples](https://github.com/bennypowers/cem/tree/main/examples)** - See working examples
