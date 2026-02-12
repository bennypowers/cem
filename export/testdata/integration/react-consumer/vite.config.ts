import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { resolve, dirname } from 'node:path';
import { readFileSync } from 'node:fs';
import ts from 'typescript';

export default defineConfig({
  plugins: [
    react(),
    {
      name: 'lit-transform',
      enforce: 'pre',
      transform(code, id) {
        if (id.includes('/examples/kitchen-sink/') && /\.ts$/.test(id)) {
          const result = ts.transpileModule(code, {
            compilerOptions: {
              target: ts.ScriptTarget.ES2022,
              module: ts.ModuleKind.ESNext,
              moduleResolution: ts.ModuleResolutionKind.Bundler,
              esModuleInterop: true,
            },
            fileName: id,
          });
          const output = result.outputText.replace(
            /import\s+(\w+)\s+from\s+["']([^"']+\.css)["']\s*with\s*\{[^}]*\}\s*;/g,
            (_match, varName, cssPath) => {
              const cssFile = resolve(dirname(id), cssPath);
              const cssContent = readFileSync(cssFile, 'utf-8');
              return `const ${varName} = new CSSStyleSheet();\n${varName}.replaceSync(${JSON.stringify(cssContent)});`;
            },
          );
          return { code: output, map: null };
        }
      },
    },
  ],
});
