import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-nav-item.css' with { type: 'css' };

/**
 * PatternFly v6 Navigation Item
 *
 * Container for navigation links with optional expandable groups.
 * The child pf-v6-nav-link handles the current state independently.
 *
 * @slot - Default slot for nav-link and optional nav-group
 */
@customElement('pf-v6-nav-item')
export class PfV6NavItem extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  #internals = this.attachInternals();

  constructor() {
    super();
    this.#internals.role = 'listitem';
  }

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('pf-nav-toggle', this.#handleToggle as EventListener);
    // Hydrate: sync nav-item expanded state from direct child nav-link (SSR sets it there)
    const navLink = this.querySelector(':scope > pf-v6-nav-link[expandable]');
    if (navLink?.hasAttribute('expanded') && !this.hasAttribute('expanded')) {
      this.expanded = true;
    }
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('pf-nav-toggle', this.#handleToggle as EventListener);
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('expanded')) {
      this.#updateExpandedState();
    }
  }

  #handleToggle = (event: Event) => {
    event.stopPropagation();
    this.expanded = !this.expanded;
  };

  #updateExpandedState() {
    const isExpanded = this.expanded;

    // Update direct child nav-group only (not nested ones)
    this.querySelector(':scope > pf-v6-nav-group')
      ?.toggleAttribute('hidden', !isExpanded);

    // Update direct child nav-link expanded state only
    // The nav-link will forward this to aria-expanded on its internal element
    this.querySelector(':scope > pf-v6-nav-link[expandable]')
      ?.toggleAttribute('expanded', isExpanded);
  }

  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-nav-item': PfV6NavItem;
  }
}
