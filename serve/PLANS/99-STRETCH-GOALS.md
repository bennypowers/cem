## Stretch Goals / Follow up Ideas

### UI
- **Manifest viewer UI** - `/cem/manifest` shows browsable CEM data
- **Component playground** - Create ad-hoc demos without files
- **Event logger** - Console showing custom events fired by elements
- **Accessibility panel** - Live a11y tree, ARIA attributes, keyboard nav testing
- **CSS custom property inspector** - Show all CSS vars, edit live
- **Slot visualizer** - Overlay that highlights available slots, show default content
- **Responsive preview** - Iframe with device presets

### Workflow
- **HMR** - inject Open-WC HMR code into ce classes
The idea would be to inject the necessary HMR code into custom element classes, rather than rely on the authors to provide it
see:
- https://open-wc.org/docs/development/hot-module-replacement/
- https://github.com/fi3ework/vite-plugin-web-components-hmr
- https://modern-web.dev/docs/dev-server/plugins/hmr/

### Logging
- **Performance metrics** - Render timing, reflow counts

### Server
1. **HTTPS support** - Modern APIs require secure context
   - Self-signed cert generation
   - `--cert`/`--key` flags
2. **Proxy configuration** - API requests need CORS workarounds
   - `/api` → `http://localhost:3000`
3. **Custom headers** - CSP, Feature-Policy, etc.
4. **Request logging format** - Machine-readable option (JSON logs)

### WebSocket Lifecycle Management

1. **Client reconnection strategy**
   - Network interruption handling
   - Exponential backoff for retries
   - Max retry attempts before giving up

2. **Connection health checks**
   - Ping/pong heartbeat to detect dead connections
   - Timeout for inactive connections
   - Cleanup of zombie connections

3. **Server-side limits**
   - Max concurrent WebSocket connections
   - Memory per connection
   - Broadcast queue depth

**WebSocket lifecycle specification:**
```typescript
// Client side
class ReloadClient {
  #reconnectAttempts = 0;
  #maxRetries = 10;
  #backoffMs = 1000;

  connect() {
    // Exponential backoff reconnection
  }
}
```

