@customElement('class-render-parts')
class ClassMethods extends LitElement {
  render() {
    return html`
      <div part="part"></slot>
      <slot name="slot" part="slot slot-part"></slot>
    `;
  }
}
