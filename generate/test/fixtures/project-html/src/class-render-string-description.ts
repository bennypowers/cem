@customElement('class-render-string-description')
class ClassRenderStringDescription extends LitElement {
  render() {
    return html`
      <!-- anonymous -->
      <slot></slot>
      <!-- name="slot" -->
      <slot name="slot"></slot>
      <!-- name="part" -->
      <a part="part"></a>
    `;
  }
}
