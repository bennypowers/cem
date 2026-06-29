import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <test-component @unknownEvent=${this.handler}></test-component>
    `;
  }
}
