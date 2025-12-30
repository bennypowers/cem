import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('test-alert')
export class TestAlert extends LitElement {
  /** The state of the alert */
  @property() state: 'info' | 'success' | 'warning' = 'info';

  render() {
    return html`<div class="alert alert-${this.state}"><slot></slot></div>`;
  }
}