/* PatternFly v6 Toolbar Component */

import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

export class Pfv6Toolbar extends HTMLElement {
  static observedAttributes = ['sticky', 'full-height', 'background', 'expandable', 'expanded'];

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('.pf-v6-c-toolbar')) {
      await this.#populateShadowRoot();
    }
    this._toolbar = this.shadowRoot.querySelector('.pf-v6-c-toolbar');
    this._expandableContent = this.shadowRoot.querySelector('.pf-v6-c-toolbar__expandable-content');
    this._updateClasses();
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
    if (oldValue !== newValue) {
      this._updateClasses();
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

  get background() {
    return this.getAttribute('background') || '';
  }

  set background(value) {
    if (value) {
      this.setAttribute('background', value);
    } else {
      this.removeAttribute('background');
    }
  }

  get expandable() {
    return this.hasAttribute('expandable');
  }

  set expandable(value) {
    if (value) {
      this.setAttribute('expandable', '');
    } else {
      this.removeAttribute('expandable');
    }
  }

  get expanded() {
    return this.hasAttribute('expanded');
  }

  set expanded(value) {
    if (value) {
      this.setAttribute('expanded', '');
    } else {
      this.removeAttribute('expanded');
    }
  }

  _updateClasses() {
    if (!this._toolbar) return;

    // Remove old modifier classes
    this._toolbar.classList.remove('pf-m-sticky', 'pf-m-full-height', 'pf-m-primary', 'pf-m-secondary', 'pf-m-no-background', 'pf-m-static');

    // Add sticky modifier
    if (this.sticky) {
      this._toolbar.classList.add('pf-m-sticky');
    }

    // Add full-height modifier
    if (this.fullHeight) {
      this._toolbar.classList.add('pf-m-full-height');
    }

    // Add background modifier
    if (this.background === 'primary') {
      this._toolbar.classList.add('pf-m-primary');
    } else if (this.background === 'secondary') {
      this._toolbar.classList.add('pf-m-secondary');
    } else if (this.background === 'no-background') {
      this._toolbar.classList.add('pf-m-no-background');
    }

    // Handle expandable content
    if (this._expandableContent) {
      if (this.expandable) {
        this._expandableContent.hidden = false;
        this._expandableContent.classList.toggle('pf-m-expanded', this.expanded);
      } else {
        this._expandableContent.hidden = true;
      }
    }
  }
}

if (!customElements.get('pfv6-toolbar')) {
  customElements.define('pfv6-toolbar', Pfv6Toolbar);
}
