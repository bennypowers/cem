import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation Container
 *
 * @attr {string} aria-label - Accessible label for navigation
 * @attr {boolean} inset - Add horizontal padding
 * @attr {boolean} horizontal - Horizontal layout
 * @attr {boolean} scrollable - Enable scroll buttons
 *
 * @slot - Default slot for nav-list
 * @customElement pf-v6-navigation
 */
class PfV6Navigation extends CemElement {
  static observedAttributes = ['aria-label', 'inset', 'variant', 'horizontal', 'scrollable'];
  static is = 'pf-v6-navigation';

  #nav;
  #navList;
  #scrollBackButton;
  #scrollForwardButton;
  #resizeObserver;

  get variant() {
    return this.getAttribute('variant') ?? '';
  }

  set variant(value) {
    this.setAttribute('variant', value);
  }

  get horizontal() {
    return this.hasAttribute('horizontal');
  }

  set horizontal(value) {
    this.toggleAttribute('horizontal', !!value);
  }

  get scrollable() {
    return this.hasAttribute('scrollable');
  }

  set scrollable(value) {
    this.toggleAttribute('scrollable', !!value);
  }

  async afterTemplateLoaded() {
    this.#nav = this.shadowRoot.querySelector('nav');
    this.#scrollBackButton = this.shadowRoot.querySelector('#scroll-back');
    this.#scrollForwardButton = this.shadowRoot.querySelector('#scroll-forward');

    this.#syncAttributes();
    this.#setupScrollButtons();
    this.#observeNavList();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#nav) return;

    if (this.hasAttribute('aria-label')) {
      this.#nav.setAttribute('aria-label', this.getAttribute('aria-label'));
    }
  }

  #observeNavList() {
    // Wait for slotted content
    const slot = this.#nav?.querySelector('slot');
    if (!slot) return;

    const initNavList = () => {
      const elements = slot.assignedElements();
      this.#navList = elements.find(el => el.tagName === 'PF-V6-NAV-LIST');
      if (this.#navList && this.horizontal) {
        this.#setupScrollObserver();
      }
    };

    slot.addEventListener('slotchange', initNavList);

    // Also check immediately in case elements are already slotted
    initNavList();
  }

  #setupScrollButtons() {
    if (!this.#scrollBackButton || !this.#scrollForwardButton) return;

    this.#scrollBackButton.addEventListener('click', () => this.#scrollBack());
    this.#scrollForwardButton.addEventListener('click', () => this.#scrollForward());
  }

  #setupScrollObserver() {
    if (!this.#navList) return;

    // Update button states on scroll
    this.#nav.addEventListener('scroll', () => {
      this.#handleScrollButtons();
    });

    // Use single ResizeObserver for all size changes
    // Container queries can't detect overflow state, so we need ResizeObserver
    this.#resizeObserver = new ResizeObserver(() => {
      this.#handleScrollButtons();
    });

    // Observe the nav container (scrollable element)
    this.#resizeObserver.observe(this.#nav);

    // Observe the nav-list (content that may overflow)
    this.#resizeObserver.observe(this.#navList);

    // Initial check after layout
    requestAnimationFrame(() => {
      this.#handleScrollButtons();
    });
  }

  #handleScrollButtons() {
    if (!this.#navList || !this.horizontal) return;

    // Get scroll dimensions from the nav element (the scrollable container)
    const { scrollLeft, scrollWidth, clientWidth } = this.#nav;

    // Check if scrollable - nav content overflows nav container
    const isOverflowing = scrollWidth > clientWidth;

    // Check scroll position
    const scrollViewAtStart = scrollLeft <= 1;
    const scrollViewAtEnd = scrollLeft + clientWidth >= scrollWidth - 1;

    // Update scrollable attribute (shows/hides buttons)
    this.scrollable = isOverflowing && (!scrollViewAtStart || !scrollViewAtEnd);

    // Update button disabled states
    if (this.#scrollBackButton) {
      this.#scrollBackButton.toggleAttribute('disabled', scrollViewAtStart);
    }

    if (this.#scrollForwardButton) {
      this.#scrollForwardButton.toggleAttribute('disabled', scrollViewAtEnd);
    }
  }

  #scrollBack() {
    if (!this.#nav) return;
    this.#nav.scrollBy({ left: -200, behavior: 'smooth' });
  }

  #scrollForward() {
    if (!this.#nav) return;
    this.#nav.scrollBy({ left: 200, behavior: 'smooth' });
  }

  disconnectedCallback() {
    if (this.#resizeObserver) {
      this.#resizeObserver.disconnect();
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
