import { playwrightLauncher } from '@web/test-runner-playwright';
import { visualRegressionPlugin } from '@web/test-runner-visual-regression/plugin';

const updateGoldens = process.argv.includes('--update-goldens');

export default {
  files: [
    'middleware/routes/templates/**/*.test.js',
    'frontend-tests/**/*.test.js'
  ],
  nodeResolve: true,
  // Only show errors in browser logs, filter out dev mode warnings
  filterBrowserLogs: ({ type }) => type === 'error',
  browsers: [
    playwrightLauncher({ product: 'chromium' }),
    playwrightLauncher({ product: 'firefox' }),
    // WebKit disabled due to missing system dependencies
    // playwrightLauncher({ product: 'webkit' }),
  ],
  testFramework: {
    config: {
      timeout: 5000,
    },
  },
  coverage: true,
  coverageConfig: {
    reportDir: 'coverage/frontend',
    exclude: [
      '**/__cem/**',  // Exclude virtual modules served by mock server
      '**/node_modules/**',
    ],
    threshold: {
      statements: 75,
      branches: 70,
      functions: 75,
      lines: 75,
    },
  },
  plugins: [
    visualRegressionPlugin({
      update: updateGoldens,
      // Only run visual tests in Chromium to avoid x-browser flakiness
      browsers: ['chromium'],
      diffOptions: {
        threshold: 0.05,  // 5% tolerance for rendering differences
      },
      baseDir: 'middleware/routes/testdata/frontend/visual-goldens',
    }),
  ],
  // Proxy /__cem/ routes to real cem serve instance
  middleware: [
    async (ctx, next) => {
      if (ctx.url.startsWith('/__cem/')) {
        // Proxy to real cem serve running on port 9876
        const proxyUrl = `http://localhost:9876${ctx.url}`;
        try {
          const response = await fetch(proxyUrl);
          ctx.status = response.status;
          ctx.type = response.headers.get('content-type');
          ctx.body = await response.text();
        } catch (error) {
          console.error(`Failed to proxy ${ctx.url}:`, error.message);
          ctx.status = 502;
          ctx.body = `Proxy error: ${error.message}`;
        }
        return;
      }
      await next();
    },
  ],
};
