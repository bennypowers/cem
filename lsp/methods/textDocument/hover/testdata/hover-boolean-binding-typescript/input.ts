import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <test-component ?disabled=${this.disabled}></test-component>
    `;
  }
}
