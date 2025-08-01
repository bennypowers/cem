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
    this.#loadSearchData();
  }

  #initializeElements() {
    const shadowRoot = this.shadowRoot;
    this.#searchInput = shadowRoot.querySelector('input[type="search"]');
    this.#resultsContainer = shadowRoot.querySelector('[role="listbox"]');
    this.#form = shadowRoot.querySelector('form');
  }

  #setupEventListeners() {
    // Prevent form submission and use Fuse search instead
    this.#form.addEventListener('submit', this.#handleFormSubmit.bind(this));

    // Handle search input
    this.#searchInput.addEventListener('input', this.#handleSearchInput.bind(this));

    // Handle keyboard navigation
    this.#searchInput.addEventListener('keydown', this.#handleKeyDown.bind(this));

    // Clear results when clicking outside
    document.addEventListener('click', this.#handleDocumentClick.bind(this));

    // Clear results on escape
    document.addEventListener('keydown', (event) => {
      if (event.code === 'Escape') {
        this.#clearResults();
        this.#searchInput.blur();
      }
    });
  }

  #handleFormSubmit(event) {
    if (this.#fuse && this.#searchInput.value.trim()) {
      // Prevent default form submission if Fuse is available
      event.preventDefault();
      return;
    }
    // Allow default form submission to DuckDuckGo if Fuse not available
  }

  #handleSearchInput(event) {
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

  #handleKeyDown(event) {
    const results = this.#resultsContainer.querySelectorAll('a');

    switch (event.code) {
      case 'ArrowDown':
        event.preventDefault();
        this.#selectedIndex = Math.min(this.#selectedIndex + 1, results.length - 1);
        this.#updateSelection(results);
        break;

      case 'ArrowUp':
        event.preventDefault();
        this.#selectedIndex = Math.max(this.#selectedIndex - 1, -1);
        this.#updateSelection(results);
        break;

      case 'Enter':
        event.preventDefault();
        if (this.#selectedIndex >= 0 && results[this.#selectedIndex]) {
          window.location.href = results[this.#selectedIndex].href;
        }
        break;

      case 'Escape':
        this.#clearResults();
        this.#searchInput.blur();
        break;
    }
  }

  #handleDocumentClick(event) {
    if (!this.contains(event.target)) {
      this.#clearResults();
    }
  }

  #updateSelection(results) {
    results.forEach((result, index) => {
      result.classList.toggle('active', index === this.#selectedIndex);
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

    limitedResults.forEach(({ item, score }) => {
      const resultElement = this.#createResultElement(item, query, score);
      this.#resultsContainer.appendChild(resultElement);
    });

    this.#resultsContainer.classList.add('active');
  }

  #createResultElement(item, query) {
    const link = document.createElement('a');
    link.href = `${item.link}?query=${encodeURIComponent(query)}`;

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
  }
}
