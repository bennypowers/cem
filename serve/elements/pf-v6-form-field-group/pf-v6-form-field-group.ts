import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-form-field-group.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

/**
 * Custom event fired when form field group toggles
 */
export class PfFormFieldGroupToggleEvent extends Event {
  expanded: boolean;
  constructor(expanded: boolean) {
    super('toggle', { bubbles: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Form Field Group
 *
 * A collapsible section within a form, with optional toggle and header.
 *
 * @slot - Form fields/groups to display in the body
 * @slot header-actions - Action buttons in the header
 *
 * @attr {boolean} expandable - Whether the field group can be expanded/collapsed
 * @attr {boolean} expanded - Whether the field group is expanded (only used with expandable)
 * @attr {string} toggle-text - Text for the title
 * @attr {string} description - Description text shown below the title
 *
 * @fires toggle - Fires when expansion state changes
 */
@customElement('pf-v6-form-field-group')
export class PfV6FormFieldGroup extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor expandable = false;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  @property({ reflect: true, attribute: 'toggle-text' })
  accessor toggleText?: string;

  @property({ reflect: true })
  accessor description?: string;

  render() {
    const hasToggle = this.expandable || this.toggleText;
    const hasHeader = this.toggleText || this.description;

    return html`
      ${hasToggle ? html`
        <pf-v6-button id="toggle-button"
                      part="toggle-button"
                      variant="plain"
                      aria-expanded=${String(this.expanded)}
                      aria-controls="body"
                      aria-label="Details"
                      @click=${this.toggle}>
          <svg id="toggle-icon"
               viewBox="0 0 256 512"
               fill="currentColor"
               role="presentation"
               width="1em"
               height="1em">
            <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
          </svg>
        </pf-v6-button>
      ` : nothing}
      ${hasHeader ? html`
        <div id="header"
             part="header">
          <div id="header-main">
            ${this.toggleText ? html`
              <div id="header-title">
                <div id="header-title-text">
                  <span id="toggle-text-default">${this.toggleText}</span>
                </div>
              </div>
            ` : nothing}
            ${this.description ? html`
              <div id="header-description">${this.description}</div>
            ` : nothing}
          </div>
          <div id="header-actions">
            <slot name="header-actions"></slot>
          </div>
        </div>
      ` : nothing}
      <div id="body"
           part="body"
           ?inert=${this.expandable && !this.expanded}>
        <slot></slot>
      </div>
    `;
  }

  /** Toggles the expanded state */
  toggle() {
    if (this.expandable) {
      this.expanded = !this.expanded;
      this.dispatchEvent(new PfFormFieldGroupToggleEvent(this.expanded));
    }
  }

  /** Expands the field group */
  show() {
    if (this.expandable) {
      this.expanded = true;
      this.dispatchEvent(new PfFormFieldGroupToggleEvent(true));
    }
  }

  /** Collapses the field group */
  hide() {
    if (this.expandable) {
      this.expanded = false;
      this.dispatchEvent(new PfFormFieldGroupToggleEvent(false));
    }
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-form-field-group': PfV6FormFieldGroup;
  }
}
