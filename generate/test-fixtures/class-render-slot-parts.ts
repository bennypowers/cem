@customElement('class-render-slot-parts')
class ClassRenderSlotParts extends LitElement {
  render() {
    return html`
      <!-- slot:
             summary: slot summary
           part:
             summary: part summary
      -->
      <slot name="slot-with-part" part="part-with-slot"></slot>
    `;
  }
}
