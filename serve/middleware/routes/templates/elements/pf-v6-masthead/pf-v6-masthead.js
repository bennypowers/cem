import { CemElement } from '/__cem/cem-element.js';

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
 * @customElement pf-v6-masthead
 */
class PfV6Masthead extends CemElement {
  static observedAttributes = ['sidebar-expanded'];
  static is = 'pf-v6-masthead';

  #internals = this.attachInternals();

  constructor() {
    super();
    this.#internals.role = 'banner';
  }

  async afterTemplateLoaded() {
    this.#attachEventListeners();
    this.#updateToggleButton();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'sidebar-expanded' && oldValue !== newValue) {
      this.#updateToggleButton();
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
    this.toggleAttribute('sidebar-expanded', !!value);
  }

  static {
    customElements.define(this.is, this);
  }
}

