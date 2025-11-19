/**
 * Base class for CEM web components that use template loading.
 *
 * Handles the common pattern of:
 * - SSR-aware shadow root creation
 * - Automatic template loading with caching
 * - Applying stylesheets via adoptedStyleSheets
 * - Error handling with console warnings
 *
 * Subclasses can override:
 * - static shadowRootOptions = { mode: 'open', delegatesFocus: false }
 * - static elementName = 'custom-name' (defaults to this.localName)
 * - async afterTemplateLoaded() - called after template is loaded and applied
 *
 * @example
 * // Simple element with no additional logic
 * export class MyElement extends CemElement {}
 *
 * @example
 * // Element with setup logic
 * export class MyElement extends CemElement {
 *   async afterTemplateLoaded() {
 *     this.button = this.shadowRoot.getElementById('button');
 *     this.button.addEventListener('click', this.#handleClick);
 *   }
 * }
 *
 * @example
 * // Element with custom shadow root options
 * export class MyElement extends CemElement {
 *   static shadowRootOptions = { mode: 'open', delegatesFocus: true };
 * }
 */
export class CemElement extends HTMLElement {
  /** Cache for Constructable Stylesheet objects shared across all instances */
  static #stylesheetCache = new Map();

  /** Cache for HTML templates shared across all instances */
  static #templateCache = new Map();

  /**
   * Shadow root options. Override in subclass to customize.
   * @type {{ mode: 'open' | 'closed', delegatesFocus?: boolean }}
   */
  static shadowRootOptions = { mode: 'open' };

  /**
   * Element name used for template loading.
   * Defaults to the element's tag name (this.localName).
   * Override if the template name differs from the tag name.
   * @type {string | null}
   */
  static elementName = null;

  /**
   * Fetch element HTML or CSS from the server
   * @param {string} name - Component name (e.g., 'pfv6-button')
   * @param {'html'|'css'} type - File type
   * @returns {Promise<string>} The content
   * @private
   */
  static async #fetchText(name, type) {
    const prettyType = type.toUpperCase();
    try {
      const response = await fetch(`/__cem/elements/${name}/${name}.${type}`);
      if (!response.ok) {
        throw new Error(`Failed to fetch ${prettyType}: ${response.status} ${response.statusText}`);
      }
      return await response.text();
    } catch (error) {
      console.error(`Failed to load ${prettyType} for ${name}:`, error);
      throw error;
    }
  }

  /**
   * Gets an existing stylesheet from cache or creates a new one
   * @param {string} name - Component name (e.g., 'pfv6-button')
   * @returns {Promise<CSSStyleSheet>} The cached or newly created stylesheet
   * @private
   */
  static async #loadCSS(name) {
    if (this.#stylesheetCache.has(name)) {
      return this.#stylesheetCache.get(name);
    }

    const css = await this.#fetchText(name, 'css');
    try {
      const sheet = new CSSStyleSheet();
      await sheet.replace(css);
      this.#stylesheetCache.set(name, sheet);
      return sheet;
    } catch (error) {
      console.error(`Failed to construct stylesheet for ${name}:`, error);
      throw error;
    }
  }

  /**
   * Gets an element template from client-side cache or from the server
   * @param {string} name - Component name (e.g., 'pfv6-button')
   * @returns {Promise<string>} The cached or fetched template string
   * @private
   */
  static async #loadHTML(name) {
    if (this.#templateCache.has(name)) {
      return this.#templateCache.get(name);
    }
    const html = await this.#fetchText(name, 'html');
    this.#templateCache.set(name, html);
    return html;
  }

  /**
   * Loads a complete component template (HTML and CSS)
   * @param {string} name - Component name (e.g., 'pfv6-button')
   * @returns {Promise<{html: string, stylesheet: CSSStyleSheet}>} The template HTML and stylesheet
   * @private
   */
  static async #loadComponentTemplate(name) {
    const [html, stylesheet] = await Promise.all([this.#loadHTML(name), this.#loadCSS(name)]);
    return { html, stylesheet };
  }

  constructor() {
    super();
    const options = this.constructor.shadowRootOptions ?? { mode: 'open' };

    // SSR-aware: only create shadow root if it doesn't exist
    if (!this.shadowRoot) {
      this.attachShadow(options);
    }
  }

  async connectedCallback() {
    // Only populate if shadow root is empty
    if (!this.shadowRoot.firstChild) {
      await this.#populateShadowRoot();
    }

    // Call lifecycle hook for subclasses
    await this.afterTemplateLoaded?.();
  }

  /**
   * Load and apply the component template.
   * Uses the element name from static elementName or falls back to this.localName.
   * @private
   */
  async #populateShadowRoot() {
    const elementName = this.constructor.elementName || this.localName;

    try {
      // Use CemElement explicitly since private static methods aren't inherited
      const { html, stylesheet } = await CemElement.#loadComponentTemplate(elementName);
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
    } catch (error) {
      console.error(`Failed to load ${elementName} template:`, error);
    }
  }

  /**
   * Lifecycle hook called after template is loaded and applied.
   * Override this in subclasses to set up event listeners, DOM references, etc.
   *
   * @example
   * async afterTemplateLoaded() {
   *   this.button = this.shadowRoot.getElementById('button');
   *   this.button.addEventListener('click', () => console.log('clicked'));
   * }
   */
  async afterTemplateLoaded() {
    // Subclasses can override this
  }
}
