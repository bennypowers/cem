import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <test-component prop="value"></test-component>
    `;
  }
}