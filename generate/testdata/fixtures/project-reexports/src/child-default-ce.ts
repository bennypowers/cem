import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('default-child-element')
export default class DefaultChildElement extends LitElement {
  render() { return html`<slot></slot>`; }
}
