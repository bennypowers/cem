import { LitElement, css, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A simple greeting component that demonstrates the basics of custom elements.
 *
 * @summary A minimal custom element example
 * @slot - Default slot for additional content after the greeting
 *
 * @cssprop {<color>} --hello-world-color - Color of the greeting text
 * @csspart greeting - The paragraph element containing the greeting
 *
 * @demo elements/hello-world/demos/basic.html Basic usage
 */
@customElement('hello-world')
export class HelloWorld extends LitElement {
  static styles = css`
    p {
      color: var(--hello-world-color, inherit);
    }
  `;

  /**
   * The name to greet
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
