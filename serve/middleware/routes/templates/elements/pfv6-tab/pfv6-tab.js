class Pfv6Tab extends HTMLElement {
  connectedCallback() {
    // Shadow root is already attached via DSD (server-side rendering)
    // Just notify parent tabs component that a tab was added
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
}

customElements.define('pfv6-tab', Pfv6Tab);
