import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <my-button .label=${this.label} ?disabled=${this.off} @change=${this.onChange}></my-button>
    `;
  }
}
