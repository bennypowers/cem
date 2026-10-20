# JavaScript API for `counter-element` 🔧

A counting element

This element exposes **4 methods** you can call from JavaScript. They are the verbs of the element: everything it can do beyond reacting to attributes and slots.

## Quick Reference 📋

| Method | Returns | Modifiers | What it does |
| ------ | ------- | --------- | ------------ |
| `increment` | `void` |  | Adds to the count |
| `reset` | `Promise<void>` |  | Resets the count to zero |
| `createRenderer` | `unknown` | static protected  | - |
| `legacyTap` | `void` | deprecated | - |

## The Full Story 📖

### `increment(by: number = 1): void`

Adds to the count

**Parameters:**
- `by`: `number` — default `1` — How much to add

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
el.increment(1);
el.reset();
```

---

This element also has **3 class fields** that are not backed by attributes. Unlike attributes, they are only reachable from JavaScript.

| Field | Type | Default | Modifiers | What it holds |
| ----- | ---- | ------- | --------- | ------------- |
| `count` | `number` | `0` |  | The current count |
| `history` | `number[]` | `[]` | readonly protected  | Every value the counter has taken |
| `registry` | `Map<string, CounterElement>` | - | static  | All live instances |

- **`count`**: `number` — The current count
- **`history`**: `number[]` *(readonly)* *(protected)* — Every value the counter has taken
- **`registry`**: `Map<string, CounterElement>` *(static)* — All live instances

## Keep Exploring 🧭

• **[Attributes](cem://element/counter-element/attributes)** — Configure the behavior
• **[Full Reference](cem://element/counter-element)** — Everything about this element
• **[Generate HTML](cem://tools/generate_html)** — Let AI create the markup for you
