/* PatternFly v6 Toolbar Group Component */

import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

export class Pfv6ToolbarGroup extends HTMLElement {
  static observedAttributes = ['variant'];

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('.pf-v6-c-toolbar__group')) {
      await this.#populateShadowRoot();
    }
    this._group = this.shadowRoot.querySelector('.pf-v6-c-toolbar__group');
    this._updateClasses();
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

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue !== newValue) {
      this._updateClasses();
    }
  }

  get variant() {
    return this.getAttribute('variant') || '';
  }

  set variant(value) {
    if (value) {
      this.setAttribute('variant', value);
    } else {
      this.removeAttribute('variant');
    }
  }

  _updateClasses() {
    if (!this._group) return;

    // Remove old variant classes
    this._group.classList.remove(
      'pf-m-filter-group',
      'pf-m-label-group',
      'pf-m-action-group',
      'pf-m-action-group-plain',
      'pf-m-action-group-inline',
      'pf-m-overflow-container'
    );

    // Add variant class
    const variant = this.variant;
    if (variant) {
      this._group.classList.add(`pf-m-${variant}`);
    }
  }
}

if (!customElements.get('pfv6-toolbar-group')) {
  customElements.define('pfv6-toolbar-group', Pfv6ToolbarGroup);
}
