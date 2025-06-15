import { LitElement, html, css, customElement } from 'lit-element';

/**
 * @element load-spinner
 * @summary Animated loading spinner.
 */
@customElement('load-spinner')
export class LoadSpinner extends LitElement {
  private _intervalId?: number;

  static styles = css`
    .spinner {
      border: 4px solid #eee;
      border-top: 4px solid #2196f3;
      border-radius: 50%;
      width: 32px; height: 32px;
      animation: spin 1s linear infinite;
      margin: 1em auto;
    }
    @keyframes spin { 100% { transform: rotate(360deg); } }
  `;

  connectedCallback() {
    super.connectedCallback();
    // Example: a fake loader tick
    this._intervalId = window.setInterval(() => this.requestUpdate(), 1000);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.clearInterval(this._intervalId);
  }

  render() {
    return html`<div class="spinner"></div>`;
  }
}