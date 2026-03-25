import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import '../pf-v6-button/pf-v6-button.js';

import styles from './cem-transform-error-overlay.css' with { type: 'css' };

/**
 * Transform Error Overlay Component.
 * Displays TypeScript transform errors and other server-side compilation errors.
 *
 * @customElement cem-transform-error-overlay
 */
@customElement('cem-transform-error-overlay')
export class CemTransformErrorOverlay extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor open = false;

  @property()
  accessor title = '';

  @property()
  accessor file = '';

  @property()
  accessor message = '';

  #handleKeydown = (e: KeyboardEvent) => {
    if (e.key === 'Escape' && this.open) {
      this.hide();
    }
  };

  connectedCallback(): void {
    super.connectedCallback();
    document.addEventListener('keydown', this.#handleKeydown);
  }

  disconnectedCallback(): void {
    super.disconnectedCallback();
    document.removeEventListener('keydown', this.#handleKeydown);
  }

  render() {
    return html`
      <div id="overlay-content">
        <div id="header">
          <h2 id="title-container">
            <span id="error-icon">\u26A0\uFE0F</span>
            <span>${this.title}</span>
          </h2>
          <pf-v6-button id="close"
                        variant="plain"
                        @click=${this.hide}>Dismiss</pf-v6-button>
        </div>
        <div id="body">
          <div id="file">${this.file ? `File: ${this.file}` : ''}</div>
          <div id="message">${this.message}</div>
        </div>
        <div id="footer">
          This error will automatically dismiss when the issue is fixed.
        </div>
      </div>
    `;
  }

  /**
   * Show the error overlay.
   * @param title - Error title
   * @param message - Error message
   * @param file - Optional file path where error occurred
   */
  show(title: string, message: string, file = ''): void {
    this.title = title;
    this.message = message;
    this.file = file;
    this.open = true;
  }

  /** Hide the error overlay */
  hide(): void {
    this.open = false;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-transform-error-overlay': CemTransformErrorOverlay;
  }
}
