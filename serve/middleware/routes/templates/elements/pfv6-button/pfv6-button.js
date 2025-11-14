import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

/**
 * PatternFly v6 inspired button component
 *
 * @attr {string} variant - Button variant: primary, secondary, tertiary (default), danger, plain, link
 * @attr {string} size - Button size: sm, md (default), lg
 * @attr {boolean} block - Make button full width
 * @attr {boolean} disabled - Disable the button
 * @attr {string} aria-label - Accessible label for the button
 *
 * @slot - Default slot for button text/content
 * @slot icon-start - Slot for icon before text
 * @slot icon-end - Slot for icon after text
 *
 * @fires click - Bubbles click events from internal button
 */
class Pfv6Button extends HTMLElement {
  static get observedAttributes() {
    return ['disabled', 'aria-label', 'aria-expanded', 'aria-controls', 'aria-haspopup', 'type'];
  }

  #internals;

  constructor() {
    super();

    // Attach ElementInternals for cross-root ARIA references
    this.#internals = this.attachInternals();

    // Don't recreate shadow root if it already exists (from SSR)
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open', delegatesFocus: true });
    }
  }

  async connectedCallback() {
    // If shadow root is empty (client-side rendering), populate it
    let button = this.shadowRoot?.querySelector('button');
    if (!button && this.shadowRoot) {
      await this.#populateShadowRoot();
      button = this.shadowRoot.querySelector('button');
    }

    if (!button) return;

    // Forward attributes to internal button
    this.#syncAttributes(button);

    // Delegate button events
    button.addEventListener('click', (e) => {
      if (this.disabled) {
        e.preventDefault();
        e.stopPropagation();
        return;
      }
      // Event bubbles naturally through shadow DOM
    });
  }

  attributeChangedCallback(name, oldValue, newValue) {
    // Only sync if shadow root is ready
    const button = this.shadowRoot?.querySelector('button');
    if (button) {
      this.#syncAttributes(button);
    }
  }

  async #populateShadowRoot() {
    try {
      // Load template using shared utility with Constructable Stylesheets
      const { html, stylesheet } = await loadComponentTemplate('pfv6-button');

      // Apply stylesheet using Constructable Stylesheets API
      // This allows stylesheet sharing across multiple button instances
      this.shadowRoot.adoptedStyleSheets = [stylesheet];

      // Populate shadow root with HTML
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-button template:', error);
      // Fallback UI for when template loading fails
      this.shadowRoot.innerHTML = '<button>Button (template failed to load)</button>';
    }
  }

  #syncAttributes(button) {
    // Sync disabled
    if (this.hasAttribute('disabled')) {
      button.disabled = true;
    } else {
      button.disabled = false;
    }

    // Cross-root ARIA references go on ElementInternals (host)
    // This allows aria-controls to reference elements outside shadow root
    if (this.hasAttribute('aria-controls')) {
      this.#internals.ariaControls = this.getAttribute('aria-controls');
    } else {
      this.#internals.ariaControls = null;
    }

    // Non-reference ARIA attributes go on internal button
    // These don't need to cross shadow boundaries
    const buttonAriaAttrs = ['aria-label', 'aria-expanded', 'aria-haspopup'];
    buttonAriaAttrs.forEach(attr => {
      if (this.hasAttribute(attr)) {
        button.setAttribute(attr, this.getAttribute(attr));
      } else {
        button.removeAttribute(attr);
      }
    });

    // Sync type attribute
    if (this.hasAttribute('type')) {
      button.setAttribute('type', this.getAttribute('type'));
    }
  }

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    if (value) {
      this.setAttribute('disabled', '');
    } else {
      this.removeAttribute('disabled');
    }
  }
}

customElements.define('pfv6-button', Pfv6Button);
