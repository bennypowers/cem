import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element toast-stack
 * @summary Stack of toasts.
 */
@customElement('toast-stack')
export class ToastStack extends LitElement {
  @property({ type: Array }) messages: string[] = [];

  static styles = css`
    .toast-stack {
      position: fixed; right: 1em; bottom: 1em;
      display: flex; flex-direction: column; gap: 0.5em;
      z-index: 100;
    }
    .toast {
      background: #444; color: white; border-radius: 6px; padding: 0.5em 1.5em;
      box-shadow: 0 2px 4px rgba(0,0,0,0.2);
    }
  `;

  clear() {
    this.messages = [];
  }

  render() {
    return html`
      <div class="toast-stack">
        ${this.messages.map(m => html`<div class="toast">${m}</div>`)}
      </div>
    `;
  }
}