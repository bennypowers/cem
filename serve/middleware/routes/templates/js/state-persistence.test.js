import { expect } from '@open-wc/testing';
import sinon from 'sinon';
import { StatePersistence, getTreeNodeId, findTreeItemById } from './state-persistence.js';

describe('StatePersistence', () => {
  beforeEach(() => {
    // Clear cookies and localStorage
    document.cookie.split(';').forEach(cookie => {
      const name = cookie.split('=')[0].trim();
      document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
    });
    localStorage.clear();
  });

  describe('cookie state management', () => {
    it('returns default state when no cookie exists', () => {
      const state = StatePersistence.getState();

      expect(state).to.deep.equal({
        colorScheme: 'system',
        drawer: { open: false, height: 400 },
        tabs: { selectedIndex: 0 },
        sidebar: { collapsed: false },
        version: 1
      });
    });

    it('sets and retrieves state from cookie', () => {
      const testState = {
        colorScheme: 'dark',
        drawer: { open: true, height: 500 },
        tabs: { selectedIndex: 1 },
        sidebar: { collapsed: true },
        version: 1
      };

      StatePersistence.setState(testState);
      const retrieved = StatePersistence.getState();

      expect(retrieved).to.deep.equal(testState);
    });

    it('updates partial state via deepMerge', () => {
      StatePersistence.setState({
        colorScheme: 'light',
        drawer: { open: false, height: 400 },
        tabs: { selectedIndex: 0 },
        sidebar: { collapsed: false },
        version: 1
      });

      StatePersistence.updateState({
        colorScheme: 'dark',
        drawer: { height: 600 }
      });

      const state = StatePersistence.getState();

      expect(state.colorScheme).to.equal('dark');
      expect(state.drawer.open).to.equal(false); // Unchanged
      expect(state.drawer.height).to.equal(600); // Updated
      expect(state.tabs.selectedIndex).to.equal(0); // Unchanged
      expect(state.sidebar.collapsed).to.equal(false); // Unchanged
    });

    it('handles invalid JSON in cookie gracefully', () => {
      document.cookie = `${StatePersistence.COOKIE_NAME}=invalid-json; Path=/`;

      const state = StatePersistence.getState();

      expect(state).to.deep.equal(StatePersistence.getDefaultState());
    });

    it('warns when cookie size exceeds 4KB', () => {
      const consoleSpy = sinon.spy(console, 'error');

      const largeState = {
        colorScheme: 'dark',
        drawer: { open: true, height: 400 },
        tabs: { selectedIndex: 0 },
        version: 1,
        // Add large data to exceed 4KB
        largeData: 'x'.repeat(5000)
      };

      StatePersistence.setState(largeState);

      expect(consoleSpy.called).to.be.true;
      expect(consoleSpy.firstCall.args[0]).to.include('4KB');

      consoleSpy.restore();
    });

    it('encodes and decodes cookie values correctly', () => {
      const stateWithSpecialChars = {
        colorScheme: 'dark & light',
        drawer: { open: true, height: 400 },
        tabs: { selectedIndex: 0 },
        version: 1
      };

      StatePersistence.setState(stateWithSpecialChars);
      const retrieved = StatePersistence.getState();

      expect(retrieved.colorScheme).to.equal('dark & light');
    });
  });

  describe('deepMerge', () => {
    it('merges nested objects', () => {
      const target = {
        a: 1,
        b: { c: 2, d: 3 }
      };

      const source = {
        b: { d: 4, e: 5 },
        f: 6
      };

      const result = StatePersistence.deepMerge(target, source);

      expect(result).to.deep.equal({
        a: 1,
        b: { c: 2, d: 4, e: 5 },
        f: 6
      });
    });

    it('handles arrays as primitive values', () => {
      const target = { arr: [1, 2, 3] };
      const source = { arr: [4, 5] };

      const result = StatePersistence.deepMerge(target, source);

      expect(result.arr).to.deep.equal([4, 5]);
    });

    it('does not mutate original objects', () => {
      const target = { a: { b: 1 } };
      const source = { a: { c: 2 } };

      const result = StatePersistence.deepMerge(target, source);

      expect(target.a).to.deep.equal({ b: 1 });
      expect(source.a).to.deep.equal({ c: 2 });
      expect(result.a).to.deep.equal({ b: 1, c: 2 });
    });
  });

  describe('tree state management', () => {
    it('returns default tree state when none exists', () => {
      const treeState = StatePersistence.getTreeState();

      expect(treeState).to.deep.equal({
        expanded: [],
        selected: ''
      });
    });

    it('sets and retrieves tree state from localStorage', () => {
      const treeState = {
        expanded: ['node-1', 'node-2'],
        selected: 'node-3'
      };

      StatePersistence.setTreeState(treeState);
      const retrieved = StatePersistence.getTreeState();

      expect(retrieved).to.deep.equal(treeState);
    });

    it('updates partial tree state', () => {
      StatePersistence.setTreeState({
        expanded: ['node-1'],
        selected: 'node-2'
      });

      StatePersistence.updateTreeState({
        expanded: ['node-1', 'node-3']
      });

      const treeState = StatePersistence.getTreeState();

      expect(treeState.expanded).to.deep.equal(['node-1', 'node-3']);
      expect(treeState.selected).to.equal('node-2'); // Unchanged
    });

    it('handles invalid JSON in localStorage gracefully', () => {
      localStorage.setItem(StatePersistence.TREE_STORAGE_KEY, 'invalid-json');

      const treeState = StatePersistence.getTreeState();

      expect(treeState).to.deep.equal({
        expanded: [],
        selected: ''
      });
    });
  });

  describe('sidebar state management', () => {
    it('persists sidebar collapsed state', () => {
      StatePersistence.setState({
        colorScheme: 'system',
        drawer: { open: false, height: 400 },
        tabs: { selectedIndex: 0 },
        sidebar: { collapsed: true },
        version: 1
      });

      const state = StatePersistence.getState();

      expect(state.sidebar.collapsed).to.equal(true);
    });

    it('persists sidebar expanded state', () => {
      StatePersistence.setState({
        colorScheme: 'system',
        drawer: { open: false, height: 400 },
        tabs: { selectedIndex: 0 },
        sidebar: { collapsed: false },
        version: 1
      });

      const state = StatePersistence.getState();

      expect(state.sidebar.collapsed).to.equal(false);
    });

    it('updates sidebar state via partial update', () => {
      StatePersistence.setState({
        colorScheme: 'system',
        drawer: { open: false, height: 400 },
        tabs: { selectedIndex: 0 },
        sidebar: { collapsed: false },
        version: 1
      });

      // Update sidebar state
      StatePersistence.updateState({
        sidebar: { collapsed: true }
      });

      const state = StatePersistence.getState();

      expect(state.sidebar.collapsed).to.equal(true);
    });

    it('provides default expanded state', () => {
      const state = StatePersistence.getDefaultState();

      // Sidebar expanded by default
      expect(state.sidebar.collapsed).to.equal(false);
    });
  });

  describe('state migration', () => {
    it('migrates old version state', () => {
      const oldState = {
        colorScheme: 'dark',
        version: 0
      };

      const migratedState = StatePersistence.migrateState(oldState);

      expect(migratedState).to.deep.equal(StatePersistence.getDefaultState());
    });
  });
});

