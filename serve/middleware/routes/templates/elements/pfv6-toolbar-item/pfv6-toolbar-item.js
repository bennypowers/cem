/* PatternFly v6 Toolbar Item Component */

import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

export class Pfv6ToolbarItem extends HTMLElement {
  static observedAttributes = ['variant'];

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('.pf-v6-c-toolbar__item')) {
      await this.#populateShadowRoot();
    }
    this._item = this.shadowRoot.querySelector('.pf-v6-c-toolbar__item');
    this._updateClasses();
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-toolbar-item');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-toolbar-item template:', error);
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
    if (!this._item) return;

    // Remove old variant classes
    this._item.classList.remove(
      'pf-m-label',
      'pf-m-pagination',
      'pf-m-expand-all',
      'pf-m-overflow-container'
    );

    // Add variant class
    const variant = this.variant;
    if (variant) {
      this._item.classList.add(`pf-m-${variant}`);
    }
  }
}

if (!customElements.get('pfv6-toolbar-item')) {
  customElements.define('pfv6-toolbar-item', Pfv6ToolbarItem);
}
