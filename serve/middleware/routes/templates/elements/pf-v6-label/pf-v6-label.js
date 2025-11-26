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

  get compact() { return this.hasAttribute('compact'); }
  set compact(value) { this.toggleAttribute('compact', !!value); }

  get disabled() { return this.hasAttribute('disabled'); }
  set disabled(value) { this.toggleAttribute('disabled', !!value); }

  get editable() { return this.hasAttribute('editable'); }
  set editable(value) { this.toggleAttribute('editable', !!value); }

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

  #clickListener;
  #slotchangeListener;

  async afterTemplateLoaded() {
    this.#content = this.#$('#content');
    this.#text = this.#$('#text');
    this.#actions = this.#$('#actions');

    // Sync href if we have an anchor element
    if (this.#content) {
      this.#syncHref();
    }

    // Hide actions if empty
    this.#updateActionsVisibility();

    // Only attach slotchange listener if actions slot exists
    const actionsSlot = this.#$('slot[name="actions"]');
    if (actionsSlot) {
      this.#slotchangeListener = () => this.#updateActionsVisibility();
      actionsSlot.addEventListener('slotchange', this.#slotchangeListener);
    }

    // Set up click handler (always attached, checks disabled state dynamically)
    this.#setupClickHandler();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue) return;

    if (name === 'href' && this.#content) {
      this.#syncHref();
    }
  }

  #syncHref() {
    if (!this.#content || !this.#content.tagName) return;

    if (this.#content.tagName === 'A') {
      if (this.hasAttribute('href')) {
        this.#content.setAttribute('href', this.getAttribute('href'));
      } else {
        this.#content.removeAttribute('href');
      }
    }
  }

  #setupClickHandler() {
    if (!this.#content) return;

    // Always attach click handler to prevent clicks when disabled
    // This works regardless of when href is added/removed
    this.#clickListener = (e) => {
      if (this.hasAttribute('disabled')) {
        e.preventDefault();
        e.stopPropagation();
      }
    };

    this.#content.addEventListener('click', this.#clickListener);
  }

  #updateActionsVisibility() {
    if (!this.#actions) return;

    const actionsSlot = this.#$('slot[name="actions"]');
    const hasActions = actionsSlot?.assignedElements().length > 0;
    this.#actions.hidden = !hasActions;
  }

  disconnectedCallback() {
    super.disconnectedCallback?.();

    // Clean up event listeners
    if (this.#clickListener && this.#content) {
      this.#content.removeEventListener('click', this.#clickListener);
      this.#clickListener = null;
    }

    if (this.#slotchangeListener) {
      this.#$('slot[name="actions"]')?.removeEventListener('slotchange', this.#slotchangeListener);
      this.#slotchangeListener = null;
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
