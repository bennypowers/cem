import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './pf-v6-form-label.css' with { type: 'css' };

/**
 * PatternFly v6 Form Label Component
 *
 * Automatically associates with control in parent pf-v6-form-group by:
 * - Setting aria-label on the control from label text content
 * - Forwarding clicks to focus the control
 *
 * @slot - Label text content
 */
@customElement('pf-v6-form-label')
export class PfV6FormLabel extends LitElement {
  static styles = styles;

  render() {
    return html`
      <label @click=${this.#handleClick}>
        <span><slot @slotchange=${this.#updateControlLabel}></slot></span>
      </label>
    `;
  }

  connectedCallback() {
    super.connectedCallback();
    // Defer initial label setup until next frame so slotted content is available
    requestAnimationFrame(() => this.#updateControlLabel());
  }

  #handleClick(e: Event) {
    e.preventDefault();
    const control = this.#getControl();
    if (control && typeof control.focus === 'function') {
      control.focus();
    }
  }

  /**
   * Updates the aria-label on the control's internal input element.
   *
   * Supports two approaches:
   * 1. Preferred: Control implements setAccessibleLabel(text) method (CemFormControl API)
   * 2. Fallback: Direct shadowRoot access for legacy controls
   *
   * Compatible controls should extend CemFormControl or implement setAccessibleLabel().
   */
  #updateControlLabel = () => {
    const control = this.#getControl();
    const labelText = this.textContent?.trim() || '';

    if (!control || !labelText) {
      return;
    }

    // Preferred: Use standard API if available
    if (typeof (control as any).setAccessibleLabel === 'function') {
      (control as any).setAccessibleLabel(labelText);
      return;
    }

    // Fallback: Direct shadowRoot access for legacy controls
    const shadowRoot = control.shadowRoot;
    if (!shadowRoot) {
      console.warn('Control has no shadow root and no setAccessibleLabel method');
      return;
    }

    const internalInput = shadowRoot.querySelector('input, select, textarea');
    if (internalInput) {
      internalInput.setAttribute('aria-label', labelText);
    } else {
      console.warn('Could not find internal input in control shadow DOM');
    }
  };

  #getControl(): Element | null {
    // Find the control element in the parent form group
    const formGroup = this.closest('pf-v6-form-group');
    if (!formGroup) {
      console.warn('pf-v6-form-label must be inside pf-v6-form-group');
      return null;
    }

    // Get the control slot (default slot content)
    const control = formGroup.querySelector(':not([slot])');
    if (!control) {
      console.warn('No control found in pf-v6-form-group');
      return null;
    }

    return control;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-form-label': PfV6FormLabel;
  }
}
