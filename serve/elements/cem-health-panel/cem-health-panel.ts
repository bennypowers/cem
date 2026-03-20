import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { state } from 'lit/decorators/state.js';

import styles from './cem-health-panel.css' with { type: 'css' };

import '../pf-v6-label/pf-v6-label.js';
import '../pf-v6-expandable-section/pf-v6-expandable-section.js';

const STATUS_COLORS: Record<string, string> = {
  pass: 'green',
  warn: 'orange',
  fail: 'red',
};

interface HealthFinding {
  message?: string;
}

interface HealthCategory {
  category: string;
  points: number;
  maxPoints: number;
  status: string;
  findings?: HealthFinding[];
}

interface HealthDeclaration {
  tagName?: string;
  name?: string;
  score: number;
  maxScore: number;
  categories: HealthCategory[];
}

interface HealthModule {
  declarations: HealthDeclaration[];
}

interface HealthResult {
  modules: HealthModule[];
  recommendations?: string[];
}

/**
 * Health panel showing documentation quality scores
 *
 * @attr {string} component - Component tag name or class name to display health for
 */
@customElement('cem-health-panel')
export class CemHealthPanel extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor component: string | null = null;

  @state()
  accessor #loading = true;

  @state()
  accessor #result: HealthResult | null = null;

  #abortController: AbortController | null = null;

  connectedCallback() {
    super.connectedCallback();
    this.fetchHealth();
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.#abortController?.abort();
  }

  async fetchHealth() {
    this.#abortController?.abort();
    this.#abortController = new AbortController();

    this.#loading = true;
    this.#result = null;

    const url = new URL('/__cem/api/health', location.origin);
    if (this.component) {
      url.searchParams.set('component', this.component);
    }

    try {
      const response = await fetch(url, { signal: this.#abortController.signal });
      if (!response.ok) throw new Error(`HTTP ${response.status}`);
      this.#result = await response.json();
    } catch (e) {
      if ((e as Error).name === 'AbortError') return;
      console.warn('cem-health-panel: failed to fetch health data', e);
    } finally {
      this.#loading = false;
    }
  }

  render() {
    if (this.#loading) {
      return html`<p>Analyzing documentation health...</p>`;
    }

    const decl = this.#findDeclaration();
    if (!decl) {
      return html`<p>No health data available.</p>`;
    }

    const pct = decl.maxScore > 0
      ? Math.round((decl.score / decl.maxScore) * 100)
      : 0;
    const overallStatus = pct >= 80 ? 'pass' : pct >= 40 ? 'warn' : 'fail';

    return html`
      <div id="overall">
        <pf-v6-label color=${STATUS_COLORS[overallStatus]}
                     size="lg">${pct}% -- ${decl.score}/${decl.maxScore}</pf-v6-label>
      </div>
      <dl id="categories"
          class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        ${decl.categories.map(cat => this.#renderCategory(cat))}
      </dl>
      ${this.#renderRecommendations()}
    `;
  }

  #findDeclaration(): HealthDeclaration | null {
    if (!this.#result) return null;
    for (const mod of this.#result.modules) {
      for (const decl of mod.declarations) {
        if (!this.component
            || decl.tagName === this.component
            || decl.name === this.component) {
          return decl;
        }
      }
    }
    return null;
  }

  #renderCategory(cat: HealthCategory) {
    const pct = cat.maxPoints > 0
      ? Math.round((cat.points / cat.maxPoints) * 100)
      : 0;
    const findings = cat.findings?.filter(f => f.message) ?? [];

    return html`
      <div class="pf-v6-c-description-list__group">
        <dt class="pf-v6-c-description-list__term">${cat.category}</dt>
        <dd class="pf-v6-c-description-list__description">
          <div class="category-bar">
            <div class="category-meter">
              <div class="category-fill fill-${cat.status}"
                   style="width: ${pct}%"></div>
            </div>
            <span class="category-score">${cat.points}/${cat.maxPoints}</span>
            <pf-v6-label color=${STATUS_COLORS[cat.status]}>${cat.status}</pf-v6-label>
          </div>
          ${findings.length > 0 ? html`
            <div class="finding-details">
              <ul>
                ${findings.map(f => html`<li>${f.message}</li>`)}
              </ul>
            </div>
          ` : nothing}
        </dd>
      </div>
    `;
  }

  #renderRecommendations() {
    const recs = this.#result?.recommendations;
    if (!recs?.length) return nothing;
    return html`
      <div id="recommendations">
        <h4>Recommendations</h4>
        <ul>
          ${recs.map(rec => html`<li>${rec}</li>`)}
        </ul>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-health-panel': CemHealthPanel;
  }
}
