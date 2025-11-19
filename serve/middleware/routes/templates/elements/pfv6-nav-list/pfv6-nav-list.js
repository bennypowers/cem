import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation List
 *
 * @slot - Default slot for nav-item or nav-group elements
 */
class Pfv6NavList extends CemElement {
  async afterTemplateLoaded() {
    this.setAttribute('role', 'list');
  }
}

customElements.define('pfv6-nav-list', Pfv6NavList);
