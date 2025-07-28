#!/bin/bash
set -euo pipefail

# Script to update embedded schemas from custom-elements-manifest package
# This should be run in CI/CD to keep schemas up to date

SCHEMAS_DIR="validate/schemas"
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

echo "Fetching available custom-elements-manifest versions..."

# Get all available versions from npm registry
VERSIONS=$(curl -s https://registry.npmjs.org/custom-elements-manifest | jq -r '.versions | keys[]' | sort -V)

echo "Available versions: $VERSIONS"

# Track if any new schemas were added
NEW_SCHEMAS=0

for version in $VERSIONS; do
    schema_file="$SCHEMAS_DIR/$version.json"
    
    # Skip if we already have this version (unless it's the speculative one)
    if [ -f "$schema_file" ] && [ "$version" != "2.1.1-speculative" ]; then
        echo "✓ Schema $version already exists"
        continue
    fi
    
    echo "Downloading schema for version $version..."
    
    # Download to temp directory first
    temp_file="$TEMP_DIR/$version.json"
    
    if curl -s "https://unpkg.com/custom-elements-manifest@$version/schema.json" -o "$temp_file"; then
        # Validate the downloaded file is valid JSON
        if jq empty "$temp_file" 2>/dev/null; then
            # Only copy if it's different or new
            if ! cmp -s "$temp_file" "$schema_file" 2>/dev/null; then
                cp "$temp_file" "$schema_file"
                echo "✓ Updated schema $version"
                NEW_SCHEMAS=$((NEW_SCHEMAS + 1))
            else
                echo "✓ Schema $version unchanged"
            fi
        else
            echo "✗ Downloaded schema $version is not valid JSON"
            exit 1
        fi
    else
        echo "✗ Failed to download schema for version $version"
        exit 1
    fi
done

echo "Schema update complete. $NEW_SCHEMAS new/updated schemas."

# If this is running in CI and we have new schemas, we could trigger a rebuild
if [ "$NEW_SCHEMAS" -gt 0 ] && [ "${CI:-}" = "true" ]; then
    echo "New schemas detected in CI environment"
    # Additional CI logic could go here
fi