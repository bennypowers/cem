import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-nav-link.css' with { type: 'css' };

/**
 * Custom event for navigation toggle
 */
export class PfNavToggleEvent extends Event {
  expanded: boolean;
  constructor(expanded: boolean) {
    super('pf-nav-toggle', { bubbles: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Navigation Link
 *
 * @slot - Default slot for link text
 * @slot icon-start - Icon before text
 * @slot icon-end - Icon after text
 *
 * @fires {PfNavToggleEvent} pf-nav-toggle - When expandable link is clicked
 */
@customElement('pf-v6-nav-link')
export class PfV6NavLink extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor href?: string;

  @property()
  accessor label?: string;

  @property({ type: Boolean, reflect: true })
  accessor current = false;

  @property({ type: Boolean, reflect: true })
  accessor expandable = false;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  connectedCallback() {
    super.connectedCallback();
    this.#markCurrentIfMatches();
  }

  #onClick(e: Event) {
    if (this.expandable) {
      e.preventDefault();
      this.dispatchEvent(new PfNavToggleEvent(!this.expanded));
    }
  }

  #markCurrentIfMatches() {
    if (this.href && this.href === window.location.pathname) {
      this.current = true;
    }
  }

  #renderLinkContent() {
    return html`
      <slot name="icon-start"></slot>
      <slot></slot>
      <slot name="icon-end"></slot>
      <span id="toggle"
            part="toggle">
        <span id="toggle-icon"
              part="toggle-icon">
          <svg class="pf-v6-svg"
               viewBox="0 0 256 512"
               fill="currentColor"
               role="presentation">
            <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
          </svg>
        </span>
      </span>
    `;
  }

  render() {
    if (this.href) {
      return html`
        <a id="link"
           part="link"
           href=${this.href}
           aria-label=${this.label ?? nothing}
           aria-current=${this.current ? 'page' : nothing}
           @click=${this.#onClick}>
          ${this.#renderLinkContent()}
        </a>
      `;
    }
    return html`
      <button id="link"
              part="link"
              type="button"
              aria-label=${this.label ?? nothing}
              aria-expanded=${this.expandable ? String(this.expanded) : nothing}
              @click=${this.#onClick}>
        ${this.#renderLinkContent()}
      </button>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-nav-link': PfV6NavLink;
  }
}
