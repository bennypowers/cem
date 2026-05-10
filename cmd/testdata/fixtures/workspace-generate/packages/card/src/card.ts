/**
 * A card
 * @summary Test card
 * @customElement test-card
 * @slot - Card content
 */
export class TestCard extends HTMLElement {
  /** Heading */
  heading = '';
}

customElements.define('test-card', TestCard);
