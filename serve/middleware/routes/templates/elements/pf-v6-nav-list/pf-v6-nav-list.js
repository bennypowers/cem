import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation List
 *
 * @slot - Default slot for nav-item or nav-group elements
 */
class PfV6NavList extends CemElement {
  async afterTemplateLoaded() {
    this.setAttribute('role', 'list');
  }
}

customElements.define('pf-v6-nav-list', PfV6NavList);
