import { setBasePath } from 'https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/utilities/base-path.js';
setBasePath('https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/');

// Theme switching is now handled by the color-mode-toggle web component

// JSON viewer lazy loading setup
document.addEventListener('DOMContentLoaded', () => {
  const disclosures = document.querySelectorAll('.json-disclosure');
  
  disclosures.forEach(disclosure => {
    const details = disclosure;
    const spinner = details.querySelector('sl-spinner');
    const viewer = details.querySelector('json-viewer');
    const jsonUrl = details.dataset.jsonUrl;
    let loaded = false;
    
    details.addEventListener('sl-show', async () => {
      if (loaded) return;
      
      try {
        const response = await fetch(jsonUrl);
        const jsonData = await response.json();
        
        viewer.data = jsonData;
        viewer.style.display = 'block';
        spinner.style.display = 'none';
        loaded = true;
      } catch (error) {
        console.error('Failed to load JSON:', error);
        spinner.innerHTML = 'Failed to load JSON data';
      }
    });
  });
});

export class BarChart extends HTMLElement {
  static is = 'bar-chart';
  static template = document.createElement('template');
  static { this.template.innerHTML = `<slot></slot>`; }
  static { customElements.define(this.is, this); }
  constructor() {
    super()
    this.attachShadow({ mode: 'open' }).append(BarChart.template.content.cloneNode(true));
  }
}

export class LineChart extends HTMLElement {
  static is = 'line-chart';
  static template = document.createElement('template');
  static {
    this.template.innerHTML = `
      <style>
        :host {
          /* Chart variables */
          --chart-axis-stroke: var(--sl-panel-border-color, #888);
          --chart-label-color: var(--sl-color-neutral-900, currentColor);
          --chart-line-stroke: var(--sl-color-primary-600, #4e79a7);
          --chart-point-fill: var(--sl-tooltip-background-color, #b3c9e5);
          --chart-point-stroke: var(--sl-tooltip-border-color, #4e79a7);
          --chart-point-highlight: var(--sl-color-primary-400, #7ea7d8);
          --chart-point-highlight-stroke: var(--sl-color-primary-700, #22527a);
          /* Tooltip variables */
          --tooltip-background: var(--sl-tooltip-background-color, #222);
          --tooltip-color: var(--sl-tooltip-color, #fff);
          --tooltip-border-radius: var(--sl-tooltip-border-radius, 4px);
          --tooltip-border-color: var(--sl-tooltip-border-color, #4e79a7);
          --tooltip-shadow: var(--sl-tooltip-shadow, 0 2px 8px 0 #2222);
          --tooltip-padding: var(--sl-spacing-x-small, 4px);
        }
        .chart-tooltip {
          position: fixed;
          background: var(--tooltip-background);
          color: var(--tooltip-color);
          padding: var(--tooltip-padding);
          border-radius: var(--tooltip-border-radius);
          border: 1.5px solid var(--tooltip-border-color);
          box-shadow: var(--tooltip-shadow);
          pointer-events: none;
          font-size: 13px;
          z-index: 10000;
          display: none;
          transition: opacity 0.1s;
        }
        .run-point {
          transition: r 0.1s, fill 0.1s, stroke 0.1s;
          cursor: pointer;
        }
        .run-point[data-hovered] {
          fill: var(--chart-point-highlight, #7ea7d8);
          stroke: var(--chart-point-highlight-stroke, #22527a);
          r: 6;
        }
      </style>
    `;
    customElements.define(this.is, this);
  }

  #tooltip;
  #points;
  #svg;
  #hoveredIndex = null;

  constructor() {
    super();
    this.attachShadow({ mode: 'open' }).append(LineChart.template.content.cloneNode(true));
    this.#tooltip = null;
    this.#points = [];
    this.#svg = null;
  }

  connectedCallback() {
    // copy svg into shadowroot
    this.#svg = this.querySelector('svg');
    this.shadowRoot.append(this.#svg);
    if (!this.#svg) return;
    this.#points = Array.from(this.#svg.querySelectorAll('circle.run-point')).map((pt, i) => ({
      el: pt,
      x: parseFloat(pt.getAttribute('cx')),
      y: parseFloat(pt.getAttribute('cy')),
      run: +pt.dataset.run,
      time: pt.dataset.time,
      i
    }));
    this.#svg.addEventListener('pointermove', this.#onPointerMove.bind(this));
    this.#svg.addEventListener('pointerleave', this.#onPointerOut.bind(this));
  }

  #onPointerMove(e) {
    if (!this.#svg || !this.#points.length) return;

    // Convert mouse to SVG coordinates
    const pt = this.#svg.createSVGPoint();
    pt.x = e.clientX;
    pt.y = e.clientY;
    const svgPt = pt.matrixTransform(this.#svg.getScreenCTM().inverse());
    const mouseX = svgPt.x;
    const mouseY = svgPt.y;

    // Find closest point (2D)
    let closest = this.#points[0];
    let minDist = Math.hypot(mouseX - closest.x, mouseY - closest.y);
    for (const pt of this.#points) {
      const dist = Math.hypot(mouseX - pt.x, mouseY - pt.y);
      if (dist < minDist) {
        closest = pt;
        minDist = dist;
      }
    }

    // Highlight and show tooltip as before...
    if (this.#hoveredIndex !== null && this.#points[this.#hoveredIndex]) {
      this.#points[this.#hoveredIndex].el.removeAttribute('data-hovered');
    }
    this.#hoveredIndex = closest.i;
    closest.el.setAttribute('data-hovered', '');

    if (!this.#tooltip) {
      this.#tooltip = document.createElement('div');
      this.#tooltip.className = 'chart-tooltip';
      this.shadowRoot.appendChild(this.#tooltip);
    }
    this.#tooltip.textContent = `Run #${closest.run}: ${closest.time}s`;
    // Position tooltip (you may want to clamp to stay inside viewport)
    const svgRect = this.#svg.getBoundingClientRect();
    this.#tooltip.style.left = `${svgRect.left + closest.x + 12}px`;
    this.#tooltip.style.top = `${svgRect.top + closest.y - 18}px`;
    this.#tooltip.style.display = 'block';
    this.#tooltip.style.opacity = '1';
  }
    #onPointerOut() {
    if (this.#tooltip) {
      this.#tooltip.style.display = 'none';
      this.#tooltip.style.opacity = '0';
    }
    if (this.#hoveredIndex !== null && this.#points[this.#hoveredIndex]) {
      this.#points[this.#hoveredIndex].el.removeAttribute('data-hovered');
    }
    this.#hoveredIndex = null;
  }
}