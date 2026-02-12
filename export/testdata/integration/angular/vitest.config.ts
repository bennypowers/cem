import { defineConfig } from 'vitest/config';
import angular from '@analogjs/vite-plugin-angular';
import { resolve } from 'node:path';

export default defineConfig({
  plugins: [angular({ tsconfig: 'tsconfig.json' })],
  resolve: {
    alias: {
      '@cem-examples/kitchen-sink': resolve(__dirname, '../../../../examples/kitchen-sink'),
    },
  },
  test: {
    globals: true,
    environment: 'happy-dom',
    setupFiles: ['@analogjs/vitest-angular/setup-zone'],
  },
});
