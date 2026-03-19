import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './pf-v6-form-group.css' with { type: 'css' };

/**
 * PatternFly v6 Form Group Component
 *
 * Groups a label, form control, and optional helper text together.
 *
 * @slot label - Label content (typically pf-v6-form-label)
 * @slot - Default slot for form control element
 * @slot helper - Helper text
 */
@customElement('pf-v6-form-group')
export class PfV6FormGroup extends LitElement {
  static styles = styles;

  render() {
    return html`
      <div id="label">
        <slot name="label"></slot>
      </div>
      <div id="control">
        <slot></slot>
      </div>
      <div id="helper-text">
        <slot name="helper"></slot>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-form-group': PfV6FormGroup;
  }
}
