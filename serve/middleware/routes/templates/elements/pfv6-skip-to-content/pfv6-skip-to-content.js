import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 Skip to Content
 *
 * @attr {string} href - Target anchor ID to skip to (e.g., "#main-content")
 *
 * @slot - Default slot for skip link text (defaults to "Skip to content")
 */
class Pfv6SkipToContent extends HTMLElement {
  static get observedAttributes() {
    return ['href'];
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('#link')) {
      await this.#populateShadowRoot();
    }
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-skip-to-content');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-skip-to-content template:', error);
      this.shadowRoot.innerHTML = '<div class="pf-v6-c-skip-to-content"><a id="link"><slot>Skip to content</slot></a></div>';
    }
  }

  #syncAttributes() {
    const link = this.shadowRoot?.querySelector('#link');
    if (!link) return;

    // Default to #main-content if no href is specified
    const href = this.getAttribute('href') || '#main-content';
    link.setAttribute('href', href);
  }
}

customElements.define('pfv6-skip-to-content', Pfv6SkipToContent);
