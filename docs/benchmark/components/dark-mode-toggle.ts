import { LitElement, html, property, customElement } from 'lit-element';

@customElement('dark-mode-toggle')
export class DarkModeToggle extends LitElement {
  @property({ type: Boolean }) dark = false;

  render() {
    return html`
      <button @click=${this.#toggle}>
        ${this.dark ? '🌙' : '☀️'} Toggle Dark Mode
      </button>
    `;
  }

  #toggle() {
    this.dark = !this.dark;
    this.dispatchEvent(new CustomEvent('dark-mode', { detail: this.dark }));
  }
} 
