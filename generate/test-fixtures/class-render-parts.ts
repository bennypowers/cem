@customElement('class-render-parts')
class ClassRenderParts extends LitElement {
  render() {
    return html`
      <div part="part"></div>
      <slot name="slot" part="slot slot-part"></slot>
    `;
  }
}
