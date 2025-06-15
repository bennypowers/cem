import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element theme-toggle
 * @summary Toggle between light and dark themes.
 */
@customElement('theme-toggle')
export class ThemeToggle extends LitElement {
  @property({ type: Boolean, reflect: true }) dark = false;

  static styles = css`
    button { padding: 0.5em 1em; }
  `;

  get icon() {
    return this.dark ? '🌙' : '☀️';
  }

  toggleTheme() {
    this.dark = !this.dark;
    this.dispatchEvent(new CustomEvent('theme-changed', { detail: this.dark }));
  }

  render() {
    return html`
      <button @click=${this.toggleTheme}>${this.icon} Toggle Theme</button>
    `;
  }
}