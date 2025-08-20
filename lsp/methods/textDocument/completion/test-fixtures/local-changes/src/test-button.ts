import {LitElement, html, css} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('test-button')
export class TestButton extends LitElement {
  /** The variant of the button */
  @property() variant: 'primary' | 'secondary' = 'primary';

  /** Whether the button is disabled */
  @property({type: Boolean}) disabled = false;

  render() {
    return html`<button class="${this.variant}" ?disabled="${this.disabled}"><slot></slot></button>`;
  }
}
