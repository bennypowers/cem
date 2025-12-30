import { customElement } from 'lit/decorators/custom-element.js';
import { LitElement, html } from 'lit';

// Import RhTab (re-exported)
import { RhTab } from './rh-tab.js';

// Import RhTabPanel (side-effect only - not re-exported but makes element available)
import { RhTabPanel } from './rh-tab-panel.js';

/**
 * Tabs component that manages tab and panel relationships
 */
@customElement('rh-tabs')
export class RhTabs extends LitElement {
  render() {
    return html`<slot></slot>`;
  }
}

// Re-export RhTab so it's available when importing rh-tabs
export { RhTab };

// Note: RhTabPanel is NOT re-exported, but importing this module
// makes rh-tab-panel available as a side effect

declare global {
  interface HTMLElementTagNameMap {
    'rh-tabs': RhTabs;
  }
}