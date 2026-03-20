import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-toolbar.css' with { type: 'css' };

type ToolbarColorVariant = 'primary' | 'secondary' | 'no-background';

/**
 * PatternFly v6 Toolbar
 *
 * @slot - Default slot for toolbar content (groups and items)
 * @slot expandable - Slot for expandable content
 */
@customElement('pf-v6-toolbar')
export class PfV6Toolbar extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor expandable = false;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  @property({ type: Boolean, reflect: true })
  accessor sticky = false;

  @property({ type: Boolean, reflect: true, attribute: 'full-height' })
  accessor fullHeight = false;

  @property({ reflect: true, attribute: 'color-variant' })
  accessor colorVariant?: ToolbarColorVariant;

  render() {
    return html`
      <div id="content">
        <div id="content-section">
          <slot></slot>
        </div>
        <div id="expandable-content"
             part="expandable-content">
          <slot name="expandable"></slot>
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-toolbar': PfV6Toolbar;
  }
}
