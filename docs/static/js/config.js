// Configuration module with Hugo templating
// This module contains all values that need Hugo template processing

// Site configuration
export const site_title = `cem`;
export let root_url = '/cem/';
root_url = root_url.startsWith('http') ? root_url : window.location.origin;

// Hugo configuration values
export const code_block_config = {"maximum": 10, "show": true};
export const iconsPath = `icons/`;

// Derived configuration
export const max_lines = code_block_config.maximum;
export const show_lines = code_block_config.show;

// Theme configuration with Hugo values
export const storageKey = `${site_title}-color-mode`;
export const mermaidThemeKey = `${site_title}-mermaid`;

// Translations from Hugo i18n
export const translations = {
  copy_text: 'Copy',
  copied_text: 'Copied!',
  toggle_line_numbers_text: 'Toggle line numbers',
  toggle_line_wrap_text: 'Toggle line wrap',
  resize_snippet: 'Resize snippet',
  not_set: 'not set',
  quick_links: 'Quick links',
  search_results_label: 'Search results',
  short_search_query: 'Query is too short',
  type_to_search: 'Type to search',
  no_matches_found: 'No matches found'
};