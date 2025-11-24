import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Form Component
 *
 * @attr {boolean} horizontal - Enables horizontal layout with grid
 * @customElement pf-v6-form
 */
class PfV6Form extends CemElement {
  static is = 'pf-v6-form';

  static {
    customElements.define(this.is, this);
  }
}
