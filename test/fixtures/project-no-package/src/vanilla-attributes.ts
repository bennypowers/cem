/** @customElement vanilla-attributes */
class VanillaAttributes extends HTMLElement {
  static observedAttributes = ['a', 'b'];
  static get observedAttributes() {
    return ['c'];
  }
}

customElements.define('vanilla-attributes', VanillaAttributes);
