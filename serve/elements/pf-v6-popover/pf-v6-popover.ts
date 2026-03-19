import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-popover.css' with { type: 'css' };

type PopoverPosition =
  | 'auto' | 'top' | 'bottom' | 'left' | 'right'
  | 'top-start' | 'top-end' | 'bottom-start' | 'bottom-end'
  | 'left-start' | 'left-end' | 'right-start' | 'right-end';

type TriggerAction = 'click' | 'hover';

/**
 * Custom event fired when popover shows
 */
export class PfPopoverShowEvent extends Event {
  open = true;
  constructor() {
    super('pf-popover-show', { bubbles: true });
  }
}

/**
 * Custom event fired when popover hides
 */
export class PfPopoverHideEvent extends Event {
  open = false;
  constructor() {
    super('pf-popover-hide', { bubbles: true });
  }
}

// Cache feature detection for CSS Anchor Positioning
const supportsAnchorPositioning = CSS.supports?.('anchor-name: --test') ?? false;

/**
 * PatternFly v6 Popover
 *
 * A popover component using the native Popover API and CSS Anchor Positioning.
 *
 * @slot trigger - The element that triggers the popover (usually a button)
 * @slot header - Header content
 * @slot - Default slot for popover body content
 * @slot footer - Footer content
 *
 * @fires pf-popover-show - Fired when popover opens
 * @fires pf-popover-hide - Fired when popover closes
 */
