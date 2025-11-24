import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Form Group Component
 *
 * @slot label - Label content (typically pf-v6-form-label)
 * @slot control - Form control element
 * @slot helper - Helper text
 * @customElement pf-v6-form-group
 */
class PfV6FormGroup extends CemElement {
  static is = 'pf-v6-form-group';

  static {
    customElements.define(this.is, this);
  }
}
