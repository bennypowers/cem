@customElement('class-render-slots')
class ClassRenderSlots extends LitElement {
  render() {
    return html`
      <slot></slot>
      <slot name="slot"></slot>
    `;
  }
}
