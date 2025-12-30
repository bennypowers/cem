import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('dialog-element')
export class DialogElement extends LitElement {
  render() {
    return html`
      <div class="dialog">
        <slot name="title"></slot>
        <slot name="content"></slot>
      </div>
    `;
  }
}