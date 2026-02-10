#!/usr/bin/env bash
# health-pr-report.sh - Generate a filtered health report for changed modules in a PR
#
# Usage: scripts/health-pr-report.sh [package-path]
#
# Arguments:
#   package-path  Path to the package (default: ".")
#
# Environment:
#   GITHUB_BASE_REF  Base branch for diff (default: "main")
#
# Output:
#   health_report.md  Markdown health report for changed modules

set -euo pipefail

PACKAGE_PATH="${1:-.}"
BASE_REF="${GITHUB_BASE_REF:-main}"

# Map source file extensions to their compiled output extensions
map_extension() {
  local file="$1"
  case "$file" in
    *.tsx)  echo "${file%.tsx}.js" ;;
    *.ts)   echo "${file%.ts}.js" ;;
    *.mts)  echo "${file%.mts}.mjs" ;;
    *.cts)  echo "${file%.cts}.cjs" ;;
    *.scss) echo "${file%.scss}.css" ;;
    *.sass) echo "${file%.sass}.css" ;;
    *)      echo "$file" ;;
  esac
}

# Get changed files relative to base branch
changed_files=$(git diff --name-only "origin/${BASE_REF}...HEAD" 2>/dev/null || \
                git diff --name-only "${BASE_REF}...HEAD" 2>/dev/null || \
                echo "")

if [ -z "$changed_files" ]; then
  cat > health_report.md <<'EOF'
### Documentation Health

No changed files detected.
EOF
  exit 0
fi

# Get all module paths from the manifest
manifest_modules=$(cem health -p "$PACKAGE_PATH" --format json 2>/dev/null | \
                   jq -r '.modules[].path' 2>/dev/null || echo "")

if [ -z "$manifest_modules" ]; then
  cat > health_report.md <<'EOF'
### Documentation Health

No modules found in manifest.
EOF
  exit 0
fi

# Match changed files against module paths
declare -A matched_modules

while IFS= read -r changed_file; do
  # Map the extension
  mapped_file=$(map_extension "$changed_file")

  while IFS= read -r module_path; do
    if [ "$mapped_file" = "$module_path" ] || [ "$changed_file" = "$module_path" ]; then
      matched_modules["$module_path"]=1
    fi
  done <<< "$manifest_modules"
done <<< "$changed_files"

if [ ${#matched_modules[@]} -eq 0 ]; then
  cat > health_report.md <<'EOF'
### Documentation Health

No changed files match any modules in the manifest.
EOF
  exit 0
fi

# Build --module flags
module_flags=()
for mod in "${!matched_modules[@]}"; do
  module_flags+=(--module "$mod")
done

# Generate the markdown report
cem health -p "$PACKAGE_PATH" --format markdown "${module_flags[@]}" > health_report.md
