import { LitElement, html, property, customElement } from 'lit-element';

@customElement('toggle-switch')
export class ToggleSwitch extends LitElement {
  @property({ type: Boolean, reflect: true }) checked = false;

  render() {
    return html`
      <label>
        <input type="checkbox" .checked=${this.checked} @change=${this.#onToggle} />
        <span><slot></slot></span>
      </label>
    `;
  }

  #onToggle(e: Event) {
    this.checked = (e.target as HTMLInputElement).checked;
    this.dispatchEvent(new CustomEvent('toggle-changed', { detail: this.checked }));
  }
}
