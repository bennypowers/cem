import { LitElement, html, property, customElement } from 'lit-element';

@customElement('tab-navigation')
export class TabNavigation extends LitElement {
  @property({ type: Array }) tabs = ['Tab 1', 'Tab 2'];
  @property({ type: Number }) selected = 0;

  render() {
    return html`
      <nav>
        ${this.tabs.map(
          (t, i) => html`
            <button
              ?disabled=${i === this.selected}
              @click=${() => this.#select(i)}
            >
              ${t}
            </button>
          `
        )}
      </nav>
      <section>
        <slot name="panel-${this.selected}"></slot>
      </section>
    `;
  }

  #select(i: number) {
    this.selected = i;
    this.dispatchEvent(new CustomEvent('tab-selected', { detail: i }));
  }
}
