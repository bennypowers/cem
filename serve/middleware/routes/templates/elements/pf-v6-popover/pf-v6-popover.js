import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event fired when popover shows
 */
export class PfPopoverShowEvent extends Event {
  constructor() {
    super('pf-popover-show', { bubbles: true });
    this.open = true;
  }
}

/**
 * Custom event fired when popover hides
 */
export class PfPopoverHideEvent extends Event {
  constructor() {
    super('pf-popover-hide', { bubbles: true });
    this.open = false;
  }
}

/**
 * PatternFly v6 Popover
 *
 * A popover component using the native Popover API and CSS Anchor Positioning.
 * Fully encapsulated in shadow DOM with slot-based anchoring.
 *
 * @slot trigger - The element that triggers the popover (usually a button)
 * @slot header - Header content
 * @slot - Default slot for popover body content
 * @slot footer - Footer content
 *
 * @attr {boolean} closeable - Show close button (default: true)
 * @attr {boolean} has-auto-width - Remove width constraints
 * @attr {boolean} open - Whether the popover is currently open (reflects state)
 * @attr {number} distance - Distance from trigger in pixels (default: 8)
 * @attr {'auto'|'top'|'bottom'|'left'|'right'|'top-start'|'top-end'|'bottom-start'|'bottom-end'|'left-start'|'left-end'|'right-start'|'right-end'} position - Position relative to trigger (default: "top")
 * @attr {'click'|'hover'} trigger-action - How to trigger: click|hover (default: "click")
 * @attr {string} close-button-label - Aria label for close button (default: "Close")
 * @attr {string} min-width - Minimum width CSS value
 * @attr {string} max-width - Maximum width CSS value
 *
 * @fires pf-popover-show - Fired when popover opens
 * @fires pf-popover-hide - Fired when popover closes
 *
 * @customElement pf-v6-popover
 */
class PfV6Popover extends CemElement {
  static is = 'pf-v6-popover';

  // Cache feature detection for CSS Anchor Positioning
  static #supportsAnchorPositioning = CSS.supports?.('anchor-name: --test') ?? false;

  static get observedAttributes() {
    return [
      'open',
      'position',
      'distance',
      'trigger-action',
      'closeable',
      'close-button-label',
      'min-width',
      'max-width',
      'has-auto-width'
    ];
  }

  // Private fields
  #triggerSlot;
  #popover;
  #closeButton;
  #headerElement;
  #hoverShowTimeout;
  #hoverHideTimeout;
  #popoverSlot;

  #$ = id => this.shadowRoot.getElementById(id);

