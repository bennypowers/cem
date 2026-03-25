import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-tab.css' with { type: 'css' };

/**
 * PatternFly v6 Tab panel content
 *
 * Used as a child of `pf-v6-tabs` to define tab panels.
 * The `title` attribute provides the tab button label.
 *
 * @slot - Default slot for tab panel content
 */
@customElement('pf-v6-tab')
export class PfV6Tab extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  override accessor title = '';

  connectedCallback() {
    super.connectedCallback();
    this.dispatchEvent(new Event('pf-v6-tab-connected', { bubbles: true }));
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.dispatchEvent(new Event('pf-v6-tab-disconnected', { bubbles: true }));
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('title')) {
      this.dispatchEvent(new Event('pf-v6-tab-title-changed', { bubbles: true }));
    }
  }

  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-tab': PfV6Tab;
  }
}
