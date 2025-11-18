import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 Navigation Group (Nested/Subnav)
 *
 * @attr {boolean} hidden - Whether the group is collapsed
 * @attr {string} aria-labelledby - ID of element labeling this group
 *
 * @slot - Default slot for nav-list containing nav-items
 */
class Pfv6NavGroup extends HTMLElement {
  static get observedAttributes() {
    return ['hidden', 'aria-labelledby'];
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('#subnav')) {
      await this.#populateShadowRoot();
    }
    this.setAttribute('role', 'region');
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-nav-group');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-nav-group template:', error);
      this.shadowRoot.innerHTML = '<section id="subnav"><div id="list" role="list"><slot></slot></div></section>';
    }
  }

  #syncAttributes() {
    const subnav = this.shadowRoot?.querySelector('#subnav');
    if (!subnav) return;

    if (this.hasAttribute('aria-labelledby')) {
      subnav.setAttribute('aria-labelledby', this.getAttribute('aria-labelledby'));
    }

    // Handle inert attribute when hidden
    if (this.hasAttribute('hidden')) {
      subnav.setAttribute('inert', '');
    } else {
      subnav.removeAttribute('inert');
    }
  }
}

customElements.define('pfv6-nav-group', Pfv6NavGroup);
