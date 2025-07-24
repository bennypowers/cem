@customElement('class-render-return-ternary')
class ClassRenderReturnTernary extends LitElement {
  render() {
    return Math.random() > 0.5 ? html`
      <!-- default slot (if you're lucky) -->
      <slot></slot>
    ` : html`
      <!-- named slot (if you're less lucky) -->
      <slot name="slot"></slot>
    `;
  }
}