@customElement('pf-v6-popover')
export class PfV6Popover extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor open = false;

  @property({ reflect: true })
  accessor position: PopoverPosition = 'top';

  @property({ type: Number, reflect: true })
  accessor distance = 8;

  @property({ reflect: true, attribute: 'trigger-action' })
  accessor triggerAction: TriggerAction = 'click';

  @property({ type: Boolean, reflect: true })
  accessor closeable = true;

  @property({ attribute: 'close-button-label' })
  accessor closeButtonLabel = 'Close';

  @property({ attribute: 'min-width' })
  accessor minWidth?: string;

  @property({ attribute: 'max-width' })
  accessor maxWidth?: string;

  @property({ type: Boolean, reflect: true, attribute: 'has-auto-width' })
  accessor hasAutoWidth = false;

  #hoverShowTimeout?: ReturnType<typeof setTimeout>;
  #hoverHideTimeout?: ReturnType<typeof setTimeout>;
  #hasHeaderContent = false;
  #hasFooterContent = false;

  render() {
    return html`
      <div id="trigger"
           @click=${this.triggerAction === 'click' ? this.#handleTriggerClick : nothing}
           @pointerenter=${this.triggerAction === 'hover' ? this.#handlePointerEnter : nothing}
           @pointerleave=${this.triggerAction === 'hover' ? this.#handlePointerLeave : nothing}>
        <slot name="trigger"></slot>
      </div>

      <div id="content"
           class="position-${this.position}"
           popover="manual"
           part="popover"
           style=${this.#contentStyle}
           @pointerenter=${this.triggerAction === 'hover' ? this.#handlePopoverPointerEnter : nothing}
           @pointerleave=${this.triggerAction === 'hover' ? this.#handlePopoverPointerLeave : nothing}
           @toggle=${this.#handlePopoverToggle}>
        <div id="arrow"></div>
        <div id="popover"
             part="content">
          <button id="close"
                  part="close-button"
                  aria-label=${this.closeButtonLabel}
                  @click=${this.#handleClose}>
            <svg class="pf-v6-svg"
                 viewBox="0 0 352 512"
                 fill="currentColor"
                 role="presentation"
                 width="1em"
                 height="1em">
              <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
            </svg>
          </button>

          <header id="header"
                  class=${this.#hasHeaderContent ? 'has-content' : ''}
                  part="header">
            <div id="title">
              <div id="header-content">
                <slot name="header"
                      @slotchange=${this.#onHeaderSlotChange}></slot>
              </div>
            </div>
          </header>

          <div id="body"
               part="body">
            <slot></slot>
          </div>

          <footer id="footer"
                  class=${this.#hasFooterContent ? 'has-content' : ''}
                  part="footer">
            <slot name="footer"
                  @slotchange=${this.#onFooterSlotChange}></slot>
          </footer>
        </div>
      </div>
    `;
  }

  get #contentStyle(): string {
    const parts: string[] = [];
    if (this.minWidth) parts.push(`--min-width: ${this.minWidth}`);
    if (this.maxWidth) parts.push(`--max-width: ${this.maxWidth}`);
    parts.push(`--distance: ${this.distance}px`);
    return parts.join('; ');
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('open')) {
      const contentEl = this.shadowRoot?.getElementById('content') as HTMLElement & { showPopover(): void; hidePopover(): void; matches(s: string): boolean } | null;
      if (!contentEl) return;
      if (this.open && !contentEl.matches(':popover-open')) {
        this.#showPopover();
      } else if (!this.open && contentEl.matches(':popover-open')) {
        this.#hidePopover();
      }
    }
  }

  toggle() {
    if (this.open) {
      this.#hidePopover();
    } else {
      this.#showPopover();
    }
  }

  async #showPopover() {
    const contentEl = this.shadowRoot?.getElementById('content') as HTMLElement & { showPopover(): void } | null;
    if (!contentEl) return;
    try {
      contentEl.showPopover();

      if (!supportsAnchorPositioning) {
        await new Promise(resolve => requestAnimationFrame(resolve));

        const trigger = this.shadowRoot?.getElementById('trigger');
        if (!trigger) return;
        const triggerRect = trigger.getBoundingClientRect();
        const popoverRect = contentEl.getBoundingClientRect();
        const dist = this.distance;

        let top: number, left: number;

        switch (this.position) {
          case 'top':
            top = triggerRect.top - popoverRect.height - dist;
            left = triggerRect.left + (triggerRect.width / 2) - (popoverRect.width / 2);
            break;
          case 'top-start':
            top = triggerRect.top - popoverRect.height - dist;
            left = triggerRect.left;
            break;
          case 'top-end':
            top = triggerRect.top - popoverRect.height - dist;
            left = triggerRect.right - popoverRect.width;
            break;
          case 'bottom':
            top = triggerRect.bottom + dist;
            left = triggerRect.left + (triggerRect.width / 2) - (popoverRect.width / 2);
            break;
          case 'bottom-start':
            top = triggerRect.bottom + dist;
            left = triggerRect.left;
            break;
          case 'bottom-end':
            top = triggerRect.bottom + dist;
            left = triggerRect.right - popoverRect.width;
            break;
          case 'left':
            top = triggerRect.top + (triggerRect.height / 2) - (popoverRect.height / 2);
            left = triggerRect.left - popoverRect.width - dist;
            break;
          case 'left-start':
            top = triggerRect.top;
            left = triggerRect.left - popoverRect.width - dist;
            break;
          case 'left-end':
            top = triggerRect.bottom - popoverRect.height;
            left = triggerRect.left - popoverRect.width - dist;
            break;
          case 'right':
            top = triggerRect.top + (triggerRect.height / 2) - (popoverRect.height / 2);
            left = triggerRect.right + dist;
            break;
          case 'right-start':
            top = triggerRect.top;
            left = triggerRect.right + dist;
            break;
          case 'right-end':
            top = triggerRect.bottom - popoverRect.height;
            left = triggerRect.right + dist;
            break;
          default:
            top = triggerRect.top - popoverRect.height - dist;
            left = triggerRect.left + (triggerRect.width / 2) - (popoverRect.width / 2);
        }

        contentEl.style.top = `${top}px`;
        contentEl.style.left = `${left}px`;
        contentEl.style.bottom = 'auto';
        contentEl.style.right = 'auto';
        contentEl.style.translate = 'none';
      }
    } catch (e) {
      console.warn('Failed to show popover:', e);
    }
  }

  #hidePopover() {
    const contentEl = this.shadowRoot?.getElementById('content') as HTMLElement & { hidePopover(): void } | null;
    try {
      contentEl?.hidePopover();
    } catch (e) {
      console.warn('Failed to hide popover:', e);
    }
  }

  #handleTriggerClick = (e: Event) => {
    const triggerSlot = this.shadowRoot?.querySelector('slot[name="trigger"]') as HTMLSlotElement | null;
    const assigned = triggerSlot?.assignedElements();
    if (assigned && assigned.length > 0) {
      e.stopPropagation();
      this.toggle();
    }
  };

  #handlePointerEnter = () => {
    clearTimeout(this.#hoverHideTimeout);
    this.#hoverShowTimeout = setTimeout(() => {
      this.#showPopover();
    }, 150);
  };

  #handlePointerLeave = () => {
    clearTimeout(this.#hoverShowTimeout);
    this.#hoverHideTimeout = setTimeout(() => {
      this.#hidePopover();
    }, 300);
  };

  #handlePopoverPointerEnter = () => {
    clearTimeout(this.#hoverHideTimeout);
  };

  #handlePopoverPointerLeave = () => {
    this.#hoverHideTimeout = setTimeout(() => {
      this.#hidePopover();
    }, 300);
  };

  #handlePopoverToggle = (e: ToggleEvent) => {
    const isOpen = e.newState === 'open';

    if (isOpen !== this.open) {
      this.open = isOpen;
    }

    const triggerSlot = this.shadowRoot?.querySelector('slot[name="trigger"]') as HTMLSlotElement | null;
    const triggerButton = triggerSlot?.assignedElements()?.[0];
    if (triggerButton) {
      triggerButton.setAttribute('aria-expanded', String(isOpen));
    }

    this.dispatchEvent(isOpen ? new PfPopoverShowEvent() : new PfPopoverHideEvent());
  };

  #handleClose = (e: Event) => {
    e.stopPropagation();
    this.#hidePopover();
  };

  #onHeaderSlotChange(e: Event) {
    const slot = e.target as HTMLSlotElement;
    this.#hasHeaderContent = slot.assignedNodes().length > 0;
    this.requestUpdate();
  }

  #onFooterSlotChange(e: Event) {
    const slot = e.target as HTMLSlotElement;
    this.#hasFooterContent = slot.assignedNodes().length > 0;
    this.requestUpdate();
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    clearTimeout(this.#hoverShowTimeout);
    clearTimeout(this.#hoverHideTimeout);
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-popover': PfV6Popover;
  }
}
