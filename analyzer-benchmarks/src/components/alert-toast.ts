import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element alert-toast
 * @summary An alert toast component for transient notifications.
 * @csspart container - The main toast container.
 */
@customElement('alert-toast')
export class AlertToast extends LitElement {
  /** Toast message text */
  @property({ type: String }) message = '';
  /** Toast visibility */
  @property({ type: Boolean, reflect: true }) open = false;

  static styles = css`
    :host { display: block; }
    .toast {
      background: #333; color: white; padding: 1em 2em; border-radius: 4px;
      position: fixed; bottom: 2em; left: 2em; opacity: 0.9;
      transition: transform 0.2s, opacity 0.2s;
    }
    :host(:not([open])) .toast { display: none; }
  `;

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('click', this.hide);
  }

  hide = () => { this.open = false; };

  render() {
    return html`<div part="container" class="toast">${this.message}</div>`;
  }
}