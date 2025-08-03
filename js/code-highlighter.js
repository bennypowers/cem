import {
  translations
} from './config.js';

// SVG icon loader
function loadSvg(icon, parent) {
  parent.innerHTML = `<svg class="icon_${icon}"><use xlink:href="#${icon}"></use></svg>`;
}

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
  }

  #addLines(block) {
    let text = block.textContent;
    const snippet_fragment = [];
    if (text.includes('\n') && block.closest('pre') && !block.children.length) {
      text = text.split('\n');
      text.forEach((text_node, index) => {
        if(text_node.trim().length) {
          const new_node = `
          <span class="line line-flex">
            <span class="ln">${index + 1}</span>
            <span class="cl">${text_node.trim()}</span>
          </span>`.trim();
          snippet_fragment.push(new_node);
          block.closest('pre').className = 'chroma';
          block.classList.add('language-unknown');
          block.dataset.lang = translations.not_set;
        }
      });
      block.innerHTML = snippet_fragment.join('').trim(' ');
    }
  }


  #addActionPanel() {
    const actions = [
      { icon: 'copy', id: 'copy', title: translations.copy_text, show: true },
      { icon: 'order', id: 'lines', title: translations.toggle_line_numbers_text, show: true },
      { icon: 'carly', id: 'wrap', title: translations.toggle_line_wrap_text, show: false }
    ];

    const panel = document.createElement('div');
    panel.className = 'panel_box';

    actions.forEach(action => {
      const btn = document.createElement('a');
      btn.href = '#';
      btn.title = action.title;
      btn.className = `icon panel_icon panel_${action.id}`;
      if (!action.show) {
        btn.classList.add('panel_hide');
      }
      loadSvg(action.icon, btn);
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
    this.addEventListener('click', this.#handleClick.bind(this));
  }

  #handleClick(event) {
    const target = event.target;
    const action = this.#getActionType(target);
    if (!action) return;
    event.preventDefault();
    this.#showActiveState(target);
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

  #getActionType(target) {
    const classes = ['panel_copy', 'panel_wrap', 'panel_lines'];
    for (const cls of classes) {
      if (target.matches(`.${cls}`) || target.closest(`.${cls}`)) {
        return cls.replace('panel_', '');
      }
    }
    return null;
  }

  #showActiveState(target) {
    const element = target.matches('.icon') ? target : target.closest('.icon');
    if (!element) return;
    element.classList.remove('active');
    setTimeout(() => {
      element.classList.toggle('active');
    }, this.animationDelay);
  }

  #copyCode(codeElement) {
    const copyBtn = this.querySelector('.panel_copy');
    if (!copyBtn) return;

    const originalTitle = copyBtn.title;
    loadSvg('check', copyBtn);
    copyBtn.title = translations.copied_text;

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
      loadSvg('copy', copyBtn);
    }, this.panelTimeout);
  }

  #toggleLineWrap(codeElement) {
    codeElement.classList.toggle('pre_wrap');
  }

  #toggleLineNumbers(codeElement) {
    const lineNumbers = codeElement.querySelectorAll('.ln');
    lineNumbers.forEach(ln => ln.parentElement?.classList.toggle('pre_nolines'));
  }
}

