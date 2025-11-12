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
    // Set up tab click handlers
    this.#setupTabs();
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
        }
      });
    });

    // Set initial state
    this.#updateTabState();
  }

  #updateTabState() {
    const currentValue = this.value;
    const tabs = this.querySelectorAll('[slot="tab"]');
    const panels = this.querySelectorAll('[slot="panel"]');

    tabs.forEach(tab => {
      const controlsId = tab.getAttribute('aria-controls');
      if (controlsId === currentValue) {
        tab.setAttribute('aria-selected', 'true');
      } else {
        tab.setAttribute('aria-selected', 'false');
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
