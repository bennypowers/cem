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
 * @attr {boolean} current - Whether this link is current/active
 * @attr {boolean} expandable - Whether this link toggles a nav-group
 * @attr {string} aria-expanded - Expanded state for expandable links
 * @attr {string} aria-label - Accessible label
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
  static observedAttributes = ['href', 'current', 'expandable', 'aria-expanded', 'aria-label'];
  static is = 'pf-v6-nav-link';

  #internals = this.attachInternals();
  #link;

  async afterTemplateLoaded() {
    this.#link = this.shadowRoot.querySelector('#link');
    this.#syncAttributes();
    this.#attachEventListeners();
    this.#markCurrentIfMatches();
  }

  attributeChangedCallback() {
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

    // Handle aria-label changes at runtime
    if (this.hasAttribute('aria-label')) {
      this.#link.setAttribute('aria-label', this.getAttribute('aria-label'));
    } else {
      this.#link.removeAttribute('aria-label');
    }

    // Handle aria-expanded for expandable buttons
    if (this.hasAttribute('expandable') && this.#link.tagName === 'BUTTON') {
      const expanded = this.getAttribute('aria-expanded') === 'true';
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
        const currentExpanded = this.getAttribute('aria-expanded') === 'true';
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

  get current() {
    return this.hasAttribute('current');
  }

  set current(value) {
    this.toggleAttribute('current', !!value);
  }

  static {
    customElements.define(this.is, this);
  }
}

