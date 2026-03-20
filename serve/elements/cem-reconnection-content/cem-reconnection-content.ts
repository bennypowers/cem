import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { state } from 'lit/decorators/state.js';

import styles from './cem-reconnection-content.css' with { type: 'css' };

/**
 * Reconnection dialog content component.
 * Displays retry information and countdown when server connection is lost.
 *
 * @customElement cem-reconnection-content
 */
@customElement('cem-reconnection-content')
export class CemReconnectionContent extends LitElement {
  static styles = styles;

  @state()
  accessor #retryText = '';

  #countdownInterval: ReturnType<typeof setInterval> | null = null;
  #remainingMs = 0;

  render() {
    return html`
      <p>
        The connection to the development server was lost.
        Automatically retrying connection...
      </p>
      <div id="retry-info">${this.#retryText}</div>
    `;
  }

  /** Clear the countdown interval */
  clearCountdown(): void {
    if (this.#countdownInterval) {
      clearInterval(this.#countdownInterval);
      this.#countdownInterval = null;
    }
  }

  /**
   * Update the retry information with countdown
   * @param attempt - Current retry attempt number
   * @param delay - Delay in milliseconds until next retry
   */
  updateRetryInfo(attempt: number, delay: number): void {
    this.clearCountdown();

    this.#remainingMs = delay;

    // If delay is 0 or negative, show connecting immediately
    if (this.#remainingMs <= 0) {
      this.#retryText = `Attempt #${attempt}. Connecting...`;
      return;
    }

    // Update immediately
    const seconds = Math.ceil(this.#remainingMs / 1000);
    this.#retryText = `Attempt #${attempt}. Retrying in ${seconds}s...`;

    // Update countdown every 100ms for smooth display
    this.#countdownInterval = setInterval(() => {
      this.#remainingMs -= 100;

      // Use < 100 instead of <= 0 to handle cases where remaining time
      // is less than one interval period (avoids showing "Retrying in 1s" when < 100ms left)
      if (this.#remainingMs < 100) {
        this.clearCountdown();
        this.#retryText = `Attempt #${attempt}. Connecting...`;
        return;
      }

      const seconds = Math.ceil(this.#remainingMs / 1000);
      this.#retryText = `Attempt #${attempt}. Retrying in ${seconds}s...`;
    }, 100);
  }

  disconnectedCallback(): void {
    super.disconnectedCallback();
    this.clearCountdown();
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-reconnection-content': CemReconnectionContent;
  }
}
