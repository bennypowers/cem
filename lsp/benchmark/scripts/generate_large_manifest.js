#!/usr/bin/env node
// Generate large custom-elements.json with 300 elements for benchmarking

const fs = require('fs');
const path = require('path');

// Element categories and their definitions
const categories = {
  buttons: {
    base: 'my-button',
    variants: ['primary', 'secondary', 'tertiary', 'ghost', 'outline', 'text', 'icon', 'fab', 'toggle', 'split', 'group', 'toolbar', 'dropdown', 'pagination', 'breadcrumb', 'close', 'expand', 'collapse', 'play', 'pause', 'stop', 'next', 'previous', 'submit', 'reset', 'cancel', 'save', 'edit', 'delete', 'copy', 'share', 'like', 'bookmark', 'favorite', 'download', 'upload', 'print', 'search', 'filter', 'sort', 'refresh', 'sync', 'settings', 'help', 'info', 'warning', 'error', 'success'],
    baseAttrs: [
      {name: 'variant', type: 'string', description: 'Button style variant'},
      {name: 'size', type: 'string', description: 'Button size'},
      {name: 'disabled', type: 'boolean', description: 'Disabled state'},
      {name: 'loading', type: 'boolean', description: 'Loading state'},
      {name: 'icon', type: 'string', description: 'Icon name'}
    ]
  },
  forms: {
    base: 'my-input',
    variants: ['text', 'email', 'password', 'number', 'tel', 'url', 'search', 'textarea', 'select', 'checkbox', 'radio', 'switch', 'slider', 'range', 'color', 'date', 'time', 'datetime', 'file', 'hidden', 'autocomplete', 'combobox', 'listbox', 'multiselect', 'typeahead', 'datepicker', 'timepicker', 'colorpicker', 'filepicker', 'dropdown', 'rating', 'tags', 'chips', 'tokens', 'editor', 'wysiwyg', 'markdown', 'code', 'json', 'numeric', 'currency', 'percentage', 'duration', 'mask', 'format', 'validation'],
    baseAttrs: [
      {name: 'value', type: 'string', description: 'Input value'},
      {name: 'placeholder', type: 'string', description: 'Placeholder text'},
      {name: 'disabled', type: 'boolean', description: 'Disabled state'},
      {name: 'required', type: 'boolean', description: 'Required field'},
      {name: 'readonly', type: 'boolean', description: 'Readonly state'}
    ]
  },
  layout: {
    base: 'my-container',
    variants: ['flex', 'grid', 'stack', 'cluster', 'sidebar', 'switcher', 'cover', 'frame', 'center', 'box', 'panel', 'section', 'article', 'aside', 'header', 'footer', 'main', 'wrapper', 'column'],
    baseAttrs: [
      {name: 'gap', type: 'string', description: 'Gap size'},
      {name: 'padding', type: 'string', description: 'Padding size'},
      {name: 'margin', type: 'string', description: 'Margin size'},
      {name: 'align', type: 'string', description: 'Alignment'},
      {name: 'justify', type: 'string', description: 'Justification'}
    ]
  },
  cards: {
    base: 'my-card',
    variants: ['basic', 'elevated', 'outlined', 'filled', 'interactive', 'media', 'product', 'profile', 'article', 'gallery', 'pricing', 'feature', 'testimonial', 'stats', 'dashboard'],
    baseAttrs: [
      {name: 'variant', type: 'string', description: 'Card style'},
      {name: 'clickable', type: 'boolean', description: 'Clickable card'},
      {name: 'selected', type: 'boolean', description: 'Selected state'},
      {name: 'loading', type: 'boolean', description: 'Loading state'}
    ]
  },
  navigation: {
    base: 'my-nav',
    variants: ['primary', 'secondary', 'breadcrumb', 'pagination', 'tabs', 'stepper', 'sidebar', 'toolbar', 'menu', 'dropdown', 'context', 'mega', 'drawer', 'rail', 'bottom'],
    baseAttrs: [
      {name: 'orientation', type: 'string', description: 'Navigation orientation'},
      {name: 'variant', type: 'string', description: 'Navigation style'},
      {name: 'active', type: 'string', description: 'Active item'},
      {name: 'collapsed', type: 'boolean', description: 'Collapsed state'}
    ]
  },
  data: {
    base: 'my-table',
    variants: ['basic', 'sortable', 'filterable', 'editable', 'virtual', 'tree', 'list', 'grid', 'calendar', 'timeline', 'kanban', 'chart', 'graph', 'diagram', 'map'],
    baseAttrs: [
      {name: 'data', type: 'string', description: 'Data source'},
      {name: 'columns', type: 'string', description: 'Column configuration'},
      {name: 'sortable', type: 'boolean', description: 'Enable sorting'},
      {name: 'filterable', type: 'boolean', description: 'Enable filtering'},
      {name: 'selectable', type: 'boolean', description: 'Enable selection'}
    ]
  },
  feedback: {
    base: 'my-alert',
    variants: ['info', 'success', 'warning', 'error', 'toast', 'banner', 'dialog', 'snackbar', 'progress', 'spinner', 'skeleton', 'badge', 'chip', 'tag', 'status'],
    baseAttrs: [
      {name: 'variant', type: 'string', description: 'Alert type'},
      {name: 'dismissible', type: 'boolean', description: 'Can be dismissed'},
      {name: 'icon', type: 'string', description: 'Alert icon'},
      {name: 'timeout', type: 'number', description: 'Auto-hide timeout'}
    ]
  },
  media: {
    base: 'my-image',
    variants: ['responsive', 'lazy', 'avatar', 'thumbnail', 'gallery', 'carousel', 'video', 'audio', 'player', 'embed', 'icon', 'logo', 'hero', 'banner', 'background'],
    baseAttrs: [
      {name: 'src', type: 'string', description: 'Media source'},
      {name: 'alt', type: 'string', description: 'Alternative text'},
      {name: 'loading', type: 'string', description: 'Loading strategy'},
      {name: 'fit', type: 'string', description: 'Object fit'},
      {name: 'ratio', type: 'string', description: 'Aspect ratio'}
    ]
  },
  overlays: {
    base: 'my-modal',
    variants: ['dialog', 'drawer', 'sidebar', 'popup', 'tooltip', 'popover', 'dropdown', 'menu', 'sheet', 'panel', 'overlay', 'backdrop', 'mask', 'scrim', 'lightbox'],
    baseAttrs: [
      {name: 'open', type: 'boolean', description: 'Open state'},
      {name: 'size', type: 'string', description: 'Modal size'},
      {name: 'placement', type: 'string', description: 'Placement position'},
      {name: 'closable', type: 'boolean', description: 'Show close button'},
      {name: 'backdrop', type: 'boolean', description: 'Show backdrop'}
    ]
  },
  charts: {
    base: 'my-chart',
    variants: ['line', 'bar', 'pie', 'doughnut', 'area', 'scatter', 'bubble', 'radar', 'polar', 'gauge', 'heatmap', 'treemap', 'sankey', 'funnel', 'candlestick'],
    baseAttrs: [
      {name: 'type', type: 'string', description: 'Chart type'},
      {name: 'data', type: 'string', description: 'Chart data'},
      {name: 'options', type: 'string', description: 'Chart options'},
      {name: 'responsive', type: 'boolean', description: 'Responsive chart'},
      {name: 'animated', type: 'boolean', description: 'Enable animations'}
    ]
  },
  utilities: {
    base: 'my-divider',
    variants: ['horizontal', 'vertical', 'spacer', 'separator', 'border', 'line', 'ruler', 'break', 'gap', 'margin', 'padding', 'offset', 'indent', 'outdent', 'gutter'],
    baseAttrs: [
      {name: 'orientation', type: 'string', description: 'Divider orientation'},
      {name: 'size', type: 'string', description: 'Divider size'},
      {name: 'color', type: 'string', description: 'Divider color'},
      {name: 'style', type: 'string', description: 'Divider style'}
    ]
  }
};

