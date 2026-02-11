import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { resolve } from 'node:path';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@cem-examples/kitchen-sink': resolve(__dirname, '../../../../examples/kitchen-sink'),
    },
  },
});
