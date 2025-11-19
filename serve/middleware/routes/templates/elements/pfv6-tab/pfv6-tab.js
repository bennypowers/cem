import { CemElement } from '/__cem/cem-element.js';

class Pfv6Tab extends CemElement {
  static observedAttributes = ['title'];

  async afterTemplateLoaded() {
    // Notify parent tabs component that a tab was added
    this.dispatchEvent(new Event('pfv6-tab-connected', { bubbles: true, composed: true }));
  }

  disconnectedCallback() {
    // Notify parent tabs component that a tab was removed
    this.dispatchEvent(new Event('pfv6-tab-disconnected', { bubbles: true, composed: true }));
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
