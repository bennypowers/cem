import fs from 'node:fs/promises';

import results from './benchmark-results.json' with { type: "json" };

function summaryCard(result) {
  return `
    <sl-card class="analyzer-summary-card" style="margin-bottom: 2rem;">
      <div slot="header" style="display:flex;align-items:center;gap:0.5rem;">
        <sl-icon name="package"></sl-icon>
        <span style="font-weight:600">${result.name}</span>
        <sl-badge variant="primary" pill style="margin-left: auto;">${result.id}</sl-badge>
      </div>
      <div style="margin-top: .5rem">
        <sl-tag size="small" pill>
          <sl-icon slot="prefix" name="code"></sl-icon>
          <a href="${result.docsUrl}" target="_blank" style="text-decoration:none;color:inherit;">
            Docs
            <sl-icon slot="suffix" name="external-link"></sl-icon>
          </a>
        </sl-tag>
      </div>
      <div style="display:flex;gap:1.5rem;align-items:center;margin-top:1.2rem;">
        <div>
          <sl-badge variant="success" pill>
            <sl-icon slot="prefix" name="clock"></sl-icon>
            ${result.averageTime}s
          </sl-badge>
          <div style="font-size:.9em;opacity:0.7;">Avg Time</div>
        </div>
        <div>
          <sl-badge variant="neutral" pill>
            <sl-icon slot="prefix" name="file-text"></sl-icon>
            ${result.averageSize}KB
          </sl-badge>
          <div style="font-size:.9em;opacity:0.7;">Avg Output Size</div>
        </div>
        <div>
          <sl-badge variant="warning" pill>
            <sl-icon slot="prefix" name="hash"></sl-icon>
            ${result.runs.length}
          </sl-badge>
          <div style="font-size:.9em;opacity:0.7;">Runs</div>
        </div>
      </div>
      <sl-details summary="Show Command" style="margin-top: 1em;">
        <sl-input readonly value="${result.command}" style="width:100%"></sl-input>
      </sl-details>
    </sl-card>
  `;
}

function detailsCard(result) {
  return `
    <sl-card class="analyzer-detail-card" style="margin-bottom: 2rem;">
      <div slot="header" style="display:flex;align-items:center;gap:0.5rem;">
        <sl-icon name="info"></sl-icon>
        <span style="font-weight:600">${result.name} Details</span>
      </div>
      <sl-details summary="Run Breakdown" open>
        <ul style="padding-left:1.2em;">
          ${result.runs.map(run =>
            run.time !== null
              ? `<li>
                  <sl-badge variant="success" pill>
                    Run ${run.run}
                  </sl-badge>
                  <sl-badge variant="primary" pill>
                    <sl-icon slot="prefix" name="clock"></sl-icon>
                    ${run.time}s
                  </sl-badge>
                  <sl-badge variant="neutral" pill>
                    <sl-icon slot="prefix" name="file-text"></sl-icon>
                    ${run.size}KB
                  </sl-badge>
                </li>`
              : `<li>
                  <sl-badge variant="danger" pill>
                    Run ${run.run} FAILED
                  </sl-badge>
                  ${run.error ? `<sl-details summary="Show Error"><pre>${run.error}</pre></sl-details>` : ''}
                </li>`
          ).join('')}
        </ul>
      </sl-details>
      <sl-details summary="Last Output (JSON)">
        <sl-copy-button value="${typeof result.lastOutput === 'object' ? JSON.stringify(result.lastOutput, null, 2) : result.lastOutput || ''}" style="float:right"></sl-copy-button>
        <pre style="max-height:350px;overflow:auto;">${result.lastOutput && typeof result.lastOutput === 'object'
            ? JSON.stringify(result.lastOutput, null, 2)
            : result.lastOutput || "No output"
        }</pre>
      </sl-details>
      ${result.lastError && result.lastError !== '""'
        ? `<sl-alert variant="danger" open style="margin-top:1em;">
              <sl-icon slot="icon" name="alert-triangle"></sl-icon>
              Last error: <pre>${result.lastError}</pre>
           </sl-alert>`
        : ''}
    </sl-card>
  `;
}

