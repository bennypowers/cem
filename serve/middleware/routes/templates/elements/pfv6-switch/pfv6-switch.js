import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

class Pfv6Switch extends HTMLElement {
  #input = null;

  constructor() {
    super();

    // Create shadow root if it doesn't exist (SSR provides it via DSD)
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    // If shadow root is empty (client-side rendering), populate it
    this.#input = this.shadowRoot?.getElementById('switch-input');
    if (!this.#input && this.shadowRoot) {
      await this.#populateShadowRoot();
      this.#input = this.shadowRoot.getElementById('switch-input');
    }

    if (!this.#input) return;

    // Sync initial state
    this.#syncAttributes();

    // Forward change events from internal checkbox and sync checked state
    this.#input.addEventListener('change', () => {
      // Sync the checked attribute to match internal input
      if (this.#input.checked) {
        this.setAttribute('checked', '');
      } else {
        this.removeAttribute('checked');
      }
      this.dispatchEvent(new Event('change', { bubbles: true }));
    });

    // Forward input events from internal checkbox
    this.#input.addEventListener('input', () => {
      this.dispatchEvent(new Event('input', { bubbles: true }));
    });
  }

  static get observedAttributes() {
    return ['checked', 'disabled', 'aria-label', 'aria-labelledby'];
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  async #populateShadowRoot() {
    try {
      // Load template using shared utility with Constructable Stylesheets
      const { html, stylesheet } = await loadComponentTemplate('pfv6-switch');

      // Apply stylesheet using Constructable Stylesheets API
      // This allows stylesheet sharing across multiple switch instances
      this.shadowRoot.adoptedStyleSheets = [stylesheet];

      // Populate shadow root with HTML
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-switch template:', error);
      // Fallback UI for when template loading fails
      this.shadowRoot.innerHTML = '<input type="checkbox" id="switch-input" role="switch"> Switch (template failed to load)';
    }
  }

  #syncAttributes() {
    if (!this.#input) return;

    // Sync checked state
    this.#input.checked = this.hasAttribute('checked');

    // Sync disabled state
    this.#input.disabled = this.hasAttribute('disabled');

    // Sync aria attributes
    const ariaLabel = this.getAttribute('aria-label');
    if (ariaLabel) {
      this.#input.setAttribute('aria-label', ariaLabel);
    } else {
      this.#input.removeAttribute('aria-label');
    }

    const ariaLabelledby = this.getAttribute('aria-labelledby');
    if (ariaLabelledby) {
      this.#input.setAttribute('aria-labelledby', ariaLabelledby);
    } else {
      this.#input.removeAttribute('aria-labelledby');
    }
  }

  // Public API
  get checked() {
    return this.hasAttribute('checked');
  }

  set checked(value) {
    if (value) {
      this.setAttribute('checked', '');
    } else {
      this.removeAttribute('checked');
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

  get value() {
    return this.checked;
  }

  set value(val) {
    this.checked = val;
  }
}

customElements.define('pfv6-switch', Pfv6Switch);
