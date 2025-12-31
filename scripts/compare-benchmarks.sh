#!/usr/bin/env bash

set -euo pipefail

# Input locations
PR_TIME_FILE="pr_bench_time.txt"
BASE_TIME_FILE="base_bench_time.txt"
PR_JSON="pr_bench.json"
BASE_JSON="base_bench.json"
PR_BENCH_TXT="pr_bench.txt"
BASE_BENCH_TXT="base_bench.txt"

# Check if all required files exist (first run case)
missing_files=()
for file in "$PR_TIME_FILE" "$BASE_TIME_FILE" "$PR_JSON" "$BASE_JSON" "$PR_BENCH_TXT" "$BASE_BENCH_TXT"; do
  if [ ! -f "$file" ]; then
    missing_files+=("$file")
  fi
done

if [ ${#missing_files[@]} -gt 0 ]; then
  cat << EOF > bench_report.md
### Benchmark Summary

⚠️ **First Run**: Benchmark comparison requires both PR and base branch results.

Missing files:
$(printf '- `%s`\n' "${missing_files[@]}")

To generate benchmarks, run:
\`\`\`bash
make bench
\`\`\`

The next benchmark run will include a comparison.
EOF
  echo "First run: missing benchmark files. Created informational report."
  exit 0
fi

# Check if base is a placeholder (first run scenario)
base_missing=false
if [ "$(cat "$BASE_TIME_FILE")" = "0" ] && jq -e '. == {}' "$BASE_JSON" >/dev/null 2>&1; then
  base_missing=true
fi

if [ "$base_missing" = true ]; then
  # Base is missing - show placeholders
  pr_size_kb=$(du -k "$PR_JSON" | cut -f1)
  base_size_kb="---"

  pr_time=$(cat "$PR_TIME_FILE")
  base_time="---"

  pr_num_runs=$(grep -c '^Benchmark' "$PR_BENCH_TXT" || echo 1)
  base_num_runs="---"

  pr_avg_time=$(awk "BEGIN {print ($pr_num_runs == 0 ? 0 : $pr_time / $pr_num_runs)}")
  base_avg_time="---"

  pr_perf_kb=$(awk "BEGIN {print ($pr_size_kb == 0 ? 0 : $pr_time / $pr_size_kb)}")
  base_perf_kb="---"

  delta_perf="---"
  delta_ratio="---"
  emoji="ℹ️"
else
  # Normal comparison
  pr_size_kb=$(du -k "$PR_JSON" | cut -f1)
  base_size_kb=$(du -k "$BASE_JSON" | cut -f1)

  pr_time=$(cat "$PR_TIME_FILE")
  base_time=$(cat "$BASE_TIME_FILE")

  pr_num_runs=$(grep -c '^Benchmark' "$PR_BENCH_TXT" || echo 1)
  base_num_runs=$(grep -c '^Benchmark' "$BASE_BENCH_TXT" || echo 1)

  pr_avg_time=$(awk "BEGIN {print ($pr_num_runs == 0 ? 0 : $pr_time / $pr_num_runs)}")
  base_avg_time=$(awk "BEGIN {print ($base_num_runs == 0 ? 0 : $base_time / $base_num_runs)}")

  pr_perf_kb=$(awk "BEGIN {print ($pr_size_kb == 0 ? 0 : $pr_time / $pr_size_kb)}")
  base_perf_kb=$(awk "BEGIN {print ($base_size_kb == 0 ? 0 : $base_time / $base_size_kb)}")

  delta_perf=$(awk "BEGIN {print $pr_perf_kb - $base_perf_kb}")
  delta_ratio=$(awk "BEGIN {if ($pr_perf_kb == 0) print 0; else print $base_perf_kb / $pr_perf_kb}")

  # Choose emoji
  if awk "BEGIN {exit !($delta_ratio > 10)}"; then
    emoji='🚀'
  elif awk "BEGIN {exit !($delta_ratio > 1)}"; then
    emoji='👍'
  elif awk "BEGIN {exit !($delta_ratio < 0.1)}"; then
    emoji='💣'
  elif awk "BEGIN {exit !($delta_ratio < 1)}"; then
    emoji='🐢'
  else
    emoji='➖'
  fi
fi

# Get branch names (or use placeholders for local testing)
base_ref="${GITHUB_BASE_REF:-main}"
head_ref="${GITHUB_HEAD_REF:-current}"

# Get workflow run URL (set by GitHub automatically)
run_url="${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY}/actions/runs/${GITHUB_RUN_ID}"

# Output markdown
if [ "$base_missing" = true ]; then
  cat << EOF > bench_report.md
### Benchmark Summary

|        | Branch    | Total Time (s) | # Runs | Avg Time/run (s) | Output Size (kb) | Perf/kb (s/kb) |
|--------|-----------|----------------|--------|------------------|------------------|----------------|
| Base   | \`${base_ref}\` | $base_time | $base_num_runs | $base_avg_time | $base_size_kb | $base_perf_kb |
| PR     | \`${head_ref}\` | $pr_time | $pr_num_runs | $pr_avg_time | $pr_size_kb | $pr_perf_kb |
| Δ      |           | --- | --- | --- | --- | --- $emoji |

ℹ️ **First Run**: Base branch does not have benchmark infrastructure. Showing PR results only. Future runs will include comparisons.

[View this benchmark run in GitHub Actions]($run_url)

_💡 Tip: Raw JSON outputs are available in [workflow artifacts]($run_url) if needed._
EOF
else
  cat << EOF > bench_report.md
### Benchmark Summary

|        | Branch    | Total Time (s) | # Runs | Avg Time/run (s) | Output Size (kb) | Perf/kb (s/kb) |
|--------|-----------|----------------|--------|------------------|------------------|----------------|
| Base   | \`${base_ref}\` | $base_time | $base_num_runs | $base_avg_time | $base_size_kb | $base_perf_kb |
| PR     | \`${head_ref}\` | $pr_time | $pr_num_runs | $pr_avg_time | $pr_size_kb | $pr_perf_kb |
| Δ      |           | $(awk "BEGIN {printf \"%+.4f\", $pr_time - $base_time}") | $(($pr_num_runs - $base_num_runs)) | $(awk "BEGIN {printf \"%+.4f\", $pr_avg_time - $base_avg_time}") | $(($pr_size_kb - $base_size_kb)) | $(awk "BEGIN {printf \"%+.4f\", $delta_perf}") $emoji |

**Perf/kb delta ratio:** $(awk "BEGIN {printf \"%.2fx\", $delta_ratio}") $emoji

[View this benchmark run in GitHub Actions]($run_url)

_💡 Tip: Raw JSON outputs are available in [workflow artifacts]($run_url) if needed._
EOF
fi

echo "✅ Benchmark comparison complete"
