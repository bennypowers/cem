# Custom Elements Manifest Browser Component - Implementation Plan

## Overview

Create a sidebar panel component (`cem-manifest-browser`) that displays a searchable, graphical tree view of the Custom Elements Manifest structure. This will be added as a new tab in the existing `cem-drawer` sidebar, similar to the knobs panel.

**Key Features:**
- Searchable tree view of entire manifest structure
- Full-text search (names, descriptions, types, defaults)
- Icons for node types + badges for metadata
- Markdown rendering for summaries and descriptions
- Hierarchical display: modules → elements → APIs (attributes, properties, methods, events, slots, CSS features)

## Component Architecture

### Main Components

1. **`cem-manifest-browser`** - Main controller component
   - Location: `serve/middleware/routes/templates/elements/cem-manifest-browser/`
   - Orchestrates search, tree, and data flow
   - Files: `.js`, `.html`, `.css`

2. **`pf-v6-tree-view`** - PatternFly tree view web component
   - Location: `serve/middleware/routes/templates/elements/pf-v6-tree-view/`
   - Translated from PatternFly React TreeView component
   - Uses PatternFly CSS (compiled from patternfly.org)
   - Files: `.js`, `.css`
   - Reference: `../patternfly/patternfly-react/packages/react-core/src/components/TreeView/TreeView.tsx`
   - CSS: `../patternfly/patternfly/src/patternfly/components/TreeView/tree-view.scss`

3. **`pf-v6-tree-view-list-item`** - Individual tree item
   - Location: `serve/middleware/routes/templates/elements/pf-v6-tree-view/` (same directory)
   - Translated from PatternFly React TreeViewListItem
   - Handles node rendering, expansion, selection, checkboxes
   - Files: Part of pf-v6-tree-view implementation
   - Reference: `../patternfly/patternfly-react/packages/react-core/src/components/TreeView/TreeViewListItem.tsx`

### Integration with cem-drawer

Add manifest browser as **third tab** in `cem-drawer` (alongside Knobs and Server Logs):

**File:** `serve/middleware/routes/templates/elements/cem-serve-chrome/cem-serve-chrome.html`

```html
<cem-drawer>
  <pf-v6-tabs>
    <pf-v6-tab title="Knobs">...</pf-v6-tab>
    <pf-v6-tab title="Server Logs">...</pf-v6-tab>
    <pf-v6-tab title="Manifest">
      <cem-manifest-browser></cem-manifest-browser>
    </pf-v6-tab>
  </pf-v6-tabs>
</cem-drawer>
```

### Data Model

**TreeNode structure (JavaScript):**
```javascript
{
  id: string,           // e.g., "pkg.module-0.element-pf-button"
  type: string,         // 'package'|'module'|'element'|'attribute'|'property'|'method'|'event'|'slot'|'css-property'|'css-part'|'css-state'|'type'
  name: string,         // Display name
  summary: string,      // Brief markdown description
  description: string,  // Full markdown description
  metadata: {
    deprecated: boolean|string,
    required: boolean,
    readonly: boolean,
    privacy: 'public'|'private'|'protected',
    type: string,       // TypeScript type
    default: string,    // Default value
  },
  children: TreeNode[],
  expanded: boolean,
}
```

**Data source:** Manifest JSON embedded in page via `window.__CEM_MANIFEST__` global (following existing debug info pattern)

### Event System

**Use pf-v6-tree-view events directly:**
- `select` - Dispatched by pf-v6-tree-view when node is selected
- `check` - Dispatched when checkbox is toggled
- `expand` - Dispatched when node is expanded
- `collapse` - Dispatched when node is collapsed

**cem-manifest-browser only needs to listen to these pf-v6-tree-view events, not create its own.**

## PatternFly Tree View Web Component Implementation

Follow patterns established in ~/Developer/cem/serve/middleware/routes/templates/elements/

### React to Web Component API Translation

**PatternFly React TreeView props → Web Component attributes:**

| React Prop | Web Component Attribute | Values |
|------------|------------------------|--------|
| `variant` | `variant` | "default" \| "compact" \| "compact-no-background" |
| `hasGuides` | `guides` | boolean |
| `isMultiSelectable` | `multi-select` | boolean |
| `hasCheckboxes` | `checkboxes` | boolean |
| `hasSelectableNodes` | `selectable` | boolean |
| `allExpanded` | `all-expanded` | boolean |

**React events → Web Component custom events:**

| React Handler | Custom Event | Event Class Fields |
|---------------|--------------|-------------------|
| `onSelect(event, item, parent)` | `select` | `item`, `parent` |
| `onCheck(event, item, parent)` | `check` | `item`, `parent`, `checked` |
| `onExpand(event, item, parent)` | `expand` | `item`, `parent` |
| `onCollapse(event, item, parent)` | `collapse` | `item`, `parent` |

#### A11y
Use Element Internals to set default role semantics for views, children. set roles on shadow containers when an element can be both item (host) and list (shadow). consider using wccg context protocol to improve role a11y depending on tree structure, or just reflect the role attr.
**Custom event classes (extend Event with class fields, NOT CustomEvent with details):**

```javascript
class PfTreeViewSelectEvent extends Event {
  item;
  parent;

  constructor(item, parent) {
    super('select', { bubbles: true, composed: true });
    this.item = item;
    this.parent = parent;
  }
}

class PfTreeViewCheckEvent extends Event {
  item;
  parent;
  checked;

  constructor(item, parent, checked) {
    super('check', { bubbles: true, composed: true });
    this.item = item;
    this.parent = parent;
    this.checked = checked;
  }
}

class PfTreeViewExpandEvent extends Event {
  item;
  parent;

  constructor(item, parent) {
    super('expand', { bubbles: true, composed: true });
    this.item = item;
    this.parent = parent;
  }
}

class PfTreeViewCollapseEvent extends Event {
  item;
  parent;

  constructor(item, parent) {
    super('collapse', { bubbles: true, composed: true });
    this.item = item;
    this.parent = parent;
  }
}
```

**Data population methods:**

**PREFER: Declarative server-side rendering** - Render tree structure in Go templates:

```html
<pf-v6-toolbar>
  <pf-v6-toolbar-item>
    <pf-v6-text-input-group type="search" placeholder="Search manifest...">
      <svg slot="icon"><!-- search icon --></svg>
    </pf-v6-text-input-group>
  </pf-v6-toolbar-item>
</pf-v6-toolbar>

<hr>

<pf-v6-tree-view variant="compact" guides>
  {{range .Modules}}
  <pf-v6-tree-view-list-item name="{{.Path}}">
    <svg slot="icon"><!-- module icon --></svg>
    <span slot="badge">{{len .Declarations}}</span>

    {{range .Declarations}}
    <pf-v6-tree-view-list-item name="{{.TagName}}">
      <svg slot="icon"><!-- element icon --></svg>
      <span slot="badge">{{.APICount}}</span>

      {{if .Attributes}}
      <pf-v6-tree-view-list-item name="Attributes">
        <svg slot="icon"><!-- folder icon --></svg>
        <span slot="badge">{{len .Attributes}}</span>

        {{range .Attributes}}
        <pf-v6-tree-view-list-item name="{{.Name}}">
          <svg slot="icon"><!-- attr icon --></svg>
          {{if .Deprecated}}<pf-v6-label compact status="warning">deprecated</pf-v6-label>{{end}}
        </pf-v6-tree-view-list-item>
        {{end}}
      </pf-v6-tree-view-list-item>
      {{end}}
    </pf-v6-tree-view-list-item>
    {{end}}
  </pf-v6-tree-view-list-item>
  {{end}}
</pf-v6-tree-view>
```

**Fallback: Imperative (data property)** - Only if declarative is not feasible:
```javascript
treeView.data = manifestTreeBuilder.build(manifest);
```

**Rationale for server-side:**
- Manifest is available at render time
- No client-side JSON parsing overhead
- Better initial render performance
- SEO-friendly (though not critical for dev tool)
- Simpler client-side code

### PatternFly CSS Structure

**Classes from PatternFly (compiled CSS from patternfly.org):**
- `.pf-v-c-tree-view` - Root container
- `.pf-v-c-tree-view__list` - List container (ul)
- `.pf-v-c-tree-view__list-item` - List item (li)
- `.pf-v-c-tree-view__node` - Node wrapper
- `.pf-v-c-tree-view__node-toggle` - Expand/collapse button
- `.pf-v-c-tree-view__node-toggle-icon` - Chevron icon
- `.pf-v-c-tree-view__node-container` - Content container
- `.pf-v-c-tree-view__node-icon` - Node icon
- `.pf-v-c-tree-view__node-text` - Node text/name
- `.pf-v-c-tree-view__node-count` - Badge/count
- `.pf-v-c-tree-view__action` - Action slot

