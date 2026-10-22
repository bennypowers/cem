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
  LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
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

NOTES_FILE="docs/releases/$VERSION.md"
if [ ! -f "$NOTES_FILE" ]; then
  echo "Error: release notes not found: $NOTES_FILE"
  echo "Run 'make notes $VERSION' first, then review and commit the notes."
  exit 1
fi
if ! git ls-files --error-unmatch "$NOTES_FILE" >/dev/null 2>&1; then
  echo "Error: $NOTES_FILE is not committed"
  echo "Commit the release notes before releasing."
  exit 1
fi
if [ -n "$(git status --porcelain -- "$NOTES_FILE")" ]; then
  echo "Error: $NOTES_FILE has uncommitted changes"
  echo "Commit the release notes before releasing."
  exit 1
fi
echo "✓ Release notes found: $NOTES_FILE"
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
gh release create "$VERSION" --notes-file "$NOTES_FILE"
