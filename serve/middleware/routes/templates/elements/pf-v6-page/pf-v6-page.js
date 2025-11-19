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
class PfV6Page extends CemElement {
  static observedAttributes = ['sidebar-collapsed'];
  static is = 'pf-v6-page';
  static match = window.matchMedia('(min-width: 75rem)');

  /** Are we a wide layout, or a narrow layout? */
  #wide = false;

  get sidebarCollapsed() {
    return this.hasAttribute('sidebar-collapsed');
  }

  set sidebarCollapsed(value) {
    this.toggleAttribute('sidebar-collapsed', !!value);
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'sidebar-collapsed' && oldValue !== newValue) {
      this.#updateSidebarState();
    }
  }

  async afterTemplateLoaded() {
    this.#attachEventListeners();
    this.#updateSidebarState();
  }

  #attachEventListeners() {
    // Listen for sidebar toggle events from masthead
    this.addEventListener('sidebar-toggle', (event) => {
      this.sidebarCollapsed = !event.expanded;
    });
    PfV6Page.match.addEventListener('change', (event) => {
      if (event.matches) {
        this.#wide = true;
      }
    });
    this.addEventListener('click', (event) => {
      if (!this.#wide &&
          !this.sidebarCollapsed &&
          !event
            .composedPath()
            .some(node =>
              node === this.querySelector('pf-v6-page-sidebar') ||
              // a little cheat, to avoid race when toggling open via button
              node === this.querySelector('pf-v6-masthead')
              .shadowRoot.getElementById('toggle-button'))) {
        this.sidebarCollapsed = true;
      }
    });
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

  static {
    customElements.define(this.is, this);
  }
}
