import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Page Main Content
 *
 * @attr {string} id - ID for the main content area (defaults to "main-content")
 *
 * @slot - Default slot for main content
 */
class PfV6PageMain extends CemElement {
  static observedAttributes = ['id'];

  #main;

  async afterTemplateLoaded() {
    this.#main = this.shadowRoot.querySelector('.pf-v6-c-page__main');
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#main) return;

    // Use provided ID or default to main-content
    const id = this.getAttribute('id') || 'main-content';
    this.#main.setAttribute('id', id);
  }
}

customElements.define('pf-v6-page-main', PfV6PageMain);
