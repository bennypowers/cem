class InfoCallout extends HTMLElement {
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

if (!customElements.get('info-callout')) {
  customElements.define('info-callout', InfoCallout);
}