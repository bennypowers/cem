import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A test custom element
 */
@customElement('test-element')
export class TestElement extends LitElement {
  @property({ type: String })
  name = 'World';

  render() {
    return html`<p>Hello, ${this.name}!</p>`;
  }
}