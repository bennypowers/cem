// CEM Serve Drawer - Collapsible drawer component

export class CemDrawerChangeEvent extends Event {
  constructor(open) {
    super('change', { bubbles: true, composed: true });
    this.open = open;
  }
}

export class CemServeDrawer extends HTMLElement {
  static is = 'cem-drawer';
  static { customElements.define(this.is, this); }

  static get observedAttributes() {
    return ['open'];
  }

  #$ = id => this.shadowRoot.getElementById(id);

  get open() {
    return this.hasAttribute('open');
  }

  set open(value) {
    const isOpen = Boolean(value);
    if (isOpen) {
      this.setAttribute('open', '');
    } else {
      this.removeAttribute('open');
    }
  }

  connectedCallback() {
    // Set up toggle button
    const toggleButton = this.#$('toggle');
    if (toggleButton) {
      toggleButton.addEventListener('click', () => {
        // Enable transitions on first click
        this.#$('content').classList.add('transitions-enabled');
        this.toggle();
      });
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'open' && oldValue !== newValue) {
      // Dispatch change event
      this.dispatchEvent(new CemDrawerChangeEvent(this.open));
    }
  }

  toggle() {
    this.open = !this.open;
  }

  openDrawer() {
    this.open = true;
  }

  close() {
    this.open = false;
  }
}
