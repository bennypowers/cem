import { defineConfig } from 'vitest/config';
import angular from '@analogjs/vite-plugin-angular';
import { resolve } from 'node:path';

export default defineConfig({
  plugins: [angular({ tsconfig: 'tsconfig.json' })],
  resolve: {
    alias: [
      { find: '@cem-examples/kitchen-sink', replacement: resolve(__dirname, '../../../../examples/kitchen-sink') },
      { find: /^lit$/, replacement: resolve(__dirname, 'node_modules/lit') },
      { find: /^lit\/(.*)/, replacement: resolve(__dirname, 'node_modules/lit/$1') },
    ],
  },
  test: {
    globals: true,
    environment: 'happy-dom',
    setupFiles: ['@analogjs/vitest-angular/setup-zone'],
  },
});
