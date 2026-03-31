import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('child-element')
export class ChildElement extends LitElement {
  render() { return html`<slot></slot>`; }
}
