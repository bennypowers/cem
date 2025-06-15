@customElement('class-render-slots')
class ClassMethods extends LitElement {
  render() {
    return html`
      <slot></slot>
      <slot name="slot"></slot>
    `;
  }
}
