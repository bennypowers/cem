// Reusable error dialog component

export class CEMErrorDialog extends HTMLElement {
  static template;

  static {
    this.template = document.createElement('template');
    this.template.innerHTML = `
      <style>
        dialog {
          border: none;
          border-radius: 8px;
          padding: 0;
          max-width: 500px;
          width: 90%;
          box-shadow: 0 20px 25px -5px rgba(0,0,0,0.5);
        }

        dialog::backdrop {
          background: rgba(0, 0, 0, 0.5);
        }

        .dialog-content {
          background: #1e293b;
          color: #e2e8f0;
          padding: 24px;
          font-family: system-ui, -apple-system, sans-serif;
        }

        h2 {
          margin: 0 0 16px 0;
          color: #60a5fa;
          font-size: 20px;
          font-weight: 600;
        }

        ::slotted(p) {
          margin: 0 0 16px 0;
          line-height: 1.6;
          color: #cbd5e1;
        }

        .buttons {
          display: flex;
          gap: 12px;
          justify-content: flex-end;
        }

        ::slotted(button) {
          padding: 8px 16px;
          border-radius: 4px;
          cursor: pointer;
          font-size: 14px;
          font-weight: 500;
          border: none;
        }

        button {
          padding: 8px 16px;
          border-radius: 4px;
          cursor: pointer;
          font-size: 14px;
          font-weight: 500;
          border: none;
        }

        button:hover {
          opacity: 0.9;
        }
      </style>

      <dialog>
        <div class="dialog-content">
          <h2 id="title"></h2>
          <slot></slot>
          <div class="buttons">
            <slot name="buttons"></slot>
          </div>
        </div>
      </dialog>
    `;

    customElements.define('cem-error-dialog', this);
  }

  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.appendChild(CEMErrorDialog.template.content.cloneNode(true));

    this._dialog = this.shadowRoot.querySelector('dialog');

    // Handle ESC key
    this._dialog.addEventListener('close', () => {
      this.dispatchEvent(new CustomEvent('close'));
    });
  }

  get title() {
    return this.shadowRoot.getElementById('title').textContent;
  }

  set title(value) {
    this.shadowRoot.getElementById('title').textContent = value;
  }

  show() {
    if (!this._dialog.open) {
      this._dialog.showModal();
    }
  }

  hide() {
    if (this._dialog.open) {
      this._dialog.close();
    }
  }

  get open() {
    return this._dialog.open;
  }
}
