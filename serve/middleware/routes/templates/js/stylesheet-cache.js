// CEM Serve - Stylesheet Cache
// Provides utilities for loading and caching Constructable Stylesheets
// This enables stylesheet sharing across multiple component instances

/** Cache for Constructable Stylesheet objects */
const stylesheetCache = new Map();

/** Cache for Element Templates */
const templateCache = new Map();

/**
 * Fetch element html or css
 * @param {string} name - Component name (e.g., 'pf-v6-button')
 * @param {'html'|'css'} type - file type
 * @returns {Promise<string>} The content
 */
async function fetchText(name, type) {
  const prettyType = type.toUpperCase();
  try {
    const html = await fetch(`/__cem/elements/${name}/${name}.${type}`).then(r => {
      if (!r.ok) {
        throw new Error(`Failed to fetch ${prettyType}: ${r.status} ${r.statusText}`);
      }
      return r.text();
    });
    return html;
  } catch (error) {
    console.error(`Failed to load ${prettyType} for ${name}:`, error);
    throw error;
  }
}

/**
 * Gets an existing stylesheet from cache or creates a new one
 * @param {string} name - Component name (e.g., 'pf-v6-button')
 * @returns {Promise<CSSStyleSheet>} The cached or newly created stylesheet
 */
async function loadCSS(name, url) {
  if (stylesheetCache.has(name)) {
    return stylesheetCache.get(name);
  }

  const css = await fetchText(name, 'css');
  try {
    const sheet = new CSSStyleSheet();
    await sheet.replace(css);
    stylesheetCache.set(name, sheet);
    return sheet;
  } catch {
    console.error(`Failed to construct stylesheet for ${name}:`, error);
    throw error;
  }
}

/**
 * Gets an element template from client-side cache or from the server
 * @param {string} name - Component name (e.g., 'pf-v6-button')
 * @returns {Promise<string>} The cached or fetched template string
 */
async function loadHTML(name) {
  if (templateCache.has(name)) {
    return templateCache.get(name);
  }
  return fetchText(name, 'html');
}

/**
 * Loads a complete component template (HTML and CSS)
 * @param {string} name - Component name (e.g., 'pf-v6-button')
 * @returns {Promise<{html: string, stylesheet: CSSStyleSheet}>} The template HTML and stylesheet
 */
export async function loadComponentTemplate(name) {
  const [html, stylesheet] = await Promise.all([loadHTML(name), loadCSS(name)]);
  return { html, stylesheet };
}
