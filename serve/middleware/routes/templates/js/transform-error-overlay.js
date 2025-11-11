// Transform Error Overlay Component
// Displays TypeScript transform errors and other server-side compilation errors

export class CEMTransformErrorOverlay extends HTMLElement {
  static template;

  static {
    this.template = document.createElement('template');
    this.template.innerHTML = `
      <style>
        :host {
          display: none;
          position: fixed;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          z-index: 999999;
          background: rgba(0, 0, 0, 0.9);
          backdrop-filter: blur(4px);
          animation: fadeIn 0.2s ease-out;
        }

        :host([open]) {
          display: flex;
          align-items: center;
          justify-content: center;
          padding: 20px;
        }

        @keyframes fadeIn {
          from {
            opacity: 0;
          }
          to {
            opacity: 1;
          }
        }

        .overlay-content {
          background: #1e1e1e;
          color: #d4d4d4;
          border: 2px solid #ef4444;
          border-radius: 8px;
          max-width: 800px;
          width: 100%;
          max-height: 90vh;
          display: flex;
          flex-direction: column;
          box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
          font-family: 'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, 'Courier New', monospace;
        }

        .overlay-header {
          display: flex;
          align-items: center;
          justify-content: space-between;
          padding: 16px 20px;
          background: #ef4444;
          color: white;
          border-radius: 6px 6px 0 0;
        }

        .overlay-title {
          display: flex;
          align-items: center;
          gap: 12px;
          font-size: 16px;
          font-weight: 600;
          margin: 0;
        }

        .error-icon {
          font-size: 24px;
        }

        .close-button {
          background: rgba(255, 255, 255, 0.2);
          border: none;
          color: white;
          cursor: pointer;
          padding: 8px 12px;
          border-radius: 4px;
          font-size: 14px;
          font-weight: 500;
          transition: background 0.2s;
        }

        .close-button:hover {
          background: rgba(255, 255, 255, 0.3);
        }

        .overlay-body {
          padding: 20px;
          overflow-y: auto;
          flex: 1;
        }

        .error-file {
          background: #2d2d2d;
          padding: 8px 12px;
          border-radius: 4px;
          margin-bottom: 16px;
          font-size: 13px;
          color: #9cdcfe;
          border-left: 3px solid #ef4444;
        }

        .error-message {
          background: #252525;
          padding: 16px;
          border-radius: 4px;
          white-space: pre-wrap;
          word-break: break-word;
          font-size: 13px;
          line-height: 1.6;
          color: #d4d4d4;
          border: 1px solid #3a3a3a;
        }

        .overlay-footer {
          padding: 16px 20px;
          background: #2d2d2d;
          border-top: 1px solid #3a3a3a;
          border-radius: 0 0 6px 6px;
          font-size: 12px;
          color: #888;
        }
      </style>

      <div class="overlay-content">
        <div class="overlay-header">
          <h2 class="overlay-title">
            <span class="error-icon">⚠️</span>
            <span id="title"></span>
          </h2>
          <button class="close-button" id="close">Dismiss</button>
        </div>
        <div class="overlay-body">
          <div class="error-file" id="file"></div>
          <div class="error-message" id="message"></div>
        </div>
        <div class="overlay-footer">
          This error will automatically dismiss when the issue is fixed.
        </div>
      </div>
    `;

    customElements.define('cem-transform-error-overlay', this);
  }

  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.appendChild(CEMTransformErrorOverlay.template.content.cloneNode(true));

    this._titleEl = this.shadowRoot.getElementById('title');
    this._fileEl = this.shadowRoot.getElementById('file');
    this._messageEl = this.shadowRoot.getElementById('message');
    this._closeButton = this.shadowRoot.getElementById('close');

    this._closeButton.addEventListener('click', () => this.hide());

    // Close on Escape key
    this._handleKeydown = (e) => {
      if (e.key === 'Escape' && this.hasAttribute('open')) {
        this.hide();
      }
    };
    document.addEventListener('keydown', this._handleKeydown);
  }

  disconnectedCallback() {
    document.removeEventListener('keydown', this._handleKeydown);
  }

  show(title, message, file = '') {
    this._titleEl.textContent = title;
    this._messageEl.textContent = message;

    if (file) {
      this._fileEl.textContent = `File: ${file}`;
      this._fileEl.style.display = 'block';
    } else {
      this._fileEl.style.display = 'none';
    }

    this.setAttribute('open', '');
  }

  hide() {
    this.removeAttribute('open');
  }

  get open() {
    return this.hasAttribute('open');
  }
}
