import { LitElement, html } from 'lit';
import { property } from 'lit/decorators.js';

/**
 * A test card element.
 * @element test-card
 * @slot header - Card header content.
 * @slot - Card body content.
 */
export class TestCardElement extends LitElement {
  @property({ type: Boolean }) elevated = false;

  render() {
    return html`
      <div id="card" elevated="${String(this.elevated)}">
        <header>
          <slot name="header"></slot>
        </header>
        <main>
          <slot></slot>
        </main>
      </div>
    `;
  }
}

customElements.define('test-card', TestCardElement);
