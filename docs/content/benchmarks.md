---
title: Analyzer Benchmarks
description: Performance benchmarks of analyzers, generated from scripts/benchmark.sh.
---

# Analyzer Benchmarks

**Number of runs per tool**: {{ .Site.Data.benchmarks.runs }}  
**Number of files analyzed per run**: {{ .Site.Data.benchmarks.file_count }}

<figure class="benchmarks-bar-chart" role="group" aria-labelledby="bar-caption">
  <figcaption id="bar-caption" style="font-weight:bold;margin-bottom:0.5em;">
    Benchmark results: lower is better (seconds)
  </figcaption>
  {{<benchmark-bar-chart>}}
</figure>

---

{{<benchmark-tool-cards>}}
