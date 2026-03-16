---
name: react
description: Show how to use custom elements in React, including workarounds for React 18 vs 19 differences. Use when the user asks to "use in React", "React wrapper", "custom elements in React", "JSX custom element", or mentions "React 18", "React 19", or "react integration".
tools: Read, Glob, Grep
---

# Custom Elements in React

Generate framework-specific guidance for using custom elements in React, adapting to the project's React version.

## Workflow

### Phase 1: Detect React Version

Search the project for React version:

```bash
# Check package.json for react version
grep -A1 '"react"' package.json
```

Determine if the project uses:
- **React 19+**: Native custom element support
- **React 18 or earlier**: Requires workarounds for properties and events

Also check for:
- TypeScript (need JSX type declarations)
- Existing wrapper patterns in the codebase
- Whether `@lit/react` or similar wrapper libraries are already in use

### Phase 2: Gather Element Data

Read the target element(s):

```
cem://element/{tagName}
cem://element/{tagName}/attributes
cem://element/{tagName}/events
```

Understand the element's API to know which features need framework-specific handling:
- **Attributes** (string-only): Work in all React versions
- **Properties** (complex types): Need workarounds in React 18
- **Events** (custom events): Need workarounds in React 18
- **Slots**: Map to `children` and named slot patterns in JSX
- **Boolean attributes**: Differ between React and HTML semantics

### Phase 3: Generate Integration Code

#### React 19+

React 19 supports custom elements natively. Properties are set correctly and custom events are handled.

```tsx
// Direct usage — no wrappers needed
function MyComponent() {
  return (
    <my-element
      variant="primary"
      complexData={someObject}
      onMyEvent={(e) => handleEvent(e)}
    >
      <span slot="icon">★</span>
      Button text
    </my-element>
  );
}
```

Key React 19 behaviors:
- Props matching an element property are set as properties, not attributes
- `on` + PascalCase event name maps to custom events (e.g., `onChange` -> `change`)
- `className` is set as the `class` attribute
- Boolean attributes work correctly

#### React 18 and Earlier

React 18 treats custom elements like unknown HTML elements — all props are set as attributes (stringified), and custom events are not supported via JSX.

**Option A: Use `@lit/react` wrappers (recommended)**

```tsx
import { createComponent } from '@lit/react';
import { MyElement } from 'my-element-library';

export const MyElementReact = createComponent({
  tagName: 'my-element',
  elementClass: MyElement,
  react: React,
  events: {
    onMyEvent: 'my-event',
    onChange: 'change',
  },
});

// Usage
function MyComponent() {
  return (
    <MyElementReact
      variant="primary"
      complexData={someObject}
      onMyEvent={(e) => handleEvent(e)}
    >
      Button text
    </MyElementReact>
  );
}
```

**Option B: Use refs for properties and events**

```tsx
function MyComponent() {
  const ref = useRef<HTMLElement>(null);

  useEffect(() => {
    const el = ref.current;
    if (!el) return;

    // Set complex properties via ref
    (el as any).complexData = someObject;

    // Add event listeners
    const handler = (e: Event) => handleEvent(e);
    el.addEventListener('my-event', handler);
    return () => el.removeEventListener('my-event', handler);
  }, [someObject]);

  return (
    <my-element ref={ref} variant="primary">
      Button text
    </my-element>
  );
}
```

**Option C: Use the CEM `export react` wrappers**

If the project uses CEM's export command, check for pre-generated React wrappers:

```bash
cem export react
```

This generates typed React wrapper components from the manifest.

#### TypeScript Declarations

For both React versions, add JSX type declarations so TypeScript knows about custom element attributes:

```tsx
// custom-elements.d.ts
import type { MyElement } from 'my-element-library';

declare global {
  namespace JSX {
    interface IntrinsicElements {
      'my-element': React.DetailedHTMLProps<
        React.HTMLAttributes<MyElement> & {
          variant?: 'primary' | 'secondary';
          disabled?: boolean;
          // For React 18: only string attributes here
          // For React 19: include property types too
        },
        MyElement
      >;
    }
  }
}
```

For React 19, the declaration can include complex property types. For React 18, only include attributes (strings/booleans) since complex types won't pass through correctly without a wrapper.

#### Slot Mapping

Map named slots to JSX:

```tsx
// HTML slots -> JSX
<my-element>
  {/* Default slot */}
  <span>Default content</span>

  {/* Named slots */}
  <span slot="header">Header content</span>
  <span slot="footer">Footer content</span>
</my-element>
```

### Phase 4: Output

Present the integration code with:

1. **Version-appropriate code**: Only show React 19 patterns if the project uses React 19+
2. **TypeScript declarations**: If the project uses TypeScript
3. **Event handling**: For each event in the manifest
4. **Property passing**: For each non-string property
5. **Import statements**: Show how to import the element registration

## React Version Reference

| Feature | React 18 | React 19+ |
|---------|----------|-----------|
| String attributes | Passed as attributes | Passed as attributes |
| Boolean attributes | `""` or `undefined` | Correct HTML behavior |
| Complex properties | Stringified (broken) | Set as properties |
| Custom events | Not supported in JSX | `onEventName` works |
| `className` | Set as attribute | Set as `class` |
| `ref` | Works | Works |
| SSR | Attributes only | Properties serialized |

## Guidelines

- **Check the React version first**: The guidance differs significantly between 18 and 19
- **Prefer `@lit/react` wrappers for React 18**: They handle properties and events correctly with minimal boilerplate
- **Generate TypeScript declarations from the manifest**: The element's attributes and their types are all available
- **Don't suggest `dangerouslySetInnerHTML`**: Custom elements use slots, not innerHTML
- **Handle SSR**: Note that server-rendered custom elements only have access to attributes, not properties
