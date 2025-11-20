import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Label
 *
 * @attr {string} color - Label color (blue, teal, green, orange, purple, red, orangered, grey, yellow)
 * @attr {'outline'|'filled'|'overflow'|'add'} variant - Label variant - default: filled
 * @attr {string} status - Status with icon (success, warning, danger, info, custom) - overrides color
 * @attr {boolean} compact - Compact size
 * @attr {boolean} disabled - Disabled state
 * @attr {boolean} editable - Editable label
 * @attr {string} href - Makes label a link
 *
 * @slot - Default slot for label text
 * @slot icon - Slot for icon (overrides status icon)
 * @slot close-btn - Slot for custom close button
 * @customElement pf-v6-label
 */
class PfV6Label extends CemElement {
  static is = 'pf-v6-label';
  static observedAttributes = ['color', 'variant', 'status', 'compact', 'disabled', 'editable', 'href'];

  #$ = selector => this.shadowRoot.querySelector(selector);
  #content;
  #text;
  #actions;

  get compact() { return this.getAttribute('compact') ?? false; }
  set compact(value) { return this.toggleAttribute('compact', !!value); }

  get disabled() { return this.getAttribute('disabled') ?? false; }
  set disabled(value) { return this.toggleAttribute('disabled', !!value); }

  get editable() { return this.getAttribute('editable') ?? false; }
  set editable(value) { return this.toggleAttribute('editable', !!value); }

  get color() {
    return this.getAttribute('color') || 'grey';
  }

  set color(value) {
    if (value) {
      this.setAttribute('color', value);
    } else {
      this.removeAttribute('color');
    }
  }

  get variant() {
    return this.getAttribute('variant') || 'filled';
  }

  set variant(value) {
    if (value) {
      this.setAttribute('variant', value);
    } else {
      this.removeAttribute('variant');
    }
  }

  get status() {
    return this.getAttribute('status') || '';
  }

  set status(value) {
    if (value) {
      this.setAttribute('status', value);
    } else {
      this.removeAttribute('status');
    }
  }

  async afterTemplateLoaded() {
    this.#content = this.#$('#content');
    this.#text = this.#$('#text');
    this.#actions = this.#$('#actions');

    // Hide actions if empty
    this.#updateActionsVisibility();
    this.#$('slot[name="actions"]')?.addEventListener('slotchange', () => {
      this.#updateActionsVisibility();
    });

    // Set up click handler for clickable labels
    this.#setupClickHandler();
  }

  #setupClickHandler() {
    // Label is clickable if it has href or onclick
    const isClickable = this.hasAttribute('href') || this.onclick;

    if (isClickable && this.#content) {
      this.#content.addEventListener('click', (e) => {
        if (this.hasAttribute('disabled')) {
          e.preventDefault();
          return;
        }
      });
    }
  }

  #updateActionsVisibility() {
    const actionsSlot = this.#$('slot[name="actions"]');
    const hasActions = actionsSlot?.assignedElements().length > 0;
    if (this.#actions) {
      this.#actions.hidden = !hasActions;
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
