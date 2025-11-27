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

3. **`pf-v6-tree-item`** - Individual tree item
   - Location: `serve/middleware/routes/templates/elements/pf-v6-tree-item/`
   - Translated from PatternFly React TreeViewListItem
   - Handles node rendering, expansion, selection, checkboxes
   - Files: `.js`, `.css`
   - Reference: `../patternfly/patternfly-react/packages/react-core/src/components/TreeView/TreeViewListItem.tsx`

### Integration with cem-drawer

Add manifest browser as **third tab** in `cem-drawer` (alongside Knobs and Server Logs):

**File:** `serve/middleware/routes/templates/elements/cem-serve-chrome/cem-serve-chrome.html`

```html
<cem-drawer>
  <pf-v6-tabs>
    <pf-v6-tab title="Knobs">...</pf-v6-tab>
    <pf-v6-tab title="Server Logs">...</pf-v6-tab>
    <pf-v6-tab title="Manifest Browser">
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
  <pf-v6-tree-item name="{{.Path}}">
    <svg slot="icon"><!-- module icon --></svg>
    <span slot="badge">{{len .Declarations}}</span>

    {{range .Declarations}}
    <pf-v6-tree-item name="{{.TagName}}">
      <svg slot="icon"><!-- element icon --></svg>
      <span slot="badge">{{.APICount}}</span>

      {{if .Attributes}}
      <pf-v6-tree-item name="Attributes">
        <svg slot="icon"><!-- folder icon --></svg>
        <span slot="badge">{{len .Attributes}}</span>

        {{range .Attributes}}
        <pf-v6-tree-item name="{{.Name}}">
          <svg slot="icon"><!-- attr icon --></svg>
          {{if .Deprecated}}<pf-v6-label compact status="warning">deprecated</pf-v6-label>{{end}}
        </pf-v6-tree-item>
        {{end}}
      </pf-v6-tree-item>
      {{end}}
    </pf-v6-tree-item>
    {{end}}
  </pf-v6-tree-item>
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
- `.pf-v6-c-tree-view` - Root container
- `.pf-v6-c-tree-view__list` - List container (ul)
- `.pf-v6-c-tree-view__list-item` - List item (li)
- `.pf-v6-c-tree-view__node` - Node wrapper
- `.pf-v6-c-tree-view__node-toggle` - Expand/collapse button
- `.pf-v6-c-tree-view__node-toggle-icon` - Chevron icon
- `.pf-v6-c-tree-view__node-container` - Content container
- `.pf-v6-c-tree-view__node-icon` - Node icon
- `.pf-v6-c-tree-view__node-text` - Node text/name
- `.pf-v6-c-tree-view__node-count` - Badge/count
- `.pf-v6-c-tree-view__action` - Action slot

**Modifiers:**
- `.pf-m-compact` - Compact variant
- `.pf-m-no-background` - No background for compact
- `.pf-m-guides` - Visual guides between nodes
- `.pf-m-expanded` - Expanded state
- `.pf-m-current` - Selected/current item

### HTML Structure (from React-rendered output)

see ./tree-view-example.html

### Compiled SASS

see ./tree-view-compiled.css

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
- `serve/middleware/routes/templates/elements/pf-v6-tree-view/pf-v6-tree-view_test.js`

**Go tests:**
- `serve/middleware/routes/manifest_test.go` - Test manifest enrichment (fixture/golden pattern)

## Implementation Phases

### Phase 1: PatternFly Tree View Component

1. **Extract PatternFly tree view CSS**
   - Download compiled CSS from patternfly.org for tree-view component
   - Reference: ../patternfly/patternfly/src/patternfly/components/TreeView/tree-view.scss
   - Copy all `.pf-v6-c-tree-view*` classes and associated styles
   - Preserve all CSS tokens and variables

2. **Implement pf-v6-tree-view web component**
   - Translate from React component in `../patternfly/patternfly-react/packages/react-core/src/components/TreeView/`
   - Support attributes: `variant`, `guides`, `multi-select`, `checkboxes`, `selectable`, `all-expanded`
   - Support both data property (imperative) and slot-based (declarative) population
   - Implement event dispatching: `pf-tree-view:select`, `pf-tree-view:check`, `pf-tree-view:expand`, `pf-tree-view:collapse`
   - Render proper PatternFly HTML structure with correct CSS classes

3. **Implement pf-v6-tree-item**
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
