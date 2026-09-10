import { LitElement, html } from 'lit';
import { property } from 'lit/decorators.js';

/**
 * A test button element.
 * @element test-button
 * @slot - The button's default slot content.
 */
export class TestButtonElement extends LitElement {
  @property({ type: String }) variant = 'primary';

  render() {
    return html`
      <button class="btn-${this.variant}">
        <slot></slot>
      </button>
    `;
  }
}

customElements.define('test-button', TestButtonElement);
