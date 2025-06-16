import { LitElement, html } from 'lit-element';

export class NoDecoratorElement extends LitElement {
  render() {
    return html`<p>No decorator used</p>`;
  }
}

customElements.define('no-decorator-element', NoDecoratorElement);
