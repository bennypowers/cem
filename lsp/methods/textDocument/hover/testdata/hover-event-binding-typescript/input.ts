import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <test-component @change=${this.onChange}></test-component>
      <!--             ^cursor -->
    `;
  }
}