**Modifiers:**
- `.pf-m-compact` - Compact variant
- `.pf-m-no-background` - No background for compact
- `.pf-m-guides` - Visual guides between nodes
- `.pf-m-expanded` - Expanded state
- `.pf-m-current` - Selected/current item

### HTML Structure (from React-rendered output)

```html
<div class="pf-v6-c-toolbar__content" style="padding: 0px;">
  <div class="pf-v6-c-toolbar__content-section">
    <div class="pf-v6-c-toolbar__item">
      <div class="pf-v6-c-tree-view__search">
        <div class="pf-v6-c-form-control pf-m-icon">
          <input id="input-search" name="search-input" aria-label="Search input example" type="search">
          <div class="pf-v6-c-form-control__utilities">
            <div class="pf-v6-c-form-control__icon">
              <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                <path d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<hr class="pf-v6-c-divider">
<ul class="pf-v6-c-tree-view__list" role="tree" aria-multiselectable="false" aria-label="Tree View with icons example">
  <li id="example4-AppLaunch" class="pf-v6-c-tree-view__list-item pf-m-expanded" aria-expanded="true" role="treeitem" tabindex="-1" aria-selected="false">
    <div class="pf-v6-c-tree-view__content">
      <button class="pf-v6-c-tree-view__node" type="button" tabindex="0">
        <span class="pf-v6-c-tree-view__node-container">
          <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
            <span class="pf-v6-c-tree-view__node-toggle-icon">
              <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
              </svg>
            </span>
          </span>
          <span class="pf-v6-c-tree-view__node-icon">
            <svg class="pf-v6-svg" viewBox="0 0 576 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
              <path d="M572.694 292.093L500.27 416.248A63.997 63.997 0 0 1 444.989 448H45.025c-18.523 0-30.064-20.093-20.731-36.093l72.424-124.155A64 64 0 0 1 152 256h399.964c18.523 0 30.064 20.093 20.73 36.093zM152 224h328v-48c0-26.51-21.49-48-48-48H272l-64-64H48C21.49 64 0 85.49 0 112v278.046l69.077-118.418C86.214 242.25 117.989 224 152 224z"></path>
            </svg>
          </span>
          <span class="pf-v6-c-tree-view__node-text">Application launcher</span>
        </span>
      </button>
    </div>
    <ul class="pf-v6-c-tree-view__list" role="group">
      <li id="example4-App1" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
        <div class="pf-v6-c-tree-view__content">
          <button class="pf-v6-c-tree-view__node" type="button">
            <span class="pf-v6-c-tree-view__node-container">
              <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
                <span class="pf-v6-c-tree-view__node-toggle-icon">
                  <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                    <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
                  </svg>
                </span>
              </span>
              <span class="pf-v6-c-tree-view__node-icon">
                <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                  <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z"></path>
                </svg>
              </span>
              <span class="pf-v6-c-tree-view__node-text">
                Application 1</span>
              <span class="pf-v6-c-tree-view__node-count">
                <span class="pf-v6-c-badge pf-m-read">2</span>
              </span>
            </span>
          </button>
        </div>
        <ul class="pf-v6-c-tree-view__list" role="group" inert="">
          <li id="example4-App1Settings" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
            <div class="pf-v6-c-tree-view__content">
              <button class="pf-v6-c-tree-view__node" type="button">
                <span class="pf-v6-c-tree-view__node-container">
                  <span class="pf-v6-c-tree-view__node-icon">
                    <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                      <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z"></path>
                    </svg>
                  </span>
                  <span class="pf-v6-c-tree-view__node-text">Settings</span>
                </span>
              </button>
            </div>
          </li>
          <li id="example4-App1Current" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
            <div class="pf-v6-c-tree-view__content">
              <button class="pf-v6-c-tree-view__node" type="button">
                <span class="pf-v6-c-tree-view__node-container">
                  <span class="pf-v6-c-tree-view__node-icon">
                    <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                      <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z"></path>
                    </svg>
                  </span>
                  <span class="pf-v6-c-tree-view__node-text">Current</span>
                </span>
              </button>
            </div>
          </li>
        </ul>
      </li>
      <li id="example4-App2" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
        <div class="pf-v6-c-tree-view__content">
          <button class="pf-v6-c-tree-view__node" type="button">
            <span class="pf-v6-c-tree-view__node-container">
              <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
                <span class="pf-v6-c-tree-view__node-toggle-icon">
                  <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                    <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
                  </svg>
                </span>
              </span>
              <span class="pf-v6-c-tree-view__node-icon">
                <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                  <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z"></path>
                </svg>
              </span>
              <span class="pf-v6-c-tree-view__node-text">Application 2</span>
            </span>
          </button>
        </div>
        <ul class="pf-v6-c-tree-view__list" role="group" inert="">
          <li id="example4-App2Settings" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
            <div class="pf-v6-c-tree-view__content">
              <button class="pf-v6-c-tree-view__node" type="button">
                <span class="pf-v6-c-tree-view__node-container">
                  <span class="pf-v6-c-tree-view__node-icon">
                    <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                      <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z"></path>
                    </svg>
                  </span>
                  <span class="pf-v6-c-tree-view__node-text">Settings</span>
                </span>
              </button>
            </div>
          </li>
          <li id="example4-App2Loader" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
            <div class="pf-v6-c-tree-view__content">
              <button class="pf-v6-c-tree-view__node" type="button">
                <span class="pf-v6-c-tree-view__node-container">
                  <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
                    <span class="pf-v6-c-tree-view__node-toggle-icon">
                      <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                        <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
                      </svg>
                    </span>
                  </span>
                  <span class="pf-v6-c-tree-view__node-icon">
                    <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                      <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                      </path>
                    </svg>
                  </span>
                  <span class="pf-v6-c-tree-view__node-text">
                    Loader</span>
                </span>
              </button>
            </div>
            <ul class="pf-v6-c-tree-view__list" role="group" inert="">
              <li id="example4-LoadApp1" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
                <div class="pf-v6-c-tree-view__content">
                  <button class="pf-v6-c-tree-view__node" type="button">
                    <span class="pf-v6-c-tree-view__node-container">
                      <span class="pf-v6-c-tree-view__node-icon">
                        <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                          <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z"></path>
                        </svg>
                      </span>
                      <span class="pf-v6-c-tree-view__node-text">Loading App 1</span>
                    </span>
                  </button>
                </div>
              </li>
              <li id="example4-LoadApp2" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
                <div class="pf-v6-c-tree-view__content">
                  <button class="pf-v6-c-tree-view__node" type="button">
                    <span class="pf-v6-c-tree-view__node-container">
                      <span class="pf-v6-c-tree-view__node-icon">
                        <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                          <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                          </path>
                        </svg>
                      </span>
                      <span class="pf-v6-c-tree-view__node-text">
                        Loading App 2</span>
                    </span>
                  </button>
                </div>
              </li>
              <li id="example4-LoadApp3" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
                <div class="pf-v6-c-tree-view__content">
                  <button class="pf-v6-c-tree-view__node" type="button">
                    <span class="pf-v6-c-tree-view__node-container">
                      <span class="pf-v6-c-tree-view__node-icon">
                        <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                          <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                          </path>
                        </svg>
                      </span>
                      <span class="pf-v6-c-tree-view__node-text">
                        Loading App 3</span>
                    </span>
                  </button>
                </div>
              </li>
            </ul>
          </li>
        </ul>
      </li>
    </ul>
  </li>
  <li id="example4-Cost" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
    <div class="pf-v6-c-tree-view__content">
      <button class="pf-v6-c-tree-view__node" type="button">
        <span class="pf-v6-c-tree-view__node-container">
          <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
            <span class="pf-v6-c-tree-view__node-toggle-icon">
              <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z">
                </path>
              </svg>
            </span>
          </span>
          <span class="pf-v6-c-tree-view__node-icon">
            <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
              <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
              </path>
            </svg>
          </span>
          <span class="pf-v6-c-tree-view__node-text">
            Cost management</span>
        </span></button>
    </div>
    <ul class="pf-v6-c-tree-view__list" role="group" inert=""><li id="example4-App3" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
      <div class="pf-v6-c-tree-view__content">
        <button class="pf-v6-c-tree-view__node" type="button">
          <span class="pf-v6-c-tree-view__node-container">
            <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
              <span class="pf-v6-c-tree-view__node-toggle-icon">
                <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                  <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
                </svg>
              </span>
            </span>
            <span class="pf-v6-c-tree-view__node-icon">
              <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                </path>
              </svg>
            </span>
            <span class="pf-v6-c-tree-view__node-text">
              Application 3</span>
          </span>
        </button>
      </div>
      <ul class="pf-v6-c-tree-view__list" role="group" inert="">
        <li id="example4-App3Settings" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
          <div class="pf-v6-c-tree-view__content">
            <button class="pf-v6-c-tree-view__node" type="button">
              <span class="pf-v6-c-tree-view__node-container">
                <span class="pf-v6-c-tree-view__node-icon">
                  <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                    <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                    </path>
                  </svg>
                </span><span class="pf-v6-c-tree-view__node-text">
                  Settings</span>
              </span>
            </button>
          </div>
        </li>
        <li id="example4-App3Current" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
          <div class="pf-v6-c-tree-view__content">
            <button class="pf-v6-c-tree-view__node" type="button">
              <span class="pf-v6-c-tree-view__node-container">
                <span class="pf-v6-c-tree-view__node-icon">
                  <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                    <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z"></path>
                  </svg>
                </span>
                
                <span class="pf-v6-c-tree-view__node-text">
                  Current</span>
              </span>
            </button>
          </div>
        </li>
      </ul>
    </li>
    </ul>
  </li>
  <li id="example4-Sources" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
    <div class="pf-v6-c-tree-view__content">
      <button class="pf-v6-c-tree-view__node" type="button">
        <span class="pf-v6-c-tree-view__node-container">
          <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
            <span class="pf-v6-c-tree-view__node-toggle-icon">
              <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z">
                </path>
              </svg>
            </span>
          </span>
          <span class="pf-v6-c-tree-view__node-icon">
            <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
              <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
              </path>
            </svg>
          </span>
          <span class="pf-v6-c-tree-view__node-text">
            Sources</span>
        </span>
      </button>
    </div>
    <ul class="pf-v6-c-tree-view__list" role="group" inert="">
      <li id="example4-App4" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
        <div class="pf-v6-c-tree-view__content">
          <button class="pf-v6-c-tree-view__node" type="button">
            <span class="pf-v6-c-tree-view__node-container">
              <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
                <span class="pf-v6-c-tree-view__node-toggle-icon">
                  <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                    <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z">
                    </path>
                  </svg>
                </span>
              </span>
              <span class="pf-v6-c-tree-view__node-icon">
                <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                  <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                  </path>
                </svg>
              </span>
              <span class="pf-v6-c-tree-view__node-text">
                Application 4</span>
            </span>
          </button>
        </div>
        <ul class="pf-v6-c-tree-view__list" role="group" inert="">
          <li id="example4-App4Settings" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
            <div class="pf-v6-c-tree-view__content">
              <button class="pf-v6-c-tree-view__node" type="button">
                <span class="pf-v6-c-tree-view__node-container">
                  <span class="pf-v6-c-tree-view__node-icon">
                    <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                      <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                      </path>
                    </svg>
                  </span>
                  <span class="pf-v6-c-tree-view__node-text">
                    Settings</span>
                </span>
              </button>
            </div>
          </li>
        </ul>
      </li>
    </ul>
  </li>
  <li id="example4-Long" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
    <div class="pf-v6-c-tree-view__content">
      <button class="pf-v6-c-tree-view__node" type="button">
        <span class="pf-v6-c-tree-view__node-container">
          <span class="pf-v6-c-tree-view__node-toggle" tabindex="-1">
            <span class="pf-v6-c-tree-view__node-toggle-icon">
              <svg class="pf-v6-svg" viewBox="0 0 256 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z">
                </path>
              </svg>
            </span>
          </span>
          <span class="pf-v6-c-tree-view__node-icon">
            <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
              <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
              </path>
            </svg>
          </span>
          <span class="pf-v6-c-tree-view__node-text">
            Really really really long folder name that overflows the container it is in</span>
        </span>
      </button>
    </div>
    <ul class="pf-v6-c-tree-view__list" role="group" inert="">
      <li id="example4-App5" class="pf-v6-c-tree-view__list-item" aria-expanded="false" role="treeitem" tabindex="-1" aria-selected="false">
        <div class="pf-v6-c-tree-view__content">
          <button class="pf-v6-c-tree-view__node" type="button">
            <span class="pf-v6-c-tree-view__node-container">
              <span class="pf-v6-c-tree-view__node-icon">
                <svg class="pf-v6-svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" role="img" width="1em" height="1em">
                  <path d="M464 128H272l-64-64H48C21.49 64 0 85.49 0 112v288c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48z">
                  </path>
                </svg>
              </span>
              <span class="pf-v6-c-tree-view__node-text">
                Application 5</span>
            </span>
          </button>
        </div>
      </li>
    </ul>
  </li>
</ul>
```

