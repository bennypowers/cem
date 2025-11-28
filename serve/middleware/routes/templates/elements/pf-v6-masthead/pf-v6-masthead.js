import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event for sidebar toggle
 */
export class SidebarToggleEvent extends Event {
  constructor(expanded) {
    super('sidebar-toggle', { bubbles: true });
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
  #toggleButton;
  #toggleListener;

  constructor() {
    super();
    this.#internals.role = 'banner';
  }

  async afterTemplateLoaded() {
    this.#toggleButton = this.shadowRoot?.querySelector('#toggle-button');
    this.#attachEventListeners();
    this.#updateToggleButton();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'sidebar-expanded' && oldValue !== newValue) {
      this.#updateToggleButton();
    }
  }

  #attachEventListeners() {
    if (!this.#toggleButton) return;

    this.#toggleListener = () => {
      const isExpanded = this.hasAttribute('sidebar-expanded');
      this.dispatchEvent(new SidebarToggleEvent(!isExpanded));
    };

    this.#toggleButton.addEventListener('click', this.#toggleListener);
  }

  #updateToggleButton() {
    if (!this.#toggleButton) return;

    const isExpanded = this.hasAttribute('sidebar-expanded');
    this.#toggleButton.setAttribute('aria-expanded', String(isExpanded));
  }

  disconnectedCallback() {
    super.disconnectedCallback?.();

    if (this.#toggleListener && this.#toggleButton) {
      this.#toggleButton.removeEventListener('click', this.#toggleListener);
      this.#toggleListener = null;
    }
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
