# JavaScript API for `counter-element` 🔧

A counting element

This element declares **6 class methods**. They are the verbs of the element: everything it does beyond reacting to attributes and slots. Only public, non-deprecated instance methods are callable on the element; static, protected, private, and deprecated members are listed for completeness.

## Quick Reference 📋

| Method | Returns | Modifiers | What it does |
| ------ | ------- | --------- | ------------ |
| `describe` | `string` |  | Returns a human-readable state |
| `increment` | `void` |  | Adds to the count |
| `log` | `void` |  | Writes a log line |
| `reset` | `Promise<void>` |  | Resets the count to zero |
| `createRenderer` | `unknown` | static protected  | - |
| `legacyTap` | `void` | deprecated | - |

## The Full Story 📖

### `describe(): string`

Returns a human-readable state

**Returns:** `string`

Inherited from `BaseCounter` (counter-element.js).

[See the declaration](https://github.com/bennypowers/cem/tree/main/mcp/testdata/fixtures/element-methods-integration/src/base-counter.ts#L4)

### `increment(by: number = 1): void`

Adds to the count

**Parameters:**
- `by`: `number` — default `1` — How much to add

**Returns:** `void`

[See the declaration](https://github.com/bennypowers/cem/tree/main/mcp/testdata/fixtures/element-methods-integration/src/counter-element.ts#L20)

### `log(...tags: string, level: "info" | "warn"): void`

Writes a log line

**Parameters:**
- `...tags`: `string` — Tags to include
- `level` (optional): `"info" | "warn"`

**Returns:** `void`

### `reset(): Promise<void>`

Resets the count to zero

**Returns:** `Promise<void>`

### `createRenderer(): unknown`

**Returns:** `unknown`

### `legacyTap(): void` ⚠️ *deprecated*

**Deprecated**: Use increment() instead.

**Returns:** `void`

## Call It From JavaScript 💻

```javascript
const el = document.querySelector('counter-element');
el.describe();
el.increment(1);
el.log(/* string */, /* "info" | "warn" */);
el.reset();
```

---

This element also has **5 class fields** that are not backed by attributes. Unlike attributes, they are only reachable from JavaScript.

| Field | Type | Default | Modifiers | What it holds |
| ----- | ---- | ------- | --------- | ------------- |
| `events` | `string[]` | `[]` | readonly  | Names of every fired event |
| `count` | `number` | `0` |  | The current count |
| `describe` | `string` | `''` |  | Cached description text; shares a name with the inherited describe() method to pin deep-link precedence |
| `history` | `number[]` | `[]` | readonly protected  | Every value the counter has taken |
| `registry` | `Map<string, CounterElement>` | - | static  | All live instances |

- **`events`**: `string[]` *(readonly)* *(inherited from `BaseCounter`)* — Names of every fired event
- **`count`**: `number` — The current count
- **`describe`**: `string` — Cached description text; shares a name with the inherited describe() method to pin deep-link precedence
- **`history`**: `number[]` *(readonly)* *(protected)* — Every value the counter has taken
- **`registry`**: `Map<string, CounterElement>` *(static)* — All live instances

## Keep Exploring 🧭

• **[Attributes](cem://element/counter-element/attributes)** — Configure the behavior
• **[Full Reference](cem://element/counter-element)** — Everything about this element
• **[Generate HTML](cem://tools/generate_html)** — Let AI create the markup for you