### Compiled SASS

```css
.pf-v6-c-tree-view {
  --pf-v6-c-tree-view__node--indent--base: calc(var(--pf-t--global--spacer--md) * 2 + var(--pf-v6-c-tree-view__node-toggle-icon--MinWidth));
  --pf-v6-c-tree-view__node--nested-indent--base: calc(var(--pf-v6-c-tree-view__node--indent--base) - var(--pf-t--global--spacer--md));
  --pf-v6-c-tree-view__content--BorderRadius: var(--pf-t--global--border--radius--small);
  --pf-v6-c-tree-view__node--PaddingBlockStart: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node--PaddingInlineEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node--PaddingBlockEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node--PaddingInlineStart: var(--pf-v6-c-tree-view__node--indent--base);
  --pf-v6-c-tree-view__node--Color: var(--pf-t--global--text--color--subtle);
  --pf-v6-c-tree-view__node--BackgroundColor: transparent;
  --pf-v6-c-tree-view__node--BorderRadius: var(--pf-v6-c-tree-view__content--BorderRadius);
  --pf-v6-c-tree-view__node--BorderColor: var(--pf-t--global--border--color--high-contrast);
  --pf-v6-c-tree-view__node--BorderWidth: var(--pf-t--global--border--width--action--plain--default);
  --pf-v6-c-tree-view__node--hover--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);
  --pf-v6-c-tree-view__node--m-current--BorderWidth: var(--pf-t--global--border--width--action--plain--clicked);
  --pf-v6-c-tree-view__node--m-current--Color: var(--pf-t--global--text--color--regular);
  --pf-v6-c-tree-view__node--m-current--BackgroundColor: var(--pf-t--global--background--color--primary--clicked);
  --pf-v6-c-tree-view__node--hover--BackgroundColor: var(--pf-t--global--background--color--primary--hover);
  --pf-v6-c-tree-view__node-container--Display: contents;
  --pf-v6-c-tree-view__node-content--RowGap: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node-content--Overflow: visible;
  --pf-v6-c-tree-view__list--TransitionDuration--expand--slide: 0s;
  --pf-v6-c-tree-view__list--TransitionDuration--expand--fade: var(--pf-t--global--motion--duration--fade--default);
  --pf-v6-c-tree-view__list--TransitionDuration--collapse--slide: 0s;
  --pf-v6-c-tree-view__list--TransitionDuration--collapse--fade: var(--pf-t--global--motion--duration--fade--short);
  --pf-v6-c-tree-view__list--TransitionDuration--slide: var(--pf-v6-c-tree-view__list--TransitionDuration--collapse--slide);
  --pf-v6-c-tree-view__list--TransitionDuration--fade: var(--pf-v6-c-tree-view__list--TransitionDuration--collapse--fade);
  --pf-v6-c-tree-view__list--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);
  --pf-v6-c-tree-view__list--Opacity: 0;
  --pf-v6-c-tree-view--m-expanded__list--Opacity: 1;
  --pf-v6-c-tree-view__list--TranslateY: 0;
  --pf-v6-c-tree-view--m-expanded__list--TranslateY: 0;
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetBlockStart: var(--pf-v6-c-tree-view__node--PaddingBlockStart);
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--TranslateX: -100%;
  --pf-v6-c-tree-view__node-toggle--Position: absolute;
  --pf-v6-c-tree-view__node-toggle--Color--base: var(--pf-t--global--icon--color--subtle);
  --pf-v6-c-tree-view__node-toggle--Color: var(--pf-v6-c-tree-view__node-toggle--Color--base);
  --pf-v6-c-tree-view__node-toggle--BackgroundColor: transparent;
  --pf-v6-c-tree-view__list-item--m-expanded__node-toggle--Color: var(--pf-t--global--icon--color--regular);
  --pf-v6-c-tree-view__node-toggle-icon--MinWidth: var(--pf-t--global--icon--size--font--body--default);
  --pf-v6-c-tree-view__node-toggle--PaddingBlockStart: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node-toggle--PaddingInlineEnd: var(--pf-t--global--spacer--md);
  --pf-v6-c-tree-view__node-toggle--PaddingBlockEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node-toggle--PaddingInlineStart: var(--pf-t--global--spacer--md);
  --pf-v6-c-tree-view__node-toggle--MarginBlockStart: calc(var(--pf-v6-c-tree-view__node-toggle--PaddingBlockStart) * -1);
  --pf-v6-c-tree-view__node-toggle--MarginBlockEnd: calc(var(--pf-v6-c-tree-view__node-toggle--PaddingBlockStart) * -1);
  --pf-v6-c-tree-view__node-toggle-icon--TransitionDuration: var(--pf-t--global--motion--duration--icon--default);
  --pf-v6-c-tree-view__node-toggle-icon--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);
  --pf-v6-c-tree-view__node-check--MarginInlineEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node-count--MarginInlineStart: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__search--PaddingBlockStart: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__search--PaddingInlineEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__search--PaddingBlockEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__search--PaddingInlineStart: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node-icon--PaddingInlineEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view__node-icon--Color: var(--pf-t--global--icon--color--subtle);
  --pf-v6-c-tree-view__node-toggle-icon--base--Rotate: 0;
  --pf-v6-c-tree-view__node-toggle-icon--Rotate: var(--pf-v6-c-tree-view__node-toggle-icon--base--Rotate);
  --pf-v6-c-tree-view__list-item--m-expanded__node-toggle-icon--Rotate: 90deg;
  --pf-v6-c-tree-view__node-text--max-lines: 1;
  --pf-v6-c-tree-view__node-title--FontWeight: var(--pf-t--global--font--weight--body--bold);
  --pf-v6-c-tree-view__action--MarginInlineEnd: var(--pf-t--global--spacer--md);
  --pf-v6-c-tree-view--m-guides--guide--InsetInlineStart: var(--pf-v6-c-tree-view--m-guides--guide-left--base);
  --pf-v6-c-tree-view--m-guides--guide-color--base: var(--pf-t--global--border--color--default);
  --pf-v6-c-tree-view--m-guides--guide-width--base: var(--pf-t--global--border--width--divider--default);
  --pf-v6-c-tree-view--m-guides--guide-left--base: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides__list-node--guide-width--base));
  --pf-v6-c-tree-view--m-guides--guide-left--base--offset: calc(var(--pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart) + var(--pf-v6-c-tree-view__node-toggle-icon--MinWidth) / 2);
  --pf-v6-c-tree-view--m-guides__list-node--guide-width--base: var(--pf-t--global--spacer--lg);
  --pf-v6-c-tree-view--m-guides__list-item--before--InsetBlockStart: 0;
  --pf-v6-c-tree-view--m-guides__list-item--before--Width: var(--pf-v6-c-tree-view--m-guides--guide-width--base);
  --pf-v6-c-tree-view--m-guides__list-item--before--Height: 100%;
  --pf-v6-c-tree-view--m-guides__list-item--before--BackgroundColor: var(--pf-v6-c-tree-view--m-guides--guide-color--base);
  --pf-v6-c-tree-view--m-guides__list-item--last-child--before--InsetBlockStart: var(--pf-v6-c-tree-view--m-guides__node--before--InsetBlockStart);
  --pf-v6-c-tree-view--m-guides__list-item--last-child--before--Height: var(--pf-v6-c-tree-view--m-guides__list-item--last-child--before--InsetBlockStart);
  --pf-v6-c-tree-view--m-guides__list-item--ZIndex: var(--pf-t--global--z-index--xs);
  --pf-v6-c-tree-view--m-guides__node--before--Width: 1rem;
  --pf-v6-c-tree-view--m-guides__node--before--Height: var(--pf-v6-c-tree-view--m-guides--guide-width--base);
  --pf-v6-c-tree-view--m-guides__node--before--InsetBlockStart: 1.125rem;
  --pf-v6-c-tree-view--m-guides__node--before--BackgroundColor: var(--pf-v6-c-tree-view--m-guides--guide-color--base);
  --pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset: var(--pf-t--global--spacer--md);
  --pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact__node--indent--base: var(--pf-v6-c-tree-view__node--indent--base);
  --pf-v6-c-tree-view--m-compact__node--nested-indent--base: var(--pf-t--global--spacer--lg);
  --pf-v6-c-tree-view--m-compact--border--InsetInlineStart: var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart);
  --pf-v6-c-tree-view--m-compact__node--Color: var(--pf-t--global--text--color--regular);
  --pf-v6-c-tree-view--m-compact__node--PaddingBlockStart: 0;
  --pf-v6-c-tree-view--m-compact__node--PaddingBlockEnd: 0;
  --pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockStart: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockEnd: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view--m-compact__list-item__list-item__node-toggle--InsetBlockStart: calc(var(--pf-v6-c-tree-view--m-compact__node-container--PaddingBlockStart));
  --pf-v6-c-tree-view--m-compact__list-item--BorderBlockEndWidth: var(--pf-t--global--border--width--divider--default);
  --pf-v6-c-tree-view--m-compact__list-item--BorderBlockEndColor: var(--pf-t--global--border--color--default);
  --pf-v6-c-tree-view--m-compact__list-item--before--InsetBlockStart: 0;
  --pf-v6-c-tree-view--m-compact__list-item--last-child--before--Height: var(--pf-v6-c-tree-view--m-compact__node--before--InsetBlockStart);
  --pf-v6-c-tree-view--m-compact__list-item--nested--before--InsetBlockStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockStart) * -1);
  --pf-v6-c-tree-view--m-compact__list-item--nested--last-child--before--Height: calc(var(--pf-v6-c-tree-view--m-compact__node--before--InsetBlockStart) + var(--pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockStart));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: var(--pf-v6-c-tree-view--m-compact__node--indent--base);
  --pf-v6-c-tree-view--m-compact__node--before--InsetBlockStart: calc(var(--pf-v6-c-tree-view--m-compact__node-container--PaddingBlockStart) + var(--pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockStart) + 0.25rem);
  --pf-v6-c-tree-view--m-compact__node--level-2--PaddingInlineStart: var(--pf-v6-c-tree-view--m-compact__node--indent--base);
  --pf-v6-c-tree-view--m-compact__node-toggle--nested--MarginInlineEnd: calc(var(--pf-v6-c-tree-view__node-toggle--PaddingInlineStart) * -.5);
  --pf-v6-c-tree-view--m-compact__node-toggle--nested--MarginInlineStart: calc(var(--pf-v6-c-tree-view__node-toggle--PaddingInlineStart) * -1.5);
  --pf-v6-c-tree-view--m-compact__node-container--Display: flex;
  --pf-v6-c-tree-view--m-compact__node-container--PaddingBlockEnd--base: var(--pf-t--global--spacer--lg);
  --pf-v6-c-tree-view--m-compact__node-container--PaddingBlockStart: var(--pf-t--global--spacer--lg);
  --pf-v6-c-tree-view--m-compact__node-container--PaddingInlineEnd: var(--pf-t--global--spacer--lg);
  --pf-v6-c-tree-view--m-compact__node-container--PaddingBlockEnd: var(--pf-v6-c-tree-view--m-compact__node-container--PaddingBlockEnd--base);
  --pf-v6-c-tree-view--m-compact__node-container--PaddingInlineStart: var(--pf-t--global--spacer--xs);
  --pf-v6-c-tree-view--m-compact__node-container--BorderRadius: var(--pf-t--global--border--radius--small);
  --pf-v6-c-tree-view--m-compact__node-container--nested--PaddingBlockStart: var(--pf-t--global--spacer--md);
  --pf-v6-c-tree-view--m-compact__node-container--nested--PaddingInlineEnd: var(--pf-t--global--spacer--lg);
  --pf-v6-c-tree-view--m-compact__node-container--nested--PaddingBlockEnd: var(--pf-t--global--spacer--md);
  --pf-v6-c-tree-view--m-compact__node-container--nested--PaddingInlineStart: var(--pf-t--global--spacer--lg);
  --pf-v6-c-tree-view--m-compact__node-container--nested--BackgroundColor: var(--pf-t--global--background--color--secondary--default);
  --pf-v6-c-tree-view--m-compact__list-item--m-expanded__node-container--PaddingBlockEnd: 0;
  --pf-v6-c-tree-view--m-no-background__node-container--BackgroundColor: transparent;
  --pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset: var(--pf-t--global--spacer--sm);
  --pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base: var(--pf-v6-c-tree-view__node--indent--base);
  --pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base: var(--pf-t--global--spacer--2xl);
  --pf-v6-c-tree-view--m-compact--m-no-background__node--nested--PaddingBlockStart: 0;
  --pf-v6-c-tree-view--m-compact--m-no-background__node--nested--PaddingBlockEnd: 0;
  --pf-v6-c-tree-view--m-compact--m-no-background__node--before--InsetBlockStart: calc(var(--pf-v6-c-tree-view--m-compact__node-container--nested--PaddingBlockStart) + var(--pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockStart) + 0.25rem);
}
@media screen and (prefers-reduced-motion: no-preference) {
  .pf-v6-c-tree-view {
    --pf-v6-c-tree-view__list--TransitionDuration--expand--slide: var(--pf-t--global--motion--duration--fade--default);
    --pf-v6-c-tree-view__list--TransitionDuration--collapse--slide: var(--pf-t--global--motion--duration--fade--short);
    --pf-v6-c-tree-view__list--TranslateY: -.5rem;
  }
}

.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item, .pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item {
  position: relative;
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item::before,
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node::before, .pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item::before,
.pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node::before {
  position: absolute;
  inset-inline-start: var(--pf-v6-c-tree-view--m-guides--guide--InsetInlineStart);
  content: "";
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item::before, .pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item::before {
  inset-block-start: var(--pf-v6-c-tree-view--m-guides__list-item--before--InsetBlockStart);
  z-index: var(--pf-v6-c-tree-view--m-guides__list-item--ZIndex);
  width: var(--pf-v6-c-tree-view--m-guides__list-item--before--Width);
  height: var(--pf-v6-c-tree-view--m-guides__list-item--before--Height);
  background-color: var(--pf-v6-c-tree-view--m-guides__list-item--before--BackgroundColor);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node::before, .pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node::before {
  inset-block-start: var(--pf-v6-c-tree-view--m-guides__node--before--InsetBlockStart);
  width: var(--pf-v6-c-tree-view--m-guides__node--before--Width);
  height: var(--pf-v6-c-tree-view--m-guides__node--before--Height);
  background-color: var(--pf-v6-c-tree-view--m-guides__node--before--BackgroundColor);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item:last-child::before, .pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item:last-child::before {
  height: var(--pf-v6-c-tree-view--m-guides__list-item--last-child--before--Height);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item, .pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view--m-guides--guide--InsetInlineStart: var(--pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item:last-child, .pf-v6-c-tree-view.pf-m-guides .pf-v6-c-tree-view__list-item:last-child {
  --pf-v6-c-tree-view--m-compact__list-item--BorderBlockEndWidth: 0;
}
.pf-v6-c-tree-view.pf-m-compact {
  --pf-v6-c-tree-view__node--Color: var(--pf-v6-c-tree-view--m-compact__node--Color);
  --pf-v6-c-tree-view__node--PaddingBlockStart: var(--pf-v6-c-tree-view--m-compact__node--PaddingBlockStart);
  --pf-v6-c-tree-view__node--PaddingBlockEnd: var(--pf-v6-c-tree-view--m-compact__node--PaddingBlockEnd);
  --pf-v6-c-tree-view__node-container--Display: var(--pf-v6-c-tree-view--m-compact__node-container--Display);
  --pf-v6-c-tree-view__node--hover--BackgroundColor: transparent;
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetBlockStart: var(--pf-v6-c-tree-view--m-compact__list-item__list-item__node-toggle--InsetBlockStart);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item {
  border-block-end: var(--pf-v6-c-tree-view--m-compact__list-item--BorderBlockEndWidth) solid var(--pf-v6-c-tree-view--m-compact__list-item--BorderBlockEndColor);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item.pf-m-expanded {
  --pf-v6-c-tree-view--m-compact__node-container--PaddingBlockEnd: var(--pf-v6-c-tree-view--m-compact__list-item--m-expanded__node-container--PaddingBlockEnd);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__node--PaddingBlockStart: var(--pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockStart);
  --pf-v6-c-tree-view__node--PaddingBlockEnd: var(--pf-v6-c-tree-view--m-compact__node--nested--PaddingBlockEnd);
  --pf-v6-c-tree-view__node-toggle--Position: static;
  --pf-v6-c-tree-view__node--PaddingInlineStart: var(--pf-v6-c-tree-view--m-compact__node--level-2--PaddingInlineStart);
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--TranslateX: 0;
  --pf-v6-c-tree-view--m-compact__list-item--BorderBlockEndWidth: 0;
  --pf-v6-c-tree-view--m-compact__node-container--PaddingBlockEnd: var(--pf-v6-c-tree-view--m-compact__node-container--PaddingBlockEnd--base);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item::before,
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node::before {
  inset-inline-start: var(--pf-v6-c-tree-view--m-compact--border--InsetInlineStart);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item::before {
  inset-block-start: var(--pf-v6-c-tree-view--m-compact__list-item--before--InsetBlockStart);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node::before {
  inset-block-start: var(--pf-v6-c-tree-view--m-compact__node--before--InsetBlockStart);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item:last-child::before {
  height: var(--pf-v6-c-tree-view--m-compact__list-item--last-child--before--Height);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__node--PaddingInlineStart: var(--pf-v6-c-tree-view--m-compact__node--PaddingInlineStart);
  --pf-v6-c-tree-view--m-compact--border--InsetInlineStart: var(--pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart);
  --pf-v6-c-tree-view--m-compact__list-item--before--InsetBlockStart: var(--pf-v6-c-tree-view--m-compact__list-item--nested--before--InsetBlockStart);
  --pf-v6-c-tree-view--m-compact__list-item--last-child--before--Height: var(--pf-v6-c-tree-view--m-compact__list-item--nested--last-child--before--Height);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node-container {
  padding-block-start: var(--pf-v6-c-tree-view--m-compact__node-container--nested--PaddingBlockStart);
  padding-block-end: var(--pf-v6-c-tree-view--m-compact__node-container--nested--PaddingBlockEnd);
  padding-inline-start: var(--pf-v6-c-tree-view--m-compact__node-container--nested--PaddingInlineStart);
  padding-inline-end: var(--pf-v6-c-tree-view--m-compact__node-container--nested--PaddingInlineEnd);
  background-color: var(--pf-v6-c-tree-view--m-compact__node-container--nested--BackgroundColor);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__node-toggle {
  margin-inline-start: var(--pf-v6-c-tree-view--m-compact__node-toggle--nested--MarginInlineStart);
  margin-inline-end: var(--pf-v6-c-tree-view--m-compact__node-toggle--nested--MarginInlineEnd);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__node-container {
  padding-block-start: var(--pf-v6-c-tree-view--m-compact__node-container--PaddingBlockStart);
  padding-block-end: var(--pf-v6-c-tree-view--m-compact__node-container--PaddingBlockEnd);
  padding-inline-start: var(--pf-v6-c-tree-view--m-compact__node-container--PaddingInlineStart);
  padding-inline-end: var(--pf-v6-c-tree-view--m-compact__node-container--PaddingInlineEnd);
}
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item:not([aria-expanded]) > .pf-v6-c-tree-view__content > .pf-v6-c-tree-view__node,
.pf-v6-c-tree-view.pf-m-compact .pf-v6-c-tree-view__list-item:not([aria-expanded]) > .pf-v6-c-tree-view__content > .pf-v6-c-tree-view__node > .pf-v6-c-tree-view__node-container {
  cursor: default;
}
.pf-v6-c-tree-view.pf-m-compact.pf-m-no-background {
  --pf-v6-c-tree-view--m-compact__node--before--InsetBlockStart: var(--pf-v6-c-tree-view--m-compact--m-no-background__node--before--InsetBlockStart);
}
.pf-v6-c-tree-view.pf-m-compact.pf-m-no-background .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__node--PaddingBlockStart: var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested--PaddingBlockStart);
  --pf-v6-c-tree-view__node--PaddingBlockEnd: var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested--PaddingBlockEnd);
}
.pf-v6-c-tree-view.pf-m-compact.pf-m-no-background .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view--m-compact--border--InsetInlineStart: var(--pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: var(--pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart);
}
.pf-v6-c-tree-view.pf-m-no-background {
  --pf-v6-c-tree-view--m-compact__node-container--nested--BackgroundColor: var(--pf-v6-c-tree-view--m-no-background__node-container--BackgroundColor);
}

.pf-v6-c-tree-view__node-toggle-icon {
  display: inline-block;
  min-width: var(--pf-v6-c-tree-view__node-toggle-icon--MinWidth);
  text-align: center;
  transition: transform var(--pf-v6-c-tree-view__node-toggle-icon--TransitionDuration) var(--pf-v6-c-tree-view__node-toggle-icon--TransitionTimingFunction);
  transform: rotate(var(--pf-v6-c-tree-view__node-toggle-icon--Rotate));
}
:where(.pf-v6-m-dir-rtl, [dir=rtl]) .pf-v6-c-tree-view__node-toggle-icon {
  scale: -1 1;
}

.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list {
  max-height: 0;
  visibility: hidden;
  opacity: var(--pf-v6-c-tree-view__list--Opacity);
  transition-delay: 0s, 0s, var(--pf-v6-c-tree-view__list--TransitionDuration--fade), var(--pf-v6-c-tree-view__list--TransitionDuration--fade);
  transition-timing-function: var(--pf-v6-c-tree-view__list--TransitionTimingFunction);
  transition-duration: var(--pf-v6-c-tree-view__list--TransitionDuration--fade), var(--pf-v6-c-tree-view__list--TransitionDuration--slide), 0s, 0s;
  transition-property: opacity, translate, visibility, max-height;
  translate: 0 var(--pf-v6-c-tree-view__list--TranslateY);
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__node-toggle-icon--Rotate: var(--pf-v6-c-tree-view__node-toggle-icon--base--Rotate);
  --pf-v6-c-tree-view__node-toggle--Color: var(--pf-v6-c-tree-view__node-toggle--Color--base);
}
.pf-v6-c-tree-view__list-item.pf-m-expanded {
  --pf-v6-c-tree-view__node-toggle--Color: var(--pf-v6-c-tree-view__list-item--m-expanded__node-toggle--Color);
  --pf-v6-c-tree-view__node-toggle-icon--Rotate: var(--pf-v6-c-tree-view__list-item--m-expanded__node-toggle-icon--Rotate);
}
.pf-v6-c-tree-view__list-item.pf-m-expanded > .pf-v6-c-tree-view__list {
  max-height: 99999px;
  visibility: revert;
  opacity: var(--pf-v6-c-tree-view--m-expanded__list--Opacity);
  transition-delay: 0s;
  transition-duration: var(--pf-v6-c-tree-view__list--TransitionDuration--expand--fade), var(--pf-v6-c-tree-view__list--TransitionDuration--expand--slide), 0s, 0s;
  translate: 0 var(--pf-v6-c-tree-view--m-expanded__list--TranslateY);
}

.pf-v6-c-tree-view__node,
.pf-v6-c-tree-view__node-container {
  flex: 1 1;
  align-items: flex-start;
  min-width: 0;
  text-align: start;
  cursor: pointer;
  border: 0;
}

.pf-v6-c-tree-view__node {
  position: relative;
  display: flex;
  padding-block-start: var(--pf-v6-c-tree-view__node--PaddingBlockStart);
  padding-block-end: var(--pf-v6-c-tree-view__node--PaddingBlockEnd);
  padding-inline-start: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  padding-inline-end: var(--pf-v6-c-tree-view__node--PaddingInlineEnd);
  color: var(--pf-v6-c-tree-view__node--Color);
  background-color: var(--pf-v6-c-tree-view__node--BackgroundColor);
}
.pf-v6-c-tree-view__node::after {
  position: absolute;
  inset: 0;
  pointer-events: none;
  content: "";
  border: var(--pf-v6-c-tree-view__node--BorderWidth) solid var(--pf-v6-c-tree-view__node--BorderColor);
  border-radius: var(--pf-v6-c-tree-view__node--BorderRadius);
}
.pf-v6-c-tree-view__node.pf-m-current {
  --pf-v6-c-tree-view__node--Color: var(--pf-v6-c-tree-view__node--m-current--Color);
  --pf-v6-c-tree-view__node--BorderWidth: var(--pf-v6-c-tree-view__node--m-current--BorderWidth);
}
.pf-v6-c-tree-view__node .pf-v6-c-tree-view__node-count {
  margin-inline-start: var(--pf-v6-c-tree-view__node-count--MarginInlineStart);
}

.pf-v6-c-tree-view__node-container {
  display: var(--pf-v6-c-tree-view__node-container--Display);
  flex-grow: 1;
  border-radius: var(--pf-v6-c-tree-view--m-compact__node-container--BorderRadius);
}

.pf-v6-c-tree-view__node-content {
  display: flex;
  flex-direction: column;
  overflow: var(--pf-v6-c-tree-view__node-content--Overflow);
}
.pf-v6-c-tree-view__node-content > * + * {
  margin-block-start: var(--pf-v6-c-tree-view__node-content--RowGap);
}

.pf-v6-c-tree-view__node-check {
  margin-inline-end: var(--pf-v6-c-tree-view__node-check--MarginInlineEnd);
}

.pf-v6-c-tree-view__node-toggle {
  position: var(--pf-v6-c-tree-view__node-toggle--Position);
  inset-block-start: var(--pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetBlockStart);
  inset-inline-start: var(--pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart);
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding-block-start: var(--pf-v6-c-tree-view__node-toggle--PaddingBlockStart);
  padding-block-end: var(--pf-v6-c-tree-view__node-toggle--PaddingBlockEnd);
  padding-inline-start: var(--pf-v6-c-tree-view__node-toggle--PaddingInlineStart);
  padding-inline-end: var(--pf-v6-c-tree-view__node-toggle--PaddingInlineEnd);
  margin-block-start: var(--pf-v6-c-tree-view__node-toggle--MarginBlockStart);
  margin-block-end: var(--pf-v6-c-tree-view__node-toggle--MarginBlockEnd);
  color: var(--pf-v6-c-tree-view__node-toggle--Color);
  background-color: var(--pf-v6-c-tree-view__node-toggle--BackgroundColor);
  border: 0;
  transform: translateX(var(--pf-v6-c-tree-view__list-item__list-item__node-toggle--TranslateX));
}
:where(.pf-v6-m-dir-rtl, [dir=rtl]) .pf-v6-c-tree-view__node-toggle {
  transform: translateX(calc(var(--pf-v6-c-tree-view__list-item__list-item__node-toggle--TranslateX) * var(--pf-v6-global--inverse--multiplier)));
}

.pf-v6-c-tree-view__node-title.pf-m-truncate,
.pf-v6-c-tree-view__node-text.pf-m-truncate {
  --pf-v6-c-tree-view__node-content--Overflow: hidden;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.pf-v6-c-tree-view.pf-m-truncate .pf-v6-c-tree-view__node-title,
.pf-v6-c-tree-view.pf-m-truncate .pf-v6-c-tree-view__node-text {
  --pf-v6-c-tree-view__node-content--Overflow: hidden;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.pf-v6-c-tree-view__node-text {
  font-weight: inherit;
  color: inherit;
  text-align: start;
  background-color: transparent;
  border: 0;
}
label.pf-v6-c-tree-view__node-text {
  cursor: pointer;
}

.pf-v6-c-tree-view__node-title {
  font-weight: var(--pf-v6-c-tree-view__node-title--FontWeight);
}

.pf-v6-c-tree-view__search {
  padding-block-start: var(--pf-v6-c-tree-view__search--PaddingBlockStart);
  padding-block-end: var(--pf-v6-c-tree-view__search--PaddingBlockEnd);
  padding-inline-start: var(--pf-v6-c-tree-view__search--PaddingInlineStart);
  padding-inline-end: var(--pf-v6-c-tree-view__search--PaddingInlineEnd);
}

.pf-v6-c-tree-view__node-icon {
  padding-inline-end: var(--pf-v6-c-tree-view__node-icon--PaddingInlineEnd);
  color: var(--pf-v6-c-tree-view__node-icon--Color);
}

.pf-v6-c-tree-view__content {
  display: flex;
  align-items: center;
  background-color: var(--pf-v6-c-tree-view__content--BackgroundColor);
  border-radius: var(--pf-v6-c-tree-view__content--BorderRadius);
}

.pf-v6-c-tree-view__content:hover,
.pf-v6-c-tree-view__content:focus-within {
  --pf-v6-c-tree-view__node--BorderWidth: var(--pf-v6-c-tree-view__node--hover--BorderWidth);
  background-color: var(--pf-v6-c-tree-view__node--hover--BackgroundColor);
}

.pf-v6-c-tree-view__action {
  margin-inline-end: var(--pf-v6-c-tree-view__action--MarginInlineEnd);
}

.pf-v6-c-tree-view__content:has(> .pf-v6-c-tree-view__node.pf-m-current) {
  --pf-v6-c-tree-view__content--BackgroundColor: var(--pf-v6-c-tree-view__node--m-current--BackgroundColor);
}

.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 1 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 2 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 1 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 1 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 3 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 2 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 2 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 4 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 3 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 3 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 5 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 4 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 4 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 6 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 5 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 5 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 7 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 6 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 6 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 8 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 7 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 7 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 9 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 8 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 8 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
.pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item .pf-v6-c-tree-view__list-item {
  --pf-v6-c-tree-view__list-item__list-item__node-toggle--InsetInlineStart: var(--pf-v6-c-tree-view__node--PaddingInlineStart);
  --pf-v6-c-tree-view__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view__node--nested-indent--base) * 10 + var(--pf-v6-c-tree-view__node--indent--base));
  --pf-v6-c-tree-view--m-guides--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-guides--guide-left--base--offset));
  --pf-v6-c-tree-view--m-compact__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact__node--nested-indent--base) * 9 + var(--pf-v6-c-tree-view--m-compact__node--indent--base));
  --pf-v6-c-tree-view--m-compact--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--base-border--InsetInlineStart--offset));
  --pf-v6-c-tree-view--m-compact--m-no-background__node--PaddingInlineStart: calc(var(--pf-v6-c-tree-view--m-compact--m-no-background__node--nested-indent--base) * 9 + var(--pf-v6-c-tree-view--m-compact--m-no-background__node--indent--base));
  --pf-v6-c-tree-view--m-compact--m-no-background--border--nested--InsetInlineStart: calc(var(--pf-v6-c-tree-view__node--PaddingInlineStart) - var(--pf-v6-c-tree-view--m-compact--m-no-background--base-border--InsetInlineStart--offset));
}
```

