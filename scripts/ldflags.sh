#!/bin/bash
set -euo pipefail

# Get version information from git
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
TAG=$(git describe --tags --exact-match 2>/dev/null || echo "unknown")
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Check if working directory is dirty
if [[ $(git status --porcelain 2>/dev/null) ]]; then
    DIRTY="dirty"
else
    DIRTY=""
fi

# Build ldflags
LDFLAGS="-X 'bennypowers.dev/cem/internal/version.Version=${VERSION}'"
LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitCommit=${COMMIT}'"
LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitTag=${TAG}'"
LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.BuildTime=${BUILD_TIME}'"
LDFLAGS="${LDFLAGS} -X 'bennypowers.dev/cem/internal/version.GitDirty=${DIRTY}'"

echo "${LDFLAGS}"