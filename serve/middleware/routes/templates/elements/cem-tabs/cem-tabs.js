// CEM Serve Tabs - Tab navigation component

export class CemTabsChangeEvent extends Event {
  constructor(value) {
    super('change', { bubbles: true, composed: true });
    this.value = value;
  }
}

export class CemServeTabs extends HTMLElement {
  static is = 'cem-tabs';
  static { customElements.define(this.is, this); }

  static get observedAttributes() {
    return ['value'];
  }

  get value() {
    return this.getAttribute('value') || '';
  }

  set value(panelId) {
    if (panelId) {
      this.setAttribute('value', panelId);
    } else {
      this.removeAttribute('value');
    }
  }

  connectedCallback() {
    // Set up tab click handlers and keyboard navigation
    this.#setupTabs();
    this.#setupKeyboardNavigation();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'value' && oldValue !== newValue) {
      this.#updateTabState();

      // Dispatch change event
      this.dispatchEvent(new CemTabsChangeEvent(newValue));
    }
  }

  select(panelId) {
    this.value = panelId;
  }

  #setupTabs() {
    const tabs = this.querySelectorAll('[slot="tab"]');

    tabs.forEach(tab => {
      tab.addEventListener('click', () => {
        const panelId = tab.getAttribute('aria-controls');
        if (panelId) {
          this.value = panelId;
          // Focus the clicked tab
          tab.focus();
        }
      });
    });

    // Set initial state
    this.#updateTabState();
  }

  #setupKeyboardNavigation() {
    this.addEventListener('keydown', (e) => {
      const tabs = Array.from(this.querySelectorAll('[slot="tab"]'));

      // Get the actual focused element, checking through shadow roots
      let currentTab = e.target;

      // Only handle if focus is on a tab
      if (!tabs.includes(currentTab)) return;

      const currentIndex = tabs.indexOf(currentTab);
      let nextIndex = -1;

      switch (e.key) {
        case 'ArrowLeft':
        case 'ArrowUp':
          e.preventDefault();
          nextIndex = currentIndex - 1;
          if (nextIndex < 0) nextIndex = tabs.length - 1;
          break;
        case 'ArrowRight':
        case 'ArrowDown':
          e.preventDefault();
          nextIndex = currentIndex + 1;
          if (nextIndex >= tabs.length) nextIndex = 0;
          break;
        case 'Home':
          e.preventDefault();
          nextIndex = 0;
          break;
        case 'End':
          e.preventDefault();
          nextIndex = tabs.length - 1;
          break;
        default:
          return;
      }

      if (nextIndex !== -1) {
        const nextTab = tabs[nextIndex];
        const panelId = nextTab.getAttribute('aria-controls');
        if (panelId) {
          this.value = panelId;
          nextTab.focus();
        }
      }
    });
  }

  #updateTabState() {
    const currentValue = this.value;
    const tabs = this.querySelectorAll('[slot="tab"]');
    const panels = this.querySelectorAll('[slot="panel"]');

    tabs.forEach(tab => {
      const controlsId = tab.getAttribute('aria-controls');
      if (controlsId === currentValue) {
        tab.setAttribute('aria-selected', 'true');
        tab.setAttribute('tabindex', '0'); // Make selected tab focusable
      } else {
        tab.setAttribute('aria-selected', 'false');
        tab.setAttribute('tabindex', '-1'); // Remove other tabs from tab order
      }
    });

    panels.forEach(panel => {
      if (panel.id === currentValue) {
        panel.setAttribute('active', '');
      } else {
        panel.removeAttribute('active');
      }
    });
  }
}
