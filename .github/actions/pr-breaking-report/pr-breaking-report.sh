#!/usr/bin/env bash
set -euo pipefail

# Generate a breaking change report comparing against the PR base ref.
# Expects env vars: PACKAGE_PATH, BREAKING_ARGS

BASE_REF="${GITHUB_BASE_REF:-main}"

# Ensure the base ref is available locally for git show
git fetch origin "$BASE_REF" --depth=1 2>/dev/null || true

# shellcheck disable=SC2086
if cem breaking -p "$PACKAGE_PATH" \
  --base "origin/${BASE_REF}" \
  --format markdown $BREAKING_ARGS > breaking_report.md; then
  # No breaking changes -- check if the report is empty/trivial
  if grep -q 'No breaking changes detected' breaking_report.md; then
    printf '### Breaking Change Report\n\nNo breaking changes detected.\n' > breaking_report.md
  fi
else
  # cem breaking may exit non-zero from --fail-on; the report is still valid
  true
fi
