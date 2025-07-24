@customElement('class-render-ternary')
class ClassRenderTernary extends LitElement {
  render() {
    const ternVar = condition ? html`
      <div part="true"></div>
    ` : html`
      <div part="false"></div>
    `;
    return Math.random() > 0.5 ? html`
      <!-- default slot (if you're lucky) -->
      <slot></slot>
      ${ternVar}
    ` : html`
      <!-- named slot (if you're less lucky) -->
      <slot name="slot"></slot>
      ${ternVar}
    `;
  }
}
