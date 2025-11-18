import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

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
 */
class Pfv6NavLink extends HTMLElement {
  static get observedAttributes() {
    return ['href', 'current', 'expandable', 'aria-expanded', 'aria-label'];
  }

  #internals;

  constructor() {
    super();
    this.#internals = this.attachInternals();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('#link')) {
      await this.#populateShadowRoot();
    }
    this.#syncAttributes();
    this.#attachEventListeners();
    this.#markCurrentIfMatches();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-nav-link');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-nav-link template:', error);
      this.shadowRoot.innerHTML = '<a id="link"><slot></slot></a>';
    }
  }

  #syncAttributes() {
    const link = this.shadowRoot?.querySelector('#link');
    if (!link) return;

    // Handle href
    if (this.hasAttribute('href')) {
      link.setAttribute('href', this.getAttribute('href'));
    } else {
      link.removeAttribute('href');
    }

    // Handle aria-label
    if (this.hasAttribute('aria-label')) {
      link.setAttribute('aria-label', this.getAttribute('aria-label'));
    } else {
      link.removeAttribute('aria-label');
    }

    // Handle aria-expanded for expandable links
    if (this.hasAttribute('expandable')) {
      const expanded = this.getAttribute('aria-expanded') === 'true';
      link.setAttribute('aria-expanded', String(expanded));
      link.setAttribute('role', 'button');
    } else {
      link.removeAttribute('aria-expanded');
      link.removeAttribute('role');
    }
  }

  #attachEventListeners() {
    const link = this.shadowRoot?.querySelector('#link');
    if (!link) return;

    link.addEventListener('click', (e) => {
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
        this.setAttribute('aria-current', 'page');
      }
    }
  }

  get current() {
    return this.hasAttribute('current');
  }

  set current(value) {
    if (value) {
      this.setAttribute('current', '');
      this.setAttribute('aria-current', 'page');
    } else {
      this.removeAttribute('current');
      this.removeAttribute('aria-current');
    }
  }
}

customElements.define('pfv6-nav-link', Pfv6NavLink);
