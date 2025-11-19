import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Page Main Content
 *
 * @attr {string} id - ID for the main content area (defaults to "main-content")
 *
 * @slot - Default slot for main content
 * @customElement pf-v6-page-main
 */
class PfV6PageMain extends CemElement {
  static observedAttributes = ['id'];
  static is = 'pf-v6-page-main';

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

  static {
    customElements.define(this.is, this);
  }
}

