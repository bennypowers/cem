import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 inspired button component
 *
 * @attr {string} variant - Button variant: primary, secondary, tertiary (default), danger, plain, link
 * @attr {string} size - Button size: sm, md (default), lg
 * @attr {boolean} block - Make button full width
 * @attr {boolean} disabled - Disable the button
 * @attr {string} aria-label - Accessible label for the button
 *
 * @slot - Default slot for button text/content
 * @slot icon-start - Slot for icon before text
 * @slot icon-end - Slot for icon after text
 *
 * @fires click - Bubbles click events from internal button
 * @customElement pf-v6-button
 */
class PfV6Button extends CemElement {
  static shadowRootOptions = { mode: 'open', delegatesFocus: true };
  static is = 'pf-v6-button';
  static observedAttributes = ['disabled', 'aria-label', 'aria-expanded', 'aria-controls', 'aria-haspopup', 'type'];

  // Attach ElementInternals for cross-root ARIA references
  #internals = this.attachInternals();

  #button;

  async afterTemplateLoaded() {
    this.#button = this.shadowRoot.querySelector('button');
    if (!this.#button) return;

    // Forward attributes to internal button
    this.#syncAttributes();

    // Delegate button events
    this.#button.addEventListener('click', (e) => {
      if (this.disabled) {
        e.preventDefault();
        e.stopPropagation();
        return;
      }
      // Event bubbles naturally through shadow DOM
    });
  }

  attributeChangedCallback(name, oldValue, newValue) {
    // Only sync if button is ready
    if (this.#button) {
      this.#syncAttributes();
    }
  }

  #syncAttributes() {
    // Sync disabled
    if (this.hasAttribute('disabled')) {
      this.#button.disabled = true;
    } else {
      this.#button.disabled = false;
    }

    // Cross-root ARIA references go on ElementInternals (host)
    // This allows aria-controls to reference elements outside shadow root
    if (this.hasAttribute('aria-controls')) {
      this.#internals.ariaControls = this.getAttribute('aria-controls');
    } else {
      this.#internals.ariaControls = null;
    }

    // Non-reference ARIA attributes go on internal button
    // These don't need to cross shadow boundaries
    const buttonAriaAttrs = ['aria-label', 'aria-expanded', 'aria-haspopup'];
    buttonAriaAttrs.forEach(attr => {
      if (this.hasAttribute(attr)) {
        this.#button.setAttribute(attr, this.getAttribute(attr));
      } else {
        this.#button.removeAttribute(attr);
      }
    });

    // Sync type attribute
    if (this.hasAttribute('type')) {
      this.#button.setAttribute('type', this.getAttribute('type'));
    }
  }

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    this.toggleAttribute('disabled', !!value);
  }

  static {
    customElements.define(this.is, this);
  }
}

