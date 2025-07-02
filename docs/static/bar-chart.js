class BarChart extends HTMLElement {
  static is = 'bar-chart';
  static template = document.createElement('template');
  static { this.template.innerHTML = `<slot></slot>`; }
  static { customElements.define(this.is, this); }
  constructor() {
    super()
    this.attachShadow({ mode: 'open' }).append(BarChart.template.content.cloneNode(true));
  }
}
