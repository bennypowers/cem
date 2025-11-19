import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * Custom event for sidebar toggle
 */
export class SidebarToggleEvent extends Event {
  constructor(expanded) {
    super('sidebar-toggle', { bubbles: true, composed: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Masthead
 *
 * @attr {boolean} sidebar-expanded - Whether the sidebar is currently expanded
 *
 * @slot logo - Logo or brand content
 * @slot toolbar - Toolbar content (actions, menus, etc.)
 *
 * @fires {SidebarToggleEvent} sidebar-toggle - When the hamburger toggle is clicked
 */
class Pfv6Masthead extends HTMLElement {
  static formAssociated = false;

  #internals;

  static get observedAttributes() {
    return ['sidebar-expanded'];
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
    this.#internals = this.attachInternals();
    this.#internals.role = 'banner';
  }

  async connectedCallback() {
    if (!this.shadowRoot.querySelector('.pf-v6-c-masthead')) {
      await this.#populateShadowRoot();
    }
    this.#attachEventListeners();
    this.#updateToggleButton();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'sidebar-expanded' && oldValue !== newValue) {
      this.#updateToggleButton();
    }
  }

  async #populateShadowRoot() {
    try {
      const { html, stylesheet } = await loadComponentTemplate('pfv6-masthead');
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-masthead template:', error);
    }
  }

  #attachEventListeners() {
    const toggleButton = this.shadowRoot?.querySelector('#toggle-button');
    if (!toggleButton) return;

    toggleButton.addEventListener('click', () => {
      const isExpanded = this.hasAttribute('sidebar-expanded');
      this.dispatchEvent(new SidebarToggleEvent(!isExpanded));
    });
  }

  #updateToggleButton() {
    const toggleButton = this.shadowRoot?.querySelector('#toggle-button');
    if (!toggleButton) return;

    const isExpanded = this.hasAttribute('sidebar-expanded');
    toggleButton.setAttribute('aria-expanded', String(isExpanded));
  }

  get sidebarExpanded() {
    return this.hasAttribute('sidebar-expanded');
  }

  set sidebarExpanded(value) {
    if (value) {
      this.setAttribute('sidebar-expanded', '');
    } else {
      this.removeAttribute('sidebar-expanded');
    }
  }
}

customElements.define('pfv6-masthead', Pfv6Masthead);
