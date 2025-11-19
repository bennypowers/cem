/* PatternFly v6 Toolbar Group Component */

import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

export class Pfv6ToolbarGroup extends HTMLElement {
  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.firstChild) {
      await this.#populateShadowRoot();
    }
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-toolbar-group');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-toolbar-group template:', error);
    }
  }
}

if (!customElements.get('pfv6-toolbar-group')) {
  customElements.define('pfv6-toolbar-group', Pfv6ToolbarGroup);
}
