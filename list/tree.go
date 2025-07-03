/*
Copyright © 2025 Benny Powers

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
package list

import (
	"reflect"

	M "bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
)

func RenderTree(title string, renderable M.Renderable) error {
	if renderable == nil || isTypedNil(renderable) {
		pterm.Println("No deprecations found")
		return nil
	}
	root := renderable.ToTreeNode()
	pterm.DefaultSection.Print(title)
	return pterm.DefaultTree.WithRoot(root).Render()
}

func isTypedNil(i any) bool {
    // Use reflection to detect typed nil
    v := reflect.ValueOf(i)
    switch v.Kind() {
    case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
        return v.IsNil()
    }
    return false
}