### Expand/Collapse Mechanics

- PatternFly CSS handles animations via transitions
- State managed via `.pf-m-expanded` modifier class
- `aria-expanded` attribute for accessibility
- Support keyboard navigation (Arrow keys, Space, Enter) per PatternFly patterns

## Icon and Badge System

### Icon Mapping

**Use PatternFly icons (not emoji)** - Inline SVG from PatternFly icon set:

| Node Type | PatternFly Icon | Description |
|-----------|-----------------|-------------|
| package | `cube-icon` or `folder-icon` | Package root |
| module | `code-icon` or `file-code-icon` | JavaScript module |
| element | `code-branch-icon` | Custom element |
| class | `object-icon` | Class/mixin |
| function | `function-icon` | Function |
| attribute | `at-icon` | HTML attribute |
| property | `key-icon` | JS property |
| method | `function-icon` | Method |
| event | `bell-icon` | Event |
| slot | `plug-icon` | Slot |
| css-property | `palette-icon` | CSS custom property |
| css-part | `cube-icon` | CSS part |
| css-state | `toggle-on-icon` | CSS custom state |
| type | `file-icon` | Type definition |
| group | `folder-icon` | API group (Attributes, Properties, etc.) |

**Implementation:**
- File: `serve/middleware/routes/templates/js/manifest-icons.js`
- Export `getIconForType(type)` returning PatternFly icon SVG strings
- Icons from PatternFly icon library (check existing pf-v6 components for icon patterns)

