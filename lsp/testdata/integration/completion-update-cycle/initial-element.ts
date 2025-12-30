import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('test-element')
export class TestElement extends LitElement {
  /** The state property of the element */
  @property() prop: 'one' | 'two' | 'three' = 'one';

  render() {
    return html`<div class="element element-${this.prop}"><slot></slot></div>`;
  }
}