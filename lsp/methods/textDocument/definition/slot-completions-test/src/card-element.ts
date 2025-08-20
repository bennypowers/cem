import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('card-element')
export class CardElement extends LitElement {
  render() {
    return html`
      <div class="card">
        <slot name="header"></slot>
        <slot></slot>
        <slot name="footer"></slot>
        <slot name="actions"></slot>
      </div>
    `;
  }
}