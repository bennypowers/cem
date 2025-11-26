import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Navigation Item
 *
 * Container for navigation links with optional expandable groups.
 * The child pf-v6-nav-link handles the current state independently.
 *
 * @attr {boolean} expanded - Whether nested group is expanded
 *
 * @slot - Default slot for nav-link and optional nav-group
 * @customElement pf-v6-nav-item
 */
class PfV6NavItem extends CemElement {
  static observedAttributes = ['expanded'];
  static is = 'pf-v6-nav-item';

  get expanded() { return this.hasAttribute('expanded'); }
  set expanded(value) { this.toggleAttribute('expanded', !!value); }

  afterTemplateLoaded() {
    this.setAttribute('role', 'listitem');
    this.addEventListener('pf-nav-toggle', this.#handleToggle);
    this.#updateExpandedState();
  }

  disconnectedCallback() {
    super.disconnectedCallback?.();
    this.removeEventListener('pf-nav-toggle', this.#handleToggle);
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'expanded' && oldValue !== newValue) {
      this.#updateExpandedState();
    }
  }

  #handleToggle = (event) => {
    event.stopPropagation();
    this.expanded = !this.expanded;
  };

  #updateExpandedState() {
    const isExpanded = this.hasAttribute('expanded');

    // Update child nav-group
    this.querySelector('pf-v6-nav-group')
      ?.toggleAttribute('hidden', isExpanded);

    // Update child nav-link expanded state
    // The nav-link will forward this to aria-expanded on its internal element
    this.querySelector('pf-v6-nav-link[expandable]')
      ?.toggleAttribute('expanded', isExpanded);
  }

  static {
    customElements.define(this.is, this);
  }
}
