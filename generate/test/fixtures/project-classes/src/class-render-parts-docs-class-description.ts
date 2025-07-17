/** class description */
@customElement('class-render-parts-docs-class-description')
class ClassRenderPartsDocsClassDescription extends LitElement {
  render() {
    return html`
      <!--
          summary: part summary
          description: |
            part description

            multiline **markdown**
      -->
      <div part="part">${childPart}</div>
      <!-- summary: summary part summary -->
      <div part="summary"></div>
      <!-- deprecated: true -->
      <div part="deprecated"></div>
      <!-- deprecated: deprecation reason -->
      <div part="reason"></div>
    `;
  }
}
