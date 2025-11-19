import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation Container
 *
 * @attr {string} aria-label - Accessible label for navigation
 * @attr {boolean} inset - Add horizontal padding
 *
 * @slot - Default slot for nav-list
 * @customElement pf-v6-navigation
 */
class PfV6Navigation extends CemElement {
  static observedAttributes = ['aria-label', 'inset'];
  static is = 'pf-v6-navigation';

  #nav;

  async afterTemplateLoaded() {
    this.#nav = this.shadowRoot.querySelector('nav');
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#nav) return;

    if (this.hasAttribute('aria-label')) {
      this.#nav.setAttribute('aria-label', this.getAttribute('aria-label'));
    }
  }

  static {
    customElements.define(this.is, this);
  }
}

