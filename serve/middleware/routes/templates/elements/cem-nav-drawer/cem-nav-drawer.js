// CEM Serve Navigation Drawer - Collapsible navigation drawer component

export class CemNavDrawerChangeEvent extends Event {
  constructor(open) {
    super('change', { bubbles: true, composed: true });
    this.open = open;
  }
}

export class CemServeNavDrawer extends HTMLElement {
  static is = 'cem-nav-drawer';
  static { customElements.define(this.is, this);  }

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
    const toggle = this.#$('toggle');
    const close = this.#$('close');
    const overlay = this.#$('overlay');
    const drawer = this.#$('drawer');

    if (!toggle || !drawer) {
      return;
    }

    // Toggle button click
    toggle.addEventListener('click', () => {
      this.toggle();
    });

    // Close button click
    close?.addEventListener('click', () => {
      this.close();
    });

    // Overlay click to close
    overlay?.addEventListener('click', () => {
      this.close();
    });

    // Escape key to close drawer
    this.shadowRoot.addEventListener('keydown', (e) => {
      if (e.key === 'Escape' && this.open) {
        this.close();
        toggle.focus(); // Return focus to toggle button
      }
    });

    // Focus trap when drawer is open
    drawer.addEventListener('keydown', (e) => {
      if (e.key !== 'Tab' || !this.open) {
        return;
      }

      const focusableElements = this.#getFocusableElements();
      if (focusableElements.length === 0) return;

      const firstElement = focusableElements[0];
      const lastElement = focusableElements[focusableElements.length - 1];

      if (e.shiftKey && document.activeElement === firstElement) {
        // Shift+Tab on first element - go to last
        e.preventDefault();
        lastElement.focus();
      } else if (!e.shiftKey && document.activeElement === lastElement) {
        // Tab on last element - go to first
        e.preventDefault();
        firstElement.focus();
      }
    });

    // Mark current page in navigation and expand parent details
    this.#markCurrentPage();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'open' && oldValue !== newValue) {
      this.#updateAriaStates();

      // Focus close button when opening
      if (this.open) {
        const close = this.#$('close');
        close?.focus();
      }

      // Dispatch change event
      this.dispatchEvent(new CemNavDrawerChangeEvent(this.open));
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

  #updateAriaStates() {
    const toggle = this.#$('toggle');
    const drawer = this.#$('drawer');
    const overlay = this.#$('overlay');

    if (this.open) {
      drawer?.setAttribute('aria-hidden', 'false');
      overlay?.setAttribute('aria-hidden', 'false');
      toggle?.setAttribute('aria-expanded', 'true');
    } else {
      drawer?.setAttribute('aria-hidden', 'true');
      overlay?.setAttribute('aria-hidden', 'true');
      toggle?.setAttribute('aria-expanded', 'false');
    }
  }

  #getFocusableElements() {
    const drawer = this.#$('drawer');
    if (!drawer) return [];

    const elements = drawer.querySelectorAll(
      'a[href], button:not([disabled]), details, [tabindex]:not([tabindex="-1"])'
    );
    return Array.from(elements);
  }

  #markCurrentPage() {
    const currentPath = window.location.pathname;
    const navLinks = this.querySelectorAll('.nav-demo-link');

    navLinks.forEach(link => {
      const href = link.getAttribute('href');
      if (href === currentPath) {
        link.setAttribute('aria-current', 'page');

        // Open the parent details element
        const details = link.closest('details');
        if (details) {
          details.setAttribute('open', '');
        }
      }
    });
  }
}
