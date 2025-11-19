import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Page Sidebar
 *
 * @attr {boolean} collapsed - Whether the sidebar is collapsed
 * @attr {boolean} expanded - Whether the sidebar is expanded
 *
 * @slot - Default slot for sidebar content (typically pfv6-navigation)
 */
class Pfv6PageSidebar extends CemElement {
  static observedAttributes = ['collapsed', 'expanded'];

  #internals = this.attachInternals();

  constructor() {
    super();
    this.#internals.role = 'navigation';
  }

  get collapsed() {
    return this.hasAttribute('collapsed');
  }

  set collapsed(value) {
    this.toggleAttribute('collapsed', !!value);
    this.toggleAttribute('expanded', !value);
  }

  get expanded() {
    return this.hasAttribute('expanded');
  }

  set expanded(value) {
    this.toggleAttribute('collapsed', !value);
    this.toggleAttribute('expanded', !!value);
  }
}

customElements.define('pfv6-page-sidebar', Pfv6PageSidebar);
