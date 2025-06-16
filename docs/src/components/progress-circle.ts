import { LitElement, html, property, customElement, css } from 'lit-element';

@customElement('progress-circle')
export class ProgressCircle extends LitElement {
  @property({ type: Number }) progress = 0;

  static styles = css`
    .circle {
      width: 48px;
      height: 48px;
      border-radius: 50%;
      border: 4px solid #eee;
      border-top-color: #2196f3;
      animation: spin 1s linear infinite;
    }
    @keyframes spin {
      to { transform: rotate(360deg); }
    }
  `;

  render() {
    return html`
      <div class="circle" style="border-top-color: ${this._color()}"></div>
      <span>${this.progress}%</span>
    `;
  }

  private _color() {
    return this.progress > 50 ? '#4caf50' : '#2196f3';
  }
}
