import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event for navigation toggle
 */
export class PfNavToggleEvent extends Event {
  constructor(expanded) {
    super('pf-nav-toggle', { bubbles: true, composed: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Navigation Link
 *
 * @attr {string} href - Link URL
 * @attr {string} label - Accessible label for the link
 * @attr {boolean} current - Whether this link is current/active
 * @attr {boolean} expandable - Whether this link toggles a nav-group
 * @attr {boolean} expanded - Whether the expandable group is expanded (for expandable links)
 *
 * @slot - Default slot for link text
 * @slot icon-start - Icon before text
 * @slot icon-end - Icon after text
 * @slot toggle - Toggle icon for expandable items
 *
 * @fires {PfNavToggleEvent} pf-nav-toggle - When expandable link is clicked
 * @customElement pf-v6-nav-link
 */
class PfV6NavLink extends CemElement {
  static is = 'pf-v6-nav-link';

  static observedAttributes = [
    'href',
    'label',
    'current',
    'expandable',
    'expanded',
  ];

  get current() { return this.hasAttribute('current'); }
  set current(value) { this.toggleAttribute('current', !!value); }

  #link;

  afterTemplateLoaded() {
    this.#link = this.shadowRoot.querySelector('#link');
    this.#syncAttributes();
    this.#attachEventListeners();
    this.#markCurrentIfMatches();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    // Skip if value hasn't actually changed
    if (oldValue === newValue) return;

    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#link) return;

    // Only sync runtime-changeable attributes
    // SSR handles initial href, aria-label, aria-current, aria-expanded

    // Handle href changes at runtime (for <a> elements)
    if (this.#link.tagName === 'A') {
      if (this.hasAttribute('href')) {
        this.#link.setAttribute('href', this.getAttribute('href'));
      } else {
        this.#link.removeAttribute('href');
      }
    }

    // Handle label -> aria-label
    if (this.hasAttribute('label')) {
      this.#link.setAttribute('aria-label', this.getAttribute('label'));
    } else {
      this.#link.removeAttribute('aria-label');
    }

    // Handle expanded -> aria-expanded for expandable buttons
    if (this.hasAttribute('expandable') && this.#link.tagName === 'BUTTON') {
      const expanded = this.hasAttribute('expanded');
      this.#link.setAttribute('aria-expanded', String(expanded));
    }

    // Handle aria-current for current links
    if (this.hasAttribute('current')) {
      this.#link.setAttribute('aria-current', 'page');
    } else {
      this.#link.removeAttribute('aria-current');
    }
  }

  #attachEventListeners() {
    if (!this.#link) return;

    this.#link.addEventListener('click', (e) => {
      if (this.hasAttribute('expandable')) {
        e.preventDefault();
        const currentExpanded = this.hasAttribute('expanded');
        this.dispatchEvent(new PfNavToggleEvent(!currentExpanded));
      }
    });
  }

  #markCurrentIfMatches() {
    if (this.hasAttribute('href')) {
      const href = this.getAttribute('href');
      const currentPath = window.location.pathname;
      if (href === currentPath) {
        this.setAttribute('current', '');
      }
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
