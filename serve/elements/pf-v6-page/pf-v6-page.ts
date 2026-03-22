import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-page.css' with { type: 'css' };

import '../pf-v6-masthead/pf-v6-masthead.js';
import { SidebarToggleEvent } from '../pf-v6-masthead/pf-v6-masthead.js';

/**
 * PatternFly v6 Page Layout
 *
 * @slot skip-to-content - Skip to content link
 * @slot masthead - Page header (pf-v6-masthead)
 * @slot sidebar - Page sidebar (pf-v6-page-sidebar)
 * @slot main - Main content area (pf-v6-page-main)
 *
 * @listens sidebar-toggle - Responds to toggle events from masthead
 */
@customElement('pf-v6-page')
export class PfV6Page extends LitElement {
  static readonly match = globalThis.window?.matchMedia?.('(min-width: 75rem)');

  static styles = styles;

  @property({ type: Boolean, reflect: true, attribute: 'sidebar-collapsed' })
  accessor sidebarCollapsed = false;

  #clickOutsideHandler = (event: MouseEvent) => {
    if (!PfV6Page.match.matches && !this.sidebarCollapsed) {
      const sidebar = this.querySelector('pf-v6-page-sidebar');
      const mastheadToggle = this.querySelector('pf-v6-masthead')
        ?.shadowRoot?.getElementById('toggle-button');

      if (!event.composedPath().some(node => node === sidebar || node === mastheadToggle)) {
        this.sidebarCollapsed = true;
        this.dispatchEvent(new SidebarToggleEvent(!this.sidebarCollapsed));
      }
    }
  };

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('sidebar-toggle', this.#onSidebarToggle as EventListener);
    document.body.addEventListener('click', this.#clickOutsideHandler);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('sidebar-toggle', this.#onSidebarToggle as EventListener);
    document.body.removeEventListener('click', this.#clickOutsideHandler);
  }

  render() {
    return html`
      <slot name="skip-to-content"></slot>
      <slot name="masthead"></slot>
      <slot name="sidebar"></slot>
      <slot name="main"></slot>
    `;
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('sidebarCollapsed')) {
      this.#updateSidebarState();
    }
  }

  #onSidebarToggle = (event: Event) => {
    const toggleEvent = event as SidebarToggleEvent;
    this.sidebarCollapsed = !toggleEvent.expanded;
  };

  #updateSidebarState() {
    const isCollapsed = this.sidebarCollapsed;
    const sidebar = this.querySelector('pf-v6-page-sidebar');
    sidebar?.toggleAttribute('collapsed', !!isCollapsed);
    sidebar?.toggleAttribute('expanded', !isCollapsed);

    const masthead = this.querySelector('pf-v6-masthead');
    masthead?.toggleAttribute('sidebar-expanded', !isCollapsed);
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-page': PfV6Page;
  }
}
