import { CemElement } from '/__cem/cem-element.js';
import '/__cem/elements/pf-v6-button/pf-v6-button.js';
import '/__cem/elements/pf-v6-menu/pf-v6-menu.js';
import '/__cem/elements/pf-v6-menu-item/pf-v6-menu-item.js';

/**
 * PatternFly v6 Dropdown component
 *
 * @attr {boolean} expanded - Open/closed state
 * @attr {boolean} disabled - Disabled state
 * @attr {string} label - Accessible label for the menu (sets aria-label on menu via ElementInternals)
 *
 * @slot toggle-text - Text content for toggle button
 * @slot - Menu items (pf-v6-menu-item elements)
 *
 * @fires {Event} expand - Fires when dropdown opens
 * @fires {Event} collapse - Fires when dropdown closes
 *
 * @customElement pf-v6-dropdown
 */
class PfV6Dropdown extends CemElement {
  static is = 'pf-v6-dropdown';

  static shadowRootOptions = { mode: 'open' };

  static observedAttributes = [
    'expanded',
    'disabled',
    'label',
  ];

  #toggle;
  #menu;
  #menuContainer;

  get expanded() {
    return this.hasAttribute('expanded');
  }

  set expanded(value) {
    this.toggleAttribute('expanded', !!value);
  }

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    this.toggleAttribute('disabled', !!value);
  }

  get label() {
    return this.getAttribute('label') || '';
  }

  set label(value) {
    if (value) {
      this.setAttribute('label', value);
    } else {
      this.removeAttribute('label');
    }
  }

  async afterTemplateLoaded() {
    this.#toggle = this.shadowRoot.getElementById('toggle');
    this.#menu = this.shadowRoot.getElementById('menu');
    this.#menuContainer = this.shadowRoot.getElementById('menu-container');

    if (!this.#toggle || !this.#menu || !this.#menuContainer) return;

    // Set up toggle button
    this.#toggle.addEventListener('click', this.#handleToggleClick);

    // Set up keyboard handlers
    this.addEventListener('keydown', this.#handleKeydown);

    // Set up menu select handler
    this.#menu.addEventListener('select', this.#handleMenuSelect);

    // Initialize ARIA attributes
    this.#updateToggleAria();

    // Pass label to menu
    if (this.label) {
      this.#menu.setAttribute('label', this.label);
    }
  }

  #handleToggleClick = (event) => {
    if (this.disabled) {
      event.preventDefault();
      return;
    }

    this.toggle();
  };

  #handleKeydown = (event) => {
    // Handle Escape to close
    if (event.key === 'Escape' && this.expanded) {
      event.preventDefault();
      this.collapse();
      this.#toggle?.focus();
    }
  };

  #handleMenuSelect = (event) => {
    // Just let the event bubble naturally
    // Event already has value and checked properties
  };

  #handleOutsideClick = (event) => {
    // Use composedPath for shadow DOM
    if (!event.composedPath().includes(this)) {
      this.collapse();
    }
  };

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.shadowRoot?.firstChild || oldValue === newValue) {
      return;
    }

    switch (name) {
      case 'expanded':
        this.#updateExpanded();
        break;
      case 'disabled':
        this.#updateDisabled();
        break;
      case 'label':
        this.#updateLabel();
        break;
    }
  }

  #updateExpanded() {
    this.#updateToggleAria();

    if (this.expanded) {
      // Show menu
      if (this.#menuContainer) {
        this.#menuContainer.style.display = 'block';
      }

      // Focus first menu item after DOM update
      requestAnimationFrame(() => {
        this.#menu?.focusFirstItem();
      });

      // Register click-outside handler
      document.addEventListener('click', this.#handleOutsideClick);

      // Dispatch expand event
      this.dispatchEvent(new Event('expand', { bubbles: true }));
    } else {
      // Hide menu
      if (this.#menuContainer) {
        this.#menuContainer.style.display = 'none';
      }

      // Remove click-outside handler
      document.removeEventListener('click', this.#handleOutsideClick);

      // Dispatch collapse event
      this.dispatchEvent(new Event('collapse', { bubbles: true }));
    }
  }

  #updateDisabled() {
    if (this.#toggle) {
      this.#toggle.disabled = this.disabled;
    }
  }

  #updateLabel() {
    if (this.#menu) {
      if (this.label) {
        this.#menu.setAttribute('label', this.label);
      } else {
        this.#menu.removeAttribute('label');
      }
    }
  }

  #updateToggleAria() {
    if (this.#toggle) {
      this.#toggle.setAttribute('aria-expanded', this.expanded.toString());
      this.#toggle.setAttribute('aria-haspopup', 'true');
    }
  }

  toggle() {
    this.expanded = !this.expanded;
  }

  expand() {
    this.expanded = true;
  }

  collapse() {
    this.expanded = false;
  }

  disconnectedCallback() {
    this.#toggle?.removeEventListener('click', this.#handleToggleClick);
    this.removeEventListener('keydown', this.#handleKeydown);
    this.#menu?.removeEventListener('select', this.#handleMenuSelect);
    document.removeEventListener('click', this.#handleOutsideClick);
  }

  static {
    customElements.define(this.is, this);
  }
}
