import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { resolve } from 'node:path';
import { litTransformPlugin } from '../lit-transform-plugin';

export default defineConfig({
  plugins: [
    react(),
    litTransformPlugin('/examples/kitchen-sink/'),
  ],
  resolve: {
    alias: {
      '@cem-examples/kitchen-sink': resolve(__dirname, '../../../../examples/kitchen-sink'),
    },
  },
});
