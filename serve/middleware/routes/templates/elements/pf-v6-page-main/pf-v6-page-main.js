import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Page Main Content
 *
 * Wraps main content area with proper semantics and skip-to-content support.
 * The id attribute should be set on the host element for skip-to-content links.
 *
 * @slot - Default slot for main content
 * @customElement pf-v6-page-main
 */
class PfV6PageMain extends CemElement {
  static is = 'pf-v6-page-main';

  #internals = this.attachInternals();
  constructor() {
    super();
    this.#internals.role = "main";
  }

  static {
    customElements.define(this.is, this);
  }
}
