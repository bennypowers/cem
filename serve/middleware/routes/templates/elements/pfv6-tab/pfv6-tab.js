import { loadComponentTemplate } from '/__cem/stylesheet-cache.js';

class Pfv6Tab extends HTMLElement {
  constructor() {
    super();

    // Create shadow root if it doesn't exist (SSR provides it via DSD)
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    // If shadow root is empty (client-side rendering), populate it
    if (this.shadowRoot && !this.shadowRoot.querySelector('slot')) {
      await this.#populateShadowRoot();
    }

    // Notify parent tabs component that a tab was added
    this.dispatchEvent(new Event('pfv6-tab-connected', { bubbles: true, composed: true }));
  }

  disconnectedCallback() {
    // Notify parent tabs component that a tab was removed
    this.dispatchEvent(new Event('pfv6-tab-disconnected', { bubbles: true, composed: true }));
  }

  static get observedAttributes() {
    return ['title'];
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'title' && oldValue !== newValue) {
      // Notify parent that title changed
      this.dispatchEvent(new Event('pfv6-tab-title-changed', { bubbles: true, composed: true }));
    }
  }

  get title() {
    return this.getAttribute('title') || '';
  }

  set title(value) {
    this.setAttribute('title', value);
  }

  async #populateShadowRoot() {
    try {
      // Load template using shared utility with Constructable Stylesheets
      const { html, stylesheet } = await loadComponentTemplate('pfv6-tab');

      // Apply stylesheet using Constructable Stylesheets API
      // This allows stylesheet sharing across multiple tab instances
      this.shadowRoot.adoptedStyleSheets = [stylesheet];

      // Populate shadow root with HTML
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error('Failed to load pfv6-tab template:', error);
      // Fallback UI for when template loading fails
      this.shadowRoot.innerHTML = '<slot></slot>';
    }
  }
}

customElements.define('pfv6-tab', Pfv6Tab);
