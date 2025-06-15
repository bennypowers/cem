import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element step-progress
 * @description Shows progress in a multi-step process.
 */
@customElement('step-progress')
export class StepProgress extends LitElement {
  @property({ type: Number }) steps = 3;
  @property({ type: Number }) current = 1;

  static styles = css`
    .steps { display: flex; gap: 0.5em; align-items: center; }
    .step {
      width: 2em; height: 2em; border-radius: 50%;
      background: #ccc; text-align: center; line-height: 2em;
    }
    .step.active { background: #2196f3; color: white; }
  `;

  get completed() {
    return this.current >= this.steps;
  }

  render() {
    return html`
      <div class="steps">
        ${Array.from({ length: this.steps }, (_, i) => html`
          <span class="step ${i + 1 === this.current ? 'active' : ''}">${i + 1}</span>
        `)}
      </div>
      ${this.completed ? html`<span>Completed!</span>` : ''}
    `;
  }
}