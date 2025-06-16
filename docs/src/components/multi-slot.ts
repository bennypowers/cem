import { LitElement, html, customElement } from 'lit-element';

@customElement('multi-slot')
export class MultiSlot extends LitElement {
  render() {
    return html`
      <slot name="header"></slot>
      <slot></slot>
      <slot name="footer"></slot>
    `;
  }
}
