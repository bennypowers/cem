#!/bin/bash
# Build the context used to generate release notes.
#
# Usage: context.sh <version>
#
# Author fields are intentionally omitted from commit and pull request data.

set -euo pipefail

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
# shellcheck source=scripts/lib/release-version.sh
source "$SCRIPT_DIR/../lib/release-version.sh"
# shellcheck source=scripts/lib/release-notes.sh
source "$SCRIPT_DIR/../lib/release-notes.sh"

VERSION="${1:-}"
if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>" >&2
  exit 1
fi

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

PREV_TAG=$(latest_release_tag)
if [ "$PREV_TAG" = "v0.0.0" ]; then
  PREV_TAG=$(git rev-list --max-parents=0 HEAD)
fi

HEAD_SHA=$(git rev-parse HEAD)
PR_NUMBERS=$(git log --no-merges --pretty=format:'%s' "$PREV_TAG..HEAD" \
  | grep -oE '\(#[0-9]+\)' \
  | tr -d '(#)' \
  | sort -n -u \
  || true)

{
  echo "# Release context"
  echo
  echo "- Version: $VERSION"
  echo "- Previous tag: $PREV_TAG"
  echo "- Date: $(date -u +%Y-%m-%d)"
  echo
  echo "## Commits since $PREV_TAG"
  echo
  git log --no-merges --pretty=format:'- %h %s' "$PREV_TAG..HEAD" || true
  echo
  echo
  echo "## Pull requests merged since $PREV_TAG"
  echo

  if [ -n "$PR_NUMBERS" ]; then
    while IFS= read -r number; do
      [ -n "$number" ] || continue
      echo "### #$number"
      run_gh pr view "$number" \
        --json number,title,body,labels \
        --jq '"Title: \(.title)\nLabels: \([.labels[].name] | join(", "))\n\n\((.body // "")[0:2000])"' \
        2>/dev/null || echo "(pull request metadata unavailable)"
      echo
    done <<< "$PR_NUMBERS"
  else
    echo "(no pull request references found in commits)"
    echo
  fi

  echo "## Diffstat"
  echo
  git diff --stat "$PREV_TAG..HEAD" || true
  echo
  echo "## GitHub generated notes"
  echo
  github_generated_notes "$VERSION" "$HEAD_SHA" 2>/dev/null \
    || echo "(GitHub generated notes unavailable)"
  echo
  echo "## Previous releases"
  echo

  if command -v gh >/dev/null 2>&1; then
    RELEASES=$(run_gh release list --limit 5 --json tagName,publishedAt \
      --jq '.[] | [.tagName, .publishedAt] | @tsv' 2>/dev/null) || RELEASES=""
    if [ -z "$RELEASES" ]; then
      echo "(previous releases unavailable)"
      echo
    else
      while IFS=$'\t' read -r tag published; do
        [ -n "$tag" ] || continue
        echo "### $tag (${published%%T*})"
        run_gh release view "$tag" --json body --jq .body 2>/dev/null || echo "(release notes unavailable)"
        echo
      done <<< "$RELEASES"
    fi
  else
    echo "(gh not available)"
    echo
  fi
}
