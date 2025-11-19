import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Page Layout
 *
 * @attr {boolean} sidebar-collapsed - Whether the sidebar is collapsed
 *
 * @slot skip-to-content - Skip to content link
 * @slot masthead - Page header (pfv6-masthead)
 * @slot sidebar - Page sidebar (pfv6-page-sidebar)
 * @slot main - Main content area (pfv6-page-main)
 *
 * @listens sidebar-toggle - Responds to toggle events from masthead
 */
class Pfv6Page extends CemElement {
  static observedAttributes = ['sidebar-collapsed'];

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
    const pageDiv = this.shadowRoot?.querySelector('.pf-v6-c-page');
    if (!pageDiv) return;

    const isCollapsed = this.hasAttribute('sidebar-collapsed');

    // Update page container class
    if (isCollapsed) {
      pageDiv.classList.add('pf-m-sidebar-collapsed');
    } else {
      pageDiv.classList.remove('pf-m-sidebar-collapsed');
    }

    // Update sidebar element
    const sidebar = this.querySelector('pfv6-page-sidebar');
    if (sidebar) {
      if (isCollapsed) {
        sidebar.setAttribute('collapsed', '');
        sidebar.removeAttribute('expanded');
      } else {
        sidebar.setAttribute('expanded', '');
        sidebar.removeAttribute('collapsed');
      }
    }

    // Update masthead element
    const masthead = this.querySelector('pfv6-masthead');
    if (masthead) {
      if (isCollapsed) {
        masthead.removeAttribute('sidebar-expanded');
      } else {
        masthead.setAttribute('sidebar-expanded', '');
      }
    }
  }

  get sidebarCollapsed() {
    return this.hasAttribute('sidebar-collapsed');
  }

  set sidebarCollapsed(value) {
    if (value) {
      this.setAttribute('sidebar-collapsed', '');
    } else {
      this.removeAttribute('sidebar-collapsed');
    }
  }
}

customElements.define('pfv6-page', Pfv6Page);
