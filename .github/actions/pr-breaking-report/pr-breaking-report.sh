#!/usr/bin/env bash
set -euo pipefail

# Generate a breaking change report comparing against the PR base ref.
# Expects env vars: PACKAGE_PATH, BREAKING_ARGS
# Optional: FAIL_ON (breaking severity threshold, passed through from workflow)

BASE_REF="${GITHUB_BASE_REF:-main}"

# Ensure the base ref is available locally for git show
git fetch origin "$BASE_REF" --depth=1 2>/dev/null || true

fail_on_args=""
if [ -n "${FAIL_ON:-}" ]; then
  fail_on_args="--fail-on $FAIL_ON"
fi

rc=0
# shellcheck disable=SC2086
cem breaking -p "$PACKAGE_PATH" \
  --base "origin/${BASE_REF}" \
  --format markdown $fail_on_args $BREAKING_ARGS > breaking_report.md || rc=$?

# Persist exit code so the workflow gate step can reuse it
echo "$rc" > breaking_exitcode.txt

if [ "$rc" -ne 0 ]; then
  if [ -s breaking_report.md ] && grep -q '### Breaking Change Report' breaking_report.md; then
    # --fail-on threshold hit: report was generated, non-zero exit is expected
    true
  else
    echo "cem breaking failed with exit code $rc" >&2
    exit "$rc"
  fi
fi

if grep -q 'No breaking changes detected' breaking_report.md; then
  printf '### Breaking Change Report\n\nNo breaking changes detected.\n' > breaking_report.md
fi
