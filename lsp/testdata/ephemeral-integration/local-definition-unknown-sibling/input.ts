import { LitElement, html, customElement } from 'lit';

@customElement('test-el')
export class TestEl extends LitElement {
  render() {
    return html`<unknown-el></unknown-el>`;
  }
}
