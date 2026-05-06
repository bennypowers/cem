import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-pf-v6-dropdown.css' with { type: 'css' };

import '../cem-pf-v6-button/cem-pf-v6-button.js';
import '../cem-pf-v6-menu/cem-pf-v6-menu.js';
import '../cem-pf-v6-menu-item/cem-pf-v6-menu-item.js';

import type { PfV6Menu } from '../cem-pf-v6-menu/cem-pf-v6-menu.js';

/**
 * PatternFly v6 Dropdown component
 *
 * @fires {Event} expand - Fires when dropdown opens
 * @fires {Event} collapse - Fires when dropdown closes
 *
 * @slot toggle-text - Text content for toggle button
 * @slot - Menu items (cem-pf-v6-menu-item elements)
 */
@customElement('cem-pf-v6-dropdown')
export class PfV6Dropdown extends LitElement {
  static readonly #instances = new Set<PfV6Dropdown>();

  static {
    document?.addEventListener?.('click', (event: Event) => {
      for (const instance of PfV6Dropdown.#instances) {
        if (instance.expanded && !event.composedPath().includes(instance)) {
          instance.collapse();
        }
      }
    });
  }

  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property()
  accessor label = '';

  connectedCallback() {
    super.connectedCallback();
    PfV6Dropdown.#instances.add(this);
    this.addEventListener('keydown', this.#onKeydown);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    PfV6Dropdown.#instances.delete(this);
    this.removeEventListener('keydown', this.#onKeydown);
  }

  render() {
    return html`
      <cem-pf-v6-button id="toggle"
                     variant="tertiary"
                     aria-haspopup="true"
                     aria-expanded="${this.expanded}"
                     ?disabled=${this.disabled}
                     @click=${this.#onToggleClick}>
        <slot name="toggle-text">Toggle</slot>
        <svg slot="icon-end"
             viewBox="0 0 320 512"
             aria-hidden="true">
          <path d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"/>
        </svg>
      </cem-pf-v6-button>

      <div id="menu-container"
           ?hidden=${!this.expanded}>
        <cem-pf-v6-menu id="menu"
                     label=${this.label}>
          <slot></slot>
        </cem-pf-v6-menu>
      </div>
    `;
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('expanded')) {
      this.dispatchEvent(new Event(this.expanded ? 'expand' : 'collapse', { bubbles: true }));
      if (this.expanded) {
        requestAnimationFrame(() => {
          const menu = this.shadowRoot?.getElementById('menu') as PfV6Menu | null;
          menu?.focusFirstItem();
        });
      }
    }
  }

  #onKeydown = (event: KeyboardEvent) => {
    if (event.key === 'Escape' && this.expanded) {
      event.preventDefault();
      this.collapse();
      (this.shadowRoot?.getElementById('toggle') as HTMLElement)?.focus();
    }
  };

  #onToggleClick() {
    if (this.disabled) return;
    this.toggle();
  }

  /** Toggle expanded state */
  toggle() {
    this.expanded = !this.expanded;
  }

  /** Expand the dropdown */
  expand() {
    this.expanded = true;
  }

  /** Collapse the dropdown */
  collapse() {
    this.expanded = false;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-pf-v6-dropdown': PfV6Dropdown;
  }
}
