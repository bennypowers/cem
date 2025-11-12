// CEM Color Scheme Toggle - Toggle between light, dark, and system color schemes

export class CemColorSchemeToggle extends HTMLElement {
  static is = 'cem-color-scheme-toggle';
  static { customElements.define(this.is, this); }

  #$ = id => this.shadowRoot.getElementById(id);

  // Storage access gatekeeper - localStorage can throw in Safari private mode
  #getStorageItem(key, defaultValue) {
    try {
      return localStorage.getItem(key) ?? defaultValue;
    } catch (e) {
      return defaultValue;
    }
  }

  #setStorageItem(key, value) {
    try {
      localStorage.setItem(key, value);
    } catch (e) {
      // Storage unavailable (private mode), silently continue
    }
  }

  connectedCallback() {
    if (!this.shadowRoot) return;

    // Restore saved color scheme preference
    const saved = this.#getStorageItem('cem-serve-color-scheme', 'system');
    this.#applyColorScheme(saved);
    this.#updateRadioButtons(saved);

    // Set up event listeners
    const radios = this.shadowRoot.querySelectorAll('input[type="radio"]');
    radios.forEach(radio => {
      radio.addEventListener('change', (e) => {
        if (e.target.checked) {
          this.#applyColorScheme(e.target.value);
        }
      });
    });
  }

  #applyColorScheme(scheme) {
    // Apply to body element
    const body = document.body;

    switch (scheme) {
      case 'light':
        body.style.colorScheme = 'light';
        break;
      case 'dark':
        body.style.colorScheme = 'dark';
        break;
      case 'system':
      default:
        body.style.colorScheme = 'light dark';
        break;
    }

    // Save preference
    this.#setStorageItem('cem-serve-color-scheme', scheme);
  }

  #updateRadioButtons(scheme) {
    const radios = this.shadowRoot.querySelectorAll('input[type="radio"]');
    radios.forEach(radio => {
      radio.checked = radio.value === scheme;
    });
  }
}
