import { LitElement, html, customElement, property } from 'lit';

/**
 * A greeting element for testing ephemeral synthesis
 * @summary Displays a greeting
 */
@customElement('test-greeting')
export class TestGreeting extends LitElement {
  /** The name to greet */
  @property()
  name: string = 'World';

  /** Whether the greeting is enthusiastic */
  @property({ type: Boolean })
  enthusiastic: boolean = false;

  render() {
    return html`<p>Hello, ${this.name}!</p>`;
  }
}
