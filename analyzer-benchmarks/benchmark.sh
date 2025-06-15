#!/usr/bin/env bash

set -e

# Take an integer argument for number of runs, default to 1 if not provided
runs=${1:-1}
if ! [[ "$runs" =~ ^[0-9]+$ ]]; then
  echo "Usage: $0 [number_of_runs]"
  exit 1
fi

# Count the number of files to be analyzed (in src/components/*.ts)
file_glob="src/components/*.ts"
file_count=$(ls $file_glob 2>/dev/null | wc -l | xargs)

# Define analyzers as arrays
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
  "./cem/cem generate -o cem/custom-elements.json $file_glob"
)
resultFiles=(
  "lit/custom-elements.json"
  "cea/custom-elements.json"
  "cem/custom-elements.json"
)
docsUrls=(
  "https://github.com/lit/lit/tree/fbda6d7b42b8acd19b388e9de0be3521da6b58bb/packages/labs/cli"
  "https://github.com/open-wc/custom-elements-manifest/tree/master/packages/analyzer"
  "https://github.com/bennypowers/cem"
)

resultsMd='| Tool | Avg Time (s) | Avg Output Size (KB) |\n|---|---|---|\n'
detailsMd=''

# Pre-warm commands to avoid first-run hangs on macOS
for i in "${!cmds[@]}"; do
  resultFile="${resultFiles[$i]}"
  if [[ -f "$resultFile" ]]; then
    rm "$resultFile"
  fi
  echo "Pre-warming: ${cmds[$i]}"
  eval "${cmds[$i]}"
done

for i in "${!ids[@]}"; do
  name="${names[$i]}"
  cmd="${cmds[$i]}"
  resultFile="${resultFiles[$i]}"
  docsUrl="${docsUrls[$i]}"

  echo "::group::Running $name ($runs runs)"
  echo "About to run: $cmd"

  sumTime=0
  sumSize=0
  successful_runs=0
  stderr_all=""
  results_run=()

  detailsMd+="<details><summary><code>${name}</code> Results</summary>\n\n"
  detailsMd+="\n\n\`${cmd}\`\n\n"

  for ((r=1; r<=runs; r++)); do
    echo "Run $r/$runs..."
    stdout=""
    stderr=""
    timeSec=""
    start=$(date +%s.%N)

    tmp_stdout=$(mktemp)
    tmp_stderr=$(mktemp)
    if eval "$cmd" >"$tmp_stdout" 2>"$tmp_stderr"; then
      run_ok=true
    else
      run_ok=false
      echo "Error running $name on run $r"
    fi
    end=$(date +%s.%N)

    stdout=$(cat "$tmp_stdout")
    stderr=$(cat "$tmp_stderr")
    rm "$tmp_stdout" "$tmp_stderr"

    # Calculate time in seconds, with 2 decimals
    timeSec=$(awk "BEGIN {printf \"%.4f\", $end - $start}")

    # Output size (if result file exists)
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
      results_run+=("Run $r: ${timeSec}s, ${size}KB")
    else
      size="0"
      echo "::error::$resultFile was not created by $name"
      results_run+=("Run $r: FAILED, 0KB")
      stderr_all+=$'\n'"Run $r stderr:\n$stderr"
    fi
  done

  # Average time and size
  if [[ $successful_runs -gt 0 ]]; then
    avgTime=$(awk "BEGIN {printf \"%.2f\", $sumTime/$successful_runs}")
    avgSize=$(awk "BEGIN {printf \"%.1f\", $sumSize/$successful_runs}")
  else
    avgTime="0"
    avgSize="0"
  fi

  # Add to markdown table
  resultsMd+="|[${name}](${docsUrl})|${avgTime}|${avgSize}|\n"

  # Add details: show each run result
  for result in "${results_run[@]}"; do
    detailsMd+="$result\n"
  done

  # Show output of last run if successful
  if [[ -f "$resultFile" ]]; then
    json=$(cat "$resultFile")
    # Pretty print JSON if jq is available
    if command -v jq >/dev/null; then
      pretty_json=$(jq . "$resultFile")
      detailsMd+="\n\n\`\`\`json\n${pretty_json}\n\`\`\`"
    else
      detailsMd+="\n\n\`\`\`json\n${json}\n\`\`\`"
    fi
  else
    detailsMd+="No result file found from last run.\n\n"
    detailsMd+="\`\`\`\n${stderr_all}\n\`\`\`"
  fi

  detailsMd+="\n</details>\n\n"
  echo "::endgroup::"
done

# Write results
echo -e "$resultsMd" > results.md
echo -e "$detailsMd" > details.md

# Add summary to README.md
{
  echo "# Benchmark Results"
  echo
  echo "- **Number of runs per tool:** $runs"
  echo "- **Number of files analyzed per run:** $file_count"
  echo
  cat results.md
  echo
  cat details.md
} > README.md

cat results.md

rm results.md
rm details.md
