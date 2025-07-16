@customElement('lit-template')
export class LitTemplate extends LitElement {
  render() {
    return html`
      <!-- summary: root level part -->
      <span part="root">
        <!-- summary: anonymous slot inside root part -->
        <slot></slot>
      </span>
      <button @click="${this.#onClick}">
        <!-- summary: inside of element with interpolated attribute -->
        <span part="inside">
          <!-- summary: named slot inside of part -->
          <slot name="inside-part"></slot>
        </span>
      </button>
      <!-- summary: after element with interpolated attribute -->
      <span part="after"></span>
    `
  }
}
