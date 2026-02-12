import { defineConfig } from 'vitest/config';
import angular from '@analogjs/vite-plugin-angular';

export default defineConfig({
  plugins: [angular({ tsconfig: 'tsconfig.json' })],
  test: {
    globals: true,
    environment: 'happy-dom',
    setupFiles: ['@analogjs/vitest-angular/setup-zone'],
  },
});
