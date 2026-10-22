#!/bin/bash
# Build the context used to generate release notes.
#
# Usage: context.sh <version>
#
# Author fields are intentionally omitted from commit and pull request data.

set -euo pipefail

VERSION="${1:-}"
if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>" >&2
  exit 1
fi

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

PREV_TAG=$(git describe --tags --abbrev=0 2>/dev/null || true)
if [ -z "$PREV_TAG" ]; then
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
      gh pr view "$number" \
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
  gh api --method POST "repos/{owner}/{repo}/releases/generate-notes" \
    -f tag_name="$VERSION" \
    -f target_commitish="$HEAD_SHA" \
    --jq .body 2>/dev/null || echo "(GitHub generated notes unavailable)"
  echo
  echo "## Previous releases"
  echo

  if command -v gh >/dev/null 2>&1; then
    {
      gh release list --limit 5 --json tagName,publishedAt \
        --jq '.[] | [.tagName, .publishedAt] | @tsv' 2>/dev/null || true
    } | while IFS=$'\t' read -r tag published; do
      [ -n "$tag" ] || continue
      echo "### $tag (${published%%T*})"
      gh release view "$tag" --json body --jq .body 2>/dev/null || true
      echo
    done
  else
    echo "(gh not available)"
    echo
  fi
}
