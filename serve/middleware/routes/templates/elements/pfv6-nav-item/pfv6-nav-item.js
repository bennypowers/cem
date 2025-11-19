import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 Navigation Item
 *
 * @attr {boolean} expanded - Whether nested group is expanded
 * @attr {boolean} current - Whether this item is current
 *
 * @slot - Default slot for nav-link and optional nav-group
 */
class Pfv6NavItem extends HTMLElement {
  static get observedAttributes() {
    return ['expanded', 'current'];
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    if (!this.shadowRoot.hasChildNodes()) {
      await this.#populateShadowRoot();
    }
    this.setAttribute('role', 'listitem');

    // Listen for toggle events from child nav-link with expandable group
    this.addEventListener('pf-nav-toggle', this.#handleToggle);

    // Initialize state on connection
    this.#updateExpandedState();
  }

  disconnectedCallback() {
    this.removeEventListener('pf-nav-toggle', this.#handleToggle);
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'expanded' && oldValue !== newValue) {
      this.#updateExpandedState();
    }
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-nav-item');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-nav-item template:', error);
      this.shadowRoot.innerHTML = '<slot></slot>';
    }
  }

  #handleToggle = (event) => {
    event.stopPropagation();
    this.expanded = !this.expanded;
  };

  get expanded() {
    return this.hasAttribute('expanded');
  }

  set expanded(value) {
    const isExpanded = Boolean(value);
    if (isExpanded) {
      this.setAttribute('expanded', '');
    } else {
      this.removeAttribute('expanded');
    }
  }

  #updateExpandedState() {
    const isExpanded = this.hasAttribute('expanded');

    // Update child nav-group
    const navGroup = this.querySelector('pfv6-nav-group');
    if (navGroup) {
      if (isExpanded) {
        navGroup.removeAttribute('hidden');
      } else {
        navGroup.setAttribute('hidden', '');
      }
    }

    // Update child nav-link aria-expanded
    // (CSS will handle toggle icon rotation via [aria-expanded] selector)
    const navLink = this.querySelector('pfv6-nav-link[expandable]');
    if (navLink) {
      navLink.setAttribute('aria-expanded', String(isExpanded));
    }
  }
}

customElements.define('pfv6-nav-item', Pfv6NavItem);
