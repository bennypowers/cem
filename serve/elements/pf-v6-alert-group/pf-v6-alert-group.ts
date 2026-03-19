import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-alert-group.css' with { type: 'css' };

/**
 * PatternFly Alert Group component
 *
 * Container for managing multiple alerts with optional toast positioning
 * and overflow handling.
 *
 * @slot - Default slot for pf-v6-alert elements
 */
@customElement('pf-v6-alert-group')
export class PfAlertGroup extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor toast = false;

  @property({ type: Boolean, reflect: true, attribute: 'live-region' })
  accessor liveRegion = false;

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('close', this.#onClose);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('close', this.#onClose);
  }

  firstUpdated() {
    if (this.liveRegion) {
      this.setAttribute('aria-live', 'polite');
      this.setAttribute('aria-atomic', 'false');
    }
  }

  render() {
    return html`<slot></slot>`;
  }

  /**
   * Add an alert to the group
   */
  addAlert(alert: HTMLElement) {
    if (this.toast) {
      alert.classList.add('pf-m-incoming');
      this.appendChild(alert);
      requestAnimationFrame(() => {
        alert.classList.remove('pf-m-incoming');
      });
    } else {
      this.appendChild(alert);
    }
  }

  #onClose = (e: Event) => {
    const alert = e.target as HTMLElement;
    if (alert?.tagName === 'PF-V6-ALERT') {
      this.#removeAlert(alert);
    }
  };

  #removeAlert(alert: HTMLElement) {
    if (!alert) return;

    if (this.toast) {
      alert.classList.add('pf-m-outgoing');
      alert.addEventListener('transitionend', () => {
        alert.remove();
      }, { once: true });
    } else {
      alert.remove();
    }
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-alert-group': PfAlertGroup;
  }
}
