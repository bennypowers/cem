// Connection status toast - specifically for WebSocket connection states

export class CEMConnectionStatus extends HTMLElement {
  static template;

  static {
    this.template = document.createElement('template');
    this.template.innerHTML = `
      <style>
        :host {
          position: fixed;
          bottom: 20px;
          right: 20px;
          padding: 8px 12px;
          border-radius: 6px;
          font-family: system-ui, -apple-system, sans-serif;
          font-size: 12px;
          font-weight: 500;
          z-index: 999999;
          display: flex;
          align-items: center;
          gap: 8px;
          box-shadow: 0 4px 6px rgba(0,0,0,0.3);
          transition: opacity 0.3s ease;
        }

        :host([state="connected"]) {
          color: #10b981;
          background-color: #064e3b;
          border: 1px solid #10b981;
        }

        :host([state="reconnecting"]) {
          color: #fbbf24;
          background-color: #78350f;
          border: 1px solid #fbbf24;
        }

        :host([state="disconnected"]) {
          color: #ef4444;
          background-color: #7f1d1d;
          border: 1px solid #ef4444;
        }

        :host([state="connected"][faded]) {
          opacity: 0.3;
        }
      </style>
      <span id="icon"></span>
      <span id="message"></span>
    `;

    customElements.define('cem-connection-status', this);
  }

  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.appendChild(CEMConnectionStatus.template.content.cloneNode(true));
  }

  show(state, message, { fadeDelay = 0 } = {}) {
    const iconEl = this.shadowRoot.getElementById('icon');
    const messageEl = this.shadowRoot.getElementById('message');

    let icon;
    switch(state) {
      case 'connected':
        icon = '🟢';
        break;
      case 'reconnecting':
        icon = '🟡';
        break;
      case 'disconnected':
        icon = '🔴';
        break;
    }

    this.setAttribute('state', state);
    this.removeAttribute('faded');

    iconEl.textContent = icon;
    messageEl.textContent = message;

    if (fadeDelay > 0 && state === 'connected') {
      setTimeout(() => {
        this.setAttribute('faded', '');
      }, fadeDelay);
    }
  }

  hide() {
    this.style.opacity = '0';
    setTimeout(() => this.remove(), 300);
  }
}
