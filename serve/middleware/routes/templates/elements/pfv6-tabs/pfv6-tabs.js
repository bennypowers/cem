import { CemElement } from '/__cem/cem-element.js';

class Pfv6Tabs extends CemElement {
  #tabsContainer;
  #selectedIndex = 0;
  #tabs = [];
  #mutationObserver;

  async afterTemplateLoaded() {
    this.#tabsContainer = this.shadowRoot.querySelector('.tabs');
    if (!this.#tabsContainer) return;

    // Listen for tab child changes
    this.addEventListener('pfv6-tab-connected', () => this.#updateTabs());
    this.addEventListener('pfv6-tab-disconnected', () => this.#updateTabs());
    this.addEventListener('pfv6-tab-title-changed', () => this.#updateTabs());

    // Observe child changes
    this.#mutationObserver = new MutationObserver(() => this.#updateTabs());
    this.#mutationObserver.observe(this, { childList: true });

    // Initial render
    this.#updateTabs();

    // Restore selected index from attribute
    if (this.hasAttribute('selected')) {
      this.selectedIndex = parseInt(this.getAttribute('selected'), 10);
    }

    // Initialize accent line position after render
    requestAnimationFrame(() => {
      this.#updateAccentLine(this.#selectedIndex);
    });
  }

  disconnectedCallback() {
    this.#mutationObserver?.disconnect();
  }

  #updateTabs() {
    // Get all pfv6-tab children
    this.#tabs = Array.from(this.querySelectorAll('pfv6-tab'));

    // Clear existing buttons and panels
    this.#tabsContainer.innerHTML = '';
    const panelsContainer = this.shadowRoot.querySelector('.panels');
    panelsContainer.innerHTML = '';

    // Generate buttons and panel slots for each tab
    this.#tabs.forEach((tab, index) => {
      const slotName = `panel-${index}`;

      // Create button in shadow DOM
      const button = document.createElement('button');
      button.className = 'tab';
      button.type = 'button';
      button.role = 'tab';
      button.id = `tab-${index}`;
      button.setAttribute('aria-controls', `panel-${index}`);
      button.setAttribute('aria-selected', index === this.#selectedIndex ? 'true' : 'false');
      button.tabIndex = index === this.#selectedIndex ? 0 : -1;

      // Get title from attribute or slot
      const titleSlot = tab.querySelector('[slot="title"]');
      if (titleSlot) {
        button.appendChild(titleSlot.cloneNode(true));
      } else {
        button.textContent = tab.title || `Tab ${index + 1}`;
      }

      // Handle click
      button.addEventListener('click', () => {
        this.selectedIndex = index;
      });

      this.#tabsContainer.appendChild(button);

      // Create panel in shadow DOM with named slot
      const panel = document.createElement('div');
      panel.className = 'panel';
      panel.id = `panel-${index}`;
      panel.role = 'tabpanel';
      panel.setAttribute('aria-labelledby', `tab-${index}`);
      panel.hidden = index !== this.#selectedIndex;

      const slot = document.createElement('slot');
      slot.name = slotName;
      panel.appendChild(slot);

      panelsContainer.appendChild(panel);

      // Assign slot name to the pfv6-tab element
      tab.setAttribute('slot', slotName);
    });

    // Handle keyboard navigation
    this.#tabsContainer.addEventListener('keydown', (e) => this.#handleKeyDown(e));
  }

  #handleKeyDown(e) {
    const buttons = Array.from(this.#tabsContainer.querySelectorAll('.tab'));
    const currentIndex = buttons.findIndex(btn => btn === document.activeElement);

    if (currentIndex === -1) return;

    let nextIndex = currentIndex;

    switch (e.key) {
      case 'ArrowLeft':
        e.preventDefault();
        nextIndex = currentIndex > 0 ? currentIndex - 1 : buttons.length - 1;
        break;
      case 'ArrowRight':
        e.preventDefault();
        nextIndex = currentIndex < buttons.length - 1 ? currentIndex + 1 : 0;
        break;
      case 'Home':
        e.preventDefault();
        nextIndex = 0;
        break;
      case 'End':
        e.preventDefault();
        nextIndex = buttons.length - 1;
        break;
      default:
        return;
    }

    this.selectedIndex = nextIndex;
    buttons[nextIndex].focus();
  }

  get selectedIndex() {
    return this.#selectedIndex;
  }

  set selectedIndex(index) {
    const newIndex = Math.max(0, Math.min(index, this.#tabs.length - 1));
    if (newIndex === this.#selectedIndex) return;

    this.#selectedIndex = newIndex;
    this.setAttribute('selected', String(newIndex));

    // Update button states
    const buttons = Array.from(this.#tabsContainer.querySelectorAll('.tab'));
    buttons.forEach((button, i) => {
      button.setAttribute('aria-selected', i === newIndex ? 'true' : 'false');
      button.tabIndex = i === newIndex ? 0 : -1;
    });

    // Update panel visibility
    const panels = Array.from(this.shadowRoot.querySelectorAll('.panel'));
    panels.forEach((panel, i) => {
      panel.hidden = i !== newIndex;
    });

    // Update animated accent line position
    this.#updateAccentLine(newIndex);

    // Dispatch change event
    this.dispatchEvent(new CustomEvent('change', {
      detail: { selectedIndex: newIndex },
      bubbles: true,
      composed: true
    }));
  }

  #updateAccentLine(index) {
    const buttons = Array.from(this.#tabsContainer.querySelectorAll('.tab'));
    const activeButton = buttons[index];

    if (!activeButton) return;

    // Get the ::before pseudo-element's position (the visual background, not the full button)
    const containerRect = this.#tabsContainer.getBoundingClientRect();
    const buttonRect = activeButton.getBoundingClientRect();

    // Account for the 8px inset from padding (item padding simulation)
    const inset = 8; // var(--pf-t--global--spacer--sm)
    const start = (buttonRect.left - containerRect.left) + inset;
    const length = buttonRect.width - (inset * 2);

    // Update private CSS custom properties for animated accent line
    this.#tabsContainer.style.setProperty('--_pfv6-tabs--link-accent--start', `${start}px`);
    this.#tabsContainer.style.setProperty('--_pfv6-tabs--link-accent--length', `${length}px`);
  }
}

customElements.define('pfv6-tabs', Pfv6Tabs);
