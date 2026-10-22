#!/bin/bash
# Generate a draft of the release notes, verify it, and open it for review.
#
# Usage: notes.sh <version|patch|minor|major>
#
# Environment:
#   CLAUDE_BIN               Claude CLI binary (default: claude)
#   RELEASE_NOTES_MODEL      Optional model override for Claude
#   RELEASE_NOTES_TIMEOUT    Claude timeout in seconds (default: 600)
#   EDITOR                   Editor command, with arguments (default: vi)
#   RELEASE_NOTES_NO_EDIT    Set to 1 to skip the editor
#
# Notes are untracked local drafts under docs/releases/. Published notes live
# in GitHub releases.

set -euo pipefail

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
# shellcheck source=scripts/lib/release-version.sh
source "$SCRIPT_DIR/lib/release-version.sh"
# shellcheck source=scripts/lib/release-notes.sh
source "$SCRIPT_DIR/lib/release-notes.sh"

usage() {
  cat >&2 <<'EOF'
Usage: notes.sh <version|patch|minor|major>

  notes.sh v0.11.4
  notes.sh 0.11.4
  notes.sh patch

Environment:
  CLAUDE_BIN               Claude CLI binary (default: claude)
  RELEASE_NOTES_MODEL      Optional model override for Claude
  RELEASE_NOTES_TIMEOUT    Claude timeout in seconds (default: 600)
  EDITOR                   Editor command, with arguments (default: vi)
  RELEASE_NOTES_NO_EDIT    Set to 1 to skip the editor
EOF
}

case "${1:-}" in
  -h|--help)
    usage
    exit 0
    ;;
esac

INPUT="${1:-}"
if [ -z "$INPUT" ]; then
  usage
  exit 1
fi

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

LATEST_TAG=$(latest_release_tag)
VERSION=$(resolve_release_version "$INPUT" "$LATEST_TAG") || exit 1

NOTES_DIR=$(release_notes_dir "$REPO_ROOT")
NOTES_FILE=$(release_notes_file "$REPO_ROOT" "$VERSION")
META_FILE=$(release_notes_meta_file "$REPO_ROOT" "$VERSION")
mkdir -p "$NOTES_DIR"

LOCK_DIR="$NOTES_DIR/.$VERSION.lock"
if ! mkdir "$LOCK_DIR" 2>/dev/null; then
  echo "Error: another make notes run is active for $VERSION" >&2
  echo "Remove $LOCK_DIR if no other run is active." >&2
  exit 1
fi

CONTEXT=""
WORKDIR=""
TMP_NOTES=""
cleanup() {
  [ -n "$CONTEXT" ] && rm -f "$CONTEXT"
  [ -n "$WORKDIR" ] && rm -rf "$WORKDIR"
  [ -n "$TMP_NOTES" ] && rm -f "$TMP_NOTES"
  rmdir "$LOCK_DIR" 2>/dev/null || true
}
trap cleanup EXIT

# Capture the commit the notes describe before the slow generation step.
BASE_SHA=$(git rev-parse HEAD)

if [ -f "$NOTES_FILE" ]; then
  if [ ! -f "$META_FILE" ]; then
    echo "Error: $NOTES_FILE exists without $META_FILE" >&2
    echo "Remove the notes file and rerun 'make notes $VERSION'." >&2
    exit 1
  fi
  echo "Release notes already exist: $NOTES_FILE"

  EXISTING_BASE=$(read_release_notes_base "$META_FILE") || EXISTING_BASE=""
  if [ -n "$EXISTING_BASE" ]; then
    set +e
    check_release_notes_freshness "$EXISTING_BASE"
    freshness=$?
    set -e
    if [ "$freshness" -ne 0 ]; then
      echo "Warning: code commits landed since these notes were generated." >&2
      echo "Remove $NOTES_FILE and $META_FILE, then rerun to regenerate." >&2
    fi
  fi
else
  rm -f "$META_FILE"
  CONTEXT=$(mktemp)
  WORKDIR=$(mktemp -d)

  echo "Building release context for $VERSION..."
  "$SCRIPT_DIR/release-notes/context.sh" "$VERSION" > "$CONTEXT"

  MODEL_FLAGS=()
  if [ -n "${RELEASE_NOTES_MODEL:-}" ]; then
    MODEL_FLAGS=(--model "$RELEASE_NOTES_MODEL")
  fi

  TMP_NOTES=$(mktemp "$NOTES_DIR/.$VERSION.XXXXXX")
  if command -v "${CLAUDE_BIN:-claude}" >/dev/null 2>&1; then
    echo "Generating release notes with Claude..."
    TIMEOUT_FLAGS=()
    if command -v timeout >/dev/null 2>&1; then
      TIMEOUT_FLAGS=(timeout "${RELEASE_NOTES_TIMEOUT:-600}")
    fi

    if ! {
      cat "$SCRIPT_DIR/release-notes/prompt.md"
      printf '\n<untrusted-context>\n'
      cat "$CONTEXT"
      printf '\n</untrusted-context>\n'
    } | (
      # Run outside the repo so project instructions, MCP config, and file
      # tools cannot influence or read the working tree.
      cd "$WORKDIR"
      "${TIMEOUT_FLAGS[@]+"${TIMEOUT_FLAGS[@]}"}" \
        "${CLAUDE_BIN:-claude}" -p --output-format text \
        ${MODEL_FLAGS[@]+"${MODEL_FLAGS[@]}"}
    ) > "$TMP_NOTES"; then
      rm -f "$TMP_NOTES"
      echo "Error: Claude failed to generate release notes." >&2
      exit 1
    fi

    if [ ! -s "$TMP_NOTES" ]; then
      rm -f "$TMP_NOTES"
      echo "Error: Claude produced no output." >&2
      exit 1
    fi
  else
    echo "claude CLI not found; falling back to GitHub generated notes." >&2
    if ! github_generated_notes "$VERSION" "$BASE_SHA" > "$TMP_NOTES"; then
      rm -f "$TMP_NOTES"
      echo "Error: failed to generate release notes." >&2
      exit 1
    fi
  fi

  mv "$TMP_NOTES" "$NOTES_FILE"
fi

if [ ! -f "$META_FILE" ]; then
  write_release_notes_meta "$META_FILE" "$VERSION" "$BASE_SHA"
fi

if ! "$SCRIPT_DIR/release-notes/verify.sh" "$NOTES_FILE" "$VERSION"; then
  echo "Draft verification failed. Fix the issues in the editor." >&2
fi

if [ "${RELEASE_NOTES_NO_EDIT:-0}" = "1" ] || [ ! -t 0 ]; then
  echo
  echo "Non-interactive: draft saved to $NOTES_FILE"
else
  echo
  echo "Review the notes in your editor. Save an empty file to abort."
  EDITOR_CMD=()
  read -r -a EDITOR_CMD <<< "${EDITOR:-vi}"
  if ! "${EDITOR_CMD[@]}" "$NOTES_FILE"; then
    rm -f "$NOTES_FILE" "$META_FILE"
    echo "Editor failed; release notes discarded." >&2
    exit 1
  fi
fi

if [ ! -s "$NOTES_FILE" ]; then
  rm -f "$NOTES_FILE" "$META_FILE"
  echo "Release notes aborted."
  exit 1
fi

"$SCRIPT_DIR/release-notes/verify.sh" "$NOTES_FILE" "$VERSION"

echo
echo "Draft saved to $NOTES_FILE"
echo "Run 'make release $VERSION' when the notes are ready."
