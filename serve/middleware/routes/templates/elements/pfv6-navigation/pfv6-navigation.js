import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 Navigation Container
 *
 * @attr {string} aria-label - Accessible label for navigation
 * @attr {boolean} inset - Add horizontal padding
 *
 * @slot - Default slot for nav-list
 */
class Pfv6Navigation extends HTMLElement {
  static get observedAttributes() {
    return ['aria-label', 'inset'];
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('nav')) {
      await this.#populateShadowRoot();
    }
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-navigation');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-navigation template:', error);
      this.shadowRoot.innerHTML = '<nav><slot></slot></nav>';
    }
  }

  #syncAttributes() {
    const nav = this.shadowRoot?.querySelector('nav');
    if (!nav) return;

    if (this.hasAttribute('aria-label')) {
      nav.setAttribute('aria-label', this.getAttribute('aria-label'));
    }
  }
}

customElements.define('pfv6-navigation', Pfv6Navigation);
