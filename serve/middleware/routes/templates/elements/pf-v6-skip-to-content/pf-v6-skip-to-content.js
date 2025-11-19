import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Skip to Content
 *
 * @attr {string} href - Target anchor ID to skip to (e.g., "#main-content")
 *
 * @slot - Default slot for skip link text (defaults to "Skip to content")
 * @customElement pf-v6-skip-to-content
 */
class PfV6SkipToContent extends CemElement {
  static observedAttributes = ['href'];
  static is = 'pf-v6-skip-to-content';

  #link;

  async afterTemplateLoaded() {
    this.#link = this.shadowRoot.querySelector('#link');
    this.#syncAttributes();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#link) return;

    // Default to #main-content if no href is specified
    const href = this.getAttribute('href') || '#main-content';
    this.#link.setAttribute('href', href);
  }

  static {
    customElements.define(this.is, this);
  }
}