describe('getTreeNodeId', () => {
  it('generates ID from type only', () => {
    const treeItem = document.createElement('div');
    treeItem.setAttribute('data-type', 'module');

    const id = getTreeNodeId(treeItem);

    expect(id).to.equal('module');
  });

  it('generates ID from type and module path', () => {
    const treeItem = document.createElement('div');
    treeItem.setAttribute('data-type', 'class');
    treeItem.setAttribute('data-module-path', 'src/components.js');

    const id = getTreeNodeId(treeItem);

    expect(id).to.equal('class:src/components.js');
  });

  it('generates ID from type, module, and tag name', () => {
    const treeItem = document.createElement('div');
    treeItem.setAttribute('data-type', 'element');
    treeItem.setAttribute('data-module-path', 'src/my-element.js');
    treeItem.setAttribute('data-tag-name', 'my-element');

    const id = getTreeNodeId(treeItem);

    expect(id).to.equal('element:src/my-element.js:my-element');
  });

  it('generates full ID with all parts', () => {
    const treeItem = document.createElement('div');
    treeItem.setAttribute('data-type', 'method');
    treeItem.setAttribute('data-module-path', 'src/my-element.js');
    treeItem.setAttribute('data-tag-name', 'my-element');
    treeItem.setAttribute('data-name', 'connectedCallback');

    const id = getTreeNodeId(treeItem);

    expect(id).to.equal('method:src/my-element.js:my-element:connectedCallback');
  });

  it('uses data-path as fallback for module path', () => {
    const treeItem = document.createElement('div');
    treeItem.setAttribute('data-type', 'function');
    treeItem.setAttribute('data-path', 'src/utils.js');

    const id = getTreeNodeId(treeItem);

    expect(id).to.equal('function:src/utils.js');
  });
});

