import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Text Input Group
 *
 * A text input with optional icons and utility buttons.
 *
 * @slot icon - Leading icon
 * @slot status-icon - Trailing status icon (success, warning, error)
 * @slot utilities - Utility buttons/actions (e.g., eyedropper, clear)
 *
 * @attr {string} value - Input value
 * @attr {string} type - Input type (default: "text")
 * @attr {string} placeholder - Placeholder text
 * @attr {string} aria-label - Accessible label
 * @attr {boolean} disabled - Disabled state
 * @attr {boolean} readonly - Readonly state
 * @attr {boolean} required - Required state
 * @attr {boolean} plain - Plain variant (no background/border)
 * @attr {boolean} icon - Has leading icon
 * @attr {string} status - Validation status: success|warning|error
 *
 * @fires input - Native input event
 * @fires change - Native change event
 *
 * @customElement pf-v6-text-input-group
 */
class PfV6TextInputGroup extends CemElement {
  static is = 'pf-v6-text-input-group';
  static formAssociated = true;

  static observedAttributes = [
    'value',
    'type',
    'placeholder',
    'disabled',
    'readonly',
    'required',
    'plain',
    'icon',
    'status'
  ];

  #input;
  #internals = this.attachInternals();

  async afterTemplateLoaded() {
    this.#input = this.shadowRoot.getElementById('input');

    // Connect ElementInternals labels to the internal input
    // This allows <label for="input-group-id"> to work properly
    if (this.#internals.labels) {
      const labelIds = Array.from(this.#internals.labels).map(label => label.id).filter(Boolean);
      if (labelIds.length > 0) {
        this.#input.setAttribute('aria-labelledby', labelIds.join(' '));
      }
    }

    // Forward input and change events
    this.#input?.addEventListener('input', (e) => {
      this.setAttribute('value', e.target.value);
      // Re-dispatch event for external listeners
      this.dispatchEvent(new InputEvent('input', {
        bubbles: true,
        composed: true,
        data: e.data,
        inputType: e.inputType
      }));
    });

    this.#input?.addEventListener('change', (e) => {
      this.dispatchEvent(new Event('change', {
        bubbles: true,
        composed: true
      }));
    });
  }

  // Getters and setters
  get value() {
    return this.getAttribute('value') || '';
  }

  set value(val) {
    if (val) {
      this.setAttribute('value', val);
    } else {
      this.removeAttribute('value');
    }
  }

  get type() {
    return this.getAttribute('type') || 'text';
  }

  set type(val) {
    this.setAttribute('type', val);
  }

  get placeholder() {
    return this.getAttribute('placeholder') || '';
  }

  set placeholder(val) {
    if (val) {
      this.setAttribute('placeholder', val);
    } else {
      this.removeAttribute('placeholder');
    }
  }

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(val) {
    this.toggleAttribute('disabled', !!val);
  }

  get readonly() {
    return this.hasAttribute('readonly');
  }

  set readonly(val) {
    this.toggleAttribute('readonly', !!val);
  }

  get required() {
    return this.hasAttribute('required');
  }

  set required(val) {
    this.toggleAttribute('required', !!val);
  }

  get status() {
    return this.getAttribute('status');
  }

  set status(val) {
    if (val && ['success', 'warning', 'error'].includes(val)) {
      this.setAttribute('status', val);
    } else {
      this.removeAttribute('status');
    }
  }

  // Public methods
  focus() {
    this.#input?.focus();
  }

  blur() {
    this.#input?.blur();
  }

  select() {
    this.#input?.select();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue || !this.#input) return;

    switch (name) {
      case 'value':
        if (this.#input.value !== newValue) {
          this.#input.value = newValue || '';
        }
        break;

      case 'type':
        this.#input.type = newValue || 'text';
        break;

      case 'placeholder':
        if (newValue) {
          this.#input.placeholder = newValue;
        } else {
          this.#input.removeAttribute('placeholder');
        }
        break;

      case 'disabled':
        this.#input.toggleAttribute('disabled', newValue !== null);
        break;

      case 'readonly':
        this.#input.toggleAttribute('readonly', newValue !== null);
        break;

      case 'required':
        this.#input.toggleAttribute('required', newValue !== null);
        break;
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
