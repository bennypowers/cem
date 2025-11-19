/* PatternFly v6 Toolbar Item Component */

import { CemElement } from '/__cem/cem-element.js';

export class Pfv6ToolbarItem extends CemElement {
  static observedAttributes = ['variant'];

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
}

customElements.define('pfv6-toolbar-item', Pfv6ToolbarItem);
