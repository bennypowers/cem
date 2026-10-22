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
META_FILE="$NOTES_DIR/$VERSION.json"
mkdir -p "$NOTES_DIR"

write_meta() {
  # Preserve the commit recorded when these notes were first generated.
  if [ -f "$META_FILE" ]; then
    return
  fi
  local base
  base=$(git rev-parse HEAD)
  printf '{\n  "version": "%s",\n  "base": "%s"\n}\n' "$VERSION" "$base" > "$META_FILE"
}

if [ -f "$NOTES_FILE" ]; then
  if [ ! -f "$META_FILE" ]; then
    echo "Error: $NOTES_FILE exists without $META_FILE" >&2
    echo "Remove the notes file and rerun 'make notes $VERSION'." >&2
    exit 1
  fi
  echo "Release notes already exist: $NOTES_FILE"
else
  rm -f "$META_FILE"
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
    if ! {
      cat "$SCRIPT_DIR/release-notes/prompt.md"
      printf '\n'
      cat "$CONTEXT"
    } | (
      # Run outside the repo so project instructions, MCP config, and file
      # tools cannot influence or read the working tree.
      cd "$WORKDIR"
      "${CLAUDE_BIN:-claude}" -p --output-format text ${MODEL_FLAGS[@]+"${MODEL_FLAGS[@]}"}
    ) > "$NOTES_FILE"; then
      rm -f "$NOTES_FILE"
      echo "Error: Claude failed to generate release notes." >&2
      exit 1
    fi

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

write_meta

if ! "$SCRIPT_DIR/release-notes/verify.sh" "$NOTES_FILE" "$VERSION"; then
  echo "Draft verification failed. Fix the issues in the editor." >&2
fi

echo
echo "Review the notes in your editor. Save an empty file to abort."
EDITOR_CMD=()
read -r -a EDITOR_CMD <<< "${EDITOR:-vi}"
"${EDITOR_CMD[@]}" "$NOTES_FILE"

if [ ! -s "$NOTES_FILE" ]; then
  rm -f "$NOTES_FILE" "$META_FILE"
  echo "Release notes aborted."
  exit 1
fi

"$SCRIPT_DIR/release-notes/verify.sh" "$NOTES_FILE" "$VERSION"

echo
echo "Draft saved to $NOTES_FILE"
echo "Run 'make release $VERSION' when the notes are ready."
