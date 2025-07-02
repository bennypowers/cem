import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element keyboard-shortcut
 * @description Registers a keyboard shortcut and emits an event when pressed.
 */
@customElement('keyboard-shortcut')
export class KeyboardShortcut extends LitElement {
  @property({ type: String }) key = 'k';
  @property({ type: Boolean }) ctrl = false;

  connectedCallback() {
    super.connectedCallback();
    window.addEventListener('keydown', this._handle);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener('keydown', this._handle);
  }

  private _handle = (e: KeyboardEvent) => {
    if ((this.ctrl ? e.ctrlKey : true) && e.key === this.key) {
      this.dispatchEvent(new CustomEvent('shortcut'));
    }
  };

  render() {
    return html`<slot></slot>`;
  }
}