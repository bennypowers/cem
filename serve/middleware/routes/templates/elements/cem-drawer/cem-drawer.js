// CEM Serve Drawer - Collapsible drawer component

export class CemDrawerChangeEvent extends Event {
  constructor(open) {
    super('change', { bubbles: true, composed: true });
    this.open = open;
  }
}

export class CemDrawerResizeEvent extends Event {
  constructor(height) {
    super('resize', { bubbles: true, composed: true });
    this.height = height;
  }
}

export class CemServeDrawer extends HTMLElement {
  static is = 'cem-drawer';
  static { customElements.define(this.is, this); }

  static get observedAttributes() {
    return ['open'];
  }

  #$ = id => this.shadowRoot.getElementById(id);
  #isDragging = false;
  #startY = 0;
  #startHeight = 0;
  #resizeDebounceTimer = null;

  get open() {
    return this.hasAttribute('open');
  }

  set open(value) {
    this.toggleAttribute('open', !!value);
  }

  connectedCallback() {
    // Set up toggle button
    const toggleButton = this.#$('toggle');
    if (toggleButton) {
      toggleButton.addEventListener('click', () => {
        // Enable transitions on first click
        this.#$('content').classList.add('transitions-enabled');
        this.toggle();
      });
    }

    // Set up resize handle
    const resizeHandle = this.#$('resize-handle');
    const content = this.#$('content');
    if (resizeHandle && content) {
      resizeHandle.addEventListener('mousedown', this.#startResize);
    }

    // Set initial height when open
    if (this.open) {
      content.style.height = '400px';
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
    if (name === 'open' && oldValue !== newValue) {
      const content = this.#$('content');
      if (content) {
        if (this.open) {
          // Set default height when opening (if not already set by resize)
          if (!content.style.height || content.style.height === '0px') {
            content.style.height = '400px';
          }
        } else {
          // Reset to 0 when closing
          content.style.height = '0px';
        }
      }

      // Dispatch change event
      this.dispatchEvent(new CemDrawerChangeEvent(this.open));
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
}
