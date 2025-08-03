const modes = ['light', 'auto', 'dark'];

export class ColorModeToggle extends HTMLElement {
  static is = 'color-mode-toggle';
  static { customElements.define(this.is, this); }

  #currentMode = 'auto';
  #storageKey = 'theme';

  connectedCallback() {
    this.#initializeMode();
    this.#setupEventListeners();
  }

  #initializeMode() {
    // Get stored preference or default to auto
    const stored = localStorage.getItem(this.#storageKey);
    this.#currentMode = modes.includes(stored) ? stored : 'auto';
    this.#applyMode(this.#currentMode);
    this.#updateButtons();
  }

  #setupEventListeners() {
    this.shadowRoot.addEventListener('change', this.#onChange.bind(this));
    // Listen for system color scheme changes
    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
    mediaQuery.addEventListener('change', () => {
      if (this.#currentMode === 'auto') {
        this.#applyMode('auto');
      }
    });
  }

  #onChange(event) {
    if (event.target.type === 'radio' && event.target.name === 'color-mode') {
      const mode = event.target.value;
      if (mode && mode !== this.#currentMode) {
        this.#setMode(mode);
      }
    }
  }

  #setMode(mode) {
    this.#currentMode = mode;
    localStorage.setItem(this.#storageKey, mode);
    this.#applyMode(mode);
    this.#updateButtons();
    // Dispatch custom event for other components
    this.dispatchEvent(new CustomEvent('color-mode-change', {
      detail: { mode },
      bubbles: true
    }));
  }

  #applyMode(mode) {
    const html = document.documentElement;
    // Remove existing mode classes/attributes
    html.removeAttribute('data-mode');
    html.classList.remove('light', 'dark');
    switch (mode) {
      case 'light':
        html.setAttribute('data-mode', 'light');
        html.classList.add('light');
        break;
      case 'dark':
        html.setAttribute('data-mode', 'dark');
        html.classList.add('dark');
        break;
      case 'auto':
        html.setAttribute('data-mode', 'auto');
        // Let CSS handle prefers-color-scheme
        break;
    }

    // Update pictures for current mode
    this.#updatePictures(this.#getEffectiveMode(mode));
    // Update external library themes
    this.#updateExternalThemes(this.#getEffectiveMode(mode));
  }

  #getEffectiveMode(mode) {
    if (mode === 'auto') {
      return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }
    return mode;
  }

  #updatePictures(effectiveMode) {
    document.querySelectorAll('picture[data-lit][data-dark]').forEach(picture => {
      const source = picture.firstElementChild;
      if (source && source.tagName === 'IMG') {
        const lightSrc = picture.dataset.lit;
        const darkSrc = picture.dataset.dark;
        source.src = effectiveMode === 'dark' ? darkSrc : lightSrc;
      }
    });
  }

  #updateButtons() {
    const radios = this.shadowRoot.querySelectorAll('input[type="radio"]');
    radios.forEach(radio => {
      radio.checked = radio.value === this.#currentMode;
    });
  }

  #updateExternalThemes(effectiveMode) {
    // Update highlight.js themes
    const hljsLight = document.getElementById('hljs-light');
    const hljsDark = document.getElementById('hljs-dark');
    if (hljsLight) hljsLight.disabled = effectiveMode === 'dark';
    if (hljsDark) hljsDark.disabled = effectiveMode !== 'dark';
    // Update Shoelace themes (for pages that use Shoelace)
    const shoelaceLight = document.getElementById('shoelace-light');
    const shoelaceDark = document.getElementById('shoelace-dark');
    if (shoelaceLight || shoelaceDark) {
      if (effectiveMode === 'dark') {
        if (shoelaceLight) shoelaceLight.disabled = true;
        if (shoelaceDark) shoelaceDark.disabled = false;
        // Add Shoelace dark theme class to html
        document.documentElement.classList.remove('sl-theme-light');
        document.documentElement.classList.add('sl-theme-dark');
      } else {
        if (shoelaceLight) shoelaceLight.disabled = false;
        if (shoelaceDark) shoelaceDark.disabled = true;
        // Add Shoelace light theme class to html
        document.documentElement.classList.remove('sl-theme-dark');
        document.documentElement.classList.add('sl-theme-light');
      }
    }
  }
}
