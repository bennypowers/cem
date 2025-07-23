@customElement('class-render-slot-anon-attributes')
class ClassRenderSlotAnonAttributes extends LitElement {
  render() {
    return html`
      <!-- summary: summary
           description: |
             \`<slot>\` -->
      <slot class="foo"></slot>
    `;
  }
}
