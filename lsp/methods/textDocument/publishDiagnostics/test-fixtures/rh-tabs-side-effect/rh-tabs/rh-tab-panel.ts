import { customElement } from 'lit/decorators/custom-element.js';
import { LitElement, html } from 'lit';

/**
 * Tab panel component - registered as side effect when module is imported
 */
@customElement('rh-tab-panel')
export class RhTabPanel extends LitElement {
  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'rh-tab-panel': RhTabPanel;
  }
}