import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 Page Main Content
 *
 * @attr {string} id - ID for the main content area (defaults to "main-content")
 *
 * @slot - Default slot for main content
 */
class Pfv6PageMain extends HTMLElement {
  static get observedAttributes() {
    return ['id'];
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('.pf-v6-c-page__main')) {
      await this.#populateShadowRoot();
    }
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-page-main');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-page-main template:', error);
      this.shadowRoot.innerHTML = '<div class="pf-v6-c-page__main-container"><main class="pf-v6-c-page__main" id="main-content" tabindex="-1"><slot></slot></main></div>';
    }
  }

  #syncAttributes() {
    const main = this.shadowRoot?.querySelector('.pf-v6-c-page__main');
    if (!main) return;

    // Use provided ID or default to main-content
    const id = this.getAttribute('id') || 'main-content';
    main.setAttribute('id', id);
  }
}

customElements.define('pfv6-page-main', Pfv6PageMain);
