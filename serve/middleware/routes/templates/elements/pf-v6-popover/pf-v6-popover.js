import { CemElement } from '/__cem/cem-element.js';

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
 * @attr {string} position - Position relative to trigger: auto|top|bottom|left|right|top-start|top-end|bottom-start|bottom-end|left-start|left-end|right-start|right-end (default: "top")
 * @attr {number} distance - Distance from trigger in pixels (default: 8)
 * @attr {string} trigger-action - How to trigger: click|hover (default: "click")
 * @attr {boolean} closeable - Show close button (default: true)
 * @attr {string} close-button-label - Aria label for close button (default: "Close")
 * @attr {string} header-level - Heading level: h1-h6 (default: "h6")
 * @attr {string} min-width - Minimum width CSS value
 * @attr {string} max-width - Maximum width CSS value
 * @attr {boolean} has-auto-width - Remove width constraints
 * @attr {string} aria-label - Accessible label (required when no header)
 * @attr {boolean} open - Whether the popover is currently open (reflects state)
 *
 * @fires pf-popover-show - Fired when popover opens
 * @fires pf-popover-hide - Fired when popover closes
 *
 * @customElement pf-v6-popover
 */
class PfV6Popover extends CemElement {
  static is = 'pf-v6-popover';

  static get observedAttributes() {
    return [
      'open',
      'position',
      'distance',
      'trigger-action',
      'closeable',
      'close-button-label',
      'header-level',
      'min-width',
      'max-width',
      'has-auto-width',
      'aria-label'
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

  async afterTemplateLoaded() {
    this.#triggerSlot = this.shadowRoot.querySelector('slot[name="trigger"]');
    this.#popoverSlot = this.shadowRoot.getElementById('content');
    this.#popover = this.shadowRoot.getElementById('popover');
    this.#closeButton = this.shadowRoot.getElementById('close');
    this.#headerElement = this.shadowRoot.getElementById('header-content');

    if (!this.#triggerSlot || !this.#popoverSlot) return;

    // The popover attribute is already set in HTML, no need to set it here

    // Apply position class for polyfill compatibility
    // The polyfill may have trouble with :host() selectors
    this.#updatePositionClass();

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

    // Update dynamic heading level
    this.#updateHeaderLevel();

    // Update ARIA attributes
    this.#updateAriaAttributes();

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

  get position() {
    return this.getAttribute('position') || 'top';
  }

  set position(value) {
    this.setAttribute('position', value);
  }

  get distance() {
    return parseInt(this.getAttribute('distance') || '8');
  }

  set distance(value) {
    this.setAttribute('distance', String(value));
  }

  get triggerAction() {
    return this.getAttribute('trigger-action') || 'click';
  }

  set triggerAction(value) {
    this.setAttribute('trigger-action', value);
  }

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

  get closeButtonLabel() {
    return this.getAttribute('close-button-label') || 'Close';
  }

  set closeButtonLabel(value) {
    this.setAttribute('close-button-label', value);
  }

  get headerLevel() {
    return this.getAttribute('header-level') || 'h6';
  }

  set headerLevel(value) {
    this.setAttribute('header-level', value);
  }

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

      // Re-apply polyfill when showing the popover, as it may need to calculate positions
      // for the now-visible element
      if (this.shadowRoot && !CSS.supports?.('anchor-name', '--test')) {
        // Small delay to let the popover become visible
        await new Promise(resolve => setTimeout(resolve, 0));
        const module = await import('/__cem/css-anchor-positioning-fn.js');
        if (module.default) {
          await module.default({ roots: [this.shadowRoot] });
        }
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
    this.dispatchEvent(new CustomEvent(`pf-popover-${isOpen ? 'show' : 'hide'}`, {
      bubbles: true,
      composed: true,
      detail: { open: isOpen }
    }));
  };

  #updateHeaderLevel() {
    if (!this.#headerElement) return;

    // Create new heading element with correct level
    const level = this.headerLevel;
    const newHeading = document.createElement(level);
    newHeading.className = 'pf-v6-c-popover__title-text';
    newHeading.innerHTML = '<slot name="header"></slot>';

    // Replace existing heading
    const oldHeading = this.#headerElement.querySelector('[class*="pf-v6-c-popover__title-text"]');
    if (oldHeading) {
      this.#headerElement.replaceChild(newHeading, oldHeading);
    } else {
      this.#headerElement.appendChild(newHeading);
    }
  }

  #updateAriaAttributes() {
    if (!this.#popoverSlot) return;

    // Set role and aria-modal on the popover element
    this.#popoverSlot.setAttribute('role', 'dialog');
    this.#popoverSlot.setAttribute('aria-modal', 'true');

    // Set aria-label or aria-labelledby
    const headerSlot = this.shadowRoot.querySelector('slot[name="header"]');
    const hasHeader = headerSlot?.assignedElements()?.length > 0;

    if (hasHeader) {
      this.#popoverSlot.setAttribute('aria-labelledby', 'header-content');
      this.#popoverSlot.removeAttribute('aria-label');
    } else if (this.ariaLabel) {
      this.#popoverSlot.setAttribute('aria-label', this.ariaLabel);
      this.#popoverSlot.removeAttribute('aria-labelledby');
    }

    // Always set aria-describedby
    this.#popoverSlot.setAttribute('aria-describedby', 'body');

    // Update close button aria-label
    if (this.#closeButton) {
      this.#closeButton.setAttribute('aria-label', this.closeButtonLabel);
    }
  }

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

      case 'header-level':
        this.#updateHeaderLevel();
        break;

      case 'aria-label':
      case 'close-button-label':
        this.#updateAriaAttributes();
        break;

      case 'min-width':
      case 'max-width':
      case 'distance':
        this.#updateCSSProperties();
        break;
    }
  }

  disconnectedCallback() {
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
