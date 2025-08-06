export class CodeHighlighter extends HTMLElement {
  static is = 'code-highlighter';
  static { customElements.define(this.is, this); }

  // Get CSS custom properties for better maintainability
  get maxHeight() {
    return getComputedStyle(document.documentElement).getPropertyValue('--code-max-height').trim() || 'initial';
  }

  get panelTimeout() {
    return parseInt(getComputedStyle(document.documentElement).getPropertyValue('--code-panel-timeout')) || 2250;
  }

  get animationDelay() {
    return parseInt(getComputedStyle(document.documentElement).getPropertyValue('--code-panel-animation-delay')) || 50;
  }

  connectedCallback() {
    this.#initializeCodeBlock();
    this.#setupEventListeners();
  }

  #initializeCodeBlock() {
    const codeElement = this.querySelector('code');
    if (!codeElement) return;

    this.#addLines(codeElement);
    this.#addActionPanel();
    this.#addLanguageLabel(codeElement);
    // Ensure accessibility attributes are set for all pre elements
    const preElement = codeElement.closest('pre');
    if (preElement) {
      this.#removeTabindex(preElement);
    }
  }

  #addLines(block) {
    let text = block.textContent;
    const snippetFragment = [];
    if (text.includes('\n') && block.closest('pre') && !block.children.length) {
      text = text.split('\n');
      text.forEach((textNode, index) => {
        if(textNode.trim().length) {
          const newNode = `
          <span class="line line-flex">
            <span class="ln">${index + 1}</span>
            <span class="cl">${textNode.trim()}</span>
          </span>`.trim();
          snippetFragment.push(newNode);
          const preElement = block.closest('pre');
          preElement.className = 'chroma';
          this.#removeTabindex(preElement);
          block.classList.add('language-unknown');
          block.dataset.lang = 'not set';
        }
      });
      block.innerHTML = snippetFragment.join('').trim(' ');
    }
  }


  #addActionPanel() {
    const actions = [
      { icon: 'copy', id: 'copy', title: 'Copy', show: true },
    ];

    const panel = document.createElement('div');
    panel.className = 'panel-box';

    actions.forEach(action => {
      const btn = document.createElement('button');
      btn.type = 'button';
      btn.title = action.title;
      btn.classList.add('panel-icon', `panel-${action.id}`);
      if (!action.show) {
        btn.classList.add('panel-hide');
      }

      // Add both icon and text label
      const iconContainer = document.createElement('span');
      iconContainer.className = 'panel-icon-svg';
      iconContainer.innerHTML = `<svg class="icon-${action.icon}" role="presentation"><use xlink:href="#${action.icon}"></use></svg>`

      const textLabel = document.createElement('span');
      textLabel.className = 'panel-icon-text';
      textLabel.textContent = action.title;

      btn.appendChild(iconContainer);
      btn.appendChild(textLabel);

      // Add click listener directly to each button
      btn.addEventListener('click', (event) => {
        event.preventDefault();
        this.#handleButtonClick(action.id, btn);
      });

      panel.appendChild(btn);
    });

    this.appendChild(panel);
  }


  #addLanguageLabel(block) {
    let label = block.dataset.lang;
    const shellBased = ['sh', 'shell', 'zsh', 'bash'];
    if (shellBased.includes(label)) {
      // Add shell class to the code block for CSS styling
      block.classList.add('language-shell');
      // Clean up shell prompts - remove $ from all lines since CSS will add it to first line
      const lines = block.querySelectorAll('.line .cl');
      Array.from(lines).forEach(line => {
        let content = line.textContent;
        // Remove leading $ and optional space
        if (content.startsWith('$ ')) {
          line.textContent = content.substring(2);
        } else if (content.startsWith('$')) {
          line.textContent = content.substring(1);
        }
      });
    }

    label = label === 'sh' ? 'shell' : label;
    if (label !== "fallback") {
      const labelEl = document.createElement('div');
      labelEl.textContent = label;
      labelEl.classList.add('lang');
      this.appendChild(labelEl);
    }
  }

  #setupEventListeners() {
    // Event listeners are now attached directly to individual buttons in #addActionPanel
  }

  #handleButtonClick(action, buttonElement) {
    this.#showActiveState(buttonElement);
    const codeElement = this.querySelector('code');
    if (!codeElement) return;
    
    switch (action) {
      case 'copy':
        this.#copyCode(codeElement);
        break;
      case 'wrap':
        this.#toggleLineWrap(codeElement);
        break;
      case 'lines':
        this.#toggleLineNumbers(codeElement);
        break;
    }
  }


  #showActiveState(buttonElement) {
    buttonElement.classList.remove('active');
    setTimeout(() => {
      buttonElement.classList.add('active');
    }, this.animationDelay);
  }

  #copyCode(codeElement) {
    const copyBtn = this.querySelector('.panel-copy');
    if (!copyBtn) return;

    const originalTitle = copyBtn.title;
    const iconContainer = copyBtn.querySelector('.panel-icon-svg');
    iconContainer.innerHTML = `<svg class="icon-check" role="presentation"><use xlink:href="#${icon}"></use></svg>`
    copyBtn.title = 'Copied!';
    const textLabel = copyBtn.querySelector('.panel-icon-text');
    textLabel.textContent = 'Copied!';

    // Clone and clean up code for copying
    const clonedCode = codeElement.cloneNode(true);
    const lineNumbers = clonedCode.querySelectorAll('.ln');
    lineNumbers.forEach(ln => ln.remove());

    // Remove leading '$' from shell snippets
    const spans = clonedCode.querySelectorAll('span');
    spans.forEach(span => {
      const text = span.textContent.trim(' ');
      if (text.indexOf('$') === 0) {
        span.textContent = span.textContent.replace('$ ', '');
      }
    });

    const snippet = clonedCode.textContent.trim(' ');
    navigator.clipboard.writeText(snippet);

    setTimeout(() => {
      copyBtn.title = originalTitle;
      iconContainer.innerHTML = `<svg class="icon-copy" role="presentation"><use xlink:href="#${icon}"></use></svg>`
      textLabel.textContent = originalTitle;
    }, this.panelTimeout);
  }

  #toggleLineWrap(codeElement) {
    codeElement.classList.toggle('pre-wrap');
  }

  #toggleLineNumbers(codeElement) {
    const lineNumbers = codeElement.querySelectorAll('.ln');
    lineNumbers.forEach(ln => ln.parentElement?.classList.toggle('pre-nolines'));
  }

  #removeTabindex(preElement) {
    // Remove tabindex from pre elements since code blocks don't need to be focusable
    preElement.removeAttribute('tabindex');
  }
}

