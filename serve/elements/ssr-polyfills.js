// Polyfill Node.js / browser globals that @lit-labs/ssr and our
// components expect but QuickJS doesn't provide.

// Add cssText getter to the SSR CSSStyleSheet shim so that
// LitElementRenderer.renderShadow() can read styles.
// The shim stores CSS in cssRules[0].cssText but doesn't expose
// a top-level cssText property that the renderer expects.
import { CSSStyleSheet } from '@lit-labs/ssr-dom-shim';

Object.defineProperty(CSSStyleSheet.prototype, 'cssText', {
  get() {
    return Array.from(this.cssRules).map(r => r.cssText).join('');
  },
  configurable: true,
});

// btoa/atob are needed for template digest computation in @lit-labs/ssr-client.
// QuickJS doesn't have them natively.
if (typeof btoa === 'undefined') {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
  globalThis.btoa = function(s) {
    let out = '';
    for (let i = 0; i < s.length; i += 3) {
      const a = s.charCodeAt(i), b = s.charCodeAt(i+1), c = s.charCodeAt(i+2);
      out += chars[a >> 2] + chars[((a & 3) << 4) | (b >> 4)]
           + (isNaN(b) ? '=' : chars[((b & 15) << 2) | (c >> 6)])
           + (isNaN(c) ? '=' : chars[c & 63]);
    }
    return out;
  };
}

// QuickJS globals that Node.js provides but QuickJS doesn't
if (typeof URL === 'undefined') {
  globalThis.URL = class URL {
    constructor(u, base) { this.href = String(u); }
    toString() { return this.href; }
  };
}

if (typeof URLSearchParams === 'undefined') {
  globalThis.URLSearchParams = class URLSearchParams {
    constructor() {}
  };
}

// Browser global used by components at module scope
if (typeof CSS === 'undefined') {
  globalThis.CSS = {
    supports() { return false; },
    escape(s) { return String(s); },
  };
}
