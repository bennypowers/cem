export class InfoCallout extends HTMLElement {
  static is = 'info-callout';
  static { customElements.define(this.is, this); }

  constructor() {
    super();
    this.setupTitle();
    this.setupCollapsible();
  }

  setupTitle() {
    const title = this.getAttribute('title');
    if (title) {
      const titleEl = this.shadowRoot.querySelector('.title');
      titleEl.textContent = title;
    }
  }

  setupCollapsible() {
    if (this.hasAttribute('collapsible')) {
      const header = this.shadowRoot.querySelector('.header');
      header.addEventListener('click', () => {
        this.toggleAttribute('expanded');
      });
    } else {
      this.setAttribute('expanded', '');
    }
  }
}
