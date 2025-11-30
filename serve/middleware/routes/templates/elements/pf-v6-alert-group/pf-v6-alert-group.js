import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly Alert Group component
 *
 * Container for managing multiple alerts with optional toast positioning
 * and overflow handling.
 *
 * @slot - Default slot for pf-v6-alert elements
 */
export class PfAlertGroup extends CemElement {
  static is = 'pf-v6-alert-group';

  static observedAttributes = ['toast', 'live-region'];

  get toast() {
    return this.hasAttribute('toast');
  }

  set toast(value) {
    this.toggleAttribute('toast', !!value);
  }

  async afterTemplateLoaded() {
    // Set up aria-live for live region mode
    if (this.hasAttribute('live-region')) {
      this.setAttribute('aria-live', 'polite');
      this.setAttribute('aria-atomic', 'false');
    }

    // Listen for close events from child alerts
    this.addEventListener('close', (e) => {
      const alert = e.target;
      if (alert && alert.tagName === 'PF-V6-ALERT') {
        this.#removeAlert(alert);
      }
    });
  }

  /**
   * Add an alert to the group
   * @param {HTMLElement} alert - Alert element to add
   */
  addAlert(alert) {
    // Wrap in alert group item for toast animations
    if (this.toast) {
      const item = document.createElement('div');
      item.className = 'pf-v6-c-alert-group__item';
      item.appendChild(alert);

      // Add incoming animation class
      item.classList.add('pf-m-incoming');
      this.appendChild(item);

      // Remove incoming class after animation
      requestAnimationFrame(() => {
        item.classList.remove('pf-m-incoming');
      });
    } else {
      this.appendChild(alert);
    }
  }

  /**
   * Remove an alert from the group with animation
   * @param {HTMLElement} alert - Alert element to remove
   */
  #removeAlert(alert) {
    const item = this.toast ? alert.closest('.pf-v6-c-alert-group__item') : alert;

    if (!item) return;

    if (this.toast) {
      // Add outgoing animation class
      item.classList.add('pf-m-outgoing');

      // Remove after animation completes
      item.addEventListener('transitionend', () => {
        item.remove();
      }, { once: true });
    } else {
      item.remove();
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
