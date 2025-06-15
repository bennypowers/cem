import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element clipboard-copy
 * @summary Copies the given text to the clipboard.
 */
@customElement('clipboard-copy')
export class ClipboardCopy extends LitElement {
  @property({ type: String }) text = '';

  static styles = css`
    button { background: #2196f3; color: white; border: none; border-radius: 4px; padding: 0.4em 1em; }
  `;

  async copy() {
    await navigator.clipboard.writeText(this.text);
    this.dispatchEvent(new CustomEvent('copied', { detail: this.text }));
  }

  render() {
    return html`
      <span>${this.text}</span>
      <button @click=${this.copy}>Copy</button>
    `;
  }
}