describe('findTreeItemById', () => {
  beforeEach(() => {
    // Create a mock DOM structure
    document.body.innerHTML = `
      <pf-v6-tree-item data-type="module" data-module-path="src/components.js"></pf-v6-tree-item>
      <pf-v6-tree-item data-type="element" data-module-path="src/my-element.js" data-tag-name="my-element"></pf-v6-tree-item>
      <pf-v6-tree-item data-type="method" data-module-path="src/my-element.js" data-tag-name="my-element" data-name="render"></pf-v6-tree-item>
    `;
  });

  afterEach(() => {
    document.body.innerHTML = '';
  });

  it('finds tree item by type only', () => {
    const item = findTreeItemById('module');

    expect(item).to.exist;
    expect(item.getAttribute('data-type')).to.equal('module');
  });

  it('finds tree item by type and module path', () => {
    const item = findTreeItemById('module:src/components.js');

    expect(item).to.exist;
    expect(item.getAttribute('data-module-path')).to.equal('src/components.js');
  });

  it('finds tree item by type, module, and tag name', () => {
    const item = findTreeItemById('element:src/my-element.js:my-element');

    expect(item).to.exist;
    expect(item.getAttribute('data-tag-name')).to.equal('my-element');
  });

  it('finds tree item with all parts', () => {
    const item = findTreeItemById('method:src/my-element.js:my-element:render');

    expect(item).to.exist;
    expect(item.getAttribute('data-name')).to.equal('render');
  });

  it('returns null for non-existent ID', () => {
    const item = findTreeItemById('nonexistent:id');

    expect(item).to.be.null;
  });

  it('escapes special characters in selectors', () => {
    // Create element with special characters
    const specialItem = document.createElement('pf-v6-tree-item');
    specialItem.setAttribute('data-type', 'class');
    specialItem.setAttribute('data-name', 'My[Special]Class');
    document.body.appendChild(specialItem);

    const item = findTreeItemById('class:::My[Special]Class');

    expect(item).to.exist;
    expect(item.getAttribute('data-name')).to.equal('My[Special]Class');
  });
});

describe('CSS.escape polyfill', () => {
  it('escapes special CSS characters', () => {
    expect(CSS.escape('[test]')).to.equal('\\[test\\]');
    expect(CSS.escape('my.class')).to.equal('my\\.class');
    expect(CSS.escape('id#123')).to.equal('id\\#123');
  });

  it('handles dash at start', () => {
    expect(CSS.escape('-webkit-transform')).to.equal('-webkit-transform');
  });

  it('handles numbers at start', () => {
    // Numbers at start should be escaped
    const escaped = CSS.escape('1st-item');
    expect(escaped).to.not.equal('1st-item');
  });
});
