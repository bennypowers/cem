/**
 * Helper for loading server-side rendered templates in tests.
 * Used for components that extend HTMLElement and expect SSR shadow DOM.
 *
 * @param {HTMLElement} element - The element to attach shadow DOM to
 * @param {string} componentName - Component name (e.g., 'cem-drawer')
 * @returns {Promise<void>}
 */
export async function loadSSRTemplate(element, componentName) {
  // Fetch real template from server
  const [htmlResponse, cssResponse] = await Promise.all([
    fetch(`/__cem/elements/${componentName}/${componentName}.html`),
    fetch(`/__cem/elements/${componentName}/${componentName}.css`),
  ]);

  const html = await htmlResponse.text();
  const css = await cssResponse.text();

  // Attach shadow DOM with real server template
  element.attachShadow({ mode: 'open' });
  element.shadowRoot.innerHTML = html;

  // Apply styles
  const styleSheet = new CSSStyleSheet();
  styleSheet.replaceSync(css);
  element.shadowRoot.adoptedStyleSheets = [styleSheet];
}
