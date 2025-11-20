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

    console.log('[pf-v6-navigation] afterTemplateLoaded - variant:', this.variant, 'horizontal:', this.horizontal, 'parent:', this.parentElement?.tagName);

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
    if (!slot) {
      console.log('[pf-v6-navigation] No slot found');
      return;
    }

    const initNavList = () => {
      const elements = slot.assignedElements();
      console.log('[pf-v6-navigation] Slot assigned elements:', elements);
      this.#navList = elements.find(el => el.tagName === 'PF-V6-NAV-LIST');
      console.log('[pf-v6-navigation] Nav list found:', this.#navList, 'horizontal:', this.horizontal);
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
    if (!this.#navList) {
      console.log('[pf-v6-navigation] setupScrollObserver: no navList');
      return;
    }

    console.log('[pf-v6-navigation] Setting up scroll observer on navList:', this.#navList);

    // Update button states on scroll - listen to nav element
    this.#nav.addEventListener('scroll', () => {
      console.log('[pf-v6-navigation] Scroll event fired');
      this.#handleScrollButtons();
    });

    // Observe size changes to detect when scrollable state should change
    this.#resizeObserver = new ResizeObserver(() => {
      console.log('[pf-v6-navigation] ResizeObserver fired');
      this.#handleScrollButtons();
    });

    this.#resizeObserver.observe(this.#navList);

    // Also observe the host element to catch window resizes
    this.#resizeObserver.observe(this);

    // Also observe each nav-item as they render
    const navItems = this.#navList.querySelectorAll('pf-v6-nav-item');
    navItems.forEach(item => {
      this.#resizeObserver.observe(item);
    });

    // Watch for new nav-items being added
    const itemObserver = new MutationObserver(() => {
      console.log('[pf-v6-navigation] Nav items mutated, re-checking');
      this.#handleScrollButtons();

      // Observe any new items
      const newItems = this.#navList.querySelectorAll('pf-v6-nav-item');
      newItems.forEach(item => {
        this.#resizeObserver.observe(item);
      });
    });

    itemObserver.observe(this.#navList, { childList: true, subtree: true });

    // Initial check - wait for layout
    requestAnimationFrame(() => {
      requestAnimationFrame(() => {
        console.log('[pf-v6-navigation] Initial handleScrollButtons check (after 2 RAF)');
        this.#handleScrollButtons();
      });
    });
  }

  #handleScrollButtons() {
    if (!this.#navList || !this.horizontal) {
      console.log('[pf-v6-navigation] handleScrollButtons early exit - navList:', this.#navList, 'horizontal:', this.horizontal);
      return;
    }

    // Get scroll dimensions from the nav element (the scrollable container)
    const { scrollLeft, scrollWidth, clientWidth } = this.#nav;

    console.log('[pf-v6-navigation] handleScrollButtons dimensions:',
      `scrollLeft=${scrollLeft}`,
      `scrollWidth=${scrollWidth}`,
      `clientWidth=${clientWidth}`,
      `nav.offsetWidth=${this.#nav.offsetWidth}`,
      `isOverflowing=${scrollWidth > clientWidth}`
    );

    // Check if scrollable - nav content overflows nav container
    const isOverflowing = scrollWidth > clientWidth;

    // Check scroll position
    const scrollViewAtStart = scrollLeft <= 1;
    const scrollViewAtEnd = scrollLeft + clientWidth >= scrollWidth - 1;

    console.log('[pf-v6-navigation] Scroll state:',
      `isOverflowing=${isOverflowing}`,
      `scrollViewAtStart=${scrollViewAtStart}`,
      `scrollViewAtEnd=${scrollViewAtEnd}`,
      `shouldBeScrollable=${isOverflowing && (!scrollViewAtStart || !scrollViewAtEnd)}`
    );

    // Update scrollable attribute (shows/hides buttons)
    this.scrollable = isOverflowing && (!scrollViewAtStart || !scrollViewAtEnd);

    console.log('[pf-v6-navigation] Set scrollable to:', this.scrollable);

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
