import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <test-component .unknownProp=${this.x}></test-component>
    `;
  }
}
