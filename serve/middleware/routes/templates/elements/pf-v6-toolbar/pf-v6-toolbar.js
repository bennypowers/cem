/* PatternFly v6 Toolbar Component */

import { CemElement } from '/__cem/cem-element.js';

/**
 * @customElement pf-v6-toolbar
 */
export class PfV6Toolbar extends CemElement {
  static is = 'pf-v6-toolbar';

  static observedAttributes = [
    'expandable',
    'expanded',
    'sticky',
    'full-height',
    'color-variant',
  ];

  get expandable() { return this.hasAttribute('expandable'); }
  set expandable(value) { this.toggleAttribute('expandable', !!value); }

  get expanded() { return this.hasAttribute('expanded'); }
  set expanded(value) { this.toggleAttribute('expanded', !!value); }

  get sticky() { return this.hasAttribute('sticky'); }
  set sticky(value) { this.toggleAttribute('sticky', !!value); }

  get fullHeight() { return this.hasAttribute('full-height'); }
  set fullHeight(value) { this.toggleAttribute('full-height', !!value); }

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

  static {
    customElements.define(this.is, this);
  }
}
