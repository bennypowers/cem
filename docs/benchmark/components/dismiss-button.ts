import { LitElement, html, css, customElement } from 'lit-element';

/**
 * @element dismiss-button
 * @summary Button for dismissing dialogs or banners.
 */
@customElement('dismiss-button')
export class DismissButton extends LitElement {
  static styles = css`
    button {
      background: transparent;
      border: none;
      font-size: 1.5em;
      cursor: pointer;
      color: #888;
    }
  `;

  render() {
    return html`<button aria-label="Dismiss" @click=${this.#dismiss}>&times;</button>`;
  }

  #dismiss() {
    this.dispatchEvent(new CustomEvent('dismiss'));
  }
}