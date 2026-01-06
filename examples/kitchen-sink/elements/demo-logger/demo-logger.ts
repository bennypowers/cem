import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { LoggingMixin } from '../../mixins/logging-mixin.js';

/**
 * A debug utility element that demonstrates mixin usage.
 *
 * This element uses the LoggingMixin to add lifecycle logging
 * and debug utilities to track element behavior.
 *
 * @summary Debug element demonstrating mixin pattern
 * @slot - Content to display
 * @csspart container - The main container
 * @cssprop --logger-border-color - Border color (default: #ccc)
 * @cssprop --logger-bg-color - Background color (default: #f5f5f5)
 */
@customElement('demo-logger')
export class DemoLogger extends LoggingMixin(LitElement) {
  static styles = css`
    :host {
      display: block;
      padding: 1rem;
      border: 2px solid var(--logger-border-color, #ccc);
      background: var(--logger-bg-color, #f5f5f5);
      border-radius: 4px;
    }

    #container {
      font-family: monospace;
    }

    .actions {
      margin-top: 1rem;
      display: flex;
      gap: 0.5rem;
    }

    button {
      padding: 0.5rem 1rem;
      cursor: pointer;
    }
  `;

  /**
   * Message to display
   */
  @property() accessor message = 'Logger initialized';

  /**
   * Log a test message to demonstrate logging
   */
  logMessage() {
    this.log('Test message from logMessage()', { timestamp: Date.now() });
  }

  /**
   * Log a test warning
   */
  logWarning() {
    this.warn('Test warning', { severity: 'medium' });
  }

  /**
   * Log a test error
   */
  logError() {
    this.error('Test error', { critical: false });
  }

  render() {
    return html`
      <div id="container" part="container">
        <p>${this.message}</p>
        <label>
          <input
            type="checkbox"
            .checked=${this.enableLogging}
            @change=${(e: Event) => {
              this.enableLogging = (e.target as HTMLInputElement).checked;
              this.log('Logging toggled', { enabled: this.enableLogging });
            }}
          />
          Enable logging
        </label>
        <div class="actions">
          <button @click=${this.logMessage}>Log Message</button>
          <button @click=${this.logWarning}>Log Warning</button>
          <button @click=${this.logError}>Log Error</button>
        </div>
        <slot></slot>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-logger': DemoLogger;
  }
}
