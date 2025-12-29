import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation List
 *
 * @slot - Default slot for nav-item or nav-group elements
 * @customElement pf-v6-nav-list
 */
class PfV6NavList extends CemElement {
  static is = 'pf-v6-nav-list';

  #internals = this.attachInternals();

  constructor() {
    super();
    this.#internals.role = 'list';
  }

  static {
    customElements.define(this.is, this);
  }
}
