// CEM Serve Knobs Manager - Manages knob groups for multi-instance demos

/**
 * Manages knob groups for multiple instances of the same custom element.
 * Watches the demo DOM for element additions/removals and dynamically creates/removes knob groups.
 *
 * Features:
 * - MutationObserver for dynamic element tracking
 * - Client-side label generation (ID → text → aria-label → fallback)
 * - Source order maintenance
 * - Knob event routing to correct element instance
 */
export class KnobsManager {
  #chrome = null;
  #demo = null;
  #tagName = '';
  #knobsContainer = null;
  #mo = null;
  #elementKnobsMap = new WeakMap(); // Maps element → knob group element
  #instanceCounter = 0;

  /**
   * @param {CemServeChrome} chrome - The chrome element
   * @param {string} tagName - Tag name to watch for
   */
  constructor(chrome, tagName) {
    this.#chrome = chrome;
    this.#tagName = tagName;
    this.#demo = chrome.querySelector('cem-serve-demo');

    if (!this.#demo) {
      console.warn('[KnobsManager] Demo element not found');
      return;
    }

    // Find knobs container in chrome's shadow root
    this.#knobsContainer = chrome.shadowRoot?.getElementById('knobs-container');
    if (!this.#knobsContainer) {
      console.warn('[KnobsManager] Knobs container not found');
      return;
    }
  }

  /**
   * Start watching for element changes
   */
  start() {
    if (!this.#demo || !this.#knobsContainer) return;

    // Discover existing elements
    this.#discoverExistingElements();

    // Set up MutationObserver to watch for element additions/removals
    this.#mo = new MutationObserver((mutations) => {
      this.#onDemoMutation(mutations);
    });

