import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-skip-to-content.css' with { type: 'css' };

/**
 * PatternFly v6 Skip to Content
 *
 * @slot - Default slot for skip link text (defaults to "Skip to content")
 */
@customElement('pf-v6-skip-to-content')
export class PfV6SkipToContent extends LitElement {
  static styles = styles;

  /** Target anchor ID to skip to (e.g., "#main-content") */
  @property({ reflect: true })
  accessor href = '#main-content';

  render() {
    return html`
      <div id="wrapper">
        <a id="link"
           part="link"
           href="${this.href}">
          <span id="text">
            <slot>Skip to content</slot>
          </span>
        </a>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-skip-to-content': PfV6SkipToContent;
  }
}
