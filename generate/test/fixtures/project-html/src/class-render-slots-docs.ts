@customElement('class-render-slots-docs')
class ClassRenderSlotsDocs extends LitElement {
  render() {
    return html`
      <!--
          summary: default slot summary
          description: |
            default slot description

            multiline **markdown**
      -->
      <slot></slot>
      <!-- summary: named slot summary -->
      <slot name="named"></slot>
      <!-- summary: named slot summary
           description: |
             named slot multiline description

             **with markdown** -->
      <slot name="named-described"></slot>
      <!-- deprecated: true -->
      <slot name="deprecated"></slot>
      <!-- deprecated: deprecation reason -->
      <slot name="reason"></slot>
    `;
  }
}
