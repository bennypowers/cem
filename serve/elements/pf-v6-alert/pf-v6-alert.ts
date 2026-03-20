import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-alert.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

type AlertVariant = 'custom' | 'success' | 'danger' | 'warning' | 'info';

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

const ICON_PATHS: Record<AlertVariant, string> = {
  success: 'M504 256c0 136.967-111.033 248-248 248S8 392.967 8 256 119.033 8 256 8s248 111.033 248 248zM227.314 387.314l184-184c6.248-6.248 6.248-16.379 0-22.627l-22.627-22.627c-6.248-6.249-16.379-6.249-22.628 0L216 308.118l-70.059-70.059c-6.248-6.248-16.379-6.248-22.628 0l-22.627 22.627c-6.248 6.248-6.248 16.379 0 22.627l104 104c6.249 6.249 16.379 6.249 22.628.001z',
  danger: 'M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42-42-18.804-42-42 18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z',
  warning: 'M569.517 440.013C587.975 472.007 564.806 512 527.94 512H48.054c-36.937 0-59.999-40.055-41.577-71.987L246.423 23.985c18.467-32.009 64.72-31.951 83.154 0l239.94 416.028zM288 354c-25.405 0-46 20.595-46 46s20.595 46 46 46 46-20.595 46-46-20.595-46-46-46zm-43.673-165.346l7.418 136c.347 6.364 5.609 11.346 11.982 11.346h48.546c6.373 0 11.635-4.982 11.982-11.346l7.418-136c.375-6.874-5.098-12.654-11.982-12.654h-63.383c-6.884 0-12.356 5.78-11.981 12.654z',
  info: 'M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42-42-18.804-42-42 18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z',
  custom: 'M464 32H48C21.49 32 0 53.49 0 80v352c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V80c0-26.51-21.49-48-48-48zm-6 400H54a6 6 0 0 1-6-6V86a6 6 0 0 1 6-6h404a6 6 0 0 1 6 6v340a6 6 0 0 1-6 6zm-42-92v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm0-96v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm0-96v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm-252 12c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36zm0 96c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36zm0 96c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36z',
};

const VARIANT_LABELS: Record<AlertVariant, string> = {
  success: 'Success alert:',
  danger: 'Danger alert:',
  warning: 'Warning alert:',
  info: 'Info alert:',
  custom: 'Custom alert:',
};

/**
 * PatternFly Alert component
 *
 * @fires expand - Fired when expandable alert is expanded
 * @fires collapse - Fired when expandable alert is collapsed
 * @fires close - Fired when alert close action is triggered
 *
 * @slot icon - Custom icon (overrides default variant icon)
 * @slot - Default slot for alert title text
 * @slot description - Alert description text
 * @slot action - Single action button/link (not available when dismissable - close button renders in shadow)
 * @slot actions - Action group for multiple action links/buttons
 */
@customElement('pf-v6-alert')
export class PfV6Alert extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor variant: AlertVariant = 'custom';

  @property({ type: Boolean, reflect: true })
  accessor inline = false;

  @property({ type: Boolean, reflect: true })
  accessor plain = false;

  @property({ type: Boolean, reflect: true })
  accessor expandable = false;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  @property({ type: Number, reflect: true })
  accessor truncate?: number;

  @property({ type: Boolean, reflect: true })
  accessor dismissable = false;

  @property({ type: Boolean, reflect: true, attribute: 'live-region' })
  accessor liveRegion = false;

  get #iconPath(): string {
    return ICON_PATHS[this.variant] ?? ICON_PATHS.custom;
  }

  get #variantLabel(): string {
    return VARIANT_LABELS[this.variant] ?? VARIANT_LABELS.custom;
  }

  get #closeButtonLabel(): string {
    const titleText = this.textContent?.trim() ?? '';
    return titleText
      ? `Close ${this.#variantLabel} ${titleText}`
      : `Close ${this.#variantLabel.replace(':', '')}`;
  }

  render() {
    return html`
      ${this.expandable ? html`
        <div id="toggle-container">
          <pf-v6-button id="toggle"
                        variant="plain"
                        aria-expanded="${this.expanded}"
                        aria-label="Toggle alert"
                        @click=${this.toggle}>
            <svg id="toggle-icon"
                 aria-hidden="true"
                 fill="currentColor"
                 height="1em"
                 width="1em"
                 viewBox="0 0 320 512">
              <path d="M285.476 272.971L91.132 467.314c-9.373 9.373-24.569 9.373-33.941 0l-22.667-22.667c-9.357-9.357-9.375-24.522-.04-33.901L188.505 256 34.484 101.255c-9.335-9.379-9.317-24.544.04-33.901l22.667-22.667c9.373-9.373 24.569-9.373 33.941 0L285.475 239.03c9.373 9.372 9.373 24.568.001 33.941z"></path>
            </svg>
          </pf-v6-button>
        </div>
      ` : nothing}

      <div id="icon-container">
        <slot name="icon">
          <svg viewBox="0 0 ${this.variant === 'warning' ? '576' : '512'} 512"
               fill="currentColor"
               aria-hidden="true"
               role="img"
               width="1em"
               height="1em">
            <path d="${this.#iconPath}"></path>
          </svg>
        </slot>
      </div>

      <h4 id="title"
          style="${this.truncate ? `--pf-v6-c-alert__title--max-lines: ${this.truncate}` : ''}">
        <span id="sr-variant">${this.#variantLabel}</span>
        <slot></slot>
      </h4>

      <div id="action">
        ${this.dismissable ? html`
          <pf-v6-button id="close"
                        variant="plain"
                        aria-label="${this.#closeButtonLabel}"
                        @click=${this.#onClose}>
            <svg viewBox="0 0 352 512"
                 fill="currentColor"
                 aria-hidden="true"
                 role="img"
                 width="1em"
                 height="1em">
              <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
            </svg>
          </pf-v6-button>
        ` : html`
          <slot name="action"></slot>
        `}
      </div>

      ${!this.expandable || this.expanded ? html`
        <div id="description">
          <slot name="description"></slot>
        </div>
      ` : nothing}

      <div id="action-group">
        <slot name="actions"></slot>
      </div>
    `;
  }

  #onClose() {
    if (this.dispatchEvent(new PfAlertCloseEvent())) {
      this.remove();
    }
  }

  /** Toggle expanded state (for expandable alerts) */
  toggle() {
    if (!this.expandable) return;
    this.expanded = !this.expanded;
    this.dispatchEvent(this.expanded ? new PfAlertExpandEvent() : new PfAlertCollapseEvent());
  }

  /** Expand the alert (for expandable alerts) */
  expand() {
    if (!this.expandable || this.expanded) return;
    this.expanded = true;
    this.dispatchEvent(new PfAlertExpandEvent());
  }

  /** Collapse the alert (for expandable alerts) */
  collapse() {
    if (!this.expandable || !this.expanded) return;
    this.expanded = false;
    this.dispatchEvent(new PfAlertCollapseEvent());
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-alert': PfV6Alert;
  }
}
