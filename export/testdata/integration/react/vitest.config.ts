import { defineConfig } from 'vitest/config';
import react from '@vitejs/plugin-react';
import { resolve } from 'node:path';

export default defineConfig({
  plugins: [react()],
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