### Badge Types

**Metadata badges** - Use `pf-v6-label` component:

| Metadata | Badge Text | Color | Compact |
|----------|-----------|-------|---------|
| deprecated | "deprecated" | warning | yes |
| required | "required" | blue | yes |
| readonly | "readonly" | grey | yes |
| private | "private" | red | yes |
| protected | "protected" | orange | yes |
| static | "static" | grey | yes |

**Number badges** - Display counts using `pf-v6-label` or badge count in tree view:

| Node Type | Count Display | Example |
|-----------|--------------|---------|
| package | Number of modules | "5" |
| module | Number of declarations | "3" |
| element | Total API surface (attrs + props + methods + events + slots) | "12" |
| group (Attributes) | Number of attributes | "4" |
| group (Properties) | Number of properties | "6" |
| group (Methods) | Number of methods | "3" |
| group (Events) | Number of events | "2" |
| group (Slots) | Number of slots | "1" |
| group (CSS Properties) | Number of CSS custom properties | "5" |
| group (CSS Parts) | Number of CSS parts | "2" |
| group (CSS States) | Number of CSS custom states | "1" |

**Link attributes and class fields:**
- When showing a property, if it reflects to an attribute, show both
- Display as: `disabled` (attribute) ↔ `disabled` (property)
- Visual indicator showing they're linked (bidirectional reflection)

