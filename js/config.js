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
