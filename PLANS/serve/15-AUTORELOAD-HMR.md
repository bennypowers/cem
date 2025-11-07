## Auto-Reload and HMR

### Auto-Reload
- Watch source files (same as generate watch mode)
- On change:
  1. Regenerate manifest (in-memory)
  2. Clear transform cache for changed files
  3. Broadcast `reload` event via WebSocket
  4. Browser reloads page

**Smart reloading:**
- CSS-only changes → inject new styles (no page reload)
- Demo HTML changes → reload demo iframe
- Element changes → full page reload

### HMR (STRETCH)
The idea would be to inject the necessary HMR code into custom element classes, rather than rely on the authors to provide it
see:
- https://open-wc.org/docs/development/hot-module-replacement/
- https://github.com/fi3ework/vite-plugin-web-components-hmr
- https://modern-web.dev/docs/dev-server/plugins/hmr/

### Config
Under `.serve.reload`
```yaml
reload: auto
# or
reload: hmr
hmr:
  inject: true # false when author already provided hooks
```
