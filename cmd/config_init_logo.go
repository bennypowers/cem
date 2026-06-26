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
	IC "bennypowers.dev/cem/internal/config"
	"charm.land/lipgloss/v2"
)

func selectASCIILogo(width int) string {
	switch {
	case width >= 100:
		return IC.LogoASCII100
	case width >= 60:
		return IC.LogoASCII60
	default:
		return IC.LogoASCII40
	}
}

func centerBlock(text string, width int) string {
	return lipgloss.PlaceHorizontal(width, lipgloss.Center, text)
}