### Visual Design

- **Indentation**: 1.5rem per level
- **Typography**: Body text (14px), Summary (12px), Types (monospace)
- **Hover**: Background highlight
- **Selected**: Blue left border + background tint

## Search Implementation

### Search Input

Use `pf-v6-text-input-group` in a toolbar above `pf-v6-tree-view`:
- Debounced input (300ms)
- Clear button
- Keyboard shortcut (Ctrl+F / Cmd+F to focus)
- Server-rendered in template (part of declarative tree structure)

### Full-text Search Algorithm

**Searchable fields:**
- Node name
- Node summary
- Node description
- Type annotations
- Default values

**Implementation:** `serve/middleware/routes/templates/js/manifest-search.js`

```javascript
class ManifestSearchIndex {
  buildIndex(rootNode) {
    // Flatten tree and extract searchable text
    this.walkTree(rootNode, (node) => {
      const text = [
        node.name,
        node.summary || '',
        node.description || '',
        node.metadata.type || '',
        node.metadata.default || '',
      ].join(' ').toLowerCase();

      this.index.set(node.id, { node, text });
    });
  }

  search(query) {
    const lowerQuery = query.toLowerCase();
    const results = [];

    for (const [id, { node, text }] of this.index.entries()) {
      if (text.includes(lowerQuery)) {
        results.push({
          node,
          score: this.calculateScore(text, lowerQuery),
        });
      }
    }

    return results.sort((a, b) => b.score - a.score);
  }
}
```

