/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package tools_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/mcp/tools"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestElementStylingTool_Integration(t *testing.T) {
	// Test successful case with comprehensive output validation
	args := tools.ElementStylingArgs{
		TagName: "button-element",
		Context: "integration-test",
	}

	testElementStylingWithGolden(t, args, "integration_test.golden.md")
}

// Helper function for testing element styling with golden files
func testElementStylingWithGolden(t *testing.T, args tools.ElementStylingArgs, goldenFile string) {
	registry := getTestRegistry(t)

	argsJSON, err := json.Marshal(args)
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "element_styling",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeElementStylingHandler(registry)
	result, err := handler(context.Background(), req)
	require.NoError(t, err, "Tool should not return error")
	require.NotNil(t, result, "Tool should return result")
	require.Len(t, result.Content, 1, "Result should have content")

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok, "Content should be text")

	output := textContent.Text

	// Handle -update flag
	goldenPath := filepath.Join("../fixtures/element-styling-integration", goldenFile)
	if *update {
		err := os.MkdirAll(filepath.Dir(goldenPath), 0755)
		require.NoError(t, err, "Failed to create golden file directory")

		err = os.WriteFile(goldenPath, []byte(output), 0644)
		require.NoError(t, err, "Failed to update golden file %s", goldenPath)
		t.Logf("Updated golden file: %s", goldenPath)
		return
	}

	// Compare with golden file
	expectedData, err := os.ReadFile(goldenPath)
	if os.IsNotExist(err) {
		t.Fatalf("Golden file %s does not exist. Run with -update to create it.", goldenPath)
	}
	require.NoError(t, err, "Should be able to read golden file: %s", goldenPath)

	assert.Equal(t, string(expectedData), output, "Generated element styling should match golden file")
}