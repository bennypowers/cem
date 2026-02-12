import { defineConfig } from 'vite';
import angular from '@analogjs/vite-plugin-angular';
import { resolve, dirname } from 'node:path';
import { readFileSync } from 'node:fs';
import ts from 'typescript';

export default defineConfig({
  plugins: [
    angular({
      tsconfig: 'tsconfig.json',
      transformFilter: (_code: string, id: string) =>
        !id.includes('/examples/kitchen-sink/'),
    }),
    {
      name: 'lit-transform',
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
          // Replace CSS import attributes with inline CSSStyleSheet construction
          // Vite can't serve CSS as native module scripts, so we inline them
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
