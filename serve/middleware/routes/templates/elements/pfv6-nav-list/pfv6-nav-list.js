import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 Navigation List
 *
 * @slot - Default slot for nav-item or nav-group elements
 */
class Pfv6NavList extends HTMLElement {
  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.hasChildNodes()) {
      await this.#populateShadowRoot();
    }
    this.setAttribute('role', 'list');
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-nav-list');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-nav-list template:', error);
      this.shadowRoot.innerHTML = '<slot></slot>';
    }
  }
}

customElements.define('pfv6-nav-list', Pfv6NavList);
