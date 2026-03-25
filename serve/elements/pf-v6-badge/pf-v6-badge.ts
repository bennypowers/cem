import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-badge.css' with { type: 'css' };

/**
 * PatternFly v6 Badge
 *
 * @slot - Badge content (typically numbers or short text)
 */
@customElement('pf-v6-badge')
export class PfV6Badge extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor read = false;

  @property({ type: Boolean, reflect: true })
  accessor unread = false;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ attribute: 'screen-reader-text' })
  accessor screenReaderText: string | undefined;

  render() {
    return html`
      <slot></slot>
      ${this.screenReaderText ? html`
        <span id="sr-text">${this.screenReaderText}</span>
      ` : ''}
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-badge': PfV6Badge;
  }
}
