/**
 * A card component
 * @summary A content card
 * @tag example-card
 * @slot - Card content
 * @slot header - Card header content
 */
export class ExampleCard extends HTMLElement {
  /** Card heading text */
  heading = '';
}

customElements.define('example-card', ExampleCard);
