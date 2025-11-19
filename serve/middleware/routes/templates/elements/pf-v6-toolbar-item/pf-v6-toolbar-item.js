/* PatternFly v6 Toolbar Item Component */

import { CemElement } from '/__cem/cem-element.js';

/**
 * @customElement pf-v6-toolbar-item
 */
export class PfV6ToolbarItem extends CemElement {
  static observedAttributes = ['variant'];
  static is = 'pf-v6-toolbar-item';

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

  static {
    customElements.define(this.is, this);
  }
}

