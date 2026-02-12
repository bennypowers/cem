import { defineConfig } from 'vitest/config';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'node:path';

export default defineConfig({
  plugins: [vue({
    template: {
      compilerOptions: {
        isCustomElement: tag => tag.startsWith('demo-'),
      },
    },
  })],
  resolve: {
    alias: [
      { find: '@cem-examples/kitchen-sink', replacement: resolve(__dirname, '../../../../examples/kitchen-sink') },
      { find: /^lit$/, replacement: resolve(__dirname, 'node_modules/lit') },
      { find: /^lit\/(.*)/, replacement: resolve(__dirname, 'node_modules/lit/$1') },
    ],
  },
  test: {
    environment: 'happy-dom',
  },
});
