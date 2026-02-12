import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
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
});
