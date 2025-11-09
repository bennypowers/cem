// Reconnection dialog content component

export class CEMReconnectionContent extends HTMLElement {
  static template;

  static {
    this.template = document.createElement('template');
    this.template.innerHTML = `
      <style>
        :host {
          display: block;
        }

        p {
          margin: 0 0 16px 0;
          line-height: 1.6;
          color: #cbd5e1;
        }

        #retry-info {
          background: #0f172a;
          padding: 12px;
          border-radius: 4px;
          margin-bottom: 16px;
          font-size: 14px;
          color: #94a3b8;
        }

        .buttons {
          display: flex;
          gap: 12px;
          justify-content: flex-end;
        }

        button {
          padding: 8px 16px;
          border-radius: 4px;
          cursor: pointer;
          font-size: 14px;
          font-weight: 500;
          border: none;
        }

        #reload-btn {
          background: #334155;
          color: #e2e8f0;
          border: 1px solid #475569;
        }

        #retry-btn {
          background: #60a5fa;
          color: #0f172a;
          font-weight: 600;
        }

        button:hover {
          opacity: 0.9;
        }
      </style>

      <p>
        The connection to the development server was lost.
        Automatically retrying connection...
      </p>

      <div id="retry-info"></div>

      <div class="buttons">
        <button id="reload-btn">Reload Page</button>
        <button id="retry-btn">Retry Now</button>
      </div>
    `;

    customElements.define('cem-reconnection-content', this);
  }

  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.appendChild(CEMReconnectionContent.template.content.cloneNode(true));

    this.countdownInterval = null;
    this.remainingMs = 0;

    // Set up event listeners
    this.shadowRoot.getElementById('reload-btn').addEventListener('click', () => {
      this.dispatchEvent(new CustomEvent('reload'));
    });

    this.shadowRoot.getElementById('retry-btn').addEventListener('click', () => {
      this.clearCountdown();
      this.dispatchEvent(new CustomEvent('retry'));
    });
  }

  clearCountdown() {
    if (this.countdownInterval) {
      clearInterval(this.countdownInterval);
      this.countdownInterval = null;
    }
  }

  updateRetryInfo(attempt, delay) {
    this.clearCountdown();

    const retryInfo = this.shadowRoot.getElementById('retry-info');
    this.remainingMs = delay;

    // Update immediately
    const seconds = Math.ceil(this.remainingMs / 1000);
    retryInfo.textContent = `Attempt #${attempt}. Retrying in ${seconds}s...`;

    // Update countdown every 100ms for smooth display
    this.countdownInterval = setInterval(() => {
      this.remainingMs -= 100;

      if (this.remainingMs <= 0) {
        this.clearCountdown();
        retryInfo.textContent = `Attempt #${attempt}. Connecting...`;
        return;
      }

      const seconds = Math.ceil(this.remainingMs / 1000);
      retryInfo.textContent = `Attempt #${attempt}. Retrying in ${seconds}s...`;
    }, 100);
  }

  disconnectedCallback() {
    this.clearCountdown();
  }
}
