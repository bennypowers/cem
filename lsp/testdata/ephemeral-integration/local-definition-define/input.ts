import { LitElement, html } from 'lit';

class TestEl extends LitElement {
  render() {
    return html`<test-el></test-el>`;
  }
}

customElements.define('test-el', TestEl);
