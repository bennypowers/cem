---
cursor:
  line: 5
  character: 23
---
import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <test-component @change=${this.onChange}></test-component>
    `;
  }
}
