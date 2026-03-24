import { readFile } from 'fs/promises';

/**
 * Mock server middleware for /__cem/ routes in tests.
 * Serves actual component JS, vendor bundles, and utility modules
 * from the templates directory.
 * @param {object} ctx - Koa-like context object
 * @returns {Promise<void>}
 */
export async function mockCemServer(ctx) {
  const url = new URL(ctx.url, 'http://localhost');

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
      ctx.status = 200;
      ctx.type = 'application/json';
      ctx.body = JSON.stringify({ schemaVersion: '1.0.0', modules: [] });
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

  // Serve utility modules (websocket-client, state-persistence, etc.)
  const utilityModules = [
    'websocket-client.js',
    'state-persistence.js',
    'manifest-search.js',
    'knob-events.js',
    'health-badges.js',
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

  // Serve compiled element JS/CSS from templates/elements/
  const elementMatch = url.pathname.match(/\/__cem\/elements\/([^/]+)\/(.+)$/);
  if (elementMatch) {
    const [, elementName, fileName] = elementMatch;
    try {
      const filePath = new URL(
        `../../../templates/elements/${elementName}/${fileName}`,
        import.meta.url
      );
      const content = await readFile(filePath, 'utf-8');
      ctx.status = 200;
      ctx.type = fileName.endsWith('.css') ? 'text/css' : 'application/javascript';
      ctx.body = content;
      return;
    } catch {
      ctx.status = 404;
      ctx.body = `Not found: ${url.pathname}`;
      return;
    }
  }

  // Serve vendor bundles (lit, hydration support, etc.)
  if (url.pathname.startsWith('/__cem/vendor/')) {
    const vendorPath = url.pathname.replace('/__cem/vendor/', '');
    try {
      const filePath = new URL(
        `../../../templates/vendor/${vendorPath}`,
        import.meta.url
      );
      const content = await readFile(filePath, 'utf-8');
      ctx.status = 200;
      ctx.type = 'application/javascript';
      ctx.body = content;
      return;
    } catch {
      ctx.status = 404;
      ctx.body = `Vendor module not found: ${vendorPath}`;
      return;
    }
  }

  // Fallback: serve any /__cem/*.js as empty module
  if (url.pathname.startsWith('/__cem/') && url.pathname.endsWith('.js')) {
    ctx.status = 200;
    ctx.type = 'application/javascript';
    ctx.body = `export default {};`;
    return;
  }

  ctx.status = 404;
  ctx.body = 'Not Found';
}
