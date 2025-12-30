import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A simple button component
 * @element my-button
 */
@customElement('my-button')
export class MyButton extends LitElement {
  /**
   * Button variant
   */
  @property({ type: String })
  variant: 'primary' | 'secondary' = 'primary';

  /**
   * Button disabled state
   */
  @property({ type: Boolean })
  disabled = false;

  render() {
    return html`
      <button ?disabled=${this.disabled}>
        <slot></slot>
      </button>
    `;
  }
}
