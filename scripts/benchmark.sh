#!/usr/bin/env bash

set -e

runs=${1:-1}
if ! [[ "$runs" =~ ^[0-9]+$ ]]; then
  echo "Usage: $0 [number_of_runs]"
  exit 1
fi

mkdir -p docs/assets
cd docs
if [[ -n "$CI" ]]; then
  pwd
  npm ci
fi

file_glob="benchmark/components/*.ts"
file_count=$(ls $file_glob 2>/dev/null | wc -l | xargs)

export PATH="../dist:$PATH"

echo "$file_glob"

ids=(
  "lit"
  "cea"
  "cem"
)
names=(
  "@lit-labs/cli"
  "@custom-elements-manifest/analyzer"
  "cem generate"
)
cmds=(
  "npx --yes @lit-labs/cli labs gen --manifest --out data/lit"
  "npx --yes @custom-elements-manifest/analyzer analyze --outdir data/cea --globs $file_glob"
  "cem generate -o data/cem/custom-elements.json $file_glob"
)
docsUrls=(
  "https://github.com/lit/lit/tree/fbda6d7b42b8acd19b388e9de0be3521da6b58bb/packages/labs/cli"
  "https://github.com/open-wc/custom-elements-manifest/tree/master/packages/analyzer"
  "https://github.com/bennypowers/cem"
)

# Pre-warm commands to avoid first-run hangs on macOS
for i in "${!cmds[@]}"; do
  id="${ids[$i]}"
  mkdir -p "data/$id"
  resultFile="data/${id}/custom-elements.json"
  if [[ -f "$resultFile" ]]; then
    rm -f "$resultFile"
  fi
  echo "Pre-warming: ${cmds[$i]}"
  eval "${cmds[$i]}"
done

# Prepare a temp file for results to avoid "argument list too long" with jq
results_tmp=$(mktemp)

# Collect results
for i in "${!ids[@]}"; do
  name="${names[$i]}"
  id="${ids[$i]}"
  cmd="${cmds[$i]}"
  resultFile="data/${id}/custom-elements.json"
  docsUrl="${docsUrls[$i]}"

  sumTime=0
  sumSize=0
  successful_runs=0
  runs_detail=()
  last_json=""
  last_stderr=""
  printf "\nRunning %s: " "$name"
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

    # Progress indicator
    if [[ "$run_ok" == true ]]; then
      printf "."
    else
      printf "x"
    fi

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
    # Optionally show 1/x progress every 10 runs
    if (( r % 10 == 0 )); then
      printf " (%d/%d)" "$r" "$runs"
    fi
  done
  printf " done [%d/%d successful]\n" "$successful_runs" "$runs"

  if [[ $successful_runs -gt 0 ]]; then
    avgTime=$(awk "BEGIN {printf \"%.2f\", $sumTime/$successful_runs}")
    avgSize=$(awk "BEGIN {printf \"%.1f\", $sumSize/$successful_runs}")
  else
    avgTime="0"
    avgSize="0"
  fi
  lastOutput="$(echo "$last_json" | jq .)"

  # Validate the final manifest once (only if we have successful runs)
  validation_result="{}"
  if [[ $successful_runs -gt 0 && -f "$resultFile" ]]; then
    validation_tmp=$(mktemp)
    if cem validate --format=json "$resultFile" >"$validation_tmp" 2>/dev/null; then
      validation_result=$(cat "$validation_tmp")
    else
      validation_result='{"valid":false,"errors":[],"warnings":[],"message":"Validation failed"}'
    fi
    rm -f "$validation_tmp"
  fi

  # Write each tool's result as a single line in the temp file
  # Sanitize last_stderr as empty string if only whitespace or empty
  clean_last_stderr="$(echo -n "$last_stderr" | sed 's/^[[:space:]]*//;s/[[:space:]]*$//')"
  if [[ -z "$clean_last_stderr" ]]; then
    clean_last_stderr=""
  fi

  jq -n \
    --arg id "$id" \
    --arg name "$name" \
    --arg docsUrl "$docsUrl" \
    --arg command "$cmd" \
    --arg lastOutputUrl "$id-last-output.md" \
    --argjson averageTime "$avgTime" \
    --argjson averageSize "$avgSize" \
    --argjson runs "[$(IFS=,; echo "${runs_detail[*]}")]" \
    --argjson lastOutput "$lastOutput" \
    --arg lastError "$clean_last_stderr" \
    --argjson validation "$validation_result" \
    '{
      id: $id,
      name: $name,
      docsUrl: $docsUrl,
      command: $command,
      averageTime: $averageTime,
      averageSize: $averageSize,
      runs: $runs,
      lastOutput: $lastOutput,
      lastOutputUrl: $lastOutputUrl,
      lastError: $lastError,
      validation: $validation
    }' >> "$results_tmp"

  output_file="assets/$id-last-output.md"
  echo "\`\`\`json" > "$output_file"
  echo "$lastOutput" >> "$output_file"
  echo "\`\`\`" >> "$output_file"

done

# Combine all tool results into a single JSON array
jq -s --arg runs "$runs" --arg file_count "$file_count" \
  '{runs: ($runs|tonumber), file_count: ($file_count|tonumber), results: .}' "$results_tmp" > data/benchmarks.json

rm "$results_tmp"

