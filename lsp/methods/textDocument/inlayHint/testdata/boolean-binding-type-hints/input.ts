import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <my-button ?disabled=${this.disabled}></my-button>
    `;
  }
}
