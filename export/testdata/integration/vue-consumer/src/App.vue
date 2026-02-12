<script setup lang="ts">
import { DemoButton, DemoCard, DemoSpinner } from '../generated';
</script>

<template>
  <div style="padding: 2rem; font-family: system-ui">
    <h1>Vue Export Demo (Consumer)</h1>

    <section>
      <h2>DemoButton</h2>
      <DemoButton variant="brand">Brand Button</DemoButton>
      <DemoButton variant="danger" :disabled="true">Disabled Danger</DemoButton>
    </section>

    <section>
      <h2>DemoCard</h2>
      <DemoCard>
        <span slot="header">Card Title</span>
        <p>Card body content goes here.</p>
      </DemoCard>
    </section>

    <section>
      <h2>DemoSpinner</h2>
      <DemoSpinner />
    </section>
  </div>
</template>
