/**
 * State persistence via cookies for SSR hydration
 *
 * Manages UI state (color scheme, drawer, tree) in a single JSON cookie
 * for server-side rendering and client-side updates.
 */
export class StatePersistence {
  static COOKIE_NAME = 'cem-serve-state';
  static MAX_AGE = 2592000; // 30 days
  static VERSION = 1;

  /**
   * Read state from cookie
   */
  static getState() {
    const cookie = document.cookie
      .split('; ')
      .find(row => row.startsWith(this.COOKIE_NAME + '='));

    if (!cookie) return this.getDefaultState();

    try {
      const json = decodeURIComponent(cookie.split('=')[1]);
      const state = JSON.parse(json);

      // Validate version and structure
      if (state.version !== this.VERSION) {
        return this.migrateState(state);
      }

      return state;
    } catch (e) {
      console.warn('[state-persistence] Failed to parse state cookie:', e);
      return this.getDefaultState();
    }
  }

  /**
   * Write state to cookie with size validation
   */
  static setState(state, retryCount = 0) {
    const json = JSON.stringify(state);
    const encoded = encodeURIComponent(json);

    // Check size (4KB limit)
    if (encoded.length > 4000) {
      // Prevent infinite recursion - only try trimming once
      if (retryCount > 0) {
        console.error('[state-persistence] Cookie still too large after trimming, clearing tree state');
        state.tree.expanded = [];
        state.tree.selected = null;
      } else {
        console.warn('[state-persistence] Cookie size exceeds limit, trimming expanded nodes');
        // Keep only first 20 expanded nodes
        state.tree.expanded = state.tree.expanded.slice(0, 20);
      }
      return this.setState(state, retryCount + 1);
    }

    document.cookie = `${this.COOKIE_NAME}=${encoded}; Path=/; SameSite=Lax; Max-Age=${this.MAX_AGE}`;
  }

  /**
   * Update partial state (deep merge)
   */
  static updateState(partial) {
    const current = this.getState();
    const updated = this.deepMerge(current, partial);
    this.setState(updated);
  }

  /**
   * Default state factory
   */
  static getDefaultState() {
    return {
      colorScheme: 'system',
      drawer: { open: false, height: 400 },
      tree: { expanded: [], selected: '' },
      tabs: { selectedIndex: 0 },
      version: this.VERSION
    };
  }

  /**
   * Deep merge helper
   */
  static deepMerge(target, source) {
    const result = { ...target };
    for (const key in source) {
      if (source[key] && typeof source[key] === 'object' && !Array.isArray(source[key])) {
        result[key] = this.deepMerge(target[key] || {}, source[key]);
      } else {
        result[key] = source[key];
      }
    }
    return result;
  }

  /**
   * Migrate from localStorage (one-time migration)
   */
  static migrateFromLocalStorage() {
    const state = this.getDefaultState();

    // Migrate color scheme
    try {
      const colorScheme = localStorage.getItem('cem-serve-color-scheme');
      if (colorScheme) state.colorScheme = colorScheme;
    } catch (e) { /* ignore */ }

    // Migrate drawer state
    try {
      const drawerOpen = localStorage.getItem('cem-serve-drawer-open');
      if (drawerOpen !== null) state.drawer.open = drawerOpen === 'true';

      const drawerHeight = localStorage.getItem('cem-serve-drawer-height');
      if (drawerHeight !== null) state.drawer.height = parseInt(drawerHeight, 10);
    } catch (e) { /* ignore */ }

    // Migrate tabs state
    try {
      const activeTab = localStorage.getItem('cem-serve-active-tab');
      if (activeTab !== null) {
        // Old localStorage stored panel IDs like 'panel-logs', 'panel-knobs', etc.
        // Map to selectedIndex (0 = Knobs, 1 = Manifest Browser, 2 = Server Logs)
        const indexMap = {
          'panel-knobs': 0,
          'panel-manifest': 1,
          'panel-logs': 2
        };
        state.tabs.selectedIndex = indexMap[activeTab] ?? 0;
      }
    } catch (e) { /* ignore */ }

    this.setState(state);
  }

  /**
   * Future: migrate state schema versions
   */
  static migrateState(oldState) {
    // For now, just return default and log
    console.warn('[state-persistence] Migrating from version', oldState.version, 'to', this.VERSION);
    return this.getDefaultState();
  }
}

/**
 * Generate tree node ID from data attributes
 * Format: {type}:{module-path}:{tag-name}:{member-name}
 */
export function getTreeNodeId(treeItem) {
  const type = treeItem.getAttribute('data-type');
  const modulePath = treeItem.getAttribute('data-module-path') || treeItem.getAttribute('data-path');
  const tagName = treeItem.getAttribute('data-tag-name');
  const name = treeItem.getAttribute('data-name');

  const parts = [type];
  if (modulePath) parts.push(modulePath);
  if (tagName) parts.push(tagName);
  if (name) parts.push(name);

  return parts.join(':');
}

/**
 * Find tree item by node ID
 * Builds a selector from the node ID parts
 */
export function findTreeItemById(nodeId, rootElement = document) {
  const parts = nodeId.split(':');
  const [type, modulePath, tagName, name] = parts;

  let selector = `pf-v6-tree-item[data-type="${type}"]`;
  if (modulePath) {
    selector += `[data-module-path="${modulePath}"]`;
  }
  if (tagName) {
    selector += `[data-tag-name="${tagName}"]`;
  }
  if (name) {
    selector += `[data-name="${name}"]`;
  }

  return rootElement.querySelector(selector);
}
