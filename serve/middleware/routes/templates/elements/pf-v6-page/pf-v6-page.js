import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Page Layout
 *
 * @attr {boolean} sidebar-collapsed - Whether the sidebar is collapsed
 *
 * @slot skip-to-content - Skip to content link
 * @slot masthead - Page header (pf-v6-masthead)
 * @slot sidebar - Page sidebar (pf-v6-page-sidebar)
 * @slot main - Main content area (pf-v6-page-main)
 *
 * @listens sidebar-toggle - Responds to toggle events from masthead
 * @customElement pf-v6-page
 */
export class PfV6Page extends CemElement {
  static is = 'pf-v6-page';

  static observedAttributes = ['sidebar-collapsed'];

  /** Are we a wide layout, or a narrow layout? */
  static match = window.matchMedia('(min-width: 75rem)');

  // Private field for click outside handler
  #clickOutsideHandler = (event) => {
    if (!PfV6Page.match.matches && !this.sidebarCollapsed) {
      const sidebar = this.querySelector('pf-v6-page-sidebar');
      // a little cheat, to avoid race when toggling open via button
      // although normally, reaching into shadow dom is forbidden
      // we do it here because page and masthead are meant to work together
      const mastheadToggle = this.querySelector('pf-v6-masthead')?.shadowRoot?.getElementById('toggle-button');

      if (!event.composedPath().some(node => node === sidebar || node === mastheadToggle)) {
        this.sidebarCollapsed = true;
      }
    }
  };

  get sidebarCollapsed() {
    return this.hasAttribute('sidebar-collapsed');
  }

  set sidebarCollapsed(value) {
    this.toggleAttribute('sidebar-collapsed', !!value);
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue) return;

    if (name === 'sidebar-collapsed') {
      this.#updateSidebarState();
    }
  }

  afterTemplateLoaded() {
    this.#attachEventListeners();
    this.#updateSidebarState();
  }

  #attachEventListeners() {
    // Listen for sidebar toggle events from masthead
    this.addEventListener('sidebar-toggle', (event) => {
      this.sidebarCollapsed = !event.expanded;
    });

    // Listen for clicks on document body to close sidebar when clicking outside
    // This is attached to body to avoid accessibility issues with clickable page elements
    document.body.addEventListener('click', this.#clickOutsideHandler);
  }

  #updateSidebarState() {
    const isCollapsed = this.hasAttribute('sidebar-collapsed');

    // Update sidebar element
    const sidebar = this.querySelector('pf-v6-page-sidebar');
    sidebar?.toggleAttribute('collapsed', !!isCollapsed);
    sidebar?.toggleAttribute('expanded', !isCollapsed);

    // Update masthead element
    const masthead = this.querySelector('pf-v6-masthead');
    masthead?.toggleAttribute('sidebar-expanded', !isCollapsed);
  }

  disconnectedCallback() {
    super.disconnectedCallback?.();

    // Clean up click outside handler
    if (this.#clickOutsideHandler) {
      document.body.removeEventListener('click', this.#clickOutsideHandler);
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
