/**
 * Current version of the kitchen-sink demo package.
 * @summary Package version string
 */
export const VERSION = '1.0.0';

/**
 * Format a custom element tag name for display.
 * @summary Format tag name for display
 * @param tagName - The tag name to format
 * @param uppercase - Whether to uppercase the result
 * @returns The formatted display name
 */
export function formatTagName(tagName: string, uppercase?: boolean): string {
  const name = tagName.replace(/^demo-/, '').replace(/-/g, ' ');
  return uppercase ? name.toUpperCase() : name;
}