function generateElement(category, variant, index) {
  const categoryData = categories[category];
  const tagName = variant === categoryData.variants[0] ? categoryData.base : `${categoryData.base}-${variant}`;
  const className = tagName.split('-').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join('');
  
  const attrs = [...categoryData.baseAttrs];
  
  // Add variant-specific attributes
  if (variant.includes('color')) {
    attrs.push({name: 'color', type: 'string', description: 'Color variant'});
  }
  if (variant.includes('size') || variant.includes('large') || variant.includes('small')) {
    attrs.push({name: 'size', type: 'string', description: 'Size variant'});
  }
  if (variant.includes('multi') || variant.includes('group')) {
    attrs.push({name: 'multiple', type: 'boolean', description: 'Multiple selection'});
  }
  
  const slots = [];
  if (category === 'layout' || category === 'cards' || category === 'navigation') {
    slots.push({name: '', description: 'Default content'});
    if (variant.includes('header') || category === 'cards') {
      slots.push({name: 'header', description: 'Header content'});
    }
    if (variant.includes('footer') || category === 'cards') {
      slots.push({name: 'footer', description: 'Footer content'});
    }
  }
  
  const events = [];
  if (category === 'buttons' || category === 'forms') {
    events.push({name: 'click', type: {text: 'MouseEvent'}, description: 'Click event'});
  }
  if (category === 'forms') {
    events.push({name: 'change', type: {text: 'Event'}, description: 'Change event'});
  }
  
  return {
    kind: 'javascript-module',
    path: `./src/components/${category}/${variant}.ts`,
    declarations: [{
      kind: 'class',
      description: `${variant.charAt(0).toUpperCase() + variant.slice(1)} ${category.slice(0, -1)} component`,
      name: className,
      tagName,
      customElement: true,
      attributes: attrs.map(attr => ({
        name: attr.name,
        type: {text: attr.type === 'string' ? 'string' : attr.type === 'boolean' ? 'boolean' : 'number'},
        description: attr.description
      })),
      ...(slots.length > 0 && {slots}),
      ...(events.length > 0 && {events})
    }],
    exports: [{
      kind: 'custom-element-definition',
      name: tagName,
      declaration: {name: className}
    }]
  };
}

function generateManifest() {
  const modules = [];
  let index = 0;
  
  for (const [category, data] of Object.entries(categories)) {
    for (const variant of data.variants) {
      modules.push(generateElement(category, variant, index++));
      if (modules.length >= 300) break;
    }
    if (modules.length >= 300) break;
  }
  
  return {
    schemaVersion: '1.0.0',
    readme: 'Comprehensive component library with 300+ elements for LSP benchmarking and stress testing',
    modules
  };
}

// Generate and write the manifest
const manifest = generateManifest();
const outputPath = path.join(__dirname, '..', 'fixtures', 'large_project', 'custom-elements.json');

fs.writeFileSync(outputPath, JSON.stringify(manifest, null, 2));
console.log(`Generated ${manifest.modules.length} elements in ${outputPath}`);