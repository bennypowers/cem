## Auto-Reload and HMR

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

