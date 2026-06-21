package cmd

import (
	"fmt"

	"bennypowers.dev/cem/internal/tui"
)

func formatStatusLine(port int, reload bool, logLevel string) string {
	reloadColor := tui.StatusDisabledStyle.Render("false")
	if reload {
		reloadColor = tui.StatusEnabledStyle.Render("true")
	}
	sep := tui.StatusSepStyle.Render("|")
	return fmt.Sprintf("Running on %s %s Live reload: %s %s Log level: %s",
		tui.StatusURLStyle.Render(fmt.Sprintf("http://localhost:%d", port)),
		sep,
		reloadColor,
		sep,
		logLevel,
	)
}
