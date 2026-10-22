#!/bin/bash
# Shared helpers for the release-notes workflow.
# Source this file; do not execute it directly.

release_notes_dir() {
  printf '%s/docs/releases' "$1"
}

release_notes_file() {
  printf '%s/docs/releases/%s.md' "$1" "$2"
}

release_notes_meta_file() {
  printf '%s/docs/releases/%s.json' "$1" "$2"
}

run_gh() {
  if command -v timeout >/dev/null 2>&1; then
    timeout "${GH_TIMEOUT:-60}" gh "$@"
  else
    gh "$@"
  fi
}

github_generated_notes() {
  local tag="$1"
  local sha="$2"
  run_gh api --method POST "repos/{owner}/{repo}/releases/generate-notes" \
    -f tag_name="$tag" \
    -f target_commitish="$sha" \
    --jq .body
}

read_release_notes_base() {
  local meta="$1"
  if command -v jq >/dev/null 2>&1; then
    jq -r '.base // empty' "$meta"
  elif command -v node >/dev/null 2>&1; then
    node -e 'const fs=require("fs"); const m=JSON.parse(fs.readFileSync(process.argv[1],"utf8")); process.stdout.write(m.base || "")' "$meta"
  else
    return 1
  fi
}

write_release_notes_meta() {
  local meta="$1"
  local version="$2"
  local base="$3"
  printf '{\n  "version": "%s",\n  "base": "%s"\n}\n' "$version" "$base" > "$meta.tmp"
  mv "$meta.tmp" "$meta"
}

# Returns 0 when the notes are current, 1 when commits landed after the base,
# and 2 when the base cannot be checked.
check_release_notes_freshness() {
  local base="$1"

  git cat-file -e "$base^{commit}" 2>/dev/null || return 2
  git merge-base --is-ancestor "$base" HEAD 2>/dev/null || return 1

  local stale
  stale=$(git log --format=%H "$base..HEAD" -- . ':(exclude)docs/releases') || return 2
  [ -z "$stale" ] || return 1
  return 0
}