    // Watch the demo's shadow root or light DOM
    const root = this.#demo.shadowRoot ?? this.#demo;
    this.#mo.observe(root, {
      childList: true,
      subtree: true,
    });

    console.log('[KnobsManager] Started watching for', this.#tagName);
  }

  /**
   * Stop watching and clean up
   */
  stop() {
    if (this.#mo) {
      this.#mo.disconnect();
      this.#mo = null;
    }
  }

  /**
   * Discover and create knob groups for existing elements
   */
  #discoverExistingElements() {
    const root = this.#demo.shadowRoot ?? this.#demo;
    const elements = root.querySelectorAll(this.#tagName);

    elements.forEach((element, index) => {
      this.#addKnobGroupForElement(element, index);
    });
  }

  /**
   * Handle DOM mutations
   */
  #onDemoMutation(mutations) {
    for (const mutation of mutations) {
      // Handle removed nodes
      for (const node of mutation.removedNodes) {
        if (node.nodeType === Node.ELEMENT_NODE) {
          this.#handleRemovedNode(node);
        }
      }

      // Handle added nodes
      for (const node of mutation.addedNodes) {
        if (node.nodeType === Node.ELEMENT_NODE) {
          this.#handleAddedNode(node);
        }
      }
    }

    // After all mutations, update instance indices (for fallback labels)
    this.#updateInstanceIndices();
  }

  /**
   * Handle a node being removed from the demo
   */
  #handleRemovedNode(node) {
    // Check if the node itself matches
    if (node.matches && node.matches(this.#tagName)) {
      this.#removeKnobGroupForElement(node);
    }

    // Check if any descendants match
    if (node.querySelectorAll) {
      const descendants = node.querySelectorAll(this.#tagName);
      descendants.forEach(element => {
        this.#removeKnobGroupForElement(element);
      });
    }
  }

  /**
   * Handle a node being added to the demo
   */
  #handleAddedNode(node) {
    // Check if the node itself matches
    if (node.matches && node.matches(this.#tagName)) {
      // Find insertion index based on source order
      const index = this.#getElementIndex(node);
      this.#addKnobGroupForElement(node, index);
    }

    // Check if any descendants match
    if (node.querySelectorAll) {
      const descendants = node.querySelectorAll(this.#tagName);
      descendants.forEach(element => {
        const index = this.#getElementIndex(element);
        this.#addKnobGroupForElement(element, index);
      });
    }
  }

  /**
   * Get the index of an element among all instances in source order
   */
  #getElementIndex(targetElement) {
    const root = this.#demo.shadowRoot ?? this.#demo;
    const allElements = Array.from(root.querySelectorAll(this.#tagName));
    return allElements.indexOf(targetElement);
  }

  /**
   * Generate a label for an element instance using the priority system
   * 1. ID (#id)
   * 2. Text content (trimmed, max 20 chars)
   * 3. aria-label
   * 4. Fallback (tag-name No. N)
   */
  #generateElementLabel(element, index) {
    // Priority 1: ID
    if (element.id) {
      return `#${element.id}`;
    }

    // Priority 2: Text content (first non-empty text node, max 20 chars)
    const text = this.#firstTextContent(element);
    if (text) {
      if (text.length > 20) {
        return text.substring(0, 20).trim() + '…';
      }
      return text;
    }

    // Priority 3: aria-label
    const ariaLabel = element.getAttribute('aria-label');
    if (ariaLabel) {
      return ariaLabel;
    }

    // Priority 4: Fallback - "tag-name No. N" (1-indexed for display)
    return `${this.#tagName} No. ${index + 1}`;
  }

  /**
   * Extract first non-empty text node from element's subtree
   */
  #firstTextContent(element) {
    const walker = document.createTreeWalker(
      element,
      NodeFilter.SHOW_TEXT,
      {
        acceptNode(node) {
          const trimmed = node.textContent.trim();
          return trimmed ? NodeFilter.FILTER_ACCEPT : NodeFilter.FILTER_SKIP;
        }
      }
    );

    const node = walker.nextNode();
    return node ? node.textContent.trim() : '';
  }

  /**
   * Add a knob group for an element
   */
  #addKnobGroupForElement(element, index) {
    // Don't create duplicates
    if (this.#elementKnobsMap.has(element)) {
      return;
    }

    const label = this.#generateElementLabel(element, index);

    // Create knob group container (details element)
    const knobGroup = document.createElement('div');
    knobGroup.classList.add('knob-group-instance');
    knobGroup.dataset.tagName = this.#tagName;
    knobGroup.dataset.instanceIndex = index;

    // Create header with label
    const header = document.createElement('div');
    header.classList.add('knob-group-instance-header');
    header.innerHTML = `
      <span class="instance-tag">${this.#tagName}</span>:
      <span class="instance-label">${label}</span>
    `;

    knobGroup.appendChild(header);

    // Create knobs container for this instance
    const instanceKnobsContainer = document.createElement('div');
    instanceKnobsContainer.classList.add('instance-knobs');
    knobGroup.appendChild(instanceKnobsContainer);

    // Store mapping
    this.#elementKnobsMap.set(element, knobGroup);

    // Insert in source order
    this.#insertKnobGroupInOrder(knobGroup, index);

    // Set up event listeners for this group's knobs
    this.#setupKnobGroupListeners(knobGroup, element);

    console.log('[KnobsManager] Added knob group for', this.#tagName, 'at index', index, 'with label:', label);
  }

  /**
   * Insert a knob group at the correct position based on source order
   */
  #insertKnobGroupInOrder(knobGroup, targetIndex) {
    const existingGroups = Array.from(this.#knobsContainer.querySelectorAll('.knob-group-instance'));

    // Find insertion point - use >= to insert before existing group at same index
    let insertBefore = null;
    for (const group of existingGroups) {
      const groupIndex = parseInt(group.dataset.instanceIndex, 10);
      if (groupIndex >= targetIndex) {
        insertBefore = group;
        break;
      }
    }

    if (insertBefore) {
      this.#knobsContainer.insertBefore(knobGroup, insertBefore);
    } else {
      this.#knobsContainer.appendChild(knobGroup);
    }
  }

  /**
   * Remove a knob group for an element
   */
  #removeKnobGroupForElement(element) {
    const knobGroup = this.#elementKnobsMap.get(element);
    if (knobGroup && knobGroup.parentNode) {
      knobGroup.parentNode.removeChild(knobGroup);
      this.#elementKnobsMap.delete(element);
      console.log('[KnobsManager] Removed knob group for', this.#tagName);
    }
  }

  /**
   * Update instance indices after mutations
   */
  #updateInstanceIndices() {
    const root = this.#demo.shadowRoot ?? this.#demo;
    const elements = Array.from(root.querySelectorAll(this.#tagName));

    elements.forEach((element, index) => {
      const knobGroup = this.#elementKnobsMap.get(element);
      if (knobGroup) {
        knobGroup.dataset.instanceIndex = index;

        // Update label if it's a fallback label
        const labelSpan = knobGroup.querySelector('.instance-label');
        if (labelSpan && labelSpan.textContent.includes('No.')) {
          const newLabel = this.#generateElementLabel(element, index);
          labelSpan.textContent = newLabel;
        }
      }
    });
  }

  /**
   * Set up event listeners for a knob group
   */
  #setupKnobGroupListeners(knobGroup, targetElement) {
    // Listen for knob change events within this group
    knobGroup.addEventListener('knob:attribute-change', (e) => {
      e.stopPropagation(); // Prevent bubbling to chrome
      this.#handleAttributeChange(targetElement, e.name, e.value);
    });

    knobGroup.addEventListener('knob:property-change', (e) => {
      e.stopPropagation();
      this.#handlePropertyChange(targetElement, e.name, e.value);
    });

    knobGroup.addEventListener('knob:css-property-change', (e) => {
      e.stopPropagation();
      this.#handleCSSPropertyChange(targetElement, e.name, e.value);
    });
  }

  /**
   * Handle attribute change for a specific element
   */
  #handleAttributeChange(element, name, value) {
    // Handle boolean attributes (presence/absence)
    if (typeof value === 'boolean') {
      if (value) {
        element.setAttribute(name, '');
      } else {
        element.removeAttribute(name);
      }
    } else if (value === '' || value === null || value === undefined) {
      element.removeAttribute(name);
    } else {
      element.setAttribute(name, value);
    }

    console.log(`[KnobsManager] Attribute changed on`, element, `${name} = ${value}`);
  }

  /**
   * Handle property change for a specific element
   */
  #handlePropertyChange(element, name, value) {
    element[name] = value;
    console.log(`[KnobsManager] Property changed on`, element, `${name} = ${value}`);
  }

  /**
   * Handle CSS property change for a specific element
   */
  #handleCSSPropertyChange(element, name, value) {
    const propertyName = name.startsWith('--') ? name : `--${name}`;
    element.style.setProperty(propertyName, value);
    console.log(`[KnobsManager] CSS property changed on`, element, `${propertyName} = ${value}`);
  }
}
