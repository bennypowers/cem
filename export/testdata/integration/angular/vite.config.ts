import { defineConfig } from 'vite';
import angular from '@analogjs/vite-plugin-angular';
import { resolve } from 'node:path';
import { litTransformPlugin } from '../lit-transform-plugin';

export default defineConfig({
  plugins: [
    angular({
      tsconfig: 'tsconfig.json',
      transformFilter: (_code: string, id: string) =>
        !id.includes('/examples/kitchen-sink/'),
    }),
    litTransformPlugin('/examples/kitchen-sink/'),
  ],
  resolve: {
    alias: {
      '@cem-examples/kitchen-sink': resolve(__dirname, '../../../../examples/kitchen-sink'),
    },
  },
});
