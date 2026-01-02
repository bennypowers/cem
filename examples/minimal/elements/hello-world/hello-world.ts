import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A simple greeting component that demonstrates the basics of custom elements.
 *
 * @summary A minimal custom element example
 * @slot - Default slot for additional content after the greeting
 *
 * @cssprop --hello-world-color - Color of the greeting text
 * @csspart greeting - The paragraph element containing the greeting
 */
@customElement('hello-world')
export class HelloWorld extends LitElement {
  /**
   * The name to greet
   * @type {string}
   */
  @property() name = 'World';

  render() {
    return html`<p part="greeting">Hello, ${this.name}! <slot></slot></p>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'hello-world': HelloWorld;
  }
}
