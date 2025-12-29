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
    if (this.toast) {
      // Add incoming animation class
      alert.classList.add('pf-m-incoming');
      this.appendChild(alert);

      // Remove incoming class after animation
      requestAnimationFrame(() => {
        alert.classList.remove('pf-m-incoming');
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
    if (!alert) return;

    if (this.toast) {
      // Add outgoing animation class
      alert.classList.add('pf-m-outgoing');

      // Remove after animation completes
      alert.addEventListener('transitionend', () => {
        alert.remove();
      }, { once: true });
    } else {
      alert.remove();
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
