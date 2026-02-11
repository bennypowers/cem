#!/bin/bash
# Release script for CEM
# Creates version commit, pushes it, then uses gh to tag and create release
#
# Usage: ./scripts/release.sh <version|patch|minor|major>
# Example: ./scripts/release.sh v0.6.6
# Example: ./scripts/release.sh patch

set -e

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

  # Remove 'v' prefix if present
  CURRENT_VERSION="${LATEST_TAG#v}"

  # Parse version components
  if [[ ! "$CURRENT_VERSION" =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    echo "Error: Latest tag '$LATEST_TAG' is not a valid semver version"
    echo "Expected format: v0.0.0"
    exit 1
  fi

  MAJOR="${BASH_REMATCH[1]}"
  MINOR="${BASH_REMATCH[2]}"
  PATCH="${BASH_REMATCH[3]}"

  # Bump the appropriate component
  case "$BUMP_TYPE" in
    patch)
      PATCH=$((PATCH + 1))
      ;;
    minor)
      MINOR=$((MINOR + 1))
      PATCH=0
      ;;
    major)
      MAJOR=$((MAJOR + 1))
      MINOR=0
      PATCH=0
      ;;
  esac

  VERSION="v${MAJOR}.${MINOR}.${PATCH}"
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

echo "Step 2: Updating version files and committing..."
./scripts/version.sh "$VERSION" || {
  echo ""
  echo "Error: Version update was rejected or failed"
  echo "Release flow aborted"
  exit 1
}
echo ""

# Strip 'v' prefix for version comparison
CHECK_VERSION="${VERSION#v}"

echo "Step 3: Verifying version sync across all packages..."
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

echo "Step 4: Pushing version commit..."
git push
echo ""

echo "Step 5: Creating GitHub release (gh will tag and push)..."
gh release create "$VERSION"
