import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-toolbar-group.css' with { type: 'css' };

type ToolbarGroupVariant =
  | 'filter-group'
  | 'label-group'
  | 'action-group'
  | 'action-group-plain'
  | 'action-group-inline'
  | 'overflow-container';

/**
 * PatternFly v6 Toolbar Group
 *
 * @slot - Default slot for toolbar items
 */
@customElement('pf-v6-toolbar-group')
export class PfV6ToolbarGroup extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor variant?: ToolbarGroupVariant;

  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-toolbar-group': PfV6ToolbarGroup;
  }
}
