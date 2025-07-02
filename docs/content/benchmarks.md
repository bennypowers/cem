---
title: Analyzer Benchmarks
description: Performance benchmarks of analyzers, generated from scripts/benchmark.sh.
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

{{<benchmark-tool-cards>}}
