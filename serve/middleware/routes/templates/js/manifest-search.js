/**
 * Search index and search functionality for Custom Elements Manifest tree
 */
export class ManifestSearchIndex {
  #index = [];
  #debounceTimer = null;

  /**
   * Build search index from tree data
   * @param {Array<TreeNode>} treeData - Tree node data structure
   */
  buildIndex(treeData) {
    this.#index = [];
    this.#indexNodes(treeData, []);
  }

  /**
   * Recursively index all nodes
   * @param {Array<TreeNode>} nodes - Nodes to index
   * @param {Array<string>} path - Current path in tree
   */
  #indexNodes(nodes, path) {
    if (!nodes || !Array.isArray(nodes)) return;

    nodes.forEach(node => {
      const currentPath = [...path, node.name || ''];

      // Create searchable text from all node properties
      const searchableText = this.#buildSearchableText(node);

      // Add to index
      this.#index.push({
        id: node.id,
        node,
        path: currentPath,
        pathString: currentPath.join(' > '),
        searchableText: searchableText.toLowerCase(),
        type: node.type,
        name: (node.name || '').toLowerCase(),
        summary: (node.summary || '').toLowerCase(),
        description: (node.description || '').toLowerCase(),
      });

      // Recursively index children
      if (node.children && node.children.length > 0) {
        this.#indexNodes(node.children, currentPath);
      }
    });
  }

  /**
   * Build searchable text string from node
   * @param {TreeNode} node - Node to extract text from
   * @returns {string} Combined searchable text
   */
  #buildSearchableText(node) {
    const parts = [];

    // Add basic properties
    if (node.name) parts.push(node.name);
    if (node.summary) parts.push(node.summary);
    if (node.description) parts.push(node.description);
    if (node.type) parts.push(node.type);

    // Add metadata properties
    if (node.metadata) {
      const meta = node.metadata;

      if (meta.tagName) parts.push(meta.tagName);
      if (meta.type?.text) parts.push(meta.type.text);
      if (meta.default) parts.push(String(meta.default));
      if (meta.fieldName) parts.push(meta.fieldName);
      if (meta.attribute) parts.push(meta.attribute);
      if (meta.syntax) parts.push(meta.syntax);

      // Add parameter info for methods/functions
      if (meta.parameters && Array.isArray(meta.parameters)) {
        meta.parameters.forEach(param => {
          if (param.name) parts.push(param.name);
          if (param.type?.text) parts.push(param.type.text);
          if (param.description) parts.push(param.description);
        });
      }

      // Add return info
      if (meta.return) {
        if (meta.return.type?.text) parts.push(meta.return.type.text);
        if (meta.return.description) parts.push(meta.return.description);
      }
    }

    return parts.join(' ');
  }

  /**
   * Search the index
   * @param {string} query - Search query
   * @returns {Array<SearchResult>} Matching results sorted by relevance
   */
  search(query) {
    if (!query || query.trim().length === 0) {
      return [];
    }

    const normalizedQuery = query.toLowerCase().trim();
    const queryTerms = normalizedQuery.split(/\s+/);

    const results = [];

    this.#index.forEach(entry => {
      const score = this.#calculateScore(entry, normalizedQuery, queryTerms);

      if (score > 0) {
        results.push({
          id: entry.id,
          node: entry.node,
          path: entry.path,
          pathString: entry.pathString,
          score,
          type: entry.type,
        });
      }
    });

    // Sort by score (descending)
    results.sort((a, b) => b.score - a.score);

    return results;
  }

  /**
   * Calculate relevance score for a search entry
   * @param {Object} entry - Index entry
   * @param {string} query - Normalized query
   * @param {Array<string>} queryTerms - Individual query terms
   * @returns {number} Relevance score (0 if no match)
   */
  #calculateScore(entry, query, queryTerms) {
    let score = 0;

    // Exact name match (highest score)
    if (entry.name === query) {
      score += 1000;
    }
    // Name starts with query
    else if (entry.name.startsWith(query)) {
      score += 500;
    }
    // Name contains query
    else if (entry.name.includes(query)) {
      score += 250;
    }

    // Type exact match
    if (entry.type === query) {
      score += 100;
    }

    // Summary starts with query
    if (entry.summary.startsWith(query)) {
      score += 75;
    }
    // Summary contains query
    else if (entry.summary.includes(query)) {
      score += 40;
    }

    // Description contains query
    if (entry.description.includes(query)) {
      score += 20;
    }

    // General searchable text contains query
    if (entry.searchableText.includes(query)) {
      score += 10;
    }

    // Score for each query term match
    queryTerms.forEach(term => {
      if (term.length < 2) return; // Skip very short terms

      // Name contains term
      if (entry.name.includes(term)) {
        score += 15;
      }

      // Summary contains term
      if (entry.summary.includes(term)) {
        score += 8;
      }

      // Description contains term
      if (entry.description.includes(term)) {
        score += 4;
      }

      // Type contains term
      if (entry.type.includes(term)) {
        score += 5;
      }
    });

    // Boost certain types
    switch (entry.type) {
      case 'element':
        score *= 1.5;
        break;
      case 'class':
        score *= 1.3;
        break;
      case 'function':
        score *= 1.2;
        break;
      case 'module':
        score *= 1.1;
        break;
    }

    return score;
  }

  /**
   * Get debounced search function
   * @param {Function} callback - Callback to invoke with search results
   * @param {number} delay - Debounce delay in milliseconds (default: 300)
   * @returns {Function} Debounced search function
   */
  createDebouncedSearch(callback, delay = 300) {
    return (query) => {
      if (this.#debounceTimer) {
        clearTimeout(this.#debounceTimer);
      }

      this.#debounceTimer = setTimeout(() => {
        const results = this.search(query);
        callback(results);
      }, delay);
    };
  }

  /**
   * Clear any pending debounced search
   */
  clearDebounce() {
    if (this.#debounceTimer) {
      clearTimeout(this.#debounceTimer);
      this.#debounceTimer = null;
    }
  }

  /**
   * Get total number of indexed entries
   * @returns {number} Number of indexed nodes
   */
  get size() {
    return this.#index.length;
  }

  /**
   * Clear the search index
   */
  clear() {
    this.#index = [];
    this.clearDebounce();
  }
}
