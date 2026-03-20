import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { classMap } from 'lit/directives/class-map.js';

import styles from './pf-v6-tree-item.css' with { type: 'css' };

import '../pf-v6-badge/pf-v6-badge.js';

/**
 * Custom event fired when tree item is selected
 */
export class PfTreeItemSelectEvent extends Event {
  constructor() {
    super('select', { bubbles: true });
  }
}

/**
 * Custom event fired when tree item is expanded
 */
export class PfTreeItemExpandEvent extends Event {
  constructor() {
    super('expand', { bubbles: true });
  }
}

/**
 * Custom event fired when tree item is collapsed
 */
export class PfTreeItemCollapseEvent extends Event {
  constructor() {
    super('collapse', { bubbles: true });
  }
}

/**
 * PatternFly v6 Tree Item
 *
 * Individual tree node that can be nested.
 *
 * @slot - Child tree items
 * @slot label - Item label text
 *
 * @fires select - Fires when item is selected
 * @fires expand - Fires when item is expanded
 * @fires collapse - Fires when item is collapsed
 */
@customElement('pf-v6-tree-item')
export class PfV6TreeItem extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor label?: string;

  @property()
  accessor icon?: string;

  @property()
  accessor badge?: string;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  @property({ type: Boolean, reflect: true })
  accessor current = false;

  @property({ type: Boolean, reflect: true, attribute: 'has-children' })
  accessor hasChildren = false;

  /** Track if has-children was explicitly set by the user */
  #explicitHasChildren = false;

  connectedCallback() {
    super.connectedCallback();
    if (this.hasAttribute('has-children')) {
      this.#explicitHasChildren = true;
    }
  }

  render() {
    return html`
      <li id="item"
          role="treeitem"
          tabindex="-1"
          aria-expanded=${this.hasChildren ? String(this.expanded) : nothing}
          aria-selected=${String(this.current)}
          aria-label=${this.label ?? nothing}>
        <div id="content">
          <button id="toggle"
                  type="button"
                  aria-label="Toggle"
                  @click=${this.#onToggleClick}>
            <span id="toggle-icon">
              <svg class="pf-v6-svg"
                   viewBox="0 0 256 512"
                   fill="currentColor"
                   role="presentation"
                   width="1em"
                   height="1em">
                <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
              </svg>
            </span>
          </button>
          <button id="node"
                  class=${classMap({ current: this.current })}
                  @click=${this.#onNodeClick}>
            <span id="node-container">
              <span id="icon"></span>
              <span id="node-text">
                <slot name="label">${this.label}</slot>
              </span>
              <span id="badge-container">
                <pf-v6-badge id="badge">
                  <slot name="badge">${this.badge ?? nothing}</slot>
                </pf-v6-badge>
              </span>
            </span>
          </button>
        </div>
        <ul id="children" role="group">
          <slot @slotchange=${this.#onSlotChange}></slot>
        </ul>
      </li>
    `;
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('icon')) {
      this.#updateIcon();
    }
    if (changed.has('expanded') && changed.get('expanded') !== undefined) {
      this.dispatchEvent(
        this.expanded ? new PfTreeItemExpandEvent() : new PfTreeItemCollapseEvent()
      );
    }
  }

  #onToggleClick(e: Event) {
    e.stopPropagation();
    this.toggle();
  }

  #onNodeClick() {
    if (this.hasChildren) {
      this.toggle();
    }
    this.select();
  }

  #onSlotChange() {
    if (this.#explicitHasChildren) return;
    const slot = this.shadowRoot?.querySelector('#children slot') as HTMLSlotElement | null;
    if (!slot) return;
    const children = slot.assignedElements();
    this.hasChildren = children.length > 0;
  }

  #updateIcon() {
    const iconEl = this.shadowRoot?.getElementById('icon');
    if (!iconEl) return;
    iconEl.innerHTML = this.icon ?? '';
  }

  /** Toggle expanded state */
  toggle() {
    if (!this.hasChildren) return;
    this.expanded = !this.expanded;
  }

  /** Expand the item */
  expand() {
    if (!this.hasChildren) return;
    this.expanded = true;
  }

  /** Collapse the item */
  collapse() {
    if (!this.hasChildren) return;
    this.expanded = false;
  }

  /** Select this item */
  select() {
    this.current = true;
    this.dispatchEvent(new PfTreeItemSelectEvent());
  }

  /** Deselect this item */
  deselect() {
    this.current = false;
  }

  /** Set tabindex for roving tabindex pattern */
  setTabindex(value: number) {
    const item = this.shadowRoot?.getElementById('item');
    item?.setAttribute('tabindex', String(value));
  }

  /** Focus this item */
  focusItem() {
    const item = this.shadowRoot?.getElementById('item');
    item?.focus();
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-tree-item': PfV6TreeItem;
  }
}
