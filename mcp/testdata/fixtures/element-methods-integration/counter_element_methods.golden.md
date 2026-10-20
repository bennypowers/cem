# JavaScript API for `counter-element` 🔧

A counting element

This element declares **6 class methods**. They are the verbs of the element: everything it does beyond reacting to attributes and slots. Only public, non-deprecated instance methods are callable on the element; static, protected, private, and deprecated members are listed for completeness.

## Quick Reference 📋

| Method | Returns | Modifiers | What it does |
| ------ | ------- | --------- | ------------ |
| [`describe`](cem://element/counter-element/methods/describe) | `string` |  | Returns a human-readable state |
| [`increment`](cem://element/counter-element/methods/increment) | `void` |  | Adds to the count |
| [`log`](cem://element/counter-element/methods/log) | `void` |  | Writes a log line |
| [`reset`](cem://element/counter-element/methods/reset) | `Promise<void>` |  | Resets the count to zero |
| [`createRenderer`](cem://element/counter-element/methods/createRenderer) | `unknown` | static protected  | - |
| [`legacyTap`](cem://element/counter-element/methods/legacyTap) | `void` | deprecated | - |

## The Full Story 📖

### [`describe(): string`](cem://element/counter-element/methods/describe)

Returns a human-readable state

**Returns:** `string`

Inherited from `BaseCounter` (counter-element.js).

[See the declaration](https://github.com/bennypowers/cem/tree/main/mcp/testdata/fixtures/element-methods-integration/src/base-counter.ts#L4)

### [`increment(by: number = 1): void`](cem://element/counter-element/methods/increment)

Adds to the count

**Parameters:**
- `by`: `number` — default `1` — How much to add

**Returns:** `void`

[See the declaration](https://github.com/bennypowers/cem/tree/main/mcp/testdata/fixtures/element-methods-integration/src/counter-element.ts#L20)

### [`log(...tags: string, level: "info" | "warn"): void`](cem://element/counter-element/methods/log)

Writes a log line

**Parameters:**
- `...tags`: `string` — Tags to include
- `level` (optional): `"info" | "warn"`

**Returns:** `void`

### [`reset(): Promise<void>`](cem://element/counter-element/methods/reset)

Resets the count to zero

**Returns:** `Promise<void>`

### [`createRenderer(): unknown`](cem://element/counter-element/methods/createRenderer)

**Returns:** `unknown`

### [`legacyTap(): void`](cem://element/counter-element/methods/legacyTap) ⚠️ *deprecated*

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
| [`events`](cem://element/counter-element/methods/events) | `string[]` | `[]` | readonly  | Names of every fired event |
| [`count`](cem://element/counter-element/methods/count) | `number` | `0` |  | The current count |
| [`describe`](cem://element/counter-element/methods/describe) | `string` | `''` |  | Cached description text; shares a name with the inherited describe() method to pin deep-link precedence |
| [`history`](cem://element/counter-element/methods/history) | `number[]` | `[]` | readonly protected  | Every value the counter has taken |
| [`registry`](cem://element/counter-element/methods/registry) | `Map<string, CounterElement>` | - | static  | All live instances |

- **[`events`](cem://element/counter-element/methods/events)**: `string[]` *(readonly)* *(inherited from `BaseCounter`)* — Names of every fired event
- **[`count`](cem://element/counter-element/methods/count)**: `number` — The current count
- **[`describe`](cem://element/counter-element/methods/describe)**: `string` — Cached description text; shares a name with the inherited describe() method to pin deep-link precedence
- **[`history`](cem://element/counter-element/methods/history)**: `number[]` *(readonly)* *(protected)* — Every value the counter has taken
- **[`registry`](cem://element/counter-element/methods/registry)**: `Map<string, CounterElement>` *(static)* — All live instances

## Keep Exploring 🧭

• **[Attributes](cem://element/counter-element/attributes)** — Configure the behavior
• **[Full Reference](cem://element/counter-element)** — Everything about this element
• **[Generate HTML](cem://tools/generate_html)** — Let AI create the markup for you
