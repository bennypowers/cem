import { build } from 'esbuild';

async function buildExtension() {
  try {
    console.log('Building extension...');

    await build({
      entryPoints: ['client/src/extension.ts'],
      bundle: true,
      outfile: 'out/client/extension.js',
      external: ['vscode'],
      format: 'esm',
      platform: 'node',
      sourcemap: true,
      minify: false,
      target: 'es2022',
      legalComments: 'none',
      banner: {
        js: `import { createRequire } from 'module'; const require = createRequire(import.meta.url);`
      },
      define: {
        'global': 'globalThis'
      }
    });

    console.log('✓ Extension built successfully');
  } catch (error) {
    console.error('Build failed:', error);
    process.exit(1);
  }
}

buildExtension();