await fs.mkdir('site', { recursive: true });
await fs.writeFile('site/index.html', /*html*/`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Analyzer Benchmarks</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <script type="module" src="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/cdn/shoelace.js"></script>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/themes/light.css">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/themes/dark.css">
  <style>
    body {
      margin: 0; min-height: 100vh;
      background: var(--sl-color-neutral-0);
      color: var(--sl-color-neutral-900);
    }
    header {
      display: flex; align-items: center; justify-content: space-between;
      padding: 1.5rem 2rem 1rem 2rem;
      background: var(--sl-color-primary-600);
      color: white;
    }
    header h1 { margin: 0; font-size: 1.6rem; letter-spacing: 2px; }
    main {
      max-width: 900px; margin: 2rem auto; padding: 2rem;
    }
    sl-card {
      display:block;margin-bottom:2rem;
    }
    .summary-cards, .detail-cards {
      display:flex;flex-direction:column;gap:2rem;
    }
    .breadcrumb {
      margin-bottom: 2rem;
    }
    @media (max-width: 700px) {
      main { padding: 0.5rem; }
      sl-card { padding: 0.5rem; }
    }
  </style>
</head>
<body>
  <header>
    <div style="display:flex;align-items:center;gap:1rem;">
      <sl-icon name="bar-chart-2"></sl-icon>
      <h1>Web Components Analyzer Benchmarks</h1>
    </div>
    <sl-switch id="theme-toggle" checked>Dark Mode</sl-switch>
  </header>
  <main>
    <div class="breadcrumb">
      <sl-breadcrumb>
        <sl-breadcrumb-item href="/">Home</sl-breadcrumb-item>
        <sl-breadcrumb-item href="#" aria-current="page">Benchmarks</sl-breadcrumb-item>
      </sl-breadcrumb>
    </div>
    <sl-alert variant="primary" open>
      <sl-icon slot="icon" name="bar-chart-2"></sl-icon>
      <strong>Analyzer Benchmarks</strong>
      <div slot="description">
        Automated, reproducible performance and output-size comparisons of leading web components analyzers.<br>
        Data is updated on every push. Try toggling the theme or exploring the metrics below!
      </div>
    </sl-alert>
    <sl-card>
      <div style="font-size: 1.1em;">
        <strong>Number of runs per tool:</strong> ${results.runs}<br>
        <strong>Number of files analyzed per run:</strong> ${results.file_count}
      </div>
    </sl-card>
    <sl-tab-group placement="top" style="margin-bottom:2rem;">
      <sl-tab slot="nav" panel="panel-summary">
        <sl-icon name="list"></sl-icon>Summary
      </sl-tab>
      <sl-tab slot="nav" panel="panel-details">
        <sl-icon name="file-bar-chart-2"></sl-icon>Details
      </sl-tab>
      <sl-tab-panel name="panel-summary">
        <div class="summary-cards">
          ${results.results.map(summaryCard).join('')}
        </div>
      </sl-tab-panel>
      <sl-tab-panel name="panel-details">
        <div class="detail-cards">
          ${results.results.map(detailsCard).join('')}
        </div>
      </sl-tab-panel>
    </sl-tab-group>
  </main>
  <script type="module">
    import { setBasePath } from 'https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/utilities/base-path.js';
    setBasePath('https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/');
    const themeToggle = document.getElementById('theme-toggle');
    const updateTheme = () => {
      const dark = themeToggle.checked;
      document.body.setAttribute('class', dark ? 'sl-theme-dark' : 'sl-theme-light');
    };
    themeToggle.addEventListener('sl-change', updateTheme);
    updateTheme();
  </script>
</body>
</html>
`);
