import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-pf-v6-page-sidebar.css' with { type: 'css' };

/**
 * PatternFly v6 Page Sidebar
 *
 * @slot - Default slot for sidebar content (typically cem-pf-v6-navigation)
 */
@customElement('cem-pf-v6-page-sidebar')
export class PfV6PageSidebar extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  @property({ type: Boolean, reflect: true })
  accessor collapsed = false;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  constructor() {
    super();
    this.#internals.role = 'navigation';
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('collapsed')) {
      this.expanded = !this.collapsed;
    } else if (changed.has('expanded')) {
      this.collapsed = !this.expanded;
    }
  }

  render() {
    return html`
      <div id="body">
        <slot></slot>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-pf-v6-page-sidebar': PfV6PageSidebar;
  }
}
