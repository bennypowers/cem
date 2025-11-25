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
  static is = 'pf-v6-button';

  static shadowRootOptions = { mode: 'open', delegatesFocus: true };

  static observedAttributes = [
    'disabled',
    'variant',
    'type',
    'href',
  ];

  // Attach ElementInternals for role and ARIA management
  #internals = this.attachInternals();

  #element; // Will be <a> or null (for host-based button)
  #isLink = false;

  get disabled() { return this.hasAttribute('disabled'); }
  set disabled(value) { this.toggleAttribute('disabled', !!value); }

  get variant() { return this.getAttribute('variant'); }
  set variant(value) {
    if (value) {
      this.setAttribute('variant', value);
    } else {
      this.removeAttribute('variant');
    }
  }

  get href() { return this.getAttribute('href') || ''; }
  set href(v) {
    if (v) {
      this.setAttribute('href', v);
    } else {
      this.removeAttribute('href');
    }
  }

  get type() { return this.getAttribute('type') || 'button'; }
  set type(v) {
    if (v) {
      this.setAttribute('type', v);
    } else {
      this.removeAttribute('type');
    }
  }

  async afterTemplateLoaded() {
    // Check if we have a link
    this.#isLink = this.hasAttribute('href');
    this.#element = this.#isLink
      ? this.shadowRoot.querySelector('a')
      : null; // No shadow element for button variant

    // Set up button or link behavior
    if (!this.#isLink) {
      this.#setupHostButton();
    } else {
      this.#setupShadowLink();
    }

    // Set up slot visibility management
    this.shadowRoot.querySelectorAll('slot').forEach(slot => {
      slot.addEventListener('slotchange', () => this.#updateSlotVisibility());
    });
    this.#updateSlotVisibility();

    // Forward attributes
    this.#syncAttributes();
  }

  #updateSlotVisibility() {
    const iconStartSlot = this.shadowRoot.querySelector('slot[name="icon-start"]');
    const iconEndSlot = this.shadowRoot.querySelector('slot[name="icon-end"]');

    // Hide slots that have no content
    if (iconStartSlot) {
      iconStartSlot.hidden = iconStartSlot.assignedNodes().length === 0;
    }
    if (iconEndSlot) {
      iconEndSlot.hidden = iconEndSlot.assignedNodes().length === 0;
    }
  }

  #setupHostButton() {
    // Set role on host for button variant
    this.#internals.role = 'button';

    // Make host focusable
    if (!this.hasAttribute('tabindex')) {
      this.setAttribute('tabindex', '0');
    }

    // Add event listeners on HOST
    this.addEventListener('click', this.#handleClick);
    this.addEventListener('keydown', this.#handleKeydown);
  }

  #setupShadowLink() {
    // For links, rely on delegatesFocus and shadow <a>
    // Add click handler to enforce disabled state
    if (this.#element) {
      this.#element.addEventListener('click', this.#handleLinkClick);
    }
  }

  #handleClick = (event) => {
    // For host-based buttons only
    if (this.disabled) {
      event.preventDefault();
      event.stopImmediatePropagation();
      return;
    }
    // Event bubbles naturally
  };

  #handleKeydown = (event) => {
    // For host-based buttons only
    if (this.disabled) {
      event.preventDefault();
      return;
    }

    // Handle Space and Enter for button activation
    if (event.key === ' ' || event.key === 'Enter') {
      event.preventDefault();
      this.click();
    }
  };

  #handleLinkClick = (event) => {
    // For shadow links only - enforce disabled
    if (this.disabled) {
      event.preventDefault();
      event.stopPropagation();
      return;
    }
  };

  attributeChangedCallback(name, oldValue, newValue) {
    switch (name) {
      case 'disabled': this.#updateDisabled(); break;
      case 'href':
      case 'variant': this.#updateRole(); break;
    }

    // Sync other attributes
    if (this.#element || !this.#isLink) {
      this.#syncAttributes();
    }
  }

  #updateRole() {
    const hasHref = this.hasAttribute('href');

    if (hasHref) {
      // Switching to link: remove button role
      this.#internals.role = null;
    } else {
      // Switching to button: set button role
      this.#internals.role = 'button';
    }
  }

  #updateDisabled() {
    if (this.#isLink) {
      // For links, set aria-disabled and prevent interaction
      if (this.#element) {
        if (this.disabled) {
          this.#element.setAttribute('aria-disabled', 'true');
          this.#element.style.pointerEvents = 'none';
        } else {
          this.#element.removeAttribute('aria-disabled');
          this.#element.style.pointerEvents = '';
        }
      }
    } else {
      // For host-based buttons, use ElementInternals
      if (this.disabled) {
        this.#internals.ariaDisabled = 'true';
        this.setAttribute('tabindex', '-1');
      } else {
        this.#internals.ariaDisabled = null;
        this.setAttribute('tabindex', '0');
      }
    }
  }

  #syncAttributes() {
    if (this.#isLink && this.#element) {
      // Sync to shadow <a>
      if (this.hasAttribute('href')) {
        this.#element.setAttribute('href', this.getAttribute('href'));
      }

      // ARIA attributes on shadow link
      const linkAriaAttrs = ['aria-label', 'aria-expanded', 'aria-haspopup'];
      linkAriaAttrs.forEach(attr => {
        if (this.hasAttribute(attr)) {
          this.#element.setAttribute(attr, this.getAttribute(attr));
        } else {
          this.#element.removeAttribute(attr);
        }
      });
    }
    // For button variant: ARIA attrs stay on host, no syncing needed
    // since we set role='button' on host via ElementInternals

    // aria-controls always goes on host (cross-root reference)
    if (this.hasAttribute('aria-controls')) {
      this.#internals.ariaControls = this.getAttribute('aria-controls');
    } else {
      this.#internals.ariaControls = null;
    }
  }

  disconnectedCallback() {
    if (!this.#isLink) {
      this.removeEventListener('click', this.#handleClick);
      this.removeEventListener('keydown', this.#handleKeydown);
    }
    if (this.#element) {
      this.#element.removeEventListener('click', this.#handleLinkClick);
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
