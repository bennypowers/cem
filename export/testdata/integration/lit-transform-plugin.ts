import { resolve, dirname } from 'node:path';
import { readFileSync } from 'node:fs';
import ts from 'typescript';
import type { Plugin } from 'vite';

/**
 * Vite plugin that compiles Lit element TypeScript sources.
 *
 * Needed because:
 * - esbuild doesn't support the `accessor` keyword used by Lit decorators
 * - Vite can't serve `import ... with { type: 'css' }` as native CSS module scripts
 *
 * This plugin uses TypeScript's transpileModule for decorator/accessor support
 * and inlines CSS imports as CSSStyleSheet construction.
 */
export function litTransformPlugin(pathMatch: string): Plugin {
  return {
    name: 'lit-transform',
    enforce: 'pre',
    transform(code, id) {
      if (id.includes(pathMatch) && /\.ts$/.test(id)) {
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
          /import\s+(\w+)\s+from\s+["']([^"']+\.css)["']\s*(?:with|assert)\s*\{[^}]*\}\s*;/g,
          (_match, varName, cssPath) => {
            const cssFile = resolve(dirname(id), cssPath);
            const cssContent = readFileSync(cssFile, 'utf-8');
            return `const ${varName} = new CSSStyleSheet();\n${varName}.replaceSync(${JSON.stringify(cssContent)});`;
          },
        );
        return { code: output, map: null };
      }
    },
  };
}
