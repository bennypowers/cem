@customElement('class-render-parts-docs')
class ClassRenderPartsDocs extends LitElement {
  render() {
    return html`
      <!--
          summary: part summary
          description: |
            part description

            multiline **markdown**
      -->
      <div part="part"></slot>
      <!-- summary: summary part summary -->
      <div part="summary"></slot>
      <!-- deprecated: true -->
      <div part="deprecated"></div>
      <!-- deprecated: deprecation reason -->
      <div part="reason"></slot>
    `;
  }
}
