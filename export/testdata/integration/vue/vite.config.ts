import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'node:path';
import { litTransformPlugin } from '../lit-transform-plugin';

export default defineConfig({
  plugins: [
    vue({
      template: {
        compilerOptions: {
          isCustomElement: tag => tag.startsWith('demo-'),
        },
      },
    }),
    litTransformPlugin('/examples/kitchen-sink/'),
  ],
  resolve: {
    alias: {
      '@cem-examples/kitchen-sink': resolve(__dirname, '../../../../examples/kitchen-sink'),
    },
  },
});
