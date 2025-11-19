import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Skip to Content
 *
 * @attr {string} href - Target anchor ID to skip to (e.g., "#main-content")
 *
 * @slot - Default slot for skip link text (defaults to "Skip to content")
 */
class Pfv6SkipToContent extends CemElement {
  static observedAttributes = ['href'];

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
}

customElements.define('pfv6-skip-to-content', Pfv6SkipToContent);
