/* PatternFly v6 Toolbar Component */

import { CemElement } from '/__cem/cem-element.js';

/**
 * @customElement pf-v6-toolbar
 */
export class PfV6Toolbar extends CemElement {
  static observedAttributes = ['sticky', 'full-height', 'color-variant', 'expandable', 'expanded'];
  static is = 'pf-v6-toolbar';

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
    this.toggleAttribute('sticky', !!value);
  }

  get fullHeight() {
    return this.hasAttribute('full-height');
  }

  set fullHeight(value) {
    this.toggleAttribute('full-height', !!value);
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
    if (!this.#expandableContent) return;

    const expandable = this.hasAttribute('expandable');
    const expanded = this.hasAttribute('expanded');

    if (expandable) {
      this.#expandableContent.hidden = false;
      this.#expandableContent.classList.toggle('pf-m-expanded', expanded);
    } else {
      this.#expandableContent.hidden = true;
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
