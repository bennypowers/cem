import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event fired when expandable section toggles
 */
export class PfExpandableSectionToggleEvent extends Event {
  constructor(expanded) {
    super('toggle', { bubbles: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Expandable Section
 *
 * A disclosure component that shows/hides content with smooth animations.
 *
 * @slot toggle-icon - Custom toggle icon (defaults to chevron)
 * @slot toggle-text - Text for the toggle button
 * @slot - Content to show/hide
 *
 * @attr {boolean} expanded - Whether the section is expanded
 * @attr {boolean} expand-top - Expand animation goes upward
 * @attr {boolean} expand-bottom - Expand animation goes downward
 * @attr {boolean} display-lg - Large display variant with background and border
 * @attr {boolean} indented - Indent content to align with toggle text
 * @attr {boolean} limit-width - Limit content width to 46.875rem
 * @attr {boolean} truncate - Show truncated content with line clamp when collapsed
 * @attr {string} toggle-text - Text for the toggle button (alternative to slot)
 *
 * @fires toggle - Fires when expansion state changes
 *
 * @customElement pf-v6-expandable-section
 */
class PfV6ExpandableSection extends CemElement {
  static is = 'pf-v6-expandable-section';

  static observedAttributes = [
    'expanded',
    'expand-top',
    'expand-bottom',
    'display-lg',
    'indented',
    'limit-width',
    'truncate',
    'toggle-text'
  ];

  #$ = id => this.shadowRoot.getElementById(id);

  #toggleButton;
  #content;

  afterTemplateLoaded() {
    this.#toggleButton = this.#$('toggle-button');
    this.#content = this.#$('content');
    this.#toggleButton?.addEventListener('click', () => this.toggle());
  }

  // Getters and setters
  get expanded() { return this.hasAttribute('expanded'); }
  set expanded(val) { this.toggleAttribute('expanded', !!val); }

  get expandTop() { return this.hasAttribute('expand-top'); }
  set expandTop(val) { this.toggleAttribute('expand-top', !!val); }

  get expandBottom() { return this.hasAttribute('expand-bottom'); }
  set expandBottom(val) { this.toggleAttribute('expand-bottom', !!val); }

  get displayLg() { return this.hasAttribute('display-lg'); }
  set displayLg(val) { this.toggleAttribute('display-lg', !!val); }

  get indented() { return this.hasAttribute('indented'); }
  set indented(val) { this.toggleAttribute('indented', !!val); }

  get limitWidth() { return this.hasAttribute('limit-width'); }
  set limitWidth(val) { this.toggleAttribute('limit-width', !!val); }

  get truncate() { return this.hasAttribute('truncate'); }
  set truncate(val) { this.toggleAttribute('truncate', !!val); }

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

  // Public methods
  toggle() {
    this.expanded = !this.expanded;
  }

  show() {
    this.expanded = true;
  }

  hide() {
    this.expanded = false;
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue || !this.#toggleButton || !this.#content) return;

    switch (name) {
      case 'expanded':
        this.#updateExpandedState();
        break;

      case 'toggle-text':
        this.#updateToggleText();
        break;

      case 'expand-top':
        this.#updateToggleIcon();
        break;
    }
  }

  #updateExpandedState() {
    const isExpanded = this.hasAttribute('expanded');

    // Update ARIA
    this.#toggleButton.setAttribute('aria-expanded', String(isExpanded));

    // Update content visibility
    this.#content.toggleAttribute('hidden', !isExpanded);

    // Dispatch toggle event
    this.dispatchEvent(new PfExpandableSectionToggleEvent(isExpanded));
  }

  #updateToggleText() {
    const textSpan = this.#$('toggle-text-default');
    if (textSpan) {
      textSpan.textContent = this.getAttribute('toggle-text') || '';
    }
  }

  #updateToggleIcon() {
    const toggleIcon = this.#$('toggle-icon');
    if (toggleIcon) {
      toggleIcon.classList.toggle('pf-m-expand-top', this.hasAttribute('expand-top'));
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