  async afterTemplateLoaded() {
    this.#triggerSlot = this.shadowRoot.querySelector('slot[name="trigger"]');
    this.#popoverSlot = this.#$('content');
    this.#popover = this.#$('popover');
    this.#closeButton = this.#$('close');
    this.#headerElement = this.#$('header-content');

    if (!this.#triggerSlot || !this.#popoverSlot) return;

    // The popover attribute is already set in HTML, no need to set it here

    // Note: We manually position popovers when native anchor positioning isn't supported
    // See #showPopover() for the fallback positioning logic

    // Check for slotted content and apply classes
    this.#updateSlotClasses();

    // Listen for slot changes
    const headerSlot = this.shadowRoot.querySelector('slot[name="header"]');
    const footerSlot = this.shadowRoot.querySelector('slot[name="footer"]');
    headerSlot?.addEventListener('slotchange', () => this.#updateSlotClasses());
    footerSlot?.addEventListener('slotchange', () => this.#updateSlotClasses());

    // Set up trigger event listeners based on trigger-action
    this.#setupTriggerListeners();

    // Listen to native popover toggle events
    this.#popoverSlot.addEventListener('toggle', this.#handlePopoverToggle);

    // Handle close button
    if (this.#closeButton) {
      this.#closeButton.addEventListener('click', (e) => {
        e.stopPropagation();
        this.#hidePopover();
      });
    }

    // Apply CSS custom properties for width/distance
    this.#updateCSSProperties();
  }

  // Getters and setters for properties
  get open() {
    return this.hasAttribute('open');
  }

  set open(value) {
    if (value) {
      this.setAttribute('open', '');
    } else {
      this.removeAttribute('open');
    }
  }

  get position() { return this.getAttribute('position') || 'top'; }
  set position(value) { this.setAttribute('position', value); }

  get distance() { return parseInt(this.getAttribute('distance') || '8'); }
  set distance(value) { this.setAttribute('distance', String(value)); }

  get triggerAction() { return this.getAttribute('trigger-action') || 'click'; }
  set triggerAction(value) { this.setAttribute('trigger-action', value); }

  get closeable() {
    return this.hasAttribute('closeable') !== false && this.getAttribute('closeable') !== 'false';
  }

  set closeable(value) {
    if (value) {
      this.setAttribute('closeable', '');
    } else {
      this.removeAttribute('closeable');
    }
  }

  get closeButtonLabel() { return this.getAttribute('close-button-label') || 'Close'; }
  set closeButtonLabel(value) { this.setAttribute('close-button-label', value); }

  get ariaLabel() {
    return this.getAttribute('aria-label');
  }

  set ariaLabel(value) {
    if (value) {
      this.setAttribute('aria-label', value);
    } else {
      this.removeAttribute('aria-label');
    }
  }

  // Public methods
  toggle() {
    if (this.open) {
      this.#hidePopover();
    } else {
      this.#showPopover();
    }
  }

  // Private methods
  async #showPopover() {
    try {
      this.#popoverSlot?.showPopover();

      // Manually position the popover since the polyfill doesn't handle dynamic elements well
      if (!PfV6Popover.#supportsAnchorPositioning) {
        // Small delay to let the popover layout
        await new Promise(resolve => requestAnimationFrame(resolve));

        const trigger = this.shadowRoot.getElementById('trigger');
        const triggerRect = trigger.getBoundingClientRect();
        const popoverRect = this.#popoverSlot.getBoundingClientRect();
        const position = this.position;
        const distance = parseInt(this.distance);

        // Calculate position based on position attribute
        let top, left, transform = '';

        switch (position) {
          case 'top':
            top = triggerRect.top - popoverRect.height - distance;
            left = triggerRect.left + (triggerRect.width / 2) - (popoverRect.width / 2);
            break;
          case 'top-start':
            top = triggerRect.top - popoverRect.height - distance;
            left = triggerRect.left;
            break;
          case 'top-end':
            top = triggerRect.top - popoverRect.height - distance;
            left = triggerRect.right - popoverRect.width;
            break;
          case 'bottom':
            top = triggerRect.bottom + distance;
            left = triggerRect.left + (triggerRect.width / 2) - (popoverRect.width / 2);
            break;
          case 'bottom-start':
            top = triggerRect.bottom + distance;
            left = triggerRect.left;
            break;
          case 'bottom-end':
            top = triggerRect.bottom + distance;
            left = triggerRect.right - popoverRect.width;
            break;
          case 'left':
            top = triggerRect.top + (triggerRect.height / 2) - (popoverRect.height / 2);
            left = triggerRect.left - popoverRect.width - distance;
            break;
          case 'left-start':
            top = triggerRect.top;
            left = triggerRect.left - popoverRect.width - distance;
            break;
          case 'left-end':
            top = triggerRect.bottom - popoverRect.height;
            left = triggerRect.left - popoverRect.width - distance;
            break;
          case 'right':
            top = triggerRect.top + (triggerRect.height / 2) - (popoverRect.height / 2);
            left = triggerRect.right + distance;
            break;
          case 'right-start':
            top = triggerRect.top;
            left = triggerRect.right + distance;
            break;
          case 'right-end':
            top = triggerRect.bottom - popoverRect.height;
            left = triggerRect.right + distance;
            break;
          default:
            // Default to top center if position is invalid
            top = triggerRect.top - popoverRect.height - distance;
            left = triggerRect.left + (triggerRect.width / 2) - (popoverRect.width / 2);
        }

        this.#popoverSlot.style.top = `${top}px`;
        this.#popoverSlot.style.left = `${left}px`;
        this.#popoverSlot.style.bottom = 'auto';
        this.#popoverSlot.style.right = 'auto';
        this.#popoverSlot.style.translate = 'none'; // Override CSS translate for manual positioning
      }
    } catch (e) {
      // Popover might already be showing or not supported
      console.warn('Failed to show popover:', e);
    }
  }

