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
 * @slot header - Optional header content
 * @slot title - Optional title content
 * @slot footer - Optional footer content
 *
 * @customElement pf-v6-card
 */
export class PfV6Card extends CemElement {
  static is = 'pf-v6-card';
  static observedAttributes = ['compact', 'variant', 'full-height'];

  #header;
  #title;
  #footer;

  async afterTemplateLoaded() {
    this.#header = this.shadowRoot.querySelector('#header');
    this.#title = this.shadowRoot.querySelector('#title');
    this.#footer = this.shadowRoot.querySelector('#footer');

    this.#updateSlotVisibility();

    // Observe slot changes to update visibility
    const slots = this.shadowRoot.querySelectorAll('slot');
    slots.forEach(slot => {
      slot.addEventListener('slotchange', () => this.#updateSlotVisibility());
    });
  }

  #updateSlotVisibility() {
    if (!this.#header || !this.#title || !this.#footer) return;

    // Show/hide header based on slotted content
    const headerSlot = this.shadowRoot.querySelector('slot[name="header"]');
    const hasHeaderContent = headerSlot && headerSlot.assignedNodes().length > 0;
    this.#header.hidden = !hasHeaderContent;

    // Show/hide title based on slotted content
    const titleSlot = this.shadowRoot.querySelector('slot[name="title"]');
    const hasTitleContent = titleSlot && titleSlot.assignedNodes().length > 0;
    this.#title.hidden = !hasTitleContent;

    // Show/hide footer based on slotted content
    const footerSlot = this.shadowRoot.querySelector('slot[name="footer"]');
    const hasFooterContent = footerSlot && footerSlot.assignedNodes().length > 0;
    this.#footer.hidden = !hasFooterContent;
  }

  static {
    customElements.define(this.is, this);
  }
}
