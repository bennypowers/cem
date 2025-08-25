#!/usr/bin/env node
// Generate medium custom-elements.json with 75 elements for benchmarking

const fs = require('fs');
const path = require('path');

// Core element categories for medium project
const categories = {
  buttons: {
    base: 'my-button',
    variants: ['primary', 'secondary', 'tertiary', 'ghost', 'outline', 'icon', 'fab', 'toggle', 'close'],
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
    variants: ['text', 'email', 'password', 'number', 'textarea', 'select', 'checkbox', 'radio', 'switch', 'slider', 'date', 'file'],
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
    variants: ['flex', 'grid', 'stack', 'cluster', 'sidebar', 'center', 'box', 'panel', 'section', 'wrapper'],
    baseAttrs: [
      {name: 'gap', type: 'string', description: 'Gap size'},
      {name: 'padding', type: 'string', description: 'Padding size'},
      {name: 'align', type: 'string', description: 'Alignment'},
      {name: 'justify', type: 'string', description: 'Justification'}
    ]
  },
  cards: {
    base: 'my-card',
    variants: ['basic', 'elevated', 'outlined', 'interactive', 'media', 'product', 'profile', 'article'],
    baseAttrs: [
      {name: 'variant', type: 'string', description: 'Card style'},
      {name: 'clickable', type: 'boolean', description: 'Clickable card'},
      {name: 'selected', type: 'boolean', description: 'Selected state'}
    ]
  },
  navigation: {
    base: 'my-nav',
    variants: ['primary', 'breadcrumb', 'tabs', 'sidebar', 'menu', 'dropdown'],
    baseAttrs: [
      {name: 'orientation', type: 'string', description: 'Navigation orientation'},
      {name: 'variant', type: 'string', description: 'Navigation style'},
      {name: 'active', type: 'string', description: 'Active item'}
    ]
  },
  data: {
    base: 'my-table',
    variants: ['basic', 'sortable', 'list', 'grid', 'calendar', 'chart'],
    baseAttrs: [
      {name: 'sortable', type: 'boolean', description: 'Enable sorting'},
      {name: 'selectable', type: 'boolean', description: 'Enable selection'}
    ]
  },
  feedback: {
    base: 'my-alert',
    variants: ['info', 'success', 'warning', 'error', 'toast', 'progress', 'spinner', 'badge'],
    baseAttrs: [
      {name: 'variant', type: 'string', description: 'Alert type'},
      {name: 'dismissible', type: 'boolean', description: 'Can be dismissed'},
      {name: 'icon', type: 'string', description: 'Alert icon'}
    ]
  },
  media: {
    base: 'my-image',
    variants: ['responsive', 'avatar', 'thumbnail', 'gallery', 'video', 'icon'],
    baseAttrs: [
      {name: 'src', type: 'string', description: 'Media source'},
      {name: 'alt', type: 'string', description: 'Alternative text'},
      {name: 'loading', type: 'string', description: 'Loading strategy'}
    ]
  },
  overlays: {
    base: 'my-modal',
    variants: ['dialog', 'drawer', 'tooltip', 'popover', 'dropdown'],
    baseAttrs: [
      {name: 'open', type: 'boolean', description: 'Open state'},
      {name: 'size', type: 'string', description: 'Modal size'},
      {name: 'closable', type: 'boolean', description: 'Show close button'}
    ]
  }
};

function generateElement(category, variant, index) {
  const categoryData = categories[category];
  const tagName = variant === categoryData.variants[0] ? categoryData.base : `${categoryData.base}-${variant}`;
  const className = tagName.split('-').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join('');
  
  const attrs = [...categoryData.baseAttrs];
  
  const slots = [];
  if (category === 'layout' || category === 'cards' || category === 'navigation') {
    slots.push({name: '', description: 'Default content'});
    if (category === 'cards') {
      slots.push({name: 'header', description: 'Header content'});
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
      if (modules.length >= 75) break;
    }
    if (modules.length >= 75) break;
  }
  
  return {
    schemaVersion: '1.0.0',
    readme: 'Medium-sized component library with 75 elements for LSP benchmarking',
    modules
  };
}

// Generate and write the manifest
const manifest = generateManifest();
const outputPath = path.join(__dirname, '..', 'fixtures', 'medium_project', 'custom-elements.json');

fs.writeFileSync(outputPath, JSON.stringify(manifest, null, 2));
console.log(`Generated ${manifest.modules.length} elements in ${outputPath}`);