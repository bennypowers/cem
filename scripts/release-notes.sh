#!/bin/bash
# Generate a draft of the release notes, verify it, and open it for review.
#
# Usage: release-notes.sh <version|patch|minor|major>
#
# Requires the Claude Code CLI for generation. When it is not installed,
# falls back to GitHub's generated release notes.

set -euo pipefail

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
# shellcheck source=scripts/lib/release-version.sh
source "$SCRIPT_DIR/lib/release-version.sh"

INPUT="${1:-}"
if [ -z "$INPUT" ]; then
  echo "Usage: $0 <version|patch|minor|major>" >&2
  echo "  $0 v0.11.4" >&2
  echo "  $0 patch" >&2
  exit 1
fi

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
VERSION=$(resolve_release_version "$INPUT" "$LATEST_TAG")

NOTES_DIR="$REPO_ROOT/docs/releases"
NOTES_FILE="$NOTES_DIR/$VERSION.md"
mkdir -p "$NOTES_DIR"

if [ -f "$NOTES_FILE" ]; then
  echo "Release notes already exist: $NOTES_FILE"
else
  CONTEXT=$(mktemp)
  WORKDIR=$(mktemp -d)
  trap 'rm -f "$CONTEXT"; rm -rf "$WORKDIR"' EXIT

  echo "Building release context for $VERSION..."
  "$SCRIPT_DIR/release-notes/context.sh" "$VERSION" > "$CONTEXT"

  MODEL_FLAGS=()
  if [ -n "${RELEASE_NOTES_MODEL:-}" ]; then
    MODEL_FLAGS=(--model "$RELEASE_NOTES_MODEL")
  fi

  if command -v "${CLAUDE_BIN:-claude}" >/dev/null 2>&1; then
    echo "Generating release notes with Claude..."
    {
      cat "$SCRIPT_DIR/release-notes/prompt.md"
      printf '\n'
      cat "$CONTEXT"
    } | (
      # Run outside the repo so project instructions, MCP config, and file
      # tools cannot influence or read the working tree.
      cd "$WORKDIR"
      "${CLAUDE_BIN:-claude}" -p --output-format text ${MODEL_FLAGS[@]+"${MODEL_FLAGS[@]}"}
    ) > "$NOTES_FILE"

    if [ ! -s "$NOTES_FILE" ]; then
      rm -f "$NOTES_FILE"
      echo "Error: Claude produced no output." >&2
      exit 1
    fi
  else
    echo "claude CLI not found; falling back to GitHub generated notes." >&2
    if ! gh api --method POST "repos/{owner}/{repo}/releases/generate-notes" \
      -f tag_name="$VERSION" \
      -f target_commitish="$(git rev-parse HEAD)" \
      --jq .body > "$NOTES_FILE"; then
      rm -f "$NOTES_FILE"
      echo "Error: failed to generate release notes." >&2
      exit 1
    fi
  fi
fi

"$SCRIPT_DIR/release-notes/verify.sh" "$NOTES_FILE" "$VERSION"

echo
echo "Review the notes in your editor. Save an empty file to abort."
"${EDITOR:-vi}" "$NOTES_FILE"

if [ ! -s "$NOTES_FILE" ]; then
  rm -f "$NOTES_FILE"
  echo "Release notes aborted."
  exit 1
fi

"$SCRIPT_DIR/release-notes/verify.sh" "$NOTES_FILE" "$VERSION"

if [ -n "$(git status --porcelain -- "$NOTES_FILE")" ]; then
  echo
  read -p "Commit release notes? (y/n) " -n 1 -r
  echo
  if [[ $REPLY =~ ^[Yy]$ ]]; then
    git add "$NOTES_FILE"
    git commit -m "docs: add release notes for $VERSION"
  else
    echo "Release notes left uncommitted. Commit them before running make release."
  fi
fi
