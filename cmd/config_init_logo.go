/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package cmd

import (
	"bytes"
	"image/png"
	"os"
	"strings"

	IC "bennypowers.dev/cem/internal/config"
	"github.com/dolmen-go/kittyimg"
	"charm.land/lipgloss/v2"
)

func isKittyTerminal() bool {
	return os.Getenv("KITTY_PID") != ""
}

func selectASCIILogo(width int) string {
	switch {
	case width < 50:
		return IC.LogoASCII40
	case width < 80:
		return IC.LogoASCII60
	default:
		return IC.LogoASCII100
	}
}

func renderLogo(width int) string {
	if isKittyTerminal() {
		if rendered := renderKittyLogo(); rendered != "" {
			return centerBlock(rendered, width)
		}
	}
	return centerBlock(strings.TrimRight(selectASCIILogo(width), "\n"), width)
}

func renderKittyLogo() string {
	img, err := png.Decode(bytes.NewReader(IC.LogoPNG))
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	if err := kittyimg.Fprint(&buf, img); err != nil {
		return ""
	}
	return buf.String()
}

func centerBlock(text string, width int) string {
	return lipgloss.PlaceHorizontal(width, lipgloss.Center, text)
}
