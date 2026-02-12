import { defineConfig } from 'vitest/config';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue({
    template: {
      compilerOptions: {
        isCustomElement: tag => tag.startsWith('demo-'),
      },
    },
  })],
  resolve: {
    dedupe: ['lit'],
  },
  test: {
    environment: 'happy-dom',
  },
});
