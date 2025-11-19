import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation Group (Nested/Subnav)
 *
 * @attr {boolean} hidden - Whether the group is collapsed
 * @attr {string} aria-labelledby - ID of element labeling this group
 *
 * @slot - Default slot for nav-list containing nav-items
 * @customElement pf-v6-nav-group
 */
class PfV6NavGroup extends CemElement {
  static observedAttributes = ['hidden', 'aria-labelledby'];
  static is = 'pf-v6-nav-group';

  #subnav;

  async afterTemplateLoaded() {
    this.#subnav = this.shadowRoot.querySelector('#subnav');
    this.setAttribute('role', 'region');
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#subnav) return;

    if (this.hasAttribute('aria-labelledby')) {
      this.#subnav.setAttribute('aria-labelledby', this.getAttribute('aria-labelledby'));
    }

    // Handle inert attribute when hidden
    if (this.hasAttribute('hidden')) {
      this.#subnav.setAttribute('inert', '');
    } else {
      this.#subnav.removeAttribute('inert');
    }
  }

  static {
    customElements.define(this.is, this);
  }
}

