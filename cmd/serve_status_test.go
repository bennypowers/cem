package cmd

import (
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

func TestFormatStatusLine(t *testing.T) {
	tests := []struct {
		name   string
		port   int
		reload bool
	}{
		{name: "reload-enabled", port: 8000, reload: true},
		{name: "reload-disabled", port: 3000, reload: false},
		{name: "port-zero", port: 0, reload: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := formatStatusLine(tc.port, tc.reload)
			testutil.CheckGolden(t, tc.name, []byte(got), testutil.GoldenOptions{
				Dir: "testdata/goldens/serve-status",
			})
		})
	}
}
