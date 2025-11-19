import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 Page Sidebar
 *
 * @attr {boolean} collapsed - Whether the sidebar is collapsed
 * @attr {boolean} expanded - Whether the sidebar is expanded
 *
 * @slot - Default slot for sidebar content (typically pfv6-navigation)
 */
class Pfv6PageSidebar extends HTMLElement {
  static formAssociated = false;

  #internals;

  static get observedAttributes() {
    return ['collapsed', 'expanded'];
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
    this.#internals = this.attachInternals();
    this.#internals.role = 'navigation';
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('.pf-v6-c-page__sidebar')) {
      await this.#populateShadowRoot();
    }
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-page-sidebar');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-page-sidebar template:', error);
      this.shadowRoot.innerHTML = '<div class="pf-v6-c-page__sidebar"><div class="pf-v6-c-page__sidebar-body"><slot></slot></div></div>';
    }
  }

  get collapsed() {
    return this.hasAttribute('collapsed');
  }

  set collapsed(value) {
    if (value) {
      this.setAttribute('collapsed', '');
      this.removeAttribute('expanded');
    } else {
      this.removeAttribute('collapsed');
      this.setAttribute('expanded', '');
    }
  }

  get expanded() {
    return this.hasAttribute('expanded');
  }

  set expanded(value) {
    if (value) {
      this.setAttribute('expanded', '');
      this.removeAttribute('collapsed');
    } else {
      this.removeAttribute('expanded');
      this.setAttribute('collapsed', '');
    }
  }
}

customElements.define('pfv6-page-sidebar', Pfv6PageSidebar);
