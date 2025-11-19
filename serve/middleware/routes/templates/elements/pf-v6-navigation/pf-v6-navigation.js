import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation Container
 *
 * @attr {string} aria-label - Accessible label for navigation
 * @attr {boolean} inset - Add horizontal padding
 *
 * @slot - Default slot for nav-list
 */
class PfV6Navigation extends CemElement {
  static observedAttributes = ['aria-label', 'inset'];

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
}

customElements.define('pf-v6-navigation', PfV6Navigation);
