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

  async afterTemplateLoaded() {
    this.#attachEventListeners();
    this.#updateSidebarState();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'sidebar-collapsed' && oldValue !== newValue) {
      this.#updateSidebarState();
    }
  }

  #attachEventListeners() {
    // Listen for sidebar toggle events from masthead
    this.addEventListener('sidebar-toggle', (event) => {
      this.sidebarCollapsed = !event.expanded;
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

  get sidebarCollapsed() {
    return this.hasAttribute('sidebar-collapsed');
  }

  set sidebarCollapsed(value) {
    this.toggleAttribute('sidebar-collapsed', !!value);
  }

  static {
    customElements.define(this.is, this);
  }
}
