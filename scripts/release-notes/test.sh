#!/bin/bash
# Regression tests for the release-notes workflow.
#
# Usage: test.sh

set -euo pipefail

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
REPO_ROOT=$(cd "$SCRIPT_DIR/../.." && pwd)
TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

pass() { echo "  ok: $*"; }
fail() { echo "  FAIL: $*" >&2; exit 1; }
assert_eq() {
  local actual="$1" expected="$2" message="$3"
  [ "$actual" = "$expected" ] || fail "$message (expected '$expected', got '$actual')"
}

echo "version resolution"
# shellcheck source=scripts/lib/release-version.sh
source "$REPO_ROOT/scripts/lib/release-version.sh"
assert_eq "$(resolve_release_version v0.11.4)" "v0.11.4" "v-prefixed explicit version"
assert_eq "$(resolve_release_version 0.11.4)" "v0.11.4" "bare explicit version"
assert_eq "$(resolve_release_version patch v0.11.3)" "v0.11.4" "patch bump"
assert_eq "$(resolve_release_version minor v0.11.3)" "v0.12.0" "minor bump"
assert_eq "$(resolve_release_version major v0.11.3)" "v1.0.0" "major bump"
if resolve_release_version nope >/dev/null 2>&1; then
  fail "invalid version was accepted"
fi
pass "explicit, bump, and invalid versions"

echo "git helpers"
(
  cd "$TMP"
  git init -q
  git config user.email test@example.com
  git config user.name Test
  printf 'one\n' > file.txt
  git add file.txt
  git commit -qm 'chore: init'
  git branch --show-current > "$TMP/branch"
  git tag v1.2.3

  assert_eq "$(latest_release_tag)" "v1.2.3" "latest tag"

  # shellcheck source=scripts/lib/release-notes.sh
  source "$REPO_ROOT/scripts/lib/release-notes.sh"

  base=$(git rev-parse HEAD)
  check_release_notes_freshness "$base" || fail "current notes were reported stale"

  printf 'two\n' > file.txt
  git add file.txt
  git commit -qm 'feat: code change'
  set +e
  check_release_notes_freshness "$base"
  status=$?
  set -e
  [ "$status" -eq 1 ] || fail "code commit was not reported stale"

  newer=$(git rev-parse HEAD)
  mkdir -p docs/releases
  printf 'notes\n' > docs/releases/v1.2.4.md
  git add docs/releases/v1.2.4.md
  git commit -qm 'docs: notes only'
  check_release_notes_freshness "$newer" || fail "notes-only commit was reported stale"

  set +e
  check_release_notes_freshness "does-not-exist"
  status=$?
  set -e
  [ "$status" -eq 2 ] || fail "invalid base did not return 2"

  git checkout -q v1.2.3
  set +e
  check_release_notes_freshness "$newer"
  status=$?
  set -e
  [ "$status" -eq 1 ] || fail "non-ancestor base was not reported stale"
)
pass "freshness, notes-only commits, and non-ancestor bases"

echo "metadata"
(
  cd "$TMP"
  # shellcheck source=scripts/lib/release-notes.sh
  source "$REPO_ROOT/scripts/lib/release-notes.sh"
  write_release_notes_meta "$TMP/meta.json" "v1.2.4" "abc123"
  assert_eq "$(read_release_notes_base "$TMP/meta.json")" "abc123" "meta round trip"
)
pass "metadata writer and reader"

echo "verify references"
(
  cd "$TMP"
  git checkout -q "$(cat "$TMP/branch")"

  printf 'one\n' > pr-one.txt
  git add pr-one.txt
  git commit -qm 'feat: one (#123)'
  printf 'two\n' > pr-two.txt
  git add pr-two.txt
  git commit -qm 'feat: two (#999)'

  stub="$TMP/gh-stub"
  mkdir -p "$stub"
  cat > "$stub/gh" <<'STUB'
#!/usr/bin/env bash
case "$*" in
  *"issues/123"*) echo "2026-01-01T00:00:00Z" ;;
  *"issues/999"*) echo "gh: Not Found (HTTP 404)" >&2; exit 1 ;;
  *) echo "gh: Server Error (HTTP 500)" >&2; exit 1 ;;
