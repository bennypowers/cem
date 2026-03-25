import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-card.css' with { type: 'css' };

/**
 * PatternFly v6 Card
 *
 * @slot - Default slot for body content
 * @slot title - Optional title content (h1-h6 headings are automatically styled)
 * @slot header - Optional header content (displayed after title)
 * @slot footer - Optional footer content
 */
@customElement('pf-v6-card')
export class PfV6Card extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor variant?: 'secondary' | 'plain';

  @property({ type: Boolean, reflect: true })
  accessor compact = false;

  @property({ type: Boolean, reflect: true, attribute: 'full-height' })
  accessor fullHeight = false;

  render() {
    return html`
      <header id="header">
        <slot name="header"></slot>
        <div id="title">
          <slot name="title"></slot>
        </div>
      </header>
      <div id="body">
        <slot></slot>
      </div>
      <footer id="footer">
        <slot name="footer"></slot>
      </footer>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-card': PfV6Card;
  }
}
