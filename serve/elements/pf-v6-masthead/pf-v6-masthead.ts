import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-masthead.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

/**
 * Custom event for sidebar toggle
 */
export class SidebarToggleEvent extends Event {
  expanded: boolean;
  constructor(expanded: boolean) {
    super('sidebar-toggle', { bubbles: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Masthead
 *
 * @slot logo - Logo or brand content
 * @slot toolbar - Toolbar content (actions, menus, etc.)
 *
 * @fires {SidebarToggleEvent} sidebar-toggle - When the hamburger toggle is clicked
 */
@customElement('pf-v6-masthead')
export class PfV6Masthead extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  @property({ type: Boolean, reflect: true, attribute: 'sidebar-expanded' })
  accessor sidebarExpanded = false;

  constructor() {
    super();
    this.#internals.role = 'banner';
  }

  render() {
    return html`
      <div id="main">
        <div id="toggle">
          <pf-v6-button id="toggle-button"
                        variant="plain"
                        aria-label="Toggle global navigation"
                        aria-expanded=${String(this.sidebarExpanded)}
                        @click=${this.#onToggle}>
            <svg id="hamburger"
                 role="presentation"
                 viewBox="0 0 10 10"
                 width="1em"
                 height="1em"
                 fill="none"
                 stroke="currentColor"
                 stroke-width="1.5"
                 stroke-linecap="round">
              <path d="M1,1 L9,1" />
              <path d="M1,5 L9,5" />
              <path d="M9,9 L1,9" />
            </svg>
          </pf-v6-button>
        </div>
        <div id="brand">
          <slot name="logo"></slot>
        </div>
      </div>
      <div id="content">
        <slot name="toolbar"></slot>
      </div>
    `;
  }

  #onToggle() {
    this.dispatchEvent(new SidebarToggleEvent(!this.sidebarExpanded));
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-masthead': PfV6Masthead;
  }
}
