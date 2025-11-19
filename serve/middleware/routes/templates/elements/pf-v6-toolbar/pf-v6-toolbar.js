/* PatternFly v6 Toolbar Component */

import { CemElement } from '/__cem/cem-element.js';

export class PfV6Toolbar extends CemElement {
  static observedAttributes = ['sticky', 'full-height', 'color-variant', 'expandable', 'expanded'];

  #expandableContent;

  async afterTemplateLoaded() {
    this.#expandableContent = this.shadowRoot.querySelector('.pf-v6-c-toolbar__expandable-content');
    this.#updateExpandableContent();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    switch (name) {
      case 'expandable':
      case 'expanded':
        if (oldValue !== newValue) {
          this.#updateExpandableContent();
        }
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

  #updateExpandableContent() {
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

customElements.define('pf-v6-toolbar', PfV6Toolbar);