### Filtering and Highlighting

**Filter behavior:**
1. Search returns matching nodes
2. Auto-expand parent nodes to reveal matches
3. Highlight matching text with `<mark>` tags
4. Show match count ("5 results")
5. hide filtered out nodes (not matching, no matching descendants)

**Performance:** <10ms search time for typical manifests

## Markdown Rendering

### Server-side Rendering (Preferred)

**Go implementation:**
- Leverage existing `markdownToHTML()` in `serve/middleware/routes/markdown.go`
- Pre-convert markdown fields before sending to client
- Use goldmark with GFM

**Approach:**
1. Add `EnrichManifestWithHTML(pkg *manifest.Package)` function
2. Walk manifest tree and convert Summary/Description fields
3. Store HTML in new fields: `SummaryHTML`, `DescriptionHTML`
4. Client renders pre-converted HTML

### Client-side Fallback

**If server-side not feasible:**
- Use `marked` library (13KB minified)
- DOMPurify for sanitization
- Cache rendered results

## Data Transformation

### Manifest to Tree Structure

**File:** `serve/middleware/routes/templates/js/manifest-tree-builder.js`

```javascript
class ManifestTreeBuilder {
  build() {
    return {
      id: 'root',
      type: 'package',
      name: this.manifest.schemaVersion || 'Package',
      children: this.manifest.modules.map(mod => this.buildModule(mod)),
      expanded: true,
    };
  }

  buildModule(module) {
    return {
      id: `module-${this.nodeIdCounter++}`,
      type: 'module',
      name: module.path,
      summary: module.summary,
      children: module.declarations.map(decl => this.buildDeclaration(decl)),
      expanded: false,
    };
  }

  buildElement(element) {
    const children = [];

    // Group APIs by category
    if (element.attributes?.length) {
      children.push({
        id: `${element.tagName}-attrs`,
        type: 'group',
        name: 'Attributes',
        children: element.attributes.map(attr => this.buildAttribute(attr)),
      });
    }

    // Similar for: properties, methods, events, slots, cssProperties, cssParts, cssStates

    return {
      id: `element-${element.tagName}`,
      type: 'element',
      name: element.tagName,
      summary: element.summary,
      children,
    };
  }
}
```

