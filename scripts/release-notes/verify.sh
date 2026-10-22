#!/bin/bash
# Verify generated release notes.
#
# Usage: verify.sh <notes-file> [version]
#
# Checks for placeholders, headings, and the state of pull requests that
# belong to this release range.

set -euo pipefail

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
# shellcheck source=scripts/lib/release-version.sh
source "$SCRIPT_DIR/../lib/release-version.sh"
# shellcheck source=scripts/lib/release-notes.sh
source "$SCRIPT_DIR/../lib/release-notes.sh"

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

if grep -qE '(^|[^[:alnum:]_])(TODO|TBD|FIXME)([^[:alnum:]_]|$)|[Ll]orem ipsum|<insert' "$NOTES_FILE"; then
  fail "notes contain placeholders"
fi

# Only validate references that belong to this release range. This avoids
# treating hex colors, URL fragments, code samples, and cross-repo references
# as issues in this repository.
PREV_TAG=$(latest_release_tag)
if [ "$PREV_TAG" = "v0.0.0" ]; then
  PREV_TAG=$(git rev-list --max-parents=0 HEAD)
fi
KNOWN_PRS=$(git log --no-merges --pretty=format:'%s' "$PREV_TAG..HEAD" 2>/dev/null \
  | grep -oE '\(#[0-9]+\)' \
  | tr -d '(#)' \
  | sort -n -u \
  || true)

if [ -n "$KNOWN_PRS" ]; then
  if command -v gh >/dev/null 2>&1; then
    while IFS= read -r number; do
      [ -n "$number" ] || continue
      if ! grep -qx "$number" <<< "$KNOWN_PRS"; then
        continue
      fi
      # Issues are allowed. Pull requests must be merged.
      if ! state=$(run_gh api "repos/{owner}/{repo}/issues/$number" \
        --jq 'if .pull_request then (.pull_request.merged_at // "unmerged") else "issue" end' \
        2>&1); then
        if grep -q 'HTTP 404' <<< "$state"; then
          state="missing"
        else
          state="error"
        fi
      fi
      if [ "$state" = "missing" ]; then
        fail "referenced #$number does not exist"
      elif [ "$state" = "error" ]; then
        fail "failed to verify referenced #$number"
      elif [ "$state" = "unmerged" ]; then
        fail "referenced pull request #$number is not merged"
      fi
    done < <(grep -oE '#[0-9]+' "$NOTES_FILE" | tr -d '#' | sort -n -u)
  else
    echo "Warning: gh not found; skipping pull request reference checks" >&2
  fi
fi

if [ "$errors" -ne 0 ]; then
  exit 1
fi

echo "✓ Release notes verified${VERSION:+ for $VERSION}"
