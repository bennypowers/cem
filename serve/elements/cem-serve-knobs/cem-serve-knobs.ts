import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './cem-serve-knobs.css' with { type: 'css' };

import '../cem-pf-v6-navigation/cem-pf-v6-navigation.js';
import '../cem-pf-v6-nav-list/cem-pf-v6-nav-list.js';
import '../cem-pf-v6-nav-item/cem-pf-v6-nav-item.js';
import '../cem-pf-v6-nav-link/cem-pf-v6-nav-link.js';
import '../cem-pf-v6-card/cem-pf-v6-card.js';
import '../cem-pf-v6-form-field-group/cem-pf-v6-form-field-group.js';

/**
 * CEM Serve Knobs - Container for multi-instance knobs panels
 *
 * @slot - Default slot for cem-pf-v6-card knob panels
 */
@customElement('cem-serve-knobs')
export class CemServeKnobs extends LitElement {
  static styles = styles;

  #navList: Element | null = null;
  #handleHashChange = () => {
    const hash = window.location.hash;
    if (!hash || hash === '#') return;

    const cardId = hash.substring(1);
    const cards = this.querySelectorAll('cem-pf-v6-card');
    const navLinks = this.#navList?.querySelectorAll('cem-pf-v6-nav-link');
    const knobsContainer = this.shadowRoot?.getElementById('knobs');

    if (navLinks) {
      for (const link of navLinks) {
        const linkHref = link.getAttribute('href');
        link.toggleAttribute('current', linkHref === hash);
      }
    }

    for (const card of cards) {
      const isActive = (card as HTMLElement).dataset.card === cardId;
      card.classList.toggle('active', isActive);
      if (isActive && knobsContainer) {
        const cardRect = card.getBoundingClientRect();
        const containerRect = knobsContainer.getBoundingClientRect();
        const scrollOffset = cardRect.top - containerRect.top + knobsContainer.scrollTop;
        knobsContainer.scrollTo({
          top: scrollOffset,
          behavior: 'smooth',
        });
      }
    }
  };

  render() {
    return html`
      <cem-pf-v6-navigation horizontal
                        variant="horizontal-subnav">
        <cem-pf-v6-nav-list id="nav-list"
                        aria-label="Elements">
        </cem-pf-v6-nav-list>
      </cem-pf-v6-navigation>
      <div id="knobs">
        <slot @slotchange=${this.#onSlotChange}></slot>
      </div>
    `;
  }

  connectedCallback() {
    super.connectedCallback();
    window.addEventListener('hashchange', this.#handleHashChange);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener('hashchange', this.#handleHashChange);
  }

  firstUpdated() {
    this.#navList = this.shadowRoot?.getElementById('nav-list') ?? null;
    this.#updateNavigation();
    this.#handleHashChange();
  }

  #onSlotChange() {
    this.#updateNavigation();
  }

  #updateNavigation() {
    const slot = this.shadowRoot?.querySelector('slot');
    if (!slot || !this.#navList) return;

    const panels = slot.assignedElements().filter(el => el.tagName === 'CEM-PF-V6-CARD');
    if (panels.length === 0) return;

    this.#navList.innerHTML = '';

    panels.forEach((panel, index) => {
      const navItem = document.createElement('cem-pf-v6-nav-item');
      const navLink = document.createElement('cem-pf-v6-nav-link');
      const label = document.createElement('span');

      const panelId = (panel as HTMLElement).dataset.card || `instance-${index}`;
      const labelText = (panel as HTMLElement).dataset.label || `Instance ${index + 1}`;

      label.className = 'instance-label';
      label.textContent = labelText;

      navLink.setAttribute('href', `#${panelId}`);
      if (index === 0) {
        navLink.setAttribute('current', '');
      }

      navLink.appendChild(label);
      navItem.appendChild(navLink);
      this.#navList!.appendChild(navItem);
    });
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-serve-knobs': CemServeKnobs;
  }
}
