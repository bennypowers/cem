import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import '@test/component-a/my-element-a.js';

/**
 * Component B element that uses Component A
 */
@customElement('my-element-b')
export class MyElementB extends LitElement {
  /**
   * Title text
   */
  @property() title = '';

  render() {
    return html`
      <div>
        <h1>${this.title}</h1>
        <my-element-a label="From Component B"></my-element-a>
      </div>
    `;
  }
}
