import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A simple card component
 * @element my-card
 */
@customElement('my-card')
export class MyCard extends LitElement {
  /**
   * Card title
   */
  @property({ type: String })
  title = '';

  /**
   * Card variant
   */
  @property({ type: String })
  variant: 'default' | 'outlined' = 'default';

  static styles = css`
    :host {
      display: block;
      border: 1px solid #ccc;
      padding: 1rem;
    }
  `;

  render() {
    return html`
      <h2>${this.title}</h2>
      <slot></slot>
    `;
  }
}
