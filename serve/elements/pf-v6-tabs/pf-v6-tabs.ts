import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { state } from 'lit/decorators/state.js';

import styles from './pf-v6-tabs.css' with { type: 'css' };

import '../pf-v6-tab/pf-v6-tab.js';

/**
 * Custom event fired when tab selection changes
 */
export class PfTabsChangeEvent extends Event {
  selectedIndex: number;
  constructor(selectedIndex: number) {
    super('change', { bubbles: true });
    this.selectedIndex = selectedIndex;
  }
}

/**
 * PatternFly v6 Tabs
 *
 * A tabbed content container. Add `pf-v6-tab` children with `title` attributes
 * to define tabs and their panel content.
 *
 * @slot - Default slot for `pf-v6-tab` children
 *
 * @fires change - Fired when the selected tab changes
 *
 * @csspart tabs - The tab button container
 * @csspart panels - The tab panel container
 */
@customElement('pf-v6-tabs')
export class PfV6Tabs extends LitElement {
  static styles = styles;

  @property({ type: Number, reflect: true })
  accessor selected = 0;

  @state()
  accessor _tabs: Element[] = [];

  #mutationObserver?: MutationObserver;

  get selectedIndex(): number {
    return this.selected;
  }

  set selectedIndex(index: number) {
    this.selected = Math.max(0, Math.min(index, this._tabs.length - 1));
  }

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('pf-v6-tab-title-changed', this.#onTabsChanged);
    this.addEventListener('pf-v6-tab-connected', this.#onTabsChanged);
    this.addEventListener('pf-v6-tab-disconnected', this.#onTabsChanged);
    this.#mutationObserver = new MutationObserver(this.#onTabsChanged);
    this.#mutationObserver.observe(this, { childList: true });
    this.#syncTabs();
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.#mutationObserver?.disconnect();
    this.removeEventListener('pf-v6-tab-title-changed', this.#onTabsChanged);
    this.removeEventListener('pf-v6-tab-connected', this.#onTabsChanged);
    this.removeEventListener('pf-v6-tab-disconnected', this.#onTabsChanged);
  }

  #onTabsChanged = () => {
    this.#syncTabs();
  };

  #syncTabs() {
    this._tabs = Array.from(this.querySelectorAll('pf-v6-tab'));
    this.requestUpdate();
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('selected') || changed.has('_tabs')) {
      // Assign slot names to tab children
      this._tabs.forEach((tab, i) => {
        tab.setAttribute('slot', `panel-${i}`);
      });
      requestAnimationFrame(() => this.#updateAccentLine());
    }
  }

  #handleTabClick(index: number) {
    const oldIndex = this.selected;
    this.selected = index;
    if (oldIndex !== index) {
      this.dispatchEvent(new PfTabsChangeEvent(index));
    }
  }

  #handleKeyDown(e: KeyboardEvent) {
    const tabsEl = this.shadowRoot?.getElementById('tabs');
    if (!tabsEl) return;
    const buttons = Array.from(tabsEl.querySelectorAll('.tab')) as HTMLButtonElement[];
    const currentIndex = buttons.findIndex(btn => btn === this.shadowRoot?.activeElement);
    if (currentIndex === -1) return;

    let nextIndex = currentIndex;

    switch (e.key) {
      case 'ArrowLeft':
        e.preventDefault();
        nextIndex = currentIndex > 0 ? currentIndex - 1 : buttons.length - 1;
        break;
      case 'ArrowRight':
        e.preventDefault();
        nextIndex = currentIndex < buttons.length - 1 ? currentIndex + 1 : 0;
        break;
      case 'Home':
        e.preventDefault();
        nextIndex = 0;
        break;
      case 'End':
        e.preventDefault();
        nextIndex = buttons.length - 1;
        break;
      default:
        return;
    }

    this.#handleTabClick(nextIndex);
    buttons[nextIndex].focus();
  }

  #updateAccentLine() {
    const tabsEl = this.shadowRoot?.getElementById('tabs');
    if (!tabsEl) return;
    const buttons = Array.from(tabsEl.querySelectorAll('.tab'));
    const activeButton = buttons[this.selected] as HTMLElement | undefined;
    if (!activeButton) return;

    const containerRect = tabsEl.getBoundingClientRect();
    const buttonRect = activeButton.getBoundingClientRect();
    const inset = 8; // var(--pf-t--global--spacer--sm)
    const start = (buttonRect.left - containerRect.left) + inset;
    const length = buttonRect.width - (inset * 2);

    tabsEl.style.setProperty('--_pf-v6-tabs--link-accent--start', `${start}px`);
    tabsEl.style.setProperty('--_pf-v6-tabs--link-accent--length', `${length}px`);
  }

  render() {
    return html`
      <div id="tabs-container">
        <div id="tabs"
             role="tablist"
             aria-label="Tabs"
             part="tabs"
             @keydown=${this.#handleKeyDown}>
          ${this._tabs.map((tab, index) => html`
            <button class="tab"
                    type="button"
                    role="tab"
                    id="tab-${index}"
                    aria-controls="panel-${index}"
                    aria-selected="${index === this.selected ? 'true' : 'false'}"
                    tabindex="${index === this.selected ? 0 : -1}"
                    @click=${() => this.#handleTabClick(index)}>
              ${(tab as HTMLElement).title || `Tab ${index + 1}`}
            </button>
          `)}
        </div>
        <div id="panels"
             part="panels">
          ${this._tabs.map((_, index) => html`
            <div class="panel"
                 id="panel-${index}"
                 role="tabpanel"
                 aria-labelledby="tab-${index}"
                 ?hidden=${index !== this.selected}>
              <slot name="panel-${index}"></slot>
            </div>
          `)}
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-tabs': PfV6Tabs;
  }
}
