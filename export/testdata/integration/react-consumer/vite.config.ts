import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { litTransformPlugin } from '../lit-transform-plugin';

export default defineConfig({
  plugins: [
    react(),
    litTransformPlugin('/examples/kitchen-sink/'),
  ],
});
