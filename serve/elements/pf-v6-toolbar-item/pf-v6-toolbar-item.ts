import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-toolbar-item.css' with { type: 'css' };

type ToolbarItemVariant =
  | 'label'
  | 'pagination'
  | 'overflow-container'
  | 'expand-all';

/**
 * PatternFly v6 Toolbar Item
 *
 * @slot - Default slot for item content
 */
@customElement('pf-v6-toolbar-item')
export class PfV6ToolbarItem extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor variant?: ToolbarItemVariant;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-toolbar-item': PfV6ToolbarItem;
  }
}
