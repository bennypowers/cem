import '/__cem/elements/pf-v6-navigation/pf-v6-navigation.js';
import '/__cem/elements/pf-v6-nav-list/pf-v6-nav-list.js';
import '/__cem/elements/pf-v6-nav-item/pf-v6-nav-item.js';
import '/__cem/elements/pf-v6-nav-link/pf-v6-nav-link.js';
import '/__cem/elements/pf-v6-card/pf-v6-card.js';

import { CemElement } from '/__cem/cem-element.js';

/**
 * CEM Serve Knobs - Container for multi-instance knobs panels
 *
 * @customElement cem-serve-knobs
 */
export class CemServeKnobs extends CemElement {
  static is = 'cem-serve-knobs';

  #$ = selector => this.shadowRoot.querySelector(selector);
  #navList = null;

  async afterTemplateLoaded() {
    this.#navList = this.#$('#nav-list');
    this.#setupNavigation();
    this.#setupHashChangeListener();

    // Handle initial hash on load
    this.#handleHashChange();
  }

  disconnectedCallback() {
    window.removeEventListener('hashchange', this.#handleHashChange);
  }

  #setupNavigation() {
    // Get all slotted knobs panels (pf-v6-card elements)
    const slot = this.#$('slot');
    if (!slot) return;

    const updateNavigation = () => {
      const panels = slot.assignedElements().filter(el => el.tagName === 'PF-V6-CARD');

      if (panels.length === 0) return;

      // Clear existing nav items
      this.#navList.innerHTML = '';

      // Create nav items for each panel
      panels.forEach((panel, index) => {
        const navItem = document.createElement('pf-v6-nav-item');
        const navLink = document.createElement('pf-v6-nav-link');
        const label = document.createElement('span');

        const panelId = panel.dataset.card || `instance-${index}`;
        const labelText = panel.dataset.label || `Instance ${index + 1}`;

        label.className = 'instance-label';
        label.textContent = labelText;

        navLink.setAttribute('href', `#${panelId}`);
        if (index === 0) {
          navLink.setAttribute('current', '');
        }

        navLink.appendChild(label);
        navItem.appendChild(navLink);
        this.#navList.appendChild(navItem);
      });
    };

    // Update navigation when slot content changes
    slot.addEventListener('slotchange', updateNavigation);

    // Initial update
    updateNavigation();
  }

  #setupHashChangeListener() {
    window.addEventListener('hashchange', this.#handleHashChange);
  }

  #handleHashChange = () => {
    const hash = window.location.hash;
    if (!hash || hash === '#') return;

    const cardId = hash.substring(1); // Remove #
    const cards = this.querySelectorAll('pf-v6-card');
    const navLinks = this.#navList?.querySelectorAll('pf-v6-nav-link');
    const knobsContainer = this.#$('#knobs');

    // Update current states for nav links
    if (navLinks) {
      for (const link of navLinks) {
        const linkHref = link.getAttribute('href');
        link.toggleAttribute('current', linkHref === hash);
      }
    }

    // Update card visibility and scroll
    for (const card of cards) {
      const isActive = card.dataset.card === cardId;
      card.classList.toggle('active', isActive);
      if (isActive && knobsContainer) {
        // Scroll the card into view within the #knobs container only
        const cardRect = card.getBoundingClientRect();
        const containerRect = knobsContainer.getBoundingClientRect();

        // Calculate scroll offset needed to bring card into view
        const scrollOffset = cardRect.top - containerRect.top + knobsContainer.scrollTop;

        knobsContainer.scrollTo({
          top: scrollOffset,
          behavior: 'smooth'
        });
      }
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
