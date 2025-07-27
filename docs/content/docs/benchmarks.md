---
title: Analyzer Benchmarks
description: Performance benchmarks of analyzers, generated from scripts/benchmark.sh.
weight: 30
---

**Number of runs per tool**: {{<benchmark-runs>}}  
**Number of files analyzed per run**: {{<benchmark-files>}}

<figure class="benchmarks-bar-chart" role="group" aria-labelledby="bar-caption">
  <figcaption id="bar-caption" style="font-weight:bold;margin-bottom:0.5em;">
    Benchmark results: lower is better (seconds)
  </figcaption>
  {{<benchmark-bar-chart>}}
</figure>

---

<style>
line-chart, bar-chart {
  display: block;
}

line-chart svg {
  background:var(--sl-panel-background-color);
  overflow:visible;
}

.analyzer-tool-stats {
  display:flex;
  gap:2em;
  margin:1em 0;
}

.run-breakdown-figure {
  margin: 2em 0;
  figcaption {
    font-weight: bold;
    margin-bottom: 0.5em;
  }
}

bar-chart {
  --bar-fill: light-dark(#e5e7eb, #23272f);
  --bar-text: light-dark(#222, #eee);
  --bar-label: light-dark(#555, #bbb);

  .bar {
    &.perf-best  { --bar-fill: light-dark(#22c55e, #15803d); }
    &.perf-good  { --bar-fill: light-dark(#a3e635, #65a30d); }
    &.perf-mid   { --bar-fill: light-dark(#facc15, #ca8a04); }
    &.perf-worst { --bar-fill: light-dark(#ef4444, #b91c1c); }
  }

  .bar-label {
    fill: var(--bar-text);
    font-weight: 500;
  }
  .bar-value {
    fill: var(--bar-label);
    font-variant-numeric: tabular-nums;
  }
  text {
    pointer-events: none;
  }
}

</style>

<div class="tool-cards">
{{% benchmark-tool-cards.inline %}}
{{ $results := site.Data.benchmarks.results }}
{{ range $idx, $result := $results }}
<section class="analyzer-tool-card">

### {{ $result.name }}

<sl-tag size="small" pill>
  <sl-icon slot="prefix" name="code"></sl-icon>
  <a href="{{ $result.docsUrl }}" target="_blank">
    Docs
    <sl-icon slot="suffix" name="box-arrow-up-right"></sl-icon>
  </a>
</sl-tag>

<div class="analyzer-tool-stats" style="">
  <div>
    <sl-badge variant="success" pill>
      <sl-icon slot="prefix" name="clock"></sl-icon>
      {{ $result.averageTime }}s
    </sl-badge>
    <div class="analyzer-tool-label">Avg Time</div>
  </div>
  <div>
    <sl-badge variant="neutral" pill>
      <sl-icon slot="prefix" name="file-text"></sl-icon>
      {{ $result.averageSize }}KB
    </sl-badge>
    <div class="analyzer-tool-label">Avg Output Size</div>
  </div>
  <div>
    <sl-badge variant="warning" pill>
      <sl-icon slot="prefix" name="hash"></sl-icon>
      {{ len $result.runs }}
    </sl-badge>
    <div class="analyzer-tool-label">Runs</div>
  </div>
  {{- if $result.validation -}}
  <div>
    {{- $errorCount := len $result.validation.errors -}}
    {{- $warningCount := len $result.validation.warnings -}}
    {{- $variant := "success" -}}
    {{- $icon := "check-circle" -}}
    {{- $text := "Valid" -}}
    {{- if gt $errorCount 0 -}}
      {{- $variant = "danger" -}}
      {{- $icon = "x-circle" -}}
      {{- $text = "Invalid" -}}
    {{- else if gt $warningCount 0 -}}
      {{- $variant = "warning" -}}
      {{- $icon = "exclamation-triangle" -}}
      {{- $text = "Warnings" -}}
    {{- end -}}
    <sl-badge variant="{{ $variant }}" pill>
      <sl-icon slot="prefix" name="{{ $icon }}"></sl-icon>
      {{ $text }}
    </sl-badge>
    <div class="analyzer-tool-label">Validation</div>
  </div>
  {{- end -}}
</div>


#### Command
```bash
{{ $result.command }}
```

<figure class="run-breakdown-figure">
  <figcaption>Run Breakdown: Time per Run (seconds)</figcaption>
  <line-chart>
    {{- $n := len $result.runs -}}
    {{- $times := slice -}}
    {{- range $result.runs -}}
      {{- $times = $times | append .time -}}
    {{- end -}}
    {{- $chart_top := 20 -}}
    {{- $chart_bottom := 120 -}}
    {{- $chart_height := sub $chart_bottom $chart_top -}}
    {{- $n_ticks := 6 -}}
    {{- $sorted := sort $times -}}
    {{- $max_time := index $sorted (sub (len $sorted) 1) | default 1 -}}
    {{- $min_time := index $sorted 0 | default 0 -}}
    {{- $y_tick_step := cond (gt (sub $n_ticks 1) 0) (div (sub $max_time $min_time) (sub $n_ticks 1)) 1.0 -}}
    <svg viewBox="0 0 540 160" data-points='[{{ delimit $times ", " }}]'>
      <!-- Axes -->
      <line x1="40" y1="120" x2="520" y2="120" stroke="var(--sl-panel-border-color, #888)" stroke-width="1"/>
      <line x1="40" y1="20" x2="40" y2="120" stroke="var(--sl-panel-border-color, #888)" stroke-width="1"/>
      <!-- Y labels and ticks -->
      {{- range $i := seq 0 (sub $n_ticks 1) -}}
        {{- $value := add $min_time (mul $y_tick_step $i) -}}
        {{- $y := sub $chart_bottom (div (mul $chart_height (sub $value $min_time)) (cond (ne $max_time $min_time) (sub $max_time $min_time) 1)) -}}
        <text x="38" y="{{ add $y 3 }}" font-size="10" text-anchor="end" fill="var(--chart-label-color, var(--sl-color-neutral-900, currentColor))">{{ printf "%.2f" $value }}</text>
        <line x1="40" y1="{{ $y }}" x2="44" y2="{{ $y }}" stroke="var(--sl-panel-border-color, #888)" stroke-width="1"/>
      {{- end -}}
      {{- $.Scratch.Set "points" "" -}}
      {{- range $i, $t := $times -}}
        {{- $x := add 40 (cond (eq $n 1) 0 (div (mul 480 $i) (sub $n 1))) -}}
        {{- $y := sub 120 (div (mul 100 (sub $t $min_time)) (cond (ne $max_time $min_time) (sub $max_time $min_time) 1)) -}}
        {{- $.Scratch.Add "points" (printf "%d,%g " $x $y) -}}
      {{- end -}}
      <polyline fill="none" stroke="var(--chart-line-stroke, var(--sl-color-primary-600, #4e79a7))" stroke-width="2" points="{{ $.Scratch.Get "points" }}"/>
      <!-- Dots for each run -->
      {{- range $i, $t := $times -}}
        <circle cx="{{ add 40 (cond (eq $n 1) 0 (div (mul 480 $i) (sub $n 1))) }}" cy="{{ sub 120 (div (mul 100 (sub $t $min_time)) (cond (ne $max_time $min_time) (sub $max_time $min_time) 1)) }}" r="3" class="run-point" data-run="{{ add $i 1 }}" data-time="{{ $t }}" fill="var(--chart-point-fill, var(--sl-tooltip-background-color, #b3c9e5))" stroke="var(--chart-point-stroke, var(--sl-tooltip-border-color, #4e79a7))" stroke-width="2" />
      {{- end -}}
    <!-- X labels (every 10 runs, always last) -->
    {{- if gt $n 1 -}}
      {{- $last := cond (gt (sub $n 1) 0) (sub $n 1) 0 -}}
      {{- range $j := seq 0 $last 10 -}}
        <text x="{{ add 40 (div (mul 480 $j) (sub $n 1)) }}" y="135" font-size="10" text-anchor="middle" fill="var(--chart-label-color, var(--sl-color-neutral-900, currentColor))">{{ add $j 1 }}</text>
      {{- end -}}
      <text x="520" y="135" font-size="10" text-anchor="middle" fill="var(--chart-label-color, var(--sl-color-neutral-900, currentColor))">{{ $n }}</text>
    {{- else -}}
      <text x="40" y="135" font-size="10" text-anchor="middle" fill="var(--chart-label-color, var(--sl-color-neutral-900, currentColor))">1</text>
    {{- end -}}
      <text x="40" y="150" font-size="12" text-anchor="start" fill="var(--chart-label-color, var(--sl-color-neutral-900, currentColor))">Run #</text>
    </svg>
  </line-chart>
</figure>

{{- if $result.validation -}}

<h4>Validation Results</h4>

{{- if gt (len $result.validation.errors) 0 -}}
<sl-alert variant="danger" open>
  <sl-icon slot="icon" name="exclamation-triangle"></sl-icon>
  <strong>{{ len $result.validation.errors }} Error{{ if ne (len $result.validation.errors) 1 }}s{{ end }}</strong>
  <ul style="margin-top: 0.5em; padding-left: 1.5em;">
  {{- range $result.validation.errors }}
    <li><strong>{{ .id }}:</strong> {{ .message }}{{ if .location }} <em>({{ .location }})</em>{{ end }}</li>
  {{- end }}
  </ul>
</sl-alert>
{{- end -}}

{{- if gt (len $result.validation.warnings) 0 -}}
<sl-alert variant="warning" open>
  <sl-icon slot="icon" name="exclamation-triangle"></sl-icon>
  <strong>{{ len $result.validation.warnings }} Warning{{ if ne (len $result.validation.warnings) 1 }}s{{ end }}</strong>
  <ul style="margin-top: 0.5em; padding-left: 1.5em;">
  {{- range $result.validation.warnings }}
    <li><strong>{{ .id }}:</strong> {{ .message }}{{ if .location }} <em>({{ .location }})</em>{{ end }}</li>
  {{- end }}
  </ul>
</sl-alert>
{{- end -}}

{{- if and (eq (len $result.validation.errors) 0) (eq (len $result.validation.warnings) 0) -}}
<sl-alert variant="success" open>
  <sl-icon slot="icon" name="check-circle"></sl-icon>
  <strong>No validation issues found</strong> - The generated manifest is valid and follows all best practices.
</sl-alert>
{{- end -}}

{{- end -}}

<sl-details summary="Last Output (JSON)">
<sl-spinner style="font-size: 3rem;"></sl-spinner>
<zero-md id="last-output-{{ $idx }}" data-src="{{ $result.lastOutputUrl | relURL }}" no-shadow><template></template></zero-md>
</sl-details>
{{- with $result.lastError -}}
  {{- $trimmed := trim . " \n\r\t" -}}
  {{- if and (ne $trimmed "") (not (eq $trimmed nil)) -}}
    <sl-alert variant="danger" open>
      <sl-icon slot="icon" name="exclamation-triangle"></sl-icon>
      Last error: <pre>{{ . }}</pre>
    </sl-alert>
  {{ end }}
{{- end -}}
</section>
  {{- end -}}
{{% /benchmark-tool-cards.inline %}}

</div>

{{< loadchart.inline >}}
<script type="module" src="{{ "charts.js" | relURL }} "></script>
{{</ loadchart.inline >}}
<script type="module" src="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/cdn/shoelace.js"></script>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/themes/light.css">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.16.0/dist/themes/dark.css">
<link id="hljs-light" rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@highlightjs/cdn-assets@11/styles/github.min.css" disabled>
<link id="hljs-dark"  rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@highlightjs/cdn-assets@11/styles/github-dark.min.css" disabled>

<script type="module">
    import ZeroMd from "https://esm.sh/zero-md@3";
    customElements.define('zero-md', ZeroMd);
</script>

<script type="module">
</script>
