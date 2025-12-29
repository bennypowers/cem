import { customElement } from 'lit/decorators/custom-element.js';
import { LitElement, html } from 'lit';

/**
 * Individual tab component
 */
@customElement('rh-tab')
export class RhTab extends LitElement {
  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'rh-tab': RhTab;
  }
}