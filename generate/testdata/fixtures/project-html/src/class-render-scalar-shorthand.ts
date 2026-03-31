@customElement('class-render-scalar-shorthand')
class ClassRenderScalarShorthand extends LitElement {
  render() {
    return html`
      <!--
        slot: Sub navigation links
        part:
          summary: Nav container
          description: The scrollable link list container
      -->
      <slot name="nav" part="nav-container"></slot>
      <!--
        part: The overlay container
        slot:
          summary: Overlay
          description: Content shown in the overlay
      -->
      <slot name="overlay" part="overlay"></slot>
      <!-- slot: Info content
           part: Info container -->
      <slot name="info" part="info"></slot>
    `;
  }
}
