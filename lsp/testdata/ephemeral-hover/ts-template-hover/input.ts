import { LitElement, html, customElement, property } from 'lit';

/**
 * A counter element
 * @summary Increments a count
 */
@customElement('click-counter')
export class ClickCounter extends LitElement {
  /** Number of clicks */
  @property({ type: Number })
  count: number = 0;

  render() {
    return html`
      <click-counter count="${this.count}"></click-counter>
    `;
  }
}
