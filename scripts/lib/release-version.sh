#!/bin/bash
# Shared release version resolution.
# Source this file; do not execute it directly.

# resolve_release_version <version|patch|minor|major> [latest-tag]
#
# Prints a v-prefixed version to stdout. The caller is responsible for
# printing user-facing status messages.
resolve_release_version() {
  local input="${1:-}"
  local latest_tag="${2:-}"

  if [[ "$input" =~ ^(patch|minor|major)$ ]]; then
    if [[ -z "$latest_tag" ]]; then
      latest_tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
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
  else
    echo "$input"
  fi
}
