import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element emoji-picker
 * @summary Basic emoji picker example.
 */
@customElement('emoji-picker')
export class EmojiPicker extends LitElement {
  @property({ type: Array }) emojis = ['😀','😂','😍','🤔','😎','😭'];
  @property({ type: String }) selected = '';

  static styles = css`
    button { font-size: 2rem; background: none; border: none; cursor: pointer; }
    button.selected { outline: 2px solid #2196f3; }
  `;

  private select(e: Event) {
    const emoji = (e.target as HTMLButtonElement).textContent || '';
    this.selected = emoji;
    this.dispatchEvent(new CustomEvent('emoji-selected', { detail: emoji }));
  }

  render() {
    return html`
      ${this.emojis.map(e => html`
        <button class=${this.selected === e ? 'selected' : ''} @click=${this.select}>${e}</button>
      `)}
    `;
  }
}