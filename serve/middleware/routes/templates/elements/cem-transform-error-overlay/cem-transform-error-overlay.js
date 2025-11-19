import { CemElement } from '/__cem/cem-element.js';

/**
 * Transform Error Overlay Component
 * Displays TypeScript transform errors and other server-side compilation errors
 *
 * @attr {boolean} open - Whether the overlay is currently visible
 *
 * @customElement cem-transform-error-overlay
 */
export class CemTransformErrorOverlay extends CemElement {
  static shadowRootOptions = { mode: 'open' };
  static is = 'cem-transform-error-overlay';
  static observedAttributes = ['open'];

  #titleEl;
  #fileEl;
  #messageEl;
  #closeButton;
  #handleKeydown;

  async afterTemplateLoaded() {
    this.#titleEl = this.shadowRoot.getElementById('title');
    this.#fileEl = this.shadowRoot.getElementById('file');
    this.#messageEl = this.shadowRoot.getElementById('message');
    this.#closeButton = this.shadowRoot.getElementById('close');

    if (!this.#closeButton) return;

    // Close button handler
    this.#closeButton.addEventListener('click', () => this.hide());

    // Close on Escape key
    this.#handleKeydown = (e) => {
      if (e.key === 'Escape' && this.hasAttribute('open')) {
        this.hide();
      }
    };
    document.addEventListener('keydown', this.#handleKeydown);
  }

  disconnectedCallback() {
    if (this.#handleKeydown) {
      document.removeEventListener('keydown', this.#handleKeydown);
    }
  }

  /**
   * Show the error overlay
   * @param {string} title - Error title
   * @param {string} message - Error message
   * @param {string} [file=''] - Optional file path where error occurred
   */
  show(title, message, file = '') {
    if (!this.#titleEl || !this.#messageEl || !this.#fileEl) return;

    this.#titleEl.textContent = title;
    this.#messageEl.textContent = message;

    if (file) {
      this.#fileEl.textContent = `File: ${file}`;
    } else {
      this.#fileEl.textContent = '';
    }

    this.setAttribute('open', '');
  }

  /**
   * Hide the error overlay
   */
  hide() {
    this.removeAttribute('open');
  }

  /**
   * Whether the overlay is currently visible
   */
  get open() {
    return this.hasAttribute('open');
  }

  static {
    customElements.define(this.is, this);
  }
}
