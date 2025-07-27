@customElement('lit-template')
export class LitTemplate extends LitElement {
  render() {
    return html`
      <!-- summary: part:root -->
      <span part="root">
        <!-- summary: anonymous slot child of part:root -->
        <slot></slot>
      </span>
      <button @click="${this.#onClick}">
        <!-- summary: part child of element with @click -->
        <span part="inside">
          <!-- summary: named slot child of part:inside -->
          <slot name="inside-part"></slot>
        </span>
      </button>
      <!-- summary: after element with @click -->
      <span part="after"></span>
    `
  }
}
