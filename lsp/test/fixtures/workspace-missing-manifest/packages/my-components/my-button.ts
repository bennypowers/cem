import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A simple button component
 */
@customElement('my-button')
export class MyButton extends LitElement {
  /**
   * Button label text
   */
  @property({ type: String }) label = '';

  render() {
    return html`<button>${this.label}</button>`;
  }
}
