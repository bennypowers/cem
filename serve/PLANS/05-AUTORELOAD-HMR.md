## Auto-Reload and HMR

## TDD First Steps

1. Remove "not implemented" guard tests from existing tests for this phase
2. Write comprehensive tests for real functionality (will fail - true red phase)
3. Implement to make tests pass

### Auto-Reload
- Watch source files (same as generate watch mode)
- On change:
  1. Regenerate manifest (in-memory)
  2. Immediately clear transform cache for changed files (and their dependency tree)
  3. Broadcast `reload` event via WebSocket
  4. Browser reloads page

### HMR (STRETCH)

Out of scope for initial release. See [Stretch goals](./99-STRETCH-GOALS.md)

### Config
Under `.serve.reload`

```yaml
reload: page | hmr | true | false
# or (stretch goal)
# reload: hmr (st)
# hmr:
#   inject: true # false when author already provided hooks
```

### WebSocket Lifecycle Management

Mostly out of scope for the first draft. See [Stretch goals](./99-STRETCH-GOALS.md). We can implement some rudimentary retry polling. Keep in mind that this is a dev server which runs locally.

### Debouncing Strategy
- 150ms debounce window (balances responsiveness and spam prevention)
- Coalesce multiple file changes into single broadcast
- Include changed file count in broadcast

#### File Change Race Conditions
- Multiple files change simultaneously (e.g., git checkout)
- File changes while manifest regeneration in progress
- Broadcast triggered before manifest regeneration completes

**Example scenario:**
```
T0: file-a.ts changes → regeneration starts
T1: file-b.ts changes → should wait or cancel T0?
T2: regeneration completes → broadcast once or twice?
```

**Recommendation:**
- Batch changes during regeneration
- Single broadcast after batch completes
- Include all changed files in event payload

---

## Acceptance Criteria

- [ ] Server starts and listens on configured port
- [ ] WebSocket endpoint `/__cem/reload` accepts connections
- [ ] WebSocket client injected into all HTML responses
- [ ] File watcher detects changes to source files
- [ ] File watcher detects changes to demo files
- [ ] File watcher detects changes to config files
- [ ] File changes trigger manifest regeneration
- [ ] Manifest regeneration broadcasts `reload` event via WebSocket
- [ ] Browser receives reload event and reloads page
- [ ] Debouncing works (rapid changes result in single reload)
- [ ] 150ms debounce window implemented
- [ ] Transform cache cleared for changed files on reload
- [ ] Multiple connected clients all receive broadcast
- [ ] Logging shows file changes and reload events
- [ ] Core errors (server crashes, fatal errors) logged to console
- [ ] Config `reload: false` disables live reload
- [ ] Tests pass for WebSocket lifecycle
