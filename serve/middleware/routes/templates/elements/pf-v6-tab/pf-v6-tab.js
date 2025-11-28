import { CemElement } from '/__cem/cem-element.js';

/**
 * @customElement pf-v6-tab
 */
class PfV6Tab extends CemElement {
  static observedAttributes = ['title'];
  static is = 'pf-v6-tab';

  async afterTemplateLoaded() {
    // Notify parent tabs component that a tab was added
    this.dispatchEvent(new Event('pf-v6-tab-connected', { bubbles: true }));
  }

  disconnectedCallback() {
    // Notify parent tabs component that a tab was removed
    this.dispatchEvent(new Event('pf-v6-tab-disconnected', { bubbles: true }));
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'title' && oldValue !== newValue) {
      // Notify parent that title changed
      this.dispatchEvent(new Event('pf-v6-tab-title-changed', { bubbles: true }));
    }
  }

  get title() {
    return this.getAttribute('title') || '';
  }

  set title(value) {
    this.setAttribute('title', value);
  }

  static {
    customElements.define(this.is, this);
  }
}
