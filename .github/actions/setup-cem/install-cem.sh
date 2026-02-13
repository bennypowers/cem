#!/usr/bin/env bash
set -euo pipefail

# Install cem binary from GitHub releases.
# Expects env vars: PACKAGE_PATH, INPUT_VERSION, GH_TOKEN

MIN_VERSION="0.9.5"

version_ge() {
  printf '%s\n%s\n' "$2" "$1" | sort -V | head -n1 | grep -qx "$2"
}

VERSION="$INPUT_VERSION"

# If no explicit version, check package.json
if [ -z "$VERSION" ] && [ -f "${PACKAGE_PATH}/package.json" ]; then
  VERSION=$(jq -r '
    (.devDependencies["@pwrs/cem"] // .dependencies["@pwrs/cem"] // empty)
  ' "${PACKAGE_PATH}/package.json" 2>/dev/null \
    | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' || true)
fi

# If still no version, use latest release
if [ -z "$VERSION" ]; then
  VERSION=$(gh release view --repo bennypowers/cem --json tagName -q '.tagName' | sed 's/^v//')
fi

# Enforce minimum version
if ! version_ge "$VERSION" "$MIN_VERSION"; then
  echo "::warning::Requested cem version $VERSION is below minimum $MIN_VERSION, using $MIN_VERSION"
  VERSION="$MIN_VERSION"
fi

echo "Installing cem v${VERSION}"
curl -fsSL "https://github.com/bennypowers/cem/releases/download/v${VERSION}/cem-linux-x64" \
  -o /usr/local/bin/cem
chmod +x /usr/local/bin/cem
cem version
