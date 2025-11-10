// Reusable status toast component

export class CEMStatusToast extends HTMLElement {
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
      </style>
      <span id="icon"></span>
      <span id="message"></span>
    `;

    customElements.define('cem-status-toast', this);
  }

  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.appendChild(CEMStatusToast.template.content.cloneNode(true));
  }

  show(state, message, { fadeDelay = 0 } = {}) {
    const iconEl = this.shadowRoot.getElementById('icon');
    const messageEl = this.shadowRoot.getElementById('message');

    let color, bgColor, icon;

    switch(state) {
      case 'success':
      case 'connected':
        color = '#10b981';
        bgColor = '#064e3b';
        icon = '🟢';
        break;
      case 'warning':
      case 'reconnecting':
        color = '#fbbf24';
        bgColor = '#78350f';
        icon = '🟡';
        break;
      case 'error':
      case 'disconnected':
        color = '#ef4444';
        bgColor = '#7f1d1d';
        icon = '🔴';
        break;
      case 'info':
        color = '#60a5fa';
        bgColor = '#1e3a8a';
        icon = 'ℹ️';
        break;
    }

    this.style.color = color;
    this.style.backgroundColor = bgColor;
    this.style.border = `1px solid ${color}`;
    this.style.opacity = '1';

    iconEl.textContent = icon;
    messageEl.textContent = message;

    if (fadeDelay > 0) {
      setTimeout(() => {
        this.style.opacity = '0.3';
      }, fadeDelay);
    }
  }

  hide() {
    this.style.opacity = '0';
    setTimeout(() => this.remove(), 300);
  }
}