esac
STUB
  chmod +x "$stub/gh"

  printf '## What is changed\n\n- Fix (#123)\n' > "$TMP/good.md"
  PATH="$stub:$PATH" "$REPO_ROOT/scripts/release-notes/verify.sh" "$TMP/good.md" v1.2.4 \
    || fail "merged reference failed verification"

  printf '## What is changed\n\n- Fix (#999)\n' > "$TMP/missing.md"
  if PATH="$stub:$PATH" "$REPO_ROOT/scripts/release-notes/verify.sh" "$TMP/missing.md" v1.2.4 >/dev/null 2>&1; then
    fail "missing reference passed verification"
  fi

  printf '## What is changed\n\n- Color from #1e90ff to #333 and todo-list support\n' > "$TMP/noise.md"
  PATH="$stub:$PATH" "$REPO_ROOT/scripts/release-notes/verify.sh" "$TMP/noise.md" v1.2.4 \
    || fail "hex colors or todo-list were treated as references or placeholders"

  printf '## What is changed\n\n- TODO\n' > "$TMP/placeholder.md"
  if PATH="$stub:$PATH" "$REPO_ROOT/scripts/release-notes/verify.sh" "$TMP/placeholder.md" v1.2.4 >/dev/null 2>&1; then
    fail "placeholder passed verification"
  fi
)
pass "known PRs only, 404 handling, hex colors, and placeholders"

echo "end-to-end make notes"
(
  e2e="$TMP/e2e"
  mkdir -p "$e2e"
  cd "$e2e"
  git init -q
  git config user.email test@example.com
  git config user.name Test
  mkdir -p scripts
  cp -R "$REPO_ROOT/scripts/lib" scripts/lib
  cp -R "$REPO_ROOT/scripts/release-notes" scripts/release-notes
  cp "$REPO_ROOT/scripts/notes.sh" scripts/notes.sh
  printf 'initial\n' > README.md
  git add README.md
  git commit -qm 'feat: initial work (#123)'

  stub="$TMP/stubs"
  mkdir -p "$stub"
  cat > "$stub/gh" <<'STUB'
#!/usr/bin/env bash
case "$*" in
  *"pr view"*) printf 'Title: Test PR\nLabels: \n\nBody\n' ;;
  *"release list"*) ;;
  *"api"*) printf 'Generated notes\n' ;;
  *) exit 0 ;;
esac
STUB
  cat > "$stub/claude" <<'STUB'
#!/usr/bin/env bash
cat >/dev/null
printf '## What is changed\n\n- Initial work (#123)\n'
STUB
  chmod +x "$stub/gh" "$stub/claude"

  PATH="$stub:$PATH" RELEASE_NOTES_NO_EDIT=1 CLAUDE_BIN="$stub/claude" \
    "$e2e/scripts/notes.sh" v1.2.4 >/dev/null

  [ -f docs/releases/v1.2.4.md ] || fail "notes file was not created"
  [ -f docs/releases/v1.2.4.json ] || fail "metadata file was not created"
  assert_eq "$(jq -r .base docs/releases/v1.2.4.json)" "$(git rev-parse HEAD)" "metadata base"

  PATH="$stub:$PATH" RELEASE_NOTES_NO_EDIT=1 CLAUDE_BIN="$stub/claude" \
    "$e2e/scripts/notes.sh" v1.2.4 >/dev/null
  assert_eq "$(jq -r .base docs/releases/v1.2.4.json)" "$(git rev-parse HEAD)" "rerun preserves base"

  if PATH="$stub:$PATH" RELEASE_NOTES_FORCE_EDIT=1 EDITOR=false CLAUDE_BIN="$stub/claude" \
    "$e2e/scripts/notes.sh" v1.2.4 >/dev/null 2>&1; then
    fail "editor failure was ignored for existing notes"
  fi
  [ -f docs/releases/v1.2.4.md ] || fail "existing notes were removed on editor failure"

  if PATH="$stub:$PATH" RELEASE_NOTES_FORCE_EDIT=1 EDITOR=false CLAUDE_BIN="$stub/claude" \
    "$e2e/scripts/notes.sh" v1.2.5 >/dev/null 2>&1; then
    fail "editor failure was ignored for a new draft"
  fi
  [ ! -f docs/releases/v1.2.5.md ] || fail "new draft was not removed on editor failure"
  [ ! -f docs/releases/v1.2.5.json ] || fail "new metadata was not removed on editor failure"

  rm docs/releases/v1.2.4.json
  if PATH="$stub:$PATH" RELEASE_NOTES_NO_EDIT=1 CLAUDE_BIN="$stub/claude" \
    "$e2e/scripts/notes.sh" v1.2.4 >/dev/null 2>&1; then
    fail "notes without metadata were accepted"
  fi
)
pass "generation, metadata, rerun, and missing-metadata error"

echo "all release-notes tests passed"
