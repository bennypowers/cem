---
name: vue
description: Show how to use custom elements in Vue, including required configuration for the Vue compiler. Use when the user asks to "use in Vue", "Vue custom elements", "custom elements in Vue", "isCustomElement", or mentions "Vue 3", "Vue 2", "Nuxt", or "vue integration".
tools: Read, Glob, Grep
---

# Custom Elements in Vue

Generate framework-specific guidance for using custom elements in Vue, including required compiler configuration.

## Workflow

### Phase 1: Detect Vue Version and Build Setup

Search the project for Vue version and tooling:

- Check `package.json` for `vue` version (2.x vs 3.x)
- Check for build tool: Vite (`vite.config`), Vue CLI (`vue.config`), Nuxt (`nuxt.config`), or webpack
- Check for TypeScript usage
- Check for existing custom element configuration

### Phase 2: Gather Element Data

Read the target element(s):

```
cem://element/{tagName}
cem://element/{tagName}/attributes
cem://element/{tagName}/events
cem://element/{tagName}/slots
```

Understand which features need framework-specific handling:
- **Attributes vs properties**: Vue can bind either with different syntax
- **Events**: Vue's `v-on` / `@` syntax works with custom events
- **Slots**: Vue slots and web component slots interact
- **v-model**: May need custom setup for two-way binding

### Phase 3: Required Configuration

#### Vue 3 + Vite

```ts
// vite.config.ts
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [
    vue({
      template: {
        compilerOptions: {
          // Tell Vue to treat tags with a dash as custom elements
          isCustomElement: (tag) => tag.includes('-'),
        },
      },
    }),
  ],
});
```

If the project only uses elements from specific libraries, be more specific:

```ts
isCustomElement: (tag) => tag.startsWith('my-') || tag.startsWith('prefix-'),
```

#### Vue 3 + Vue CLI (webpack)

```js
// vue.config.js
module.exports = {
  chainWebpack: (config) => {
    config.module
      .rule('vue')
      .use('vue-loader')
      .tap((options) => ({
        ...options,
        compilerOptions: {
          isCustomElement: (tag) => tag.includes('-'),
        },
      }));
  },
};
```

#### Nuxt 3

```ts
// nuxt.config.ts
export default defineNuxtConfig({
  vue: {
    compilerOptions: {
      isCustomElement: (tag) => tag.includes('-'),
    },
  },
});
```

#### Vue 2

Vue 2 uses `Vue.config.ignoredElements`:

```js
// main.js
Vue.config.ignoredElements = [
  'my-element',
  /^prefix-/,  // regex for all elements with a prefix
];
```

Note: Vue 2 reached end-of-life on December 31, 2023. Consider migrating to Vue 3.

### Phase 4: Generate Integration Code

#### Basic Usage

```vue
<template>
  <my-element
    variant="primary"
    :disabled="isDisabled"
    @my-event="handleEvent"
  >
    Default slot content
  </my-element>
</template>

<script setup lang="ts">
import 'my-element-library';

const isDisabled = ref(false);

function handleEvent(e: CustomEvent) {
  console.log(e.detail);
}
</script>
```

#### Attribute vs Property Binding

```vue
<template>
  <!-- Attribute binding (string coercion) -->
  <my-element variant="primary" />

  <!-- Property binding (passes JS values directly) -->
  <my-element :complex-data="myObject" />
  <!-- or explicitly with .prop modifier -->
  <my-element :complex-data.prop="myObject" />

  <!-- Boolean attribute -->
  <my-element :disabled="true" />
  <!-- or just present for true -->
  <my-element disabled />
</template>
```

Vue 3 passes `:bound` values as properties if the element has a matching property, and as attributes otherwise. Use the `.prop` modifier to force property binding.

In Vue 2, use the `.prop` modifier explicitly:

```vue
<my-element :complex-data.prop="myObject" />
```

#### Event Handling

Vue's `@` / `v-on` works with custom element events:

```vue
<template>
  <!-- kebab-case event names work directly -->
  <my-element @my-event="handleMyEvent" />

  <!-- Access event detail -->
  <my-element @change="(e) => value = e.detail.value" />
</template>
```

Note: Vue converts `@myEvent` to listen for `my-event` (camelCase to kebab-case). If the custom element fires a camelCase event name, use the explicit form:

```vue
<my-element v-on:myEvent="handler" />
```

#### Slots

Web component slots and Vue slots coexist but work differently:

**Important**: Do not use Vue's `v-slot` / `#slotname` syntax for web component slots. Use the native HTML `slot="name"` attribute instead.

```vue
<template>
  <my-element>
    <!-- Default slot -->
    <span>Content goes to the default slot</span>

    <!-- Named slots — use the native HTML slot attribute -->
    <span slot="header">This goes to the web component's header slot</span>
    <span slot="footer">This goes to the web component's footer slot</span>
  </my-element>
</template>
```

#### v-model

Custom elements don't support `v-model` out of the box. Wire it up manually:

```vue
<template>
  <my-input
    :value="inputValue"
    @input="inputValue = $event.detail.value"
  />
</template>
```

Or create a composable for reusable v-model binding:

```ts
function useCustomElementModel(eventName = 'input', detailKey = 'value') {
  return {
    modelValue: (props: any) => props.modelValue,
    listener: (emit: any) => (e: CustomEvent) => emit('update:modelValue', e.detail[detailKey]),
  };
}
```

#### TypeScript

Add type declarations for custom elements in Vue templates:

```ts
// custom-elements.d.ts
declare module 'vue' {
  interface GlobalComponents {
    'my-element': import('my-element-library').MyElement;
  }
}
```

For more detailed type checking, use the element's type with Vue's `HTMLAttributes`:

```ts
declare module 'vue' {
  interface GlobalComponents {
    'my-element': DefineComponent<{
      variant?: 'primary' | 'secondary';
      disabled?: boolean;
    }>;
  }
}
```

### Phase 5: Output

Present the integration with:

1. **Required config** appropriate to the project's build tool
2. **Usage examples** for each element the user asked about
3. **Property binding** for complex (non-string) properties
4. **Event handling** for each event in the manifest
5. **Slot usage** with the correct native `slot` attribute (not Vue `v-slot`)
6. **TypeScript declarations** if the project uses TypeScript

## Guidelines

- **Always include the `isCustomElement` config**: Without it, Vue warns about unresolved components
- **Use native `slot` attribute, not `v-slot`**: This is the most common mistake when using web components in Vue
- **Prefer `:prop` binding for non-strings**: Vue's property binding passes values correctly to custom elements
- **Check for Nuxt**: Nuxt has its own config location for compiler options
- **Note Vue 2 EOL**: If the project uses Vue 2, mention the EOL status but still provide working guidance
