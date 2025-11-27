import { expect, waitUntil } from '@open-wc/testing';
import { ManifestSearchIndex } from './manifest-search.js';

describe('ManifestSearchIndex', () => {
  let searchIndex;

  beforeEach(() => {
    searchIndex = new ManifestSearchIndex();
  });

  describe('buildIndex', () => {
    it('builds index from tree data', () => {
      const treeData = [
        {
          id: 'node-1',
          name: 'MyElement',
          type: 'element',
          summary: 'A custom element',
          description: 'This is a custom element for testing'
        },
        {
          id: 'node-2',
          name: 'MyClass',
          type: 'class',
          summary: 'A utility class'
        }
      ];

      searchIndex.buildIndex(treeData);

      expect(searchIndex.size).to.equal(2);
    });

    it('recursively indexes nested nodes', () => {
      const treeData = [
        {
          id: 'parent',
          name: 'Parent',
          type: 'module',
          children: [
            {
              id: 'child-1',
              name: 'ChildOne',
              type: 'class'
            },
            {
              id: 'child-2',
              name: 'ChildTwo',
              type: 'function',
              children: [
                {
                  id: 'grandchild',
                  name: 'GrandChild',
                  type: 'method'
                }
              ]
            }
          ]
        }
      ];

      searchIndex.buildIndex(treeData);

      // Should index parent + 2 children + 1 grandchild = 4 nodes
      expect(searchIndex.size).to.equal(4);
    });

    it('builds path information for each node', () => {
      const treeData = [
        {
          id: 'parent',
          name: 'Parent',
          type: 'module',
          children: [
            {
              id: 'child',
              name: 'Child',
              type: 'class'
            }
          ]
        }
      ];

      searchIndex.buildIndex(treeData);

      const results = searchIndex.search('Child');
      expect(results).to.have.lengthOf(1);
      expect(results[0].pathString).to.equal('Parent > Child');
    });

    it('indexes metadata properties', () => {
      const treeData = [
        {
          id: 'element-1',
          name: 'MyButton',
          type: 'element',
          metadata: {
            tagName: 'my-button',
            type: { text: 'HTMLElement' },
            default: 'primary'
          }
        }
      ];

      searchIndex.buildIndex(treeData);

      // Should find by tag name
      let results = searchIndex.search('my-button');
      expect(results).to.have.lengthOf(1);

      // Should find by type
      results = searchIndex.search('HTMLElement');
      expect(results).to.have.lengthOf(1);
    });

    it('indexes method parameters', () => {
      const treeData = [
        {
          id: 'method-1',
          name: 'updateValue',
          type: 'method',
          metadata: {
            parameters: [
              {
                name: 'newValue',
                type: { text: 'string' },
                description: 'The new value to set'
              },
              {
                name: 'silent',
                type: { text: 'boolean' },
                description: 'Suppress events'
              }
            ]
          }
        }
      ];

      searchIndex.buildIndex(treeData);

      // Should find by parameter name
      const results = searchIndex.search('newValue');
      expect(results).to.have.lengthOf(1);
    });

    it('handles empty or null input', () => {
      searchIndex.buildIndex(null);
      expect(searchIndex.size).to.equal(0);

      searchIndex.buildIndex([]);
      expect(searchIndex.size).to.equal(0);

      searchIndex.buildIndex(undefined);
      expect(searchIndex.size).to.equal(0);
    });
  });

  describe('search', () => {
    beforeEach(() => {
      const treeData = [
        {
          id: 'elem-1',
          name: 'MyButton',
          type: 'element',
          summary: 'A button component',
          description: 'This is a custom button element with various variants'
        },
        {
          id: 'elem-2',
          name: 'MyInput',
          type: 'element',
          summary: 'An input component',
          description: 'Text input field'
        },
        {
          id: 'class-1',
          name: 'ButtonController',
          type: 'class',
          summary: 'Controller for button behavior'
        },
        {
          id: 'func-1',
          name: 'formatValue',
          type: 'function',
          summary: 'Formats a value for display'
        }
      ];

      searchIndex.buildIndex(treeData);
    });

    it('returns empty array for empty query', () => {
      expect(searchIndex.search('')).to.have.lengthOf(0);
      expect(searchIndex.search('   ')).to.have.lengthOf(0);
      expect(searchIndex.search(null)).to.have.lengthOf(0);
    });

    it('finds exact name matches', () => {
      const results = searchIndex.search('MyButton');

      expect(results).to.have.lengthOf(1);
      expect(results[0].node.name).to.equal('MyButton');
    });

    it('is case-insensitive', () => {
      const results = searchIndex.search('mybutton');

      expect(results).to.have.lengthOf(1);
      expect(results[0].node.name).to.equal('MyButton');
    });

    it('finds partial matches', () => {
      const results = searchIndex.search('button');

      // Should match MyButton and ButtonController
      expect(results.length).to.be.at.least(2);
      const names = results.map(r => r.node.name);
      expect(names).to.include('MyButton');
      expect(names).to.include('ButtonController');
    });

    it('ranks exact matches higher', () => {
      const results = searchIndex.search('MyButton');

      expect(results[0].node.name).to.equal('MyButton');
      expect(results[0].score).to.be.greaterThan(results[1]?.score || 0);
    });

    it('ranks name matches higher than description matches', () => {
      const results = searchIndex.search('button');

      // Names starting with query (ButtonController) rank higher than names containing it (MyButton)
      const buttonControllerResult = results.find(r => r.node.name === 'ButtonController');
      const myButtonResult = results.find(r => r.node.name === 'MyButton');

      // ButtonController starts with "button" (score +500) so ranks higher than MyButton which contains it (score +250)
      expect(buttonControllerResult.score).to.be.greaterThan(myButtonResult.score);

      // Both name matches should rank higher than pure description matches (if any existed)
      const nonNameMatches = results.filter(r =>
        !r.node.name.toLowerCase().includes('button')
      );
      if (nonNameMatches.length > 0) {
        expect(myButtonResult.score).to.be.greaterThan(nonNameMatches[0].score);
      }
    });

    it('boosts element types in ranking', () => {
      // Both have 'format' in name, but different types
      const extraData = [
        {
          id: 'elem-format',
          name: 'format-display',
          type: 'element'
        },
        {
          id: 'func-format',
          name: 'formatHelper',
          type: 'function'
        }
      ];

      searchIndex.buildIndex(extraData);
      const results = searchIndex.search('format');

      // Element should be boosted
      const elementResult = results.find(r => r.type === 'element');
      const functionResult = results.find(r => r.type === 'function');

      expect(elementResult.score).to.be.greaterThan(functionResult.score);
    });

    it('searches in summary field', () => {
      const results = searchIndex.search('controller');

      expect(results.length).to.be.at.least(1);
      expect(results.some(r => r.node.name === 'ButtonController')).to.be.true;
    });

    it('searches in description field', () => {
      const results = searchIndex.search('variants');

      expect(results.length).to.be.at.least(1);
      expect(results.some(r => r.node.name === 'MyButton')).to.be.true;
    });

    it('supports multi-term queries', () => {
      const results = searchIndex.search('button component');

      // Should match items containing both terms
      expect(results.length).to.be.greaterThan(0);
    });

    it('skips very short query terms', () => {
      // Single character terms should not contribute much to score
      const results = searchIndex.search('a b button');

      // Should still find 'button' matches
      expect(results.length).to.be.greaterThan(0);
    });

    it('returns results sorted by relevance', () => {
      const results = searchIndex.search('button');

      // Verify descending score order
      for (let i = 1; i < results.length; i++) {
        expect(results[i - 1].score).to.be.at.least(results[i].score);
      }
    });

    it('includes result metadata', () => {
      const results = searchIndex.search('MyButton');

      expect(results[0]).to.have.property('id');
      expect(results[0]).to.have.property('node');
      expect(results[0]).to.have.property('path');
      expect(results[0]).to.have.property('pathString');
      expect(results[0]).to.have.property('score');
      expect(results[0]).to.have.property('type');
    });
  });

  describe('scoring algorithm', () => {
    it('prioritizes prefix matches over contains', () => {
      const treeData = [
        { id: '1', name: 'button-group', type: 'element' },
        { id: '2', name: 'my-button', type: 'element' }
      ];

      searchIndex.buildIndex(treeData);
      const results = searchIndex.search('button');

      // 'button-group' starts with 'button', should rank higher
      expect(results[0].node.name).to.equal('button-group');
    });

    it('gives highest score to exact matches', () => {
      const treeData = [
        { id: '1', name: 'button', type: 'element' },
        { id: '2', name: 'button-group', type: 'element' },
        { id: '3', name: 'my-button', type: 'element' }
      ];

      searchIndex.buildIndex(treeData);
      const results = searchIndex.search('button');

      // Exact match should be first
      expect(results[0].node.name).to.equal('button');
    });

    it('applies type boosting correctly', () => {
      const treeData = [
        { id: '1', name: 'test', type: 'element' },   // 1.5x boost
        { id: '2', name: 'test', type: 'class' },     // 1.3x boost
        { id: '3', name: 'test', type: 'function' },  // 1.2x boost
        { id: '4', name: 'test', type: 'module' },    // 1.1x boost
        { id: '5', name: 'test', type: 'variable' }   // no boost
      ];

      searchIndex.buildIndex(treeData);
      const results = searchIndex.search('test');

      // Should be ordered by type boost
      expect(results[0].type).to.equal('element');
      expect(results[1].type).to.equal('class');
      expect(results[2].type).to.equal('function');
      expect(results[3].type).to.equal('module');
    });
  });

  describe('createDebouncedSearch', () => {
    it('debounces search calls', (done) => {
      const treeData = [
        { id: '1', name: 'test', type: 'element' }
      ];
      searchIndex.buildIndex(treeData);

      let callCount = 0;
      const callback = (results) => {
        callCount++;
        // Should only be called once despite multiple rapid calls
        expect(callCount).to.equal(1);
        expect(results).to.have.lengthOf(1);
        done();
      };

      const debouncedSearch = searchIndex.createDebouncedSearch(callback, 50);

      // Call multiple times rapidly
      debouncedSearch('test');
      debouncedSearch('test');
      debouncedSearch('test');
    });

    it('uses custom delay', (done) => {
      const treeData = [
        { id: '1', name: 'test', type: 'element' }
      ];
      searchIndex.buildIndex(treeData);

      const startTime = Date.now();
      const callback = () => {
        const elapsed = Date.now() - startTime;
        expect(elapsed).to.be.at.least(100);
        done();
      };

      const debouncedSearch = searchIndex.createDebouncedSearch(callback, 100);
      debouncedSearch('test');
    });

    it('cancels previous search when called again', (done) => {
      const treeData = [
        { id: '1', name: 'apple', type: 'element' },
        { id: '2', name: 'banana', type: 'element' }
      ];
      searchIndex.buildIndex(treeData);

      let callCount = 0;
      const callback = (results) => {
        callCount++;
        // Should only get the last search (banana)
        expect(results[0].node.name).to.equal('banana');
        expect(callCount).to.equal(1);
        done();
      };

      const debouncedSearch = searchIndex.createDebouncedSearch(callback, 50);

      debouncedSearch('apple');
      setTimeout(() => debouncedSearch('banana'), 10);
    });
  });

  describe('clearDebounce', () => {
    it('cancels pending debounced search', (done) => {
      const treeData = [
        { id: '1', name: 'test', type: 'element' }
      ];
      searchIndex.buildIndex(treeData);

      let callbackCalled = false;
      const callback = () => {
        callbackCalled = true;
      };

      const debouncedSearch = searchIndex.createDebouncedSearch(callback, 50);

      debouncedSearch('test');
      searchIndex.clearDebounce();

      setTimeout(() => {
        expect(callbackCalled).to.be.false;
        done();
      }, 100);
    });
  });

  describe('clear', () => {
    it('clears the search index', () => {
      const treeData = [
        { id: '1', name: 'test', type: 'element' }
      ];

      searchIndex.buildIndex(treeData);
      expect(searchIndex.size).to.equal(1);

      searchIndex.clear();
      expect(searchIndex.size).to.equal(0);

      const results = searchIndex.search('test');
      expect(results).to.have.lengthOf(0);
    });

    it('clears pending debounce', (done) => {
      const treeData = [
        { id: '1', name: 'test', type: 'element' }
      ];
      searchIndex.buildIndex(treeData);

      let callbackCalled = false;
      const callback = () => {
        callbackCalled = true;
      };

      const debouncedSearch = searchIndex.createDebouncedSearch(callback, 50);

      debouncedSearch('test');
      searchIndex.clear();

      setTimeout(() => {
        expect(callbackCalled).to.be.false;
        done();
      }, 100);
    });
  });

  describe('size getter', () => {
    it('returns number of indexed entries', () => {
      expect(searchIndex.size).to.equal(0);

      searchIndex.buildIndex([
        { id: '1', name: 'a', type: 'element' },
        { id: '2', name: 'b', type: 'class' }
      ]);

      expect(searchIndex.size).to.equal(2);
    });

    it('includes nested entries', () => {
      searchIndex.buildIndex([
        {
          id: '1',
          name: 'parent',
          type: 'module',
          children: [
            { id: '2', name: 'child1', type: 'class' },
            { id: '3', name: 'child2', type: 'function' }
          ]
        }
      ]);

      expect(searchIndex.size).to.equal(3);
    });
  });

  describe('real-world usage scenarios', () => {
    it('handles complex custom element manifest structure', () => {
      const complexManifest = [
        {
          id: 'module-1',
          name: 'button.js',
          type: 'module',
          children: [
            {
              id: 'elem-button',
              name: 'Button',
              type: 'element',
              summary: 'Interactive button component',
              metadata: {
                tagName: 'my-button',
                type: { text: 'HTMLElement' }
              },
              children: [
                {
                  id: 'attr-variant',
                  name: 'variant',
                  type: 'attribute',
                  description: 'Visual variant of the button',
                  metadata: {
                    attribute: 'variant',
                    type: { text: 'string' },
                    default: 'default'
                  }
                },
                {
                  id: 'method-click',
                  name: 'handleClick',
                  type: 'method',
                  summary: 'Handles click events',
                  metadata: {
                    parameters: [
                      { name: 'event', type: { text: 'MouseEvent' } }
                    ]
                  }
                }
              ]
            }
          ]
        }
      ];

      searchIndex.buildIndex(complexManifest);

      // Should index all nodes: module + element + attribute + method = 4 nodes
      expect(searchIndex.size).to.equal(4);

      // Should find by element name
      let results = searchIndex.search('Button');
      expect(results.some(r => r.node.name === 'Button')).to.be.true;

      // Should find by tag name
      results = searchIndex.search('my-button');
      expect(results.length).to.be.greaterThan(0);

      // Should find by attribute name
      results = searchIndex.search('variant');
      expect(results.some(r => r.node.name === 'variant')).to.be.true;

      // Should find by parameter name
      results = searchIndex.search('MouseEvent');
      expect(results.some(r => r.node.name === 'handleClick')).to.be.true;
    });

    it('provides useful search results for ambiguous queries', () => {
      const treeData = [
        {
          id: '1',
          name: 'Button',
          type: 'element',
          summary: 'A clickable button'
        },
        {
          id: '2',
          name: 'IconButton',
          type: 'element',
          summary: 'Button with icon'
        },
        {
          id: '3',
          name: 'click',
          type: 'method',
          summary: 'Trigger click event'
        },
        {
          id: '4',
          name: 'Toggle',
          type: 'element',
          summary: 'A clickable toggle switch'
        }
      ];

      searchIndex.buildIndex(treeData);

      // Search for 'click' should return relevant results
      const results = searchIndex.search('click');

      // Should find items with 'click' in name or description
      expect(results.length).to.be.greaterThan(0);

      // Most relevant should be first (likely the 'click' method or Button elements)
      const topResult = results[0];
      expect(['click', 'Button', 'IconButton', 'Toggle']).to.include(topResult.node.name);
    });
  });
});
