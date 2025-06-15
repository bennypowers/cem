import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * A button element with customizable label and click counter.
 */
@customElement('button-element')
export class ButtonElement extends LitElement {
  @property({ type: String }) label = 'Click me!';
  @property({ type: Number }) clicks = 0;

  static styles = css`:host { display: inline-block; }`;

  render() {
    return html`<button @click=${this.onClick}>${this.label} (${this.clicks})</button>`;
  }

  onClick() {
    this.clicks++;
    this.dispatchEvent(new CustomEvent('button-pressed', { detail: this.clicks }));
  }
}
