/* PatternFly v6 Toolbar Component */

import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

export class Pfv6Toolbar extends HTMLElement {
  static observedAttributes = ['sticky', 'full-height', 'color-variant', 'expandable', 'expanded'];

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
    this._expandableContent = this.shadowRoot.querySelector('.pf-v6-c-toolbar__expandable-content');
    this._updateExpandableContent();
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-toolbar');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-toolbar template:', error);
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue !== newValue && (name === 'expandable' || name === 'expanded')) {
      this._updateExpandableContent();
    }
  }

  get sticky() {
    return this.hasAttribute('sticky');
  }

  set sticky(value) {
    if (value) {
      this.setAttribute('sticky', '');
    } else {
      this.removeAttribute('sticky');
    }
  }

  get fullHeight() {
    return this.hasAttribute('full-height');
  }

  set fullHeight(value) {
    if (value) {
      this.setAttribute('full-height', '');
    } else {
      this.removeAttribute('full-height');
    }
  }

  get colorVariant() {
    return this.getAttribute('color-variant') || '';
  }

  set colorVariant(value) {
    if (value) {
      this.setAttribute('color-variant', value);
    } else {
      this.removeAttribute('color-variant');
    }
  }

  _updateExpandableContent() {
    if (!this._expandableContent) return;

    const expandable = this.hasAttribute('expandable');
    const expanded = this.hasAttribute('expanded');

    if (expandable) {
      this._expandableContent.hidden = false;
      this._expandableContent.classList.toggle('pf-m-expanded', expanded);
    } else {
      this._expandableContent.hidden = true;
    }
  }
}

if (!customElements.get('pfv6-toolbar')) {
  customElements.define('pfv6-toolbar', Pfv6Toolbar);
}
