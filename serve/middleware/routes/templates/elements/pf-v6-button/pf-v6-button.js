import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 inspired button component
 *
 * @attr {string} variant - Button variant: primary, secondary, tertiary (default), danger, plain, link
 * @attr {string} size - Button size: sm, md (default), lg
 * @attr {boolean} block - Make button full width
 * @attr {boolean} disabled - Disable the button
 * @attr {string} href - URL for anchor links (automatically renders as <a> when present)
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
  static observedAttributes = ['disabled', 'aria-label', 'aria-expanded', 'aria-controls', 'aria-haspopup', 'type', 'href'];

  // Attach ElementInternals for cross-root ARIA references
  #internals = this.attachInternals();

  #button;
  #isAnchor = false;

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    this.toggleAttribute('disabled', !!value);
  }

  get variant() {
    return this.getAttribute('variant');
  }

  set variant(value) {
    if (value) {
      this.setAttribute('variant', value);
    } else {
      this.removeAttribute('variant');
    }
  }

  async afterTemplateLoaded() {
    // Check if we have an anchor or button (SSR renders the correct element)
    this.#isAnchor = this.hasAttribute('href');
    this.#button = this.#isAnchor
      ? this.shadowRoot.querySelector('a')
      : this.shadowRoot.querySelector('button');

    // Forward attributes to internal element
    this.#syncAttributes();

    // Delegate events
    this.#button.addEventListener('click', (e) => {
      if (this.disabled && !this.#isAnchor) {
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
    if (this.#isAnchor) {
      // Sync href for anchor elements
      if (this.hasAttribute('href')) {
        this.#button.setAttribute('href', this.getAttribute('href'));
      } else {
        this.#button.removeAttribute('href');
      }

      // Anchors can't be disabled in the same way, but we can prevent clicks
      if (this.hasAttribute('disabled')) {
        this.#button.setAttribute('aria-disabled', 'true');
        this.#button.style.pointerEvents = 'none';
      } else {
        this.#button.removeAttribute('aria-disabled');
        this.#button.style.pointerEvents = '';
      }
    } else {
      // Sync disabled for button elements
      if (this.hasAttribute('disabled')) {
        this.#button.disabled = true;
      } else {
        this.#button.disabled = false;
      }

      // Sync type attribute
      if (this.hasAttribute('type')) {
        this.#button.setAttribute('type', this.getAttribute('type'));
      }
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
  }

  static {
    customElements.define(this.is, this);
  }
}
