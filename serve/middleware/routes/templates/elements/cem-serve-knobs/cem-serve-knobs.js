import '/__cem/elements/pf-v6-navigation/pf-v6-navigation.js';
import '/__cem/elements/pf-v6-nav-list/pf-v6-nav-list.js';
import '/__cem/elements/pf-v6-nav-item/pf-v6-nav-item.js';
import '/__cem/elements/pf-v6-nav-link/pf-v6-nav-link.js';

import { CemElement } from '/__cem/cem-element.js';

/**
 * CEM Serve Knobs - Container for multi-instance knobs panels
 *
 * @customElement cem-serve-knobs
 */
export class CemServeKnobs extends CemElement {
  static is = 'cem-serve-knobs';

  #navList = null;

  async afterTemplateLoaded() {
    this.#navList = this.shadowRoot.getElementById('nav-list');
    this.#setupNavigation();
    this.#setupPanelSwitching();
  }

  #setupNavigation() {
    // Get all slotted knobs panels (pf-v6-card elements)
    const slot = this.shadowRoot.querySelector('slot');
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

        const panelId = panel.dataset.panel || `instance-${index}`;
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

  #setupPanelSwitching() {
    // Delegate click handling to the nav list
    this.#navList?.addEventListener('click', (e) => {
      const navLink = e.target.closest('pf-v6-nav-link');
      if (!navLink) return;

      const href = navLink.getAttribute('href');
      if (!href || !href.startsWith('#')) return;

      e.preventDefault();

      const panelId = href.substring(1); // Remove #
      const panels = this.querySelectorAll('pf-v6-card');
      const navLinks = this.#navList.querySelectorAll('pf-v6-nav-link');

      // Update current states
      for (const link of navLinks) {
        link.toggleAttribute('current', link === navLink);
      }

      // Update panel visibility
      for (const panel of panels) {
        const isActive = panel.dataset.panel === panelId
        panel.classList.toggle('active', isActive);
        if (isActive) {
          panel.scrollIntoView();
        }
      }
    });
  }

  static {
    customElements.define(this.is, this);
  }
}
