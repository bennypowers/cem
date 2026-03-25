import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-modal.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

type ModalVariant = 'small' | 'medium' | 'large';
type ModalPosition = 'top';

/**
 * Custom event fired when modal opens
 */
export class PfModalOpenEvent extends Event {
  constructor() {
    super('open', { bubbles: true });
  }
}

/**
 * Custom event fired when modal closes normally
 */
export class PfModalCloseEvent extends Event {
  returnValue: string;
  constructor(returnValue: string) {
    super('close', { bubbles: true });
    this.returnValue = returnValue;
  }
}

/**
 * Custom event fired when modal is cancelled (ESC key)
 */
export class PfModalCancelEvent extends Event {
  constructor() {
    super('cancel', { bubbles: true });
  }
}

/**
 * PatternFly v6 Modal Component
 *
 * A modal displays important information to a user without requiring them to
 * navigate to a new page. Uses native `<dialog>` element for proper modal
 * behavior.
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
 */
@customElement('pf-v6-modal')
export class PfV6Modal extends LitElement {
  static shadowRootOptions: ShadowRootInit = {
    ...LitElement.shadowRootOptions,
    delegatesFocus: true,
  };

  static styles = styles;

  #dialog: HTMLDialogElement | null = null;
  #cancelling = false;
  #hasHeader = false;
  #hasDescription = false;
  #hasFooter = false;

  /** @see https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/returnValue */
  returnValue = '';

  @property({ reflect: true })
  accessor variant?: ModalVariant;

  @property({ reflect: true })
  accessor position?: ModalPosition;

  @property({ type: Boolean, reflect: true })
  accessor open = false;

  render() {
    return html`
      <dialog id="dialog"
              part="dialog"
              @close=${this.#onDialogClose}
              @cancel=${this.#onDialogCancel}>
        <pf-v6-button id="close"
                      part="close"
                      variant="plain"
                      aria-label="Close"
                      @click=${this.#onCloseClick}>
          <svg viewBox="0 0 352 512"
               role="presentation">
            <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
          </svg>
        </pf-v6-button>
        <header id="header"
                part="header"
                ?hidden=${!this.#hasHeader}>
          <div id="header-main">
            <slot name="header"
                  @slotchange=${this.#onHeaderSlotChange}></slot>
            <div id="description"
                 part="description"
                 ?hidden=${!this.#hasDescription}>
              <slot name="description"
                    @slotchange=${this.#onDescriptionSlotChange}></slot>
            </div>
          </div>
        </header>
        <div id="body"
             part="body">
          <slot></slot>
        </div>
        <footer id="footer"
                part="footer"
                ?hidden=${!this.#hasFooter}>
          <slot name="footer"
                @slotchange=${this.#onFooterSlotChange}></slot>
        </footer>
      </dialog>
    `;
  }

  protected updated(changed: Map<string, unknown>) {
    if (changed.has('open')) {
      this.#syncDialogState();
    }
  }

  protected firstUpdated() {
    this.#dialog = this.shadowRoot!.getElementById('dialog') as HTMLDialogElement;
    this.#syncDialogState();
  }

  #syncDialogState() {
    if (!this.#dialog) return;
    if (this.open) {
      if (!this.#dialog.open) {
        this.#dialog.showModal();
        this.dispatchEvent(new PfModalOpenEvent());
      }
    } else {
      if (this.#dialog.open) {
        this.#dialog.close();
      }
    }
  }

  #onDialogClose = () => {
    // Sync open property (which reflects to attribute)
    if (this.open) {
      this.open = false;
    }

    if (this.#cancelling) {
      this.dispatchEvent(new PfModalCancelEvent());
      this.#cancelling = false;
    } else {
      this.dispatchEvent(new PfModalCloseEvent(this.returnValue));
    }
  };

  #onDialogCancel = (e: Event) => {
    e.preventDefault();
    this.#cancelling = true;
    this.close();
  };

  #onCloseClick = () => {
    this.#cancelling = true;
    this.close('cancel');
  };

  #onHeaderSlotChange = (e: Event) => {
    const slot = e.target as HTMLSlotElement;
    const had = this.#hasHeader;
    this.#hasHeader = slot.assignedNodes().length > 0;
    if (had !== this.#hasHeader) this.requestUpdate();
  };

  #onDescriptionSlotChange = (e: Event) => {
    const slot = e.target as HTMLSlotElement;
    const had = this.#hasDescription;
    this.#hasDescription = slot.assignedNodes().length > 0;
    if (had !== this.#hasDescription) this.requestUpdate();
  };

  #onFooterSlotChange = (e: Event) => {
    const slot = e.target as HTMLSlotElement;
    const had = this.#hasFooter;
    this.#hasFooter = slot.assignedNodes().length > 0;
    if (had !== this.#hasFooter) this.requestUpdate();
  };

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
    this.open = true;
  }

  /**
   * Closes the modal
   * @param returnValue - Optional return value
   */
  close(returnValue?: string) {
    if (typeof returnValue === 'string') {
      this.returnValue = returnValue;
      if (this.#dialog) {
        this.#dialog.returnValue = returnValue;
      }
    }
    this.open = false;
  }

  /**
   * Toggles the modal open/closed
   */
  toggle() {
    if (this.open) {
      this.close();
    } else {
      this.showModal();
    }
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-modal': PfV6Modal;
  }
}
