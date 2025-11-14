// CEM Serve - Stylesheet Cache
// Provides utilities for loading and caching Constructable Stylesheets
// This enables stylesheet sharing across multiple component instances

// Cache for Constructable Stylesheet objects
const stylesheetCache = new Map();

/**
 * Gets an existing stylesheet from cache or creates a new one
 * @param {string} name - Component name (e.g., 'pfv6-button')
 * @param {string} url - URL to fetch CSS from (e.g., '/__cem/elements/pfv6-button/pfv6-button.css')
 * @returns {Promise<CSSStyleSheet>} The cached or newly created stylesheet
 */
export async function getOrCreateStylesheet(name, url) {
  if (stylesheetCache.has(name)) {
    return stylesheetCache.get(name);
  }

  try {
    const css = await fetch(url).then(r => {
      if (!r.ok) {
        throw new Error(`Failed to fetch CSS: ${r.status} ${r.statusText}`);
      }
      return r.text();
    });

    const sheet = new CSSStyleSheet();
    await sheet.replace(css);
    stylesheetCache.set(name, sheet);
    return sheet;
  } catch (error) {
    console.error(`Failed to load stylesheet for ${name}:`, error);
    throw error;
  }
}

/**
 * Loads a complete component template (HTML and CSS)
 * @param {string} name - Component name (e.g., 'pfv6-button')
 * @returns {Promise<{html: string, stylesheet: CSSStyleSheet}>} The template HTML and stylesheet
 */
export async function loadComponentTemplate(name) {
  try {
    const [html, stylesheet] = await Promise.all([
      fetch(`/__cem/elements/${name}/${name}.html`).then(r => {
        if (!r.ok) {
          throw new Error(`Failed to fetch HTML: ${r.status} ${r.statusText}`);
        }
        return r.text();
      }),
      getOrCreateStylesheet(name, `/__cem/elements/${name}/${name}.css`)
    ]);

    return { html, stylesheet };
  } catch (error) {
    console.error(`Failed to load template for ${name}:`, error);
    throw error;
  }
}
