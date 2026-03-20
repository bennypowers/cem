import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-navigation.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

/**
 * PatternFly v6 Navigation Container
 *
 * @slot - Default slot for nav-list
 *
 * @attr {string} variant - Navigation variant
 * @attr {boolean} inset - Add horizontal padding
 * @attr {boolean} horizontal - Horizontal layout
 * @attr {boolean} scrollable - Enable scroll buttons
 */
@customElement('pf-v6-navigation')
export class PfV6Navigation extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor variant?: string;

  @property({ type: Boolean, reflect: true })
  accessor horizontal = false;

  @property({ type: Boolean, reflect: true })
  accessor scrollable = false;

  @property({ type: Boolean, reflect: true })
  accessor inset = false;

  #nav?: HTMLElement;
  #navList?: Element;
  #resizeObserver?: ResizeObserver;

  render() {
    return html`
      <pf-v6-button id="scroll-back"
                    class="scroll-button"
                    variant="plain"
                    aria-label="Scroll back"
                    @click=${this.#scrollBack}>
        <svg viewBox="0 0 256 512" role="presentation">
          <path d="M31.7 239l136-136c9.4-9.4 24.6-9.4 33.9 0l22.6 22.6c9.4 9.4 9.4 24.6 0 33.9L127.9 256l96.4 96.4c9.4 9.4 9.4 24.6 0 33.9L201.7 409c-9.4 9.4-24.6 9.4-33.9 0l-136-136c-9.5-9.4-9.5-24.6-.1-34z"></path>
        </svg>
      </pf-v6-button>
      <nav part="nav"
           aria-label=${this.getAttribute('aria-label') || ''}
           @scroll=${this.#handleScrollButtons}>
        <slot @slotchange=${this.#onSlotChange}></slot>
      </nav>
      <pf-v6-button id="scroll-forward"
                    class="scroll-button"
                    variant="plain"
                    aria-label="Scroll forward"
                    @click=${this.#scrollForward}>
        <svg viewBox="0 0 256 512" role="presentation">
          <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
        </svg>
      </pf-v6-button>
    `;
  }

  firstUpdated() {
    this.#nav = this.shadowRoot?.querySelector('nav') as HTMLElement;
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('horizontal')) {
      if (this.horizontal && this.#navList) {
        this.#setupScrollObserver();
      } else {
        this.#teardownScrollObserver();
      }
    }
  }

  #onSlotChange(e: Event) {
    const slot = e.target as HTMLSlotElement;
    const elements = slot.assignedElements();
    this.#navList = elements.find(el => el.tagName === 'PF-V6-NAV-LIST');
    if (this.#navList && this.horizontal) {
      this.#setupScrollObserver();
    }
  }

  #setupScrollObserver() {
    if (!this.#navList || !this.#nav) return;
    if (this.#resizeObserver) return;

    this.#resizeObserver = new ResizeObserver(() => {
      this.#handleScrollButtons();
    });

    this.#resizeObserver.observe(this.#nav);
    this.#resizeObserver.observe(this.#navList);

    requestAnimationFrame(() => {
      this.#handleScrollButtons();
    });
  }

  #handleScrollButtons = () => {
    if (!this.#navList || !this.horizontal || !this.#nav) return;

    const { scrollLeft, scrollWidth, clientWidth } = this.#nav;
    const isOverflowing = scrollWidth > clientWidth;
    const scrollViewAtStart = scrollLeft <= 1;
    const scrollViewAtEnd = scrollLeft + clientWidth >= scrollWidth - 1;

    this.scrollable = isOverflowing && (!scrollViewAtStart || !scrollViewAtEnd);

    const scrollBack = this.shadowRoot?.getElementById('scroll-back');
    const scrollForward = this.shadowRoot?.getElementById('scroll-forward');
    scrollBack?.toggleAttribute('disabled', scrollViewAtStart);
    scrollForward?.toggleAttribute('disabled', scrollViewAtEnd);
  };

  #scrollBack() {
    this.#nav?.scrollBy({ left: -200, behavior: 'smooth' });
  }

  #scrollForward() {
    this.#nav?.scrollBy({ left: 200, behavior: 'smooth' });
  }

  #teardownScrollObserver() {
    this.#resizeObserver?.disconnect();
    this.#resizeObserver = undefined;
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.#teardownScrollObserver();
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-navigation': PfV6Navigation;
  }
}
