import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event fired when form field group toggles
 */
export class PfFormFieldGroupToggleEvent extends Event {
  constructor(expanded) {
    super('toggle', { bubbles: true, composed: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Form Field Group
 *
 * A collapsible section within a form, with optional toggle and header.
 *
 * @slot - Form fields/groups to display in the body
 * @slot header-actions - Action buttons in the header
 *
 * @attr {boolean} expandable - Whether the field group can be expanded/collapsed
 * @attr {boolean} expanded - Whether the field group is expanded (only used with expandable)
 * @attr {string} toggle-text - Text for the title
 * @attr {string} description - Description text shown below the title
 *
 * @fires toggle - Fires when expansion state changes
 *
 * @customElement pf-v6-form-field-group
 */
class PfV6FormFieldGroup extends CemElement {
  static is = 'pf-v6-form-field-group';

  static observedAttributes = [
    'expandable',
    'expanded',
    'toggle-text',
    'description'
  ];

  #$ = id => this.shadowRoot.getElementById(id);

  #toggleButton;
  #body;

  async afterTemplateLoaded() {
    this.#toggleButton = this.#$('toggle-button');
    this.#body = this.#$('body');

    // Handle toggle button clicks if expandable
    this.#toggleButton?.addEventListener('click', () => {
      this.toggle();
    });

    // Set expandable attribute if not explicitly set but has toggle button
    if (this.#toggleButton && !this.hasAttribute('expandable')) {
      this.setAttribute('expandable', '');
    }
  }

  // Getters and setters
  get expandable() { return this.hasAttribute('expandable'); }
  set expandable(val) { this.toggleAttribute('expandable', !!val); }

  get expanded() { return this.hasAttribute('expanded'); }
  set expanded(val) { this.toggleAttribute('expanded', !!val); }

  get toggleText() {
    return this.getAttribute('toggle-text') || '';
  }

  set toggleText(val) {
    if (val) {
      this.setAttribute('toggle-text', val);
    } else {
      this.removeAttribute('toggle-text');
    }
  }

  get description() {
    return this.getAttribute('description') || '';
  }

  set description(val) {
    if (val) {
      this.setAttribute('description', val);
    } else {
      this.removeAttribute('description');
    }
  }

  // Public methods
  toggle() {
    if (this.expandable) {
      this.expanded = !this.expanded;
    }
  }

  show() {
    if (this.expandable) {
      this.expanded = true;
    }
  }

  hide() {
    if (this.expandable) {
      this.expanded = false;
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue) return;

    switch (name) {
      case 'expanded':
        this.#updateExpandedState();
        break;

      case 'toggle-text':
        this.#updateToggleText();
        break;
    }
  }

  #updateExpandedState() {
    if (!this.#toggleButton || !this.#body) return;

    const isExpanded = this.hasAttribute('expanded');

    // Update ARIA
    this.#toggleButton.setAttribute('aria-expanded', String(isExpanded));

    // Update body inert state
    this.#body.toggleAttribute('inert', !!isExpanded);

    // Dispatch toggle event
    this.dispatchEvent(new PfFormFieldGroupToggleEvent(isExpanded));
  }

  #updateToggleText() {
    const textSpan = this.#$('toggle-text-default');
    if (textSpan) {
      textSpan.textContent = this.getAttribute('toggle-text') || '';
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
