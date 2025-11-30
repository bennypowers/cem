import { CemElement } from '/__cem/cem-element.js';
import '/__cem/elements/pf-v6-button/pf-v6-button.js';

/**
 * Custom event fired when alert is expanded
 */
export class PfAlertExpandEvent extends Event {
  constructor() {
    super('expand', { bubbles: true });
  }
}

/**
 * Custom event fired when alert is collapsed
 */
export class PfAlertCollapseEvent extends Event {
  constructor() {
    super('collapse', { bubbles: true });
  }
}

/**
 * Custom event fired when alert is closed
 */
export class PfAlertCloseEvent extends Event {
  constructor() {
    super('close', { bubbles: true });
  }
}

/**
 * PatternFly Alert component
 *
 * @fires expand - Fired when expandable alert is expanded
 * @fires collapse - Fired when expandable alert is collapsed
 * @fires close - Fired when alert close action is triggered
 *
 * @slot icon - Custom icon (overrides default variant icon)
 * @slot title - Alert title text
 * @slot - Default slot for alert description
 * @slot action - Alert action button (typically close button)
 * @slot actions - Alert action group (action links/buttons)
 */
export class PfAlert extends CemElement {
  static is = 'pf-v6-alert';

  static observedAttributes = [
    'variant',
    'inline',
    'plain',
    'expandable',
    'expanded',
    'truncate',
    'live-region'
  ];

  // Default icon paths for each variant
  static #iconPaths = {
    success: 'M504 256c0 136.967-111.033 248-248 248S8 392.967 8 256 119.033 8 256 8s248 111.033 248 248zM227.314 387.314l184-184c6.248-6.248 6.248-16.379 0-22.627l-22.627-22.627c-6.248-6.249-16.379-6.249-22.628 0L216 308.118l-70.059-70.059c-6.248-6.248-16.379-6.248-22.628 0l-22.627 22.627c-6.248 6.248-6.248 16.379 0 22.627l104 104c6.249 6.249 16.379 6.249 22.628.001z',
    danger: 'M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42-42-18.804-42-42 18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z',
    warning: 'M569.517 440.013C587.975 472.007 564.806 512 527.94 512H48.054c-36.937 0-59.999-40.055-41.577-71.987L246.423 23.985c18.467-32.009 64.72-31.951 83.154 0l239.94 416.028zM288 354c-25.405 0-46 20.595-46 46s20.595 46 46 46 46-20.595 46-46-20.595-46-46-46zm-43.673-165.346l7.418 136c.347 6.364 5.609 11.346 11.982 11.346h48.546c6.373 0 11.635-4.982 11.982-11.346l7.418-136c.375-6.874-5.098-12.654-11.982-12.654h-63.383c-6.884 0-12.356 5.78-11.981 12.654z',
    info: 'M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42-42-18.804-42-42 18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z',
    custom: 'M464 32H48C21.49 32 0 53.49 0 80v352c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V80c0-26.51-21.49-48-48-48zm-6 400H54a6 6 0 0 1-6-6V86a6 6 0 0 1 6-6h404a6 6 0 0 1 6 6v340a6 6 0 0 1-6 6zm-42-92v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm0-96v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm0-96v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm-252 12c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36zm0 96c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36zm0 96c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36z'
  };

  static #variantLabels = {
    success: 'Success alert:',
    danger: 'Danger alert:',
    warning: 'Warning alert:',
    info: 'Info alert:',
    custom: 'Custom alert:'
  };

  #$(id) {
    return this.shadowRoot?.getElementById(id);
  }

  get variant() {
    return this.getAttribute('variant') || 'custom';
  }

  set variant(value) {
    this.setAttribute('variant', value);
  }

  get expanded() {
    return this.hasAttribute('expanded');
  }

  set expanded(value) {
    this.toggleAttribute('expanded', !!value);
  }

  get truncate() {
    return parseInt(this.getAttribute('truncate') || '0', 10);
  }

  set truncate(value) {
    if (value) {
      this.setAttribute('truncate', value.toString());
    } else {
      this.removeAttribute('truncate');
    }
  }

  async afterTemplateLoaded() {
    // Set default icon based on variant if no custom icon provided
    const iconSlot = this.shadowRoot.querySelector('slot[name="icon"]');
    const hasCustomIcon = iconSlot && iconSlot.assignedNodes().length > 0;

    if (!hasCustomIcon) {
      this.#updateDefaultIcon();
    }

    // Set up toggle button for expandable alerts
    if (this.hasAttribute('expandable')) {
      const toggleButton = this.#$('toggle');
      if (toggleButton) {
        toggleButton.addEventListener('click', () => {
          this.toggle();
        });
      }
    }

    // Apply truncate styles if specified
    if (this.truncate > 0) {
      const title = this.#$('title');
      if (title) {
        title.style.setProperty('--pf-v6-c-alert__title--max-lines', this.truncate.toString());
      }
    }

    // Set up action slot listener for close events
    const actionSlot = this.shadowRoot.querySelector('slot[name="action"]');
    if (actionSlot) {
      actionSlot.addEventListener('slotchange', () => {
        const actions = actionSlot.assignedElements();
        actions.forEach(action => {
          action.addEventListener('click', (e) => {
            // Check if this is a close action
            if (action.hasAttribute('close') || action.getAttribute('aria-label')?.includes('Close')) {
              this.dispatchEvent(new PfAlertCloseEvent());
            }
          });
        });
      });
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.shadowRoot?.firstChild) return;

    if (name === 'variant' && oldValue !== newValue) {
      this.#updateDefaultIcon();
    }

    if (name === 'truncate' && oldValue !== newValue) {
      const title = this.#$('title');
      if (title && newValue) {
        title.style.setProperty('--pf-v6-c-alert__title--max-lines', newValue);
      }
    }
  }

  #updateDefaultIcon() {
    const iconPath = this.#$('default-icon-path');
    if (iconPath) {
      const pathData = PfAlert.#iconPaths[this.variant] || PfAlert.#iconPaths.custom;
      iconPath.setAttribute('d', pathData);
    }
  }

  /**
   * Toggle expanded state (for expandable alerts)
   */
  toggle() {
    if (!this.hasAttribute('expandable')) return;

    const willExpand = !this.expanded;
    this.expanded = willExpand;

    this.dispatchEvent(willExpand ? new PfAlertExpandEvent() : new PfAlertCollapseEvent());
  }

  /**
   * Expand the alert (for expandable alerts)
   */
  expand() {
    if (!this.hasAttribute('expandable')) return;
    if (!this.expanded) {
      this.expanded = true;
      this.dispatchEvent(new PfAlertExpandEvent());
    }
  }

  /**
   * Collapse the alert (for expandable alerts)
   */
  collapse() {
    if (!this.hasAttribute('expandable')) return;
    if (this.expanded) {
      this.expanded = false;
      this.dispatchEvent(new PfAlertCollapseEvent());
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
