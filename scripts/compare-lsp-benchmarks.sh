#!/usr/bin/env bash

set -euo pipefail

# Input files
PR_JSON="pr_lsp_bench.json"
BASE_JSON="base_lsp_bench.json"

# Check if required files exist (first run case)
if [ ! -f "$PR_JSON" ] || [ ! -f "$BASE_JSON" ]; then
  cat << EOF > lsp_bench_report.md
### LSP Benchmark Results

⚠️ **First Run**: LSP benchmark comparison requires both PR and base branch results.

Missing files:
$([ ! -f "$PR_JSON" ] && echo "- \`$PR_JSON\`")
$([ ! -f "$BASE_JSON" ] && echo "- \`$BASE_JSON\`")

To generate LSP benchmarks, run:
\`\`\`bash
make bench-lsp-cem
\`\`\`

The next benchmark run will include a comparison.
EOF
  echo "First run: missing LSP benchmark files. Created informational report."
  exit 0
fi

# Helper: Extract metric with jq, fallback to 0 if missing
get_metric() {
  local file=$1
  local path=$2
  jq -r "$path // 0" "$file" 2>/dev/null || echo "0"
}

# Helper: Get success rate (0-100%)
get_success_rate() {
  local file=$1
  local path=$2
  local rate
  rate=$(jq -r "$path // 0" "$file" 2>/dev/null || echo "0")
  awk "BEGIN {printf \"%.0f%%\", $rate * 100}"
}

# Check if benchmarks failed completely
if jq -e '.error' "$PR_JSON" >/dev/null 2>&1 && jq -e '.error' "$BASE_JSON" >/dev/null 2>&1; then
  cat << EOF > lsp_bench_report.md
### LSP Benchmark Results

⚠️ LSP benchmarks failed on both PR and base branches. See workflow logs for details.
EOF
  exit 0
fi

# If only one failed, note it but continue with partial results
pr_failed=false
base_failed=false
jq -e '.error' "$PR_JSON" >/dev/null 2>&1 && pr_failed=true
jq -e '.error' "$BASE_JSON" >/dev/null 2>&1 && base_failed=true

# Extract metrics for each benchmark
# 1. Startup
pr_startup=$(get_metric "$PR_JSON" '.benchmarks.startup.statistics.mean')
base_startup=$(get_metric "$BASE_JSON" '.benchmarks.startup.statistics.mean')
startup_success=$(get_success_rate "$PR_JSON" '.benchmarks.startup.success_rate')

# 2. Hover
pr_hover=$(get_metric "$PR_JSON" '.benchmarks.hover.overall_statistics.mean')
base_hover=$(get_metric "$BASE_JSON" '.benchmarks.hover.overall_statistics.mean')
hover_success=$(get_success_rate "$PR_JSON" '.benchmarks.hover.success_rate')

# 3. Completion
pr_completion=$(get_metric "$PR_JSON" '.benchmarks.completion.overall_statistics.mean')
base_completion=$(get_metric "$BASE_JSON" '.benchmarks.completion.overall_statistics.mean')
completion_success=$(get_success_rate "$PR_JSON" '.benchmarks.completion.success_rate')

# 4. Diagnostics
pr_diagnostics=$(get_metric "$PR_JSON" '.benchmarks.diagnostics.duration_ms')
base_diagnostics=$(get_metric "$BASE_JSON" '.benchmarks.diagnostics.duration_ms')
diagnostics_success=$(get_success_rate "$PR_JSON" '.benchmarks.diagnostics.success_rate')

# 5. Attribute Hover
pr_attr_hover=$(get_metric "$PR_JSON" '.benchmarks.hover_attribute.overall_statistics.mean')
base_attr_hover=$(get_metric "$BASE_JSON" '.benchmarks.hover_attribute.overall_statistics.mean')
attr_hover_success=$(get_success_rate "$PR_JSON" '.benchmarks.hover_attribute.success_rate')

# 6. References
pr_references=$(get_metric "$PR_JSON" '.benchmarks.references.average_search_time')
base_references=$(get_metric "$BASE_JSON" '.benchmarks.references.average_search_time')
references_success=$(get_success_rate "$PR_JSON" '.benchmarks.references.success_rate')

# Helper: Calculate delta and choose emoji
calc_delta() {
  local pr=$1
  local base=$2
  local delta
  delta=$(awk "BEGIN {printf \"%.2f\", $pr - $base}")
  local pct=0

  if awk "BEGIN {exit !($base != 0)}"; then
    pct=$(awk "BEGIN {printf \"%.1f\", ($pr - $base) / $base * 100}")
  fi

  # Choose emoji based on percentage change
  local emoji=""
  if awk "BEGIN {exit !($pct < 0)}"; then
    emoji='✅'
  elif awk "BEGIN {exit !($pct >= 0 && $pct < 5)}"; then
    emoji='➖'
  elif awk "BEGIN {exit !($pct >= 5 && $pct < 10)}"; then
    emoji='⚠️'
  elif awk "BEGIN {exit !($pct >= 10 && $pct < 25)}"; then
    emoji='🐢'
  else
    emoji='❌'
  fi

  echo "${delta} (${pct}%) ${emoji}"
}

# Calculate deltas for all benchmarks
startup_delta=$(calc_delta "$pr_startup" "$base_startup")
hover_delta=$(calc_delta "$pr_hover" "$base_hover")
completion_delta=$(calc_delta "$pr_completion" "$base_completion")
diagnostics_delta=$(calc_delta "$pr_diagnostics" "$base_diagnostics")
attr_hover_delta=$(calc_delta "$pr_attr_hover" "$base_attr_hover")
references_delta=$(calc_delta "$pr_references" "$base_references")

# Get workflow run URL (or use placeholder for local testing)
if [ -n "${GITHUB_SERVER_URL:-}" ] && [ -n "${GITHUB_REPOSITORY:-}" ] && [ -n "${GITHUB_RUN_ID:-}" ]; then
  run_url="${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY}/actions/runs/${GITHUB_RUN_ID}"
else
  run_url="[Local test run - no GitHub Actions URL available]"
fi

# Generate markdown report
cat << EOF > lsp_bench_report.md
### LSP Benchmark Results

| Benchmark       | PR Mean (ms) | Base Mean (ms) | Delta          | Success Rate | Status |
|-----------------|--------------|----------------|----------------|--------------|--------|
| Startup         | $pr_startup  | $base_startup  | $startup_delta | $startup_success | |
| Hover           | $pr_hover    | $base_hover    | $hover_delta   | $hover_success   | |
| Completion      | $pr_completion | $base_completion | $completion_delta | $completion_success | |
| Diagnostics     | $pr_diagnostics | $base_diagnostics | $diagnostics_delta | $diagnostics_success | |
| Attribute Hover | $pr_attr_hover | $base_attr_hover | $attr_hover_delta | $attr_hover_success | |
| References      | $pr_references | $base_references | $references_delta | $references_success | |

[View this benchmark run in GitHub Actions]($run_url)

<details>
<summary>Raw PR LSP Results</summary>

\`\`\`json
$(cat "$PR_JSON")
\`\`\`
</details>

<details>
<summary>Raw Base LSP Results</summary>

\`\`\`json
$(cat "$BASE_JSON")
\`\`\`
</details>
EOF

if [ "$pr_failed" = true ]; then
  echo "" >> lsp_bench_report.md
  echo "⚠️ **Note**: PR branch LSP benchmarks failed. Showing base branch results only." >> lsp_bench_report.md
fi

if [ "$base_failed" = true ]; then
  echo "" >> lsp_bench_report.md
  echo "⚠️ **Note**: Base branch LSP benchmarks failed. Showing PR branch results only." >> lsp_bench_report.md
fi

echo "✅ LSP benchmark comparison complete"
