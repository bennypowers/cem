#!/usr/bin/env bash

set -e

runs=${1:-1}
if ! [[ "$runs" =~ ^[0-9]+$ ]]; then
  echo "Usage: $0 [number_of_runs]"
  exit 1
fi

cd analyzer-benchmarks
if [[ -n "$CI" ]]; then
  pwd
  npm ci
fi

file_glob="src/components/*.ts"
file_count=$(ls $file_glob 2>/dev/null | wc -l | xargs)

echo "$file_glob"

ids=(
  "lit"
  "cea"
  "cem"
)
names=(
  "@lit-labs/cli"
  "@custom-elements-manifest/analyzer"
  "bennypowers/cem generate"
)
cmds=(
  "npx --yes @lit-labs/cli labs gen --manifest --out lit"
  "npx --yes @custom-elements-manifest/analyzer analyze --outdir cea --globs $file_glob"
  "../cem generate -o cem/custom-elements.json $file_glob"
)
docsUrls=(
  "https://github.com/lit/lit/tree/fbda6d7b42b8acd19b388e9de0be3521da6b58bb/packages/labs/cli"
  "https://github.com/open-wc/custom-elements-manifest/tree/master/packages/analyzer"
  "https://github.com/bennypowers/cem"
)

# Pre-warm commands to avoid first-run hangs on macOS
for i in "${!cmds[@]}"; do
  id="${ids[$i]}"
  mkdir -p "$id"
  resultFile="${id}/custom-elements.json"
  if [[ -f "$resultFile" ]]; then
    rm -f "$resultFile"
  fi
  echo "Pre-warming: ${cmds[$i]}"
  eval "${cmds[$i]}"
done

# Collect results
results_array=()
for i in "${!ids[@]}"; do
  id="${ids[$i]}"
  name="${names[$i]}"
  id="${ids[$i]}"
  cmd="${cmds[$i]}"
  resultFile="${id}/custom-elements.json"
  docsUrl="${docsUrls[$i]}"

  sumTime=0
  sumSize=0
  successful_runs=0
  runs_detail=()
  last_json=""
  last_stderr=""
  for ((r=1; r<=runs; r++)); do
    start=$(date +%s.%N)
    tmp_stdout=$(mktemp)
    tmp_stderr=$(mktemp)
    if eval "$cmd" >"$tmp_stdout" 2>"$tmp_stderr"; then
      run_ok=true
    else
      run_ok=false
      last_stderr=$(cat "$tmp_stderr")
    fi
    end=$(date +%s.%N)
    timeSec=$(awk "BEGIN {printf \"%.4f\", $end - $start}")

    if [[ -f "$resultFile" ]]; then
      if stat --version >/dev/null 2>&1; then
        file_size=$(stat -c%s "$resultFile")
      else
        file_size=$(stat -f%z "$resultFile")
      fi
      size=$(awk "BEGIN {printf \"%.1f\", $file_size / 1024 }")
      sumTime=$(awk "BEGIN {print $sumTime + $timeSec}")
      sumSize=$(awk "BEGIN {print $sumSize + $size}")
      successful_runs=$((successful_runs+1))
      runs_detail+=("{\"run\":$r,\"time\":$timeSec,\"size\":$size}")
      last_json=$(cat "$resultFile")
    else
      runs_detail+=("{\"run\":$r,\"time\":null,\"size\":null,\"error\":\"$last_stderr\"}")
    fi
    rm "$tmp_stdout" "$tmp_stderr"
  done

  if [[ $successful_runs -gt 0 ]]; then
    avgTime=$(awk "BEGIN {printf \"%.2f\", $sumTime/$successful_runs}")
    avgSize=$(awk "BEGIN {printf \"%.1f\", $sumSize/$successful_runs}")
  else
    avgTime="0"
    avgSize="0"
  fi

  results_array+=("{
    \"id\": \"$id\",
    \"name\": \"$name\",
    \"docsUrl\": \"$docsUrl\",
    \"command\": \"$cmd\",
    \"averageTime\": $avgTime,
    \"averageSize\": $avgSize,
    \"runs\": [$(IFS=,; echo "${runs_detail[*]}")],
    \"lastOutput\": $(echo "$last_json" | jq . 2>/dev/null || echo "null"),
    \"lastError\": $(jq -Rs . <<<"$last_stderr")
  }")
done

jq -n --argjson results "[$(IFS=,; echo "${results_array[*]}")]" \
  --arg runs "$runs" \
  --arg file_count "$file_count" \
  '{runs: ($runs|tonumber), file_count: ($file_count|tonumber), results: $results}' > benchmark-results.json

if [[ -n "$CI" ]]; then
  npm run generate-site
fi

