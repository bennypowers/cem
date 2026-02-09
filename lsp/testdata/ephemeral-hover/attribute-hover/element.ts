import { LitElement, html, customElement, property } from 'lit';

/**
 * A greeting card element
 * @summary Shows a personalized greeting
 * @slot - Default slot for greeting content
 * @slot header - Header content
 */
@customElement('greeting-card')
export class GreetingCard extends LitElement {
  /** The person to greet */
  @property()
  name: string = 'World';

  /** Visual style variant */
  @property({ attribute: 'color-scheme' })
  colorScheme: 'light' | 'dark' = 'light';

  render() {
    return html`
      <div>
        <slot name="header"></slot>
        <p>Hello, ${this.name}!</p>
        <slot></slot>
      </div>
    `;
  }
}
