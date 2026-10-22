#!/bin/bash
# Verify generated release notes.
#
# Usage: verify.sh <notes-file> [version]
#
# Checks for placeholders, headings, and missing or unmerged pull request
# references.

set -euo pipefail

NOTES_FILE="${1:-}"
VERSION="${2:-}"

if [ -z "$NOTES_FILE" ] || [ ! -f "$NOTES_FILE" ]; then
  echo "Usage: $0 <notes-file> [version]" >&2
  exit 1
fi

errors=0

fail() {
  echo "Error: $*" >&2
  errors=$((errors + 1))
}

if [ ! -s "$NOTES_FILE" ]; then
  fail "notes file is empty"
fi

if ! grep -qE '^#{1,3} ' "$NOTES_FILE"; then
  fail "notes have no Markdown heading"
fi

if grep -qiE 'TODO|TBD|FIXME|lorem ipsum|<insert' "$NOTES_FILE"; then
  fail "notes contain placeholders"
fi

if command -v gh >/dev/null 2>&1; then
  while IFS= read -r number; do
    [ -n "$number" ] || continue
    # Issues are allowed. Pull requests must be merged.
    if ! response=$(gh api "repos/{owner}/{repo}/issues/$number" \
      --jq 'if .pull_request then (.pull_request.merged_at // "unmerged") else "issue" end' \
      2>&1); then
      if grep -q 'HTTP 404' <<< "$response"; then
        state="missing"
      else
        state="error"
      fi
    else
      state="$response"
    fi
    if [ "$state" = "missing" ]; then
      fail "referenced #$number does not exist"
    elif [ "$state" = "error" ]; then
      fail "failed to verify referenced #$number"
    elif [ "$state" = "unmerged" ]; then
      fail "referenced pull request #$number is not merged"
    fi
  done < <(grep -oE '#[0-9]+' "$NOTES_FILE" | tr -d '#' | sort -n -u || true)
else
  echo "Warning: gh not found; skipping pull request reference checks" >&2
fi

if [ "$errors" -ne 0 ]; then
  exit 1
fi

echo "✓ Release notes verified${VERSION:+ for $VERSION}"
