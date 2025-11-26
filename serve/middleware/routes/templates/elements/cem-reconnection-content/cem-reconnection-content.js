import { CemElement } from '/__cem/cem-element.js';

/**
 * Reconnection dialog content component
 * Displays retry information and countdown when server connection is lost
 *
 * @customElement cem-reconnection-content
 */
export class CemReconnectionContent extends CemElement {
  static shadowRootOptions = { mode: 'open' };
  static is = 'cem-reconnection-content';

  #retryInfo;
  #countdownInterval = null;
  #remainingMs = 0;

  async afterTemplateLoaded() {
    this.#retryInfo = this.shadowRoot.getElementById('retry-info');
  }

  /**
   * Clear the countdown interval
   */
  clearCountdown() {
    if (this.#countdownInterval) {
      clearInterval(this.#countdownInterval);
      this.#countdownInterval = null;
    }
  }

  /**
   * Update the retry information with countdown
   * @param {number} attempt - Current retry attempt number
   * @param {number} delay - Delay in milliseconds until next retry
   */
  updateRetryInfo(attempt, delay) {
    this.clearCountdown();

    if (!this.#retryInfo) return;

    this.#remainingMs = delay;

    // Update immediately
    const seconds = Math.ceil(this.#remainingMs / 1000);
    this.#retryInfo.textContent = `Attempt #${attempt}. Retrying in ${seconds}s...`;

    // Update countdown every 100ms for smooth display
    this.#countdownInterval = setInterval(() => {
      this.#remainingMs -= 100;

      if (this.#remainingMs <= 0) {
        this.clearCountdown();
        this.#retryInfo.textContent = `Attempt #${attempt}. Connecting...`;
        return;
      }

      const seconds = Math.ceil(this.#remainingMs / 1000);
      this.#retryInfo.textContent = `Attempt #${attempt}. Retrying in ${seconds}s...`;
    }, 100);
  }

  disconnectedCallback() {
    this.clearCountdown();
  }

  static {
    customElements.define(this.is, this);
  }
}
