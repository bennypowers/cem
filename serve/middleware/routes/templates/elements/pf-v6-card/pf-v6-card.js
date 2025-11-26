import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Card Component
 *
 * A card is a square or rectangular container that can contain any kind of content.
 *
 * @attr {boolean} compact - Makes the card compact with reduced padding
 * @attr {string} variant - Card variant: "secondary" or "plain"
 * @attr {boolean} full-height - Makes the card fill its container height
 *
 * @slot - Default slot for body content
 * @slot title - Optional title content (h1-h6 headings are automatically styled)
 * @slot header - Optional header content (displayed after title)
 * @slot footer - Optional footer content
 *
 * @customElement pf-v6-card
 */
export class PfV6Card extends CemElement {
  static is = 'pf-v6-card';
  static observedAttributes = ['compact', 'variant', 'full-height'];

  static {
    customElements.define(this.is, this);
  }
}
