import { CemElement } from '/__cem/cem-element.js';

class Pfv6Switch extends CemElement {
  static observedAttributes = ['checked', 'disabled', 'aria-label', 'aria-labelledby'];

  #input;

  async afterTemplateLoaded() {
    this.#input = this.shadowRoot.getElementById('switch-input');
    if (!this.#input) return;

    // Sync initial state
    this.#syncAttributes();

    // Forward change events from internal checkbox and sync checked state
    this.#input.addEventListener('change', () => {
      // Sync the checked attribute to match internal input
      if (this.#input.checked) {
        this.setAttribute('checked', '');
      } else {
        this.removeAttribute('checked');
      }
      this.dispatchEvent(new Event('change', { bubbles: true }));
    });

    // Forward input events from internal checkbox
    this.#input.addEventListener('input', () => {
      this.dispatchEvent(new Event('input', { bubbles: true }));
    });
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#input) return;

    // Sync checked state
    this.#input.checked = this.hasAttribute('checked');

    // Sync disabled state
    this.#input.disabled = this.hasAttribute('disabled');

    // Sync aria attributes
    const ariaLabel = this.getAttribute('aria-label');
    if (ariaLabel) {
      this.#input.setAttribute('aria-label', ariaLabel);
    } else {
      this.#input.removeAttribute('aria-label');
    }

    const ariaLabelledby = this.getAttribute('aria-labelledby');
    if (ariaLabelledby) {
      this.#input.setAttribute('aria-labelledby', ariaLabelledby);
    } else {
      this.#input.removeAttribute('aria-labelledby');
    }
  }

  // Public API
  get checked() {
    return this.hasAttribute('checked');
  }

  set checked(value) {
    if (value) {
      this.setAttribute('checked', '');
    } else {
      this.removeAttribute('checked');
    }
  }

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    if (value) {
      this.setAttribute('disabled', '');
    } else {
      this.removeAttribute('disabled');
    }
  }

  get value() {
    return this.checked;
  }

  set value(val) {
    this.checked = val;
  }
}

customElements.define('pfv6-switch', Pfv6Switch);
