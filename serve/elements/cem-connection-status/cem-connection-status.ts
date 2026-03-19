import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-connection-status.css' with { type: 'css' };

type ConnectionState = 'connected' | 'reconnecting' | 'disconnected';

const ICONS: Record<ConnectionState, string> = {
  connected: '\u{1F7E2}',
  reconnecting: '\u{1F7E1}',
  disconnected: '\u{1F534}',
};

/**
 * Connection status toast for WebSocket connection states.
 *
 * @customElement cem-connection-status
 */
@customElement('cem-connection-status')
export class CemConnectionStatus extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor state: ConnectionState | undefined;

  @property({ type: Boolean, reflect: true })
  accessor faded = false;

  @property()
  accessor message = '';

  render() {
    return html`
      <span id="icon">${this.state ? ICONS[this.state] : ''}</span>
      <span id="message">${this.message}</span>
    `;
  }

  show(state: ConnectionState, message: string, { fadeDelay = 0 } = {}): void {
    this.state = state;
    this.faded = false;
    this.message = message;

    if (fadeDelay > 0 && state === 'connected') {
      setTimeout(() => {
        this.faded = true;
      }, fadeDelay);
    }
  }

  hide(): void {
    this.style.opacity = '0';
    setTimeout(() => this.remove(), 300);
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-connection-status': CemConnectionStatus;
  }
}
