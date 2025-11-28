import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event fired when drawer expands
 */
export class PfDrawerExpandEvent extends Event {
  constructor() {
    super('expand', { bubbles: true });
  }
}

/**
 * Custom event fired when drawer collapses
 */
export class PfDrawerCollapseEvent extends Event {
  constructor() {
    super('collapse', { bubbles: true });
  }
}

/**
 * Custom event fired when drawer close button is clicked
 */
export class PfDrawerCloseEvent extends Event {
  constructor() {
    super('close', { bubbles: true });
  }
}

/**
 * PatternFly v6 Drawer Component
 *
 * A drawer is a sliding panel that enters from the edge of the screen to display
 * additional content without navigating away from the current page.
 *
 * @attr {boolean} expanded - Whether the drawer panel is expanded
 * @attr {string} panel-position - Position of panel: "end" (default), "start", "bottom"
 * @attr {string} variant - Display variant: "default", "inline", "static"
 * @attr {string} panel-width - Panel width: "25", "33", "50", "66", "75", "100"
 * @attr {string} panel-width-on-lg - Panel width at lg breakpoint
 * @attr {string} panel-width-on-xl - Panel width at xl breakpoint
 * @attr {string} panel-width-on-2xl - Panel width at 2xl breakpoint
 * @attr {boolean} no-border - Remove panel border
 * @attr {boolean} resizable - Make panel resizable
 *
 * @slot - Main content area
 * @slot panel-header - Panel header content
 * @slot panel-description - Panel description
 * @slot panel-body - Panel body content
 *
 * @fires {PfDrawerExpandEvent} expand - When panel expands
 * @fires {PfDrawerCollapseEvent} collapse - When panel collapses
 * @fires {PfDrawerCloseEvent} close - When close button is clicked
 *
 * @customElement pf-v6-drawer
 */
export class PfV6Drawer extends CemElement {
  static is = 'pf-v6-drawer';

  static observedAttributes = [
    'expanded',
    'panel-position',
    'variant',
    'panel-width',
    'panel-width-on-lg',
    'panel-width-on-xl',
    'panel-width-on-2xl',
    'no-border',
    'resizable'
  ];

  #$ = id => this.shadowRoot.getElementById(id);

  #closeButton;
  #panelHeader;
  #panelDescription;
  #panelBody;

  async afterTemplateLoaded() {
    this.#closeButton = this.#$('close');
    this.#panelHeader = this.#$('panel-header');
    this.#panelDescription = this.#$('panel-description');
    this.#panelBody = this.#$('panel-body');

    // Handle close button click
    this.#closeButton?.addEventListener('click', () => {
      this.close();
    });

    // Manage slot visibility
    this.#updateSlotVisibility();

    // Observe slot changes to update visibility
    const slots = this.shadowRoot.querySelectorAll('slot[name]');
    slots.forEach(slot => {
      slot.addEventListener('slotchange', () => this.#updateSlotVisibility());
    });
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue) return;

    if (name === 'expanded') {
      this.#updateExpandedState();
    }
  }

  #updateExpandedState() {
    const isExpanded = this.hasAttribute('expanded');

    // Dispatch appropriate event
    if (isExpanded) {
      this.dispatchEvent(new PfDrawerExpandEvent());
    } else {
      this.dispatchEvent(new PfDrawerCollapseEvent());
    }
  }

  #updateSlotVisibility() {
    if (!this.#panelHeader || !this.#panelDescription || !this.#panelBody) return;

    // Hide panel header if no header slot content
    const headerSlot = this.shadowRoot.querySelector('slot[name="panel-header"]');
    const hasHeaderContent = headerSlot && headerSlot.assignedNodes().length > 0;
    this.#panelHeader.hidden = !hasHeaderContent;

    // Hide panel description if no description slot content
    const descriptionSlot = this.shadowRoot.querySelector('slot[name="panel-description"]');
    const hasDescriptionContent = descriptionSlot && descriptionSlot.assignedNodes().length > 0;
    this.#panelDescription.hidden = !hasDescriptionContent;

    // Hide panel body if no body slot content
    const bodySlot = this.shadowRoot.querySelector('slot[name="panel-body"]');
    const hasBodyContent = bodySlot && bodySlot.assignedNodes().length > 0;
    this.#panelBody.hidden = !hasBodyContent;
  }

  // Public API

  /**
   * Opens the drawer panel
   */
  open() {
    this.expanded = true;
  }

  /**
   * Closes the drawer panel and fires close event
   */
  close() {
    this.expanded = false;
    this.dispatchEvent(new PfDrawerCloseEvent());
  }

  /**
   * Toggles the drawer panel
   */
  toggle() {
    this.expanded = !this.expanded;
  }

  // Getters and setters

  get expanded() {
    return this.hasAttribute('expanded');
  }

  set expanded(val) {
    this.toggleAttribute('expanded', !!val);
  }

  get panelPosition() {
    return this.getAttribute('panel-position') || 'end';
  }

  set panelPosition(val) {
    if (val) {
      this.setAttribute('panel-position', val);
    } else {
      this.removeAttribute('panel-position');
    }
  }

  get variant() {
    return this.getAttribute('variant') || 'default';
  }

  set variant(val) {
    if (val) {
      this.setAttribute('variant', val);
    } else {
      this.removeAttribute('variant');
    }
  }

  get panelWidth() {
    return this.getAttribute('panel-width');
  }

  set panelWidth(val) {
    if (val) {
      this.setAttribute('panel-width', val);
    } else {
      this.removeAttribute('panel-width');
    }
  }

  get noBorder() {
    return this.hasAttribute('no-border');
  }

  set noBorder(val) {
    this.toggleAttribute('no-border', !!val);
  }

  get resizable() {
    return this.hasAttribute('resizable');
  }

  set resizable(val) {
    this.toggleAttribute('resizable', !!val);
  }

  static {
    customElements.define(this.is, this);
  }
}
