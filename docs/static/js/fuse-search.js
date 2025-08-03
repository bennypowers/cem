import Fuse from 'fuse.js';

export class FuseSearch extends HTMLElement {
  static #searchKeys = ['body', 'title', 'link', 'section', 'id'];

  static #minQueryLength = 2;

  static is = 'fuse-search'
  static { customElements.define(this.is, this); }

  #fuse = null;
  #searchInput = null;
  #resultsContainer = null;
  #form = null;
  #searchData = [];
  #selectedIndex = -1;
  #debounceTimer = null;

  connectedCallback() {
    this.#initializeElements();
    this.#setupEventListeners();
    this.#setupAccessibility();
    this.#loadSearchData();
  }

  #initializeElements() {
    const shadowRoot = this.shadowRoot;
    this.#searchInput = shadowRoot.querySelector('input[type="search"]');
    this.#resultsContainer = shadowRoot.querySelector('[role="listbox"]');
    this.#form = shadowRoot.querySelector('form');
  }

  #setupAccessibility() {
    // Set up ARIA attributes for combobox pattern
    this.#searchInput.setAttribute('role', 'combobox');
    this.#searchInput.setAttribute('aria-expanded', 'false');
    this.#searchInput.setAttribute('aria-haspopup', 'listbox');
    this.#searchInput.setAttribute('aria-autocomplete', 'list');
    // Generate unique ID for the listbox
    const listboxId = `search-listbox-${Math.random().toString(36).substr(2, 9)}`;
    this.#resultsContainer.id = listboxId;
    this.#searchInput.setAttribute('aria-controls', listboxId);
  }

  #setupEventListeners() {
    this.#form.addEventListener('submit', this.#onFormSubmit);
    this.#searchInput.addEventListener('input', this.#onSearchInput);
    this.#searchInput.addEventListener('keydown', this.#onKeyDown);
    document.addEventListener('click', this.#onDocumentClick);

    // Clear results on escape
    document.addEventListener('keydown', (event) => {
      if (event.code === 'Escape') {
        this.#clearResults();
        this.#searchInput.blur();
      }
    });
  }

  #onFormSubmit = (event) => {
    if (this.#fuse && this.#searchInput.value.trim()) {
      // Prevent default form submission if Fuse is available
      event.preventDefault();
      return;
    }
    // Allow default form submission to DuckDuckGo if Fuse not available
  }

  #onSearchInput = (event) => {
    const query = event.target.value.trim();

    // Clear existing debounce timer
    if (this.#debounceTimer) {
      clearTimeout(this.#debounceTimer);
    }

    // Debounce search to avoid excessive API calls
    this.#debounceTimer = setTimeout(() => {
      this.#performSearch(query);
    }, 150);
  }

  #onKeyDown = (event) => {
    const results = this.#resultsContainer.querySelectorAll('a[role="option"]');
    // Only handle navigation if there are results
    if (results.length === 0) return;

    switch (event.code) {
      case 'ArrowDown':
        event.preventDefault();
        if (this.#selectedIndex < results.length - 1) {
          this.#selectedIndex++;
        } else {
          this.#selectedIndex = 0; // Wrap to first
        }
        this.#updateSelection(results);
        break;

      case 'ArrowUp':
        event.preventDefault();
        if (this.#selectedIndex > 0) {
          this.#selectedIndex--;
        } else if (this.#selectedIndex === 0) {
          this.#selectedIndex = results.length - 1; // Wrap to last
        } else {
          this.#selectedIndex = results.length - 1; // Start from last if none selected
        }
        this.#updateSelection(results);
        break;

      case 'Enter':
        event.preventDefault();
        if (this.#selectedIndex >= 0 && results[this.#selectedIndex]) {
          results[this.#selectedIndex].click();
        }
        break;

      case 'Escape':
        event.preventDefault();
        this.#clearResults();
        this.#searchInput.blur();
        break;
    }
  }

  #onDocumentClick = (event) => {
    if (!this.contains(event.target)) {
      this.#clearResults();
    }
  }

  #updateSelection(results) {
    // Remove previous active descendant
    this.#searchInput.removeAttribute('aria-activedescendant');
    results.forEach((result, index) => {
      const isSelected = index === this.#selectedIndex;
      result.classList.toggle('active', isSelected);
      result.setAttribute('aria-selected', isSelected.toString());
      // Set active descendant
      if (isSelected) {
        this.#searchInput.setAttribute('aria-activedescendant', result.id);
      }
    });
  }

  async #loadSearchData() {
    try {
      const pageLanguage = document.documentElement.lang || 'en';
      const searchIndex = pageLanguage === 'en' ? 'index.json' : `${pageLanguage}/index.json`;
      const baseUrl = this.getAttribute('base-url') || window.location.origin;

      const response = await fetch(new URL(searchIndex, baseUrl).href);

      if (!response.ok) {
        throw new Error(`Failed to load search index: ${response.status}`);
      }

      this.#searchData = await response.json();

      // Initialize Fuse with the loaded data
      this.#fuse = new Fuse(this.#searchData, {
        ignoreLocation: true,
        findAllMatches: true,
        includeScore: true,
        shouldSort: true,
        keys: FuseSearch.#searchKeys,
        threshold: 0.1
      });

    } catch (error) {
      console.warn('Search index could not be loaded, falling back to external search:', error);
      // Fuse search will remain unavailable, form will submit to DuckDuckGo
    }
  }

  #performSearch(query) {
    if (!query || query.length < FuseSearch.#minQueryLength) {
      this.#clearResults();
      return;
    }

    if (!this.#fuse) {
      // Show loading state or fallback message
      this.#showMessage('Search index loading...');
      return;
    }

    const results = this.#fuse.search(query);
    this.#displayResults(results, query);
  }

  #displayResults(results, query) {
    this.#clearResults();
    this.#selectedIndex = -1;

    if (results.length === 0) {
      this.#showMessage('No results found');
      return;
    }

    // Limit results for performance
    const limitedResults = results.slice(0, 8);

    limitedResults.forEach(({ item, score }, index) => {
      const resultElement = this.#createResultElement(item, query, score, index);
      this.#resultsContainer.appendChild(resultElement);
    });

    this.#resultsContainer.classList.add('active');
    this.#searchInput.setAttribute('aria-expanded', 'true');
  }

  #createResultElement(item, query, _score, index) {
    const link = document.createElement('a');
    link.href = `${item.link}?query=${encodeURIComponent(query)}`;
    
    // Set up accessibility attributes
    link.setAttribute('role', 'option');
    link.setAttribute('aria-selected', 'false');
    link.id = `search-result-${index}`;

    const title = document.createElement('h3');
    title.textContent = item.title;
    link.appendChild(title);

    // Add description if available
    if (item.body) {
      const description = document.createElement('p');
      const queryIndex = item.body.toLowerCase().indexOf(query.toLowerCase());
      const start = Math.max(0, queryIndex - 50);
      const end = Math.min(item.body.length, queryIndex + 150);
      const excerpt = item.body.substring(start, end);
      description.textContent = start > 0 ? `...${excerpt}` : excerpt;
      link.appendChild(description);
    }

    return link;
  }

  #showMessage(message) {
    this.#clearResults();

    const messageElement = document.createElement('div');
    messageElement.setAttribute('role', 'status');
    messageElement.textContent = message;

    this.#resultsContainer.appendChild(messageElement);
    this.#resultsContainer.classList.add('active');
  }

  #clearResults() {
    this.#resultsContainer.innerHTML = '';
    this.#resultsContainer.classList.remove('active');
    this.#selectedIndex = -1;
    this.#searchInput.setAttribute('aria-expanded', 'false');
    this.#searchInput.removeAttribute('aria-activedescendant');
  }
}
