@customElement('class-render-escaped-backticks')
class ClassRenderEscapedBackticks extends LitElement {
  render() {
    return html`
      <!-- \`<slot>\` -->
      <slot></slot>
      <!-- summary: "\`<slot>\`" -->
      <slot name="slot"></slot>
      <!-- \`<p>\` -->
      <p part="part"></p>
      <!-- summary: "\`<p>\`" -->
      <p part="summary"></p>
    `;
  }
}
