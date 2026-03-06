import '/__cem/elements/pf-v6-label/pf-v6-label.js';
import '/__cem/elements/pf-v6-expandable-section/pf-v6-expandable-section.js';
import { CemElement } from '/__cem/cem-element.js';

const STATUS_COLORS = {
  pass: 'green',
  warn: 'orange',
  fail: 'red',
};

const categoryTemplate = document.createElement('template');
categoryTemplate.innerHTML = `
  <div class="pf-v6-c-description-list__group">
    <dt class="pf-v6-c-description-list__term"></dt>
    <dd class="pf-v6-c-description-list__description">
      <div class="category-bar">
        <div class="category-meter">
          <div class="category-fill"></div>
        </div>
        <span class="category-score"></span>
        <pf-v6-label></pf-v6-label>
      </div>
      <div class="finding-details" hidden>
        <ul></ul>
      </div>
    </dd>
  </div>
`;

/**
 * Health panel showing documentation quality scores
 * @customElement cem-health-panel
 */
export class CemHealthPanel extends CemElement {
  static is = 'cem-health-panel';

  #loading;
  #empty;
  #content;
  #overallLabel;
  #categories;
  #recommendations;
  #recommendationsList;
  #component = null;
  #abortController = null;

  afterTemplateLoaded() {
    this.#loading = this.shadowRoot.getElementById('loading');
    this.#empty = this.shadowRoot.getElementById('empty');
    this.#content = this.shadowRoot.getElementById('content');
    this.#overallLabel = this.shadowRoot.getElementById('overall-label');
    this.#categories = this.shadowRoot.getElementById('categories');
    this.#recommendations = this.shadowRoot.getElementById('recommendations');
    this.#recommendationsList = this.shadowRoot.getElementById('recommendations-list');

    // Read component from host attribute
    this.#component = this.getAttribute('component');

    // Fetch on connect
    this.fetchHealth();
  }

  disconnectedCallback() {
    this.#abortController?.abort();
  }

  async fetchHealth() {
    this.#abortController?.abort();
    this.#abortController = new AbortController();

    this.#loading.hidden = false;
    this.#empty.hidden = true;
    this.#content.hidden = true;

    const url = new URL('/__cem/api/health', location.origin);
    if (this.#component) {
      url.searchParams.set('component', this.#component);
    }

    try {
      const response = await fetch(url, { signal: this.#abortController.signal });
      if (!response.ok) throw new Error(`HTTP ${response.status}`);
      const result = await response.json();
      this.#render(result);
    } catch (e) {
      if (e.name === 'AbortError') return;
      console.warn('cem-health-panel: failed to fetch health data', e);
      this.#loading.hidden = true;
      this.#empty.hidden = false;
    }
  }

  #render(result) {
    this.#loading.hidden = true;

    // Find the declaration to display
    const decl = this.#findDeclaration(result);
    if (!decl) {
      this.#empty.hidden = false;
      this.#content.hidden = true;
      return;
    }

    this.#empty.hidden = true;
    this.#content.hidden = false;

    // Overall score
    const pct = decl.maxScore > 0
      ? Math.round((decl.score / decl.maxScore) * 100)
      : 0;
    const overallStatus = pct >= 80 ? 'pass' : pct >= 40 ? 'warn' : 'fail';
    this.#overallLabel.textContent = `${pct}% — ${decl.score}/${decl.maxScore}`;
    this.#overallLabel.setAttribute('color', STATUS_COLORS[overallStatus]);

    // Categories
    this.#categories.replaceChildren();
    for (const cat of decl.categories) {
      this.#categories.append(this.#renderCategory(cat));
    }

    // Recommendations
    if (result.recommendations?.length > 0) {
      this.#recommendations.hidden = false;
      this.#recommendationsList.replaceChildren();
      for (const rec of result.recommendations) {
        const li = document.createElement('li');
        li.textContent = rec;
        this.#recommendationsList.append(li);
      }
    } else {
      this.#recommendations.hidden = true;
    }
  }

  #findDeclaration(result) {
    for (const mod of result.modules) {
      for (const decl of mod.declarations) {
        if (!this.#component || decl.tagName === this.#component || decl.name === this.#component) {
          return decl;
        }
      }
    }
    return null;
  }

  #renderCategory(cat) {
    const fragment = categoryTemplate.content.cloneNode(true);

    const dt = fragment.querySelector('dt');
    dt.textContent = cat.category;

    const fill = fragment.querySelector('.category-fill');
    const pct = cat.maxPoints > 0
      ? Math.round((cat.points / cat.maxPoints) * 100)
      : 0;
    fill.classList.add(`fill-${cat.status}`);
    fill.style.width = `${pct}%`;

    const score = fragment.querySelector('.category-score');
    score.textContent = `${cat.points}/${cat.maxPoints}`;

    const label = fragment.querySelector('pf-v6-label');
    label.setAttribute('color', STATUS_COLORS[cat.status]);
    label.textContent = cat.status;

    // Add findings if any have messages
    const findings = cat.findings?.filter(f => f.message);
    if (findings?.length > 0) {
      const details = fragment.querySelector('.finding-details');
      details.hidden = false;
      const ul = details.querySelector('ul');
      for (const finding of findings) {
        const li = document.createElement('li');
        li.textContent = finding.message;
        ul.append(li);
      }
    }

    return fragment;
  }
}

customElements.define(CemHealthPanel.is, CemHealthPanel);
