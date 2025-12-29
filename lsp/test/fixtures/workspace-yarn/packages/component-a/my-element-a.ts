import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * Component A element
 */
@customElement('my-element-a')
export class MyElementA extends LitElement {
  /**
   * Label text
   */
  @property() label = '';

  render() {
    return html`<div>${this.label}</div>`;
  }
}
