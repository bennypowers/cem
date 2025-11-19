import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event fired when modal opens
 */
export class PfModalOpenEvent extends Event {
  constructor() {
    super('open', { bubbles: true, composed: true });
  }
}

/**
 * Custom event fired when modal closes normally
 */
export class PfModalCloseEvent extends Event {
  constructor(returnValue) {
    super('close', { bubbles: true, composed: true });
    this.returnValue = returnValue;
  }
}

/**
 * Custom event fired when modal is cancelled (ESC key)
 */
export class PfModalCancelEvent extends Event {
  constructor() {
    super('cancel', { bubbles: true, composed: true });
  }
}

/**
 * PatternFly v6 Modal Component
 *
 * A modal displays important information to a user without requiring them to navigate to a new page.
 * Uses native <dialog> element for proper modal behavior.
 *
 * @attr {string} variant - Modal width variant: small, medium, large
 * @attr {string} position - Modal position: top (default is centered)
 * @attr {boolean} open - Whether the modal is currently open
 *
 * @slot header - Modal title (typically h2-h6)
 * @slot description - Optional description below title
 * @slot - Default slot for body content
 * @slot footer - Action buttons
 *
 * @fires {PfModalOpenEvent} open - When modal opens
 * @fires {PfModalCloseEvent} close - When modal closes normally
 * @fires {PfModalCancelEvent} cancel - When modal is cancelled (ESC key)
 *
 * @customElement pf-v6-modal
 */
export class PfV6Modal extends CemElement {
  static shadowRootOptions = { mode: 'open', delegatesFocus: true };
  static is = 'pf-v6-modal';
  static observedAttributes = ['open', 'variant', 'position'];

  #dialog;
  #closeButton;
  #header;
  #description;
  #footer;
  #cancelling = false;

  /** @see https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/returnValue */
  returnValue = '';

  async afterTemplateLoaded() {
    this.#dialog = this.shadowRoot.querySelector('#dialog');
    this.#closeButton = this.shadowRoot.querySelector('#close');
    this.#header = this.shadowRoot.querySelector('#header');
    this.#description = this.shadowRoot.querySelector('#description');
    this.#footer = this.shadowRoot.querySelector('#footer');

    if (!this.#dialog) return;

    // Handle native dialog events
    this.#dialog.addEventListener('close', () => {
      this.#onDialogClose();
    });

    this.#dialog.addEventListener('cancel', (e) => {
      e.preventDefault(); // Prevent default close, we'll handle it
      this.#cancelling = true;
      this.close();
    });

    // Manage slot visibility
    this.#updateSlotVisibility();

    // Observe slot changes to update visibility
    const slots = this.shadowRoot.querySelectorAll('slot');
    slots.forEach(slot => {
      slot.addEventListener('slotchange', () => this.#updateSlotVisibility());
    });

    // Sync initial open state
    if (this.hasAttribute('open') && !this.#dialog.open) {
      this.showModal();
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.#dialog) return;

    if (name === 'open' && oldValue !== newValue) {
      if (this.hasAttribute('open')) {
        if (!this.#dialog.open) {
          this.showModal();
        }
      } else {
        if (this.#dialog.open) {
          this.#dialog.close();
        }
      }
    }
  }

  #updateSlotVisibility() {
    if (!this.#header || !this.#description || !this.#footer) return;

    // Hide header if no header slot content
    const headerSlot = this.shadowRoot.querySelector('slot[name="header"]');
    const hasHeaderContent = headerSlot && headerSlot.assignedNodes().length > 0;
    this.#header.hidden = !hasHeaderContent;

    // Hide description if no description slot content
    const descriptionSlot = this.shadowRoot.querySelector('slot[name="description"]');
    const hasDescriptionContent = descriptionSlot && descriptionSlot.assignedNodes().length > 0;
    this.#description.hidden = !hasDescriptionContent;

    // Hide footer if no footer slot content
    const footerSlot = this.shadowRoot.querySelector('slot[name="footer"]');
    const hasFooterContent = footerSlot && footerSlot.assignedNodes().length > 0;
    this.#footer.hidden = !hasFooterContent;
  }

  #onDialogClose() {
    // Sync open attribute
    if (this.hasAttribute('open')) {
      this.removeAttribute('open');
    }

    // Restore body scroll
    document.body.style.overflow = '';

    // Dispatch appropriate event
    if (this.#cancelling) {
      this.dispatchEvent(new PfModalCancelEvent());
      this.#cancelling = false;
    } else {
      this.dispatchEvent(new PfModalCloseEvent(this.returnValue));
    }
  }

  /**
   * Opens the modal
   */
  show() {
    this.showModal();
  }

  /**
   * Opens the modal as a modal dialog (with backdrop)
   */
  showModal() {
    if (!this.#dialog || this.#dialog.open) return;

    // Prevent body scroll
    document.body.style.overflow = 'hidden';

    // Open dialog
    this.#dialog.showModal();

    // Sync open attribute
    if (!this.hasAttribute('open')) {
      this.setAttribute('open', '');
    }

    // Dispatch open event
    this.dispatchEvent(new PfModalOpenEvent());
  }

  /**
   * Closes the modal
   * @param {string} [returnValue] - Optional return value
   */
  close(returnValue) {
    if (!this.#dialog || !this.#dialog.open) return;

    if (typeof returnValue === 'string') {
      this.returnValue = returnValue;
      this.#dialog.returnValue = returnValue;
    }

    this.#dialog.close();
  }

  /**
   * Toggles the modal open/closed
   */
  toggle() {
    if (this.#dialog?.open) {
      this.close();
    } else {
      this.showModal();
    }
  }

  get open() {
    return this.#dialog?.open ?? false;
  }

  set open(value) {
    this.toggleAttribute('open', !!value);
  }

  static {
    customElements.define(this.is, this);
  }
}
