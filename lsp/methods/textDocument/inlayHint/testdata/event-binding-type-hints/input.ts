import { html } from 'lit';

export class MyComponent {
  render() {
    return html`
      <my-button @change=${this.onChange}></my-button>
    `;
  }
}