  #hidePopover() {
    try {
      this.#popoverSlot?.hidePopover();
    } catch (e) {
      // Popover might already be hidden
      console.warn('Failed to hide popover:', e);
    }
  }

  #setupTriggerListeners() {
    // Remove existing listeners
    this.#triggerSlot?.removeEventListener('click', this.#handleTriggerClick);
    this.#triggerSlot?.removeEventListener('pointerenter', this.#handlePointerEnter);
    this.#triggerSlot?.removeEventListener('pointerleave', this.#handlePointerLeave);
    this.#popoverSlot?.removeEventListener('pointerenter', this.#handlePopoverPointerEnter);
    this.#popoverSlot?.removeEventListener('pointerleave', this.#handlePopoverPointerLeave);

    // Add appropriate listeners based on trigger action
    if (this.triggerAction === 'hover') {
      this.#triggerSlot?.addEventListener('pointerenter', this.#handlePointerEnter);
      this.#triggerSlot?.addEventListener('pointerleave', this.#handlePointerLeave);
      // Keep popover open when hovering over it
      this.#popoverSlot?.addEventListener('pointerenter', this.#handlePopoverPointerEnter);
      this.#popoverSlot?.addEventListener('pointerleave', this.#handlePopoverPointerLeave);
    } else {
      this.#triggerSlot?.addEventListener('click', this.#handleTriggerClick);
    }
  }

  #handleTriggerClick = (e) => {
    // Only toggle if the click came from a slotted element
    const assignedElements = this.#triggerSlot.assignedElements();
    if (assignedElements && assignedElements.length > 0) {
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

  #handlePopoverToggle = (e) => {
    const isOpen = e.newState === 'open';

    // Sync the open attribute
    if (isOpen !== this.open) {
      this.open = isOpen;
    }

    // Update aria-expanded on trigger
    const triggerButton = this.#triggerSlot?.assignedElements()?.[0];
    if (triggerButton) {
      triggerButton.setAttribute('aria-expanded', String(isOpen));
    }

    // Dispatch custom events
    this.dispatchEvent(isOpen ? new PfPopoverShowEvent() : new PfPopoverHideEvent());
  };

  #updateCSSProperties() {
    const minWidth = this.getAttribute('min-width');
    const maxWidth = this.getAttribute('max-width');
    const distance = this.distance;

    if (minWidth) {
      this.style.setProperty('--min-width', minWidth);
    }
    if (maxWidth) {
      this.style.setProperty('--max-width', maxWidth);
    }
    this.style.setProperty('--distance', `${distance}px`);
  }

  #updatePositionClass() {
    if (!this.#popoverSlot) return;

    // Remove all position classes
    const positionClasses = [
      'position-top', 'position-bottom', 'position-left', 'position-right',
      'position-top-start', 'position-top-end',
      'position-bottom-start', 'position-bottom-end',
      'position-left-start', 'position-left-end',
      'position-right-start', 'position-right-end',
      'position-auto'
    ];
    this.#popoverSlot.classList.remove(...positionClasses);

    // Add current position class
    const position = this.position;
    this.#popoverSlot.classList.add(`position-${position}`);
  }

  #updateSlotClasses() {
    // Check header slot
    const headerSlot = this.shadowRoot.querySelector('slot[name="header"]');
    const headerContainer = this.shadowRoot.getElementById('header');
    if (headerSlot && headerContainer) {
      const hasHeaderContent = headerSlot.assignedNodes().length > 0;
      headerContainer.classList.toggle('has-content', hasHeaderContent);
    }

    // Check footer slot
    const footerSlot = this.shadowRoot.querySelector('slot[name="footer"]');
    const footerContainer = this.shadowRoot.getElementById('footer');
    if (footerSlot && footerContainer) {
      const hasFooterContent = footerSlot.assignedNodes().length > 0;
      footerContainer.classList.toggle('has-content', hasFooterContent);
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue) return;

    switch (name) {
      case 'open':
        // Sync with popover state if needed
        if (this.open && !this.#popoverSlot?.matches(':popover-open')) {
          this.#showPopover();
        } else if (!this.open && this.#popoverSlot?.matches(':popover-open')) {
          this.#hidePopover();
        }
        break;

      case 'position':
        this.#updatePositionClass();
        break;

      case 'trigger-action':
        this.#setupTriggerListeners();
        break;

      case 'close-button-label':
        // Update close button label
        if (this.#closeButton) {
          this.#closeButton.setAttribute('aria-label', this.closeButtonLabel);
        }
        break;

      case 'min-width':
      case 'max-width':
      case 'distance':
        this.#updateCSSProperties();
        break;
    }
  }

  disconnectedCallback() {
    super.disconnectedCallback?.();

    // Clean up timeouts
    clearTimeout(this.#hoverShowTimeout);
    clearTimeout(this.#hoverHideTimeout);

    // Remove event listeners
    this.#popoverSlot?.removeEventListener('toggle', this.#handlePopoverToggle);
    this.#triggerSlot?.removeEventListener('click', this.#handleTriggerClick);
    this.#triggerSlot?.removeEventListener('pointerenter', this.#handlePointerEnter);
    this.#triggerSlot?.removeEventListener('pointerleave', this.#handlePointerLeave);
    this.#popoverSlot?.removeEventListener('pointerenter', this.#handlePopoverPointerEnter);
    this.#popoverSlot?.removeEventListener('pointerleave', this.#handlePopoverPointerLeave);
  }

  static {
    customElements.define(this.is, this);
  }
}
