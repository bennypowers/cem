#!/bin/bash
# Shared release version resolution.
# Source this file; do not execute it directly.

# resolve_release_version <version|patch|minor|major> [latest-tag]
#
# Prints a v-prefixed version to stdout. The caller is responsible for
# printing user-facing status messages.
latest_release_tag() {
  git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"
}

resolve_release_version() {
  local input="${1:-}"
  local latest_tag="${2:-}"

  # Accept both vX.Y.Z and X.Y.Z, and always return the v-prefixed form.
  if [[ "$input" =~ ^v?([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    echo "v${BASH_REMATCH[1]}.${BASH_REMATCH[2]}.${BASH_REMATCH[3]}"
    return 0
  fi

  if [[ "$input" =~ ^(patch|minor|major)$ ]]; then
    if [[ -z "$latest_tag" ]]; then
      latest_tag=$(latest_release_tag)
    fi

    local current="${latest_tag#v}"
    if [[ ! "$current" =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
      echo "Error: Latest tag '$latest_tag' is not a valid semver version" >&2
      echo "Expected format: v0.0.0" >&2
      return 1
    fi

    local major="${BASH_REMATCH[1]}"
    local minor="${BASH_REMATCH[2]}"
    local patch="${BASH_REMATCH[3]}"

    case "$input" in
      patch)
        patch=$((patch + 1))
        ;;
      minor)
        minor=$((minor + 1))
        patch=0
        ;;
      major)
        major=$((major + 1))
        minor=0
        patch=0
        ;;
    esac

    echo "v${major}.${minor}.${patch}"
    return 0
  fi

  echo "Error: version must be X.Y.Z, vX.Y.Z, or patch|minor|major" >&2
  return 1
}
