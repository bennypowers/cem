import { readFile } from 'fs/promises';
import { fileURLToPath } from 'url';

/**
 * Mock server middleware for /__cem/ routes
 * Handles component template loading and manifest endpoints
 * @param {object} ctx - Koa-like context object
 * @returns {Promise<void>}
 */
export async function mockCemServer(ctx) {
  const url = new URL(ctx.url, 'http://localhost');

  // Mock element HTML/CSS loading
  // Pattern: /__cem/elements/{element-name}/{element-name}.{html|css}
  const elementMatch = url.pathname.match(/\/__cem\/elements\/([^/]+)\/\1\.(html|css)$/);
  if (elementMatch) {
    const [, elementName, type] = elementMatch;

    try {
      // Try to load from actual templates directory
      const templatePath = new URL(
        `../../../templates/elements/${elementName}/${elementName}.${type}`,
        import.meta.url
      );
      const content = await readFile(templatePath, 'utf-8');

      ctx.status = 200;
      ctx.type = type === 'html' ? 'text/html' : 'text/css';
      ctx.body = content;
      return;
    } catch (error) {
      // If file doesn't exist, return minimal content
      if (type === 'html') {
        ctx.status = 200;
        ctx.type = 'text/html';
        ctx.body = `<div id="mock-${elementName}">Mock ${elementName}</div>`;
        return;
      } else {
        ctx.status = 200;
        ctx.type = 'text/css';
        ctx.body = `/* Mock styles for ${elementName} */`;
        return;
      }
    }
  }

  // Mock manifest endpoint
  if (url.pathname === '/__cem/manifest.json') {
    try {
      const manifestPath = new URL(
        '../fixtures/mock-manifests/default-manifest.json',
        import.meta.url
      );
      const content = await readFile(manifestPath, 'utf-8');
      ctx.status = 200;
      ctx.type = 'application/json';
      ctx.body = content;
      return;
    } catch {
      // Return minimal manifest
      ctx.status = 200;
      ctx.type = 'application/json';
      ctx.body = JSON.stringify({
        schemaVersion: '1.0.0',
        modules: []
      });
      return;
    }
  }

  // Mock logs endpoint
  if (url.pathname === '/__cem/logs') {
    ctx.status = 200;
    ctx.type = 'application/json';
    ctx.body = JSON.stringify([]);
    return;
  }

  // Serve actual CemElement source
  if (url.pathname === '/__cem/cem-element.js') {
    try {
      const cemElementPath = new URL(
        '../../../templates/js/cem-element.js',
        import.meta.url
      );
      const content = await readFile(cemElementPath, 'utf-8');
      ctx.status = 200;
      ctx.type = 'application/javascript';
      ctx.body = content;
      return;
    } catch (error) {
      console.error('Failed to load cem-element.js:', error);
      ctx.status = 500;
      ctx.body = 'Failed to load cem-element.js';
      return;
    }
  }

  // Serve actual utility modules
  const utilityModules = [
    'websocket-client.js',
    'state-persistence.js',
    'manifest-search.js',
    'knob-events.js'
  ];

  const utilityMatch = utilityModules.find(mod => url.pathname === `/__cem/${mod}`);
  if (utilityMatch) {
    try {
      const modulePath = new URL(
        `../../../templates/js/${utilityMatch}`,
        import.meta.url
      );
      const content = await readFile(modulePath, 'utf-8');
      ctx.status = 200;
      ctx.type = 'application/javascript';
      ctx.body = content;
      return;
    } catch (error) {
      console.error(`Failed to load ${utilityMatch}:`, error);
      ctx.status = 404;
      ctx.body = `Module not found: ${utilityMatch}`;
      return;
    }
  }

  // Serve actual CEM component source files
  const cemComponentMatch = url.pathname.match(/\/__cem\/elements\/(cem-[^/]+)\/\1\.js$/);
  if (cemComponentMatch) {
    const componentName = cemComponentMatch[1];
    try {
      const componentPath = new URL(
        `../../../templates/elements/${componentName}/${componentName}.js`,
        import.meta.url
      );
      const content = await readFile(componentPath, 'utf-8');
      ctx.status = 200;
      ctx.type = 'application/javascript';
      ctx.body = content;
      return;
    } catch (error) {
      console.error(`Failed to load ${componentName}.js:`, error);
      ctx.status = 404;
      ctx.body = `Component not found: ${componentName}`;
      return;
    }
  }

  // Mock PatternFly and other component JavaScript imports
  // Pattern: /__cem/{path}.js or /__cem/elements/{name}/{name}.js
  if (url.pathname.startsWith('/__cem/') && url.pathname.endsWith('.js')) {
    ctx.status = 200;
    ctx.type = 'application/javascript';
    // Return empty module that defines the custom element if it's a component
    const elementMatch = url.pathname.match(/elements\/([^/]+)\/\1\.js$/);
    if (elementMatch) {
      const elementName = elementMatch[1];
      ctx.body = `
// Mock component for ${elementName}
if (!customElements.get('${elementName}')) {
  customElements.define('${elementName}', class extends HTMLElement {
    constructor() {
      super();
      this.attachShadow({ mode: 'open' });
    }
  });
}
export default {};
`;
    } else {
      ctx.body = `// Mock module for ${url.pathname}\nexport default {};`;
    }
    return;
  }

  // Default 404 for unknown paths
  ctx.status = 404;
  ctx.body = 'Not Found';
}
