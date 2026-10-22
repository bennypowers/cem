#!/bin/bash
# Release script for CEM
# Creates version commit, pushes it, then uses gh to tag and create release
#
# Usage: ./scripts/release.sh <version|patch|minor|major>
# Example: ./scripts/release.sh v0.6.6
# Example: ./scripts/release.sh patch

set -e

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
# shellcheck source=scripts/lib/release-version.sh
source "$SCRIPT_DIR/lib/release-version.sh"
# shellcheck source=scripts/lib/release-notes.sh
source "$SCRIPT_DIR/lib/release-notes.sh"

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

if [ -z "$1" ]; then
  echo "Error: VERSION or bump type is required"
  echo "Usage: $0 <version|patch|minor|major>"
  echo "  $0 v0.6.6   - Release explicit version"
  echo "  $0 patch    - Bump patch version (0.0.x)"
  echo "  $0 minor    - Bump minor version (0.x.0)"
  echo "  $0 major    - Bump major version (x.0.0)"
  exit 1
fi

INPUT="$1"

# Check if input is a bump type (patch/minor/major)
if [[ "$INPUT" =~ ^(patch|minor|major)$ ]]; then
  BUMP_TYPE="$INPUT"

  # Get the latest tag
  LATEST_TAG=$(latest_release_tag)
  echo "Latest tag: $LATEST_TAG"

  VERSION=$(resolve_release_version "$BUMP_TYPE" "$LATEST_TAG") || exit 1
  echo "Bumping $BUMP_TYPE: $LATEST_TAG → $VERSION"
  echo ""
else
  # Use explicit version
  VERSION="$INPUT"
fi

echo "Checking if tag $VERSION already exists..."
if git rev-parse "$VERSION" >/dev/null 2>&1; then
  echo "Error: Tag $VERSION already exists"
  echo "Use 'git tag -d $VERSION' to delete locally if needed"
  exit 1
fi
echo "✓ Tag $VERSION does not exist"
echo ""

if git remote get-url origin >/dev/null 2>&1; then
  if git ls-remote --tags origin "refs/tags/$VERSION" 2>/dev/null | grep -q .; then
    echo "Error: Tag $VERSION already exists on origin"
    exit 1
  fi
fi

NOTES_FILE=$(release_notes_file "$REPO_ROOT" "$VERSION")
META_FILE=$(release_notes_meta_file "$REPO_ROOT" "$VERSION")
if [ ! -f "$NOTES_FILE" ]; then
  echo "Error: release notes not found: $NOTES_FILE"
  echo "Run 'make notes $VERSION' first, then review the draft."
  exit 1
fi
if [ ! -f "$META_FILE" ]; then
  echo "Error: release notes metadata not found: $META_FILE"
  echo "Run 'make notes $VERSION' to regenerate the notes and metadata."
  exit 1
fi
"$SCRIPT_DIR/release-notes/verify.sh" "$NOTES_FILE" "$VERSION"

BASE_SHA=$(read_release_notes_base "$META_FILE") || BASE_SHA=""
if [ -z "$BASE_SHA" ]; then
  echo "Error: $META_FILE has no valid base commit"
  echo "Run 'make notes $VERSION' to regenerate the notes and metadata."
  exit 1
fi

set +e
check_release_notes_freshness "$BASE_SHA"
freshness=$?
set -e
if [ "$freshness" -eq 1 ]; then
  echo "Error: release notes are stale"
  if git merge-base --is-ancestor "$BASE_SHA" HEAD 2>/dev/null; then
    echo "Commits landed after the notes were generated:"
    git log --oneline "$BASE_SHA..HEAD" -- . ':(exclude)docs/releases'
  else
    echo "The notes were generated from a commit that is not an ancestor of HEAD."
  fi
  echo ""
  echo "Remove $NOTES_FILE and $META_FILE, then rerun 'make notes $VERSION'."
  exit 1
elif [ "$freshness" -ne 0 ]; then
  echo "Error: cannot verify release notes base $BASE_SHA"
  exit 1
fi

echo "✓ Release notes found: $NOTES_FILE"
echo "✓ Release notes are current"
echo ""

echo "Creating release $VERSION..."
echo ""

echo "Step 1: Verifying npm packaging structure..."
make verify-npm || {
  echo ""
  echo "Error: npm packaging verification failed"
  echo "Fix the issue above before releasing."
  exit 1
}
echo ""

echo "Step 2: Ensuring generated files are up to date..."
go generate ./...
if ! git diff --quiet serve/middleware/routes/templates/; then
  echo "  Generated files have changed, committing..."
  git add serve/middleware/routes/templates/
  git commit -m "chore: regenerate embedded assets"
  write_release_notes_meta "$META_FILE" "$VERSION" "$(git rev-parse HEAD)"
  echo "  Done."
else
  echo "  Generated files are up to date."
fi
echo ""

echo "Step 3: Updating version files and committing..."
./scripts/version.sh "$VERSION" || {
  echo ""
  echo "Error: Version update was rejected or failed"
  echo "Release flow aborted"
  exit 1
}
write_release_notes_meta "$META_FILE" "$VERSION" "$(git rev-parse HEAD)"
echo ""

# Strip 'v' prefix for version comparison
CHECK_VERSION="${VERSION#v}"

echo "Step 4: Verifying version sync across all packages..."
MISMATCH=0

VSCODE_VERSION=$(node -e "process.stdout.write(JSON.parse(require('fs').readFileSync('extensions/vscode/package.json','utf8')).version)")
ZED_VERSION=$(grep '^version = ' extensions/zed/extension.toml | cut -d'"' -f2)
NPM_VERSION=$(node -e "process.stdout.write(JSON.parse(require('fs').readFileSync('npm/package.json','utf8')).version)")

CLAUDE_VERSION=""
if [ -f .claude-plugin/marketplace.json ]; then
  CLAUDE_VERSION=$(node -e "process.stdout.write(JSON.parse(require('fs').readFileSync('.claude-plugin/marketplace.json','utf8')).plugins[0].version)")
fi

for PAIR in "extensions/vscode/package.json:$VSCODE_VERSION" "extensions/zed/extension.toml:$ZED_VERSION" "npm/package.json:$NPM_VERSION" ".claude-plugin/marketplace.json:$CLAUDE_VERSION"; do
  FILE="${PAIR%%:*}"
  FILE_VERSION="${PAIR##*:}"
  if [ ! -f "$FILE" ]; then
    echo "  - $FILE: skipped (file not found)"
  elif [ -n "$FILE_VERSION" ] && [ "$FILE_VERSION" != "$CHECK_VERSION" ]; then
    echo "  ✗ $FILE has version $FILE_VERSION (expected $CHECK_VERSION)"
    MISMATCH=1
  else
    echo "  ✓ $FILE: $FILE_VERSION"
  fi
done

if [ "$MISMATCH" -eq 1 ]; then
  echo ""
  echo "Error: version mismatch detected after running version.sh"
  echo "Release aborted. Fix version.sh and try again."
  exit 1
fi
echo ""

echo "Step 5: Pushing version commit..."
git push
echo ""

echo "Step 6: Creating GitHub release (gh will tag and push)..."
run_gh release create "$VERSION" \
  --target "$(git rev-parse HEAD)" \
  --notes-file "$NOTES_FILE"
