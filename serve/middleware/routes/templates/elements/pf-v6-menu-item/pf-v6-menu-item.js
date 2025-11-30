import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Menu Item component
 *
 * @attr {string} value - Item value
 * @attr {boolean} disabled - Disabled state
 * @attr {boolean} checked - Checked state (for checkbox variant)
 * @attr {string} variant - Item variant: "default" | "checkbox"
 * @attr {string} description - Optional description text
 *
 * @slot - Default slot for item text
 *
 * @fires {PfMenuItemSelectEvent} select - Fires when item is selected/checked
 * @customElement pf-v6-menu-item
 */

export class PfMenuItemSelectEvent extends Event {
  constructor(value, checked) {
    super('select', { bubbles: true });
    this.value = value;
    this.checked = checked;
  }
}

class PfV6MenuItem extends CemElement {
  static is = 'pf-v6-menu-item';

  static shadowRootOptions = { mode: 'open', delegatesFocus: true };

  static observedAttributes = [
    'disabled',
    'checked',
    'variant',
    'value',
  ];

  #internals = this.attachInternals();
  #input;

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    this.toggleAttribute('disabled', !!value);
  }

  get checked() {
    return this.hasAttribute('checked');
  }

  set checked(value) {
    this.toggleAttribute('checked', !!value);
  }

  get variant() {
    return this.getAttribute('variant') || 'default';
  }

  set variant(value) {
    if (value) {
      this.setAttribute('variant', value);
    } else {
      this.removeAttribute('variant');
    }
  }

  get value() {
    return this.getAttribute('value') || '';
  }

  set value(v) {
    if (v) {
      this.setAttribute('value', v);
    } else {
      this.removeAttribute('value');
    }
  }

  async afterTemplateLoaded() {
    this.#input = this.shadowRoot.getElementById('input');

    // Set up ARIA role based on variant
    this.#updateRole();

    // Set up event listeners on host
    this.addEventListener('click', this.#handleClick);
    this.addEventListener('keydown', this.#handleKeydown);

    // Initialize ARIA state
    this.#updateAriaChecked();
    this.#updateDisabled();

    // Sync checkbox with initial checked state
    this.#syncCheckbox();
  }

  #updateRole() {
    const isCheckbox = this.variant === 'checkbox';
    this.#internals.role = isCheckbox ? 'menuitemcheckbox' : 'menuitem';
  }

  #handleClick = (event) => {
    if (this.disabled) {
      event.preventDefault();
      event.stopImmediatePropagation();
      return;
    }

    if (this.variant === 'checkbox') {
      // Toggle checked state for checkbox variant
      this.checked = !this.checked;
    }

    // Dispatch select event
    this.dispatchEvent(new PfMenuItemSelectEvent(this.value, this.checked));
  };

  #handleKeydown = (event) => {
    if (this.disabled) {
      return;
    }

    // Handle Space and Enter for item activation
    if (event.key === ' ' || event.key === 'Enter') {
      event.preventDefault();
      this.#handleClick(event);
    }
  };

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.shadowRoot?.firstChild) {
      return;
    }

    switch (name) {
      case 'disabled':
        this.#updateDisabled();
        break;
      case 'checked':
        this.#updateAriaChecked();
        this.#syncCheckbox();
        break;
      case 'variant':
        this.#updateRole();
        this.#updateAriaChecked();
        break;
    }
  }

  #updateDisabled() {
    if (this.disabled) {
      this.#internals.ariaDisabled = 'true';
      this.setAttribute('tabindex', '-1');
    } else {
      this.#internals.ariaDisabled = null;
      this.setAttribute('tabindex', '-1'); // Will be managed by menu RTI
    }

    // Update native checkbox if present
    if (this.#input) {
      this.#input.disabled = this.disabled;
    }
  }

  #updateAriaChecked() {
    if (this.variant === 'checkbox') {
      this.#internals.ariaChecked = this.checked ? 'true' : 'false';
    } else {
      this.#internals.ariaChecked = null;
    }
  }

  #syncCheckbox() {
    if (this.#input) {
      this.#input.checked = this.checked;
    }
  }

  disconnectedCallback() {
    this.removeEventListener('click', this.#handleClick);
    this.removeEventListener('keydown', this.#handleKeydown);
  }

  static {
    customElements.define(this.is, this);
  }
}
