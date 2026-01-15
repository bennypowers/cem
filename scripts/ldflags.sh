#!/bin/bash
set -euo pipefail

# Check if working directory has ANY changes (tracked or untracked)
if [[ $(git status --porcelain 2>/dev/null) ]]; then
    # Dirty repo - let Go's build info provide version with timestamp + dirty suffix
    # Only set metadata fields, not Version itself
    COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
    TAG=$(git describe --tags --exact-match 2>/dev/null || echo "unknown")
    BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

    LDFLAGS="-X 'bennypowers.dev/cem/internal/version.GitCommit=${COMMIT}'"
    LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitTag=${TAG}'"
    LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.BuildTime=${BUILD_TIME}'"
    LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitDirty=dirty'"
else
    # Clean repo - prefer exact tag, fallback to git describe
    VERSION=$(git describe --tags --exact-match 2>/dev/null || git describe --tags --always 2>/dev/null || echo "dev")
    COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
    TAG=$(git describe --tags --exact-match 2>/dev/null || echo "unknown")
    BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

    LDFLAGS="-X 'bennypowers.dev/cem/internal/version.Version=${VERSION}'"
    LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitCommit=${COMMIT}'"
    LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitTag=${TAG}'"
    LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.BuildTime=${BUILD_TIME}'"
    LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitDirty='"
fi

# Strip debug info and symbol table for smaller binaries
LDFLAGS="${LDFLAGS} -s -w"

echo "${LDFLAGS}"