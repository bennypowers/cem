#!/usr/bin/env bash
set -euo pipefail

# Generate a PR-scoped health report by detecting changed files.
# Expects env vars: PACKAGE_PATH, HEALTH_ARGS

BASE_REF="${GITHUB_BASE_REF:-main}"

changed_files=$(git diff --name-only "origin/${BASE_REF}...HEAD" 2>/dev/null \
             || git diff --name-only "${BASE_REF}...HEAD" 2>/dev/null \
             || echo "")

if [ -z "$changed_files" ]; then
  printf '### Documentation Health\n\nNo changed files detected.\n' > health_report.md
  exit 0
fi

# Pass changed files as positional args — cem health resolves
# source extensions and package.json exports internally.
files=()
while IFS= read -r f; do
  files+=("$f")
done <<< "$changed_files"

printf '%s\n' "${files[@]}" > changed_files.txt
# shellcheck disable=SC2086
cem health -p "$PACKAGE_PATH" --format markdown $HEALTH_ARGS "${files[@]}" > health_report.md
