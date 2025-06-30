const PALETTE_RE = /(er|est)+/g;

@customElement('has-will-update')
export class HasWillUpdate extends LitElement {
  static styles = [styles];

  @property({ reflect: true, attribute: 'color-palette' }) colorPalette?: ColorPalette;

  /**
   * Change the style of the card to be a "Promo"
   */
  @property({ reflect: true }) variant?: 'promo';

  /**
   * Change a promo with an image + body + footer to use the `full-width` style
   */
  @property({ reflect: true, attribute: 'full-width', type: Boolean }) fullWidth = false;

  willUpdate() {
    this.#isPromo = this.variant === 'promo';
    this.#isStandardPromo =
         this.#isPromo
      && this.#slots.hasSlotted(null)
      && this.#slots.isEmpty('image')
      && this.#slots.isEmpty('header');
  }
}