### Caching Strategy

1. **Manifest JSON**: Cache raw manifest (from server)
2. **Tree structure**: Cache transformed tree (sessionStorage)
3. **Search index**: Cache in memory
4. **Rendered nodes**: Cache HTML fragments

## Files to Create/Modify

### New Web Component Files

**cem-manifest-browser:**
- `serve/middleware/routes/templates/elements/cem-manifest-browser/cem-manifest-browser.js`
- `serve/middleware/routes/templates/elements/cem-manifest-browser/cem-manifest-browser.html`
- `serve/middleware/routes/templates/elements/cem-manifest-browser/cem-manifest-browser.css`

**pf-v6-tree-view (NEW PatternFly component):**
- `serve/middleware/routes/templates/elements/pf-v6-tree-view/pf-v6-tree-view.js`
- `serve/middleware/routes/templates/elements/pf-v6-tree-view/pf-v6-tree-view.css`
  - Compiled CSS from patternfly.org (tree-view component)
  - Source reference: `../patternfly/patternfly/src/patternfly/components/TreeView/tree-view.scss`

### Shared JavaScript Files

- `serve/middleware/routes/templates/js/manifest-icons.js` - Icon library
- `serve/middleware/routes/templates/js/manifest-tree-builder.js` - Tree transformation
- `serve/middleware/routes/templates/js/manifest-search.js` - Search index

### Template Files to Modify

- `serve/middleware/routes/templates/elements/cem-serve-chrome/cem-serve-chrome.html` - Add manifest tab

### Go Middleware Changes

**Add manifest to template data:**
- `serve/middleware/routes/chrome.go` - Add manifest JSON to `ChromeData` struct

**Manifest enrichment (NEW):**
- `serve/middleware/routes/manifest.go` - Implement `EnrichManifestWithHTML()`

**Update demo route:**
- `serve/middleware/routes/demo.go` - Pass manifest to chrome template

### Test Files

**Component tests:**
- `serve/middleware/routes/templates/elements/cem-manifest-browser/cem-manifest-browser_test.js`
- `serve/middleware/routes/templates/elements/cem-tree-view/cem-tree-view_test.js`

**Go tests:**
- `serve/middleware/routes/manifest_test.go` - Test manifest enrichment (fixture/golden pattern)

## Implementation Phases

### Phase 1: PatternFly Tree View Component

1. **Extract PatternFly tree view CSS**
   - Download compiled CSS from patternfly.org for tree-view component
   - Reference: ../patternfly/patternfly/src/patternfly/components/TreeView/tree-view.scss
   - Copy all `.pf-v-c-tree-view*` classes and associated styles
   - Preserve all CSS tokens and variables

2. **Implement pf-v6-tree-view web component**
   - Translate from React component in `../patternfly/patternfly-react/packages/react-core/src/components/TreeView/`
   - Support attributes: `variant`, `guides`, `multi-select`, `checkboxes`, `selectable`, `all-expanded`
   - Support both data property (imperative) and slot-based (declarative) population
   - Implement event dispatching: `pf-tree-view:select`, `pf-tree-view:check`, `pf-tree-view:expand`, `pf-tree-view:collapse`
   - Render proper PatternFly HTML structure with correct CSS classes

3. **Implement pf-v6-tree-view-list-item (internal)**
   - Node rendering with PatternFly classes
   - Expand/collapse toggle button
   - Support for icon, name, badge, action slots
   - Recursive nesting for children
   - Checkbox support (if enabled)

4. **Create manifest-icons.js**
   - Implement `getIconForType()` function
   - All node type icons as SVG strings (following PatternFly icon style)

### Phase 2: Data Integration

5. **Create manifest-tree-builder.js**
   - Transform manifest JSON to tree nodes
   - Handle all declaration types
   - Group APIs by category (Attributes, Properties, etc.)
   - Generate data structure compatible with pf-v6-tree-view

6. **Go middleware changes**
   - Add `ManifestJSON` field to `ChromeData` struct
   - Serialize manifest to JSON
   - Embed in template via `<script>` tag

7. **Test pf-v6-tree-view with real manifest**
   - Verify tree structure correctness
   - Test with multi-module manifests
   - Verify PatternFly styling works correctly

### Phase 3: Manifest Browser Component

8. **Create cem-manifest-browser component**
   - pf-v6-tree-view with search slot populated
   - Server-side rendered tree structure (Go template)
   - Custom styling for manifest-specific needs
   - Event listeners for tree interactions (if needed)

9. **Integrate with chrome template**
   - Add "Manifest" tab to cem-drawer
   - Test in dev server
   - Verify tab switching
   - Verify tree view renders in sidebar

10. **Style and polish**
    - Ensure PatternFly tree view fits in sidebar
    - Add custom icons for manifest node types
    - Responsive behavior for drawer

### Phase 4: Search Functionality

11. **Create manifest-search.js**
    - Implement search index builder
    - Search algorithm with ranking
    - Debounced input handling

12. **Integrate search with pf-v6-tree-view**
    - Search input is in tree view's search slot
    - pf-v6-tree-view handles search input changes
    - Auto-expand matched nodes
    - Highlight matching text in tree nodes
    - Display result count (or use tree view's built-in handling)

### Phase 5: Markdown Rendering

13. **Client-side markdown (interim solution)**
    - Import `marked` library
    - Render in node summaries/descriptions
    - Sanitize with DOMPurify

14. **Display in tree nodes**
    - Show summary in pf-v6-tree-view nodes (custom slot or text content)
    - Truncate long summaries
    - Expand to show full description

### Phase 6: Polish & Testing

15. **Accessibility**
    - Verify pf-v6-tree-view ARIA implementation
    - Test keyboard navigation (arrows, space, enter)
    - Screen reader support
    - Focus management in sidebar context

16. **Test coverage**
    - Unit tests for pf-v6-tree-view component
    - Unit tests for cem-manifest-browser
    - Golden file tests for tree transformation
    - Integration tests with dev server
    - Cross-browser testing

17. **Performance optimization**
    - Test with large manifests (100+ elements)
    - Virtual scrolling if needed (may require pf-v6-tree-view enhancement)
    - Lazy loading of children
    - Debounce search input

## Critical Files

**Core implementation:**
1. `serve/middleware/routes/templates/elements/pf-v6-tree-view/pf-v6-tree-view.js` - NEW PatternFly web component (translate from React)
2. `serve/middleware/routes/templates/elements/pf-v6-tree-view/pf-v6-tree-view.css` - PatternFly tree view CSS (from patternfly.org)
3. `serve/middleware/routes/templates/elements/cem-manifest-browser/cem-manifest-browser.js` - Main controller (uses pf-v6-tree-view)
4. `serve/middleware/routes/templates/js/manifest-tree-builder.js` - Data transformation (manifest → tree data)
5. `serve/middleware/routes/chrome.go` - Backend integration (embed manifest JSON)
6. `serve/middleware/routes/templates/elements/cem-serve-chrome/cem-serve-chrome.html` - UI integration (add manifest tab)

**Reference files (for translation):**
- `../patternfly/patternfly-react/packages/react-core/src/components/TreeView/TreeView.tsx` - React implementation
- `../patternfly/patternfly-react/packages/react-core/src/components/TreeView/TreeViewListItem.tsx` - React list item
- `../patternfly/patternfly/src/patternfly/components/TreeView/tree-view.scss` - SASS source

## Success Criteria

- [ ] pf-v6-tree-view web component successfully translates PatternFly React component
- [ ] pf-v6-tree-view uses PatternFly CSS classes and styling correctly
- [ ] Tree view displays full manifest structure (modules → elements → APIs)
- [ ] All node types have correct icons (manifest-specific)
- [ ] Badges display metadata (deprecated, required, readonly, etc.)
- [ ] Full-text search works across all fields
- [ ] Markdown summaries render correctly
- [ ] Expand/collapse works smoothly (PatternFly animations)
- [ ] Keyboard navigation fully functional (PatternFly patterns)
- [ ] Accessible (ARIA, screen readers per PatternFly standards)
- [ ] Persists expanded state across refreshes
- [ ] Performance acceptable for large manifests (100+ elements)
- [ ] Test coverage >80%
- [ ] pf-v6-tree-view is reusable for other purposes (not manifest-specific)
