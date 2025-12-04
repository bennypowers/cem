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

    // If delay is 0 or negative, show connecting immediately
    if (this.#remainingMs <= 0) {
      this.#retryInfo.textContent = `Attempt #${attempt}. Connecting...`;
      return;
    }

    // Update immediately
    const seconds = Math.ceil(this.#remainingMs / 1000);
    this.#retryInfo.textContent = `Attempt #${attempt}. Retrying in ${seconds}s...`;

    // Update countdown every 100ms for smooth display
    this.#countdownInterval = setInterval(() => {
      this.#remainingMs -= 100;

      // Use < 100 instead of <= 0 to handle cases where remaining time
      // is less than one interval period (avoids showing "Retrying in 1s" when < 100ms left)
      if (this.#remainingMs < 100) {
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
