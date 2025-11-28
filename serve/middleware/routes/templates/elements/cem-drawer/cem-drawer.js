import '/__cem/elements/pf-v6-button/pf-v6-button.js';

// CEM Serve Drawer - Collapsible drawer component

export class CemDrawerChangeEvent extends Event {
  constructor(open) {
    super('change', { bubbles: true });
    this.open = open;
  }
}

export class CemDrawerResizeEvent extends Event {
  constructor(height) {
    super('resize', { bubbles: true });
    this.height = height;
  }
}

export class CemServeDrawer extends HTMLElement {
  static is = 'cem-drawer';

  static observedAttributes = [
    'open',
    'drawer-height',
  ];

  #$(id) {
    return this.shadowRoot?.getElementById(id);
  }

  #isDragging = false;
  #startY = 0;
  #startHeight = 0;
  #resizeDebounceTimer = null;
  #initialized = false;

  get open() {
    return this.hasAttribute('open');
  }

  set open(value) {
    this.toggleAttribute('open', !!value);
  }

  connectedCallback() {
    // Guard against calling before shadowRoot is populated
    if (!this.shadowRoot?.firstChild) {
      return;
    }

    // Only initialize once
    if (this.#initialized) {
      return;
    }
    this.#initialized = true;

    // Apply initial attribute states that may have been set before shadowRoot was ready
    this.#applyInitialState();

    // Set up toggle button
    const toggleButton = this.#$('toggle');
    if (toggleButton) {
      toggleButton.addEventListener('click', () => {
        // Enable transitions on first click
        const content = this.#$('content');
        if (content) {
          content.classList.add('transitions-enabled');
        }
        this.toggle();
      });
    }

    // Set up resize handle
    const resizeHandle = this.#$('resize-handle');
    const content = this.#$('content');
    if (resizeHandle && content) {
      resizeHandle.addEventListener('mousedown', this.#startResize);
      resizeHandle.addEventListener('keydown', this.#handleKeydown);
    }
  }

  #applyInitialState() {
    // Apply open attribute state if present
    if (this.hasAttribute('open')) {
      const content = this.#$('content');
      const toggleButton = this.#$('toggle');

      if (content) {
        const height = parseInt(this.getAttribute('drawer-height') || '400', 10);
        content.style.height = `${height}px`;
      }

      if (toggleButton) {
        toggleButton.setAttribute('aria-expanded', 'true');
      }
    }
  }

  #startResize = (e) => {
    this.#isDragging = true;
    this.#startY = e.clientY;
    const content = this.#$('content');
    this.#startHeight = content.offsetHeight;

    // Disable transitions during drag
    content.classList.remove('transitions-enabled');

    // Add global event listeners
    document.addEventListener('mousemove', this.#handleResize);
    document.addEventListener('mouseup', this.#stopResize);

    // Prevent text selection during drag
    e.preventDefault();
  }

  #handleResize = (e) => {
    if (!this.#isDragging) return;

    const deltaY = this.#startY - e.clientY; // Inverted because drawer grows upward
    // Allow resizing from 100px minimum up to just below the header (56px)
    const headerHeight = 56; // --cem-dev-server-header-height
    const maxHeight = window.innerHeight - headerHeight;
    const newHeight = Math.max(100, Math.min(maxHeight, this.#startHeight + deltaY));

    const content = this.#$('content');
    content.style.height = `${newHeight}px`;

    // Update ARIA attribute
    this.#updateAriaValueNow(newHeight);
  }

  #handleKeydown = (e) => {
    const content = this.#$('content');
    if (!content) return;

    const step = e.shiftKey ? 50 : 10; // Larger steps with Shift
    let currentHeight = parseInt(content.style.height, 10) || 400;
    let newHeight = currentHeight;

    switch (e.key) {
      case 'ArrowUp':
        e.preventDefault();
        newHeight = currentHeight + step;
        break;
      case 'ArrowDown':
        e.preventDefault();
        newHeight = currentHeight - step;
        break;
      case 'Home': {
        e.preventDefault();
        newHeight = 100; // Minimum height
        break;
      }
      case 'End': {
        e.preventDefault();
        const headerHeight = 56;
        newHeight = window.innerHeight - headerHeight; // Maximum height
        break;
      }
      default:
        return;
    }

    // Clamp to min/max
    const headerHeight = 56;
    const maxHeight = window.innerHeight - headerHeight;
    newHeight = Math.max(100, Math.min(maxHeight, newHeight));

    content.style.height = `${newHeight}px`;
    this.#updateAriaValueNow(newHeight);

    // Debounce resize event (300ms)
    clearTimeout(this.#resizeDebounceTimer);
    this.#resizeDebounceTimer = setTimeout(() => {
      this.dispatchEvent(new CemDrawerResizeEvent(newHeight));
    }, 300);
  }

  #updateAriaValueNow(height) {
    const resizeHandle = this.#$('resize-handle');
    if (resizeHandle) {
      resizeHandle.setAttribute('aria-valuenow', Math.round(height));

      // Update max based on window size
      const headerHeight = 56;
      const maxHeight = window.innerHeight - headerHeight;
      resizeHandle.setAttribute('aria-valuemax', Math.round(maxHeight));
    }
  }

  #stopResize = () => {
    this.#isDragging = false;

    // Re-enable transitions
    const content = this.#$('content');
    content.classList.add('transitions-enabled');

    // Debounce the resize event dispatch (300ms)
    clearTimeout(this.#resizeDebounceTimer);
    this.#resizeDebounceTimer = setTimeout(() => {
      const height = parseInt(content.style.height, 10);
      this.dispatchEvent(new CemDrawerResizeEvent(height));
    }, 300);

    // Remove global event listeners
    document.removeEventListener('mousemove', this.#handleResize);
    document.removeEventListener('mouseup', this.#stopResize);
  }

  attributeChangedCallback(name, oldValue, newValue) {
    // Guard against calling before shadowRoot is populated
    if (!this.shadowRoot?.firstChild) {
      return;
    }

    if (name === 'open' && oldValue !== newValue) {
      const content = this.#$('content');
      const toggleButton = this.#$('toggle');

      const {open} = this;
      if (content) {
        if (open) {
          // Restore height from drawer-height attribute (default 400 if not set)
          const height = parseInt(this.getAttribute('drawer-height') || '400', 10);
          content.style.height = `${height}px`;
        } else {
          // Reset to 0 when closing
          content.style.height = '0px';
        }
      }

      // Update aria-expanded on toggle button
      if (toggleButton) {
        toggleButton.setAttribute('aria-expanded', open.toString());
      }

      // Dispatch change event
      this.dispatchEvent(new CemDrawerChangeEvent(open));
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

  static {
    customElements.define(this.is, this);
  }
}
